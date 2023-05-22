package main

import (
	"app/handler"
	"app/infrastructure"
	"app/usecase"
	"fmt"

	appvalidator "app/handler/validator"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	d, err := infrastructure.NewDB()
	if err != nil {
		fmt.Printf("failed to start server. db setup failed, err = %s", err.Error())
		return
	}
	r := setupRouter(d)
	if err := appvalidator.SetupValidator(); err != nil {
		fmt.Printf("failed to start server. validator setup failed, err = %s", err.Error())
		return
	}
	r.Run()
}

func setupRouter(d *gorm.DB) *gin.Engine {
	r := gin.Default()

	repository := infrastructure.NewTodo(d)
	usecase := usecase.NewTodo(repository)
	handler := handler.NewTodo(usecase)

	todo := r.Group("/todo")
	{
		todo.POST("", handler.Create)
		todo.GET("", handler.FindAll)
		todo.GET("/:id", handler.Find)
		todo.PUT("/:id", handler.Update)
		todo.DELETE("/:id", handler.Delete)
	}
	return r
}
