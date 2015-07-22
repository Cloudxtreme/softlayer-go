package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
)

// SoftLayer_Network_Subnet_Registration_Details - The SoftLayer_Network_Subnet_Registration_Details
// objects are used to relate [[SoftLayer_Account_Regional_Registry_Detail]] objects to a
// [[SoftLayer_Network_Subnet_Registration]] object. This allows for easy reuse of registration
// details. It is important to note that only one detail object per type may be associated to a
// registration object.
type SoftLayer_Network_Subnet_Registration_Details struct {

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// Detail - The related [[SoftLayer_Account_Regional_Registry_Detail|detail object]].
	Detail *SoftLayer_Account_Regional_Registry_Detail `json:"detail"`

	// DetailId - Numeric ID of the related [[SoftLayer_Account_Regional_Registry_Detail]] object
	DetailId int `json:"detailId"`

	// Id - no documentation
	Id int `json:"id"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate"`

	// Registration - The related [[SoftLayer_Network_Subnet_Registration|registration object]].
	Registration *SoftLayer_Network_Subnet_Registration `json:"registration"`

	// RegistrationId - Numeric ID of the related [[SoftLayer_Network_Subnet_Registration]] object
	RegistrationId int `json:"registrationId"`
}

// CreateObject - This method will create a new SoftLayer_Network_Subnet_Registration_Details object.
// Input - [[SoftLayer_Network_Subnet_Registration_Details
// (type)|SoftLayer_Network_Subnet_Registration_Details]] The numeric ID of the
// [[SoftLayer_Account_Regional_Registry_Detail|detail]] object to relate. The numeric ID of the
// [[SoftLayer_Network_Subnet_Registration|registration]] object to relate.
func (softlayer_network_subnet_registration_details *SoftLayer_Network_Subnet_Registration_Details) CreateObject(commonOptions *slapi.CommonOptions, templateObject SoftLayer_Network_Subnet_Registration_Details) (*SoftLayer_Network_Subnet_Registration_Details, error) {
	var returnValue *SoftLayer_Network_Subnet_Registration_Details
	return returnValue, nil
}

// DeleteObject - This method will delete an existing SoftLayer_Account_Regional_Registry_Detail
// object.
func (softlayer_network_subnet_registration_details *SoftLayer_Network_Subnet_Registration_Details) DeleteObject(commonOptions *slapi.CommonOptions) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_network_subnet_registration_details *SoftLayer_Network_Subnet_Registration_Details) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Network_Subnet_Registration_Details, error) {
	var returnValue *SoftLayer_Network_Subnet_Registration_Details
	return returnValue, nil
}
