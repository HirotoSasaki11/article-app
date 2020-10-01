package application_test

import (
	"database/sql"
	"encoding/json"
	"hexagonal-architecture-sample/server/adapter/mysql"
	"hexagonal-architecture-sample/server/adapter/router"
	"hexagonal-architecture-sample/server/application/model"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_GetAll(t *testing.T) {
	type want struct {
		user []model.User
		err  error
	}
	tests := []struct {
		name string
		want want
	}{
		{
			name: "正常系_全てのデータを取得できること",
			want: want{
				user: []model.User{
					{
						ID:        1,
						FirstName: "sa",
						LastName:  "hi",
						Email:     "aa@co.jp",
					},
					{
						ID:        2,
						FirstName: "saa",
						LastName:  "hia",
						Email:     "aaa@co.jp",
					},
				},
				err: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resources := NewResource()
			resources.Initialize()
			defer resources.Finalize()
			defer finalize(resources.DB)

			initialize(resources.DB)

			ts := httptest.NewServer(router.NewRouter(resources))
			defer ts.Close()
			res, err := http.Get(ts.URL + "/user/list")
			if err != nil {
				t.Fatalf("http Get failed:%s", err.Error())
			}

			getting, err := ioutil.ReadAll(res.Body)
			defer res.Body.Close()
			if err != nil {
				t.Fatalf("read from HTTP Response Body failed:%s", err.Error())
			}

			data, err := json.Marshal(tt.want.user)
			if err != nil {
				log.Println(err)
			}
			if string(getting) != string(data) {
				t.Fatalf("return : %s , want %v", string(getting), string(data))
			}
			// req:= httptest.NewRequest("Get","/user/list",nil)
			// rec := httptest.NewRecorder()
		})
	}
}

func NewResource() mysql.Resource {
	return mysql.Resource{
		Config: initializeConfig(),
	}
}

func initializeConfig() *mysql.DatabaseConfig {
	return &mysql.DatabaseConfig{
		User:       "root",
		Pass:       "root",
		Protocol:   "tcp",
		Connection: "127.0.0.1:3307",
		// Connection: os.Getenv("MYSQL_HOST"),
	}
}

func initialize(d *sql.DB) {
	_, err := d.Exec("create database test_sample")
	if err != nil {
		panic(err)
	}
	_, err = d.Exec("use test_sample")
	if err != nil {
		panic(err)
	}
	initSql, err := ioutil.ReadFile("./testdata/initialize.sql")
	if err != nil {
		panic(err)
	}
	for _, v := range strings.Split(string(initSql), ";") {
		d.Exec(v)
		if err != nil {
			panic(err)
		}
	}

	insertTestdata(d)
}

func finalize(d *sql.DB) {
	_, err := d.Exec("drop database test_sample")
	if err != nil {
		panic(err)
	}
}

var testData string = `INSERT INTO users(id,first_name, last_name, email) VALUES(1,"sa" ,"hi", "aa@co.jp"),(2,"saa", "hia", "aaa@co.jp")`

func insertTestdata(d *sql.DB) {
	_, err := d.Exec(testData)
	if err != nil {
		panic(err)
	}
}
