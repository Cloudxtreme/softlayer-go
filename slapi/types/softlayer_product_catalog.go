package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Catalog - A Catalog is defined as a set of prices for products that SoftLayer
// offers for sale. These prices are organized into packages which represent the different servers and
// services that SoftLayer offers.
type SoftLayer_Product_Catalog struct {
}

func (softlayer_product_catalog *SoftLayer_Product_Catalog) String() string {
	return "SoftLayer_Product_Catalog"
}

// SoftLayer_Product_Catalog_Extended is SoftLayer_Product_Catalog with all maskable types.
type SoftLayer_Product_Catalog_Extended struct {
	SoftLayer_Product_Catalog

	// PriceCount - no documentation
	PriceCount uint64 `json:"priceCount"`

	// ProductCount - no documentation
	ProductCount uint64 `json:"productCount"`

	// Brands - no documentation
	Brands []*SoftLayer_Brand `json:"brands"`

	// Packages - no documentation
	Packages []*SoftLayer_Product_Package `json:"packages"`

	// Prices - no documentation
	Prices []*SoftLayer_Product_Item_Price `json:"prices"`

	// Products - no documentation
	Products []*SoftLayer_Product_Item `json:"products"`

	// BrandCount - no documentation
	BrandCount uint64 `json:"brandCount"`

	// PackageCount - no documentation
	PackageCount uint64 `json:"packageCount"`
}

func (softlayer_product_catalog *SoftLayer_Product_Catalog_Extended) String() string {
	return "SoftLayer_Product_Catalog"
}
