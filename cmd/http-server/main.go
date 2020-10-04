package main

import (
	"flag"
	"log"
)

func main() {

	port := flag.String("port", "8000", "Port to listen to")
	flag.Parse()

	listeningPort := ":" + *port
	log.Println(listeningPort)

}
