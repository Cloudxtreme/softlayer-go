package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Billing_Invoice_Tax_Info - Invoice tax information contains top-level information about
// the taxes recorded for a particular invoice.
type SoftLayer_Billing_Invoice_Tax_Info struct {

	// Id - The internal identifier for this invoice tax information.
	Id int `json:"id"`

	// ReportedFlag - A flag to indicate whether the invoice will be auditable.
	ReportedFlag int `json:"reportedFlag"`

	// CreateDate - The date and time this tax information was recorded.
	CreateDate *time.Time `json:"createDate"`

	// CurrencyId - The currency code that the invoice should be recorded in.
	CurrencyId int `json:"currencyId"`

	// InvoiceId - no documentation
	InvoiceId int `json:"invoiceId"`

	// ModifyDate - The date and time this tax information was updated.
	ModifyDate *time.Time `json:"modifyDate"`
}

// SoftLayer_Billing_Invoice_Tax_Info_Extended is SoftLayer_Billing_Invoice_Tax_Info with all maskable types.
type SoftLayer_Billing_Invoice_Tax_Info_Extended struct {
	SoftLayer_Billing_Invoice_Tax_Info

	// FunctionalCurrency - This is the functional currency used for the invoice.
	FunctionalCurrency *SoftLayer_Billing_Currency `json:"functionalCurrency"`

	// Items - This is the collection of tax information for each of the related invoice items.
	Items []*SoftLayer_Billing_Invoice_Item_Tax_Info `json:"items"`

	// TotalTaxAmountToCurrency - This the total tax amount (converted to the 'to' currency) for the
	// invoice.
	TotalTaxAmountToCurrency float32 `json:"totalTaxAmountToCurrency"`

	// ItemCount - A count of this is the collection of tax information for each of the related invoice
	// items.
	ItemCount uint64 `json:"itemCount"`

	// Currency - no documentation
	Currency *SoftLayer_Billing_Currency `json:"currency"`

	// Invoice - This is the related invoice for this tax-related information.
	Invoice *SoftLayer_Billing_Invoice `json:"invoice"`

	// ItemWithCurrencyInfo - This tax information on the invoice item that includes currency details.
	ItemWithCurrencyInfo *SoftLayer_Billing_Invoice_Item_Tax_Info `json:"itemWithCurrencyInfo"`
}

func (softlayer_billing_invoice_tax_info *SoftLayer_Billing_Invoice_Tax_Info) String() string {
	return "SoftLayer_Billing_Invoice_Tax_Info"
}
