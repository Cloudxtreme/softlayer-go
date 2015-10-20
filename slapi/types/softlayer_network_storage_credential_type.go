package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Storage_Credential_Type - <nil>
type SoftLayer_Network_Storage_Credential_Type struct {

	// KeyName - no documentation
	KeyName string `json:"keyName,omitempty"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Description - no documentation
	Description string `json:"description,omitempty"`
}

func (softlayer_network_storage_credential_type *SoftLayer_Network_Storage_Credential_Type) String() string {
	return "SoftLayer_Network_Storage_Credential_Type"
}
