package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
)

// SoftLayer_User_External_Binding - The SoftLayer_User_External_Binding data type contains general
// information for a single external binding. This includes the 3rd party vendor, type of binding, and
// a unique identifier and password that is used to authenticate against the 3rd party service.
type SoftLayer_User_External_Binding struct {

	// Active - The flag that determines whether the external binding is active will be used for
	// authentication or not.
	Active bool `json:"active"`

	// AttributeCount - A count of attributes of an external authentication binding.
	AttributeCount uint64 `json:"attributeCount"`

	// Attributes - no documentation
	Attributes []*SoftLayer_User_External_Binding_Attribute `json:"attributes"`

	// BillingItem - Information regarding the billing item for external authentication.
	BillingItem *SoftLayer_Billing_Item `json:"billingItem"`

	// CreateDate - The date that the external authentication binding was created.
	CreateDate *time.Time `json:"createDate"`

	// ExternalId - The identifier used to identify this binding to an external authentication source.
	ExternalId string `json:"externalId"`

	// Id - An external authentication binding's internal identifier.
	Id int `json:"id"`

	// Note - An optional note for identifying the external binding.
	Note string `json:"note"`

	// Password - The password used to authenticate the external id at an external authentication source.
	Password string `json:"password"`

	// Type - no documentation
	Type *SoftLayer_User_External_Binding_Type `json:"type"`

	// TypeId - The [[SoftLayer_User_External_Binding_Type|type]] identifier of an external authentication
	// binding.
	TypeId int `json:"typeId"`

	// UserId - An external authentication binding's associated [[SoftLayer_User_Customer|user account]]
	// id.
	UserId int `json:"userId"`

	// Vendor - no documentation
	Vendor *SoftLayer_User_External_Binding_Vendor `json:"vendor"`

	// VendorId - The [[SoftLayer_User_External_Binding_Vendor|vendor]] identifier of an external
	// authentication binding.
	VendorId int `json:"vendorId"`
}

// DeleteObject - Delete an external authentication binding. If the external binding currently has an
// active billing item associated you will be prevented from deleting the binding. The alternative
// method to remove an external authentication binding is to use the service cancellation form.
func (softlayer_user_external_binding *SoftLayer_User_External_Binding) DeleteObject(commonOptions *slapi.CommonOptions) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_user_external_binding *SoftLayer_User_External_Binding) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_User_External_Binding, error) {
	var returnValue *SoftLayer_User_External_Binding
	return returnValue, nil
}

// UpdateNote - Update the note of an external binding. The note is an optional property that is used
// to store information about a binding.
func (softlayer_user_external_binding *SoftLayer_User_External_Binding) UpdateNote(commonOptions *slapi.CommonOptions, text string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}
