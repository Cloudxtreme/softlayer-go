package main

import (
	"log"

	slapi "github.com/sudorandom/softlayer-go/slapi"
	types "github.com/sudorandom/softlayer-go/slapi/types"
)

func main() {

	client := slapi.Client{
		Endpoint: "https://api.softlayer.com/rest/v3.1/",
		Username: "INSERT_USERNAME",
		APIKey:   "INSERT_API_KEY",
	}

	makeBasicAPICall(client)
	makeExtendedAPICall(client)
	makeCustomAPICall(client)
}

// makeBasicAPICall demonstrates using a basic type
func makeBasicAPICall(client slapi.Client) {
	req := slapi.Request{
		Service: "SoftLayer_Account",
		Method:  "getObject",
		Mask:    `mask[id,companyName]`,
	}

	// Make API call with basic account type
	basicAccount := &types.SoftLayer_Account{}
	err := client.Call(req, basicAccount)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Basic:")
	log.Println(basicAccount.Id)
	log.Println(basicAccount.CompanyName)
}

// makeExtendedAPICall demonstrates using the extended type
func makeExtendedAPICall(client slapi.Client) {
	req := slapi.Request{
		Service: "SoftLayer_Account",
		Method:  "getObject",
		Mask:    `mask[id,virtualGuests]`,
	}

	// Make API call with extended account type
	extendedAccount := &types.SoftLayer_Account_Extended{}
	err := client.Call(req, extendedAccount)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Extended:")
	log.Println(extendedAccount.Id)
	log.Println(extendedAccount.VirtualGuests)
}

type CustomAccount struct {
	types.SoftLayer_Account

	VirtualGuests []*types.SoftLayer_Virtual_Guest `json:"virtualGuests"`
}

// makeCustomAPICall demonstrates using a user-defined type
func makeCustomAPICall(client slapi.Client) {

	req := slapi.Request{
		Service: "SoftLayer_Account",
		Method:  "getObject",
		Mask:    `mask[id,virtualGuests,companyName]`,
	}

	// Make API call with custom type based on the basic account type
	customAccount := &CustomAccount{}
	err := client.Call(req, customAccount)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Custom Extended:")
	log.Println(customAccount.Id)
	log.Println(customAccount.CompanyName)
	log.Println(customAccount.VirtualGuests)
}
