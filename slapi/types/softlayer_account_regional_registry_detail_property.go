package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
)

// SoftLayer_Account_Regional_Registry_Detail_Property - Subnet registration properties are used to
// define various attributes of the [[SoftLayer_Account_Regional_Registry_Detail|detail objects]].
// These properties are defined by the [[SoftLayer_Account_Regional_Registry_Detail_Property_Type]]
// objects, which describe the available value formats.
type SoftLayer_Account_Regional_Registry_Detail_Property struct {

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// Detail - The [[SoftLayer_Account_Regional_Registry_Detail]] object this property belongs to
	Detail *SoftLayer_Account_Regional_Registry_Detail `json:"detail"`

	// Id - no documentation
	Id int `json:"id"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate"`

	// PropertyType - The [[SoftLayer_Account_Regional_Registry_Detail_Property_Type]] object this property
	// belongs to
	PropertyType *SoftLayer_Account_Regional_Registry_Detail_Property_Type `json:"propertyType"`

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

// CreateObject - This method will create a new SoftLayer_Account_Regional_Registry_Detail_Property
// object. Input - [[SoftLayer_Account_Regional_Registry_Detail_Property
// (type)|SoftLayer_Account_Regional_Registry_Detail_Property]] The numeric ID of the
// [[SoftLayer_Account_Regional_Registry_Detail|detail object]] this property belongs to The numeric ID
// of the associated [[SoftLayer_Account_Regional_Registry_Detail_Property_Type]] object When more than
// one property of the same type exists on a detail object, this value determines the position in that
// collection. This can be thought of more as a sort order.
func (softlayer_account_regional_registry_detail_property *SoftLayer_Account_Regional_Registry_Detail_Property) CreateObject(ctx *slapi.RequestContext, templateObject SoftLayer_Account_Regional_Registry_Detail_Property) (*SoftLayer_Account_Regional_Registry_Detail_Property, error) {
	var returnValue *SoftLayer_Account_Regional_Registry_Detail_Property
	return returnValue, nil
}

// CreateObjects - Edit multiple [[SoftLayer_Account_Regional_Registry_Detail_Property]] objects.
func (softlayer_account_regional_registry_detail_property *SoftLayer_Account_Regional_Registry_Detail_Property) CreateObjects(ctx *slapi.RequestContext, templateObjects []SoftLayer_Account_Regional_Registry_Detail_Property) ([]*SoftLayer_Account_Regional_Registry_Detail_Property, error) {
	var returnValue []*SoftLayer_Account_Regional_Registry_Detail_Property
	return returnValue, nil
}

// DeleteObject - This method will delete an existing
// SoftLayer_Account_Regional_Registry_Detail_Property object.
func (softlayer_account_regional_registry_detail_property *SoftLayer_Account_Regional_Registry_Detail_Property) DeleteObject(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObject - This method will edit an existing SoftLayer_Account_Regional_Registry_Detail_Property
// object. For more detail, see
// [[SoftLayer_Account_Regional_Registry_Detail_Property::createObject|createObject]].
func (softlayer_account_regional_registry_detail_property *SoftLayer_Account_Regional_Registry_Detail_Property) EditObject(ctx *slapi.RequestContext, templateObject SoftLayer_Account_Regional_Registry_Detail_Property) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObjects - Edit multiple [[SoftLayer_Account_Regional_Registry_Detail_Property]] objects.
func (softlayer_account_regional_registry_detail_property *SoftLayer_Account_Regional_Registry_Detail_Property) EditObjects(ctx *slapi.RequestContext, templateObjects []SoftLayer_Account_Regional_Registry_Detail_Property) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_account_regional_registry_detail_property *SoftLayer_Account_Regional_Registry_Detail_Property) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Account_Regional_Registry_Detail_Property, error) {
	var returnValue *SoftLayer_Account_Regional_Registry_Detail_Property
	return returnValue, nil
}
