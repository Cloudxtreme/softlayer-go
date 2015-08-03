package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Package - The SoftLayer_Product_Package data type contains information about
// packages from which orders can be generated. Packages contain general information regarding what is
// in them, where they are currently sold, availability, and pricing.
type SoftLayer_Product_Package struct {

	// Description - A generic description of the processor type and count. This includes so you may want
	// to strip these tags if you plan to use it.
	Description string `json:"description"`

	// SubDescription - This currently contains no information but is here for future use.
	SubDescription string `json:"subDescription"`

	// FirstOrderStepId - This is only needed for step-based order verification. We use this for the order
	// forms, but it is not required. This step is the first SoftLayer_Product_Package_Step for this
	// package. Use this for for filtering which item categories are returned as a part of
	// SoftLayer_Product_Package_Order_Configuration.
	FirstOrderStepId int `json:"firstOrderStepId"`

	// Id - A package's internal identifier. Everything regarding a SoftLayer_Product_Package is tied back
	// to this id.
	Id int `json:"id"`

	// UnitSize - no documentation
	UnitSize int `json:"unitSize"`

	// Name - The description of the package. For server packages, this is usually a detailed description
	// of processor type and count.
	Name string `json:"name"`

	// IsActive - <nil>
	IsActive int `json:"isActive"`
}

