package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Billing_Currency - <nil>
type SoftLayer_Billing_Currency struct {

	// Id - <nil>
	Id int `json:"id"`

	// KeyName - <nil>
	KeyName string `json:"keyName"`

	// Name - <nil>
	Name string `json:"name"`
}

// GetAllObjects - <nil>
func (softlayer_billing_currency *SoftLayer_Billing_Currency) GetAllObjects(ctx *slapi.RequestContext) ([]*SoftLayer_Billing_Currency, error) {
	var returnValue []*SoftLayer_Billing_Currency
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_billing_currency *SoftLayer_Billing_Currency) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Billing_Currency, error) {
	var returnValue *SoftLayer_Billing_Currency
	return returnValue, nil
}

// GetPrice - <nil>
func (softlayer_billing_currency *SoftLayer_Billing_Currency) GetPrice(ctx *slapi.RequestContext, price float32, formatOptions SoftLayer_Container_Billing_Currency_Format) (string, error) {
	var returnValue string
	return returnValue, nil
}
