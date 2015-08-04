package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Billing_Invoice_Tax_Info - Invoice tax information contains top-level information about
// the taxes recorded for a particular invoice.
type SoftLayer_Billing_Invoice_Tax_Info struct {

	// CreateDate - The date and time this tax information was recorded.
	CreateDate *time.Time `json:"createDate,omitempty"`

	// ModifyDate - The date and time this tax information was updated.
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// CurrencyId - The currency code that the invoice should be recorded in.
	CurrencyId int `json:"currencyId,omitempty"`

	// Id - The internal identifier for this invoice tax information.
	Id int `json:"id,omitempty"`

	// InvoiceId - no documentation
	InvoiceId int `json:"invoiceId,omitempty"`

	// ReportedFlag - A flag to indicate whether the invoice will be auditable.
	ReportedFlag int `json:"reportedFlag,omitempty"`
}

func (softlayer_billing_invoice_tax_info *SoftLayer_Billing_Invoice_Tax_Info) String() string {
	return "SoftLayer_Billing_Invoice_Tax_Info"
}

// SoftLayer_Billing_Invoice_Tax_Info_Extended is SoftLayer_Billing_Invoice_Tax_Info with all maskable types.
type SoftLayer_Billing_Invoice_Tax_Info_Extended struct {
	SoftLayer_Billing_Invoice_Tax_Info

	// Currency - no documentation
	Currency *SoftLayer_Billing_Currency `json:"currency,omitempty"`

	// FunctionalCurrency - This is the functional currency used for the invoice.
	FunctionalCurrency *SoftLayer_Billing_Currency `json:"functionalCurrency,omitempty"`

	// ItemWithCurrencyInfo - This tax information on the invoice item that includes currency details.
	ItemWithCurrencyInfo *SoftLayer_Billing_Invoice_Item_Tax_Info `json:"itemWithCurrencyInfo,omitempty"`

	// Items - This is the collection of tax information for each of the related invoice items.
	Items []*SoftLayer_Billing_Invoice_Item_Tax_Info `json:"items,omitempty"`

	// TotalTaxAmountToCurrency - This the total tax amount (converted to the 'to' currency) for the
	// invoice.
	TotalTaxAmountToCurrency float32 `json:"totalTaxAmountToCurrency,omitempty"`

	// ItemCount - A count of this is the collection of tax information for each of the related invoice
	// items.
	ItemCount uint64 `json:"itemCount,omitempty"`

	// Invoice - This is the related invoice for this tax-related information.
	Invoice *SoftLayer_Billing_Invoice `json:"invoice,omitempty"`
}

func (softlayer_billing_invoice_tax_info *SoftLayer_Billing_Invoice_Tax_Info_Extended) String() string {
	return "SoftLayer_Billing_Invoice_Tax_Info"
}
