package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Billing_Invoice_Receivable_Payment - The SoftLayer_Billing_Invoice_Receivable_Payment data
// type contains general information relating to payments made against invoices.
type SoftLayer_Billing_Invoice_Receivable_Payment struct {

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`

	// Amount - no documentation
	Amount float64 `json:"amount"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// CreditCardLastFourDigits - <nil>
	CreditCardLastFourDigits int `json:"creditCardLastFourDigits"`

	// CreditCardRequestId - <nil>
	CreditCardRequestId string `json:"creditCardRequestId"`

	// CreditCardTransaction - <nil>
	CreditCardTransaction *SoftLayer_Billing_Payment_Card_Transaction `json:"creditCardTransaction"`

	// ExchangeRate - <nil>
	ExchangeRate *SoftLayer_Billing_Currency_ExchangeRate `json:"exchangeRate"`

	// Invoice - <nil>
	Invoice *SoftLayer_Billing_Invoice `json:"invoice"`

	// InvoiceId - no documentation
	InvoiceId int `json:"invoiceId"`

	// PaypalTransaction - <nil>
	PaypalTransaction *SoftLayer_Billing_Payment_PayPal_Transaction `json:"paypalTransaction"`

	// TypeCode - no documentation
	TypeCode string `json:"typeCode"`
}
