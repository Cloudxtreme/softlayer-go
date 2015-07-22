package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
)

// SoftLayer_Account_Note_Type - <nil>
type SoftLayer_Account_Note_Type struct {

	// BrandId - <nil>
	BrandId int `json:"brandId"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// Description - <nil>
	Description string `json:"description"`

	// Id - <nil>
	Id int `json:"id"`

	// KeyName - <nil>
	KeyName string `json:"keyName"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate"`

	// Name - <nil>
	Name string `json:"name"`

	// ValueExpression - <nil>
	ValueExpression string `json:"valueExpression"`
}

// CreateObject - <nil>
func (softlayer_account_note_type *SoftLayer_Account_Note_Type) CreateObject(ctx *slapi.RequestContext, templateObject SoftLayer_Account_Note_Type) (*SoftLayer_Account_Note_Type, error) {
	var returnValue *SoftLayer_Account_Note_Type
	return returnValue, nil
}

// DeleteObject - <nil>
func (softlayer_account_note_type *SoftLayer_Account_Note_Type) DeleteObject(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObject - <nil>
func (softlayer_account_note_type *SoftLayer_Account_Note_Type) EditObject(ctx *slapi.RequestContext, templateObject SoftLayer_Account_Note_Type) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetAllObjects - <nil>
func (softlayer_account_note_type *SoftLayer_Account_Note_Type) GetAllObjects(ctx *slapi.RequestContext) ([]*SoftLayer_Account_Note_Type, error) {
	var returnValue []*SoftLayer_Account_Note_Type
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_account_note_type *SoftLayer_Account_Note_Type) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Account_Note_Type, error) {
	var returnValue *SoftLayer_Account_Note_Type
	return returnValue, nil
}
