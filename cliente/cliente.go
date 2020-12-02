package main

import (
	"log"
	"fmt"	
	"google.golang.org/grpc"
	"context"
	"bufio"
	"io/ioutil"
	"strings"
	"math"
	"math/rand"
	"os"
	"strconv"
	"github.com/RodrigoCaya/SD-2/dn_proto"
	"github.com/RodrigoCaya/SD-2/nn_proto"
)

//Funcion que se conecta a un DataNode aleatorio

func data_node(chunk_libro []byte, algoritmo string, probabilidad int, part int, total int, nombrelibro string)int{
	var conn *grpc.ClientConn
	maquina := strconv.Itoa(probabilidad+4)
	puerto := strconv.Itoa(probabilidad+1)
	conexion:= "dist1"
	conexion = conexion + maquina + ":900" + puerto
	conn, err := grpc.Dial(conexion, grpc.WithInsecure())
	if err != nil {
		log.Printf("no se pudo conectar al dn %s", maquina)
		return 0
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
	log.Printf("tam chunk: %d", len(message.Chunk))
	log.Printf("tipo: %s", message.Tipo)
	log.Printf("parte: %s", message.Parte)
	log.Printf("cantidad: %s", message.Cantidad)
	log.Printf("maquina: %s", message.Machine)
	log.Printf("nombre libro: %s", message.Nombrel)

	response, err := c.EnviarChunks(context.Background(), &message)
	if err != nil {
		log.Printf("no se pudo conectar al dn %s", maquina)
		return 0
	}

	log.Printf("%s", response.Code)
	return 1
}

//Funcion que separa el libro indicado en chunks de 250kB

func separarlibro(algoritmo string, librosinpdf string, libroconpdf string){
	log.Printf("sin pdf %s con pdf %s ", librosinpdf,libroconpdf)
	nombrelibro := librosinpdf
	fileToBeChunked := "../libros_cliente/" + libroconpdf// change here!

	file, err := os.Open(fileToBeChunked)

	if err != nil {
			fmt.Println(err)
			os.Exit(1)
	}

	defer file.Close()

	fileInfo, _ := file.Stat()
	var fileSize int64 = fileInfo.Size()
	const fileChunk = 256000 // 250 kb, change this to your requirement
	totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))
	vivo := 1
	probabilidad := rand.Intn(3)
	var probabilidad2 int
	var probabilidad3 int
	for i := uint64(0); i < totalPartsNum; i++ {

		partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
		partBuffer := make([]byte, partSize)

		file.Read(partBuffer)
		log.Printf("CANTIDAD %d", int(totalPartsNum))
		vivo = data_node(partBuffer, algoritmo, probabilidad,int(i) , int(totalPartsNum), nombrelibro)
		if vivo == 0 {
			probabilidad2 = rand.Intn(2)
			if probabilidad == 0 {
				probabilidad2 = probabilidad2 + 1
			}
			if probabilidad == 1 {
				if probabilidad2 == 1 {
					probabilidad2 = 2
				}
			}
			vivo = data_node(partBuffer, algoritmo, probabilidad2,int(i) , int(totalPartsNum), nombrelibro)
		}
		if vivo == 0 {
			probabilidad3 = 0
			if (probabilidad == 0 && probabilidad2 == 1) || (probabilidad == 1 && probabilidad2 == 0){
				probabilidad3 = 2
			}
			if (probabilidad == 0 && probabilidad2 == 2) || (probabilidad == 2 && probabilidad2 == 0){
				probabilidad3 = 1
			}
			if (probabilidad == 1 && probabilidad2 == 2) || (probabilidad == 2 && probabilidad2 == 1){
				probabilidad3 = 0
			}
			vivo = data_node(partBuffer, algoritmo, probabilidad3,int(i) , int(totalPartsNum), nombrelibro)
		}
	}

	// now, we close the newFileName
	file.Close()
}

//Funcion que transforma los chunks de un libro en un libro entero

