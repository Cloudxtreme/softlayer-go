package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Billing_Invoice_Item_Tax_Info - Information about the tax rates that apply to a particular
// invoice item.
type SoftLayer_Billing_Invoice_Item_Tax_Info struct {

	// InvoiceTaxInfoId - A reference to the tax information for the parent invoice.
	InvoiceTaxInfoId int `json:"invoiceTaxInfoId,omitempty"`

	// ToCurrencyId - The currency code that the invoice is being converted to.
	ToCurrencyId int `json:"toCurrencyId,omitempty"`

	// EffectiveTaxRate - The tax rate that can be multiplied by the subtotal to get the
	EffectiveTaxRate float32 `json:"effectiveTaxRate,omitempty"`

	// ExemptAmount - no documentation
	ExemptAmount float32 `json:"exemptAmount,omitempty"`

	// TaxAmountToCurrency - The tax amount (converted to the 'to' currency) associated with this line
	// item.
	TaxAmountToCurrency float32 `json:"taxAmountToCurrency,omitempty"`

	// ModifyDate - The date and time the tax information was modified.
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// NonTaxableBasis - no documentation
	NonTaxableBasis float32 `json:"nonTaxableBasis,omitempty"`

	// TaxRate - The tax rate used. Note that this might apply to only part of the
	TaxRate float32 `json:"taxRate,omitempty"`

	// TaxableBasis - no documentation
	TaxableBasis float32 `json:"taxableBasis,omitempty"`

	// FeeProperty - The type of fee being tracked for this particular set of tax information.
	FeeProperty string `json:"feeProperty,omitempty"`

	// Id - An invoice item's tax information internal identifier.
	Id int `json:"id,omitempty"`

	// InvoiceItemId - no documentation
	InvoiceItemId int `json:"invoiceItemId,omitempty"`

	// SellerRegistration - The registration that the seller will use to report the invoice.
	SellerRegistration string `json:"sellerRegistration,omitempty"`

	// TaxAmount - no documentation
	TaxAmount float32 `json:"taxAmount,omitempty"`

	// CreateDate - The date and time the tax information was recorded.
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Description - The invoice description with special information about the invoice.
	Description string `json:"description,omitempty"`
}

func (softlayer_billing_invoice_item_tax_info *SoftLayer_Billing_Invoice_Item_Tax_Info) String() string {
	return "SoftLayer_Billing_Invoice_Item_Tax_Info"
}

// SoftLayer_Billing_Invoice_Item_Tax_Info_Extended is SoftLayer_Billing_Invoice_Item_Tax_Info with all maskable types.
type SoftLayer_Billing_Invoice_Item_Tax_Info_Extended struct {
	SoftLayer_Billing_Invoice_Item_Tax_Info

	// InvoiceTaxInfo - <nil>
	InvoiceTaxInfo *SoftLayer_Billing_Invoice_Tax_Info `json:"invoiceTaxInfo,omitempty"`

	// InvoiceItem - <nil>
	InvoiceItem *SoftLayer_Billing_Invoice_Item `json:"invoiceItem,omitempty"`

	// ToCurrency - This is the currency the invoice will be converted to.
	ToCurrency *SoftLayer_Billing_Currency `json:"toCurrency,omitempty"`
}

func (softlayer_billing_invoice_item_tax_info *SoftLayer_Billing_Invoice_Item_Tax_Info_Extended) String() string {
	return "SoftLayer_Billing_Invoice_Item_Tax_Info"
}
