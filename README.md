# go-gen-backend
A generic backend developed using go - gin

# Aim
Explore go programming creating backend rest api server.

# Prerequisites 
go installed see [Go Dev Installation Page](https://go.dev/doc/install)

# Authentication
Rest services are exposed under token-authentication,
login is verified against an ldap server

# Folder
- keys contains keys used to implement tls

# Services 
- login
- healthcheck
- gettask

# Configuration
- update ./config yml file accordingly your environment
- set up the env variable ENVIRONMENT_TYPE pointing to
  - DEV       development env values
  - PROD      production env values
  - TEST      test env values
  - AUTOTEST  automatic test env values

# Commands

- build:     go build 
- run:       go run main.go
- unit test: go test ./...

# Contents
- internal
  - config: configuration loading
  - handlers: exposed functions handlers
    - get.task: based on the received id return all task data ( tbd: if a body is present it must be validated and should contains requested properties )
    - healthcheck: healtcheck service
  - middleware
    - middleware.auth: receive the
  - runtimes: 
    - runtime.backend1.client: this comes from a requirement I struggle with some time ago. I should invoke a backend web service accessible using a basic authentication ( user/password ). To make it a little more performant I intend to use a persistent http but I'm not able to maintain it accross many routins ( automatically instantiated by gin ). The idea here is instantiate a number of this runtime setting up a connection to the backend , an incoming channel will receive new requests must be forwarded to the backend, each request will contains also the response channel. Here is not implemented the web service call.

- after we created the main folder we have to initialize the new module assigning it a name:
    go mod init github.com/a11dev/go-gen-backend
- I have chosen a project structure pointing to the modularization, we are going to create:
  - config: it contains environments configuration files
  - user : it contains user actions like login
  - task : it contains task actions
  - utilities : common functions
  - interfaces : backend to southbound system interface
  - middleware : contains gin middleware
- add gin to our module: go get github.com/gin-gonic/gin
- Into the root create main.go
- 
- create authentication 


# Run

## Run Test

go test <test path>
or
go test <test path>/...  
to test starting from a path