package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Package_Server - The SoftLayer_Product_Package_Server data type contains
// summarized information for bare metal servers regarding pricing, processor stats, and feature sets.
type SoftLayer_Product_Package_Server struct {

	// PackageType - no documentation
	PackageType string `json:"packageType"`

	// PresetId - The unique identifier of a [[SoftLayer_Product_Package_Preset]].
	PresetId int `json:"presetId"`

	// ProcessorName - no documentation
	ProcessorName string `json:"processorName"`

	// DefaultRamCapacity - The minimum amount of RAM the server is configured with.
	DefaultRamCapacity float64 `json:"defaultRamCapacity"`

	// ProcessorCores - no documentation
	ProcessorCores int `json:"processorCores"`

	// ProcessorCount - no documentation
	ProcessorCount int `json:"processorCount"`

	// TxtTpmFlag - Flag to indicate if the server configuration supports
	TxtTpmFlag bool `json:"txtTpmFlag"`

	// UnitSize - no documentation
	UnitSize int `json:"unitSize"`

	// ProcessorManufacturer - no documentation
	ProcessorManufacturer string `json:"processorManufacturer"`

	// ProcessorSpeed - no documentation
	ProcessorSpeed string `json:"processorSpeed"`

	// TotalCoreCount - The total number of processor cores available for the server.
	TotalCoreCount int `json:"totalCoreCount"`

	// ProcessorBusSpeed - no documentation
	ProcessorBusSpeed string `json:"processorBusSpeed"`

	// ItemPriceId - The unique identifier of a [[SoftLayer_Product_Item_Price]].
	ItemPriceId int `json:"itemPriceId"`

	// MaximumPortSpeed - The maximum available network speed for the server.
	MaximumPortSpeed float64 `json:"maximumPortSpeed"`

	// ProcessorCache - no documentation
	ProcessorCache string `json:"processorCache"`

	// ProductName - no documentation
	ProductName string `json:"productName"`

	// HourlyBillingFlag - Flag to determine if a server is available for hourly billing.
	HourlyBillingFlag bool `json:"hourlyBillingFlag"`

	// MaximumRamCapacity - no documentation
	MaximumRamCapacity float64 `json:"maximumRamCapacity"`

	// ProcessorModel - no documentation
	ProcessorModel string `json:"processorModel"`

	// RedundantPowerFlag - Indicates whether or not the server has the capability to support a redundant
	// power supply.
	RedundantPowerFlag bool `json:"redundantPowerFlag"`

	// StartingMonthlyPrice - The monthly starting price for the server. This includes a sum of all the
	// minimum required items, including RAM and hard drives.
	StartingMonthlyPrice float64 `json:"startingMonthlyPrice"`

	// CatalogId - The unique identifier of a [[SoftLayer_Product_Catalog]].
	CatalogId int `json:"catalogId"`

	// Id - The unique identifier of a [[SoftLayer_Product_Package_Server]].
	Id int `json:"id"`

	// DualPathNetworkFlag - Flag to indicate if the server configuration supports dual path network
	// routing.
	DualPathNetworkFlag bool `json:"dualPathNetworkFlag"`

	// PackageId - The unique identifier of a [[SoftLayer_Product_Package]].
	PackageId int `json:"packageId"`

	// PrivateNetworkOnlyFlag - Indicates whether or not the server can only be configured with a private
	// network.
	PrivateNetworkOnlyFlag bool `json:"privateNetworkOnlyFlag"`

	// ItemId - The unique identifier of a [[SoftLayer_Product_Item]].
	ItemId int `json:"itemId"`

	// MaximumDriveCount - The maximum number of hard drives the server can support.
	MaximumDriveCount int `json:"maximumDriveCount"`

	// MinimumPortSpeed - The minimum available network speed for the server.
	MinimumPortSpeed float64 `json:"minimumPortSpeed"`

	// OutletFlag - Indicates whether or not the server is being sold as part of an outlet package.
	OutletFlag bool `json:"outletFlag"`

	// StartingHourlyPrice - The hourly starting price for the server. This includes a sum of all the
	// minimum required items, including RAM and hard drives. Not all servers are available hourly.
	StartingHourlyPrice float64 `json:"startingHourlyPrice"`

	// GpuFlag - no documentation
	GpuFlag bool `json:"gpuFlag"`
}

// SoftLayer_Product_Package_Server_Extended is SoftLayer_Product_Package_Server with all maskable types.
type SoftLayer_Product_Package_Server_Extended struct {
	SoftLayer_Product_Package_Server

	// Catalog - <nil>
	Catalog *SoftLayer_Product_Catalog `json:"catalog"`

	// Package - <nil>
	Package *SoftLayer_Product_Package `json:"package"`

	// Preset - <nil>
	Preset *SoftLayer_Product_Package_Preset `json:"preset"`

	// Item - <nil>
	Item *SoftLayer_Product_Item `json:"item"`

	// ItemPrice - <nil>
	ItemPrice *SoftLayer_Product_Item_Price `json:"itemPrice"`
}

func (softlayer_product_package_server *SoftLayer_Product_Package_Server) String() string {
	return "SoftLayer_Product_Package_Server"
}
