package cdt

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"testing"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func TestGetAllRealms(t *testing.T) {
	jwt := login()
	c := http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8000/admin/realms", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", "Bearer "+jwt)
	resp, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("Body : %s", body)
}

func login() string {
	data := url.Values{
		"username":   {os.Getenv("USERNAME")},
		"password":   {os.Getenv("PASSWORD")},
		"client_id":  {os.Getenv("CLIENT_ID")},
		"grant_type": {"password"},
	}
	resp, err := http.PostForm("http://localhost:8080/realms/wingding/protocol/openid-connect/token", data)
	if err != nil {
		log.Fatal(err)
	}

	var res map[string]string
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		log.Fatal(err)
	}
	return res["access_token"]
}
