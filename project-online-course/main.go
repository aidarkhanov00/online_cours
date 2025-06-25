package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"project-online-course/controllers"
	"project-online-course/middleware"
	_ "project-online-course/middleware"
)

// Строка подключения к БД
const dbConnStr = "host=localhost port=5432 user=postgres password=12345 dbname=ent_school sslmode=disable"

func main() {
	// Подключаемся к БД
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}
	defer db.Close()

	// Проверяем подключение
	err = db.Ping()
	if err != nil {
		log.Fatal("БД не отвечает:", err)
	}
	fmt.Println("Успешное подключение к БД")

	r := gin.Default()

	// Разрешаем CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	})

	// Эндпоинт для логина
	r.POST("/login", func(c *gin.Context) {
		var request struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат запроса"})
			return
		}
		api := r.Group("/api")
		api.Use(middleware.JWTMiddleware()) // Применяем middleware

		api.GET("/profile", func(c *gin.Context) {
			username, _ := c.Get("username")
			role, _ := c.Get("role")

			c.JSON(http.StatusOK, gin.H{
				"message":  "Профиль",
				"username": username,
				"role":     role,
			})
		})

		// Проверяем пользователя в БД
		var dbPassword, role string
		err := db.QueryRow("SELECT password, role FROM users WHERE name=$1", request.Username).Scan(&dbPassword, &role)
		if err != nil {
			log.Println("Ошибка запроса к БД:", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный логин или пароль"})
			return
		}

		// Сравниваем пароли (если они не хешированы)
		if request.Password != dbPassword {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный логин или пароль"})
			return
		}

		// Генерируем JWT-токен
		token, err := controllers.GenerateJWT(request.Username, role)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания токена"})
			return
		}

		// Отправляем ответ с токеном
		c.JSON(http.StatusOK, gin.H{
			"message": "Успешный вход",
			"role":    role,
			"token":   token,
		})
	})

	r.Run(":8080")
}
