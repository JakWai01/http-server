# How to deploy a containerized Go application in a Docker container image and run that container on a Kubernetes cluster 

## 0. Table of Contents

0. Table of Contents
1. Introduction
2. File structure 
3. Write HTTP server in Go
4. Dockerize HTTP server :whale:
5. Kubernetes yaml setup
6. Run contsainer on Kubernetes cluster
7. Conclusion

## 1. Introduction

The following Tutorial will instruct the reader on deploying a containerized Go apllication in a Docker container image and run that container on a Kubernetes cluster. Therefore we will build a simple HTTP Server.

### Requirements

1. A working Kubernetes cluster
2. Docker (on your local machine)
3. Go language fundamentals

## 2. File structure

This section will be about the file structure used in this tutorial. "It's not an official standard defined by the core Go dev team; however, it is a set of common historical and emerging project layout patterns in the Go ecosystem." (https://github.com/golang-standards/project-layout)

### File structure 

![img](https://github.com/JakWai01/http-server/blob/main/images/file-structure.png "File structure")

The picture above displays the file structure used for this project as well as all files you will need. You can also check out the repository in general above the README to get into it.

## 3. Write HTTP server in Go

The almost simplest HTTP Server you can write in Go looks like this: 
```
package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}

```
In this tutorial though, we split our application into two parts. The `main.go` and the `http.go`. The `http.go` looks like this: 
```
package server

import (
	"io"
	"net/http"
)

// HTTPServer creates a http server and can be reached through the porvided port
type HTTPServer struct {
	port string
}

// NewHTTPServer initializes variables
func NewHTTPServer(port string) *HTTPServer {
	return &HTTPServer{port}
}

// Open creates the http server
func (s HTTPServer) Open() error {
	http.HandleFunc("/", home)
	http.ListenAndServe(s.port, nil)

	return nil
}

func home(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World")
}
```
and the `main.go` looks like this: 
```
package main

import (
	"flag"
	"log"

	"github.com/JakWai01/http-server/pkg/server"
)

func main() {

	port := flag.String("port", "8080", "Port to listen to")
	flag.Parse()

	listeningPort := ":" + *port
	log.Println(listeningPort)

	httpServer := server.NewHTTPServer(listeningPort)

	if err := httpServer.Open(); err != nil {
		log.Fatal("could not open httpServer", err)
	}

}
```
## 4. Dockerize HTTP Server


