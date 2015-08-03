package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Item_Bundles - The SoftLayer_Product_Item_Bundles contains item to price cross
// references Relates a category, price and item to a bundle. Match bundle ids to see all items and
// prices in a particular bundle.
type SoftLayer_Product_Item_Bundles struct {

	// BundleItem - no documentation
	BundleItem *SoftLayer_Product_Item `json:"bundleItem"`

	// BundleItemId - no documentation
	BundleItemId int `json:"bundleItemId"`

	// Category - no documentation
	Category *SoftLayer_Product_Item_Category `json:"category"`

	// Id - no documentation
	Id int `json:"id"`

	// ItemPrice - no documentation
	ItemPrice *SoftLayer_Product_Item_Price `json:"itemPrice"`

	// ItemPriceId - no documentation
	ItemPriceId int `json:"itemPriceId"`
}

func (softlayer_product_item_bundles *SoftLayer_Product_Item_Bundles) String() string {
	return "SoftLayer_Product_Item_Bundles"
}
