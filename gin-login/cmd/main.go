package main

import (
	"context"
	"log"

	http "gin-login/internal/delivery/http/handler"
	"gin-login/internal/delivery/http/middleware"
	"gin-login/internal/infrastructure/mongo"
	"gin-login/internal/usecase"

	"github.com/gin-gonic/gin"
	mongoClient "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongoClient.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("test")

	userRepo := mongo.NewUserRepository(db)
	authUsecase := usecase.NewAuthUsecase(userRepo)
	authHandler := http.NewAuthHandler(authUsecase)

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	r.POST("/login", authHandler.Login)
	r.POST("/register", authHandler.Register)

	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/profile", authHandler.GetProfile)
	}

	r.Run(":8000")
}
