package nn_proto

import (
	"os"
	"fmt"
	"log"
	"bufio"
	"golang.org/x/net/context"
)

type Server struct{
}


var nombres []string

func (s *Server) DisplayLista(ctx context.Context, message *CodeRequest) (*Lista, error) {
	file, err := os.Open("log.txt")
	if err != nil {
		log.fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		nombres = append(nombres, scanner.Text())
	}

	return &Lista{L: nombres}, nil
}

func (s *Server) EnviarPropuesta(ctx context.Context, message *Propuesta) (*CodeRequest, error) {
	log.Printf("Propuesta recibida")
	return &CodeRequest{Code: "xd"}, nil
}