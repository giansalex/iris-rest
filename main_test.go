package main

import (
	"testing"

	"github.com/kataras/iris/v12/httptest"
)

func TestNewApp(t *testing.T) {
	app := newApp()
	e := httptest.New(t, app)

	e.GET("/").Expect().Status(httptest.StatusOK)
	e.GET("/hello/giansalex").Expect().Status(httptest.StatusOK).
		Body().Equal("Hello giansalex")

	e.POST("/api/v1/users").Expect().Status(httptest.StatusNotFound)
	e.GET("/api/v1/users").Expect().Status(httptest.StatusUnauthorized)

	user := map[string]string{
		"username": "admin",
		"password": "123456",
	}

	schema := `{
		"type": "object",
		"properties": {
			"token":  {"type": "string"},
			"expire": {"type": "number"}
		},
		"required": ["token", "expire"]
	}`

	schemaUsers := `{
		"type": "array",
		"items": {
			"type": "object",
			"properties": {
				"name":  {"type": "string"},
				"email": {"type": "string"}
			},
			"required": ["name", "email"]
		}
	}`

	result := e.POST("/api/login").WithJSON(user).Expect().Status(httptest.StatusOK).
		JSON()

	result.Schema(schema)
	token := result.Object().Value("token").String().Raw()
	users := e.GET("/api/v1/users").WithHeader("Authorization", "Bearer "+token).
		Expect().Status(httptest.StatusOK).JSON()

	users.Schema(schemaUsers)
}
