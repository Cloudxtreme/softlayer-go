package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Tax_Cache - These are the results of a tax calculation. The tax calculation was
// kicked off but allowed to run in the background. This type stores the results so that an interface
// can be updated with up-to-date information.
type SoftLayer_Container_Tax_Cache struct {

	// Status - The status of the tax request. This should be or
	Status string `json:"status"`

	// TotalTaxAmount - no documentation
	TotalTaxAmount float64 `json:"totalTaxAmount"`

	// EffectiveTaxRate - The percentage of the final total that should be tax.
	EffectiveTaxRate float64 `json:"effectiveTaxRate"`

	// Items - The container that holds the four actual tax rates, one for each fee type.
	Items []*SoftLayer_Container_Tax_Cache_Item `json:"items"`
}

func (softlayer_container_tax_cache *SoftLayer_Container_Tax_Cache) String() string {
	return "SoftLayer_Container_Tax_Cache"
}
