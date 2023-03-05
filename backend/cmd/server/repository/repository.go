package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	"testKwd/cmd/server/config"
	"testKwd/libs/errors"

	proto_bookAuthor_service "testKwd/proto"
)

type Database interface {
	GetBooks(nameAuthor string) ([]*proto_bookAuthor_service.Book, error)
	GetAuthors(nameBook string) ([]*proto_bookAuthor_service.Author, error)
}

type DatabaseConn struct {
	conn *sql.DB
}

func NewDatabase() (*DatabaseConn, error) {
	conf := config.ServerConfig()

	connStr := conf.PgConnString

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		errors.HandleFatalError(err, "failed to connect to database")
	}

	err = db.Ping()
	if err != nil {
		errors.HandleFatalError(err, "failed to connect to database")
	}

	return &DatabaseConn{conn: db}, err
}

func (db DatabaseConn) GetBooks(nameAuthor string) ([]*proto_bookAuthor_service.Book, error) {
	query := `
			SELECT book_name  
			FROM authors
			WHERE name = $1
		`

	rows, err := db.conn.Query(query, nameAuthor)
	if err != nil {
		return nil, errors.HandleDatabaseError(err)
	}

	books := make([]*proto_bookAuthor_service.Book, 0)

	for rows.Next() {
		book := &proto_bookAuthor_service.Book{}

		err = rows.Scan(
			&book.BookName,
		)
		if err != nil {
			return nil, errors.HandleDatabaseError(err)
		}

		books = append(books, book)
	}

	err = rows.Err()

	if err != nil {
		return nil, errors.HandleDatabaseError(err)
	}

	return books, err
}

func (db DatabaseConn) GetAuthors(nameBook string) ([]*proto_bookAuthor_service.Author, error) {
	query := `
			SELECT name 
			FROM authors
			WHERE book_name = $1
		`

	rows, err := db.conn.Query(query, nameBook)
	if err != nil {
		return nil, errors.HandleDatabaseError(err)
	}

	authors := make([]*proto_bookAuthor_service.Author, 0)

	for rows.Next() {
		author := &proto_bookAuthor_service.Author{}

		err = rows.Scan(
			&author.AuthorName,
		)
		if err != nil {
			return nil, errors.HandleDatabaseError(err)
		}

		authors = append(authors, author)
	}

	err = rows.Err()

	if err != nil {
		return nil, errors.HandleDatabaseError(err)
	}

	return authors, err
}
