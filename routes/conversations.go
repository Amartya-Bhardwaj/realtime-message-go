package routes

import (
	"log"
	"net/http"
	"time"

	"github.com/Amartya-Bhardwaj/RealTime-message/middleware"
	"github.com/Amartya-Bhardwaj/RealTime-message/models"
	"github.com/Amartya-Bhardwaj/RealTime-message/views"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ConversationJSON struct {
	SenderEmail string `json:"senderEmail"`
	ReceiverEmail string `json:"receiverEmail"`
	Message string `json:"message"`
}

func ConversationInUsers(c *gin.Context) {
	conversationJSON := ConversationJSON
	err := c.ShouldBindJSON(&conversationJSON)
	if err != nil {
		log.Println("Error occured while binding json payload: ", err)
		return
	}
	
	token := c.GetHeader("authorization")
	if token == "" {
		log.Println("JWT Token is empty")
		c.JSON(403, "Invalid Token")
		return
	}

	if middleware.VerifyToken(token) == false {
		log.Println("Invalid Token")
		c.JSON(403, "Invalid Token")
		return
	}
	//Check If the User is Premium or not
	email := string(conversationJSON.SenderEmail)
	var user = views.GetUserByEmail(email)
	if !user.IsPremium {
		if user.NonPremiumCount > 0 {
			changedCount := user.NonPremiumCount - 1
			views.UpdateNonPremiumCount(changedCount, email)
			sender_id := convertStringtoHex(views.GetIdByEmail(conversationJSON.SenderEmail))
	
			receiver_id := convertStringtoHex(views.GetIdByEmail(conversationJSON.ReceiverEmail))

			conversation := models.Conversation {
				Participants: []primitive.ObjectID{receiver_id},
				LastMessage: models.Message{
					Sender: sender_id,
					Text: conversationJSON.Message,
					Timestamp: time.Now(),
				},
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}

			views.ConversationInUsers(conversation)
			c.JSON(http.StatusOK, conversation)
		} else {
			response := map[string] interface{} {"desc": "Not a premium user"}
			c.JSON(http.StatusNotAcceptable, response)
		}
	} else {
		sender_id := convertStringtoHex(views.GetIdByEmail(conversationJSON.SenderEmail))
	
		receiver_id := convertStringtoHex(views.GetIdByEmail(conversationJSON.ReceiverEmail))

		conversation := models.Conversation {
			Participants: []primitive.ObjectID{receiver_id},
			LastMessage: models.Message{
				Sender: sender_id,
				Text: conversationJSON.Message,
				Timestamp: time.Now(),
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		views.ConversationInUsers(conversation)
		c.JSON(http.StatusOK, conversation)
	}
}

func convertStringtoHex(input string) primitive.ObjectID{
	objectID, err := primitive.ObjectIDFromHex(input)
	if err != nil {
		log.Println("Error converting string to ObjectID:", err)
	}
	return objectID
}