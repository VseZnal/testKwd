# testKwd

## Installation

Make sure that docker is installed (for windows or mac install docker desktop)

1. Clone this repo
2. Copy `.env.template.local` file contents to file `.env.local`, and write variables values
3. Run `source .env.local`
4. Run `docker-compose up --build -d` to start backend
5. Run `make migrate` to apply all migrations to database
6. Run `make feed` to insert test data to database
7. Run `docker-compose ps` to check if all services are UP


open terminal
1. Run cd backend
2. Run `go run ./cmd/server main.go` to start server
3. Run `go run ./cmd/client/main.go` to start client


Client connection:

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
