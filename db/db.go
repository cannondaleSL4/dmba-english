package db

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/pgxpool"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Dict struct {
	Dict_id        int
	Word           string
	Word_translate string
}

func init() {
	createTable()
}

//DATABASE_URL - postgres://postgres:postgres@localhost:5432/test_dmba

func PgConnection() *pgxpool.Pool {
	conn, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn
}

func createTable() {
	var connection *pgxpool.Pool
	connection = PgConnection()
	path := filepath.Join("create.sql")
	c, _ := ioutil.ReadFile(path)
	sql := string(c)
	_, err := connection.Exec(context.Background(), sql)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error durig creation table in database: %v\n", err)
		os.Exit(1)
	}
	defer connection.Close()
}

func InsertData(data map[string]string) {
	var connection *pgxpool.Pool
	connection = PgConnection()
	for k, v := range data {
		connection.Exec(context.Background(), "INSERT INTO dict (word, word_translate) VALUES ($1, $2)", k, v)
	}
	defer connection.Close()
}

func CheckUserExist(userId int) bool {
	var connection *pgxpool.Pool
	connection = PgConnection()
	userExist, err := connection.Exec(context.Background(), "Select * from  users where user_number=$1", userId)
	if err != nil {
		fmt.Println("could not make request")
	}
	defer connection.Close()
	return userExist.String() != "SELECT 0"
}

func SaveNewUser(userId int) {
	var connection *pgxpool.Pool
	connection = PgConnection()
	connection.Exec(context.Background(), "INSERT INTO users (user_number) VALUES ($1)", userId)
	defer connection.Close()
}

func GetWords(userId int) []*Dict {
	var result []*Dict
	var connection *pgxpool.Pool
	connection = PgConnection()
	//result, err := connection.Exec(context.Background(),  "Select * from dict_user where users_id=$1 and done=false", userId)
	err := pgxscan.Select(context.Background(), connection, &result, "Select * from dict_user where users_id=$1 and done=false", userId)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting new words for user: %v\n", err)
		return nil
	}
	if result == nil {
		err = pgxscan.Select(context.Background(), connection, &result, "Select * from dict  where dict_id not in (select dict_id from dict_user where users_id=$1) ORDER BY random() limit 10", userId)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting new words for user: %v\n", err)
			return nil
		}
	}
	defer connection.Close()
	return result
}
