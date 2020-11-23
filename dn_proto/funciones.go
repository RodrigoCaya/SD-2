package dn_proto

import (
	"log"
	"golang.org/x/net/context"
)

type Server struct{
}

func distribuido(){
	log.Printf("algoritmo distribuido")
}

func centralizado(){
	log.Printf("algoritmo centralizado")
}

func (s *Server) EnviarChunks(ctx context.Context, message *ChunkRequest) (*CodeRequest, error) {
	if message.Tipo == "1"{
		distribuido()
	}else{
		centralizado()
	}
	return &CodeRequest{Code: "chunk recibido"}, nil
}
