package main

import (
	"flag"
	"log"

	"github.com/JakWai01/http-server/pkg/server"
)

func main() {

	port := flag.String("port", "8000", "Port to listen to")
	flag.Parse()

	listeningPort := ":" + *port
	log.Println(listeningPort)

	httpServer := server.NewHTTPServer(listeningPort)

	if err := httpServer.Open(); err != nil {
		log.Fatal("could not open httpServer", err)
	}

}
