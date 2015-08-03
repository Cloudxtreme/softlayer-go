package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Product_Item_Category - The SoftLayer_Product_Item_Category data type contains general
// category information for prices.
type SoftLayer_Product_Item_Category struct {

	// BillingItemCount - A count of the billing items associated with an account that share a category
	// code with an item category's category code.
	BillingItemCount uint64 `json:"billingItemCount"`

	// BillingItems - The billing items associated with an account that share a category code with an item
	// category's category code.
	BillingItems []*SoftLayer_Billing_Item `json:"billingItems"`

	// CategoryCode - no documentation
	CategoryCode string `json:"categoryCode"`

	// Group - no documentation
	Group *SoftLayer_Product_Item_Category_Group `json:"group"`

	// GroupCount - A count of a collection of service offering category groups. Each group contains a
	// collection of items associated with this category.
	GroupCount uint64 `json:"groupCount"`

	// Groups - A collection of service offering category groups. Each group contains a collection of items
	// associated with this category.
	Groups []*SoftLayer_Product_Package_Item_Category_Group `json:"groups"`

	// Id - no documentation
	Id int `json:"id"`

	// Name - The friendly, descriptive name of the category as seen on the order forms and on invoices.
	Name string `json:"name"`

	// OrderOptionCount - A count of any unique options associated with an itme category.
	OrderOptionCount uint64 `json:"orderOptionCount"`

	// OrderOptions - Any unique options associated with an itme category.
	OrderOptions []*SoftLayer_Product_Item_Category_Order_Option_Type `json:"orderOptions"`

	// PackageConfigurationCount - A count of a list of configuration available in this category.'
	PackageConfigurationCount uint64 `json:"packageConfigurationCount"`

	// PackageConfigurations - A list of configuration available in this category.'
	PackageConfigurations []*SoftLayer_Product_Package_Order_Configuration `json:"packageConfigurations"`

	// PresetConfigurationCount - A count of a list of preset configurations this category is used in.'
	PresetConfigurationCount uint64 `json:"presetConfigurationCount"`

	// PresetConfigurations - A list of preset configurations this category is used in.'
	PresetConfigurations []*SoftLayer_Product_Package_Preset_Configuration `json:"presetConfigurations"`

	// QuantityLimit - Quantity that can be ordered. If 0, it will inherit the quantity from the server
	// quantity ordered. Otherwise it can be specified with the order separately
	QuantityLimit int `json:"quantityLimit"`

	// QuestionCount - A count of the questions that are associated with an item category.
	QuestionCount uint64 `json:"questionCount"`

	// QuestionReferenceCount - A count of the question references that are associated with an item
	// category.
	QuestionReferenceCount uint64 `json:"questionReferenceCount"`

	// QuestionReferences - The question references that are associated with an item category.
	QuestionReferences []*SoftLayer_Product_Item_Category_Question_Xref `json:"questionReferences"`

	// Questions - The questions that are associated with an item category.
	Questions []*SoftLayer_Product_Item_Category_Question `json:"questions"`
}

func (softlayer_product_item_category *SoftLayer_Product_Item_Category) String() string {
	return "SoftLayer_Product_Item_Category"
}

// GetAdditionalProductsForCategory - Returns a list of of active Items in the "Additional Services"
// package with their active prices for a given product item category and sorts them by price.
func (softlayer_product_item_category *SoftLayer_Product_Item_Category) GetAdditionalProductsForCategory(ctx *slapi.RequestContext) ([]*SoftLayer_Product_Item, error) {
	var returnValue []*SoftLayer_Product_Item
	return returnValue, nil
}

// GetBandwidthCategories - <nil>
func (softlayer_product_item_category *SoftLayer_Product_Item_Category) GetBandwidthCategories(ctx *slapi.RequestContext) ([]*SoftLayer_Product_Item_Category, error) {
	var returnValue []*SoftLayer_Product_Item_Category
	return returnValue, nil
}

// GetComputingCategories - This method returns a collection of computing categories. These categories
// are also top level items in a service offering.
func (softlayer_product_item_category *SoftLayer_Product_Item_Category) GetComputingCategories(ctx *slapi.RequestContext, resetCache bool) ([]*SoftLayer_Product_Item_Category, error) {
	var returnValue []*SoftLayer_Product_Item_Category
	return returnValue, nil
}

// GetCustomUsageRatesCategories - <nil>
func (softlayer_product_item_category *SoftLayer_Product_Item_Category) GetCustomUsageRatesCategories(ctx *slapi.RequestContext) ([]*SoftLayer_Product_Item_Category, error) {
	var returnValue []*SoftLayer_Product_Item_Category
	return returnValue, nil
}

// GetObject - Each product item price must be tied to a category for it to be sold. These categories
// describe how a particular product item is sold. For example, the 250GB hard drive can be sold as
// disk0, disk1, ... disk11. There are different prices for this product item depending on which
// category it is. This keeps down the number of products in total.
func (softlayer_product_item_category *SoftLayer_Product_Item_Category) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Product_Item_Category, error) {
	var returnValue *SoftLayer_Product_Item_Category
	return returnValue, nil
}

// GetSoftwareCategories - <nil>
func (softlayer_product_item_category *SoftLayer_Product_Item_Category) GetSoftwareCategories(ctx *slapi.RequestContext) ([]*SoftLayer_Product_Item_Category, error) {
	var returnValue []*SoftLayer_Product_Item_Category
	return returnValue, nil
}

// GetSubnetCategories - <nil>
func (softlayer_product_item_category *SoftLayer_Product_Item_Category) GetSubnetCategories(ctx *slapi.RequestContext) ([]*SoftLayer_Product_Item_Category, error) {
	var returnValue []*SoftLayer_Product_Item_Category
	return returnValue, nil
}

// GetTopLevelCategories - This method returns a collection of computing categories. These categories
// are also top level items in a service offering.
func (softlayer_product_item_category *SoftLayer_Product_Item_Category) GetTopLevelCategories(ctx *slapi.RequestContext, resetCache bool) ([]*SoftLayer_Product_Item_Category, error) {
	var returnValue []*SoftLayer_Product_Item_Category
	return returnValue, nil
}

// GetValidCancelableServiceItemCategories - This method returns service product categories that can be
// canceled via You can use these categories to find the billing items you wish to cancel.
func (softlayer_product_item_category *SoftLayer_Product_Item_Category) GetValidCancelableServiceItemCategories(ctx *slapi.RequestContext) ([]*SoftLayer_Product_Item_Category, error) {
	var returnValue []*SoftLayer_Product_Item_Category
	return returnValue, nil
}

// GetVlanCategories - <nil>
func (softlayer_product_item_category *SoftLayer_Product_Item_Category) GetVlanCategories(ctx *slapi.RequestContext) ([]*SoftLayer_Product_Item_Category, error) {
	var returnValue []*SoftLayer_Product_Item_Category
	return returnValue, nil
}
