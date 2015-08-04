package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Container_Tax_Cache - These are the results of a tax calculation. The tax calculation was
// kicked off but allowed to run in the background. This type stores the results so that an interface
// can be updated with up-to-date information.
type SoftLayer_Container_Tax_Cache struct {

	// Status - The status of the tax request. This should be or
	Status string `json:"status,omitempty"`

	// TotalTaxAmount - no documentation
	TotalTaxAmount slapi.Float64 `json:"totalTaxAmount,omitempty"`

	// EffectiveTaxRate - The percentage of the final total that should be tax.
	EffectiveTaxRate slapi.Float64 `json:"effectiveTaxRate,omitempty"`

	// Items - The container that holds the four actual tax rates, one for each fee type.
	Items []*SoftLayer_Container_Tax_Cache_Item `json:"items,omitempty"`
}

func (softlayer_container_tax_cache *SoftLayer_Container_Tax_Cache) String() string {
	return "SoftLayer_Container_Tax_Cache"
}
