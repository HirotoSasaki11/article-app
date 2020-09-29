package dao

import (
	"database/sql"
	"hexagonal-architecture-sample/server/application/model"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	db *sql.DB
}

func ProveideUser(db *sql.DB) *User {
	return &User{db: db}
}
func (u *User) Create(user model.User) error {
	stmt, err := u.db.Prepare("INSERT INTO users(first_name, last_name, email) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email)
	if err != nil {
		log.Println("インサート処理失敗")
		return err
	}
	return nil
}
func (u *User) GetAll() ([]model.User, error) {
	var users []model.User

	result, err := u.db.Query("select * from users")
	if err != nil {
		return nil, err
	}
	for result.Next() {
		var user model.User
		err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
