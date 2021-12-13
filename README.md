# Generic Auth Service Project
The code structure and conventions inspired by tutorialedge.net

## Requirements to run the application in development environment
- Docker (required)

## Run the application in the development environment
1. Create a .env file at the root level of the project with the following environment variables.
```
MYSQL_CONNECTION_STRING=username:password@tcp(host:port)/dbname?parseTime=true
AUTH_SECRET=yourauthsecretkey
```
4. You should be able to run the server with the following command
``` 
go run cmd/server/main.go
```
