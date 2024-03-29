package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Auxiliary_Shipping_Courier - The SoftLayer_Auxiliary_Shipping_Courier data type contains
// general information relating the different (major) couriers that SoftLayer may use for shipping.
type SoftLayer_Auxiliary_Shipping_Courier struct {

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// KeyName - no documentation
	KeyName string `json:"keyName,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// Url - no documentation
	Url string `json:"url,omitempty"`
}

func (softlayer_auxiliary_shipping_courier *SoftLayer_Auxiliary_Shipping_Courier) String() string {
	return "SoftLayer_Auxiliary_Shipping_Courier"
}
