package application_test

import (
	"hexagonal-architecture-sample/server/application"
	"hexagonal-architecture-sample/server/application/model"
	"net/http"
	"net/http/httptest"
	"testing"
	"hexagonal-architecture-sample/server/application/router"
)

func Test_GetAll(t *testing.T) {
	tests := []struct {
		name   string
		want []model.User
		args   args
	}{
		{
			name: "正常系_全てのデータを取得できること",
			want: []model.User{
				{
					ID: 1,
					FirstName: "sa",
					LastName: "hi",
					Email: "aa@co.jp",
				},
				{
					ID: 2,
					FirstName: "saa",
					LastName: "hia",
					Email: "aaa@co.jp",
				}
			},
		}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := router.NewRouter()
			
			req:= httptest.NewRequest("Get","/user/list",nil)
			rec := httptest.NewRecorder()
			
			assert.Equal()
		})
	}
}
