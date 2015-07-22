package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
)

// SoftLayer_Network_Subnet_Registration_Status - Subnet Registration Status objects describe the
// current status of a subnet registration. The standard values for these objects are as follows: -
// Indicates that the registration object is new and has yet to be submitted to the - Indicates that
// the registration object has been submitted to the RIR and is awaiting response - Indicates that the
// RIR action has completed - Indicates that the registration object has been gracefully removed is no
// longer valid - Indicates that the registration object has been abruptly removed is no longer valid
type SoftLayer_Network_Subnet_Registration_Status struct {

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// Id - no documentation
	Id int `json:"id"`

	// KeyName - no documentation
	KeyName string `json:"keyName"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate"`

	// Name - no documentation
	Name string `json:"name"`
}

// GetAllObjects - <nil>
func (softlayer_network_subnet_registration_status *SoftLayer_Network_Subnet_Registration_Status) GetAllObjects(ctx *slapi.RequestContext) ([]*SoftLayer_Network_Subnet_Registration_Status, error) {
	var returnValue []*SoftLayer_Network_Subnet_Registration_Status
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_network_subnet_registration_status *SoftLayer_Network_Subnet_Registration_Status) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Network_Subnet_Registration_Status, error) {
	var returnValue *SoftLayer_Network_Subnet_Registration_Status
	return returnValue, nil
}
