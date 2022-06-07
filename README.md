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

# All steps
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