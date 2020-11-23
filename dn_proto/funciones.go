package dn_proto

import (
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/RodrigoCaya/SD-2/dn_proto"
)

type Server struct{
}

func distribuido(){
	log.Printf("algoritmo distribuido")
}

func (s *Server) Propuesta(ctx context.Context, message *PropRequest) (*CodeRequest, error) {
	return &CodeRequest{Code: "equis de"}, nil
}

func conectar(maquina string, prop string){
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
	var maquina string = ""
	var prop string = ""
	if machine == "4" {
		maquina = "dist15:9002"
		conectar(maquina, prop)
		maquina = "dist16:9003"
		conectar(maquina, prop)

	}else{
		if machine == "5" {
			maquina = "dist14:9001"
			conectar(maquina, prop)
			maquina = "dist16:9003"
			conectar(maquina, prop)

		}else{
			maquina = "dist14:9001"
			conectar(maquina, prop)
			maquina = "dist15:9002"
			conectar(maquina, prop)

		}
	}
	log.Printf("algoritmo centralizado")
}

func (s *Server) EnviarChunks(ctx context.Context, message *ChunkRequest) (*CodeRequest, error) {
	//si es el ultimo chunk
	if message.Tipo == "1"{
		distribuido()
	}else{
		centralizado(message.Machine)
	}
	return &CodeRequest{Code: "chunk recibido"}, nil
}
