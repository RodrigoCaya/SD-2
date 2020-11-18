package main

import (
	"github.com/streadway/amqp"
	"golang.org/x/net/context"
)

func (s *Server) Buscar(ctx context.Context, message *CodeRequest) (*CodeRequest, error) {
	return &CodeRequest{Code: "xd"}, nil
}
