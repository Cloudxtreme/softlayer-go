package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Billing_Currency_ExchangeRate - <nil>
type SoftLayer_Billing_Currency_ExchangeRate struct {

	// EffectiveDate - <nil>
	EffectiveDate *time.Time `json:"effectiveDate,omitempty"`

	// ExpirationDate - <nil>
	ExpirationDate *time.Time `json:"expirationDate,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// Rate - <nil>
	Rate float64 `json:"rate,omitempty"`
}

func (softlayer_billing_currency_exchangerate *SoftLayer_Billing_Currency_ExchangeRate) String() string {
	return "SoftLayer_Billing_Currency_ExchangeRate"
}

// SoftLayer_Billing_Currency_ExchangeRate_Extended is SoftLayer_Billing_Currency_ExchangeRate with all maskable types.
type SoftLayer_Billing_Currency_ExchangeRate_Extended struct {
	SoftLayer_Billing_Currency_ExchangeRate

	// FundingCurrency - <nil>
	FundingCurrency *SoftLayer_Billing_Currency `json:"fundingCurrency,omitempty"`

	// LocalCurrency - <nil>
	LocalCurrency *SoftLayer_Billing_Currency `json:"localCurrency,omitempty"`
}

func (softlayer_billing_currency_exchangerate *SoftLayer_Billing_Currency_ExchangeRate_Extended) String() string {
	return "SoftLayer_Billing_Currency_ExchangeRate"
}
