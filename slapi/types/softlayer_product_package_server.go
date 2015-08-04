package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Product_Package_Server - The SoftLayer_Product_Package_Server data type contains
// summarized information for bare metal servers regarding pricing, processor stats, and feature sets.
type SoftLayer_Product_Package_Server struct {

	// DefaultRamCapacity - The minimum amount of RAM the server is configured with.
	DefaultRamCapacity slapi.Float64 `json:"defaultRamCapacity,omitempty"`

	// DualPathNetworkFlag - Flag to indicate if the server configuration supports dual path network
	// routing.
	DualPathNetworkFlag bool `json:"dualPathNetworkFlag,omitempty"`

	// Id - The unique identifier of a [[SoftLayer_Product_Package_Server]].
	Id int `json:"id,omitempty"`

	// ItemPriceId - The unique identifier of a [[SoftLayer_Product_Item_Price]].
	ItemPriceId int `json:"itemPriceId,omitempty"`

	// ProductName - no documentation
	ProductName string `json:"productName,omitempty"`

	// StartingMonthlyPrice - The monthly starting price for the server. This includes a sum of all the
	// minimum required items, including RAM and hard drives.
	StartingMonthlyPrice slapi.Float64 `json:"startingMonthlyPrice,omitempty"`

	// StartingHourlyPrice - The hourly starting price for the server. This includes a sum of all the
	// minimum required items, including RAM and hard drives. Not all servers are available hourly.
	StartingHourlyPrice slapi.Float64 `json:"startingHourlyPrice,omitempty"`

	// MaximumPortSpeed - The maximum available network speed for the server.
	MaximumPortSpeed slapi.Float64 `json:"maximumPortSpeed,omitempty"`

	// PresetId - The unique identifier of a [[SoftLayer_Product_Package_Preset]].
	PresetId int `json:"presetId,omitempty"`

	// RedundantPowerFlag - Indicates whether or not the server has the capability to support a redundant
	// power supply.
	RedundantPowerFlag bool `json:"redundantPowerFlag,omitempty"`

	// CatalogId - The unique identifier of a [[SoftLayer_Product_Catalog]].
	CatalogId int `json:"catalogId,omitempty"`

	// HourlyBillingFlag - Flag to determine if a server is available for hourly billing.
	HourlyBillingFlag bool `json:"hourlyBillingFlag,omitempty"`

	// OutletFlag - Indicates whether or not the server is being sold as part of an outlet package.
	OutletFlag bool `json:"outletFlag,omitempty"`

	// GpuFlag - no documentation
	GpuFlag bool `json:"gpuFlag,omitempty"`

	// PrivateNetworkOnlyFlag - Indicates whether or not the server can only be configured with a private
	// network.
	PrivateNetworkOnlyFlag bool `json:"privateNetworkOnlyFlag,omitempty"`

	// ProcessorCores - no documentation
	ProcessorCores int `json:"processorCores,omitempty"`

	// UnitSize - no documentation
	UnitSize int `json:"unitSize,omitempty"`

	// ItemId - The unique identifier of a [[SoftLayer_Product_Item]].
	ItemId int `json:"itemId,omitempty"`

	// MaximumRamCapacity - no documentation
	MaximumRamCapacity slapi.Float64 `json:"maximumRamCapacity,omitempty"`

	// PackageType - no documentation
	PackageType string `json:"packageType,omitempty"`

	// ProcessorBusSpeed - no documentation
	ProcessorBusSpeed string `json:"processorBusSpeed,omitempty"`

	// ProcessorCount - no documentation
	ProcessorCount int `json:"processorCount,omitempty"`

	// ProcessorCache - no documentation
	ProcessorCache string `json:"processorCache,omitempty"`

	// ProcessorModel - no documentation
	ProcessorModel string `json:"processorModel,omitempty"`

	// ProcessorName - no documentation
	ProcessorName string `json:"processorName,omitempty"`

	// ProcessorSpeed - no documentation
	ProcessorSpeed string `json:"processorSpeed,omitempty"`

	// TotalCoreCount - The total number of processor cores available for the server.
	TotalCoreCount int `json:"totalCoreCount,omitempty"`

	// TxtTpmFlag - Flag to indicate if the server configuration supports
	TxtTpmFlag bool `json:"txtTpmFlag,omitempty"`

	// MaximumDriveCount - The maximum number of hard drives the server can support.
	MaximumDriveCount int `json:"maximumDriveCount,omitempty"`

	// MinimumPortSpeed - The minimum available network speed for the server.
	MinimumPortSpeed slapi.Float64 `json:"minimumPortSpeed,omitempty"`

	// PackageId - The unique identifier of a [[SoftLayer_Product_Package]].
	PackageId int `json:"packageId,omitempty"`

	// ProcessorManufacturer - no documentation
	ProcessorManufacturer string `json:"processorManufacturer,omitempty"`
}

func (softlayer_product_package_server *SoftLayer_Product_Package_Server) String() string {
	return "SoftLayer_Product_Package_Server"
}

// SoftLayer_Product_Package_Server_Extended is SoftLayer_Product_Package_Server with all maskable types.
type SoftLayer_Product_Package_Server_Extended struct {
	SoftLayer_Product_Package_Server

	// ItemPrice - <nil>
	ItemPrice *SoftLayer_Product_Item_Price `json:"itemPrice,omitempty"`

	// Preset - <nil>
	Preset *SoftLayer_Product_Package_Preset `json:"preset,omitempty"`

	// Catalog - <nil>
	Catalog *SoftLayer_Product_Catalog `json:"catalog,omitempty"`

	// Item - <nil>
	Item *SoftLayer_Product_Item `json:"item,omitempty"`

	// Package - <nil>
	Package *SoftLayer_Product_Package `json:"package,omitempty"`
}

func (softlayer_product_package_server *SoftLayer_Product_Package_Server_Extended) String() string {
	return "SoftLayer_Product_Package_Server"
}
