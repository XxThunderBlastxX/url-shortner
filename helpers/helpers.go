package helpers

//func CheckUrlInDatabase(url string) bool {
//	now := time.Now()
//	collection := database.ConnectDb().Client.Database("url-shortner").Collection("url")
//
//	var result response
//	err := collection.FindOne(context.TODO(), bson.M{{}})
//	return false
//}

func EnforceHttps(url string) string {
	if url[:4] != "http" || url[:5] != "https" {
		url = "https://" + url
	}
	return url
}
