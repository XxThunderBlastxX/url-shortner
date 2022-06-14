package controller

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"url-shortner/database"
)

func ResolveUrl(c *fiber.Ctx) error {
	urlId := c.Params("urlId")

	collection := database.ConnectDb().Client.Database("url-shortner").Collection("url")

	var result response
	err := collection.FindOne(context.Background(), bson.D{{"urlId", urlId}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return err
		}
		return err
	}

	return c.Redirect(result.Url)
}
