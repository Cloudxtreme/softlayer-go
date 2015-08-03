package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Regional_Registry_Detail_Property - Subnet registration properties are used to
// define various attributes of the [[SoftLayer_Account_Regional_Registry_Detail|detail objects]].
// These properties are defined by the [[SoftLayer_Account_Regional_Registry_Detail_Property_Type]]
// objects, which describe the available value formats.
type SoftLayer_Account_Regional_Registry_Detail_Property struct {

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// Id - no documentation
	Id int `json:"id"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate"`

	// PropertyTypeId - The numeric ID of the related
	// [[SoftLayer_Account_Regional_Registry_Detail_Property_Type|property type object]]
	PropertyTypeId int `json:"propertyTypeId"`

	// RegistrationDetailId - The numeric ID of the related
	// [[SoftLayer_Account_Regional_Registry_Detail|detail object]]
	RegistrationDetailId int `json:"registrationDetailId"`

	// SequencePosition - When multiple properties exist for a property type, defines the position in the
	// sequence of those properties
	SequencePosition int `json:"sequencePosition"`

	// Value - no documentation
	Value string `json:"value"`
}

// SoftLayer_Account_Regional_Registry_Detail_Property_Extended is SoftLayer_Account_Regional_Registry_Detail_Property with all maskable types.
type SoftLayer_Account_Regional_Registry_Detail_Property_Extended struct {
	SoftLayer_Account_Regional_Registry_Detail_Property

	// Detail - The [[SoftLayer_Account_Regional_Registry_Detail]] object this property belongs to
	Detail *SoftLayer_Account_Regional_Registry_Detail `json:"detail"`

	// PropertyType - The [[SoftLayer_Account_Regional_Registry_Detail_Property_Type]] object this property
	// belongs to
	PropertyType *SoftLayer_Account_Regional_Registry_Detail_Property_Type `json:"propertyType"`
}

func (softlayer_account_regional_registry_detail_property *SoftLayer_Account_Regional_Registry_Detail_Property) String() string {
	return "SoftLayer_Account_Regional_Registry_Detail_Property"
}
