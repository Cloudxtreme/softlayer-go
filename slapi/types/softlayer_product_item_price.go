package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Item_Price - The SoftLayer_Product_Item_Price data type contains general
// information relating to a single SoftLayer product item price. You can find out what packages each
// price is in as well as which category under which this price is sold. All prices are returned in
// floating point values measured in US Dollars
type SoftLayer_Product_Item_Price struct {

	// RecurringFee - A recurring fee is a fee that happens every billing period. This fee is represented
	// as a floating point decimal in US dollars
	RecurringFee float64 `json:"recurringFee"`

	// ItemId - no documentation
	ItemId int `json:"itemId"`

	// Quantity - <nil>
	Quantity int `json:"quantity"`

	// UsageRate - no documentation
	UsageRate float64 `json:"usageRate"`

	// LaborFee - no documentation
	LaborFee float64 `json:"laborFee"`

	// OneTimeFee - no documentation
	OneTimeFee float64 `json:"oneTimeFee"`

	// OneTimeFeeTax - A price's total tax amount of the one time fees (oneTimeFee, laborFee, and
	// setupFee). This is only populated after the order is verified via
	// SoftLayer_Product_Order::verifyOrder()
	OneTimeFeeTax float64 `json:"oneTimeFeeTax"`

	// OrderOptions - Order options for the category that this price is associated with.
	OrderOptions []*SoftLayer_Product_Item_Category_Order_Option_Type `json:"orderOptions"`

	// LocationGroupId - The id of the [[SoftLayer_Location_Group_Pricing]] that this price is part of. If
	// set to null, the price is considered a standard price, which can be used with any location when
	// ordering. During order [[SoftLayer_Product_Order/verifyOrder|verification]] and
	// [[SoftLayer_Product_Order/placeOrder|placement]], if a standard price is used, that price may be
	// replaced with a location based price, which does not have this property set to null. The location
	// based price must be part of a [[SoftLayer_Location_Group_Pricing]] that has the location being
	// ordered in order for this to happen.
	LocationGroupId int `json:"locationGroupId"`

	// RecurringFeeTax - A price's tax amount of the recurring fee. This is only populated after the order
	// is verified via SoftLayer_Product_Order::verifyOrder()
	RecurringFeeTax float64 `json:"recurringFeeTax"`

	// Sort - no documentation
	Sort int `json:"sort"`

	// OnSaleFlag - no documentation
	OnSaleFlag bool `json:"onSaleFlag"`

	// HourlyRecurringFee - The hourly price for this item, should this item be part of an hourly pricing
	// package.
	HourlyRecurringFee float64 `json:"hourlyRecurringFee"`

	// Id - no documentation
	Id int `json:"id"`

	// ProratedRecurringFee - A recurring fee is a fee that happens every billing period. This fee is
	// represented as a floating point decimal in US dollars
	ProratedRecurringFee float64 `json:"proratedRecurringFee"`

	// ProratedRecurringFeeTax - A price's tax amount of the recurring fee. This is only populated after
	// the order is verified via SoftLayer_Product_Order::verifyOrder()
	ProratedRecurringFeeTax float64 `json:"proratedRecurringFeeTax"`

	// CurrentPriceFlag - This flag is used by the [[SoftLayer_Hardware::getUpgradeItems|getUpgradeItems]]
	// method to indicate if a product price is used for the current billing item.
	CurrentPriceFlag bool `json:"currentPriceFlag"`

	// SetupFee - The setup fee associated with a product item price.
	SetupFee float64 `json:"setupFee"`
}

// SoftLayer_Product_Item_Price_Extended is SoftLayer_Product_Item_Price with all maskable types.
type SoftLayer_Product_Item_Price_Extended struct {
	SoftLayer_Product_Item_Price

	// PackageReferences - no documentation
	PackageReferences []*SoftLayer_Product_Package_Item_Prices `json:"packageReferences"`

	// BigDataOsJournalDiskFlag - Whether the price is for Big Data OS/Journal disks only. (Deprecated)
	BigDataOsJournalDiskFlag bool `json:"bigDataOsJournalDiskFlag"`

	// PackageReferenceCount - no documentation
	PackageReferenceCount uint64 `json:"packageReferenceCount"`

	// Attributes - <nil>
	Attributes []*SoftLayer_Product_Item_Price_Attribute `json:"attributes"`

	// BundleReferences - no documentation
	BundleReferences []*SoftLayer_Product_Item_Bundles `json:"bundleReferences"`

	// CapacityRestrictionMinimum - The minimum capacity value for which this price is suitable.
	CapacityRestrictionMinimum string `json:"capacityRestrictionMinimum"`

	// PresetConfigurations - A list of preset configurations this price is used in.'
	PresetConfigurations []*SoftLayer_Product_Package_Preset_Configuration `json:"presetConfigurations"`

	// AccountRestrictionCount - A count of the account that the item price is restricted to.
	AccountRestrictionCount uint64 `json:"accountRestrictionCount"`

	// AttributeCount - no documentation
	AttributeCount uint64 `json:"attributeCount"`

	// InventoryCount - A count of an item price's inventory status per datacenter.
	InventoryCount uint64 `json:"inventoryCount"`

	// DefinedSoftwareLicenseFlag - Whether this price defines a software license for its product item.
	DefinedSoftwareLicenseFlag bool `json:"definedSoftwareLicenseFlag"`

	// CategoryCount - A count of all categories which this item is a member.
	CategoryCount uint64 `json:"categoryCount"`

	// Categories - no documentation
	Categories []*SoftLayer_Product_Item_Category `json:"categories"`

	// AccountRestrictions - no documentation
	AccountRestrictions []*SoftLayer_Product_Item_Price_Account_Restriction `json:"accountRestrictions"`

	// CapacityRestrictionType - The type of capacity restriction by which this price must abide.
	CapacityRestrictionType string `json:"capacityRestrictionType"`

	// Inventory - no documentation
	Inventory []*SoftLayer_Product_Package_Inventory `json:"inventory"`

	// PackageCount - A count of a price's packages under which this item is sold.
	PackageCount uint64 `json:"packageCount"`

	// PresetConfigurationCount - A count of a list of preset configurations this price is used in.'
	PresetConfigurationCount uint64 `json:"presetConfigurationCount"`

	// Item - no documentation
	Item *SoftLayer_Product_Item `json:"item"`

	// OrderPremiums - <nil>
	OrderPremiums []*SoftLayer_Product_Item_Price_Premium `json:"orderPremiums"`

	// Packages - no documentation
	Packages []*SoftLayer_Product_Package `json:"packages"`

	// PricingLocationGroup - The pricing location group that this price is applicable for. Prices that
	// have a pricing location group will only be available for ordering with the locations specified on
	// the location group.
	PricingLocationGroup *SoftLayer_Location_Group_Pricing `json:"pricingLocationGroup"`

	// RequiredCoreCount - The number of server cores required to order this item.
	RequiredCoreCount string `json:"requiredCoreCount"`

	// CapacityRestrictionMaximum - The maximum capacity value for which this price is suitable.
	CapacityRestrictionMaximum string `json:"capacityRestrictionMaximum"`

	// BundleReferenceCount - no documentation
	BundleReferenceCount uint64 `json:"bundleReferenceCount"`

	// OrderPremiumCount - no documentation
	OrderPremiumCount uint64 `json:"orderPremiumCount"`
}

func (softlayer_product_item_price *SoftLayer_Product_Item_Price) String() string {
	return "SoftLayer_Product_Item_Price"
}