func unirchunks(nombrel string, partes int){
	// just for fun, let's recombine back the chunked files in a new file

	newFileName := nombrel + ".pdf"
	_, err := os.Create(newFileName)

	if err != nil {
			fmt.Println(err)
			os.Exit(1)
	}

	//set the newFileName file to APPEND MODE!!
	// open files r and w

	file, err := os.OpenFile(newFileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err != nil {
			fmt.Println(err)
			os.Exit(1)
	}

	// IMPORTANT! do not defer a file.Close when opening a file for APPEND mode!
	// defer file.Close()

	// just information on which part of the new file we are appending
	var writePosition int64 = 0

	totalPartsNum := partes
	
	for j := uint64(0); j < uint64(totalPartsNum); j++ {

			//read a chunk
			currentChunkFileName := nombrel + "_" + strconv.FormatUint(j, 10)

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
	}

	// now, we close the newFileName
	file.Close()
}

//Funcion que almacena los chunks que le corresponden a la maquina dada

func pedirchunksaldn(maquina string, parte string, nombrel string){
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(maquina, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := dn_proto.NewDnServiceClient(conn)
		
	message := dn_proto.ChunkRequestDN{
		Nombrel: nombrel,
		Partes: parte,
	}

	response, err := c.PedirChunks(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling PedirChunks: %s", err)
	}
	//descagando el chunk
	// write to disk
	fileName := response.Nombrel+ "_" + response.Partes
	_, err = os.Create(fileName)

	if err != nil {
			fmt.Println(err)
			os.Exit(1)
	}

	// write/save buffer to disk
	ioutil.WriteFile(fileName, response.Chunk, os.ModeAppend)
	// log.Printf("chunkkk %d", len(response.Chunk))
}

//Funcion que borra los chunks despues de reensamblarlo de la carpeta del cliente

func borrarchunks(partes int, nombrel string){
	cont := 0
	for{
		if cont == partes{
			break
		}
		contaux := strconv.Itoa(cont)
		path := nombrel + "_" + contaux
		err := os.Remove(path)
	
		if err != nil {
			fmt.Println(err)
			return
		}
		cont = cont + 1
	}
}
	
//Funcion que se conecta con el NameNode para entregar la lista de libros que el cliente puede descargar y las direcciones donde se encuentra cada chunk de un libro escogido

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
		log.Fatalf("Error when calling DisplayLista: %s", err)
	}

	log.Printf("%s", response.Code)
	split := strings.Split(response.Code, "\n")
	var first string
	fmt.Scanln(&first)
	auxfirst, err := strconv.Atoi(first)
	if err != nil {
		log.Fatalf("Error al convertir", err)
	}
	nombredellibro := split[auxfirst+2]
	paratrim := "("+first+")"
	nombrefinal := strings.Trim(nombredellibro, paratrim)

	fmt.Println(nombrefinal)

	mensaje := nn_proto.CodeRequest{
		Code: nombrefinal,
	}

	respuesta, err := c.DisplayDirecciones(context.Background(), &mensaje)
	if err != nil {
		log.Fatalf("Error when calling DisplayDirecciones: %s", err)
	}
	fmt.Println(respuesta)
	
	canttotal := 0
	//pal dn1
	if respuesta.Partes1 != "0," {
		partesdn := strings.Split(respuesta.Partes1, ",")
		tamdn := len(partesdn)-1
		canttotal = canttotal + tamdn
		cont := 0
		maquina := "dist14:9001"
		for{
			if cont == tamdn{
				break
			}
			parte := partesdn[cont]
			pedirchunksaldn(maquina, parte, nombrefinal)
			cont = cont + 1
		}
	}
	//pal dn2
	if respuesta.Partes2 != "0," {
		partesdn := strings.Split(respuesta.Partes2, ",")
		tamdn := len(partesdn)-1
		canttotal = canttotal + tamdn
		cont := 0
		maquina := "dist15:9002"
		for{
			if cont == tamdn{
				break
			}
			parte := partesdn[cont]
			pedirchunksaldn(maquina, parte, nombrefinal)
			cont = cont + 1
		}
	}
	//pal dn3
	if respuesta.Partes3 != "0," {
		partesdn := strings.Split(respuesta.Partes3, ",")
		tamdn := len(partesdn)-1
		canttotal = canttotal + tamdn
		cont := 0
		maquina := "dist16:9003"
		for{
			if cont == tamdn{
				break
			}
			parte := partesdn[cont]
			pedirchunksaldn(maquina, parte, nombrefinal)
			cont = cont + 1
		}
	}
	unirchunks(nombrefinal,canttotal)
	borrarchunks(canttotal,nombrefinal)

	//hacer la funcion del nn para qe le pase las direcciones (jean) (listoko, el cliente pide las direcciones mandando el nombre del libro y las recibe en respuestas)
	//recibir cual libro

	//hacer la funcion del dn para qe le envien los chunks
	//recibir cual libro y sus partes

	//hacer la funcion del cliente para qe descargue los chunks
	//hacer la funcion del cliente para unir los chunks
	//borrar los chunks
}
var listalibros []string

//Funcion que despliega una lista de los libros que el cliente puede subir

 func mostrarlibros() {
 	files, err := ioutil.ReadDir("../libros_cliente")
 	if err != nil {
 		log.Fatal(err)
 	}

	i := 0
 	for _, f := range files {
		 i = i+1
		 fmt.Println("(",i,")", f.Name())
		 listalibros = append(listalibros, f.Name())
	 }
 }

//Funcion que retorna el nombre del libro, con ".pdf" y sin ".pdf"

 func escogerlibro(eleccion string) (string, string){
	var libro string
	elexion, err := strconv.Atoi(eleccion)
	if err != nil{
		log.Fatal(err)
	}
	libropdf := listalibros[elexion-1]
	last := len(libropdf) - 4
	libro = libropdf[:last]
	fmt.Println(libro)
	return libro, libropdf
 }

func main() {
	
	var first string
	var second string
	var libro string

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
			fmt.Println("Escoge un libro de la lista de libros disponibles:")
			mostrarlibros()
			fmt.Println("-----------------")
			fmt.Scanln(&libro)
			librosinpdf, libroconpdf := escogerlibro(libro)
			fmt.Println("-----------------")
			fmt.Println("Escoge: ") 
			fmt.Println("(1) Algoritmo Distribuido") 
			fmt.Println("(2) Algoritmo Centralizado")
			fmt.Println("(0) Salir")
			fmt.Println("-----------------")
			fmt.Scanln(&second)
			if second != "0"{
				separarlibro(second, librosinpdf, libroconpdf)
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
