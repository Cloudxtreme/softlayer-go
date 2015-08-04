package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"

	time "time"
)

// SoftLayer_Billing_Currency_ExchangeRate - <nil>
type SoftLayer_Billing_Currency_ExchangeRate struct {

	// ExpirationDate - <nil>
	ExpirationDate *time.Time `json:"expirationDate,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// Rate - <nil>
	Rate slapi.Float64 `json:"rate,omitempty"`

	// EffectiveDate - <nil>
	EffectiveDate *time.Time `json:"effectiveDate,omitempty"`
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
