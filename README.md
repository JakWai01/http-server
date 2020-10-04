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

(Insert Image here)

The picture above displays the file structure used for this project as well as all files you will need. You can also check out the repository in general above the README to get the hang of it. 


