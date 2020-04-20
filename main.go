package main

import "go-user-auth-api/infrastructure/server"

// @title Go-user-auth-api
// @version 1.0
// @description Go User Auth Api ( Returns all user,permission and roles information differs by application etc )
func main() {
	server.StartServer()
}
