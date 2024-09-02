package main

import (
	"database/sql"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

func registerUser(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	log.Printf("Username: %s", username)
	log.Printf("Password: %s", password)

	// passwordハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error during database connection:", err)
		return c.String(http.StatusInternalServerError, "Error hashing password")
	}

	// DBにユーザーを保存
	db, err := sql.Open("sqlite3", "./quiz_app.db")
	if err != nil {
		return c.String(http.StatusInternalServerError, "Database connection error")
	}
	defer db.Close()

	// 既にユーザー名が存在するかチェック
    var existingUsername string
    err = db.QueryRow("SELECT username FROM users WHERE username = ?", username).Scan(&existingUsername)
    if err == nil {
        return c.JSON(http.StatusConflict, echo.Map{"message": "Username already exists"})
    }

	_, err = db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, hashedPassword)
	if err != nil {
		log.Println("Error during user registration:", err)
		return c.String(http.StatusInternalServerError, "Error saving user to database")
	}

	return c.String(http.StatusOK, "User registered successfully")
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func login(c echo.Context) error {
	req := new(LoginRequest)
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	var storedHash string

	db, err := sql.Open("sqlite3", "./quiz_app.db")
	if err != nil {
		return c.String(http.StatusInternalServerError, "Database connection error")
	}
	defer db.Close()

	err = db.QueryRow("SELECT password FROM users WHERE username = ?", req.Username).Scan(&storedHash)

	if err != nil {
        return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid username or password"})
    }
	// パスワード比較
	if err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(req.Password)) ; err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid username or password"})
	}

	// セッション開始やトークン発行などの処理
	// session, _ := store.Get()

	return c.JSON(http.StatusOK, map[string]string{"message": "Login successful"})
}

func saveQuizResult(c echo.Context) error {
	userID := c.FormValue("user_id")
	quizID := c.FormValue("quiz_id")
	correctAnswers := c.FormValue("correct_answers")
	totalQuestions := c.FormValue("total_questions")

	db, err := sql.Open("sqlite3", "./quiz_app.db")
	if err != nil {
		return c.String(http.StatusInternalServerError, "Database connection error")
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO quiz_results (user_id, quiz_id, correct_answers, total_questions, created_at) VALUES (?, ?, ?, ?, ?)", userID, quizID, correctAnswers, totalQuestions, time.Now())
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error saving quiz result to database")
	}

	return c.String(http.StatusOK, "Quiz result saved successfully")
}

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
	// 問題取得
	e.GET("/api/questions", getQuestions)

	// ユーザー登録
	e.POST("/api/register", registerUser)

	// ログイン
	e.POST("/api/login", login)

	// クイズ結果の保存
	e.POST("/api/save", saveQuizResult)
	e.Logger.Fatal(e.Start(":8080"))
}
