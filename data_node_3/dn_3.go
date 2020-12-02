package main 
import(
	"log"
	"net"
	"context"
	"strconv"
	"math"
	"os"//agregao
	"fmt"//agregao
	"io/ioutil"//agregao
	"google.golang.org/grpc"
	"github.com/RodrigoCaya/SD-2/dn_proto"
	"github.com/RodrigoCaya/SD-2/nn_proto"
)

var estado string = "RELEASED" //pal Ricardo
var id int = 3 //pal Ricardo
var msgenviados int = 0

type Pagina struct{
	chunks []byte
	id_libro string
}

var libroactual []Pagina

type Server struct{
	dn_proto.UnimplementedDnServiceServer
}

//Funcion que utiliza el algoritmo de Ricart & Agrawala para resolver el conflicto 
//cuando más de 1 dn quiere escribir en el nn, haciendo que uno espere al otro dependiendo de la Id

func (s *Server) Ricardo(ctx context.Context, message *dn_proto.RicRequest) (*dn_proto.CodeRequest, error) {
	for{
		if estado == "RELEASED"{
			break
		}else if estado == "WANTED" && int(message.Id) > id{
			break
		}
	}
	return &dn_proto.CodeRequest{Code: "OK"}, nil
}

//Funcion que llama a la funcion de Ricart & Agrawala

