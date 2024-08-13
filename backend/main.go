package main

import (
	"database/sql"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
)

type Question struct {
	ID       int
	Question string
	Options  []string
	Answer   int
}

func getQuestions(c echo.Context) error {
	db, err := sql.Open("sqlite3", "./quiz.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, question, options, answer FROM questions")
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error fetching questions")
	}
	defer rows.Close()

	var questions []Question
	for rows.Next() {
		var question Question
		var options string
		if err := rows.Scan(&question.ID, &question.Question, &options, &question.Answer); err != nil {
			return c.String(http.StatusInternalServerError, "Error scanning question")
		}
		question.Options = strings.Split(options, ",")
		questions = append(questions, question)
	}
	return c.JSON(http.StatusOK, questions)

	// question := []Question{
	// 	{ID: 1, Question: "人間の赤ちゃんの妊娠期間はどれくらいですか?", Options: []string{"38 weeks", "40 weeks", "42 weeks"}, Answer: 1},
	// }
	// return c.JSON(http.StatusOK, question)
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("/api/questions", getQuestions)
	e.Logger.Fatal(e.Start(":8080"))
}
