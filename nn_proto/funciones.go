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
	cantidad1 = strconv.Atoi(c1)
	for i = 0 ; i < cantidad1 ; i++{
		file.WriteString(nombrelibro + "_" + strconv.Itoa(i) + " " + "dist14:9001\n")
	}
}
if c2 != ""{
	cantidad2 = strconv.Atoi(c2)
	for j = (0 + i) ; j < (cantidad2 + i) ; j++ {
		file.WriteString(nombrelibro + "_" + strconv.Itoa(j) + " " + "dist15:9002\n")
	}
}

if c3 != ""{
	cantidad3 = strconv.Atoi(c3)
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
	return &CodeRequest{Code: "Propuesta aceptada"}, nil
}
