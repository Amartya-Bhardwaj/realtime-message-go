package main

import (
	"context"
	"log"

	"github.com/Amartya-Bhardwaj/RealTime-message/db"
	"github.com/Amartya-Bhardwaj/RealTime-message/routes"
	"github.com/Amartya-Bhardwaj/RealTime-message/views"
	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background()
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	db.ConnectDatabase()
	err := db.DB.Ping(ctx,  nil)
	if err != nil {
		log.Fatalf("Error occured: %v", err)
	}
	log.Println("Database Connected")
	
	log.Println(views.GetUsers())

	r.GET("/users", routes.GetAllUsers)
	r.POST("/createUser", routes.CreateUser)
	r.POST("/loginUser", routes.LoginUser)
	r.POST("/conversation", routes.ConversationInUsers)
	r.Run()
}
