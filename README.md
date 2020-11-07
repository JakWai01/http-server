# How to deploy a containerized Go application in a Docker container image and run that container on a Kubernetes cluster 

## 0. Table of Contents

0. Table of Contents
1. Introduction
2. File structure 
3. Write HTTP server in Go
4. Dockerize HTTP server :whale:
5. Kubernetes yaml setup
6. Run container on Kubernetes cluster
7. Conclusion

## 1. Introduction

The following Tutorial will instruct the reader on deploying a containerized Go application in a Docker container image and run that container on a Kubernetes cluster. Therefore we will build a simple HTTP Server.

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

Now we need to containerize our application with Docker. Therefore, create a "Dockerfile" as shown in the file structure above. Before you do this, be sure to type `go mod init` into a terminal in your project folder to create a `go.mod` file. Then include the following code: 
```
FROM golang:alpine AS build

RUN apk add git

RUN mkdir /src
ADD . /src
WORKDIR /src

RUN go build -o /tmp/http-server ./cmd/http-server/main.go

FROM alpine:edge

COPY --from=build /tmp/http-server /sbin/http-server

CMD /sbin/http-server
```
If you are currently working on your own application just replace "http-server" with your projects name and "./cmd/http-server/main.go" with the path to your executable main.

Next you need to create the docker image. You need to name your project following this convention: "$DOCKERHUB_USERNAME/$PROJECTNAME" (If you do not have a Dockerhub account by now, create one)

```
docker build -t jakwai01/http-server .
```
Now check if your container works
```
docker run -p 8080:8080 jakwai01/http-server
```
If this works, login to Dockerhub
```
docker login
```
Finally, push your image
```
docker push jakwai01/http-server
```
If you got trouble at any of these steps, maybe consider double checking with another tutorial(https://hackernoon.com/publish-your-docker-image-to-docker-hub-10b826793faf)

## Kubernetes yaml setup

Now we need to create a stack.yaml file (you can choose the name of the yaml file by yourself). 

```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: http-server
spec:
  selector:
    matchLabels:
      app: http-server
  template:
    metadata:
      labels:
        app: http-server
    spec:
      containers:
        - name: http-server
          image: jakwai01/http-server
          resources:
            limits:
              memory: "128Mi"
              cpu: "500m"
          ports:
            - containerPort: 8080

---
apiVersion: v1
kind: Service
metadata:
  name: http-server
spec:
  selector:
    app: http-server
  ports:
    - port: 8080
      targetPort: 8080

---
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: http-server
spec:
  backend:
    serviceName: http-server
    servicePort: 8080
  rules:
    - host: http-server.services.jakobwaibel.com
      http:
        paths:
          - backend:
              serviceName: http-server
              servicePort: 8080
```
If you have trouble understanding why we use Ingress right here, consider checking out https://medium.com/google-cloud/kubernetes-nodeport-vs-loadbalancer-vs-ingress-when-should-i-use-what-922f010849e0. Furthermore, the "host: http-server.services.jakobwaibel.com" is just an address assigned to all of my services. 

## 6. Run container on Kubernetes cluster

To run the container on your Kubernetes cluster, just use the following command in your project folder where your `stack.yaml` is located.
```
kubectl apply -f .
```
Now there should be a pod running the application on your Kubernetes cluster.

## 7. Conclusion

If everything worked, you just deployed a containerized Go application in a Docker container image and ran that container on a Kubernetes cluster. If you had any trouble following this tutorial, please let me know. I will try to update this tutorial from time to time to make it easy to follow as my own understanding of Docker and Kubernetes improves. Thanks for reading :).

