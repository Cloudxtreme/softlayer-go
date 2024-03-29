package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Product_Package_Server - The SoftLayer_Product_Package_Server data type contains
// summarized information for bare metal servers regarding pricing, processor stats, and feature sets.
type SoftLayer_Product_Package_Server struct {

	// GpuFlag - no documentation
	GpuFlag bool `json:"gpuFlag,omitempty"`

	// ItemId - The unique identifier of a [[SoftLayer_Product_Item]].
	ItemId int `json:"itemId,omitempty"`

	// MaximumPortSpeed - The maximum available network speed for the server.
	MaximumPortSpeed slapi.Float64 `json:"maximumPortSpeed,omitempty"`

	// StartingMonthlyPrice - The monthly starting price for the server. This includes a sum of all the
	// minimum required items, including RAM and hard drives.
	StartingMonthlyPrice slapi.Float64 `json:"startingMonthlyPrice,omitempty"`

	// DualPathNetworkFlag - Flag to indicate if the server configuration supports dual path network
	// routing.
	DualPathNetworkFlag bool `json:"dualPathNetworkFlag,omitempty"`

	// HourlyBillingFlag - Flag to determine if a server is available for hourly billing.
	HourlyBillingFlag bool `json:"hourlyBillingFlag,omitempty"`

	// MaximumDriveCount - The maximum number of hard drives the server can support.
	MaximumDriveCount int `json:"maximumDriveCount,omitempty"`

	// MinimumPortSpeed - The minimum available network speed for the server.
	MinimumPortSpeed slapi.Float64 `json:"minimumPortSpeed,omitempty"`

	// PresetId - The unique identifier of a [[SoftLayer_Product_Package_Preset]].
	PresetId int `json:"presetId,omitempty"`

	// ProcessorCount - no documentation
	ProcessorCount int `json:"processorCount,omitempty"`

	// StartingHourlyPrice - The hourly starting price for the server. This includes a sum of all the
	// minimum required items, including RAM and hard drives. Not all servers are available hourly.
	StartingHourlyPrice slapi.Float64 `json:"startingHourlyPrice,omitempty"`

	// ProcessorName - no documentation
	ProcessorName string `json:"processorName,omitempty"`

	// ProductName - no documentation
	ProductName string `json:"productName,omitempty"`

	// RedundantPowerFlag - Indicates whether or not the server has the capability to support a redundant
	// power supply.
	RedundantPowerFlag bool `json:"redundantPowerFlag,omitempty"`

	// ProcessorSpeed - no documentation
	ProcessorSpeed string `json:"processorSpeed,omitempty"`

	// UnitSize - no documentation
	UnitSize int `json:"unitSize,omitempty"`

	// PackageType - no documentation
	PackageType string `json:"packageType,omitempty"`

	// DefaultRamCapacity - The minimum amount of RAM the server is configured with.
	DefaultRamCapacity slapi.Float64 `json:"defaultRamCapacity,omitempty"`

	// MaximumRamCapacity - no documentation
	MaximumRamCapacity slapi.Float64 `json:"maximumRamCapacity,omitempty"`

	// OutletFlag - Indicates whether or not the server is being sold as part of an outlet package.
	OutletFlag bool `json:"outletFlag,omitempty"`

	// ProcessorCores - no documentation
	ProcessorCores int `json:"processorCores,omitempty"`

	// ProcessorManufacturer - no documentation
	ProcessorManufacturer string `json:"processorManufacturer,omitempty"`

	// TotalCoreCount - The total number of processor cores available for the server.
	TotalCoreCount int `json:"totalCoreCount,omitempty"`

	// TxtTpmFlag - Flag to indicate if the server configuration supports
	TxtTpmFlag bool `json:"txtTpmFlag,omitempty"`

	// ProcessorCache - no documentation
	ProcessorCache string `json:"processorCache,omitempty"`

	// ProcessorBusSpeed - no documentation
	ProcessorBusSpeed string `json:"processorBusSpeed,omitempty"`

	// Id - The unique identifier of a [[SoftLayer_Product_Package_Server]].
	Id int `json:"id,omitempty"`

	// ItemPriceId - The unique identifier of a [[SoftLayer_Product_Item_Price]].
	ItemPriceId int `json:"itemPriceId,omitempty"`

	// PackageId - The unique identifier of a [[SoftLayer_Product_Package]].
	PackageId int `json:"packageId,omitempty"`

	// PrivateNetworkOnlyFlag - Indicates whether or not the server can only be configured with a private
	// network.
	PrivateNetworkOnlyFlag bool `json:"privateNetworkOnlyFlag,omitempty"`

	// ProcessorModel - no documentation
	ProcessorModel string `json:"processorModel,omitempty"`

	// CatalogId - The unique identifier of a [[SoftLayer_Product_Catalog]].
	CatalogId int `json:"catalogId,omitempty"`

	// Preset - <nil>
	Preset *SoftLayer_Product_Package_Preset `json:"preset,omitempty"`

	// ItemPrice - <nil>
	ItemPrice *SoftLayer_Product_Item_Price `json:"itemPrice,omitempty"`

	// Item - <nil>
	Item *SoftLayer_Product_Item `json:"item,omitempty"`

	// Package - <nil>
	Package *SoftLayer_Product_Package `json:"package,omitempty"`

	// Catalog - <nil>
	Catalog *SoftLayer_Product_Catalog `json:"catalog,omitempty"`
}

func (softlayer_product_package_server *SoftLayer_Product_Package_Server) String() string {
	return "SoftLayer_Product_Package_Server"
}
