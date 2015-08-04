package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Account_Media - The SoftLayer_Account_Media data type contains information on a single
// piece of media associated with a Data Transfer Service request.
type SoftLayer_Account_Media struct {

	// SerialNumber - no documentation
	SerialNumber string `json:"serialNumber,omitempty"`

	// RequestId - no documentation
	RequestId int `json:"requestId,omitempty"`

	// TypeId - no documentation
	TypeId int `json:"typeId,omitempty"`

	// Description - no documentation
	Description string `json:"description,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`
}

func (softlayer_account_media *SoftLayer_Account_Media) String() string {
	return "SoftLayer_Account_Media"
}

// SoftLayer_Account_Media_Extended is SoftLayer_Account_Media with all maskable types.
type SoftLayer_Account_Media_Extended struct {
	SoftLayer_Account_Media

	// Volume - A guest's associated EVault network storage service account.
	Volume *SoftLayer_Network_Storage `json:"volume,omitempty"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account,omitempty"`

	// CreateUser - no documentation
	CreateUser *SoftLayer_User_Customer `json:"createUser,omitempty"`

	// ModifyUser - no documentation
	ModifyUser *SoftLayer_User_Customer `json:"modifyUser,omitempty"`

	// Type - no documentation
	Type *SoftLayer_Account_Media_Type `json:"type,omitempty"`

	// Datacenter - no documentation
	Datacenter *SoftLayer_Location `json:"datacenter,omitempty"`

	// ModifyEmployee - no documentation
	ModifyEmployee *SoftLayer_User_Employee `json:"modifyEmployee,omitempty"`

	// Request - no documentation
	Request *SoftLayer_Account_Media_Data_Transfer_Request `json:"request,omitempty"`
}

func (softlayer_account_media *SoftLayer_Account_Media_Extended) String() string {
	return "SoftLayer_Account_Media"
}
