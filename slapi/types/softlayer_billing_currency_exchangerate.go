package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"

	time "time"
)

// SoftLayer_Billing_Currency_ExchangeRate - <nil>
type SoftLayer_Billing_Currency_ExchangeRate struct {

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// Rate - <nil>
	Rate slapi.Float64 `json:"rate,omitempty"`

	// EffectiveDate - <nil>
	EffectiveDate *time.Time `json:"effectiveDate,omitempty"`

	// ExpirationDate - <nil>
	ExpirationDate *time.Time `json:"expirationDate,omitempty"`

	// FundingCurrency - <nil>
	FundingCurrency *SoftLayer_Billing_Currency `json:"fundingCurrency,omitempty"`

	// LocalCurrency - <nil>
	LocalCurrency *SoftLayer_Billing_Currency `json:"localCurrency,omitempty"`
}

func (softlayer_billing_currency_exchangerate *SoftLayer_Billing_Currency_ExchangeRate) String() string {
	return "SoftLayer_Billing_Currency_ExchangeRate"
}
