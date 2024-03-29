package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Package_Order_Configuration - This datatype describes the item categories that are
// required for each package to be ordered. For instance, for package 2, there will be many required
// categories. When submitting an order for a server, there must be at most 1 price for each category
// whose "isRequired" is set. Examples of required categories: - server - ram - bandwidth - disk0 There
// are others, but these are the main ones. For each required category, a SoftLayer_Product_Item_Price
// must be chosen that is valid for the package.
type SoftLayer_Product_Package_Order_Configuration struct {

	// OrderStepId - The order step ID for this particular option in the package.
	OrderStepId int `json:"orderStepId,omitempty"`

	// ErrorMessage - The error message displayed if the submitted order does not contain this item
	// category, if it is required.
	ErrorMessage string `json:"errorMessage,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// ItemCategoryId - no documentation
	ItemCategoryId int `json:"itemCategoryId,omitempty"`

	// PackageId - no documentation
	PackageId int `json:"packageId,omitempty"`

	// Sort - This is an integer used to show the order in which each item Category should be displayed.
	// This is merely the suggested order.
	Sort int `json:"sort,omitempty"`

	// IsRequired - This is a flag which tells SoftLayer_Product_Order::verifyOrder() whether or not this
	// category is required. If this is set, then the order submitted must contain a
	// SoftLayer_Product_Item_Price with this category as part of the order.
	IsRequired int `json:"isRequired,omitempty"`

	// Step - no documentation
	Step *SoftLayer_Product_Package_Order_Step `json:"step,omitempty"`

	// Package - no documentation
	Package *SoftLayer_Product_Package `json:"package,omitempty"`

	// ItemCategory - no documentation
	ItemCategory *SoftLayer_Product_Item_Category `json:"itemCategory,omitempty"`
}

func (softlayer_product_package_order_configuration *SoftLayer_Product_Package_Order_Configuration) String() string {
	return "SoftLayer_Product_Package_Order_Configuration"
}
