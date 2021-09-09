package main

import (
	"database/sql"
	"log"
	"os"
	"simplebank/api"
	"simplebank/api/sqlc"
	"simplebank/api/util"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
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
	config,err := util.LoadConfig(".")
	if err != nil{
		log.Fatal("cannot load config:",err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := sqlc.NewStore(conn)
	server,err := api.NewServer(config,store)
	if err != nil {
		log.Fatal("cannot create new server:",err)
	}

	server.Start(config.ServerAddress)
}
