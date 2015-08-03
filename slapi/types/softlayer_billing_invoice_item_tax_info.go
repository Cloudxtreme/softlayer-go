package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Billing_Invoice_Item_Tax_Info - Information about the tax rates that apply to a particular
// invoice item.
type SoftLayer_Billing_Invoice_Item_Tax_Info struct {

	// CreateDate - The date and time the tax information was recorded.
	CreateDate *time.Time `json:"createDate"`

	// Description - The invoice description with special information about the invoice.
	Description string `json:"description"`

	// EffectiveTaxRate - The tax rate that can be multiplied by the subtotal to get the
	EffectiveTaxRate float32 `json:"effectiveTaxRate"`

	// ExemptAmount - no documentation
	ExemptAmount float32 `json:"exemptAmount"`

	// FeeProperty - The type of fee being tracked for this particular set of tax information.
	FeeProperty string `json:"feeProperty"`

	// Id - An invoice item's tax information internal identifier.
	Id int `json:"id"`

	// InvoiceItem - <nil>
	InvoiceItem *SoftLayer_Billing_Invoice_Item `json:"invoiceItem"`

	// InvoiceItemId - no documentation
	InvoiceItemId int `json:"invoiceItemId"`

	// InvoiceTaxInfo - <nil>
	InvoiceTaxInfo *SoftLayer_Billing_Invoice_Tax_Info `json:"invoiceTaxInfo"`

	// InvoiceTaxInfoId - A reference to the tax information for the parent invoice.
	InvoiceTaxInfoId int `json:"invoiceTaxInfoId"`

	// ModifyDate - The date and time the tax information was modified.
	ModifyDate *time.Time `json:"modifyDate"`

	// NonTaxableBasis - no documentation
	NonTaxableBasis float32 `json:"nonTaxableBasis"`

	// SellerRegistration - The registration that the seller will use to report the invoice.
	SellerRegistration string `json:"sellerRegistration"`

	// TaxAmount - no documentation
	TaxAmount float32 `json:"taxAmount"`

	// TaxAmountToCurrency - The tax amount (converted to the 'to' currency) associated with this line
	// item.
	TaxAmountToCurrency float32 `json:"taxAmountToCurrency"`

	// TaxRate - The tax rate used. Note that this might apply to only part of the
	TaxRate float32 `json:"taxRate"`

	// TaxableBasis - no documentation
	TaxableBasis float32 `json:"taxableBasis"`

	// ToCurrency - This is the currency the invoice will be converted to.
	ToCurrency *SoftLayer_Billing_Currency `json:"toCurrency"`

	// ToCurrencyId - The currency code that the invoice is being converted to.
	ToCurrencyId int `json:"toCurrencyId"`
}

func (softlayer_billing_invoice_item_tax_info *SoftLayer_Billing_Invoice_Item_Tax_Info) String() string {
	return "SoftLayer_Billing_Invoice_Item_Tax_Info"
}
