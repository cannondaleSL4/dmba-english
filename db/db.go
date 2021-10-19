package db

import (
	"context"
	"fmt"
	"github.com/dmba-english/dict"
	"github.com/jackc/pgx"
	"io/ioutil"
	"os"
	"path/filepath"
)

type UserDict struct {
	Words  *[]dict.Words
	UserId string
}

var connection *pgx.Conn

func init() {
	connection = pgConnection()
	createTable()
}

//DATABASE_URL - postgres://postgres:postgres@localhost:5432/test_dmba

func pgConnection() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn
}

func createTable() {
	path := filepath.Join("create.sql")
	c, _ := ioutil.ReadFile(path)
	sql := string(c)
	_, err := connection.Exec(context.Background(), sql)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error durig creation table in database: %v\n", err)
		os.Exit(1)
	}
	defer connection.Close(context.Background())
}

func Test() {
	fmt.Println("test")
}
