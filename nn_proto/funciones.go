package nn_proto

import (
	"golang.org/x/net/context"
)

type Server struct{
}

var nombres string = ""

func (s *Server) Buscar(ctx context.Context, message *CodeRequest) (*CodeRequest, error) {
	return &CodeRequest{Code: "xd"}, nil
}


func (s *Server) DisplayLista(ctx context.Context, message *CodeRequest) (*Lista, error) {
/*	f, err := os.Open("log.txt")
	if err != nil{
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
*/
	nombres = nombres + "Dracula," + "Frankenstein,"
	

	return &Lista{L: nombres}, nil
}
