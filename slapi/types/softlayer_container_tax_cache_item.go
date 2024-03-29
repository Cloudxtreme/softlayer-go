package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Tax_Cache_Item - This represents one order item in a tax calculation.
type SoftLayer_Container_Tax_Cache_Item struct {

	// CategoryCode - no documentation
	CategoryCode string `json:"categoryCode,omitempty"`

	// ContainerHash - This hash will match to the hash on an order container.
	ContainerHash string `json:"containerHash,omitempty"`

	// ItemPriceId - no documentation
	ItemPriceId int `json:"itemPriceId,omitempty"`

	// TaxRates - This is the container containing the individual tax rates.
	TaxRates *SoftLayer_Container_Tax_Rates `json:"taxRates,omitempty"`
}

func (softlayer_container_tax_cache_item *SoftLayer_Container_Tax_Cache_Item) String() string {
	return "SoftLayer_Container_Tax_Cache_Item"
}
