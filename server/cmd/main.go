package main

import (
	"hexagonal-architecture-sample/server/adapter/mysql"
	"hexagonal-architecture-sample/server/adapter/router"
	"log"
	"net/http"
)

func main() {

	resources := mysql.NewResource()
	resources.Initialize()
	defer resources.Finalize()

	err := http.ListenAndServe(":8080", router.NewRouter(resources))
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
