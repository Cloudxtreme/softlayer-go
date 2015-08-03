package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Account_Media - The SoftLayer_Account_Media data type contains information on a single
// piece of media associated with a Data Transfer Service request.
type SoftLayer_Account_Media struct {

	// Id - no documentation
	Id int `json:"id"`

	// RequestId - no documentation
	RequestId int `json:"requestId"`

	// SerialNumber - no documentation
	SerialNumber string `json:"serialNumber"`

	// TypeId - no documentation
	TypeId int `json:"typeId"`

	// Description - no documentation
	Description string `json:"description"`
}

// SoftLayer_Account_Media_Extended is SoftLayer_Account_Media with all maskable types.
type SoftLayer_Account_Media_Extended struct {
	SoftLayer_Account_Media

	// CreateUser - no documentation
	CreateUser *SoftLayer_User_Customer `json:"createUser"`

	// ModifyUser - no documentation
	ModifyUser *SoftLayer_User_Customer `json:"modifyUser"`

	// Type - no documentation
	Type *SoftLayer_Account_Media_Type `json:"type"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// ModifyEmployee - no documentation
	ModifyEmployee *SoftLayer_User_Employee `json:"modifyEmployee"`

	// Request - no documentation
	Request *SoftLayer_Account_Media_Data_Transfer_Request `json:"request"`

	// Volume - A guest's associated EVault network storage service account.
	Volume *SoftLayer_Network_Storage `json:"volume"`

	// Datacenter - no documentation
	Datacenter *SoftLayer_Location `json:"datacenter"`
}

func (softlayer_account_media *SoftLayer_Account_Media) String() string {
	return "SoftLayer_Account_Media"
}
