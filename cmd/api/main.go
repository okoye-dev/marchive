package main

import (
	"log"

	"github.com/okoye-dev/marchive/internal/server"
)

func main(){
	server := server.NewServer()
	log.Fatal(server.Start())
}