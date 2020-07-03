package main

import (
	"net/http"
	"os"
	"time"

	"github.com/giansalex/iris-rest/model"
	"github.com/kataras/iris/v12"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/iris-contrib/middleware/cors"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

func newApp() *iris.Application {
	app := iris.New()

	jwtHandler := jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	app.Logger().SetLevel("debug")
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())

	app.Get("/", index)
	app.Get("/hello/{name:string}", hello)

	api := app.Party("/api", corsHandler).AllowMethods(iris.MethodOptions)
	{
		api.Post("/login", login)

		v1 := api.Party("/v1")
		{
			v1.Use(jwtHandler.Serve)
			v1.Get("/users", users)
		}
	}

	return app
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	app := newApp()
	app.Run(iris.Addr(":"+port), iris.WithoutServerError(iris.ErrServerClosed))
}

// Handler
func index(c iris.Context) {
	c.StatusCode(http.StatusOK)
	c.WriteString("Iris REST API")
}

func hello(c iris.Context) {
	name := c.Params().Get("name")

	c.WriteString("Hello " + name)
}

func users(c iris.Context) {
	list := []*model.User{
		{
			Name:  "Jon",
			Email: "jon@labstack.com",
		},
		{
			Name:  "GianCarlos",
			Email: "giansalex@gmail.com",
		},
	}
	c.JSON(list)
}

func login(c iris.Context) {
	auth := new(model.Auth)
	err := c.ReadJSON(auth)
	if err != nil {
		c.StatusCode(iris.StatusBadRequest)
		c.WriteString(err.Error())
		return
	}

	if auth.Username == "admin" && auth.Password == "123456" {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "Giancarlos"
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			c.StatusCode(iris.StatusInternalServerError)
			c.WriteString(err.Error())
		}

		c.JSON(map[string]interface{}{
			"token":  t,
			"expire": claims["exp"],
		})
	}

	c.StatusCode(iris.StatusUnauthorized)
}
