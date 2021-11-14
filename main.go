package main

import (
	"fmt"
	"io/ioutil"

	"github.com/anish-sinha1/httpclient/httpc"
)

func main() {
	client := httpc.CreateClient()
	res, err := client.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.StatusCode)
	bytes, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(bytes))
}
