# go-gen-backend
A generic backend developed using go - gin

# aim
Explore go programming creating backend rest api server.

# authentication
Rest services are exposed under token-authentication,
login is verified against a ldap server

# services 
- login
- logout
- healthcheck
- gettask



# All steps
- after we created the main folder we have to initialize the new module assigning it a nam:
    go mod init github.com/a11dev/go-gen-backend
- I have chosen a project structure pointing to the modularization, we are going to create:
  - user : it contains user actions like login
  - task : it contains task actions
  - utilities : common functions
  - interfaces : backend to southbound system interface
  - middleware : contains gin middleware
- add gin to our module: go get github.com/gin-gonic/gin
- Into the root create main.go
- 
- create authentication 
