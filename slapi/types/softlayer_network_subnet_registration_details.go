package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Subnet_Registration_Details - The SoftLayer_Network_Subnet_Registration_Details
// objects are used to relate [[SoftLayer_Account_Regional_Registry_Detail]] objects to a
// [[SoftLayer_Network_Subnet_Registration]] object. This allows for easy reuse of registration
// details. It is important to note that only one detail object per type may be associated to a
// registration object.
type SoftLayer_Network_Subnet_Registration_Details struct {

	// DetailId - Numeric ID of the related [[SoftLayer_Account_Regional_Registry_Detail]] object
	DetailId int `json:"detailId"`

	// Id - no documentation
	Id int `json:"id"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate"`

	// RegistrationId - Numeric ID of the related [[SoftLayer_Network_Subnet_Registration]] object
	RegistrationId int `json:"registrationId"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`
}

func (softlayer_network_subnet_registration_details *SoftLayer_Network_Subnet_Registration_Details) String() string {
	return "SoftLayer_Network_Subnet_Registration_Details"
}

// SoftLayer_Network_Subnet_Registration_Details_Extended is SoftLayer_Network_Subnet_Registration_Details with all maskable types.
type SoftLayer_Network_Subnet_Registration_Details_Extended struct {
	SoftLayer_Network_Subnet_Registration_Details

	// Detail - The related [[SoftLayer_Account_Regional_Registry_Detail|detail object]].
	Detail *SoftLayer_Account_Regional_Registry_Detail `json:"detail"`

	// Registration - The related [[SoftLayer_Network_Subnet_Registration|registration object]].
	Registration *SoftLayer_Network_Subnet_Registration `json:"registration"`
}

func (softlayer_network_subnet_registration_details *SoftLayer_Network_Subnet_Registration_Details_Extended) String() string {
	return "SoftLayer_Network_Subnet_Registration_Details"
}
