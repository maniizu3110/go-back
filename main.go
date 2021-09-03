package main

import (
	"database/sql"
	"fmt"
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
	fmt.Printf("%+v",conn)
	fmt.Printf("%+v",err)
	if err != nil {
		log.Fatal("cannot connect to db:",err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	server.Start(serverAddress)
}