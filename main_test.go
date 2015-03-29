package main

import (
	"encoding/json"
	"fmt"
	sl "go-softlayer/gen"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	username = ""
	apiKey   = ""
)

// go run ./bin/sl-generate-types/*.go
func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.softlayer.com/rest/v3.1/SoftLayer_Account.json", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(username, apiKey)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))

	var account sl.SoftLayer_Account
	err = json.Unmarshal(body, &account)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%#v\n", account)
	fmt.Println(account.Id)
	fmt.Println(account.Address1)
}
