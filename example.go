package main

import (
	"log"

	slapi "github.com/sudorandom/softlayer-go/slapi"
	types "github.com/sudorandom/softlayer-go/slapi/types"
)

func main() {

	client := slapi.Client{}
	client.Endpoint = "https://api.softlayer.com/rest/v3.1/"
	client.Username = "INSERT_USERNAME"
	client.APIKey = "INSERT_API_KEY"

	makeBasicAPICall(client)
	makeExtendedAPICall(client)
	makeCustomAPICall(client)
}

func makeBasicAPICall(client slapi.Client) {
	ctx := client.Request()
	ctx.Mask = `mask[id,companyName]`
	ctx.Service = "SoftLayer_Account"
	ctx.Method = "getObject"

	// Make API call with basic account type
	basicAccount := &types.SoftLayer_Account{}
	err := ctx.Do(basicAccount)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Basic:")
	log.Println(basicAccount.Id)
	log.Println(basicAccount.CompanyName)
}

func makeExtendedAPICall(client slapi.Client) {
	ctx := client.Request()
	ctx.Mask = `mask[id,virtualGuests]`
	ctx.Service = "SoftLayer_Account"
	ctx.Method = "getObject"

	// Make API call with extended account type
	extendedAccount := &types.SoftLayer_Account_Extended{}
	err := ctx.Do(extendedAccount)
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

func makeCustomAPICall(client slapi.Client) {

	ctx := client.Request()
	ctx.Mask = `mask[id,virtualGuests,companyName]`
	ctx.Service = "SoftLayer_Account"
	ctx.Method = "getObject"

	// Make API call with custom type based on the basic account type
	customAccount := &CustomAccount{}
	err := ctx.Do(customAccount)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Custom Extended:")
	log.Println(customAccount.Id)
	log.Println(customAccount.CompanyName)
	log.Println(customAccount.VirtualGuests)
}
