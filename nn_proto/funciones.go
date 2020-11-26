package nn_proto

import (
	"os"
	"fmt"
	"strconv"
	"io"
	"log"
	"bufio"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	"github.com/RodrigoCaya/SD-2/dn_proto"
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
		return &CodeRequest{Code: nombres}, nil
	}
	
var cantidad1 string
var cantidad2 string
var cantidad3 string
var i int
var j int
var k int

	
func agregarlog(c1 string, c2 string, c3 string, cantidadtotal string, nombrelibro string){
	file, err := os.Create("log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	file.WriteString(nombrelibro + " " + cantidadtotal + "\n")
	if c1 != ""{
		cantidad1, err := strconv.Atoi(c1)
		if err != nil {
			log.Fatal(err)
		}
		for i = 0 ; i < cantidad1 ; i++{
			file.WriteString(nombrelibro + "_" + strconv.Itoa(i) + " " + "dist14:9001\n")
		}
	}
	if c2 != ""{
		cantidad2, err := strconv.Atoi(c2)
		if err != nil {
			log.Fatal(err)
		}
		for j = (0 + i) ; j < (cantidad2 + i) ; j++ {
			file.WriteString(nombrelibro + "_" + strconv.Itoa(j) + " " + "dist15:9002\n")
		}
	}

	if c3 != ""{
		cantidad3, err := strconv.Atoi(c3)
		if err != nil {
			log.Fatal(err)
		}
		for k = (0 + i + j) ; k < (cantidad3 + i + j) ; k++ {
			file.WriteString(nombrelibro + "_" + strconv.Itoa(k) + " " + "dist16:9003\n")
		}
	}

//cree el log a partir de los c1, c2, c3
//c1 tiene los primeros, despues el c2, despues el c3
//IP dn1 = dist14:9001
//IP dn2 = dist15:9002
//IP dn3 = dist16:9003
}

func ualive(maquina string) string {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(maquina, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := dn_proto.NewDnServiceClient(conn)

	message := dn_proto.CodeRequest{
		Code: "¿Estas vivo?",
	}
	response, err := c.Estado(context.Background(), &message)
	respuesta := ""
	if err != nil {
		log.Printf("Se cayó el %s", maquina)
		respuesta = "gg"
	}else{
		respuesta = response.Code
	}
	return respuesta
}

func (s *Server) EnviarPropuesta(ctx context.Context, message *Propuesta) (*CodeRequest, error) {
	log.Printf("Propuesta recibida")
	
	log.Printf("C1: %s", message.Cantidadn1)
	log.Printf("C2: %s", message.Cantidadn2)
	log.Printf("C3: %s", message.Cantidadn3)
	log.Printf("Cantidad: %s", message.Cantidadtotal)
	agregarlog(message.Cantidadn1, message.Cantidadn2, message.Cantidadn3, message.Cantidadtotal, message.Nombrel)
	//revisar qe los dn involucrados esten activos
	//si estan activos, entonces hacer el log y responder "propuesta aceptada"
	//si no estan activos, no hacer el log y responder "se cayo un wn"
	respuesta := ""
	flag := 1
	resp := ualive("dist14:9001")
	if resp == "gg" {
		respuesta = respuesta + "dn1"
		flag = 0
	}
	resp = ualive("dist15:9002")
	if resp == "gg" {
		respuesta = respuesta + "dn2"
		flag = 0
	}
	resp = ualive("dist16:9003")
	if resp == "gg" {
		respuesta = respuesta + "dn3"
		flag = 0
	}

	if flag == 1 {
		return &CodeRequest{Code: "Propuesta aceptada"}, nil
	}
	return &CodeRequest{Code: respuesta}, nil	
}
