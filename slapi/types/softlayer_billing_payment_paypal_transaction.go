package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"

	time "time"
)

// SoftLayer_Billing_Payment_PayPal_Transaction - The SoftLayer_Billing_Payment_PayPal_Transaction data
// type contains general information relating to attempted PayPal transactions.
type SoftLayer_Billing_Payment_PayPal_Transaction struct {

	// AddressCountry - Country given in the named address of the PayPal user.
	AddressCountry string `json:"addressCountry,omitempty"`

	// AddressCityName - no documentation
	AddressCityName string `json:"addressCityName,omitempty"`

	// AddressStatus - PayPal defined status of the address of the PayPal user.
	AddressStatus string `json:"addressStatus,omitempty"`

	// ExchangeRate - no documentation
	ExchangeRate string `json:"exchangeRate,omitempty"`

	// PayerLastName - no documentation
	PayerLastName string `json:"payerLastName,omitempty"`

	// SerializedReply - A serialized, delimited string of the reply received from PayPal.
	SerializedReply string `json:"serializedReply,omitempty"`

	// AddressStreet2 - Second line of the street address of the PayPal user.
	AddressStreet2 string `json:"addressStreet2,omitempty"`

	// GrossAmount - The total amount of the payment executed by PayPal, represented in decimal format as
	// US Dollars
	GrossAmount slapi.Float64 `json:"grossAmount,omitempty"`

	// InvoiceId - Unique identifier of the invoice to which funds will be applied.
	InvoiceId int `json:"invoiceId,omitempty"`

	// OrderTotal - The amount of the payment submitted through the SoftLayer interface, represented in
	// decimal format as US Dollars
	OrderTotal slapi.Float64 `json:"orderTotal,omitempty"`

	// PayerCountry - no documentation
	PayerCountry string `json:"payerCountry,omitempty"`

	// SerializedRequest - A serialized, delimited string of the request submitted to PayPal.
	SerializedRequest string `json:"serializedRequest,omitempty"`

	// Token - Value issued by PayPal for referencing the attempted transaction.
	Token string `json:"token,omitempty"`

	// AddressName - Name given to the address provided for the PayPal user.
	AddressName string `json:"addressName,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// LastPaypalCommand - The name of the command issued to PayPal with regards to the attempted
	// transaction.
	LastPaypalCommand string `json:"lastPaypalCommand,omitempty"`

	// PaymentType - PayPal defined code used to identify the type of payment. Provided in a PayPal
	// response.
	PaymentType string `json:"paymentType,omitempty"`

	// FeeAmount - no documentation
	FeeAmount slapi.Float64 `json:"feeAmount,omitempty"`

	// PayerFirstName - no documentation
	PayerFirstName string `json:"payerFirstName,omitempty"`

	// PayerStatus - Current PayPal status associated with the user account.
	PayerStatus string `json:"payerStatus,omitempty"`

	// TransactionType - PayPal defined code used to identify the type of transaction. Provided in a PayPal
	// response.
	TransactionType string `json:"transactionType,omitempty"`

	// AddressPostalCode - no documentation
	AddressPostalCode string `json:"addressPostalCode,omitempty"`

	// AddressStateProvence - State or Province in the address of the PayPal user.
	AddressStateProvence string `json:"addressStateProvence,omitempty"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// OrderFromIpAddress - The IP address from where the PayPal payment request originated.
	OrderFromIpAddress string `json:"orderFromIpAddress,omitempty"`

	// Payer - The PayPal user account name (email address) associated with the customer account.
	Payer string `json:"payer,omitempty"`

	// PayerId - no documentation
	PayerId string `json:"payerId,omitempty"`

	// TaxAmount - no documentation
	TaxAmount slapi.Float64 `json:"taxAmount,omitempty"`

	// TransactionId - Unique transaction ID provided in a PayPal response.
	TransactionId string `json:"transactionId,omitempty"`

	// Id - The unique identifier for a single PayPal transaction request.
	Id int `json:"id,omitempty"`

	// PayerBusiness - The name of the business associated with the PayPal user.
	PayerBusiness string `json:"payerBusiness,omitempty"`

	// PaymentDate - Date that the payment was confirmed in PayPal by the user.
	PaymentDate *time.Time `json:"paymentDate,omitempty"`

	// PendingReason - Reason provided by PayPal for a payment given a pending status.
	PendingReason string `json:"pendingReason,omitempty"`

	// SettleAmount - no documentation
	SettleAmount slapi.Float64 `json:"settleAmount,omitempty"`

	// AccountId - The account ID to which the PayPal and billing information is associated with.
	AccountId int `json:"accountId,omitempty"`

	// AddressStreet1 - First line of the street address of the PayPal user.
	AddressStreet1 string `json:"addressStreet1,omitempty"`

	// ContactPhone - no documentation
	ContactPhone string `json:"contactPhone,omitempty"`

	// PaymentStatus - no documentation
	PaymentStatus string `json:"paymentStatus,omitempty"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account,omitempty"`

	// Order - <nil>
	Order *SoftLayer_Billing_Order `json:"order,omitempty"`
}

func (softlayer_billing_payment_paypal_transaction *SoftLayer_Billing_Payment_PayPal_Transaction) String() string {
	return "SoftLayer_Billing_Payment_PayPal_Transaction"
}
