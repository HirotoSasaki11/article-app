package router

import (
	"hexagonal-architecture-sample/server/adapter/mysql"
	"hexagonal-architecture-sample/server/adapter/mysql/dao"
	"hexagonal-architecture-sample/server/application"
	"net/http"
)

func NewRouter(r mysql.Resource) *http.ServeMux {
	u := &user{
		user: application.User{
			Interface: dao.ProveideUser(r.DB),
		},
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/user/create", u.Create)
	mux.HandleFunc("/user/list", u.GetAll)
	return mux
}
