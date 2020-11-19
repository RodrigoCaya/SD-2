package main

import (
        "log"
        //"fmt"
        "google.golang.org/grpc"
        "context"
        "github.com/RodrigoCaya/SD-2/dn_proto"
)

func main() {
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

