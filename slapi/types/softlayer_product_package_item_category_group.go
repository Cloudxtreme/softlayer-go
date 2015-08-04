package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Package_Item_Category_Group - This class is used to organize categories for a
// service offering. A service offering (usually) contains multiple categories (e.g., server, os,
// disk0, ram). This class allows us to organize the prices into related item category groups.
type SoftLayer_Product_Package_Item_Category_Group struct {

	// Title - An optional title associated with this group. E.g., for operating systems, this will be the
	// manufacturer.
	Title string `json:"title,omitempty"`

	// ItemCategoryId - no documentation
	ItemCategoryId int `json:"itemCategoryId,omitempty"`

	// PackageId - The service offering id associated with this group.
	PackageId int `json:"packageId,omitempty"`

	// Sort - no documentation
	Sort int `json:"sort,omitempty"`
}

func (softlayer_product_package_item_category_group *SoftLayer_Product_Package_Item_Category_Group) String() string {
	return "SoftLayer_Product_Package_Item_Category_Group"
}

// SoftLayer_Product_Package_Item_Category_Group_Extended is SoftLayer_Product_Package_Item_Category_Group with all maskable types.
type SoftLayer_Product_Package_Item_Category_Group_Extended struct {
	SoftLayer_Product_Package_Item_Category_Group

	// Category - <nil>
	Category *SoftLayer_Product_Item_Category `json:"category,omitempty"`

	// Package - <nil>
	Package *SoftLayer_Product_Package `json:"package,omitempty"`

	// Prices - <nil>
	Prices []*SoftLayer_Product_Item_Price `json:"prices,omitempty"`

	// PriceCount - no documentation
	PriceCount uint64 `json:"priceCount,omitempty"`
}

func (softlayer_product_package_item_category_group *SoftLayer_Product_Package_Item_Category_Group_Extended) String() string {
	return "SoftLayer_Product_Package_Item_Category_Group"
}
