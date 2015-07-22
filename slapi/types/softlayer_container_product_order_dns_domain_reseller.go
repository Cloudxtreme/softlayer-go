package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Order_Dns_Domain_Reseller - This is the datatype that needs to be
// populated and sent to SoftLayer_Product_Order::placeOrder. The
// SoftLayer_Container_Product_Order_Dns_Domain_Reseller datatype contains everything required to place
// a domain reseller credit order with SoftLayer.
type SoftLayer_Container_Product_Order_Dns_Domain_Reseller struct {

	// CreditAmount - Amount to be credited to the domain reseller account.
	CreditAmount float64 `json:"creditAmount"`
}
