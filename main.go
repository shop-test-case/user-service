package main

import (
	"user-service/config"
	"user-service/controller"
	"user-service/database"
	"user-service/handler"
	"user-service/middleware"
	"user-service/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	db := database.Connect(cfg)

	var userRepo repository.IUserRepo = &repository.UserRepo{DB: db}
	userCtrl := &controller.UserController{Repo: userRepo, JWTSecret: cfg.JWTSecret}
	userHandler := &handler.UserHandler{Ctrl: userCtrl}

	r := gin.Default()
	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)

	auth := r.Group("/auth")
	auth.Use(middleware.JWT(cfg.JWTSecret))

	if err := r.Run(":" + cfg.Port); err != nil {
		panic("failed to start server: " + err.Error())
	}
}
