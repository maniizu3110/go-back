package main

import (
	"database/sql"
	"log"
	"simplebank/api"
	db "simplebank/db/sqlc"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = ":8080"
)

func main(){
	conn,err := sql.Open(dbDriver,dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:",err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	server.Start(serverAddress)
}