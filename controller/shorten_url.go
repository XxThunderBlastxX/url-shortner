package controller

import (
	"context"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"os"
	"time"
	"url-shortner/database"
	"url-shortner/helpers"
)

type request struct {
	Url       string `json:"url"`
	CustomUrl string `json:"customId"`
	//Expiry    time.Duration `json:"expiry"`
}

type response struct {
	Url         string    `json:"url"`
	ShortUrl    string    `json:"shortUrl"`
	CreatedTime time.Time `json:"createdTime"`
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

	//Check the URL already exist or not
	//if false = duplicate url exist else true = does not exit .... all ok to go :)
	//if !helpers.CheckUrlInDatabase(body.Url) {
	//	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "sharam bech khaye ho ka ??"})
	//}

	// enforce https
	// all url will be converted to https before storing in database
	body.Url = helpers.EnforceHttps(body.Url)
	fmt.Println(body.Url)

	//check if user has given any customUrl or not
	var id string
	if body.CustomUrl == "" {
		id = uuid.New().String()[:6]
	} else {
		id = body.CustomUrl
	}

	//Get MongoDB collection
	collection := database.ConnectDb().Client.Database("url-shortner").Collection("url")

	res := response{Url: body.Url, ShortUrl: "", CreatedTime: time.Now()}

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
