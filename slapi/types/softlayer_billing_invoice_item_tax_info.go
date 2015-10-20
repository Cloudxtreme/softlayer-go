package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"

	time "time"
)

// SoftLayer_Billing_Invoice_Item_Tax_Info - Information about the tax rates that apply to a particular
// invoice item.
type SoftLayer_Billing_Invoice_Item_Tax_Info struct {

	// ReportedFlag - A flag to indicate whether this is the official record for this invoice item.
	ReportedFlag bool `json:"reportedFlag,omitempty"`

	// SellerRegistration - The registration that the seller will use to report the invoice.
	SellerRegistration string `json:"sellerRegistration,omitempty"`

	// TaxRate - The tax rate used. Note that this might apply to only part of the
	TaxRate slapi.Float64 `json:"taxRate,omitempty"`

	// EffectiveTaxRate - The tax rate that can be multiplied by the subtotal to get the
	EffectiveTaxRate slapi.Float64 `json:"effectiveTaxRate,omitempty"`

	// FeeProperty - The type of fee being tracked for this particular set of tax information.
	FeeProperty string `json:"feeProperty,omitempty"`

	// Id - An invoice item's tax information internal identifier.
	Id int `json:"id,omitempty"`

	// InvoiceTaxInfoId - A reference to the tax information for the parent invoice.
	InvoiceTaxInfoId int `json:"invoiceTaxInfoId,omitempty"`

	// ModifyDate - The date and time the tax information was modified.
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// TaxAmountToCurrency - The tax amount (converted to the 'to' currency) associated with this line
	// item.
	TaxAmountToCurrency slapi.Float64 `json:"taxAmountToCurrency,omitempty"`

	// ToCurrencyId - The currency code that the invoice is being converted to.
	ToCurrencyId int `json:"toCurrencyId,omitempty"`

	// CreateDate - The date and time the tax information was recorded.
	CreateDate *time.Time `json:"createDate,omitempty"`

	// ExemptAmount - no documentation
	ExemptAmount slapi.Float64 `json:"exemptAmount,omitempty"`

	// NonTaxableBasis - no documentation
	NonTaxableBasis slapi.Float64 `json:"nonTaxableBasis,omitempty"`

	// TaxAmount - no documentation
	TaxAmount slapi.Float64 `json:"taxAmount,omitempty"`

	// Description - The invoice description with special information about the invoice.
	Description string `json:"description,omitempty"`

	// InvoiceItemId - no documentation
	InvoiceItemId int `json:"invoiceItemId,omitempty"`

	// TaxableBasis - no documentation
	TaxableBasis slapi.Float64 `json:"taxableBasis,omitempty"`

	// InvoiceTaxInfo - <nil>
	InvoiceTaxInfo *SoftLayer_Billing_Invoice_Tax_Info `json:"invoiceTaxInfo,omitempty"`

	// ToCurrency - This is the currency the invoice will be converted to.
	ToCurrency *SoftLayer_Billing_Currency `json:"toCurrency,omitempty"`

	// InvoiceItem - <nil>
	InvoiceItem *SoftLayer_Billing_Invoice_Item `json:"invoiceItem,omitempty"`
}

func (softlayer_billing_invoice_item_tax_info *SoftLayer_Billing_Invoice_Item_Tax_Info) String() string {
	return "SoftLayer_Billing_Invoice_Item_Tax_Info"
}
