package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Billing_Payment_Card_Transaction - The SoftLayer_Billing_Payment_Card_Transaction data
// type contains general information relating to attempted credit card transactions.
type SoftLayer_Billing_Payment_Card_Transaction struct {

	// BillingNameFirst - no documentation
	BillingNameFirst string `json:"billingNameFirst"`

	// ReturnStatus - The status code returned from the financial institution.
	ReturnStatus int `json:"returnStatus"`

	// BillingAddressLine1 - The physical street address. Reserve information such as "apartment #123" or
	// "Suite 2" for line 1.
	BillingAddressLine1 string `json:"billingAddressLine1"`

	// BillingPhoneVoice - The phone number associated with a customer account.
	BillingPhoneVoice string `json:"billingPhoneVoice"`

	// CardType - The type of payment issued (i.e. Visa, MasterCard, American Express).
	CardType string `json:"cardType"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// RequestId - The unique identifier of the request submitted to the financial institution.
	RequestId string `json:"requestId"`

	// BillingCity - no documentation
	BillingCity string `json:"billingCity"`

	// BillingPostalCode - The Zip or Postal Code for the billing address on an account.
	BillingPostalCode string `json:"billingPostalCode"`

	// CardAccountLast4 - no documentation
	CardAccountLast4 int `json:"cardAccountLast4"`

	// CardExpirationYear - The year in which a customer's payment card will expire.
	CardExpirationYear int `json:"cardExpirationYear"`

	// BillingPhoneFax - no documentation
	BillingPhoneFax string `json:"billingPhoneFax"`

	// SerializedRequest - A serialized, delimited string of the transaction request sent to the financial
	// institution.
	SerializedRequest string `json:"serializedRequest"`

	// AccountId - The account ID to which the credit card and billing information is associated with.
	AccountId int `json:"accountId"`

	// BillingAddressLine2 - The second line in the address. Information such as suite number goes here.
	BillingAddressLine2 string `json:"billingAddressLine2"`

	// BillingEmail - The email address associated with a customer account.
	BillingEmail string `json:"billingEmail"`

	// Amount - The total amount of the attempted transaction, represented in decimal format as US Dollars
	Amount float64 `json:"amount"`

	// BillingNameLast - no documentation
	BillingNameLast string `json:"billingNameLast"`

	// OrderFromIpAddress - The IP address from which the transaction originates.
	OrderFromIpAddress string `json:"orderFromIpAddress"`

	// ReferenceCode - A code used by the financial institution to refer to the requested transaction.
	ReferenceCode string `json:"referenceCode"`

	// SerializedReply - A serialized, delimited string of the transaction request sent to the financial
	// institution.
	SerializedReply string `json:"serializedReply"`

	// BillingCountryCode - The 2-character Country code for an account's address. (i.e.
	BillingCountryCode string `json:"billingCountryCode"`

	// BillingState - no documentation
	BillingState string `json:"billingState"`

	// CardExpirationMonth - The month in which a customer's payment card will expire.
	CardExpirationMonth int `json:"cardExpirationMonth"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// InvoiceId - Unique identifier of the invoice to which funds will be applied.
	InvoiceId int `json:"invoiceId"`

	// BillingNameCompany - no documentation
	BillingNameCompany string `json:"billingNameCompany"`

	// Id - The unique identifier for a single credit card transaction request.
	Id int `json:"id"`
}

func (softlayer_billing_payment_card_transaction *SoftLayer_Billing_Payment_Card_Transaction) String() string {
	return "SoftLayer_Billing_Payment_Card_Transaction"
}

// SoftLayer_Billing_Payment_Card_Transaction_Extended is SoftLayer_Billing_Payment_Card_Transaction with all maskable types.
type SoftLayer_Billing_Payment_Card_Transaction_Extended struct {
	SoftLayer_Billing_Payment_Card_Transaction

	// Order - <nil>
	Order *SoftLayer_Billing_Order `json:"order"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`
}

func (softlayer_billing_payment_card_transaction *SoftLayer_Billing_Payment_Card_Transaction_Extended) String() string {
	return "SoftLayer_Billing_Payment_Card_Transaction"
}
