package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Package - The SoftLayer_Product_Package data type contains information about
// packages from which orders can be generated. Packages contain general information regarding what is
// in them, where they are currently sold, availability, and pricing.
type SoftLayer_Product_Package struct {

	// AccountRestrictedCategories - The results from this call are similar to
	// [[SoftLayer_Product_Package/getCategories|getCategories]], but these include account-restricted
	// prices. Not all accounts have restricted pricing.
	AccountRestrictedCategories []*SoftLayer_Product_Item_Category `json:"accountRestrictedCategories"`

	// AccountRestrictedCategoryCount - A count of the results from this call are similar to
	// [[SoftLayer_Product_Package/getCategories|getCategories]], but these include account-restricted
	// prices. Not all accounts have restricted pricing.
	AccountRestrictedCategoryCount uint64 `json:"accountRestrictedCategoryCount"`

	// AccountRestrictedPricesFlag - The flag to indicate if there are any restricted prices in a package
	// for the currently-active account.
	AccountRestrictedPricesFlag bool `json:"accountRestrictedPricesFlag"`

	// ActivePresetCount - A count of the available preset configurations for this package.
	ActivePresetCount uint64 `json:"activePresetCount"`

	// ActivePresets - The available preset configurations for this package.
	ActivePresets []*SoftLayer_Product_Package_Preset `json:"activePresets"`

	// ActiveRamItemCount - A count of a collection of valid RAM items available for purchase in this
	// package.
	ActiveRamItemCount uint64 `json:"activeRamItemCount"`

	// ActiveRamItems - A collection of valid RAM items available for purchase in this package.
	ActiveRamItems []*SoftLayer_Product_Item `json:"activeRamItems"`

	// ActiveServerItemCount - A count of a collection of valid server items available for purchase in this
	// package.
	ActiveServerItemCount uint64 `json:"activeServerItemCount"`

	// ActiveServerItems - A collection of valid server items available for purchase in this package.
	ActiveServerItems []*SoftLayer_Product_Item `json:"activeServerItems"`

	// ActiveSoftwareItemCount - A count of a collection of valid software items available for purchase in
	// this package.
	ActiveSoftwareItemCount uint64 `json:"activeSoftwareItemCount"`

	// ActiveSoftwareItems - A collection of valid software items available for purchase in this package.
	ActiveSoftwareItems []*SoftLayer_Product_Item `json:"activeSoftwareItems"`

	// ActiveUsagePriceCount - A count of a collection of [[SoftLayer_Product_Item_Price]] objects for
	// pay-as-you-go usage.
	ActiveUsagePriceCount uint64 `json:"activeUsagePriceCount"`

	// ActiveUsagePrices - A collection of [[SoftLayer_Product_Item_Price]] objects for pay-as-you-go
	// usage.
	ActiveUsagePrices []*SoftLayer_Product_Item_Price `json:"activeUsagePrices"`

	// AdditionalServiceFlag - This flag indicates that the package is an additional service.
	AdditionalServiceFlag bool `json:"additionalServiceFlag"`

	// AttributeCount - no documentation
	AttributeCount uint64 `json:"attributeCount"`

	// Attributes - <nil>
	Attributes []*SoftLayer_Product_Package_Attribute `json:"attributes"`

	// AvailableLocationCount - A count of a collection of valid locations for this package.
	AvailableLocationCount uint64 `json:"availableLocationCount"`

	// AvailableLocations - no documentation
	AvailableLocations []*SoftLayer_Product_Package_Locations `json:"availableLocations"`

	// AvailableStorageUnits - The maximum number of available disk storage units associated with the
	// servers in a package.
	AvailableStorageUnits uint `json:"availableStorageUnits"`

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

	// Configuration - The item categories associated with a package, including information detailing which
	// item categories are required as part of a SoftLayer product order.
	Configuration []*SoftLayer_Product_Package_Order_Configuration `json:"configuration"`

	// ConfigurationCount - A count of the item categories associated with a package, including information
	// detailing which item categories are required as part of a SoftLayer product order.
	ConfigurationCount uint64 `json:"configurationCount"`

	// DefaultRamItemCount - A count of a collection of valid RAM items available for purchase in this
	// package.
	DefaultRamItemCount uint64 `json:"defaultRamItemCount"`

	// DefaultRamItems - A collection of valid RAM items available for purchase in this package.
	DefaultRamItems []*SoftLayer_Product_Item `json:"defaultRamItems"`

	// DeploymentCount - A count of the package that represents a multi-server solution. (Deprecated)
	DeploymentCount uint64 `json:"deploymentCount"`

	// DeploymentNodeType - The node type for a package in a solution deployment.
	DeploymentNodeType string `json:"deploymentNodeType"`

	// DeploymentPackageCount - A count of the packages that are allowed in a multi-server solution.
	// (Deprecated)
	DeploymentPackageCount uint64 `json:"deploymentPackageCount"`

	// DeploymentPackages - The packages that are allowed in a multi-server solution. (Deprecated)
	DeploymentPackages []*SoftLayer_Product_Package `json:"deploymentPackages"`

	// DeploymentType - no documentation
	DeploymentType string `json:"deploymentType"`

	// Deployments - The package that represents a multi-server solution. (Deprecated)
	Deployments []*SoftLayer_Product_Package `json:"deployments"`

	// Description - A generic description of the processor type and count. This includes so you may want
	// to strip these tags if you plan to use it.
	Description string `json:"description"`

	// DisallowCustomDiskPartitions - This flag indicates the package does not allow custom disk
	// partitions.
	DisallowCustomDiskPartitions bool `json:"disallowCustomDiskPartitions"`

	// FirstOrderStep - The Softlayer order step is optionally step-based. This returns the first
	// SoftLayer_Product_Package_Order_Step in the step-based order process.
	FirstOrderStep *SoftLayer_Product_Package_Order_Step `json:"firstOrderStep"`

	// FirstOrderStepId - This is only needed for step-based order verification. We use this for the order
	// forms, but it is not required. This step is the first SoftLayer_Product_Package_Step for this
	// package. Use this for for filtering which item categories are returned as a part of
	// SoftLayer_Product_Package_Order_Configuration.
	FirstOrderStepId int `json:"firstOrderStepId"`

	// GatewayApplianceFlag - Whether the package is a specialized network gateway appliance package.
	GatewayApplianceFlag bool `json:"gatewayApplianceFlag"`

	// GpuFlag - This flag indicates that the package supports GPUs.
	GpuFlag bool `json:"gpuFlag"`

	// HourlyBillingAvailableFlag - Determines whether the package contains prices that can be ordered
	// hourly.
	HourlyBillingAvailableFlag bool `json:"hourlyBillingAvailableFlag"`

	// Id - A package's internal identifier. Everything regarding a SoftLayer_Product_Package is tied back
	// to this id.
	Id int `json:"id"`

	// IsActive - <nil>
	IsActive int `json:"isActive"`

	// ItemConflictCount - A count of the item-item conflicts associated with a package.
	ItemConflictCount uint64 `json:"itemConflictCount"`

	// ItemConflicts - no documentation
	ItemConflicts []*SoftLayer_Product_Item_Resource_Conflict `json:"itemConflicts"`

	// ItemCount - A count of a collection of valid items available for purchase in this package.
	ItemCount uint64 `json:"itemCount"`

	// ItemLocationConflictCount - A count of the item-location conflicts associated with a package.
	ItemLocationConflictCount uint64 `json:"itemLocationConflictCount"`

	// ItemLocationConflicts - The item-location conflicts associated with a package.
	ItemLocationConflicts []*SoftLayer_Product_Item_Resource_Conflict `json:"itemLocationConflicts"`

	// ItemPriceCount - A count of a collection of SoftLayer_Product_Item_Prices that are valid for this
	// package.
	ItemPriceCount uint64 `json:"itemPriceCount"`

	// ItemPriceReferenceCount - no documentation
	ItemPriceReferenceCount uint64 `json:"itemPriceReferenceCount"`

	// ItemPriceReferences - no documentation
	ItemPriceReferences []*SoftLayer_Product_Package_Item_Prices `json:"itemPriceReferences"`

	// ItemPrices - A collection of SoftLayer_Product_Item_Prices that are valid for this package.
	ItemPrices []*SoftLayer_Product_Item_Price `json:"itemPrices"`

	// Items - A collection of valid items available for purchase in this package.
	Items []*SoftLayer_Product_Item `json:"items"`

	// LocationCount - A count of a collection of valid locations for this package.
	LocationCount uint64 `json:"locationCount"`

	// Locations - no documentation
	Locations []*SoftLayer_Location `json:"locations"`

	// LowestServerPrice - The lowest server [[SoftLayer_Product_Item_Price]] related to this package.
	LowestServerPrice *SoftLayer_Product_Item_Price `json:"lowestServerPrice"`

	// MaximumPortSpeed - The maximum available network speed associated with the package.
	MaximumPortSpeed uint `json:"maximumPortSpeed"`

	// MinimumPortSpeed - The minimum available network speed associated with the package.
	MinimumPortSpeed uint `json:"minimumPortSpeed"`

	// MongoDbEngineeredFlag - This flag indicates that this is a MongoDB engineered package. (Deprecated)
	MongoDbEngineeredFlag bool `json:"mongoDbEngineeredFlag"`

	// Name - The description of the package. For server packages, this is usually a detailed description
	// of processor type and count.
	Name string `json:"name"`

	// OrderPremiumCount - A count of the premium price modifiers associated with the
	// [[SoftLayer_Product_Item_Price]] and [[SoftLayer_Location]] objects in a package.
	OrderPremiumCount uint64 `json:"orderPremiumCount"`

	// OrderPremiums - The premium price modifiers associated with the [[SoftLayer_Product_Item_Price]] and
	// [[SoftLayer_Location]] objects in a package.
	OrderPremiums []*SoftLayer_Product_Item_Price_Premium `json:"orderPremiums"`

	// PreconfiguredFlag - This flag indicates the package is pre-configured. (Deprecated)
	PreconfiguredFlag bool `json:"preconfiguredFlag"`

	// PresetConfigurationRequiredFlag - Whether the package requires the user to define a preset
	// configuration.
	PresetConfigurationRequiredFlag bool `json:"presetConfigurationRequiredFlag"`

	// PreventVlanSelectionFlag - Whether the package prevents the user from specifying a Vlan.
	PreventVlanSelectionFlag bool `json:"preventVlanSelectionFlag"`

	// PrivateHostedCloudPackageFlag - This flag indicates the package is for a private hosted cloud
	// deployment. (Deprecated)
	PrivateHostedCloudPackageFlag bool `json:"privateHostedCloudPackageFlag"`

	// PrivateHostedCloudPackageType - The server role of the private hosted cloud deployment. (Deprecated)
	PrivateHostedCloudPackageType string `json:"privateHostedCloudPackageType"`

	// PrivateNetworkOnlyFlag - Whether the package only has access to the private network.
	PrivateNetworkOnlyFlag bool `json:"privateNetworkOnlyFlag"`

	// QuantaStorPackageFlag - Whether the package is a specialized mass storage QuantaStor package.
	QuantaStorPackageFlag bool `json:"quantaStorPackageFlag"`

	// RaidDiskRestrictionFlag - This flag indicates the package does not allow different disks with
	RaidDiskRestrictionFlag bool `json:"raidDiskRestrictionFlag"`

	// RedundantPowerFlag - This flag determines if the package contains a redundant power supply product.
	RedundantPowerFlag bool `json:"redundantPowerFlag"`

	// RegionCount - A count of the regional locations that a package is available in.
	RegionCount uint64 `json:"regionCount"`

	// Regions - The regional locations that a package is available in.
	Regions []*SoftLayer_Location_Region `json:"regions"`

	// ResourceGroupTemplate - The resource group template that describes a multi-server solution.
	// (Deprecated)
	ResourceGroupTemplate *SoftLayer_Resource_Group_Template `json:"resourceGroupTemplate"`

	// SubDescription - This currently contains no information but is here for future use.
	SubDescription string `json:"subDescription"`

	// TopLevelItemCategoryCode - The top level category code for this service offering.
	TopLevelItemCategoryCode string `json:"topLevelItemCategoryCode"`

	// Type - The type of service offering. This property can be used to help filter packages.
	Type *SoftLayer_Product_Package_Type `json:"type"`

	// UnitSize - no documentation
	UnitSize int `json:"unitSize"`
}

