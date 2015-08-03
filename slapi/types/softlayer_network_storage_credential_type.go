package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Storage_Credential_Type - <nil>
type SoftLayer_Network_Storage_Credential_Type struct {

	// KeyName - no documentation
	KeyName string `json:"keyName"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// Name - no documentation
	Name string `json:"name"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Description - no documentation
	Description string `json:"description"`
}

func (softlayer_network_storage_credential_type *SoftLayer_Network_Storage_Credential_Type) String() string {
	return "SoftLayer_Network_Storage_Credential_Type"
}
