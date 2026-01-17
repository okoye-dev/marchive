package main

import (
	"log"

	"github.com/okoye-dev/marchive/internal/config"
	"github.com/okoye-dev/marchive/internal/server"
)

func main(){
	cfg := config.LoadConfig()
	server := server.NewServer(cfg)
	
	log.Fatal(server.Start())
}