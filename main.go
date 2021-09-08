package main

import (
	"database/sql"
	"log"
	"os"
	"simplebank/api"
	"simplebank/api/sqlc"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

//TODO:設定ファイルから読み込む
const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = ":8080"
)

func init(){
	logrus.SetOutput(os.Stdout)
	logLevel,err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logLevel = logrus.InfoLevel
	}
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logLevel)
}


func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := sqlc.NewStore(conn)
	server := api.NewServer(store)

	server.Start(serverAddress)
}
