package helper

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckEmail(email string) (error){
	result := bson.M{}
	if err := godotenv.Load(".env");err != nil {
		return err
	}
	reqUrl := "https://api.emailable.com/v1/verify?"
	url := bytes.Buffer{}
	url.WriteString(reqUrl)
	url.WriteString("&email=")
	url.WriteString(email)
	url.WriteString("&api_key=")
	url.WriteString(os.Getenv("EMAIL_API"))
	resp,err := http.Get(url.String())
	if err != nil{
		return err
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&result)
	return nil
}
