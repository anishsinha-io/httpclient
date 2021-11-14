package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/anish-sinha1/httpclient/httpc"
)

func createClient() httpc.HttpClient {
	client := httpc.CreateClient()
	commonHeaders := make(http.Header)
	commonHeaders.Set("Authorization", "Bearer <token>")
	client.SetHeaders(commonHeaders)
	return client
}

var githubHttpClient = createClient()

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {
	headers := make(http.Header)
	headers.Set("Authorization", "Bearer <token>")
	res, err := githubHttpClient.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.StatusCode)
	bytes, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(bytes))
}
