package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Subnet_Registration_Status - Subnet Registration Status objects describe the
// current status of a subnet registration. The standard values for these objects are as follows: -
// Indicates that the registration object is new and has yet to be submitted to the - Indicates that
// the registration object has been submitted to the RIR and is awaiting response - Indicates that the
// RIR action has completed - Indicates that the registration object has been gracefully removed is no
// longer valid - Indicates that the registration object has been abruptly removed is no longer valid
type SoftLayer_Network_Subnet_Registration_Status struct {

	// KeyName - no documentation
	KeyName string `json:"keyName"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate"`

	// Name - no documentation
	Name string `json:"name"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// Id - no documentation
	Id int `json:"id"`
}

func (softlayer_network_subnet_registration_status *SoftLayer_Network_Subnet_Registration_Status) String() string {
	return "SoftLayer_Network_Subnet_Registration_Status"
}
