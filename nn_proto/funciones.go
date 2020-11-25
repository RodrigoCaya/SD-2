package nn_proto

import (
	"os"
	"fmt"
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
var ultimo int

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
		fmt.Println(ultimostring)
		if ultimostring == ""{
			break
		}
		ultimo, err = strconv.Atoi(ultimostring)
		if err != nil{
			log.Fatal(err)
		}
		fmt.Println(ultimo)
		for j := 0 ; j < (len(split) - 1) ; j++{
			nombres = nombres + split[j] + " "
			fmt.Println("xd")
		}
		fmt.Println(nombres)
		nombres = nombres +  "\n"
		for i := 0 ; i < ultimo ; i++{
			scanner.Scan()
			fmt.Println("xd")
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
	log.Printf("C1: %s", message.Cantidadn1)
	log.Printf("C2: %s", message.Cantidadn2)
	log.Printf("C3: %s", message.Cantidadn3)
	log.Printf("Cantidad: %s", message.Cantidadtotal)
	return &CodeRequest{Code: "xd"}, nil
}
