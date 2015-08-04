package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Package_Items - This data type is a cross-reference between the
// SoftLayer_Product_Package and the SoftLayer_Product_Item(s) that belong in the
// SoftLayer_Product_Package.
type SoftLayer_Product_Package_Items struct {

	// Id - The unique identifier for this object. It is not used anywhere but in this object.
	Id string `json:"id,omitempty"`

	// ItemId - The SoftLayer_Product_Item id to which this instance of the object belongs.
	ItemId int `json:"itemId,omitempty"`

	// PackageId - The SoftLayer_Product_Package id to which this instance of the object belongs.
	PackageId int `json:"packageId,omitempty"`

	// Item - no documentation
	Item *SoftLayer_Product_Item `json:"item,omitempty"`

	// Package - no documentation
	Package *SoftLayer_Product_Package `json:"package,omitempty"`
}

func (softlayer_product_package_items *SoftLayer_Product_Package_Items) String() string {
	return "SoftLayer_Product_Package_Items"
}
