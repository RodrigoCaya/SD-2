package main 
import(
	"log"
	"net"
	"google.golang.org/grpc"
	"github.com/RodrigoCaya/SD-2/nn_proto"
)

//Funcion que permite al NameNode actuar como servidor

func conexioncl(){
	liscliente, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}
	s := nn_proto.Server{}
	grpcServer := grpc.NewServer()
	nn_proto.RegisterHelloworldServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(liscliente); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}
}

func main(){
	conexioncl()
}
