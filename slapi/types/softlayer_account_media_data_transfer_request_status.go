package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Account_Media_Data_Transfer_Request_Status - The
// SoftLayer_Account_Media_Data_Transfer_Request_Status data type contains general information relating
// to the statuses to which a Data Transfer Request may be set.
type SoftLayer_Account_Media_Data_Transfer_Request_Status struct {

	// KeyName - no documentation
	KeyName string `json:"keyName,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// Description - no documentation
	Description string `json:"description,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`
}

func (softlayer_account_media_data_transfer_request_status *SoftLayer_Account_Media_Data_Transfer_Request_Status) String() string {
	return "SoftLayer_Account_Media_Data_Transfer_Request_Status"
}
