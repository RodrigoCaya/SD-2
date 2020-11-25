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


var nombres string

func (s *Server) DisplayLista(ctx context.Context, message *CodeRequest) (*CodeRequest, error) {
	file, err := os.Open("log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		nombres = nombres + "," + scanner.Text()
	}

	return &CodeRequest{Code: nombres}, nil
}

func (s *Server) EnviarPropuesta(ctx context.Context, message *Propuesta) (*CodeRequest, error) {
	log.Printf("Propuesta recibida")
	return &CodeRequest{Code: "xd"}, nil
}
