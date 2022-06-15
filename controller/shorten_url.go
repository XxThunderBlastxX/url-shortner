package controller

import (
	"context"
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"os"
	"time"
	"url-shortner/database"
	"url-shortner/helpers"
)

type request struct {
	Url       string `json:"url" bson:"url"`
	CustomUrl string `json:"customUrl" bson:"customUrl"`
	//Expiry    time.Duration `json:"expiry"`
}

type response struct {
	Url         string    `json:"url" bson:"url"`
	ShortUrl    string    `json:"shortUrl" bson:"shortUrl"`
	UrlId       string    `json:"urlId" bson:"urlId"`
	CreatedTime time.Time `json:"createdTime" bson:"createdTime"`
	//Expiry          time.Duration `json:"expiry"`
	//XRateRemaining  int           `json:"rate_limit"`
	//XRateLimitReset time.Duration `json:"rate_limit_reset"`
}

func ShortenUrl(c *fiber.Ctx) error {
	body := new(request)

	// check for the incoming request body
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	//Check if the incoming link is correct URL
	if !govalidator.IsURL(body.Url) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid URL !!"})
	}

	// enforce https
	// all url will be converted to https before storing in database
	body.Url = helpers.EnforceHttps(body.Url)

	//Check the URL already exist or not
	//if true = duplicate url exist else false= does not exit ... all Ok to go :)
	UrlVal, condVal := helpers.CheckUrlInDatabase(body.Url)
	if condVal {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Url already exist !!", "url": body.Url, "short-url": UrlVal})
	}

	//check if user has given any customUrl or not
	var id string
	if body.CustomUrl == "" {
		id = uuid.New().String()[:6]
	} else {
		if !helpers.CheckGivenIdInDatabase(body.CustomUrl) {
			id = body.CustomUrl
		} else {
			return c.Status(fiber.StatusFound).JSON(fiber.Map{"message": "short url already exist with same custom-url"})
		}
	}

	//Get MongoDB collection
	collection := database.ConnectDb().Client.Database("url-shortner").Collection("url")

	res := response{Url: body.Url, ShortUrl: "", UrlId: id, CreatedTime: time.Now()}

	res.ShortUrl = os.Getenv("DOMAIN") + "/" + id

	_, err := collection.InsertOne(context.TODO(), res)
	if err != nil {
		return err
	}
	defer func() {
		if err = database.ConnectDb().Client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	return c.Status(fiber.StatusOK).JSON(res)
}
