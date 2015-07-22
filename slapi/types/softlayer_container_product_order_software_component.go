package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Order_Software_Component - This is the datatype that needs to be
// populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to
// place an Urchin order with SoftLayer.
type SoftLayer_Container_Product_Order_Software_Component struct {

	// VirtualGuestPremiumOs - A flag to determine if an attempt is being made to license a premium OS for
	// a virtual machine.
	VirtualGuestPremiumOs bool `json:"virtualGuestPremiumOs"`
}
