package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_User_Customer_Prospect_Type - <nil>
type SoftLayer_User_Customer_Prospect_Type struct {

	// Id - <nil>
	Id int `json:"id"`

	// KeyName - <nil>
	KeyName string `json:"keyName"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate"`

	// Name - <nil>
	Name string `json:"name"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// Description - <nil>
	Description string `json:"description"`
}

func (softlayer_user_customer_prospect_type *SoftLayer_User_Customer_Prospect_Type) String() string {
	return "SoftLayer_User_Customer_Prospect_Type"
}
