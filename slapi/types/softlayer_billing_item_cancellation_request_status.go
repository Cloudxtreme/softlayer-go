package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Billing_Item_Cancellation_Request_Status -
// SoftLayer_Billing_Item_Cancellation_Request_Status data type represents the status of a service
// cancellation request.
type SoftLayer_Billing_Item_Cancellation_Request_Status struct {

	// Name - no documentation
	Name string `json:"name"`

	// Description - The short description of a cancellation request status
	Description string `json:"description"`

	// Id - The internal identifier of a cancellation request status.
	Id int `json:"id"`

	// KeyName - no documentation
	KeyName string `json:"keyName"`
}

func (softlayer_billing_item_cancellation_request_status *SoftLayer_Billing_Item_Cancellation_Request_Status) String() string {
	return "SoftLayer_Billing_Item_Cancellation_Request_Status"
}
