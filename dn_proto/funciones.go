package dn_proto

import (
	"golang.org/x/net/context"
)

type Server struct{
}

func (s *Server) EnviarChunks(ctx context.Context, message *ChunkRequest) (*CodeRequest, error) {
	return &CodeRequest{Code: "chunk recibido"}, nil
}
