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

var id int = 0

type Pagina struct{
	chunks []byte
	id_libro int
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

func distribuido(){
	/*var maquina string = ""
	var prop string = "xd"
	maquina = "dist15:9002"
	conectardn(maquina, prop)
	maquina = "dist16:9003"
	conectardn(maquina, prop)*/
	log.Printf("algoritmo distribuido")
}

func (s *Server) Propuesta(ctx context.Context, message *dn_proto.PropRequest) (*dn_proto.CodeRequest, error) {
	log.Printf("me llegó una propuesta de un dn")
	return &dn_proto.CodeRequest{Code: "Recibido"}, nil
}

func conectardn(maquina string, message dn_proto.PropRequest){
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(maquina, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := dn_proto.NewDnServiceClient(conn)

	response, err := c.Propuesta(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling Buscar: %s", err)
	}

	log.Printf("%s", response.Code)
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

func (s *Server) EnviarChunks(ctx context.Context, message *dn_proto.ChunkRequest) (*dn_proto.CodeRequest, error) {
	
	parte, err := strconv.Atoi(message.Parte)
	cantidad, err := strconv.Atoi(message.Cantidad)
	if err != nil {
		log.Fatalf("Error convirtiendo: %s", err)
	}

	pagina1 := Pagina{
		chunks: message.Chunk,
		id_libro: id,
	}

	libroactual = append(libroactual, pagina1)

	if cantidad == (parte + 1){
		id = id + 1
		if message.Tipo == "1"{
			distribuido()
		}else{
			centralizado(cantidad, message.Nombrel)
		}
	}
	return &dn_proto.CodeRequest{Code: "chunk recibido"}, nil
}


func conexioncl(){
	liscliente, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("Failed to listen on port 9001: %v", err)
	}
	// s := dn_proto.Server{}
	grpcServer := grpc.NewServer()
	dn_proto.RegisterDnServiceServer(grpcServer, &Server{})
	if err := grpcServer.Serve(liscliente); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9001: %v", err)
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
			if message.Cantidadn2 != "0"{
				maquina = "dist15:9002"
				conectardn(maquina, messagedn)
			}
			if message.Cantidadn3 != "0"{
				maquina = "dist16:9003"
				conectardn(maquina, messagedn)
			}
			break
		}else{
			if response.Code == "dn2"{
				chunk2 = "0"
			}else{
				if response.Code == "dn3"{
					chunk3 = "0"
				}else{
					chunk2 = "0"
					chunk3 = "0"
				}
			}
			message := nn_proto.Propuesta{
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
