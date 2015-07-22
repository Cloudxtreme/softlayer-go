package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Package_Item_Category_Group - This class is used to organize categories for a
// service offering. A service offering (usually) contains multiple categories (e.g., server, os,
// disk0, ram). This class allows us to organize the prices into related item category groups.
type SoftLayer_Product_Package_Item_Category_Group struct {

	// Category - <nil>
	Category *SoftLayer_Product_Item_Category `json:"category"`

	// ItemCategoryId - no documentation
	ItemCategoryId int `json:"itemCategoryId"`

	// Package - <nil>
	Package *SoftLayer_Product_Package `json:"package"`

	// PackageId - The service offering id associated with this group.
	PackageId int `json:"packageId"`

	// PriceCount - no documentation
	PriceCount uint64 `json:"priceCount"`

	// Prices - <nil>
	Prices []*SoftLayer_Product_Item_Price `json:"prices"`

	// Sort - no documentation
	Sort int `json:"sort"`

	// Title - An optional title associated with this group. E.g., for operating systems, this will be the
	// manufacturer.
	Title string `json:"title"`
}