// SoftLayer_Product_Package_Extended is SoftLayer_Product_Package with all maskable types.
type SoftLayer_Product_Package_Extended struct {
	SoftLayer_Product_Package

	// ActivePresets - The available preset configurations for this package.
	ActivePresets []*SoftLayer_Product_Package_Preset `json:"activePresets"`

	// DeploymentPackages - The packages that are allowed in a multi-server solution. (Deprecated)
	DeploymentPackages []*SoftLayer_Product_Package `json:"deploymentPackages"`

	// FirstOrderStep - The Softlayer order step is optionally step-based. This returns the first
	// SoftLayer_Product_Package_Order_Step in the step-based order process.
	FirstOrderStep *SoftLayer_Product_Package_Order_Step `json:"firstOrderStep"`

	// MinimumPortSpeed - The minimum available network speed associated with the package.
	MinimumPortSpeed uint `json:"minimumPortSpeed"`

	// PrivateHostedCloudPackageFlag - This flag indicates the package is for a private hosted cloud
	// deployment. (Deprecated)
	PrivateHostedCloudPackageFlag bool `json:"privateHostedCloudPackageFlag"`

	// TopLevelItemCategoryCode - The top level category code for this service offering.
	TopLevelItemCategoryCode string `json:"topLevelItemCategoryCode"`

	// OrderPremiumCount - A count of the premium price modifiers associated with the
	// [[SoftLayer_Product_Item_Price]] and [[SoftLayer_Location]] objects in a package.
	OrderPremiumCount uint64 `json:"orderPremiumCount"`

	// Attributes - <nil>
	Attributes []*SoftLayer_Product_Package_Attribute `json:"attributes"`

	// AvailableStorageUnits - The maximum number of available disk storage units associated with the
	// servers in a package.
	AvailableStorageUnits uint `json:"availableStorageUnits"`

	// Locations - no documentation
	Locations []*SoftLayer_Location `json:"locations"`

	// ItemPriceReferenceCount - no documentation
	ItemPriceReferenceCount uint64 `json:"itemPriceReferenceCount"`

	// DefaultRamItems - A collection of valid RAM items available for purchase in this package.
	DefaultRamItems []*SoftLayer_Product_Item `json:"defaultRamItems"`

	// GpuFlag - This flag indicates that the package supports GPUs.
	GpuFlag bool `json:"gpuFlag"`

	// LowestServerPrice - The lowest server [[SoftLayer_Product_Item_Price]] related to this package.
	LowestServerPrice *SoftLayer_Product_Item_Price `json:"lowestServerPrice"`

	// DefaultRamItemCount - A count of a collection of valid RAM items available for purchase in this
	// package.
	DefaultRamItemCount uint64 `json:"defaultRamItemCount"`

	// DeploymentPackageCount - A count of the packages that are allowed in a multi-server solution.
	// (Deprecated)
	DeploymentPackageCount uint64 `json:"deploymentPackageCount"`

	// ItemConflictCount - A count of the item-item conflicts associated with a package.
	ItemConflictCount uint64 `json:"itemConflictCount"`

	// ActiveSoftwareItems - A collection of valid software items available for purchase in this package.
	ActiveSoftwareItems []*SoftLayer_Product_Item `json:"activeSoftwareItems"`

	// ItemConflicts - no documentation
	ItemConflicts []*SoftLayer_Product_Item_Resource_Conflict `json:"itemConflicts"`

	// ActiveRamItemCount - A count of a collection of valid RAM items available for purchase in this
	// package.
	ActiveRamItemCount uint64 `json:"activeRamItemCount"`

	// AdditionalServiceFlag - This flag indicates that the package is an additional service.
	AdditionalServiceFlag bool `json:"additionalServiceFlag"`

	// DisallowCustomDiskPartitions - This flag indicates the package does not allow custom disk
	// partitions.
	DisallowCustomDiskPartitions bool `json:"disallowCustomDiskPartitions"`

	// RedundantPowerFlag - This flag determines if the package contains a redundant power supply product.
	RedundantPowerFlag bool `json:"redundantPowerFlag"`

	// ConfigurationCount - A count of the item categories associated with a package, including information
	// detailing which item categories are required as part of a SoftLayer product order.
	ConfigurationCount uint64 `json:"configurationCount"`

	// DeploymentCount - A count of the package that represents a multi-server solution. (Deprecated)
	DeploymentCount uint64 `json:"deploymentCount"`

	// Items - A collection of valid items available for purchase in this package.
	Items []*SoftLayer_Product_Item `json:"items"`

	// ActiveUsagePriceCount - A count of a collection of [[SoftLayer_Product_Item_Price]] objects for
	// pay-as-you-go usage.
	ActiveUsagePriceCount uint64 `json:"activeUsagePriceCount"`

	// AvailableLocationCount - A count of a collection of valid locations for this package.
	AvailableLocationCount uint64 `json:"availableLocationCount"`

	// PreventVlanSelectionFlag - Whether the package prevents the user from specifying a Vlan.
	PreventVlanSelectionFlag bool `json:"preventVlanSelectionFlag"`

	// Categories - This is a collection of categories ([[SoftLayer_Product_Item_Category]]) associated
	// with a package which can be used for ordering. These categories have several objects prepopulated
	// which are useful when determining the available products for purchase. The categories contain groups
	// ([[SoftLayer_Product_Package_Item_Category_Group]]) that organize the products and prices by similar
	// features. For example, operating systems will be grouped by their manufacturer and virtual server
	// disks will be grouped by their disk type vs. local). Each group will contain prices
	// ([[SoftLayer_Product_Item_Price]]) which you can use determine the cost of each product. Each price
	// has a product ([[SoftLayer_Product_Item]]) which provides the name and other useful information
	// about the server, service or software you may purchase.
	Categories []*SoftLayer_Product_Item_Category `json:"categories"`

	// ItemPrices - A collection of SoftLayer_Product_Item_Prices that are valid for this package.
	ItemPrices []*SoftLayer_Product_Item_Price `json:"itemPrices"`

	// MaximumPortSpeed - The maximum available network speed associated with the package.
	MaximumPortSpeed uint `json:"maximumPortSpeed"`

	// PresetConfigurationRequiredFlag - Whether the package requires the user to define a preset
	// configuration.
	PresetConfigurationRequiredFlag bool `json:"presetConfigurationRequiredFlag"`

	// RaidDiskRestrictionFlag - This flag indicates the package does not allow different disks with
	RaidDiskRestrictionFlag bool `json:"raidDiskRestrictionFlag"`

	// ResourceGroupTemplate - The resource group template that describes a multi-server solution.
	// (Deprecated)
	ResourceGroupTemplate *SoftLayer_Resource_Group_Template `json:"resourceGroupTemplate"`

	// ActiveSoftwareItemCount - A count of a collection of valid software items available for purchase in
	// this package.
	ActiveSoftwareItemCount uint64 `json:"activeSoftwareItemCount"`

	// DeploymentNodeType - The node type for a package in a solution deployment.
	DeploymentNodeType string `json:"deploymentNodeType"`

	// GatewayApplianceFlag - Whether the package is a specialized network gateway appliance package.
	GatewayApplianceFlag bool `json:"gatewayApplianceFlag"`

	// ItemPriceReferences - no documentation
	ItemPriceReferences []*SoftLayer_Product_Package_Item_Prices `json:"itemPriceReferences"`

	// LocationCount - A count of a collection of valid locations for this package.
	LocationCount uint64 `json:"locationCount"`

	// RegionCount - A count of the regional locations that a package is available in.
	RegionCount uint64 `json:"regionCount"`

	// MongoDbEngineeredFlag - This flag indicates that this is a MongoDB engineered package. (Deprecated)
	MongoDbEngineeredFlag bool `json:"mongoDbEngineeredFlag"`

	// QuantaStorPackageFlag - Whether the package is a specialized mass storage QuantaStor package.
	QuantaStorPackageFlag bool `json:"quantaStorPackageFlag"`

	// ActiveServerItemCount - A count of a collection of valid server items available for purchase in this
	// package.
	ActiveServerItemCount uint64 `json:"activeServerItemCount"`

	// AttributeCount - no documentation
	AttributeCount uint64 `json:"attributeCount"`

	// Configuration - The item categories associated with a package, including information detailing which
	// item categories are required as part of a SoftLayer product order.
	Configuration []*SoftLayer_Product_Package_Order_Configuration `json:"configuration"`

	// PreconfiguredFlag - This flag indicates the package is pre-configured. (Deprecated)
	PreconfiguredFlag bool `json:"preconfiguredFlag"`

	// Type - The type of service offering. This property can be used to help filter packages.
	Type *SoftLayer_Product_Package_Type `json:"type"`

	// ItemCount - A count of a collection of valid items available for purchase in this package.
	ItemCount uint64 `json:"itemCount"`

	// CategoryCount - A count of this is a collection of categories ([[SoftLayer_Product_Item_Category]])
	// associated with a package which can be used for ordering. These categories have several objects
	// prepopulated which are useful when determining the available products for purchase. The categories
	// contain groups ([[SoftLayer_Product_Package_Item_Category_Group]]) that organize the products and
	// prices by similar features. For example, operating systems will be grouped by their manufacturer and
	// virtual server disks will be grouped by their disk type vs. local). Each group will contain prices
	// ([[SoftLayer_Product_Item_Price]]) which you can use determine the cost of each product. Each price
	// has a product ([[SoftLayer_Product_Item]]) which provides the name and other useful information
	// about the server, service or software you may purchase.
	CategoryCount uint64 `json:"categoryCount"`

	// ActiveUsagePrices - A collection of [[SoftLayer_Product_Item_Price]] objects for pay-as-you-go
	// usage.
	ActiveUsagePrices []*SoftLayer_Product_Item_Price `json:"activeUsagePrices"`

	// PrivateNetworkOnlyFlag - Whether the package only has access to the private network.
	PrivateNetworkOnlyFlag bool `json:"privateNetworkOnlyFlag"`

	// Regions - The regional locations that a package is available in.
	Regions []*SoftLayer_Location_Region `json:"regions"`

	// AccountRestrictedCategories - The results from this call are similar to
	// [[SoftLayer_Product_Package/getCategories|getCategories]], but these include account-restricted
	// prices. Not all accounts have restricted pricing.
	AccountRestrictedCategories []*SoftLayer_Product_Item_Category `json:"accountRestrictedCategories"`

	// ItemLocationConflicts - The item-location conflicts associated with a package.
	ItemLocationConflicts []*SoftLayer_Product_Item_Resource_Conflict `json:"itemLocationConflicts"`

	// OrderPremiums - The premium price modifiers associated with the [[SoftLayer_Product_Item_Price]] and
	// [[SoftLayer_Location]] objects in a package.
	OrderPremiums []*SoftLayer_Product_Item_Price_Premium `json:"orderPremiums"`

	// ItemPriceCount - A count of a collection of SoftLayer_Product_Item_Prices that are valid for this
	// package.
	ItemPriceCount uint64 `json:"itemPriceCount"`

	// AccountRestrictedPricesFlag - The flag to indicate if there are any restricted prices in a package
	// for the currently-active account.
	AccountRestrictedPricesFlag bool `json:"accountRestrictedPricesFlag"`

	// ActiveServerItems - A collection of valid server items available for purchase in this package.
	ActiveServerItems []*SoftLayer_Product_Item `json:"activeServerItems"`

	// DeploymentType - no documentation
	DeploymentType string `json:"deploymentType"`

	// AccountRestrictedCategoryCount - A count of the results from this call are similar to
	// [[SoftLayer_Product_Package/getCategories|getCategories]], but these include account-restricted
	// prices. Not all accounts have restricted pricing.
	AccountRestrictedCategoryCount uint64 `json:"accountRestrictedCategoryCount"`

	// ActivePresetCount - A count of the available preset configurations for this package.
	ActivePresetCount uint64 `json:"activePresetCount"`

	// ItemLocationConflictCount - A count of the item-location conflicts associated with a package.
	ItemLocationConflictCount uint64 `json:"itemLocationConflictCount"`

	// HourlyBillingAvailableFlag - Determines whether the package contains prices that can be ordered
	// hourly.
	HourlyBillingAvailableFlag bool `json:"hourlyBillingAvailableFlag"`

	// PrivateHostedCloudPackageType - The server role of the private hosted cloud deployment. (Deprecated)
	PrivateHostedCloudPackageType string `json:"privateHostedCloudPackageType"`

	// ActiveRamItems - A collection of valid RAM items available for purchase in this package.
	ActiveRamItems []*SoftLayer_Product_Item `json:"activeRamItems"`

	// AvailableLocations - no documentation
	AvailableLocations []*SoftLayer_Product_Package_Locations `json:"availableLocations"`

	// Deployments - The package that represents a multi-server solution. (Deprecated)
	Deployments []*SoftLayer_Product_Package `json:"deployments"`
}

func (softlayer_product_package *SoftLayer_Product_Package) String() string {
	return "SoftLayer_Product_Package"
}