// GetActiveItems - Return a list of Items in the package with their active prices.
func (softlayer_product_package *SoftLayer_Product_Package) GetActiveItems() ([]*SoftLayer_Product_Item, error) {
	var returnValue []*SoftLayer_Product_Item
	return returnValue, nil
}

// GetActivePackagesByAttribute - This method is deprecated and should not be used in production code.
// This method will return the [[SoftLayer_Product_Package]] objects from which you can order a bare
// metal server, virtual server, service (such as CDN or Object Storage) or other software filtered by
// an attribute type associated with the package. Once you have the package you want to order from, you
// may query one of various endpoints from that package to get specific information about its products
// and pricing. See [[SoftLayer_Product_Package/getCategories|getCategories]] or
// [[SoftLayer_Product_Package/getItems|getItems]] for more information.
func (softlayer_product_package *SoftLayer_Product_Package) GetActivePackagesByAttribute(attributeKeyName string) ([]*SoftLayer_Product_Package, error) {
	var returnValue []*SoftLayer_Product_Package
	return returnValue, nil
}

// GetActivePrivateHostedCloudPackages - This method pulls all the active private hosted cloud
// packages. This will give you a basic description of the packages that are currently active and from
// which you can order private hosted cloud configurations.
func (softlayer_product_package *SoftLayer_Product_Package) GetActivePrivateHostedCloudPackages() ([]*SoftLayer_Product_Package, error) {
	var returnValue []*SoftLayer_Product_Package
	return returnValue, nil
}

