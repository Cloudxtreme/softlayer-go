package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "github.com/sudorandom/softlayer-go/slapi"
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

// GetAllCurrencyExchangeRates - <nil>
func (softlayer_billing_currency_exchangerate *SoftLayer_Billing_Currency_ExchangeRate) GetAllCurrencyExchangeRates(ctx *slapi.RequestContext, stringDate string) ([]*SoftLayer_Billing_Currency_ExchangeRate, error) {
	var returnValue []*SoftLayer_Billing_Currency_ExchangeRate
	return returnValue, nil
}

// GetCurrencies - <nil>
func (softlayer_billing_currency_exchangerate *SoftLayer_Billing_Currency_ExchangeRate) GetCurrencies(ctx *slapi.RequestContext) ([]*SoftLayer_Billing_Currency, error) {
	var returnValue []*SoftLayer_Billing_Currency
	return returnValue, nil
}

// GetExchangeRate - <nil>
func (softlayer_billing_currency_exchangerate *SoftLayer_Billing_Currency_ExchangeRate) GetExchangeRate(ctx *slapi.RequestContext, to string, from string, effectiveDate time.Time) (*SoftLayer_Billing_Currency_ExchangeRate, error) {
	var returnValue *SoftLayer_Billing_Currency_ExchangeRate
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_billing_currency_exchangerate *SoftLayer_Billing_Currency_ExchangeRate) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Billing_Currency_ExchangeRate, error) {
	var returnValue *SoftLayer_Billing_Currency_ExchangeRate
	return returnValue, nil
}

// GetPrice - <nil>
func (softlayer_billing_currency_exchangerate *SoftLayer_Billing_Currency_ExchangeRate) GetPrice(ctx *slapi.RequestContext, price float32, formatOptions SoftLayer_Container_Billing_Currency_Format) (string, error) {
	var returnValue string
	return returnValue, nil
}
