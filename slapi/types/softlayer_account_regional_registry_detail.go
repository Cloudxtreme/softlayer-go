package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
)

// SoftLayer_Account_Regional_Registry_Detail - <nil>
type SoftLayer_Account_Regional_Registry_Detail struct {

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// AccountId - The detail object's associated [[SoftLayer_Account|account]] id
	AccountId int `json:"accountId"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// DetailCount - A count of references to the [[SoftLayer_Network_Subnet_Registration|registration
	// objects]] that consume this detail object.
	DetailCount uint64 `json:"detailCount"`

	// DetailType - no documentation
	DetailType *SoftLayer_Account_Regional_Registry_Detail_Type `json:"detailType"`

	// DetailTypeId - The detail object's associated
	// [[SoftLayer_Account_Regional_Registry_Detail_Type|type]] id
	DetailTypeId int `json:"detailTypeId"`

	// Details - References to the [[SoftLayer_Network_Subnet_Registration|registration objects]] that
	// consume this detail object.
	Details []*SoftLayer_Network_Subnet_Registration_Details `json:"details"`

	// Id - no documentation
	Id int `json:"id"`

	// ModifyDate - The date and time the detail object was last modified
	ModifyDate *time.Time `json:"modifyDate"`

	// Properties - The individual properties that define this detail object's values.
	Properties []*SoftLayer_Account_Regional_Registry_Detail_Property `json:"properties"`

	// PropertyCount - A count of the individual properties that define this detail object's values.
	PropertyCount uint64 `json:"propertyCount"`

	// RegionalInternetRegistryHandle - The associated RWhois handle of this detail object. Used only when
	// detailed reassignments are necessary.
	RegionalInternetRegistryHandle *SoftLayer_Account_Rwhois_Handle `json:"regionalInternetRegistryHandle"`

	// RegionalInternetRegistryHandleId - The detail object's associated
	// [[SoftLayer_Account_Rwhois_Handle|RIR handle]] id
	RegionalInternetRegistryHandleId int `json:"regionalInternetRegistryHandleId"`
}

// CreateObject - This method will create a new SoftLayer_Account_Regional_Registry_Detail object.
// Input - [[SoftLayer_Account_Regional_Registry_Detail
// (type)|SoftLayer_Account_Regional_Registry_Detail]] The
// [[SoftLayer_Account_Regional_Registry_Detail_Type|type id]] of this detail object The id of the
// [[SoftLayer_Account_Rwhois_Handle|RWhois handle]] object. This is only to be used for detailed
// registrations, where a subnet is registered to an organization. The associated handle will be
// required to be a valid organization object id at the relevant registry. In this case, the detail
// object will only be valid for the registry the organization belongs to.
func (softlayer_account_regional_registry_detail *SoftLayer_Account_Regional_Registry_Detail) CreateObject(ctx *slapi.RequestContext, templateObject SoftLayer_Account_Regional_Registry_Detail) (*SoftLayer_Account_Regional_Registry_Detail, error) {
	var returnValue *SoftLayer_Account_Regional_Registry_Detail
	return returnValue, nil
}

// DeleteObject - This method will delete an existing SoftLayer_Account_Regional_Registry_Detail
// object.
func (softlayer_account_regional_registry_detail *SoftLayer_Account_Regional_Registry_Detail) DeleteObject(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObject - This method will edit an existing SoftLayer_Account_Regional_Registry_Detail object.
// For more detail, see [[SoftLayer_Account_Regional_Registry_Detail::createObject|createObject]].
func (softlayer_account_regional_registry_detail *SoftLayer_Account_Regional_Registry_Detail) EditObject(ctx *slapi.RequestContext, templateObject SoftLayer_Account_Regional_Registry_Detail) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_account_regional_registry_detail *SoftLayer_Account_Regional_Registry_Detail) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Account_Regional_Registry_Detail, error) {
	var returnValue *SoftLayer_Account_Regional_Registry_Detail
	return returnValue, nil
}

// UpdateReferencedRegistrations - This method will create a bulk transaction to update any
// registrations that reference this detail object. It should only be called from a child class such as
// [[SoftLayer_Account_Regional_Registry_Detail_Person]] or
// [[SoftLayer_Account_Regional_Registry_Detail_Network]]. The registrations should be in the Open or
// Registration_Complete status.
func (softlayer_account_regional_registry_detail *SoftLayer_Account_Regional_Registry_Detail) UpdateReferencedRegistrations(ctx *slapi.RequestContext) (*SoftLayer_Container_Network_Subnet_Registration_TransactionDetails, error) {
	var returnValue *SoftLayer_Container_Network_Subnet_Registration_TransactionDetails
	return returnValue, nil
}
