package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
)

// SoftLayer_Account_Shipment_Status - <nil>
type SoftLayer_Account_Shipment_Status struct {

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// Id - <nil>
	Id int `json:"id"`

	// KeyName - <nil>
	KeyName string `json:"keyName"`

	// Name - <nil>
	Name string `json:"name"`
}

// GetObject - <nil>
func (softlayer_account_shipment_status *SoftLayer_Account_Shipment_Status) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Account_Shipment_Status, error) {
	var returnValue *SoftLayer_Account_Shipment_Status
	return returnValue, nil
}
