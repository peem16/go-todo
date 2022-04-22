package main

import (
	"context"
	"go-todo-service/handler"
	"go-todo-service/repository"
	"go-todo-service/router"
	"go-todo-service/service"

	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading environment: %v", err)
	}
}

func main() {
	_, err := os.Create("/tmp/live")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove("/tmp/live")
	db := initDatabase()

	todoRepository := repository.NewRepositoryDB(db)
	todoService := service.NewTodoService(todoRepository)
	todoHandler := handler.NewTodoHandler(todoService)

	r := router.New()

	r.GET("api/v1/todo", todoHandler.GetAll)
	r.GET("api/v1/todo/:id", todoHandler.GetByID)
	r.POST("api/v1/todo", todoHandler.NewTodo)
	r.PUT("api/v1/todo/:id", todoHandler.UpdateByID)
	r.DEL("api/v1/todo/:id", todoHandler.Delete)

	r.ListenAndServe()()
}

func initDatabase() *mongo.Database {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("DSN")))
	if err != nil {
		panic("failed to connect database")
	}
	database := client.Database("myapp")

	return database
}
