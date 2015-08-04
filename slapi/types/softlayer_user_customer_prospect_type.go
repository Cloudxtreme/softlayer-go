package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_User_Customer_Prospect_Type - <nil>
type SoftLayer_User_Customer_Prospect_Type struct {

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Description - <nil>
	Description string `json:"description,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// KeyName - <nil>
	KeyName string `json:"keyName,omitempty"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// Name - <nil>
	Name string `json:"name,omitempty"`
}

func (softlayer_user_customer_prospect_type *SoftLayer_User_Customer_Prospect_Type) String() string {
	return "SoftLayer_User_Customer_Prospect_Type"
}
