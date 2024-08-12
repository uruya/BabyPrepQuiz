package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Question struct {
	ID       int
	Question string
	Options  []string
	Answer   int
}

func getQuestions(c echo.Context) error {
	question := []Question{
		{ID: 1, Question: "人間の赤ちゃんの妊娠期間はどれくらいですか?", Options: []string{"38 weeks", "40 weeks", "42 weeks"}, Answer: 1},
	}
	return c.JSON(http.StatusOK, question)
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("/api/questions", getQuestions)
	e.Logger.Fatal(e.Start(":8080"))
}
