package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Subnet_Registration - The subnet registration data type contains general
// information relating to a single subnet registration instance. These registration instances can be
// updated to reflect changes, and will record the changes in the
// [[SoftLayer_Network_Subnet_Registration_Event|events]].
type SoftLayer_Network_Subnet_Registration struct {

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// AccountId - The registration object's associated [[SoftLayer_Account|account]] id
	AccountId int `json:"accountId"`

	// Cidr - no documentation
	Cidr int `json:"cidr"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// DetailReferenceCount - A count of the cross-reference records that tie the
	// [[SoftLayer_Account_Regional_Registry_Detail]] objects to the registration object.
	DetailReferenceCount uint64 `json:"detailReferenceCount"`

	// DetailReferences - The cross-reference records that tie the
	// [[SoftLayer_Account_Regional_Registry_Detail]] objects to the registration object.
	DetailReferences []*SoftLayer_Network_Subnet_Registration_Details `json:"detailReferences"`

	// EventCount - no documentation
	EventCount uint64 `json:"eventCount"`

	// Events - no documentation
	Events []*SoftLayer_Network_Subnet_Registration_Event `json:"events"`

	// Id - no documentation
	Id int `json:"id"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate"`

	// NetworkDetail - no documentation
	NetworkDetail *SoftLayer_Account_Regional_Registry_Detail `json:"networkDetail"`

	// NetworkHandle - The RIR-specific handle or name of the registered subnet. This field is read-only.
	NetworkHandle string `json:"networkHandle"`

	// NetworkIdentifier - no documentation
	NetworkIdentifier string `json:"networkIdentifier"`

	// PersonDetail - no documentation
	PersonDetail *SoftLayer_Account_Regional_Registry_Detail `json:"personDetail"`

	// RegionalInternetRegistry - no documentation
	RegionalInternetRegistry *SoftLayer_Network_Regional_Internet_Registry `json:"regionalInternetRegistry"`

	// RegionalInternetRegistryHandle - The RIR handle that this registration object belongs to. This field
	// may not be populated until the registration is complete.
	RegionalInternetRegistryHandle *SoftLayer_Account_Rwhois_Handle `json:"regionalInternetRegistryHandle"`

	// RegionalInternetRegistryHandleId - The registration object's associated
	// [[SoftLayer_Account_Rwhois_Handle|RIR handle]] id
	RegionalInternetRegistryHandleId int `json:"regionalInternetRegistryHandleId"`

	// RegionalInternetRegistryId - The registration object's associated
	// [[SoftLayer_Network_Regional_Internet_Registry|RIR]] id
	RegionalInternetRegistryId int `json:"regionalInternetRegistryId"`

	// Status - no documentation
	Status *SoftLayer_Network_Subnet_Registration_Status `json:"status"`

	// StatusId - The registration object's associated
	// [[SoftLayer_Network_Subnet_Registration_Status|status]] id
	StatusId int `json:"statusId"`

	// Subnet - no documentation
	Subnet *SoftLayer_Network_Subnet `json:"subnet"`
}

// ClearRegistration - This method will initiate the removal of a subnet registration.
func (softlayer_network_subnet_registration *SoftLayer_Network_Subnet_Registration) ClearRegistration() (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// CreateObject - This method will create a new SoftLayer_Network_Subnet_Registration object. Input -
// [[SoftLayer_Network_Subnet_Registration (type)|SoftLayer_Network_Subnet_Registration]] The base
// address of the [[SoftLayer_Network_Subnet|subnet]] being registered. This can be derived directly
// from the SoftLayer_Network_Subnet object's networkIdentifier property. The prefix of the
// [[SoftLayer_Network_Subnet|subnet]] being registered. This can be derived directly from the
// SoftLayer_Network_Subnet object's cidr property.
func (softlayer_network_subnet_registration *SoftLayer_Network_Subnet_Registration) CreateObject(templateObject SoftLayer_Network_Subnet_Registration) (*SoftLayer_Network_Subnet_Registration, error) {
	var returnValue *SoftLayer_Network_Subnet_Registration
	return returnValue, nil
}

// EditObject - This method will edit an existing SoftLayer_Network_Subnet_Registration object. For
// more detail, see [[SoftLayer_Network_Subnet_Registration::createObject|createObject]].
func (softlayer_network_subnet_registration *SoftLayer_Network_Subnet_Registration) EditObject(templateObject SoftLayer_Network_Subnet_Registration) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditRegistrationAttachedDetails - This method modifies a single registration by modifying the
// current [[SoftLayer_Network_Subnet_Registration_Details]] objects that are linked to that
// registration.
func (softlayer_network_subnet_registration *SoftLayer_Network_Subnet_Registration) EditRegistrationAttachedDetails(personObjectSkeleton SoftLayer_Network_Subnet_Registration_Details, networkObjectSkeleton SoftLayer_Network_Subnet_Registration_Details) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_network_subnet_registration *SoftLayer_Network_Subnet_Registration) GetObject() (*SoftLayer_Network_Subnet_Registration, error) {
	var returnValue *SoftLayer_Network_Subnet_Registration
	return returnValue, nil
}
