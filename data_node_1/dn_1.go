package main 
import(
	"log"
	"net"
	"google.golang.org/grpc"
	"github.com/RodrigoCaya/SD-2/dn_proto"
)

func conexioncl(){
	liscliente, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("Failed to listen on port 9001: %v", err)
	}
	s := dn_proto.Server{}
	grpcServer := grpc.NewServer()
	dn_proto.RegisterHelloworldServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(liscliente); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9001: %v", err)
	}
}

func main(){
	conexioncl()
}
