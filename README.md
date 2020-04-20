go-user-auth-api
===============

Go User Auth Api ( Returns all user,permission and roles information differs by application etc )

## Build

go mod vendor

go build main.go

## Run 

First Run docker-compose up -d to start sqlserver

Second run DDL under the infrastructure/resource/go_user_auth_api_schema.sql


## Swagger
http://localhost:6161/swagger/index.html

## Postman
Use infrastructure/resource/go-user-auth-api.postman_collection.json to run sample requests