// GetActiveUsageRatePrices - This method returns a collection of active usage rate
// [[SoftLayer_Product_Item_Price]] objects for the current package and specified datacenter.
// Optionally you can retrieve the active usage rate prices for a particular
// [[SoftLayer_Product_Item_Category]] by specifying a category code as the first parameter. This
// information is useful so that you can see "pay as you go" rates (if any) for the current package,
// location and optionally category.
func (softlayer_product_package *SoftLayer_Product_Package) GetActiveUsageRatePrices(locationId int, categoryCode string) ([]*SoftLayer_Product_Item_Price, error) {
	var returnValue []*SoftLayer_Product_Item_Price
	return returnValue, nil
}

// GetAllObjects - This method pulls all the active packages. This will give you a basic description of
// the packages that are currently active
func (softlayer_product_package *SoftLayer_Product_Package) GetAllObjects() ([]*SoftLayer_Product_Package, error) {
	var returnValue []*SoftLayer_Product_Package
	return returnValue, nil
}

// GetAvailablePackagesForImageTemplate - <nil>
func (softlayer_product_package *SoftLayer_Product_Package) GetAvailablePackagesForImageTemplate(imageTemplate SoftLayer_Virtual_Guest_Block_Device_Template_Group) ([]*SoftLayer_Product_Package, error) {
	var returnValue []*SoftLayer_Product_Package
	return returnValue, nil
}

