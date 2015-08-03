package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Account_Media_Type - The SoftLayer_Account_Media_Type data type contains general
// information relating to the different types of media devices that SoftLayer currently supports, as
// part of the Data Transfer Request Service. Such devices as USB hard drives and flash drives, as well
// as optical media such as CD and DVD are currently supported.
type SoftLayer_Account_Media_Type struct {

	// Description - no documentation
	Description string `json:"description"`

	// Id - no documentation
	Id int `json:"id"`

	// KeyName - no documentation
	KeyName string `json:"keyName"`

	// Name - no documentation
	Name string `json:"name"`
}

func (softlayer_account_media_type *SoftLayer_Account_Media_Type) String() string {
	return "SoftLayer_Account_Media_Type"
}