func llamarRicardo(maquina string){
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(maquina, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := dn_proto.NewDnServiceClient(conn)

	message := dn_proto.RicRequest{
		Id: int32(id),
	}

	_, err = c.Ricardo(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling Ricardo: %s", err)
	}
	msgenviados = msgenviados + 1
}

//Funcion que realiza el algoritmo de Exclusion Mutua Distribuido, 
//considerando las propuestas dependiendo de los data nodes activos

func distribuido(cantidad int, nombrelibro string){
	chunksxcadauno := cantidad/3
	c1 := chunksxcadauno
	c2 := chunksxcadauno
	c3 := chunksxcadauno
	if math.Mod(float64(cantidad), 3) == 1 {
		c1 = c1 + 1
	} else{
		if math.Mod(float64(cantidad), 3) == 2 {
			c1 = c1 + 1
			c2 = c2 + 1
		}
	}
	message := dn_proto.PropRequest{ 
		Cantidadn1: strconv.Itoa(c1),
		Cantidadn2: strconv.Itoa(c2),
		Cantidadn3: strconv.Itoa(c3),
		Nombrel: nombrelibro,
		Cantidadtotal: strconv.Itoa(cantidad),
	}
	maquina15 := "dist15:9002"
	maquina14 := "dist14:9001"
	respuesta1 := ""
	respuesta2 := ""
	for {
		flag := 0
		if c2 != 0 {
			respuesta1 = propuestadn(maquina15, message)
			if respuesta1 == "Rechazado" {
				flag = 1
			}
		}
		if c1 != 0 {
			respuesta2 = propuestadn(maquina14, message)
			if respuesta2 == "Rechazado" {
				flag = 1
			}
		}

		if flag == 0{
			estado = "WANTED" //pal Ricardo
			llamarRicardo("dist15:9002")
			llamarRicardo("dist14:9001")
			estado = "HELD"

			var conn *grpc.ClientConn
			conn, err := grpc.Dial("dist13:9000", grpc.WithInsecure())
			if err != nil {
				log.Fatalf("could not connect: %s", err)
			}
			defer conn.Close()
	
			c := nn_proto.NewHelloworldServiceClient(conn) 
	
			messagenn := nn_proto.Propuesta{ 
				Cantidadn1: strconv.Itoa(c1),
				Cantidadn2: strconv.Itoa(c2),
				Cantidadn3: strconv.Itoa(c3),
				Nombrel: nombrelibro,
				Cantidadtotal: strconv.Itoa(cantidad),
			}
	
			_, err = c.AgregarAlLog(context.Background(), &messagenn)
			if err != nil {
				log.Fatalf("Error when calling AgregarAlLog: %s", err)
			}
			msgenviados = msgenviados + 1
			estado = "RELEASED"
			break
		}
		if respuesta1 == "Rechazado" && respuesta2 == "Aceptado" {
			chunksxcadauno = cantidad/2
			c1 = chunksxcadauno
			c2 = 0
			c3 = chunksxcadauno
			if math.Mod(float64(cantidad), 2) == 1 {
				c1 = c1 + 1
			}

		} else if respuesta2 == "Rechazado" && respuesta1 == "Aceptado" {
			chunksxcadauno = cantidad/2
			c1 = 0
			c2 = chunksxcadauno
			c3 = chunksxcadauno
			if math.Mod(float64(cantidad), 2) == 1 {
				c2 = c2 + 1
			}

		} else{
			c1 = 0
			c2 = 0
			c3 = cantidad
		}
		message = dn_proto.PropRequest{ 
			Cantidadn1: strconv.Itoa(c1),
			Cantidadn2: strconv.Itoa(c2),
			Cantidadn3: strconv.Itoa(c3),
			Nombrel: nombrelibro,
			Cantidadtotal: strconv.Itoa(cantidad),
		}
		log.Printf("algoritmo distribuido")
	}
	if message.Cantidadn2 != "0"{
		maquina := "dist15:9002" 
		conectardn(maquina, message)
	}
	if message.Cantidadn1 != "0"{
		maquina := "dist14:9001" 
		conectardn(maquina, message)
	}
	descargarlocal(message)
}

//Funcion que manda la propuestas a los otros data nodes para el algoritmo distribuido,
// los cuales aceptan o rechazan dependiendo de su estado (vivo o muerto),
// una vez reciba las respuestas de aceptacion, se acepta la propuesta

func propuestadn(maquina string, message dn_proto.PropRequest) string {
	respuesta := "Aceptado"
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(maquina, grpc.WithInsecure())
	if err != nil {
		log.Printf("could not connect: %s", err)
	}
	defer conn.Close()

	c := dn_proto.NewDnServiceClient(conn)

	response, err := c.PropuestasDN(context.Background(), &message)
	if err != nil {
		log.Printf("Se cayó la máquina: %s", maquina)
	}else {
		msgenviados = msgenviados + 1
		if response.Code == "Propuesta aceptada"{
			return respuesta
		}
	}
	respuesta = "Rechazado"
	return respuesta

}

//Funcion que recibe los chunks que le corresponden a este data node y lo escribe en su carpeta de registro "chunks"

func (s *Server) ChunksDN(ctx context.Context, message *dn_proto.ChunkRequest) (*dn_proto.CodeRequest, error) {
	parteaux, err := strconv.Atoi(message.Parte)
	if err != nil {
		log.Fatalf("Error convirtiendo: %s", err)
	}
	parteaux = parteaux - 1
	partee := strconv.Itoa(parteaux)
	fileName := "chunks/" + message.Nombrel + "_" + partee
	_, err = os.Create(fileName)

	if err != nil {
			fmt.Println(err)
			os.Exit(1)
	}
	ioutil.WriteFile(fileName, message.Chunk, os.ModeAppend)
	
	return &dn_proto.CodeRequest{Code: "Recibido"}, nil
}

//Funcion que recibe una propuesta y la acepta, si es que el nodo no esta caido

func (s *Server) PropuestasDN(ctx context.Context, message *dn_proto.PropRequest) (*dn_proto.CodeRequest, error) {
	log.Printf("Propuesta recibida")
	return &dn_proto.CodeRequest{Code: "Propuesta aceptada"}, nil
}

//Funcion que realiza la descarga las partes que le corresponden a este data node y las almacena localmente

func descargarlocal(message dn_proto.PropRequest){ 
	mensaje := dn_proto.ChunkRequest{}
	paldn1, err := strconv.Atoi(message.Cantidadn1) 
	if err != nil {
		log.Fatalf("Error convirtiendo: %s", err)
	}
	paldn2, err := strconv.Atoi(message.Cantidadn2) 
	if err != nil {
		log.Fatalf("Error convirtiendo: %s", err)
	}
	paldn3, err := strconv.Atoi(message.Cantidadn3) 
	if err != nil {
		log.Fatalf("Error convirtiendo: %s", err)
	}
	part3 := "" 
	contdn3 := 0 
	for{
		if paldn3 != 0 && contdn3 < paldn3 {
			aux := paldn1+paldn2+contdn3+1
			part3 = strconv.Itoa(aux)
			mensaje = dn_proto.ChunkRequest{
				Chunk: libroactual[paldn1+paldn2+contdn3].chunks,
				Parte: part3,
				Cantidad: message.Cantidadtotal,
				Nombrel: message.Nombrel,
			}
			contdn3 = contdn3 + 1
		}else{
			break
		}
		parteaux, err := strconv.Atoi(mensaje.Parte)
		if err != nil {
			log.Fatalf("Error convirtiendo: %s", err)
		}
		parteaux = parteaux - 1
		partee := strconv.Itoa(parteaux)
		fileName := "chunks/" + mensaje.Nombrel + "_" + partee
		_, err = os.Create(fileName)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		ioutil.WriteFile(fileName, mensaje.Chunk, os.ModeAppend)
	}
	var librovacio []Pagina
	libroactual = librovacio
}

//Funcion que se conecta con los otros data nodes para enviarle los chunks que les corresponden

func conectardn(maquina string, message dn_proto.PropRequest){
	mensaje := dn_proto.ChunkRequest{}
	paldn1, err := strconv.Atoi(message.Cantidadn1) 
	if err != nil {
		log.Fatalf("Error convirtiendo: %s", err)
	}
	part1 := "" 
	contdn1 := 0 

	paldn2, err := strconv.Atoi(message.Cantidadn2) 
	if err != nil {
		log.Fatalf("Error convirtiendo: %s", err)
	}
	part2 := "" 
	contdn2 := 0 
	
	for{
		if maquina == "dist14:9001"{ 
			if paldn1 != 0 && contdn1 < paldn1 { 
				aux := contdn1+1 
				part1 = strconv.Itoa(aux) 
				mensaje = dn_proto.ChunkRequest{
					Chunk: libroactual[contdn1].chunks, 
					Parte: part1, 
					Cantidad: message.Cantidadtotal,
					Nombrel: message.Nombrel,
				}
				contdn1 = contdn1 + 1 
			}else{
				break
			}
		}
		if maquina == "dist15:9002"{
			if paldn2 != 0 && contdn2 < paldn2 {
				aux := paldn1+contdn2+1
				part2 = strconv.Itoa(aux)
				mensaje = dn_proto.ChunkRequest{
					Chunk: libroactual[paldn1+contdn2].chunks,
					Parte: part2,
					Cantidad: message.Cantidadtotal,
					Nombrel: message.Nombrel,
				}
				contdn2 = contdn2 + 1
			}else{
				break
			}
		}
		var conn *grpc.ClientConn
		conn, err := grpc.Dial(maquina, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("could not connect: %s", err)
		}
		defer conn.Close()
	
		c := dn_proto.NewDnServiceClient(conn)
	
		_, err = c.ChunksDN(context.Background(), &mensaje)
		if err != nil {
			log.Fatalf("Error when calling ChunksDN: %s", err)
		}
		msgenviados = msgenviados + 1
	}
}

//Funcion que realiza el algoritmo de Exclusion Mutua Centralizado,
// el cual llama a que se conecte con el Name Node para que este genere la propuesta

func centralizado(cantidad int, nombrelibro string){
	chunksxcadauno := cantidad/3
	c1 := chunksxcadauno
	c2 := chunksxcadauno
	c3 := chunksxcadauno
	if math.Mod(float64(cantidad), 3) == 1 {
		c1 = c1 + 1
	} else{
		if math.Mod(float64(cantidad), 3) == 2 {
			c1 = c1 + 1
			c2 = c2 + 1
		}
	}
	message := nn_proto.Propuesta{
		Cantidadn1: strconv.Itoa(c1),
		Cantidadn2: strconv.Itoa(c2),
		Cantidadn3: strconv.Itoa(c3),
		Nombrel: nombrelibro,
		Cantidadtotal: strconv.Itoa(cantidad),
	}
	name_node(message)
}

//Funcion que recibe la peticion del algoritmo que el cliente desea implementar
// y guarda los chunks dentro de una lista llamada "libroactual"

func (s *Server) EnviarChunks(ctx context.Context, message *dn_proto.ChunkRequest) (*dn_proto.CodeRequest, error) {
	parte, err := strconv.Atoi(message.Parte)
	cantidad, err := strconv.Atoi(message.Cantidad)
	if err != nil {
		log.Fatalf("Error convirtiendo: %s", err)
	}

	pagina1 := Pagina{
		chunks: message.Chunk,
		id_libro: message.Nombrel,
	}

	libroactual = append(libroactual, pagina1)

	if cantidad == (parte + 1){
		
		if message.Tipo == "1"{
			distribuido(cantidad, message.Nombrel)
		}else{
			centralizado(cantidad, message.Nombrel)
		}
	}
	log.Printf("La cantidad de mensajes enviados es: %d", msgenviados)
	return &dn_proto.CodeRequest{Code: "chunk recibido"}, nil
}

//Funcion que permite al data node actuar como servidor

func conexioncl(){
	liscliente, err := net.Listen("tcp", ":9003")
	if err != nil {
		log.Fatalf("Failed to listen on port 9003: %v", err)
	}
	grpcServer := grpc.NewServer()
	dn_proto.RegisterDnServiceServer(grpcServer, &Server{})
	if err := grpcServer.Serve(liscliente); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9003: %v", err)
	}
}

