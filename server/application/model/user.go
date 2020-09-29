package model

type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first-name"`
	LastName  string `json:"last-name"`
	Email     string `json:"email"`
}
