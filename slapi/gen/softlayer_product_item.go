package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Item - The SoftLayer_Product_Item data type contains general information relating
// to a single SoftLayer product.
type SoftLayer_Product_Item struct {

	// ActivePresaleEventCount - no documentation
	ActivePresaleEventCount uint64 `json:"activePresaleEventCount"`

	// ActivePresaleEvents - <nil>
	ActivePresaleEvents []*SoftLayer_Sales_Presale_Event `json:"activePresaleEvents"`

	// ActiveUsagePriceCount - no documentation
	ActiveUsagePriceCount uint64 `json:"activeUsagePriceCount"`

	// ActiveUsagePrices - no documentation
	ActiveUsagePrices []*SoftLayer_Product_Item_Price `json:"activeUsagePrices"`

	// AttributeCount - A count of the attribute values for a product item. These are additional properties
	// that give extra information about the product being sold.
	AttributeCount uint64 `json:"attributeCount"`

	// Attributes - The attribute values for a product item. These are additional properties that give
	// extra information about the product being sold.
	Attributes []*SoftLayer_Product_Item_Attribute `json:"attributes"`

	// AvailabilityAttributeCount - A count of attributes that govern when an item may no longer be
	// available.
	AvailabilityAttributeCount uint64 `json:"availabilityAttributeCount"`

	// AvailabilityAttributes - Attributes that govern when an item may no longer be available.
	AvailabilityAttributes []*SoftLayer_Product_Item_Attribute `json:"availabilityAttributes"`

	// BillingType - no documentation
	BillingType string `json:"billingType"`

	// Bundle - An item's included products. Some items have other items included in them that we
	// specifically detail. They are here called Bundled Items. An example is Plesk unlimited. It as a
	// bundled item labeled 'SiteBuilder'. These are the SoftLayer_Product_Item_Bundles objects.
	Bundle []*SoftLayer_Product_Item_Bundles `json:"bundle"`

	// BundleCount - A count of an item's included products. Some items have other items included in them
	// that we specifically detail. They are here called Bundled Items. An example is Plesk unlimited. It
	// as a bundled item labeled 'SiteBuilder'. These are the SoftLayer_Product_Item_Bundles objects.
	BundleCount uint64 `json:"bundleCount"`

	// Capacity - Some Product Items have capacity information such as RAM and bandwidth, and others. This
	// provides the numerical representation of the capacity given in the description of this product item.
	Capacity float64 `json:"capacity"`

	// Categories - no documentation
	Categories []*SoftLayer_Product_Item_Category `json:"categories"`

	// CategoryCount - no documentation
	CategoryCount uint64 `json:"categoryCount"`

	// ConfigurationTemplateCount - A count of some product items have configuration templates which can be
	// used to during provisioning of that product.
	ConfigurationTemplateCount uint64 `json:"configurationTemplateCount"`

	// ConfigurationTemplates - Some product items have configuration templates which can be used to during
	// provisioning of that product.
	ConfigurationTemplates []*SoftLayer_Configuration_Template `json:"configurationTemplates"`

	// ConflictCount - A count of an item's conflicts. For example, McAfee LinuxShield cannot be ordered
	// with Windows. It was not meant for that operating system and as such is a conflict.
	ConflictCount uint64 `json:"conflictCount"`

	// Conflicts - An item's conflicts. For example, McAfee LinuxShield cannot be ordered with Windows. It
	// was not meant for that operating system and as such is a conflict.
	Conflicts []*SoftLayer_Product_Item_Resource_Conflict `json:"conflicts"`

	// CoreRestrictedItemFlag - This flag indicates that this product is restricted by the number of cores
	// on the compute instance.
	CoreRestrictedItemFlag bool `json:"coreRestrictedItemFlag"`

	// Description - no documentation
	Description string `json:"description"`

	// DowngradeItem - Some product items have a downgrade path. This is the first product item in the
	// downgrade path.
	DowngradeItem *SoftLayer_Product_Item `json:"downgradeItem"`

	// DowngradeItemCount - A count of some product items have a downgrade path. These are those product
	// items.
	DowngradeItemCount uint64 `json:"downgradeItemCount"`

	// DowngradeItems - Some product items have a downgrade path. These are those product items.
	DowngradeItems []*SoftLayer_Product_Item `json:"downgradeItems"`

	// GlobalCategoryConflictCount - A count of an item's category conflicts. For example, 10 Gbps
	// redundant network functionality cannot be ordered with a secondary GPU and as such is a conflict.
	GlobalCategoryConflictCount uint64 `json:"globalCategoryConflictCount"`

	// GlobalCategoryConflicts - An item's category conflicts. For example, 10 Gbps redundant network
	// functionality cannot be ordered with a secondary GPU and as such is a conflict.
	GlobalCategoryConflicts []*SoftLayer_Product_Item_Resource_Conflict `json:"globalCategoryConflicts"`

	// HardwareGenericComponentModel - The generic hardware component that this item represents.
	HardwareGenericComponentModel *SoftLayer_Hardware_Component_Model_Generic `json:"hardwareGenericComponentModel"`

	// HideFromPortalFlag - <nil>
	HideFromPortalFlag bool `json:"hideFromPortalFlag"`

	// Id - no documentation
	Id int `json:"id"`

	// Inventory - no documentation
	Inventory []*SoftLayer_Product_Package_Inventory `json:"inventory"`

	// InventoryCount - A count of an item's inventory status per datacenter.
	InventoryCount uint64 `json:"inventoryCount"`

	// IsEngineeredServerProduct - Flag to indicate the server product is engineered for a multi-server
	// solution. (Deprecated)
	IsEngineeredServerProduct bool `json:"isEngineeredServerProduct"`

	// ItemCategory - no documentation
	ItemCategory *SoftLayer_Product_Item_Category `json:"itemCategory"`

	// ItemTaxCategoryId - A products tax category internal identification number
	ItemTaxCategoryId int `json:"itemTaxCategoryId"`

	// KeyName - no documentation
	KeyName string `json:"keyName"`

	// LocationConflictCount - A count of an item's location conflicts. For example, Dual Path network
	// functionality cannot be ordered in WDC and as such is a conflict.
	LocationConflictCount uint64 `json:"locationConflictCount"`

	// LocationConflicts - An item's location conflicts. For example, Dual Path network functionality
	// cannot be ordered in WDC and as such is a conflict.
	LocationConflicts []*SoftLayer_Product_Item_Resource_Conflict `json:"locationConflicts"`

	// LongDescription - no documentation
	LongDescription string `json:"longDescription"`

	// ObjectStorageItemFlag - <nil>
	ObjectStorageItemFlag bool `json:"objectStorageItemFlag"`

	// PackageCount - A count of a collection of all the SoftLayer_Product_Package(s) in which this item
	// exists.
	PackageCount uint64 `json:"packageCount"`

	// Packages - A collection of all the SoftLayer_Product_Package(s) in which this item exists.
	Packages []*SoftLayer_Product_Package `json:"packages"`

	// PhysicalCoreCapacity - no documentation
	PhysicalCoreCapacity string `json:"physicalCoreCapacity"`

	// PresaleEventCount - no documentation
	PresaleEventCount uint64 `json:"presaleEventCount"`

	// PresaleEvents - <nil>
	PresaleEvents []*SoftLayer_Sales_Presale_Event `json:"presaleEvents"`

	// PriceCount - no documentation
	PriceCount uint64 `json:"priceCount"`

	// Prices - no documentation
	Prices []*SoftLayer_Product_Item_Price `json:"prices"`

	// RequirementCount - A count of if an item must be ordered with another item, it will have a
	// requirement item here.
	RequirementCount uint64 `json:"requirementCount"`

	// Requirements - If an item must be ordered with another item, it will have a requirement item here.
	Requirements []*SoftLayer_Product_Item_Requirement `json:"requirements"`

	// SoftwareDescription - The SoftLayer_Software_Description tied to this item. This will only be
	// populated for software items.
	SoftwareDescription *SoftLayer_Software_Description `json:"softwareDescription"`

	// SoftwareDescriptionId - The unique identifier of the SoftLayer_Software_Description tied to this
	// item.
	SoftwareDescriptionId int `json:"softwareDescriptionId"`

	// TaxCategory - no documentation
	TaxCategory *SoftLayer_Product_Item_Tax_Category `json:"taxCategory"`

	// ThirdPartySupportVendor - The 3rd party vendor for a support subscription item. (Deprecated)
	ThirdPartySupportVendor string `json:"thirdPartySupportVendor"`

	// TotalPhysicalCoreCapacity - The total number of physical processing cores (excluding virtual cores /
	// hyperthreads) for this server.
	TotalPhysicalCoreCapacity int `json:"totalPhysicalCoreCapacity"`

	// TotalPhysicalCoreCount - <nil>
	TotalPhysicalCoreCount string `json:"totalPhysicalCoreCount"`

	// TotalProcessorCapacity - no documentation
	TotalProcessorCapacity int `json:"totalProcessorCapacity"`

	// Units - The unit of measurement that a product item is measured in.
	Units string `json:"units"`

	// UpgradeItem - Some product items have an upgrade path. This is the next product item in the upgrade
	// path.
	UpgradeItem *SoftLayer_Product_Item `json:"upgradeItem"`

	// UpgradeItemCount - A count of some product items have an upgrade path. These are those upgrade
	// product items.
	UpgradeItemCount uint64 `json:"upgradeItemCount"`

	// UpgradeItemId - A products upgrade item's internal identification number
	UpgradeItemId int `json:"upgradeItemId"`

	// UpgradeItems - Some product items have an upgrade path. These are those upgrade product items.
	UpgradeItems []*SoftLayer_Product_Item `json:"upgradeItems"`
}
