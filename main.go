package main

import (
	"github.com/wandersonsantos01/go-gin/databases"
	"github.com/wandersonsantos01/go-gin/routes"
)

func main() {
	databases.DbConnect()
	routes.HandleRequests()
}
