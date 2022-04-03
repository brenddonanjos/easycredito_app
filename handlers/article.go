package handlers

import (
	"net/http"

	"github.com/brenddonanjos/easycredito_app/models"
	"github.com/labstack/echo/v4"
)

func getArticles(c echo.Context) error {

	a := models.Articles{}

	// if err := db.Find(&products).Error; err != nil {
	// 	return echo.NewHTTPError(http.StatusNoContent)
	// }

	return c.JSON(http.StatusOK, a)
}

func showArticles(c echo.Context) error {

	a := new(models.Article)
	// id := c.Param("id")
	// db.First(&p, id)
	return c.JSON(http.StatusOK, a)
}

func setArticle(c echo.Context) error {
	a := new(models.Article)

	return c.JSON(http.StatusCreated, a)
}

func updateArticle(c echo.Context) error {
	// db := database.StartGorm()
	// sqlDB, _ := db.DB()
	// defer sqlDB.Close()

	s := new(models.Article)

	return c.JSON(http.StatusCreated, s)
}

func deleteArticle(c echo.Context) error {

	s := new(models.Article)
	return c.String(http.StatusOK, "Produto: "+s.Title+" inativado")
}
