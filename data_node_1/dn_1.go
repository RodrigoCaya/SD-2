package main 
import(
	"log"
	"net"
	"context"
	"strconv"
	"google.golang.org/grpc"
	"github.com/RodrigoCaya/SD-2/dn_proto"
	"github.com/RodrigoCaya/SD-2/nn_proto"
)

var id int = 0

type Libro struct{
	chunks []byte
	id_libro string
}

var libros []Libro

type Server struct{
	dn_proto.UnimplementedDnServiceServer
}

func distribuido(){
	var maquina string = ""
	var prop string = "xd"
	maquina = "dist15:9002"
	conectardn(maquina, prop)
	maquina = "dist16:9003"
	conectardn(maquina, prop)
	log.Printf("algoritmo distribuido")
}

func (s *Server) Propuesta(ctx context.Context, message *dn_proto.PropRequest) (*dn_proto.CodeRequest, error) {
	log.Printf("hola")
	return &dn_proto.CodeRequest{Code: "equis de"}, nil
}

func conectardn(maquina string, prop string){
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(maquina, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := dn_proto.NewDnServiceClient(conn)
		
	message := dn_proto.PropRequest{
		Propuesta: prop,
	}

	response, err := c.Propuesta(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling Buscar: %s", err)
	}

	log.Printf("%s", response.Code)
}

func centralizado(machine string){
	name_node()
	log.Printf("algoritmo centralizado")
}

func (s *Server) EnviarChunks(ctx context.Context, message *dn_proto.ChunkRequest) (*dn_proto.CodeRequest, error) {
	//si es el ultimo chunk
	parte, err := strconv.Atoi(message.Parte)
	cantidad, err := strconv.Atoi(message.Cantidad)
	if err != nil {
		log.Fatalf("Error convirtiendo: %s", err)
	}
	if cantidad == (parte + 1){
		if message.Tipo == "1"{
			distribuido()
		}else{
			centralizado(message.Machine)
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

func name_node(){
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("dist13:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := nn_proto.NewHelloworldServiceClient(conn)
		
	message := nn_proto.Propuesta{
		Cantidadn1: "1",
		Cantidadn2: "2",
		Cantidadn3: "2",
		Cantidadtotal: "5",
	}

	response, err := c.EnviarPropuesta(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling Buscar: %s", err)
	}

	log.Printf("%s", response.Code)
}

func main(){
	//go name_node()
	conexioncl()
}