//Funcion que se conecta al name node para enviar la propuesta base en el caso del algoritmo centralizado

func name_node(message nn_proto.Propuesta){
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("vt13:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := nn_proto.NewHelloworldServiceClient(conn)

	response, err := c.EnviarPropuesta(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling EnviarPropuesta: %s", err)
	}
	msgenviados = msgenviados + 1
	messagedn := dn_proto.PropRequest{
		Cantidadn1: response.Cantidadn1,
		Cantidadn2: response.Cantidadn2,
		Cantidadn3: response.Cantidadn3,
		Nombrel: response.Nombrel,
		Cantidadtotal: response.Cantidadtotal,
	}
	if response.Nombrel == "Propuesta aceptada" {
		messagedn = dn_proto.PropRequest{
			Cantidadn1: message.Cantidadn1,
			Cantidadn2: message.Cantidadn2,
			Cantidadn3: message.Cantidadn3,
			Nombrel: message.Nombrel,
			Cantidadtotal: message.Cantidadtotal,
		}
	}
	var maquina string = ""
	if message.Cantidadn1 != "0"{
		maquina = "dist14:9001"
		conectardn(maquina, messagedn)
	}
	if message.Cantidadn2 != "0"{
		maquina = "dist15:9002"
		conectardn(maquina, messagedn)
	}
	descargarlocal(messagedn)
}

//Funcion que recibe un ping para verificar si esta activo el DN

func (s *Server) Estado(ctx context.Context, message *dn_proto.CodeRequest) (*dn_proto.CodeRequest, error) {
	return &dn_proto.CodeRequest{Code: "Estoy vivo"}, nil
}

//Funcion que recibe y almacena los chunks que otro data node le envió

func (s *Server) PedirChunks(ctx context.Context, message *dn_proto.ChunkRequestDN) (*dn_proto.ChunkRequestDN, error) {
	parte := message.Partes
	nombrelibro := message.Nombrel
	chunkname := "chunks/" + nombrelibro + "_" + parte

	file, err := os.Open(chunkname)

	if err != nil {
			fmt.Println(err)
			os.Exit(1)
	}

	defer file.Close()
	fileInfo, _ := file.Stat()
	var fileSize int64 = fileInfo.Size()
	const fileChunk = 256000 // 250 kb
	partBuffer := make([]byte, fileSize)
	file.Read(partBuffer)
	
	return &dn_proto.ChunkRequestDN{Nombrel: nombrelibro, Partes: parte, Chunk: partBuffer,}, nil
}

func main(){
	conexioncl()
}
