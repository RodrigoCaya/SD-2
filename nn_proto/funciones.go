package nn_proto

import (
	"os"
	"strconv"
	"io"
	"log"
	"bufio"
	"golang.org/x/net/context"
	"strings"
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

	for err != io.EOF {
		scanner.Scan()
		split := strings.Split(scanner.Text(), " ")
		ultimostring := split[(len(split)) - 1]
		ultimo, err := strconv.Atoi(ultimostring)
		if err != nil{
			log.Fatal(err)
		}
		for j := 0 ; j < ultimo - 1 ; j++{
			nombres = nombres + split[j]
		}
		nombres = nombres + "\n"
		for i := 0 ; i < ultimo ; i++{
			scanner.Scan()
		}
	} 
	
	/*for scanner.Scan() {
		fmt.Println(scanner.Text())
		nombres = nombres + "," + scanner.Text()
	}*/

	return &CodeRequest{Code: nombres}, nil
}

func (s *Server) EnviarPropuesta(ctx context.Context, message *Propuesta) (*CodeRequest, error) {
	log.Printf("Propuesta recibida")
	return &CodeRequest{Code: "xd"}, nil
}
