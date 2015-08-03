package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Billing_Currency_ExchangeRate - <nil>
type SoftLayer_Billing_Currency_ExchangeRate struct {

	// EffectiveDate - <nil>
	EffectiveDate *time.Time `json:"effectiveDate"`

	// ExpirationDate - <nil>
	ExpirationDate *time.Time `json:"expirationDate"`

	// FundingCurrency - <nil>
	FundingCurrency *SoftLayer_Billing_Currency `json:"fundingCurrency"`

	// Id - no documentation
	Id int `json:"id"`

	// LocalCurrency - <nil>
	LocalCurrency *SoftLayer_Billing_Currency `json:"localCurrency"`

	// Rate - <nil>
	Rate float64 `json:"rate"`
}

func (softlayer_billing_currency_exchangerate *SoftLayer_Billing_Currency_ExchangeRate) String() string {
	return "SoftLayer_Billing_Currency_ExchangeRate"
}
