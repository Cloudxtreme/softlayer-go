package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Product_Package_Server - The SoftLayer_Product_Package_Server data type contains
// summarized information for bare metal servers regarding pricing, processor stats, and feature sets.
type SoftLayer_Product_Package_Server struct {

	// Catalog - <nil>
	Catalog *SoftLayer_Product_Catalog `json:"catalog"`

	// CatalogId - The unique identifier of a [[SoftLayer_Product_Catalog]].
	CatalogId int `json:"catalogId"`

	// DefaultRamCapacity - The minimum amount of RAM the server is configured with.
	DefaultRamCapacity float64 `json:"defaultRamCapacity"`

	// DualPathNetworkFlag - Flag to indicate if the server configuration supports dual path network
	// routing.
	DualPathNetworkFlag bool `json:"dualPathNetworkFlag"`

	// GpuFlag - no documentation
	GpuFlag bool `json:"gpuFlag"`

	// HourlyBillingFlag - Flag to determine if a server is available for hourly billing.
	HourlyBillingFlag bool `json:"hourlyBillingFlag"`

	// Id - The unique identifier of a [[SoftLayer_Product_Package_Server]].
	Id int `json:"id"`

	// Item - <nil>
	Item *SoftLayer_Product_Item `json:"item"`

	// ItemId - The unique identifier of a [[SoftLayer_Product_Item]].
	ItemId int `json:"itemId"`

	// ItemPrice - <nil>
	ItemPrice *SoftLayer_Product_Item_Price `json:"itemPrice"`

	// ItemPriceId - The unique identifier of a [[SoftLayer_Product_Item_Price]].
	ItemPriceId int `json:"itemPriceId"`

	// MaximumDriveCount - The maximum number of hard drives the server can support.
	MaximumDriveCount int `json:"maximumDriveCount"`

	// MaximumPortSpeed - The maximum available network speed for the server.
	MaximumPortSpeed float64 `json:"maximumPortSpeed"`

	// MaximumRamCapacity - no documentation
	MaximumRamCapacity float64 `json:"maximumRamCapacity"`

	// MinimumPortSpeed - The minimum available network speed for the server.
	MinimumPortSpeed float64 `json:"minimumPortSpeed"`

	// OutletFlag - Indicates whether or not the server is being sold as part of an outlet package.
	OutletFlag bool `json:"outletFlag"`

	// Package - <nil>
	Package *SoftLayer_Product_Package `json:"package"`

	// PackageId - The unique identifier of a [[SoftLayer_Product_Package]].
	PackageId int `json:"packageId"`

	// PackageType - no documentation
	PackageType string `json:"packageType"`

	// Preset - <nil>
	Preset *SoftLayer_Product_Package_Preset `json:"preset"`

	// PresetId - The unique identifier of a [[SoftLayer_Product_Package_Preset]].
	PresetId int `json:"presetId"`

	// PrivateNetworkOnlyFlag - Indicates whether or not the server can only be configured with a private
	// network.
	PrivateNetworkOnlyFlag bool `json:"privateNetworkOnlyFlag"`

	// ProcessorBusSpeed - no documentation
	ProcessorBusSpeed string `json:"processorBusSpeed"`

	// ProcessorCache - no documentation
	ProcessorCache string `json:"processorCache"`

	// ProcessorCores - no documentation
	ProcessorCores int `json:"processorCores"`

	// ProcessorCount - no documentation
	ProcessorCount int `json:"processorCount"`

	// ProcessorManufacturer - no documentation
	ProcessorManufacturer string `json:"processorManufacturer"`

	// ProcessorModel - no documentation
	ProcessorModel string `json:"processorModel"`

	// ProcessorName - no documentation
	ProcessorName string `json:"processorName"`

	// ProcessorSpeed - no documentation
	ProcessorSpeed string `json:"processorSpeed"`

	// ProductName - no documentation
	ProductName string `json:"productName"`

	// RedundantPowerFlag - Indicates whether or not the server has the capability to support a redundant
	// power supply.
	RedundantPowerFlag bool `json:"redundantPowerFlag"`

	// StartingHourlyPrice - The hourly starting price for the server. This includes a sum of all the
	// minimum required items, including RAM and hard drives. Not all servers are available hourly.
	StartingHourlyPrice float64 `json:"startingHourlyPrice"`

	// StartingMonthlyPrice - The monthly starting price for the server. This includes a sum of all the
	// minimum required items, including RAM and hard drives.
	StartingMonthlyPrice float64 `json:"startingMonthlyPrice"`

	// TotalCoreCount - The total number of processor cores available for the server.
	TotalCoreCount int `json:"totalCoreCount"`

	// TxtTpmFlag - Flag to indicate if the server configuration supports
	TxtTpmFlag bool `json:"txtTpmFlag"`

	// UnitSize - no documentation
	UnitSize int `json:"unitSize"`
}

func (softlayer_product_package_server *SoftLayer_Product_Package_Server) String() string {
	return "SoftLayer_Product_Package_Server"
}

// GetAllObjects - no documentation
func (softlayer_product_package_server *SoftLayer_Product_Package_Server) GetAllObjects(ctx *slapi.RequestContext) ([]*SoftLayer_Product_Package_Server, error) {
	var returnValue []*SoftLayer_Product_Package_Server
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_product_package_server *SoftLayer_Product_Package_Server) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Product_Package_Server, error) {
	var returnValue *SoftLayer_Product_Package_Server
	return returnValue, nil
}