// GetCdnItems - <nil>
func (softlayer_product_package *SoftLayer_Product_Package) GetCdnItems() ([]*SoftLayer_Product_Item, error) {
	var returnValue []*SoftLayer_Product_Item
	return returnValue, nil
}

// GetCloudStorageItems - <nil>
func (softlayer_product_package *SoftLayer_Product_Package) GetCloudStorageItems(provider int) ([]*SoftLayer_Product_Item, error) {
	var returnValue []*SoftLayer_Product_Item
	return returnValue, nil
}

// GetItemAvailabilityTypes - Returns a collection of SoftLayer_Product_Item_Attribute_Type objects.
// These item attribute types specifically deal with when an item, SoftLayer_Product_Item, from the
// product catalog may no longer be available. The keynames for these attribute types start with where
// the '*' may represent any string. For example, signifies that the item is not available for new
// orders. There is a catch all attribute type, If an item has one of these availability attributes
// set, the value should be a valid date in indicating the date after which the item will no longer be
// available.
func (softlayer_product_package *SoftLayer_Product_Package) GetItemAvailabilityTypes() ([]*SoftLayer_Product_Item_Attribute_Type, error) {
	var returnValue []*SoftLayer_Product_Item_Attribute_Type
	return returnValue, nil
}

// GetItemPricesFromSoftwareDescriptions - Return a collection of SoftLayer_Item_Price objects from a
// collection of SoftLayer_Software_Description
func (softlayer_product_package *SoftLayer_Product_Package) GetItemPricesFromSoftwareDescriptions(softwareDescriptions []SoftLayer_Software_Description, includeTranslationsFlag bool, returnAllPricesFlag bool) ([]*SoftLayer_Product_Item_Price, error) {
	var returnValue []*SoftLayer_Product_Item_Price
	return returnValue, nil
}

