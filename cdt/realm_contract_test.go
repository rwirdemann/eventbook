package cdt

import (
	"bytes"
	"encoding/json"
	"eventbook/core/domain"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"io"
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
	realms := getAllRealms(t, jwt)
	len := len(realms)
	println("LEN", len)

	realm := domain.Realm{
		Name: "runbuddies",
	}
	jsonData, err := json.Marshal(realm)
	if err != nil {
		log.Fatal(err)
	}
	request, error := http.NewRequest("POST", "http://localhost:8000/admin/realms", bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("Authorization", "Bearer "+jwt)

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
}

func getAllRealms(t *testing.T, jwt string) []domain.Realm {
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
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var realms []domain.Realm
	err = json.NewDecoder(resp.Body).Decode(&realms)
	if err != nil {
		log.Fatal(err)
	}
	return realms
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

	var res map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		log.Fatal(err)
	}
	accessToken := fmt.Sprintf("%v", res["access_token"])
	return accessToken
}
