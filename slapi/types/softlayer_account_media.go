package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Account_Media - The SoftLayer_Account_Media data type contains information on a single
// piece of media associated with a Data Transfer Service request.
type SoftLayer_Account_Media struct {

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// CreateUser - no documentation
	CreateUser *SoftLayer_User_Customer `json:"createUser"`

	// Datacenter - no documentation
	Datacenter *SoftLayer_Location `json:"datacenter"`

	// Description - no documentation
	Description string `json:"description"`

	// Id - no documentation
	Id int `json:"id"`

	// ModifyEmployee - no documentation
	ModifyEmployee *SoftLayer_User_Employee `json:"modifyEmployee"`

	// ModifyUser - no documentation
	ModifyUser *SoftLayer_User_Customer `json:"modifyUser"`

	// Request - no documentation
	Request *SoftLayer_Account_Media_Data_Transfer_Request `json:"request"`

	// RequestId - no documentation
	RequestId int `json:"requestId"`

	// SerialNumber - no documentation
	SerialNumber string `json:"serialNumber"`

	// Type - no documentation
	Type *SoftLayer_Account_Media_Type `json:"type"`

	// TypeId - no documentation
	TypeId int `json:"typeId"`

	// Volume - A guest's associated EVault network storage service account.
	Volume *SoftLayer_Network_Storage `json:"volume"`
}

// EditObject - Edit the properties of a media record by passing in a modified instance of a
// SoftLayer_Account_Media object.
func (softlayer_account_media *SoftLayer_Account_Media) EditObject(commonOptions *slapi.CommonOptions, templateObject SoftLayer_Account_Media) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetAllMediaTypes - Retrieve a list supported media types for SoftLayer's Data Transfer Service.
func (softlayer_account_media *SoftLayer_Account_Media) GetAllMediaTypes(commonOptions *slapi.CommonOptions) ([]*SoftLayer_Account_Media_Type, error) {
	var returnValue []*SoftLayer_Account_Media_Type
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_account_media *SoftLayer_Account_Media) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Account_Media, error) {
	var returnValue *SoftLayer_Account_Media
	return returnValue, nil
}

// RemoveMediaFromList - Remove a media from a SoftLayer account's list of media. The media record is
// not deleted.
func (softlayer_account_media *SoftLayer_Account_Media) RemoveMediaFromList(commonOptions *slapi.CommonOptions, mediaTemplate SoftLayer_Account_Media) (int, error) {
	var returnValue int
	return returnValue, nil
}
