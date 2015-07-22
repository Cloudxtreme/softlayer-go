package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Virtual_Guest_Configuration_Option - An option found within a
// [[SoftLayer_Container_Virtual_Guest_Configuration (type)]] structure.
type SoftLayer_Container_Virtual_Guest_Configuration_Option struct {

	// ItemPrice - Provides hourly and monthly costs (if either are applicable), and a description of the
	// option.
	ItemPrice *SoftLayer_Product_Item_Price `json:"itemPrice"`

	// Template - Provides a fragment of the request with the properties and values that must be sent when
	// creating a computing instance with the option.
	Template *SoftLayer_Virtual_Guest `json:"template"`
}
