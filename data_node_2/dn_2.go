package main 
import(
	"log"
	"net"
	"context"
	"strconv"
	"math"
	"google.golang.org/grpc"
	"github.com/RodrigoCaya/SD-2/dn_proto"
	"github.com/RodrigoCaya/SD-2/nn_proto"
)



type Pagina struct{
	chunks []byte
	id_libro string
}

var libroactual []Pagina

type Book struct{
	books []Pagina
}
//aqui va solo lo qe se va a almacenar al final
var biblioteca []Book

type Server struct{
	dn_proto.UnimplementedDnServiceServer
}

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
	maquina15 := "dist15:9002"
	maquina16 := "dist16:9003"
	respuesta1 := propuestadn(maquina15, message)
	respuesta2 := propuestadn(maquina16, message)
	if respuesta1 == "Aceptado" && respuesta2 == "Aceptado" {
		var conn *grpc.ClientConn
		conn, err := grpc.Dial("dist13:9000", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("could not connect: %s", err)
		}
		defer conn.Close()

		c := nn_proto.NewHelloworldServiceClient(conn) // lo cambie de dn_proto a nn_proto

		message := nn_proto.Propuesta{ //agregue este msj, porqe el otro era tipo dn_proto
			Cantidadn1: strconv.Itoa(c1),
			Cantidadn2: strconv.Itoa(c2),
			Cantidadn3: strconv.Itoa(c3),
			Nombrel: nombrelibro,
			Cantidadtotal: strconv.Itoa(cantidad),
		}

		response, err := c.AgregarAlLog(context.Background(), &message)
		if err != nil {
			log.Fatalf("Error when calling Buscar: %s", err)
		}
		log.Printf("%s", response.Code)
	} else if respuesta1 == "Rechazado" {
		
	} else {
		
	}
	log.Printf("algoritmo distribuido")
}

func (s *Server) ChunksDN(ctx context.Context, message *dn_proto.ChunkRequest) (*dn_proto.CodeRequest, error) {
	log.Printf("me llegó la parte %s del libro %s",message.Parte, message.Nombrel)
	//descargar los chunks
	return &dn_proto.CodeRequest{Code: "Recibido"}, nil
}

func (s *Server) PropuestasDN(ctx context.Context, message *dn_proto.PropRequest) (*dn_proto.CodeRequest, error) {
	log.Printf("Propuesta recibida")
	
	log.Printf("C1: %s", message.Cantidadn1)
	log.Printf("C2: %s", message.Cantidadn2)
	log.Printf("C3: %s", message.Cantidadn3)
	log.Printf("Cantidad: %s", message.Cantidadtotal)
	log.Printf("me llegó una propuesta de un dn")
	return &dn_proto.CodeRequest{Code: "Recibido"}, nil
}

func ualive(maquina string) string {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(maquina, grpc.WithInsecure())
	if err != nil {
		log.Printf("could not connect: %s", err)
	}
	defer conn.Close()

	c := dn_proto.NewDnServiceClient(conn)

	message := dn_proto.CodeRequest{
		Code: "¿Estas vivo?",
	}
	response, err := c.Estado(context.Background(), &message)
	respuesta := ""
	if err != nil {
		log.Printf("Se cayó el %s", maquina)
		respuesta = "gg"
	}else{
		respuesta = response.Code
	}
	return respuesta
}

func propuestadn(maquina string, message dn_proto.PropRequest) string {
	respuesta := "Aceptado"
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(maquina, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := dn_proto.NewDnServiceClient(conn)

	response, err := c.PropuestasDN(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling Buscar: %s", err)
	}

	log.Printf("%s", response.Code)
	if response.Code == "Propuesta aceptada"{
		return respuesta
	}else{
		respuesta = "Rechazado"
		return respuesta
	}
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
			log.Fatalf("Error when calling Buscar: %s", err)
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
	chunk1 := message.Cantidadn1
	chunk2 := message.Cantidadn2
	chunk3 := message.Cantidadn3
	nombre := message.Nombrel
	cantidadtotal := message.Cantidadtotal
	for{
		var conn *grpc.ClientConn
		conn, err := grpc.Dial("dist13:9000", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("could not connect: %s", err)
		}
		defer conn.Close()
	
		c := nn_proto.NewHelloworldServiceClient(conn)
	
		response, err := c.EnviarPropuesta(context.Background(), &message)
		if err != nil {
			log.Fatalf("Error when calling Buscar: %s", err)
		}
		if response.Code == "Propuesta aceptada" {
			messagedn := dn_proto.PropRequest{
				Cantidadn1: message.Cantidadn1,
				Cantidadn2: message.Cantidadn2,
				Cantidadn3: message.Cantidadn3,
				Nombrel: message.Nombrel,
				Cantidadtotal: message.Cantidadtotal,
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
			break
		}else{
			if response.Code == "dn1"{
				chunk1 = "0"
			}else{
				if response.Code == "dn3"{
					chunk3 = "0"
				}else{
					chunk1 = "0"
					chunk3 = "0"
				}
			}
			message = nn_proto.Propuesta{
				Cantidadn1: chunk1,
				Cantidadn2: chunk2,
				Cantidadn3: chunk3,
				Nombrel: nombre,
				Cantidadtotal: cantidadtotal,
			}
		}
		log.Printf("%s", response.Code)
	}
}

func (s *Server) Estado(ctx context.Context, message *dn_proto.CodeRequest) (*dn_proto.CodeRequest, error) {
	return &dn_proto.CodeRequest{Code: "Estoy vivo"}, nil
}

func main(){
	//go name_node()
	conexioncl()
}
