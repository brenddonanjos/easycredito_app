package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func App() {
	e := echo.New()

	//TODO: Code Middleware here
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{ //enable cors
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	//Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Back-end Challenge 2021 🏅 - Space Flight News")
	})
	//Articles
	e.GET("/articles", getArticles)
	e.GET("/articles/{id}", showArticles)
	e.POST("/articles", setArticle)
	e.PUT("/articles", updateArticle)
	e.DELETE("/articles", deleteArticle)
	//Start aplication
	e.Logger.Fatal(e.Start(":8000"))
}
