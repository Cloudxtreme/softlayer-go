package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Package_Server - The SoftLayer_Product_Package_Server data type contains
// summarized information for bare metal servers regarding pricing, processor stats, and feature sets.
type SoftLayer_Product_Package_Server struct {

	// ProcessorCache - no documentation
	ProcessorCache string `json:"processorCache"`

	// ProcessorModel - no documentation
	ProcessorModel string `json:"processorModel"`

	// ProcessorSpeed - no documentation
	ProcessorSpeed string `json:"processorSpeed"`

	// ProductName - no documentation
	ProductName string `json:"productName"`

	// StartingHourlyPrice - The hourly starting price for the server. This includes a sum of all the
	// minimum required items, including RAM and hard drives. Not all servers are available hourly.
	StartingHourlyPrice float64 `json:"startingHourlyPrice"`

	// UnitSize - no documentation
	UnitSize int `json:"unitSize"`

	// ProcessorBusSpeed - no documentation
	ProcessorBusSpeed string `json:"processorBusSpeed"`

	// TotalCoreCount - The total number of processor cores available for the server.
	TotalCoreCount int `json:"totalCoreCount"`

	// TxtTpmFlag - Flag to indicate if the server configuration supports
	TxtTpmFlag bool `json:"txtTpmFlag"`

	// DualPathNetworkFlag - Flag to indicate if the server configuration supports dual path network
	// routing.
	DualPathNetworkFlag bool `json:"dualPathNetworkFlag"`

	// MaximumDriveCount - The maximum number of hard drives the server can support.
	MaximumDriveCount int `json:"maximumDriveCount"`

	// StartingMonthlyPrice - The monthly starting price for the server. This includes a sum of all the
	// minimum required items, including RAM and hard drives.
	StartingMonthlyPrice float64 `json:"startingMonthlyPrice"`

	// CatalogId - The unique identifier of a [[SoftLayer_Product_Catalog]].
	CatalogId int `json:"catalogId"`

	// GpuFlag - no documentation
	GpuFlag bool `json:"gpuFlag"`

	// MaximumPortSpeed - The maximum available network speed for the server.
	MaximumPortSpeed float64 `json:"maximumPortSpeed"`

	// PresetId - The unique identifier of a [[SoftLayer_Product_Package_Preset]].
	PresetId int `json:"presetId"`

	// ProcessorCount - no documentation
	ProcessorCount int `json:"processorCount"`

	// HourlyBillingFlag - Flag to determine if a server is available for hourly billing.
	HourlyBillingFlag bool `json:"hourlyBillingFlag"`

	// MaximumRamCapacity - no documentation
	MaximumRamCapacity float64 `json:"maximumRamCapacity"`

	// PackageId - The unique identifier of a [[SoftLayer_Product_Package]].
	PackageId int `json:"packageId"`

	// ProcessorCores - no documentation
	ProcessorCores int `json:"processorCores"`

	// ProcessorManufacturer - no documentation
	ProcessorManufacturer string `json:"processorManufacturer"`

	// DefaultRamCapacity - The minimum amount of RAM the server is configured with.
	DefaultRamCapacity float64 `json:"defaultRamCapacity"`

	// ItemId - The unique identifier of a [[SoftLayer_Product_Item]].
	ItemId int `json:"itemId"`

	// MinimumPortSpeed - The minimum available network speed for the server.
	MinimumPortSpeed float64 `json:"minimumPortSpeed"`

	// PackageType - no documentation
	PackageType string `json:"packageType"`

	// ProcessorName - no documentation
	ProcessorName string `json:"processorName"`

	// PrivateNetworkOnlyFlag - Indicates whether or not the server can only be configured with a private
	// network.
	PrivateNetworkOnlyFlag bool `json:"privateNetworkOnlyFlag"`

	// RedundantPowerFlag - Indicates whether or not the server has the capability to support a redundant
	// power supply.
	RedundantPowerFlag bool `json:"redundantPowerFlag"`

	// Id - The unique identifier of a [[SoftLayer_Product_Package_Server]].
	Id int `json:"id"`

	// ItemPriceId - The unique identifier of a [[SoftLayer_Product_Item_Price]].
	ItemPriceId int `json:"itemPriceId"`

	// OutletFlag - Indicates whether or not the server is being sold as part of an outlet package.
	OutletFlag bool `json:"outletFlag"`
}

func (softlayer_product_package_server *SoftLayer_Product_Package_Server) String() string {
	return "SoftLayer_Product_Package_Server"
}

// SoftLayer_Product_Package_Server_Extended is SoftLayer_Product_Package_Server with all maskable types.
type SoftLayer_Product_Package_Server_Extended struct {
	SoftLayer_Product_Package_Server

	// Catalog - <nil>
	Catalog *SoftLayer_Product_Catalog `json:"catalog"`

	// ItemPrice - <nil>
	ItemPrice *SoftLayer_Product_Item_Price `json:"itemPrice"`

	// Preset - <nil>
	Preset *SoftLayer_Product_Package_Preset `json:"preset"`

	// Item - <nil>
	Item *SoftLayer_Product_Item `json:"item"`

	// Package - <nil>
	Package *SoftLayer_Product_Package `json:"package"`
}

func (softlayer_product_package_server *SoftLayer_Product_Package_Server_Extended) String() string {
	return "SoftLayer_Product_Package_Server"
}
