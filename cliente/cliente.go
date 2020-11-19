package main

import (
	"log"	
	"google.golang.org/grpc"
	"context"
	"github.com/RodrigoCaya/SD-2/dn_proto"
	"github.com/RodrigoCaya/SD-2/nn_proto"
)

func data_node(){
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("dist14:9001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := dn_proto.NewHelloworldServiceClient(conn)
		
	message := dn_proto.CodeRequest{
		Code: "hola",
	}

	response, err := c.Buscar(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling Buscar: %s", err)
	}

	log.Printf("%s", response.Code)
}

func name_node(){
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("dist13:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := nn_proto.NewHelloworldServiceClient(conn)
		
	message := nn_proto.CodeRequest{
		Code: "hola",
	}

	response, err := c.Buscar(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling Buscar: %s", err)
	}

	log.Printf("%s", response.Code)
}

func main() {
	
	var first string

	for{
		fmt.Println("-----------------")
		fmt.Println("Escoge: ") 
		fmt.Println("(1) Subir libro") 
		fmt.Println("(2) Descargar libro")
		fmt.Println("(0) Salir")
		fmt.Println("-----------------")
		 	  
		fmt.Scanln(&first)
		if first == "1"{
			go data_node()
		}
		if first == "2"{
			go name_node()
			
		}
		if first == "0"{
			break
		}
	}
}
