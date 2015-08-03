package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Billing_Invoice_Item_Hardware - The SoftLayer_Billing_Invoice_Item_Hardware data type
// contains a "resource". This resource is a link to the hardware tied to a SoftLayer_Billing_item
// whose category code is "server".
type SoftLayer_Billing_Invoice_Item_Hardware struct {
}

// SoftLayer_Billing_Invoice_Item_Hardware_Extended is SoftLayer_Billing_Invoice_Item_Hardware with all maskable types.
type SoftLayer_Billing_Invoice_Item_Hardware_Extended struct {
	SoftLayer_Billing_Invoice_Item_Hardware

	// Resource - no documentation
	Resource *SoftLayer_Hardware `json:"resource"`
}

func (softlayer_billing_invoice_item_hardware *SoftLayer_Billing_Invoice_Item_Hardware) String() string {
	return "SoftLayer_Billing_Invoice_Item_Hardware"
}
