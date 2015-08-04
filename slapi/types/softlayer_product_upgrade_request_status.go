package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Upgrade_Request_Status - The SoftLayer_Product_Upgrade_Request_Status data type
// contains detailed information relating to an hardware or software upgrade request.
type SoftLayer_Product_Upgrade_Request_Status struct {

	// StatusCode - no documentation
	StatusCode string `json:"statusCode,omitempty"`

	// Description - The detailed description of an upgrade request status.
	Description string `json:"description,omitempty"`

	// Id - An internal identifier of an upgrade request status.
	Id int `json:"id,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`
}

func (softlayer_product_upgrade_request_status *SoftLayer_Product_Upgrade_Request_Status) String() string {
	return "SoftLayer_Product_Upgrade_Request_Status"
}
