package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Product_Item_Price - The SoftLayer_Product_Item_Price data type contains general
// information relating to a single SoftLayer product item price. You can find out what packages each
// price is in as well as which category under which this price is sold. All prices are returned in
// floating point values measured in US Dollars
type SoftLayer_Product_Item_Price struct {

	// LaborFee - no documentation
	LaborFee slapi.Float64 `json:"laborFee,omitempty"`

	// LocationGroupId - The id of the [[SoftLayer_Location_Group_Pricing]] that this price is part of. If
	// set to null, the price is considered a standard price, which can be used with any location when
	// ordering. During order [[SoftLayer_Product_Order/verifyOrder|verification]] and
	// [[SoftLayer_Product_Order/placeOrder|placement]], if a standard price is used, that price may be
	// replaced with a location based price, which does not have this property set to null. The location
	// based price must be part of a [[SoftLayer_Location_Group_Pricing]] that has the location being
	// ordered in order for this to happen.
	LocationGroupId int `json:"locationGroupId,omitempty"`

	// OnSaleFlag - no documentation
	OnSaleFlag bool `json:"onSaleFlag,omitempty"`

	// OrderOptions - Order options for the category that this price is associated with.
	OrderOptions []*SoftLayer_Product_Item_Category_Order_Option_Type `json:"orderOptions,omitempty"`

	// SetupFee - The setup fee associated with a product item price.
	SetupFee slapi.Float64 `json:"setupFee,omitempty"`

	// UsageRate - no documentation
	UsageRate slapi.Float64 `json:"usageRate,omitempty"`

	// ProratedRecurringFeeTax - A price's tax amount of the recurring fee. This is only populated after
	// the order is verified via SoftLayer_Product_Order::verifyOrder()
	ProratedRecurringFeeTax slapi.Float64 `json:"proratedRecurringFeeTax,omitempty"`

	// ItemId - no documentation
	ItemId int `json:"itemId,omitempty"`

	// OneTimeFeeTax - A price's total tax amount of the one time fees (oneTimeFee, laborFee, and
	// setupFee). This is only populated after the order is verified via
	// SoftLayer_Product_Order::verifyOrder()
	OneTimeFeeTax slapi.Float64 `json:"oneTimeFeeTax,omitempty"`

	// Quantity - <nil>
	Quantity int `json:"quantity,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// CurrentPriceFlag - This flag is used by the [[SoftLayer_Hardware::getUpgradeItems|getUpgradeItems]]
	// method to indicate if a product price is used for the current billing item.
	CurrentPriceFlag bool `json:"currentPriceFlag,omitempty"`

	// HourlyRecurringFee - The hourly price for this item, should this item be part of an hourly pricing
	// package.
	HourlyRecurringFee slapi.Float64 `json:"hourlyRecurringFee,omitempty"`

	// RecurringFeeTax - A price's tax amount of the recurring fee. This is only populated after the order
	// is verified via SoftLayer_Product_Order::verifyOrder()
	RecurringFeeTax slapi.Float64 `json:"recurringFeeTax,omitempty"`

	// Sort - no documentation
	Sort int `json:"sort,omitempty"`

	// OneTimeFee - no documentation
	OneTimeFee slapi.Float64 `json:"oneTimeFee,omitempty"`

	// ProratedRecurringFee - A recurring fee is a fee that happens every billing period. This fee is
	// represented as a floating point decimal in US dollars
	ProratedRecurringFee slapi.Float64 `json:"proratedRecurringFee,omitempty"`

	// RecurringFee - A recurring fee is a fee that happens every billing period. This fee is represented
	// as a floating point decimal in US dollars
	RecurringFee slapi.Float64 `json:"recurringFee,omitempty"`
}

func (softlayer_product_item_price *SoftLayer_Product_Item_Price) String() string {
	return "SoftLayer_Product_Item_Price"
}

// SoftLayer_Product_Item_Price_Extended is SoftLayer_Product_Item_Price with all maskable types.
type SoftLayer_Product_Item_Price_Extended struct {
	SoftLayer_Product_Item_Price

	// Categories - no documentation
	Categories []*SoftLayer_Product_Item_Category `json:"categories,omitempty"`

	// PackageReferences - no documentation
	PackageReferences []*SoftLayer_Product_Package_Item_Prices `json:"packageReferences,omitempty"`

	// AccountRestrictions - no documentation
	AccountRestrictions []*SoftLayer_Product_Item_Price_Account_Restriction `json:"accountRestrictions,omitempty"`

	// BigDataOsJournalDiskFlag - Whether the price is for Big Data OS/Journal disks only. (Deprecated)
	BigDataOsJournalDiskFlag bool `json:"bigDataOsJournalDiskFlag,omitempty"`

	// CapacityRestrictionMaximum - The maximum capacity value for which this price is suitable.
	CapacityRestrictionMaximum string `json:"capacityRestrictionMaximum,omitempty"`

	// DefinedSoftwareLicenseFlag - Whether this price defines a software license for its product item.
	DefinedSoftwareLicenseFlag bool `json:"definedSoftwareLicenseFlag,omitempty"`

	// Inventory - no documentation
	Inventory []*SoftLayer_Product_Package_Inventory `json:"inventory,omitempty"`

	// PricingLocationGroup - The pricing location group that this price is applicable for. Prices that
	// have a pricing location group will only be available for ordering with the locations specified on
	// the location group.
	PricingLocationGroup *SoftLayer_Location_Group_Pricing `json:"pricingLocationGroup,omitempty"`

	// BundleReferences - no documentation
	BundleReferences []*SoftLayer_Product_Item_Bundles `json:"bundleReferences,omitempty"`

	// CategoryCount - A count of all categories which this item is a member.
	CategoryCount uint64 `json:"categoryCount,omitempty"`

	// PackageCount - A count of a price's packages under which this item is sold.
	PackageCount uint64 `json:"packageCount,omitempty"`

	// PackageReferenceCount - no documentation
	PackageReferenceCount uint64 `json:"packageReferenceCount,omitempty"`

	// Attributes - <nil>
	Attributes []*SoftLayer_Product_Item_Price_Attribute `json:"attributes,omitempty"`

	// Packages - no documentation
	Packages []*SoftLayer_Product_Package `json:"packages,omitempty"`

	// BundleReferenceCount - no documentation
	BundleReferenceCount uint64 `json:"bundleReferenceCount,omitempty"`

	// OrderPremiumCount - no documentation
	OrderPremiumCount uint64 `json:"orderPremiumCount,omitempty"`

	// OrderPremiums - <nil>
	OrderPremiums []*SoftLayer_Product_Item_Price_Premium `json:"orderPremiums,omitempty"`

	// RequiredCoreCount - The number of server cores required to order this item.
	RequiredCoreCount string `json:"requiredCoreCount,omitempty"`

	// AttributeCount - no documentation
	AttributeCount uint64 `json:"attributeCount,omitempty"`

	// CapacityRestrictionType - The type of capacity restriction by which this price must abide.
	CapacityRestrictionType string `json:"capacityRestrictionType,omitempty"`

	// PresetConfigurations - A list of preset configurations this price is used in.'
	PresetConfigurations []*SoftLayer_Product_Package_Preset_Configuration `json:"presetConfigurations,omitempty"`

	// CapacityRestrictionMinimum - The minimum capacity value for which this price is suitable.
	CapacityRestrictionMinimum string `json:"capacityRestrictionMinimum,omitempty"`

	// Item - no documentation
	Item *SoftLayer_Product_Item `json:"item,omitempty"`

	// AccountRestrictionCount - A count of the account that the item price is restricted to.
	AccountRestrictionCount uint64 `json:"accountRestrictionCount,omitempty"`

	// InventoryCount - A count of an item price's inventory status per datacenter.
	InventoryCount uint64 `json:"inventoryCount,omitempty"`

	// PresetConfigurationCount - A count of a list of preset configurations this price is used in.'
	PresetConfigurationCount uint64 `json:"presetConfigurationCount,omitempty"`
}

func (softlayer_product_item_price *SoftLayer_Product_Item_Price_Extended) String() string {
	return "SoftLayer_Product_Item_Price"
}
