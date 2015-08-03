package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Package_Items - This data type is a cross-reference between the
// SoftLayer_Product_Package and the SoftLayer_Product_Item(s) that belong in the
// SoftLayer_Product_Package.
type SoftLayer_Product_Package_Items struct {

	// Id - The unique identifier for this object. It is not used anywhere but in this object.
	Id string `json:"id"`

	// ItemId - The SoftLayer_Product_Item id to which this instance of the object belongs.
	ItemId int `json:"itemId"`

	// PackageId - The SoftLayer_Product_Package id to which this instance of the object belongs.
	PackageId int `json:"packageId"`
}

// SoftLayer_Product_Package_Items_Extended is SoftLayer_Product_Package_Items with all maskable types.
type SoftLayer_Product_Package_Items_Extended struct {
	SoftLayer_Product_Package_Items

	// Item - no documentation
	Item *SoftLayer_Product_Item `json:"item"`

	// Package - no documentation
	Package *SoftLayer_Product_Package `json:"package"`
}

func (softlayer_product_package_items *SoftLayer_Product_Package_Items) String() string {
	return "SoftLayer_Product_Package_Items"
}
