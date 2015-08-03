package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Product_Item_Category_Group - The SoftLayer_Product_Item_Category_Group data type contains
// general category group information.
type SoftLayer_Product_Item_Category_Group struct {

	// Id - no documentation
	Id int `json:"id"`

	// Name - The friendly, descriptive name of the category group as seen on the order forms and on
	// invoices.
	Name string `json:"name"`
}

func (softlayer_product_item_category_group *SoftLayer_Product_Item_Category_Group) String() string {
	return "SoftLayer_Product_Item_Category_Group"
}

// GetObject - Each product item category must be tied to a category group. These category groups
// describe how a particular product item category is categorized. For example, the disk0, disk1, ...
// disk11 can be categorized as Server and Attached Services. There are different groups for each of
// this product item category depending on the function of the item product in the subject category.
func (softlayer_product_item_category_group *SoftLayer_Product_Item_Category_Group) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Product_Item_Category_Group, error) {
	var returnValue *SoftLayer_Product_Item_Category_Group
	return returnValue, nil
}
