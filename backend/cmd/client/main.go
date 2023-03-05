package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"testKwd/cmd/server/config"
	proto_bookAuthor_service "testKwd/proto"
)

func main() {
	ctx := context.Background()

	conf := config.ServerConfig()

	host := conf.HostServer
	port := conf.PortServer

	clientServiceAddress := fmt.Sprintf("%s:%s", host, port)

	conn, err := grpc.Dial(clientServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}

	client := proto_bookAuthor_service.NewBookAuthorServiceClient(conn)
	author, err := client.GetAuthors(ctx, &proto_bookAuthor_service.GetAuthorRequest{BookName: "testBook3"})
	if err != nil {
		log.Fatalln(err)
	}

	book, err := client.GetBooks(ctx, &proto_bookAuthor_service.GetBookRequest{AuthorName: "testAuthor4"})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(author)
	log.Println(book)
}