// GetItemsFromImageTemplate - Return a collection of [[SoftLayer_Product_Item]] objects from a
// [[SoftLayer_Virtual_Guest_Block_Device_Template_Group]] object
func (softlayer_product_package *SoftLayer_Product_Package) GetItemsFromImageTemplate(imageTemplate SoftLayer_Virtual_Guest_Block_Device_Template_Group) ([]*SoftLayer_Product_Item, error) {
	var returnValue []*SoftLayer_Product_Item
	return returnValue, nil
}

// GetMessageQueueItems - <nil>
func (softlayer_product_package *SoftLayer_Product_Package) GetMessageQueueItems() ([]*SoftLayer_Product_Item, error) {
	var returnValue []*SoftLayer_Product_Item
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_product_package *SoftLayer_Product_Package) GetObject() (*SoftLayer_Product_Package, error) {
	var returnValue *SoftLayer_Product_Package
	return returnValue, nil
}

// GetObjectStorageDatacenters - This method will return a collection of
// [[SoftLayer_Container_Product_Order_Network_Storage_Hub_Datacenter]] objects which contain a
// datacenter location and all the associated active usage rate prices where object storage is
// available. This method is really only applicable to the object storage additional service package
// which has a [[SoftLayer_Product_Package_Type]] of This information is useful so that you can see the
// "pay as you go" rates per datacenter.
func (softlayer_product_package *SoftLayer_Product_Package) GetObjectStorageDatacenters() ([]*SoftLayer_Container_Product_Order_Network_Storage_Hub_Datacenter, error) {
	var returnValue []*SoftLayer_Container_Product_Order_Network_Storage_Hub_Datacenter
	return returnValue, nil
}

// GetStandardCategories - This call is similar to
// [[SoftLayer_Product_Package/getCategories|getCategories]], except that it does not include
// account-restricted pricing. Not all accounts have restricted pricing.
func (softlayer_product_package *SoftLayer_Product_Package) GetStandardCategories() ([]*SoftLayer_Product_Item_Category, error) {
	var returnValue []*SoftLayer_Product_Item_Category
	return returnValue, nil
}
