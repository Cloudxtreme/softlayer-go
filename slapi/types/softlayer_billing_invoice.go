package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Billing_Invoice - The SoftLayer_Billing_Invoice data type contains general information
// relating to an individual invoice applied to a SoftLayer customer account. Personal information in
// this type such as names, addresses, and phone numbers are taken from the account's contact
// information at the time the invoice is generated.
type SoftLayer_Billing_Invoice struct {

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// AccountId - The SoftLayer customer account that an invoice belongs to.
	AccountId int `json:"accountId"`

	// Address1 - The first line of an address belonging to an account at the time an invoice is created.
	Address1 string `json:"address1"`

	// Address2 - The second line of an address belonging to an account at the time an invoice is created.
	Address2 string `json:"address2"`

	// Amount - no documentation
	Amount float64 `json:"amount"`

	// BrandAtInvoiceCreation - <nil>
	BrandAtInvoiceCreation *SoftLayer_Brand `json:"brandAtInvoiceCreation"`

	// City - The city portion of an address belonging to an account at the time an invoice is created.
	City string `json:"city"`

	// ClaimedTaxExemptTxFlag - Whether an account was exempt from taxes on their invoices at the time an
	// invoice is created.
	ClaimedTaxExemptTxFlag bool `json:"claimedTaxExemptTxFlag"`

	// ClosedDate - The date an invoice was closed. Open invoices have a null closed date.
	ClosedDate *time.Time `json:"closedDate"`

	// CompanyName - The company name belonging to an account at the time an invoice is created.
	CompanyName string `json:"companyName"`

	// Country - A two-letter abbreviation of the country portion of an address belonging to an account at
	// the time an invoice is created.
	Country string `json:"country"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// DetailedPdfGeneratedFlag - A flag that will reflect whether the detailed version of the pdf has been
	// generated.
	DetailedPdfGeneratedFlag bool `json:"detailedPdfGeneratedFlag"`

	// DocumentsGeneratedFlag - <nil>
	DocumentsGeneratedFlag bool `json:"documentsGeneratedFlag"`

	// Email - The email address belonging to an account at the time an invoice is created.
	Email string `json:"email"`

	// EndingBalance - An SoftLayer account's balance at the time an invoice is closed. This value is
	// measured in US Dollar currency.
	EndingBalance float64 `json:"endingBalance"`

	// FaxPhone - The fax telephone number belonging to an account at the time an invoice is created.
	FaxPhone string `json:"faxPhone"`

	// FirstName - The first name of the account holder at the time an invoice is created.
	FirstName string `json:"firstName"`

	// Id - no documentation
	Id int `json:"id"`

	// InvoiceTopLevelItemCount - A count of a list of top-level invoice items that are on the currently
	// pending invoice.
	InvoiceTopLevelItemCount uint64 `json:"invoiceTopLevelItemCount"`

	// InvoiceTopLevelItems - A list of top-level invoice items that are on the currently pending invoice.
	InvoiceTopLevelItems []*SoftLayer_Billing_Invoice_Item `json:"invoiceTopLevelItems"`

	// InvoiceTotalAmount - no documentation
	InvoiceTotalAmount float64 `json:"invoiceTotalAmount"`

	// InvoiceTotalOneTimeAmount - The total one-time charges for this invoice. This is the sum of one-time
	// charges + setup fees + labor fees. This does not include taxes.
	InvoiceTotalOneTimeAmount float64 `json:"invoiceTotalOneTimeAmount"`

	// InvoiceTotalOneTimeTaxAmount - A sum of all the taxes related to one time charges for this invoice.
	InvoiceTotalOneTimeTaxAmount float64 `json:"invoiceTotalOneTimeTaxAmount"`

	// InvoiceTotalPreTaxAmount - The total amount of this invoice. This does not include taxes.
	InvoiceTotalPreTaxAmount float64 `json:"invoiceTotalPreTaxAmount"`

	// InvoiceTotalRecurringAmount - The total Recurring amount of this invoice. This amount does not
	// include taxes or one time charges.
	InvoiceTotalRecurringAmount float64 `json:"invoiceTotalRecurringAmount"`

	// InvoiceTotalRecurringTaxAmount - The total amount of the recurring taxes on this invoice.
	InvoiceTotalRecurringTaxAmount float64 `json:"invoiceTotalRecurringTaxAmount"`

	// ItemCount - no documentation
	ItemCount uint64 `json:"itemCount"`

	// Items - no documentation
	Items []*SoftLayer_Billing_Invoice_Item `json:"items"`

	// LastName - The last name of the account holder at the time an invoice is created.
	LastName string `json:"lastName"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// OfficePhone - The telephone number belonging to an account at the time an invoice is created.
	OfficePhone string `json:"officePhone"`

	// Payment - no documentation
	Payment float64 `json:"payment"`

	// PaymentCount - no documentation
	PaymentCount uint64 `json:"paymentCount"`

	// Payments - no documentation
	Payments []*SoftLayer_Billing_Invoice_Receivable_Payment `json:"payments"`

	// PostalCode - The postal code portion of an address belonging to an account at the time an invoice is
	// created.
	PostalCode string `json:"postalCode"`

	// PurchaseOrderNumber - <nil>
	PurchaseOrderNumber string `json:"purchaseOrderNumber"`

	// SellerRegistration - no documentation
	SellerRegistration string `json:"sellerRegistration"`

	// StartingBalance - An SoftLayer account's balance at the time an invoice is created. This value is
	// measured in US Dollar currency.
	StartingBalance float64 `json:"startingBalance"`

	// State - A two-letter abbreviation of the state portion of an address belonging to an account at the
	// time an invoice is created. If the account that the invoice was generated for resides outside a
	// province then this is set to "other".
	State string `json:"state"`

	// StatusCode - An invoice's status. The status means SoftLayer has not yet received payment for this
	// invoice. status means that SoftLayer has received payment and closed the invoice. The status code
	// means SoftLayer closed the invoice without receiving a payment. Invoices are usually set to status
	// in cases where customer accounts are terminated for non-payment.
	StatusCode string `json:"statusCode"`

	// TaxInfo - This is the tax information that applies to tax auditing. This is the official tax record
	// for this invoice.
	TaxInfo *SoftLayer_Billing_Invoice_Tax_Info `json:"taxInfo"`

	// TaxInfoHistory - This is the set of tax information for any tax calculation for this invoice. Note
	// that not all of these are necessarily official, so use the taxInfo key to get the final information.
	TaxInfoHistory []*SoftLayer_Billing_Invoice_Tax_Info `json:"taxInfoHistory"`

	// TaxInfoHistoryCount - A count of this is the set of tax information for any tax calculation for this
	// invoice. Note that not all of these are necessarily official, so use the taxInfo key to get the
	// final information.
	TaxInfoHistoryCount uint64 `json:"taxInfoHistoryCount"`

	// TaxMessage - This is a message explaining the tax treatment for this invoice.
	TaxMessage string `json:"taxMessage"`

	// TaxStatusId - <nil>
	TaxStatusId int `json:"taxStatusId"`

	// TaxType - This is the strategy used to calculate tax on this invoice.
	TaxType *SoftLayer_Billing_Invoice_Tax_Type `json:"taxType"`

	// TaxTypeId - <nil>
	TaxTypeId int `json:"taxTypeId"`

	// TypeCode - An invoice's type. SoftLayer invoices and service credits are differentiated by their
	// type. The type code signifies an invoice for new service. A SoftLayer customer's first invoice has
	// the NEW type code. invoices are generated on a SoftLayer customer's anniversary billing date for
	// monthly services. invoices are generated when one-time charges are applied to an account. invoices
	// are generated whenever SoftLayer applies a credit against an account's balance. There are two
	// special types of service credits. type credits are applied against a customer's account balance
	// along with the receivables on their account. invoice credits are generated whenever a customer makes
	// an unscheduled payment.
	TypeCode string `json:"typeCode"`
}

func (softlayer_billing_invoice *SoftLayer_Billing_Invoice) String() string {
	return "SoftLayer_Billing_Invoice"
}
