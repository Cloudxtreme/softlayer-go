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

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// SequencePosition - When multiple properties exist for a property type, defines the position in the
	// sequence of those properties
	SequencePosition int `json:"sequencePosition,omitempty"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate,omitempty"`

	// PropertyTypeId - The numeric ID of the related
	// [[SoftLayer_Account_Regional_Registry_Detail_Property_Type|property type object]]
	PropertyTypeId int `json:"propertyTypeId,omitempty"`

	// RegistrationDetailId - The numeric ID of the related
	// [[SoftLayer_Account_Regional_Registry_Detail|detail object]]
	RegistrationDetailId int `json:"registrationDetailId,omitempty"`

	// Value - no documentation
	Value string `json:"value,omitempty"`

	// Detail - The [[SoftLayer_Account_Regional_Registry_Detail]] object this property belongs to
	Detail *SoftLayer_Account_Regional_Registry_Detail `json:"detail,omitempty"`

	// PropertyType - The [[SoftLayer_Account_Regional_Registry_Detail_Property_Type]] object this property
	// belongs to
	PropertyType *SoftLayer_Account_Regional_Registry_Detail_Property_Type `json:"propertyType,omitempty"`
}

func (softlayer_account_regional_registry_detail_property *SoftLayer_Account_Regional_Registry_Detail_Property) String() string {
	return "SoftLayer_Account_Regional_Registry_Detail_Property"
}
