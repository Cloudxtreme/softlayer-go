package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Billing_Item_Account_Media_Data_Transfer_Request - The
// SoftLayer_Billing_Item_Account_Media_Data_Transfer_Request data type contains general information
// relating to a single SoftLayer billing item for a data transfer request.
type SoftLayer_Billing_Item_Account_Media_Data_Transfer_Request struct {

	// Resource - The data transfer request to which the billing item points.
	Resource *SoftLayer_Account_Media_Data_Transfer_Request `json:"resource"`
}

func (softlayer_billing_item_account_media_data_transfer_request *SoftLayer_Billing_Item_Account_Media_Data_Transfer_Request) String() string {
	return "SoftLayer_Billing_Item_Account_Media_Data_Transfer_Request"
}
