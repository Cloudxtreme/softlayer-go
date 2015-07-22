package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Hardware_Configuration_Option - An option found within a
// [[SoftLayer_Container_Hardware_Configuration (type)]] structure.
type SoftLayer_Container_Hardware_Configuration_Option struct {

	// ItemPrice - Provides hourly and monthly costs (if either are applicable), and a description of the
	// option.
	ItemPrice *SoftLayer_Product_Item_Price `json:"itemPrice"`

	// Preset - Provides a description of a fixed configuration preset with monthly and hourly costs.
	Preset *SoftLayer_Product_Package_Preset `json:"preset"`

	// Template - Provides a fragment of the request with the properties and values that must be sent when
	// creating a server with the option.
	Template *SoftLayer_Hardware `json:"template"`
}
