package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Billing_Invoice_Receivable_Payment - The SoftLayer_Billing_Invoice_Receivable_Payment data
// type contains general information relating to payments made against invoices.
type SoftLayer_Billing_Invoice_Receivable_Payment struct {

	// Amount - no documentation
	Amount float64 `json:"amount,omitempty"`

	// TypeCode - no documentation
	TypeCode string `json:"typeCode,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// InvoiceId - no documentation
	InvoiceId int `json:"invoiceId,omitempty"`
}

func (softlayer_billing_invoice_receivable_payment *SoftLayer_Billing_Invoice_Receivable_Payment) String() string {
	return "SoftLayer_Billing_Invoice_Receivable_Payment"
}

// SoftLayer_Billing_Invoice_Receivable_Payment_Extended is SoftLayer_Billing_Invoice_Receivable_Payment with all maskable types.
type SoftLayer_Billing_Invoice_Receivable_Payment_Extended struct {
	SoftLayer_Billing_Invoice_Receivable_Payment

	// Invoice - <nil>
	Invoice *SoftLayer_Billing_Invoice `json:"invoice,omitempty"`

	// CreditCardLastFourDigits - <nil>
	CreditCardLastFourDigits int `json:"creditCardLastFourDigits,omitempty"`

	// ExchangeRate - <nil>
	ExchangeRate *SoftLayer_Billing_Currency_ExchangeRate `json:"exchangeRate,omitempty"`

	// CreditCardTransaction - <nil>
	CreditCardTransaction *SoftLayer_Billing_Payment_Card_Transaction `json:"creditCardTransaction,omitempty"`

	// PaypalTransaction - <nil>
	PaypalTransaction *SoftLayer_Billing_Payment_PayPal_Transaction `json:"paypalTransaction,omitempty"`

	// Account - <nil>
	Account *SoftLayer_Account `json:"account,omitempty"`

	// CreditCardRequestId - <nil>
	CreditCardRequestId string `json:"creditCardRequestId,omitempty"`
}

func (softlayer_billing_invoice_receivable_payment *SoftLayer_Billing_Invoice_Receivable_Payment_Extended) String() string {
	return "SoftLayer_Billing_Invoice_Receivable_Payment"
}
