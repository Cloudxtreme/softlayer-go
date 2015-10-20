package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"

	time "time"
)

// SoftLayer_Billing_Invoice - The SoftLayer_Billing_Invoice data type contains general information
// relating to an individual invoice applied to a SoftLayer customer account. Personal information in
// this type such as names, addresses, and phone numbers are taken from the account's contact
// information at the time the invoice is generated.
type SoftLayer_Billing_Invoice struct {

	// Email - The email address belonging to an account at the time an invoice is created.
	Email string `json:"email,omitempty"`

	// AccountId - The SoftLayer customer account that an invoice belongs to.
	AccountId int `json:"accountId,omitempty"`

	// Address1 - The first line of an address belonging to an account at the time an invoice is created.
	Address1 string `json:"address1,omitempty"`

	// State - A two-letter abbreviation of the state portion of an address belonging to an account at the
	// time an invoice is created. If the account that the invoice was generated for resides outside a
	// province then this is set to "other".
	State string `json:"state,omitempty"`

	// TaxStatusId - <nil>
	TaxStatusId int `json:"taxStatusId,omitempty"`

	// Address2 - The second line of an address belonging to an account at the time an invoice is created.
	Address2 string `json:"address2,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// StartingBalance - An SoftLayer account's balance at the time an invoice is created. This value is
	// measured in US Dollar currency.
	StartingBalance slapi.Float64 `json:"startingBalance,omitempty"`

	// PostalCode - The postal code portion of an address belonging to an account at the time an invoice is
	// created.
	PostalCode string `json:"postalCode,omitempty"`

	// City - The city portion of an address belonging to an account at the time an invoice is created.
	City string `json:"city,omitempty"`

	// ClaimedTaxExemptTxFlag - Whether an account was exempt from taxes on their invoices at the time an
	// invoice is created.
	ClaimedTaxExemptTxFlag bool `json:"claimedTaxExemptTxFlag,omitempty"`

	// EndingBalance - An SoftLayer account's balance at the time an invoice is closed. This value is
	// measured in US Dollar currency.
	EndingBalance slapi.Float64 `json:"endingBalance,omitempty"`

	// StatusCode - An invoice's status. The status means SoftLayer has not yet received payment for this
	// invoice. status means that SoftLayer has received payment and closed the invoice. The status code
	// means SoftLayer closed the invoice without receiving a payment. Invoices are usually set to status
	// in cases where customer accounts are terminated for non-payment.
	StatusCode string `json:"statusCode,omitempty"`

	// ClosedDate - The date an invoice was closed. Open invoices have a null closed date.
	ClosedDate *time.Time `json:"closedDate,omitempty"`

	// Country - A two-letter abbreviation of the country portion of an address belonging to an account at
	// the time an invoice is created.
	Country string `json:"country,omitempty"`

	// DocumentsGeneratedFlag - <nil>
	DocumentsGeneratedFlag bool `json:"documentsGeneratedFlag,omitempty"`

	// LastName - The last name of the account holder at the time an invoice is created.
	LastName string `json:"lastName,omitempty"`

	// TypeCode - An invoice's type. SoftLayer invoices and service credits are differentiated by their
	// type. The type code signifies an invoice for new service. A SoftLayer customer's first invoice has
	// the NEW type code. invoices are generated on a SoftLayer customer's anniversary billing date for
	// monthly services. invoices are generated when one-time charges are applied to an account. invoices
	// are generated whenever SoftLayer applies a credit against an account's balance. There are two
	// special types of service credits. type credits are applied against a customer's account balance
	// along with the receivables on their account. invoice credits are generated whenever a customer makes
	// an unscheduled payment.
	TypeCode string `json:"typeCode,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// FaxPhone - The fax telephone number belonging to an account at the time an invoice is created.
	FaxPhone string `json:"faxPhone,omitempty"`

	// OfficePhone - The telephone number belonging to an account at the time an invoice is created.
	OfficePhone string `json:"officePhone,omitempty"`

	// TaxTypeId - <nil>
	TaxTypeId int `json:"taxTypeId,omitempty"`

	// FirstName - The first name of the account holder at the time an invoice is created.
	FirstName string `json:"firstName,omitempty"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// CompanyName - The company name belonging to an account at the time an invoice is created.
	CompanyName string `json:"companyName,omitempty"`

	// PurchaseOrderNumber - <nil>
	PurchaseOrderNumber string `json:"purchaseOrderNumber,omitempty"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account,omitempty"`

	// DetailedPdfGeneratedFlag - A flag that will reflect whether the detailed version of the pdf has been
	// generated.
	DetailedPdfGeneratedFlag bool `json:"detailedPdfGeneratedFlag,omitempty"`

	// Payment - no documentation
	Payment slapi.Float64 `json:"payment,omitempty"`

	// TaxInfo - This is the tax information that applies to tax auditing. This is the official tax record
	// for this invoice.
	TaxInfo *SoftLayer_Billing_Invoice_Tax_Info `json:"taxInfo,omitempty"`

	// TaxMessage - This is a message explaining the tax treatment for this invoice.
	TaxMessage string `json:"taxMessage,omitempty"`

	// TaxInfoHistoryCount - A count of this is the set of tax information for any tax calculation for this
	// invoice. Note that not all of these are necessarily official, so use the taxInfo key to get the
	// final information.
	TaxInfoHistoryCount uint64 `json:"taxInfoHistoryCount,omitempty"`

	// InvoiceTopLevelItemCount - A count of a list of top-level invoice items that are on the currently
	// pending invoice.
	InvoiceTopLevelItemCount uint64 `json:"invoiceTopLevelItemCount,omitempty"`

	// InvoiceTotalOneTimeTaxAmount - A sum of all the taxes related to one time charges for this invoice.
	InvoiceTotalOneTimeTaxAmount slapi.Float64 `json:"invoiceTotalOneTimeTaxAmount,omitempty"`

	// Payments - no documentation
	Payments []*SoftLayer_Billing_Invoice_Receivable_Payment `json:"payments,omitempty"`

	// InvoiceTotalPreTaxAmount - The total amount of this invoice. This does not include taxes.
	InvoiceTotalPreTaxAmount slapi.Float64 `json:"invoiceTotalPreTaxAmount,omitempty"`

	// SellerRegistration - no documentation
	SellerRegistration string `json:"sellerRegistration,omitempty"`

	// PaymentCount - no documentation
	PaymentCount uint64 `json:"paymentCount,omitempty"`

	// Items - no documentation
	Items []*SoftLayer_Billing_Invoice_Item `json:"items,omitempty"`

	// ItemCount - no documentation
	ItemCount uint64 `json:"itemCount,omitempty"`

	// InvoiceTopLevelItems - A list of top-level invoice items that are on the currently pending invoice.
	InvoiceTopLevelItems []*SoftLayer_Billing_Invoice_Item `json:"invoiceTopLevelItems,omitempty"`

	// InvoiceTotalRecurringAmount - The total Recurring amount of this invoice. This amount does not
	// include taxes or one time charges.
	InvoiceTotalRecurringAmount slapi.Float64 `json:"invoiceTotalRecurringAmount,omitempty"`

	// TaxType - This is the strategy used to calculate tax on this invoice.
	TaxType *SoftLayer_Billing_Invoice_Tax_Type `json:"taxType,omitempty"`

	// BrandAtInvoiceCreation - <nil>
	BrandAtInvoiceCreation *SoftLayer_Brand `json:"brandAtInvoiceCreation,omitempty"`

	// InvoiceTotalOneTimeAmount - The total one-time charges for this invoice. This is the sum of one-time
	// charges + setup fees + labor fees. This does not include taxes.
	InvoiceTotalOneTimeAmount slapi.Float64 `json:"invoiceTotalOneTimeAmount,omitempty"`

	// TaxInfoHistory - This is the set of tax information for any tax calculation for this invoice. Note
	// that not all of these are necessarily official, so use the taxInfo key to get the final information.
	TaxInfoHistory []*SoftLayer_Billing_Invoice_Tax_Info `json:"taxInfoHistory,omitempty"`

	// InvoiceTotalRecurringTaxAmount - The total amount of the recurring taxes on this invoice.
	InvoiceTotalRecurringTaxAmount slapi.Float64 `json:"invoiceTotalRecurringTaxAmount,omitempty"`

	// Amount - no documentation
	Amount slapi.Float64 `json:"amount,omitempty"`

	// InvoiceTotalAmount - no documentation
	InvoiceTotalAmount slapi.Float64 `json:"invoiceTotalAmount,omitempty"`
}

func (softlayer_billing_invoice *SoftLayer_Billing_Invoice) String() string {
	return "SoftLayer_Billing_Invoice"
}
