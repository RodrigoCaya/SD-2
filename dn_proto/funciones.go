package dn_proto

import (
	"golang.org/x/net/context"
)

type Server struct{
}

func (s *Server) Buscar(ctx context.Context, message *CodeRequest) (*CodeRequest, error) {
	return &CodeRequest{Code: "xd 2"}, nil
}
