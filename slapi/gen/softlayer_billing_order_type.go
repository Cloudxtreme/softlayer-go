package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Billing_Order_Type - The SoftLayer_Billing_Oder_Type data type contains general
// information relating to all the different types of orders that exist. This data pertains only to
// where an order was generated from, from any of the SoftLayer websites with ordering interfaces or
// directly through the SoftLayer
type SoftLayer_Billing_Order_Type struct {

	// Description - A brief description of where a SoftLayer order originated from.
	Description string `json:"description"`

	// Id - no documentation
	Id int `json:"id"`

	// Type - A simple keyname stating where a SoftLayer order originated from.
	Type string `json:"type"`
}