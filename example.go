package main

import (
	"log"

	slapi "github.com/sudorandom/softlayer-go/slapi"
	types "github.com/sudorandom/softlayer-go/slapi/types"
)

type Account struct {
	types.SoftLayer_Account

	VirtualGuests []*types.SoftLayer_Virtual_Guest `json:"virtualGuests"`
}

func main() {
	var err error

	client := slapi.Client{}
	client.Endpoint = "https://api.softlayer.com/rest/v3.1/"
	client.Username = "INSERT_USERNAME"
	client.APIKey = "INSERT_API_KEY"

	ctx := client.Context()
	ctx.Mask = `mask[id,virtualGuests]`
	ctx.Service = "SoftLayer_Account"

	ctx.Method = "getObject"

	// Make API call with basic account type
	basicAccount := &types.SoftLayer_Account{}
	err = ctx.Call(basicAccount)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Basic:")
	log.Println(basicAccount.Id)
	// types.SoftLayer_Account has no virtualGuests field
	// log.Println(basicAccount.VirtualGuests)

	// Make API call with extended account type
	extendedAccount := &types.SoftLayer_Account_Extended{}
	err = ctx.Call(extendedAccount)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Extended:")
	log.Println(extendedAccount.Id)
	log.Println(extendedAccount.VirtualGuests)

	// Make API call with custom type based on the basic account type
	customAccount := &Account{}
	err = ctx.Call(customAccount)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Custom Extended:")
	log.Println(customAccount.Id)
	log.Println(customAccount.VirtualGuests)
}
