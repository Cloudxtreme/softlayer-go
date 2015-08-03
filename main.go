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

	ctx := client.Context()
	ctx.Mask = `mask[id,virtualGuests]`
	ctx.Service = "SoftLayer_Account"
	ctx.Method = "getObject"

	account := &types.SoftLayer_Account{}
	err := ctx.Call(account)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(account.VirtualGuests)
	log.Println(account.Id)
	log.Println(account.CompanyName)
}
