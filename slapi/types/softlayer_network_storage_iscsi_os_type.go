package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Storage_Iscsi_OS_Type - <nil>
type SoftLayer_Network_Storage_Iscsi_OS_Type struct {

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Description - no documentation
	Description string `json:"description,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// KeyName - no documentation
	KeyName string `json:"keyName,omitempty"`
}

func (softlayer_network_storage_iscsi_os_type *SoftLayer_Network_Storage_Iscsi_OS_Type) String() string {
	return "SoftLayer_Network_Storage_Iscsi_OS_Type"
}
