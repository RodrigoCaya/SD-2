package nn_proto

import (
	"golang.org/x/net/context"
)

type Server struct{
}


var nombres []string

func (s *Server) Buscar(ctx context.Context, message *CodeRequest) (*CodeRequest, error) {
	file, err := os.Open("log.txt")
	if err != nil {
		log.fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		nombres = append(nombres, scanner.Text())
	}

	return &CodeRequest{Code: nombres}, nil
}