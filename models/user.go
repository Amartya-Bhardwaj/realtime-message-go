package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	IsPremium bool				`bson:"isPremium"`
	NonPremiumCount int 		`bson:"nonPremiumCount"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}

// Conversation represents a chat conversation
type Conversation struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Participants []primitive.ObjectID `bson:"participants"`
	LastMessage  Message            `bson:"lastMessage"`
	CreatedAt    time.Time          `bson:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt"`
}

// Message represents a chat message
type Message struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	ConversationID primitive.ObjectID `bson:"conversationId"`
	Sender        primitive.ObjectID `bson:"sender"`
	Receiver      primitive.ObjectID `bson:"receiver"`
	Text          string             `bson:"text"`
	Timestamp     time.Time          `bson:"timestamp"`
}
