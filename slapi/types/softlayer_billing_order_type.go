package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Billing_Order_Type - The SoftLayer_Billing_Oder_Type data type contains general
// information relating to all the different types of orders that exist. This data pertains only to
// where an order was generated from, from any of the SoftLayer websites with ordering interfaces or
// directly through the SoftLayer
type SoftLayer_Billing_Order_Type struct {

	// Type - A simple keyname stating where a SoftLayer order originated from.
	Type string `json:"type,omitempty"`

	// Description - A brief description of where a SoftLayer order originated from.
	Description string `json:"description,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`
}

func (softlayer_billing_order_type *SoftLayer_Billing_Order_Type) String() string {
	return "SoftLayer_Billing_Order_Type"
}
