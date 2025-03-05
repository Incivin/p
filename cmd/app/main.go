package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"p/internal/database"
	"p/internal/handlers"
	"p/internal/taskService"
	"p/internal/userService"
	"p/internal/web/tasks"
	"p/internal/web/users"
)

func main() {
	database.InitDB()
	if err := database.DB.AutoMigrate(&userService.User{}, &taskService.Task{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	tasksRepo := taskService.NewTaskRepository(database.DB)
	tasksService := taskService.NewService(tasksRepo)
	tasksHandler := handlers.NewHandler(tasksService)

	userRepo := userService.NewUserRepository(database.DB)
	userService := userService.NewUserService(userRepo)

	userHandler := handlers.NewUserHandler(userService)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictTaskHandler := tasks.NewStrictHandler(tasksHandler, nil)
	tasks.RegisterHandlers(e, strictTaskHandler)
	strictUserHandler := users.NewStrictHandler(userHandler, nil)
	users.RegisterHandlers(e, strictUserHandler)
	if err := e.Start(":8080"); err != nil {
		log.Fatal("Error starting server:", err)
	}

}
