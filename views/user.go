package views

import (
	"context"
	"log"

	"github.com/Amartya-Bhardwaj/RealTime-message/db"
	"github.com/Amartya-Bhardwaj/RealTime-message/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUsers() []models.User{ 
	ctx := context.Background()
	var users []models.User
	collection := db.DB.Database("test").Collection("users")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Println(err)
	}
	defer cursor.Close(ctx)
	if err := cursor.All(ctx, &users); err != nil {
		log.Println(err)
	}
	return users
}

func CreateUser(user interface{}) {
	ctx := context.Background()
	collection, err := db.DB.Database("test").Collection("users").InsertOne(ctx, user)
	if err != nil {
		log.Println(err)
	}
	log.Println("Collection: {}", collection)
}

func LoginUser (creds models.Credential) models.User{
	ctx := context.Background()
	var user models.User
	collection := db.DB.Database("test").Collection("users")
	err := collection.FindOne(ctx, bson.M{"email": creds.Email}).Decode(&user)
	if err != nil {
		log.Println("Error Occured: {}", err)
	}
	return user
}

func GetIdByEmail (email string) string{
	ctx := context.Background()
	var user models.User
	collection := db.DB.Database("test").Collection("users")
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		log.Println("Error occured: ", err)
	}

	return string(user.ID.Hex())
}

func GetUserByEmail (email string) models.User{
	ctx := context.Background()
	var user models.User
	collection := db.DB.Database("test").Collection("users")
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		log.Println("Error occured: ", err)
	}

	return user
}

func UpdateNonPremiumCount(count int, email string) {
	ctx := context.Background()
	collection := db.DB.Database("test").Collection("users")
	update := bson.M{
		"$set": bson.M{
			"nonPremiumCount": count,
		},
	}
	res,err := collection.UpdateOne(ctx, bson.M{"email": email}, update)
	if err != nil {
		log.Println("Error occured: ", err)
	}
	log.Println("Successfully updated the count", res.ModifiedCount)
}