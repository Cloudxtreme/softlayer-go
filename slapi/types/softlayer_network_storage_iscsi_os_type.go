package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Network_Storage_Iscsi_OS_Type - <nil>
type SoftLayer_Network_Storage_Iscsi_OS_Type struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Description - no documentation
	Description string `json:"description"`

	// Id - no documentation
	Id int `json:"id"`

	// KeyName - no documentation
	KeyName string `json:"keyName"`

	// Name - no documentation
	Name string `json:"name"`
}

func (softlayer_network_storage_iscsi_os_type *SoftLayer_Network_Storage_Iscsi_OS_Type) String() string {
	return "SoftLayer_Network_Storage_Iscsi_OS_Type"
}

// GetAllObjects - no documentation
func (softlayer_network_storage_iscsi_os_type *SoftLayer_Network_Storage_Iscsi_OS_Type) GetAllObjects(ctx *slapi.RequestContext) ([]*SoftLayer_Network_Storage_Iscsi_OS_Type, error) {
	var returnValue []*SoftLayer_Network_Storage_Iscsi_OS_Type
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_network_storage_iscsi_os_type *SoftLayer_Network_Storage_Iscsi_OS_Type) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Network_Storage_Iscsi_OS_Type, error) {
	var returnValue *SoftLayer_Network_Storage_Iscsi_OS_Type
	return returnValue, nil
}
