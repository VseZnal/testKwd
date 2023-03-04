package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"testKwd/backend/cmd/server/config"
	bookAuthor_service "testKwd/backend/proto"
)

type Server struct {
	bookAuthor_service.UnimplementedBookAuthorServiceServer
}

func main() {
	conf := config.ServerConfig()

	serverPort := conf.PortServer

	lis, err := net.Listen("tcp", ":"+serverPort)
	if err != nil {
		log.Fatal(err)
	}
	defer lis.Close()

	server := grpc.NewServer()

	bookAuthor_service.RegisterBookAuthorServiceServer(server, Server{})

	log.Fatal(server.Serve(lis))
}
