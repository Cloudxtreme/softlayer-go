package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Tax_Rates - This contains the four tax rates, one for each fee type.
type SoftLayer_Container_Tax_Rates struct {

	// LaborTaxRate - no documentation
	LaborTaxRate float64 `json:"laborTaxRate,omitempty"`

	// LocationId - no documentation
	LocationId float64 `json:"locationId,omitempty"`

	// OneTimeTaxRate - no documentation
	OneTimeTaxRate float64 `json:"oneTimeTaxRate,omitempty"`

	// RecurringTaxRate - no documentation
	RecurringTaxRate float64 `json:"recurringTaxRate,omitempty"`

	// SetupTaxRate - no documentation
	SetupTaxRate float64 `json:"setupTaxRate,omitempty"`
}

func (softlayer_container_tax_rates *SoftLayer_Container_Tax_Rates) String() string {
	return "SoftLayer_Container_Tax_Rates"
}
