package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"

	time "time"
)

// SoftLayer_Billing_Invoice_Receivable_Payment - The SoftLayer_Billing_Invoice_Receivable_Payment data
// type contains general information relating to payments made against invoices.
type SoftLayer_Billing_Invoice_Receivable_Payment struct {

	// TypeCode - no documentation
	TypeCode string `json:"typeCode,omitempty"`

	// Amount - no documentation
	Amount slapi.Float64 `json:"amount,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// InvoiceId - no documentation
	InvoiceId int `json:"invoiceId,omitempty"`

	// CreditCardLastFourDigits - <nil>
	CreditCardLastFourDigits int `json:"creditCardLastFourDigits,omitempty"`

	// Account - <nil>
	Account *SoftLayer_Account `json:"account,omitempty"`

	// CreditCardRequestId - <nil>
	CreditCardRequestId string `json:"creditCardRequestId,omitempty"`

	// CreditCardTransaction - <nil>
	CreditCardTransaction *SoftLayer_Billing_Payment_Card_Transaction `json:"creditCardTransaction,omitempty"`

	// ExchangeRate - <nil>
	ExchangeRate *SoftLayer_Billing_Currency_ExchangeRate `json:"exchangeRate,omitempty"`

	// Invoice - <nil>
	Invoice *SoftLayer_Billing_Invoice `json:"invoice,omitempty"`

	// PaypalTransaction - <nil>
	PaypalTransaction *SoftLayer_Billing_Payment_PayPal_Transaction `json:"paypalTransaction,omitempty"`
}

func (softlayer_billing_invoice_receivable_payment *SoftLayer_Billing_Invoice_Receivable_Payment) String() string {
	return "SoftLayer_Billing_Invoice_Receivable_Payment"
}
