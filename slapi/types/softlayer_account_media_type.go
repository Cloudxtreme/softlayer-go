package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Account_Media_Type - The SoftLayer_Account_Media_Type data type contains general
// information relating to the different types of media devices that SoftLayer currently supports, as
// part of the Data Transfer Request Service. Such devices as USB hard drives and flash drives, as well
// as optical media such as CD and DVD are currently supported.
type SoftLayer_Account_Media_Type struct {

	// KeyName - no documentation
	KeyName string `json:"keyName,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// Description - no documentation
	Description string `json:"description,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`
}

func (softlayer_account_media_type *SoftLayer_Account_Media_Type) String() string {
	return "SoftLayer_Account_Media_Type"
}
