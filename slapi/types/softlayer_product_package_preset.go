package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Package_Preset - Package presets are used to simplify ordering by eliminating the
// need for price ids when submitting orders. Orders submitted without prices or a preset id defined
// will use the preset for the package id. The default package presets include the base options
// required for a package configuration. Orders submitted with a preset id defined will use the prices
// included in the package preset. Prices submitted on an order with a preset id will replace the
// prices included in the package preset for that prices category. If the package preset has a fixed
// configuration flag (fixedConfigurationFlag) set then the prices included in the preset configuration
// cannot be replaced by prices submitted on the order. The only exception to the fixed configuration
// flag would be if a price submitted on the order is an account-restricted price for the same product
// item.
type SoftLayer_Product_Package_Preset struct {

	// AvailableStorageUnits - <nil>
	AvailableStorageUnits uint `json:"availableStorageUnits"`

	// Categories - The item categories that are included in this package preset configuration.
	Categories []*SoftLayer_Product_Item_Category `json:"categories"`

	// CategoryCount - A count of the item categories that are included in this package preset
	// configuration.
	CategoryCount uint64 `json:"categoryCount"`

	// Configuration - no documentation
	Configuration []*SoftLayer_Product_Package_Preset_Configuration `json:"configuration"`

	// ConfigurationCount - A count of the preset configuration (category and price).
	ConfigurationCount uint64 `json:"configurationCount"`

	// Description - no documentation
	Description string `json:"description"`

	// FixedConfigurationFlag - A package preset with this flag set will not allow the price's defined in
	// the preset configuration to be overriden during order placement.
	FixedConfigurationFlag bool `json:"fixedConfigurationFlag"`

	// Id - A preset's internal identifier. Everything regarding a SoftLayer_Product_Package_Preset is tied
	// back to this id.
	Id int `json:"id"`

	// IsActive - no documentation
	IsActive string `json:"isActive"`

	// KeyName - The key name of the package preset. For the base configuration of a package the preset key
	// name is
	KeyName string `json:"keyName"`

	// LowestPresetServerPrice - The lowest server prices related to this package preset.
	LowestPresetServerPrice *SoftLayer_Product_Item_Price `json:"lowestPresetServerPrice"`

	// Name - no documentation
	Name string `json:"name"`

	// Package - no documentation
	Package *SoftLayer_Product_Package `json:"package"`

	// PackageConfiguration - The item categories associated with a package preset, including information
	// detailing which item categories are required as part of a SoftLayer product order.
	PackageConfiguration []*SoftLayer_Product_Package_Order_Configuration `json:"packageConfiguration"`

	// PackageConfigurationCount - A count of the item categories associated with a package preset,
	// including information detailing which item categories are required as part of a SoftLayer product
	// order.
	PackageConfigurationCount uint64 `json:"packageConfigurationCount"`

	// PackageId - The package id for the package this preset belongs to.
	PackageId int `json:"packageId"`

	// PriceCount - A count of the item prices that are included in this package preset configuration.
	PriceCount uint64 `json:"priceCount"`

	// Prices - The item prices that are included in this package preset configuration.
	Prices []*SoftLayer_Product_Item_Price `json:"prices"`

	// StorageGroupTemplateArrayCount - A count of describes how all disks in this preset will be
	// configured.
	StorageGroupTemplateArrayCount uint64 `json:"storageGroupTemplateArrayCount"`

	// StorageGroupTemplateArrays - Describes how all disks in this preset will be configured.
	StorageGroupTemplateArrays []*SoftLayer_Configuration_Storage_Group_Template_Group `json:"storageGroupTemplateArrays"`

	// TotalMinimumHourlyFee - The starting hourly price for this configuration. Additional options not
	// defined in the preset may increase the cost.
	TotalMinimumHourlyFee uint `json:"totalMinimumHourlyFee"`

	// TotalMinimumRecurringFee - The starting monthly price for this configuration. Additional options not
	// defined in the preset may increase the cost.
	TotalMinimumRecurringFee uint `json:"totalMinimumRecurringFee"`
}

func (softlayer_product_package_preset *SoftLayer_Product_Package_Preset) String() string {
	return "SoftLayer_Product_Package_Preset"
}
