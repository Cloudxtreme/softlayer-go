package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Network_Message_Delivery - <nil>
type SoftLayer_Network_Message_Delivery struct {

	// Account - The SoftLayer customer account that a network message delivery account belongs to.
	Account *SoftLayer_Account `json:"account"`

	// AccountId - <nil>
	AccountId int `json:"accountId"`

	// BillingItem - The billing item for a network message delivery account.
	BillingItem *SoftLayer_Billing_Item `json:"billingItem"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// Id - <nil>
	Id int `json:"id"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate"`

	// Password - <nil>
	Password string `json:"password"`

	// Type - The message delivery type of a network message delivery account.
	Type *SoftLayer_Network_Message_Delivery_Type `json:"type"`

	// TypeId - <nil>
	TypeId int `json:"typeId"`

	// Username - <nil>
	Username string `json:"username"`

	// Vendor - no documentation
	Vendor *SoftLayer_Network_Message_Delivery_Vendor `json:"vendor"`

	// VendorId - <nil>
	VendorId int `json:"vendorId"`
}

func (softlayer_network_message_delivery *SoftLayer_Network_Message_Delivery) String() string {
	return "SoftLayer_Network_Message_Delivery"
}

// EditObject - <nil>
func (softlayer_network_message_delivery *SoftLayer_Network_Message_Delivery) EditObject(ctx *slapi.RequestContext, templateObject SoftLayer_Network_Message_Delivery) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_network_message_delivery *SoftLayer_Network_Message_Delivery) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Network_Message_Delivery, error) {
	var returnValue *SoftLayer_Network_Message_Delivery
	return returnValue, nil
}
