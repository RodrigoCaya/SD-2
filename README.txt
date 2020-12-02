2020-2 San Joaquin
Rodrigo Cayazaya Marin, ROL: 201773538-4
Jean-Franco Zárate, ROL: 201773524-4

La tarea consistía en crear una biblioteca de libros digitales para subir y descargar libros, estos libros se encuentran distribuidos en chunks de 250Kb entre 3 data nodes.
Para esta tarea se utlizaron 5 servicios: cliente, data node 1, data node 2, data node 3 y name node. Cada una se encuentra en una carpeta con el mismo nombre.
Se utilizaron 2 tipos de algoritmos, uno de exclusión mutua distribuida y otro de exclusión mutua centralizada.
El registro de las localizaciones de los chunks se encuentra en log.txt dentro de la carpeta name_node, donde es este último el encargado de llevar el registro.
Para trabajar los conflictos producidos cuando 2 data nodes intenten escribir en el log.txt, se utilizó el algoritmo de Ricart y Agrawala, pero en vez de utilizar el tiempo como medida de cambio, se utilizó una id para cada data node.
Dentro de las carpetas "dn_proto" y "nn_proto" se encuentran las funciones utilizadas por los data nodes y el name node, tambien se encuentran sus .proto y .pb.go. correspondientes.

La conexión entre servicios se realizó de la siguiente manera:
  Los data nodes son servidores grpc, al igual que el name node.
  El cliente se puede conectar tanto a los data nodes como al name node.
  Los data nodes y name node también se comunican entre ellos a través de sus funciones.

Requisitos:
	-Go
	-Grpc
	-Protocol-Buffer
  
Instrucciones:	
	Dentro de cada servicio se encuentra un makefile que se ejecuta de la siguiente manera:
	-Para los data nodes
	*make run para ejecutar.
	-Para name node
	*make run para ejecutar y make clean para limpiar los archivos csv.
	-Para cliente
	*make run para ejecutar y make clean para limpiar los archivos csv.
  
Consideraciones generales:
  -Los libros disponibles para subir se encuentran en la carpeta "libros_cliente" y son los que se muestran al intentar subir un libro (deben ser archivos pdf).
  -Dentro de cada data nodes, los chunks se descargan dentro de una carpeta con el mismo nombre.
  -Los libros descargados se descargan en formato pdf y se encuentran dentro de la carpeta cliente.  

Consideraciones máquinas virtuales:
  El name node se debe correr en la máquina dist13.
  El data node 1 se debe correr en la máquina dist14.
  El data node 2 se debe correr en la máquina dist15.
  El data node 3 se debe correr en la máquina dist16.
  El cliente se puede correr en cualquiera de las 4 máquinas.
  
Consideraciones compilación:
	-Se utiliza "protoc --go_out=plugins=grpc:../dn_proto dn.proto" para compilar el archivo dn.proto.
  -Se utiliza "protoc --go_out=plugins=grpc:../nn_proto nn.proto" para compilar el archivo nn.proto.
	-Si aparece un error "Plugin failed with status code 1" al compilar, se deben utilizar los siguientes comandos:
      export GOROOT=/usr/local/go
      export GOPATH=$HOME/go
      export GOBIN=$GOPATH/bin
      export PATH=$PATH:$GOROOT:$GOPATH:$GOBIN
  
Consideraciones al ejecutar:
  -El orden de ejecución es: primero data_nodes/name_node y luego cliente.
	-Al ejecutar el cliente, este le pedirá si desea subir o descargar un libro, debe especificar que libro quiere subir o descargar escogiendo dentro de la interfaz. Si desea subir un libro, también le pedirá con que algoritmo quiere subirlo (distribuido o centralizado).
	
Consideraciones al "caer" una máquina:
  -Las máquinas se pueden caer utilizando control+c.
  -También puede correr 2 o 1 o ninguna máquina si se quiere.
  -El name node no se puede caer, solo los data nodes.
  -La máquina no se puede caer mientras realiza el algoritmo de distribución (el resto de las máquinas si).
  
  
