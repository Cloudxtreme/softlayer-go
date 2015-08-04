package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Container_Billing_Currency_Format - <nil>
type SoftLayer_Container_Billing_Currency_Format struct {

	// Script - <nil>
	Script string `json:"script,omitempty"`

	// Service - <nil>
	Service string `json:"service,omitempty"`

	// Display - <nil>
	Display int `json:"display,omitempty"`

	// Format - <nil>
	Format string `json:"format,omitempty"`

	// Locale - <nil>
	Locale string `json:"locale,omitempty"`

	// Name - <nil>
	Name string `json:"name,omitempty"`

	// Tag - <nil>
	Tag string `json:"tag,omitempty"`

	// Value - <nil>
	Value slapi.Float64 `json:"value,omitempty"`

	// Currency - <nil>
	Currency string `json:"currency,omitempty"`

	// Position - <nil>
	Position int `json:"position,omitempty"`

	// Precision - <nil>
	Precision int `json:"precision,omitempty"`

	// Symbol - <nil>
	Symbol string `json:"symbol,omitempty"`
}

func (softlayer_container_billing_currency_format *SoftLayer_Container_Billing_Currency_Format) String() string {
	return "SoftLayer_Container_Billing_Currency_Format"
}
