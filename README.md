# Docker_Test
This is a simple implementation of setting up a Docker Container for a Simple Webserver in Golang

#### Starting the server
```bash
  cd {Directory}/docker_test
```
- Build the project
```bash
  docker build -t docker_test
```
- Run the main.go file
```bash
  docker run -p 8080:8081 -it docker_test
```
