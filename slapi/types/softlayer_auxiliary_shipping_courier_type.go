package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Auxiliary_Shipping_Courier_Type - <nil>
type SoftLayer_Auxiliary_Shipping_Courier_Type struct {

	// Name - <nil>
	Name string `json:"name,omitempty"`

	// Description - <nil>
	Description string `json:"description,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// KeyName - <nil>
	KeyName string `json:"keyName,omitempty"`
}

func (softlayer_auxiliary_shipping_courier_type *SoftLayer_Auxiliary_Shipping_Courier_Type) String() string {
	return "SoftLayer_Auxiliary_Shipping_Courier_Type"
}

// SoftLayer_Auxiliary_Shipping_Courier_Type_Extended is SoftLayer_Auxiliary_Shipping_Courier_Type with all maskable types.
type SoftLayer_Auxiliary_Shipping_Courier_Type_Extended struct {
	SoftLayer_Auxiliary_Shipping_Courier_Type

	// Courier - <nil>
	Courier []*SoftLayer_Auxiliary_Shipping_Courier `json:"courier,omitempty"`

	// CourierCount - no documentation
	CourierCount uint64 `json:"courierCount,omitempty"`
}

func (softlayer_auxiliary_shipping_courier_type *SoftLayer_Auxiliary_Shipping_Courier_Type_Extended) String() string {
	return "SoftLayer_Auxiliary_Shipping_Courier_Type"
}
