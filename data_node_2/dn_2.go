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
var id int = 2 //pal Ricardo

type Pagina struct{
	chunks []byte
	id_libro string
}

var libroactual []Pagina

type Server struct{
	dn_proto.UnimplementedDnServiceServer
}

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

func llamarRicardo(maquina string){
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(maquina, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := dn_proto.NewDnServiceClient(conn)

	message := dn_proto.RicRequest{ //agregue este msj, porqe el otro era tipo dn_proto
		Id: int32(id),
	}

	_, err = c.Ricardo(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling Ricardo: %s", err)
	}
}

//Funcion que realiza el algoritmo de Exclusion Mutua Distribuido, considerando las propuestas dependiendo de los data nodes activos

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
	message := dn_proto.PropRequest{ // lo cambie de nn_proto a dn_proto
		Cantidadn1: strconv.Itoa(c1),
		Cantidadn2: strconv.Itoa(c2),
		Cantidadn3: strconv.Itoa(c3),
		Nombrel: nombrelibro,
		Cantidadtotal: strconv.Itoa(cantidad),
	}
	maquina14 := "dist14:9001"
	maquina16 := "dist16:9003"
	respuesta1 := ""
	respuesta2 := ""
	for {
		flag := 0
		if c1 != 0 {
			respuesta1 = propuestadn(maquina14, message)
			if respuesta1 == "Rechazado" {
				flag = 1
			}
		}
		if c3 != 0 {
			respuesta2 = propuestadn(maquina16, message)
			if respuesta2 == "Rechazado" {
				flag = 1
			}
		}

		if flag == 0{
			estado = "WANTED" //pal Ricardo
			llamarRicardo("dist14:9001")
			llamarRicardo("dist16:9003")
			estado = "HELD"

			var conn *grpc.ClientConn
			conn, err := grpc.Dial("dist13:9000", grpc.WithInsecure())
			if err != nil {
				log.Fatalf("could not connect: %s", err)
			}
			defer conn.Close()
	
			c := nn_proto.NewHelloworldServiceClient(conn) // lo cambie de dn_proto a nn_proto
	
			messagenn := nn_proto.Propuesta{ //agregue este msj, porqe el otro era tipo dn_proto
				Cantidadn1: strconv.Itoa(c1),
				Cantidadn2: strconv.Itoa(c2),
				Cantidadn3: strconv.Itoa(c3),
				Nombrel: nombrelibro,
				Cantidadtotal: strconv.Itoa(cantidad),
			}
	
			response, err := c.AgregarAlLog(context.Background(), &messagenn)
			if err != nil {
				log.Fatalf("Error when calling AgregarAlLog: %s", err)
			}
			estado = "RELEASED"
			log.Printf("%s", response.Code)
			break
		}
		if respuesta1 == "Rechazado" && respuesta2 == "Aceptado" {
			chunksxcadauno = cantidad/2
			c1 = 0
			c2 = chunksxcadauno
			c3 = chunksxcadauno
			if math.Mod(float64(cantidad), 2) == 1 {
				c2 = c2 + 1
			}

		} else if respuesta2 == "Rechazado" && respuesta1 == "Aceptado" {
			chunksxcadauno = cantidad/2
			c1 = chunksxcadauno
			c2 = chunksxcadauno
			c3 = 0
			if math.Mod(float64(cantidad), 2) == 1 {
				c1 = c1 + 1
			}

		} else{
			c1 = 0
			c2 = cantidad
			c3 = 0
		}
		message = dn_proto.PropRequest{ // lo cambie de nn_proto a dn_proto
			Cantidadn1: strconv.Itoa(c1),
			Cantidadn2: strconv.Itoa(c2),
			Cantidadn3: strconv.Itoa(c3),
			Nombrel: nombrelibro,
			Cantidadtotal: strconv.Itoa(cantidad),
		}
		log.Printf("algoritmo distribuido")
	}
	if message.Cantidadn1 != "0"{
		maquina := "dist14:9001"
		conectardn(maquina, message)
	}
	if message.Cantidadn3 != "0"{
		maquina := "dist16:9003"
		conectardn(maquina, message)
	}
	descargarlocal(message)
}

//Funcion que manda la propuestas a los otros data nodes, los cuales aceptan o rechazan dependiendo de su estado, una vez reciba las respuestas de aceptacion, se acepta la propuesta


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
	}else{
		if response.Code == "Propuesta aceptada"{
			return respuesta
		}
	}
	respuesta = "Rechazado"
	return respuesta
	
}

func (s *Server) ChunksDN(ctx context.Context, message *dn_proto.ChunkRequest) (*dn_proto.CodeRequest, error) { //modificado
	log.Printf("me llegó la parte %s del libro %s",message.Parte, message.Nombrel)
	log.Printf("tamaño del chunk num %s es de %d", message.Parte, len(message.Chunk))
	// write to disk
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

	// write/save buffer to disk
	ioutil.WriteFile(fileName, message.Chunk, os.ModeAppend)
	
	return &dn_proto.CodeRequest{Code: "Recibido"}, nil
}

func (s *Server) PropuestasDN(ctx context.Context, message *dn_proto.PropRequest) (*dn_proto.CodeRequest, error) {
	log.Printf("Propuesta recibida")
	
	log.Printf("C1: %s", message.Cantidadn1)
	log.Printf("C2: %s", message.Cantidadn2)
	log.Printf("C3: %s", message.Cantidadn3)
	log.Printf("Cantidad: %s", message.Cantidadtotal)
	log.Printf("me llegó una propuesta de un dn")

	return &dn_proto.CodeRequest{Code: "Propuesta aceptada"}, nil
}

