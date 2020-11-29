package main

import (
	"log"
	"fmt"	
	"google.golang.org/grpc"
	"context"
	//"bufio"
	//"io/ioutil"
	"strings"
	"math"
	"math/rand"
	"os"
	"strconv"
	"github.com/RodrigoCaya/SD-2/dn_proto"
	"github.com/RodrigoCaya/SD-2/nn_proto"
)

func data_node(chunk_libro []byte, algoritmo string, probabilidad int, part int, total int, nombrelibro string){
	var conn *grpc.ClientConn
	maquina := strconv.Itoa(probabilidad+4)
	puerto := strconv.Itoa(probabilidad+1)
	conexion:= "dist1"
	conexion = conexion + maquina + ":900" + puerto
	conn, err := grpc.Dial(conexion, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := dn_proto.NewDnServiceClient(conn)
		
	message := dn_proto.ChunkRequest{
		Chunk: chunk_libro,
		Tipo: algoritmo,
		Parte: strconv.Itoa(part),
		Cantidad: strconv.Itoa(total),
		Machine: maquina,
		Nombrel: nombrelibro,
	}

	response, err := c.EnviarChunks(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling Buscar: %s", err)
	}

	log.Printf("%s", response.Code)
}

func separarlibro(algoritmo string){
	nombrelibro := "Peter_Pan-J._M._Barrie"
	fileToBeChunked := "../libros_cliente/Peter_Pan-J._M._Barrie.pdf" // change here!

	file, err := os.Open(fileToBeChunked)

	if err != nil {
			fmt.Println(err)
			os.Exit(1)
	}

	defer file.Close()

	fileInfo, _ := file.Stat()

	var fileSize int64 = fileInfo.Size()

	const fileChunk = 256000 // 250 kb, change this to your requirement

	// calculate total number of parts the file will be chunked into

	totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))

	
	//fmt.Printf("Splitting to %d pieces.\n", totalPartsNum)

	probabilidad := rand.Intn(3)
	for i := uint64(0); i < totalPartsNum; i++ {

			partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
			partBuffer := make([]byte, partSize)

			file.Read(partBuffer)

			// write to disk
			/*fileName := "Dracula-Stoker_Bram_" + strconv.FormatUint(i, 10)
			_, err := os.Create(fileName)

			if err != nil {
					fmt.Println(err)
					os.Exit(1)
			}

			// write/save buffer to disk
			ioutil.WriteFile(fileName, partBuffer, os.ModeAppend)

			fmt.Println("Split to : ", fileName)*/

			data_node(partBuffer, algoritmo, probabilidad,int(i) , int(totalPartsNum), nombrelibro)
	}
	/*
	// just for fun, let's recombine back the chunked files in a new file

	newFileName := "NEWbigfile.pdf"
	_, err = os.Create(newFileName)

	if err != nil {
			fmt.Println(err)
			os.Exit(1)
	}

	//set the newFileName file to APPEND MODE!!
	// open files r and w

	file, err = os.OpenFile(newFileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err != nil {
			fmt.Println(err)
			os.Exit(1)
	}

	// IMPORTANT! do not defer a file.Close when opening a file for APPEND mode!
	// defer file.Close()

	// just information on which part of the new file we are appending
	var writePosition int64 = 0

	for j := uint64(0); j < totalPartsNum; j++ {

			//read a chunk
			currentChunkFileName := "Dracula-Stoker_Bram_" + strconv.FormatUint(j, 10)

			newFileChunk, err := os.Open(currentChunkFileName)

			if err != nil {
					fmt.Println(err)
					os.Exit(1)
			}

			defer newFileChunk.Close()

			chunkInfo, err := newFileChunk.Stat()

			if err != nil {
					fmt.Println(err)
					os.Exit(1)
			}

			// calculate the bytes size of each chunk
			// we are not going to rely on previous data and constant

			var chunkSize int64 = chunkInfo.Size()
			chunkBufferBytes := make([]byte, chunkSize)

			fmt.Println("Appending at position : [", writePosition, "] bytes")
			writePosition = writePosition + chunkSize

			// read into chunkBufferBytes
			reader := bufio.NewReader(newFileChunk)
			_, err = reader.Read(chunkBufferBytes)

			if err != nil {
					fmt.Println(err)
					os.Exit(1)
			}

			// DON't USE ioutil.WriteFile -- it will overwrite the previous bytes!
			// write/save buffer to disk
			//ioutil.WriteFile(newFileName, chunkBufferBytes, os.ModeAppend)

			n, err := file.Write(chunkBufferBytes)

			if err != nil {
					fmt.Println(err)
					os.Exit(1)
			}

			file.Sync() //flush to disk

			// free up the buffer for next cycle
			// should not be a problem if the chunk size is small, but
			// can be resource hogging if the chunk size is huge.
			// also a good practice to clean up your own plate after eating

			chunkBufferBytes = nil // reset or empty our buffer

			fmt.Println("Written ", n, " bytes")

			fmt.Println("Recombining part [", j, "] into : ", newFileName)
	}*/

	// now, we close the newFileName
	file.Close()
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

	response, err := c.DisplayLista(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling Buscar: %s", err)
	}

	log.Printf("%s", response.Code)
	split := strings.Split(response.Code, "\n")
	var first string
	fmt.Scanln(&first)
	auxfirst, err := strconv.Atoi(first)
	if err != nil {
		log.Fatalf("Error when calling Buscar: %s", err)
	}
	nombredellibro := split[auxfirst+1]
	paratrim := "("+first+")"
	nombrefinal := strings.Trim(nombredellibro, paratrim)

	mensaje := nn_proto.CodeRequest{
		Code: nombrefinal
	}

	respuesta, err := c.DisplayDirecciones(context.Background(), &mensaje)
	if err != nil {
		log.Fatalf("Error when calling DisplayDirecciones: %s", err)
	}
	mensajedirecciones := nn_proto.Partes{
		Partes1: mensaje.Partes1,
		Partes2: mensaje.Partes2,
		Partes3: mensaje.Partes3,
	}
	
	//hacer la funcion del nn para qe le pase las direcciones (jean) (listoko)
	//recibir cual libro

	//hacer la funcion del dn para qe le envien los chunks
	//recibir cual libro y sus partes

	//hacer la funcion del cliente para qe descargue los chunks
	//hacer la funcion del cliente para unir los chunks
	//borrar los chunks
}

func recibirdirecciones(nombrelibro string){

}

func main() {
	
	var first string
	var second string

	for{
		fmt.Println("-----------------")
		fmt.Println("Escoge: ") 
		fmt.Println("(1) Subir libro") 
		fmt.Println("(2) Descargar libro")
		fmt.Println("(0) Salir")
		fmt.Println("-----------------")
		 	  
		fmt.Scanln(&first)
		if first == "1"{
			fmt.Println("-----------------")
			fmt.Println("Escoge: ") 
			fmt.Println("(1) Algoritmo Distribuido") 
			fmt.Println("(2) Algoritmo Centralizado")
			fmt.Println("(0) Salir")
			fmt.Println("-----------------")
			fmt.Scanln(&second)
			if second != "0"{
				separarlibro(second)
			}
		}
		if first == "2"{
			name_node()
			
		}
		if first == "0"{
			break
		}
	}
}
