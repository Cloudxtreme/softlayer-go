package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Catalog - A Catalog is defined as a set of prices for products that SoftLayer
// offers for sale. These prices are organized into packages which represent the different servers and
// services that SoftLayer offers.
type SoftLayer_Product_Catalog struct {

	// Products - no documentation
	Products []*SoftLayer_Product_Item `json:"products,omitempty"`

	// BrandCount - no documentation
	BrandCount uint64 `json:"brandCount,omitempty"`

	// PackageCount - no documentation
	PackageCount uint64 `json:"packageCount,omitempty"`

	// PriceCount - no documentation
	PriceCount uint64 `json:"priceCount,omitempty"`

	// ProductCount - no documentation
	ProductCount uint64 `json:"productCount,omitempty"`

	// Brands - no documentation
	Brands []*SoftLayer_Brand `json:"brands,omitempty"`

	// Packages - no documentation
	Packages []*SoftLayer_Product_Package `json:"packages,omitempty"`

	// Prices - no documentation
	Prices []*SoftLayer_Product_Item_Price `json:"prices,omitempty"`
}

func (softlayer_product_catalog *SoftLayer_Product_Catalog) String() string {
	return "SoftLayer_Product_Catalog"
}
