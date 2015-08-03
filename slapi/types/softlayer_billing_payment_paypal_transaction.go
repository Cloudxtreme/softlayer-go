package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Billing_Payment_PayPal_Transaction - The SoftLayer_Billing_Payment_PayPal_Transaction data
// type contains general information relating to attempted PayPal transactions.
type SoftLayer_Billing_Payment_PayPal_Transaction struct {

	// PayerCountry - no documentation
	PayerCountry string `json:"payerCountry"`

	// SerializedReply - A serialized, delimited string of the reply received from PayPal.
	SerializedReply string `json:"serializedReply"`

	// SettleAmount - no documentation
	SettleAmount float64 `json:"settleAmount"`

	// AddressPostalCode - no documentation
	AddressPostalCode string `json:"addressPostalCode"`

	// AddressStateProvence - State or Province in the address of the PayPal user.
	AddressStateProvence string `json:"addressStateProvence"`

	// PayerStatus - Current PayPal status associated with the user account.
	PayerStatus string `json:"payerStatus"`

	// PaymentType - PayPal defined code used to identify the type of payment. Provided in a PayPal
	// response.
	PaymentType string `json:"paymentType"`

	// PendingReason - Reason provided by PayPal for a payment given a pending status.
	PendingReason string `json:"pendingReason"`

	// AddressName - Name given to the address provided for the PayPal user.
	AddressName string `json:"addressName"`

	// FeeAmount - no documentation
	FeeAmount float64 `json:"feeAmount"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// AddressCityName - no documentation
	AddressCityName string `json:"addressCityName"`

	// LastPaypalCommand - The name of the command issued to PayPal with regards to the attempted
	// transaction.
	LastPaypalCommand string `json:"lastPaypalCommand"`

	// Payer - The PayPal user account name (email address) associated with the customer account.
	Payer string `json:"payer"`

	// PayerBusiness - The name of the business associated with the PayPal user.
	PayerBusiness string `json:"payerBusiness"`

	// TransactionId - Unique transaction ID provided in a PayPal response.
	TransactionId string `json:"transactionId"`

	// AddressStreet1 - First line of the street address of the PayPal user.
	AddressStreet1 string `json:"addressStreet1"`

	// ContactPhone - no documentation
	ContactPhone string `json:"contactPhone"`

	// PayerFirstName - no documentation
	PayerFirstName string `json:"payerFirstName"`

	// PayerId - no documentation
	PayerId string `json:"payerId"`

	// SerializedRequest - A serialized, delimited string of the request submitted to PayPal.
	SerializedRequest string `json:"serializedRequest"`

	// TaxAmount - no documentation
	TaxAmount float64 `json:"taxAmount"`

	// Token - Value issued by PayPal for referencing the attempted transaction.
	Token string `json:"token"`

	// AccountId - The account ID to which the PayPal and billing information is associated with.
	AccountId int `json:"accountId"`

	// TransactionType - PayPal defined code used to identify the type of transaction. Provided in a PayPal
	// response.
	TransactionType string `json:"transactionType"`

	// InvoiceId - Unique identifier of the invoice to which funds will be applied.
	InvoiceId int `json:"invoiceId"`

	// PayerLastName - no documentation
	PayerLastName string `json:"payerLastName"`

	// AddressStreet2 - Second line of the street address of the PayPal user.
	AddressStreet2 string `json:"addressStreet2"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// ExchangeRate - no documentation
	ExchangeRate string `json:"exchangeRate"`

	// OrderFromIpAddress - The IP address from where the PayPal payment request originated.
	OrderFromIpAddress string `json:"orderFromIpAddress"`

	// AddressCountry - Country given in the named address of the PayPal user.
	AddressCountry string `json:"addressCountry"`

	// GrossAmount - The total amount of the payment executed by PayPal, represented in decimal format as
	// US Dollars
	GrossAmount float64 `json:"grossAmount"`

	// Id - The unique identifier for a single PayPal transaction request.
	Id int `json:"id"`

	// OrderTotal - The amount of the payment submitted through the SoftLayer interface, represented in
	// decimal format as US Dollars
	OrderTotal float64 `json:"orderTotal"`

	// PaymentDate - Date that the payment was confirmed in PayPal by the user.
	PaymentDate *time.Time `json:"paymentDate"`

	// PaymentStatus - no documentation
	PaymentStatus string `json:"paymentStatus"`

	// AddressStatus - PayPal defined status of the address of the PayPal user.
	AddressStatus string `json:"addressStatus"`
}

// SoftLayer_Billing_Payment_PayPal_Transaction_Extended is SoftLayer_Billing_Payment_PayPal_Transaction with all maskable types.
type SoftLayer_Billing_Payment_PayPal_Transaction_Extended struct {
	SoftLayer_Billing_Payment_PayPal_Transaction

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// Order - <nil>
	Order *SoftLayer_Billing_Order `json:"order"`
}

func (softlayer_billing_payment_paypal_transaction *SoftLayer_Billing_Payment_PayPal_Transaction) String() string {
	return "SoftLayer_Billing_Payment_PayPal_Transaction"
}
