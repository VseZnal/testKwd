package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"testKwd/cmd/server/config"
	"testKwd/cmd/server/repository"
	"testKwd/libs/errors"
	proto_bookAuthor_service "testKwd/proto"
)

type Server struct {
	proto_bookAuthor_service.UnimplementedBookAuthorServiceServer
}

var db repository.Database

func Init() error {
	var err error

	db, err = repository.NewDatabase()
	return err
}

func main() {
	err := Init()
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	conf := config.ServerConfig()

	serverPort := conf.PortServer

	lis, err := net.Listen("tcp", ":"+serverPort)
	if err != nil {
		log.Fatal(err)
	}
	defer lis.Close()

	server := grpc.NewServer()

	proto_bookAuthor_service.RegisterBookAuthorServiceServer(server, Server{})

	log.Fatal(server.Serve(lis))
}

func (s Server) GetBooks(ctx context.Context,
	r *proto_bookAuthor_service.GetBookRequest,
) (*proto_bookAuthor_service.GetBookResponse, error) {
	books, err := db.GetBooks(r.AuthorName)
	if err != nil {
		return nil, errors.LogError(err)
	}

	return &proto_bookAuthor_service.GetBookResponse{Books: books}, err
}

func (s Server) GetAuthors(ctx context.Context,
	r *proto_bookAuthor_service.GetAuthorRequest,
) (*proto_bookAuthor_service.GetAuthorResponse, error) {
	authors, err := db.GetAuthors(r.BookName)
	if err != nil {
		return nil, errors.LogError(err)
	}

	return &proto_bookAuthor_service.GetAuthorResponse{Authors: authors}, err
}
