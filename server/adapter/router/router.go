package router

import (
	"hexagonal-architecture-sample/server/adapter/mysql"
	"net/http"
)

type Router interface{}
type Provide struct {
	User User
}

func NewRouter(r mysql.Resource, p Provide) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			p.User.Create(w, r)
		case http.MethodPut:
			p.User.Update(w, r)
		case http.MethodGet:
			p.User.GetByID(w, r)
		}
	})
	// mux.HandleFunc("/users/update", u.Update)
	mux.HandleFunc("/users/list", p.User.GetAll)
	return mux
}