func descargarlocal(message dn_proto.PropRequest){ // debe ir despues de llamar a conectardn
	mensaje := dn_proto.ChunkRequest{}
	//modificado
	paldn1, err := strconv.Atoi(message.Cantidadn1) 
	if err != nil {
		log.Fatalf("Error convirtiendo: %s", err)
	}
	// part1 := "" 
	// contdn1 := 0 

	paldn2, err := strconv.Atoi(message.Cantidadn2) 
	if err != nil {
		log.Fatalf("Error convirtiendo: %s", err)
	}
	part2 := "" 
	contdn2 := 0 
	
	// paldn3, err := strconv.Atoi(message.Cantidadn3) 
	// if err != nil {
	// 	log.Fatalf("Error convirtiendo: %s", err)
	// }
	// part3 := "" 
	// contdn3 := 0 
	for{
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
		// write to disk
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
	
		// write/save buffer to disk
		ioutil.WriteFile(fileName, mensaje.Chunk, os.ModeAppend)
	}
	var librovacio []Pagina
	libroactual = librovacio
}

func conectardn(maquina string, message dn_proto.PropRequest){
	mensaje := dn_proto.ChunkRequest{}
	contdn1 := 0 //cambiar aca
	paldn2, err := strconv.Atoi(message.Cantidadn2) 
	if err != nil {
		log.Fatalf("Error convirtiendo: %s", err)
	}
	
	paldn1, err := strconv.Atoi(message.Cantidadn1) 
	if err != nil {
		log.Fatalf("Error convirtiendo: %s", err)
	}
	part1 := "" //cambiar aca

	contdn3 := 0 
	paldn3, err := strconv.Atoi(message.Cantidadn3) 
	if err != nil {
		log.Fatalf("Error convirtiendo: %s", err)
	}
	part3 := "" 
	
	for{
		if maquina == "dist14:9001"{ //cambiar aca
			if paldn1 != 0 && contdn1 < paldn1 { //cambiar aca
				aux := contdn1+1 //cambiar aca
				part1 = strconv.Itoa(aux) //cambiar aca
				mensaje = dn_proto.ChunkRequest{
					Chunk: libroactual[contdn1].chunks, //cambiar aca
					Parte: part1, //cambiar aca
					Cantidad: message.Cantidadtotal,
					Nombrel: message.Nombrel,
				}
				contdn1 = contdn1 + 1 //cambiar aca
			}else{
				break
			}
		}
		if maquina == "dist16:9003"{ //cambiar aca
			if paldn3 != 0 && contdn3 < paldn3 { //cambiar aca
				aux := paldn1+paldn2+contdn3+1 //cambiar aca
				part3 = strconv.Itoa(aux) //cambiar aca
				mensaje = dn_proto.ChunkRequest{
					Chunk: libroactual[paldn1+paldn2+contdn3].chunks, //cambiar aca
					Parte: part3, //cambiar aca
					Cantidad: message.Cantidadtotal,
					Nombrel: message.Nombrel,
				}
				contdn3 = contdn3 + 1 //cambiar aca
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
	
		response, err := c.ChunksDN(context.Background(), &mensaje)
		if err != nil {
			log.Fatalf("Error when calling ChunksDN: %s", err)
		}
	
		log.Printf("%s", response.Code)

	}

	//agregar la parte de dn2 para descargar y hacer libroactual = vacio
}

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
	log.Printf("algoritmo centralizado")
}

//aqui se conecta el cliente
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
	return &dn_proto.CodeRequest{Code: "chunk recibido"}, nil
}


func conexioncl(){
	liscliente, err := net.Listen("tcp", ":9002")
	if err != nil {
		log.Fatalf("Failed to listen on port 9002: %v", err)
	}
	// s := dn_proto.Server{}
	grpcServer := grpc.NewServer()
	dn_proto.RegisterDnServiceServer(grpcServer, &Server{})
	if err := grpcServer.Serve(liscliente); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9002: %v", err)
	}
}

func name_node(message nn_proto.Propuesta){

	var conn *grpc.ClientConn
	conn, err := grpc.Dial("dist13:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := nn_proto.NewHelloworldServiceClient(conn)

	response, err := c.EnviarPropuesta(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling EnviarPropuesta: %s", err)
	}
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
	//var prop string = "xd"
	if message.Cantidadn1 != "0"{
		maquina = "dist14:9001"
		conectardn(maquina, messagedn)
	}
	if message.Cantidadn3 != "0"{
		maquina = "dist16:9003"
		conectardn(maquina, messagedn)
	}
	descargarlocal(messagedn) //modificar aca
		
}

func (s *Server) Estado(ctx context.Context, message *dn_proto.CodeRequest) (*dn_proto.CodeRequest, error) {
	return &dn_proto.CodeRequest{Code: "Estoy vivo"}, nil
}

func (s *Server) PedirChunks(ctx context.Context, message *dn_proto.ChunkRequestDN) (*dn_proto.ChunkRequestDN, error) {
	parte := message.Partes

	nombrelibro := message.Nombrel
	chunkname := "chunks/" + nombrelibro + "_" + parte // change here!

	file, err := os.Open(chunkname)

	if err != nil {
			fmt.Println(err)
			os.Exit(1)
	}

	defer file.Close()

	fileInfo, _ := file.Stat()

	var fileSize int64 = fileInfo.Size()

	const fileChunk = 256000 // 250 kb, change this to your requirement

	partBuffer := make([]byte, fileSize)

	file.Read(partBuffer)

	log.Printf("tamaño del chunk num %s es de %d", parte, len(partBuffer))	
	return &dn_proto.ChunkRequestDN{Nombrel: nombrelibro, Partes: parte, Chunk: partBuffer,}, nil
}

func main(){
	//go name_node()
	conexioncl()
}
