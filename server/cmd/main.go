package main

import (
	"hexagonal-architecture-sample/server/adapter/mysql"
	"hexagonal-architecture-sample/server/adapter/mysql/dao"
	"hexagonal-architecture-sample/server/adapter/router"
	"hexagonal-architecture-sample/server/application"
	"log"
	"net/http"
)

func main() {

	resources := mysql.NewResource()
	resources.Initialize()
	defer resources.Finalize()
	p := &router.Provide{
		User: router.User{
			User: application.User{
				Interface: dao.ProveideUser(resources.DB),
			},
		},
	}

	err := http.ListenAndServe(":8080", router.NewRouter(resources, *p))
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
