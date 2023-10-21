package helpers

import (
	"context"
	"time"
	"url-shortner/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type response struct {
	Url         string    `json:"url" bson:"url"`
	ShortUrl    string    `json:"shortUrl" bson:"shortUrl"`
	UrlId       string    `json:"urlId" bson:"urlId"`
	CreatedTime time.Time `json:"createdTime" bson:"createdTime"`
	//Expiry          time.Duration `json:"expiry"`
	//XRateRemaining  int           `json:"rate_limit"`
	//XRateLimitReset time.Duration `json:"rate_limit_reset"`
}

func CheckUrlInDatabase(url string) (string, bool) {

	collection := database.ConnectDb().Client.Database("url-shortner").Collection("url")

	var result response
	err := collection.FindOne(context.Background(), bson.D{{"url", url}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", false
		}
	}
	return result.ShortUrl, true
}

func CheckGivenIdInDatabase(short string) bool {
	collection := database.ConnectDb().Client.Database("url-shortner").Collection("url")

	var result response
	err := collection.FindOne(context.Background(), bson.D{{"urlId", short}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false
		}
	}
	return true
}

func EnforceHttps(url string) string {
	if url[:4] != "http" || url[:5] != "https" {
		url = "https://" + url
	}
	return url
}
