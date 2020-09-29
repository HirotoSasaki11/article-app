package router

import (
	"encoding/json"
	"hexagonal-architecture-sample/server/application"
	"hexagonal-architecture-sample/server/application/model"
	"log"
	"net/http"
)

type user struct {
	user application.User
}

func (u *user) Create(w http.ResponseWriter, r *http.Request) {
	var user *model.User
	var body []byte
	_, err := r.Body.Read(body)
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(body, user)
	if err != nil {
		log.Println(err)
	}
	err = u.user.Create(*user)
	if err != nil {
		w.WriteHeader(http.StatusCreated)
	}
}

func (u *user) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := u.user.GetAll()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	data, err := json.Marshal(users)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
