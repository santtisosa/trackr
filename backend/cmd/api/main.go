package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/santtisosa/trackr/backend/internal/db"
	"github.com/santtisosa/trackr/backend/internal/handler"
	"github.com/santtisosa/trackr/backend/internal/middleware"
	"github.com/santtisosa/trackr/backend/internal/repository"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	database := db.Connect()
	defer database.Close()

	categoryRepo := repository.NewCategoryRepository(database)
	expenseRepo := repository.NewExpenseRepository(database)

	categoryHandler := handler.NewCategoryHandler(categoryRepo)
	expenseHandler := handler.NewExpenseHandler(expenseRepo)

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	api := r.Group("/api", middleware.AuthMiddleware())
	{
		api.GET("/categories", categoryHandler.GetCategories)

		api.POST("/expenses", expenseHandler.CreateExpense)
		api.GET("/expenses", expenseHandler.GetExpenses)
		api.GET("/expenses/:id", expenseHandler.GetExpense)
		api.PUT("/expenses/:id", expenseHandler.UpdateExpense)
		api.DELETE("/expenses/:id", expenseHandler.DeleteExpense)
	}

	log.Printf("Starting server on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
