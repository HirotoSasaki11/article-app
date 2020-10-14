package router

import (
	"encoding/json"
	"hexagonal-architecture-sample/server/application"
	"hexagonal-architecture-sample/server/application/model"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"
)

type User struct {
	User application.User
}

func (u *User) Create(w http.ResponseWriter, r *http.Request) {
	log.Println("insert started")
	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var user model.User
	body := make([]byte, length)

	_, err = r.Body.Read(body)
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println(err)
	}
	err = u.User.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusCreated)
	}
}

func (u *User) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := u.User.GetAll()
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

func (u *User) GetByID(w http.ResponseWriter, r *http.Request) {
	_, id := ShiftPath(r.URL.Path)
	user, err := u.User.GetByID(id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
func (u *User) Update(w http.ResponseWriter, r *http.Request) {
	log.Println("update started")
	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var user model.User
	body := make([]byte, length)
	log.Println(r.Body)
	r.Body.Read(body)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	err = u.User.Update(user)
	if err != nil {
		w.WriteHeader(http.StatusCreated)
	}
	w.WriteHeader(http.StatusOK)
}

func ShiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i+1:]
}
