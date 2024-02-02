package views

import (
	"context"
	"log"

	"github.com/Amartya-Bhardwaj/RealTime-message/db"
	"github.com/Amartya-Bhardwaj/RealTime-message/models"
)

// Views function for conversation between users
func ConversationInUsers(msg models.Conversation) {
	ctx := context.Background()
	collection := db.DB.Database("test").Collection("conversations")
	_, err := collection.InsertOne(ctx, msg)
	if err != nil {
		log.Println("Error occured: ", err)
	}
}