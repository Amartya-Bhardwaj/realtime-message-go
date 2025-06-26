package main

import (
	"context"
	"log"

	"github.com/Amartya-Bhardwaj/RealTime-message/db"
	"github.com/Amartya-Bhardwaj/RealTime-message/routes"
	"github.com/Amartya-Bhardwaj/RealTime-message/views"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background()
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	r.Use(cors.New(config))
	gin.SetMode(gin.DebugMode)
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
	r.POST("/payment/order", routes.CreateOrder)
	r.POST("/order/events", routes.OrderWebhookEvent)
	r.Run()
}
