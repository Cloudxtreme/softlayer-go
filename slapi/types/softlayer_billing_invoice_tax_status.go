package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Billing_Invoice_Tax_Status - The invoice tax status data type models a single status or
// state that an invoice can reflect in regard to an integration with a third-party tax calculation
// service.
type SoftLayer_Billing_Invoice_Tax_Status struct {

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// Id - <nil>
	Id int `json:"id"`

	// KeyName - <nil>
	KeyName string `json:"keyName"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate"`

	// Name - <nil>
	Name string `json:"name"`
}

func (softlayer_billing_invoice_tax_status *SoftLayer_Billing_Invoice_Tax_Status) String() string {
	return "SoftLayer_Billing_Invoice_Tax_Status"
}
