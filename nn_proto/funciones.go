package nn_proto

import (
	"os"
	"fmt"
	"strconv"
	"io"
	"math"
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
var direcciones string

//Funcion que recorre el archivo log.txt y entrega las ubicaciones de los chunks 
//que tiene cada DataNode de un libro en especifico

func (s *Server) DisplayDirecciones(ctx context.Context, message *CodeRequest) (*Partes, error) {
	file, err := os.Open("log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	nombrelibro := message.Code
	scanner := bufio.NewScanner(file)
	partes1 := ""
	partes2 := ""
	partes3 := ""
	for err != io.EOF{
		scanner.Scan()
		split := strings.Split(scanner.Text(), " ")
		ultimostring := split[(len(split)) - 1]
		if ultimostring == ""{
			break
		}
		if nombrelibro == split[0]{
			cantidadtotal, err := strconv.Atoi(ultimostring)
			if err != nil{
				log.Fatal(err)
			}
			for i := 0 ; i < cantidadtotal ; i++{
				scanner.Scan()
				split1 := strings.Split(scanner.Text(), " ")
				direccion := split1[1]
				if direccion == "dist14:9001" {
					parte := strconv.Itoa(i)
					partes1 = partes1 + parte + ","
				}
				if direccion == "dist15:9002"{
					parte := strconv.Itoa(i)
					partes2 = partes2 + parte + ","
				}
				if direccion == "dist16:9003"{
					parte := strconv.Itoa(i)
					partes3 = partes3 + parte + ","
				}
			}
			break
		}
	}
	return &Partes{Partes1: partes1, Partes2: partes2, Partes3: partes3}, nil
}

//Funcion que lee el archivo log.txt ubicado en la carpeta NameNode 
//y despliega una lista con los libros disponibles para descargar

func (s *Server) DisplayLista(ctx context.Context, message *CodeRequest) (*CodeRequest, error) {
	file, err := os.Open("log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	cont := 0
	nombres = "\nEscoge libro para descargar" + "\n----------------- \n"
	
	for err != io.EOF {
		scanner.Scan()
		split := strings.Split(scanner.Text(), " ")
		ultimostring := split[(len(split)) - 1]
		if ultimostring == ""{
			break
		}
		ultimo, err = strconv.Atoi(ultimostring)
		if err != nil{
			log.Fatal(err)
		}
		cont = cont + 1
		contaux := strconv.Itoa(cont)
		for j := 0 ; j < (len(split) - 1) ; j++{
			nombres =  nombres + "(" + contaux + ")" + split[j] + "\n"
		}
		for i := 0 ; i < ultimo ; i++{
			scanner.Scan()
		}
	}
	nombres = nombres + "-----------------"
	return &CodeRequest{Code: nombres}, nil
}
	
var cantidad1 string
var cantidad2 string
var cantidad3 string
var i int
var j int
var k int

//Funcion global que llama a la funcion de agregar al log

func (s *Server) AgregarAlLog(ctx context.Context, message *Propuesta) (*CodeRequest, error) {
	agregarlog(message.Cantidadn1, message.Cantidadn2, message.Cantidadn3, message.Cantidadtotal, message.Nombrel)
	return &CodeRequest{Code: "Agregado al Log, revisalo"}, nil
}

//Funcion que agrega el libro al registro log.txt en el formato del enunciado
	
func agregarlog(c1 string, c2 string, c3 string, cantidadtotal string, nombrelibro string){
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	file.WriteString(nombrelibro + " " + cantidadtotal + "\n")
	if c1 != "0"{
		cantidad1, err := strconv.Atoi(c1)
		if err != nil {
			log.Fatal(err)
		}
		for i = 0 ; i < cantidad1 ; i++{
			file.WriteString(nombrelibro + "_" + strconv.Itoa(i) + " " + "dist14:9001\n")
		}
	}
	if c2 != "0"{
		cantidad2, err := strconv.Atoi(c2)
		if err != nil {
			log.Fatal(err)
		}
		for j = (0 + i) ; j < (cantidad2 + i) ; j++ {
			file.WriteString(nombrelibro + "_" + strconv.Itoa(j) + " " + "dist15:9002\n")
		}
	}

	if c3 != "0"{
		cantidad3, err := strconv.Atoi(c3)
		if err != nil {
			log.Fatal(err)
		}
		for k = (0 + j) ; k < (cantidad3 + j) ; k++ {
			file.WriteString(nombrelibro + "_" + strconv.Itoa(k) + " " + "dist16:9003\n")
			log.Printf("%d", k)
		}
	}
}

//Funcion que pinguea una maquina para saber si esta activa

func ualive(maquina string) string {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(maquina, grpc.WithInsecure())
	if err != nil {
		log.Printf("could not connect: %s", err)
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

//Funcion que recalcula la cantidad de chunks que recibe cada DataNode por si uno o varios se encuentran caidos

func recalcular(cantidad string, c1 string, c2 string, c3 string, nombrelibro string)(Propuesta){
	cant, err := strconv.Atoi(cantidad)
	if err != nil {
		log.Fatal(err)
	}
	chunksvivos := 3
	if c1 == "0" {
		chunksvivos = chunksvivos - 1
	}
	if c2 == "0" {
		chunksvivos = chunksvivos - 1
	}
	if c3 == "0" {
		chunksvivos = chunksvivos - 1
	}
	chunksxcadauno := cant/chunksvivos
	ch1 := chunksxcadauno
	ch2 := chunksxcadauno
	ch3 := chunksxcadauno
	if chunksvivos == 3 {
		if math.Mod(float64(cant), float64(chunksvivos)) == 1 {
			ch1 = ch1 + 1
		} else{
			if math.Mod(float64(cant), float64(chunksvivos)) == 2 {
				ch1 = ch1 + 1
				ch2 = ch2 + 1
			}
		}
	}else if chunksvivos == 2 {
		if c1 == "0" {
			ch1 = 0
			if math.Mod(float64(cant), float64(chunksvivos)) == 1 {
				ch2 = ch2 + 1
			}
		}
		if c2 == "0"{
			ch2 = 0
			if math.Mod(float64(cant), float64(chunksvivos)) == 1 {
				ch1 = ch1 + 1
			}
		}
		if c3 == "0"{
			ch3 = 0
			if math.Mod(float64(cant), float64(chunksvivos)) == 1 {
				ch1 = ch1 + 1
			}
		}
	}else if chunksvivos == 1 {
		if c1 != "0"{
			ch2 = 0
			ch3 = 0
		}
		if c2 != "0"{
			ch1 = 0
			ch3 = 0
		}
		if c3 != "0"{
			ch1 = 0
			ch2 = 0
		}
	}
	message := Propuesta{
		Cantidadn1: strconv.Itoa(ch1),
		Cantidadn2: strconv.Itoa(ch2),
		Cantidadn3: strconv.Itoa(ch3),
		Nombrel: nombrelibro,
		Cantidadtotal: cantidad,
	}
	return message
}

//Funcion que genera propuestas a partir de los DataNodes activos

func (s *Server) EnviarPropuesta(ctx context.Context, message *Propuesta) (*Propuesta, error) {
	log.Printf("Propuesta recibida")
	
	c1 := message.Cantidadn1
	c2 := message.Cantidadn2
	c3 := message.Cantidadn3
	flag := 1
	flag2 := 1
	if message.Cantidadn1 != "0"{
		resp := ualive("dist14:9001")
		if resp == "gg" {
			c1 = "0"
			flag = 0
		}
	}
	if message.Cantidadn2 != "0"{
		resp := ualive("dist15:9002")
		if resp == "gg" {
			c2 = "0"
			flag = 0
			flag2 = 0
		}
	}else if flag == 0{ //si hay solo 1 chunk
		resp := ualive("dist15:9002")
		if resp == "gg" {
			c2 = "0"
			flag = 0
		}else{
			c2 = "1"
		}
	}
	if message.Cantidadn3 != "0"{
		resp := ualive("dist16:9003")
		if resp == "gg" {
			c3 = "0"
			flag = 0
		}
	}else{
		if flag == 0{ //si hay solo 2 chunks
			c3 = "1"
			if flag2 == 0{
				c3 = "2"
			}
		}
		if flag2 == 0 && flag == 1{
			c3 = "1"
		} 
	}
	if flag == 1 {
		agregarlog(message.Cantidadn1, message.Cantidadn2, message.Cantidadn3, message.Cantidadtotal, message.Nombrel)
		return &Propuesta{Nombrel: "Propuesta aceptada"}, nil
	}
	mensaje := recalcular(message.Cantidadtotal,c1,c2,c3,message.Nombrel)
	agregarlog(mensaje.Cantidadn1, mensaje.Cantidadn2, mensaje.Cantidadn3, mensaje.Cantidadtotal, mensaje.Nombrel)
	return &mensaje, nil	
}
