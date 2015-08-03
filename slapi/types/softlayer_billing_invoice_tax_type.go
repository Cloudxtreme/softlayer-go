package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Billing_Invoice_Tax_Type - The invoice tax type data type models a single strategy for
// handling tax calculations.
type SoftLayer_Billing_Invoice_Tax_Type struct {

	// Id - A tax type's internal identifier. Each type of tax calculation strategy has a unique ID value.
	Id int `json:"id"`

	// KeyName - A unique string that identifies each strategy and is guaranteed to be stable over time.
	KeyName string `json:"keyName"`

	// Name - no documentation
	Name string `json:"name"`
}

func (softlayer_billing_invoice_tax_type *SoftLayer_Billing_Invoice_Tax_Type) String() string {
	return "SoftLayer_Billing_Invoice_Tax_Type"
}
