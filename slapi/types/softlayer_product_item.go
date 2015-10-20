package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Product_Item - The SoftLayer_Product_Item data type contains general information relating
// to a single SoftLayer product.
type SoftLayer_Product_Item struct {

	// Capacity - Some Product Items have capacity information such as RAM and bandwidth, and others. This
	// provides the numerical representation of the capacity given in the description of this product item.
	Capacity slapi.Float64 `json:"capacity,omitempty"`

	// SoftwareDescriptionId - The unique identifier of the SoftLayer_Software_Description tied to this
	// item.
	SoftwareDescriptionId int `json:"softwareDescriptionId,omitempty"`

	// LongDescription - no documentation
	LongDescription string `json:"longDescription,omitempty"`

	// Description - no documentation
	Description string `json:"description,omitempty"`

	// ItemTaxCategoryId - A products tax category internal identification number
	ItemTaxCategoryId int `json:"itemTaxCategoryId,omitempty"`

	// Units - The unit of measurement that a product item is measured in.
	Units string `json:"units,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// KeyName - no documentation
	KeyName string `json:"keyName,omitempty"`

	// UpgradeItemId - A products upgrade item's internal identification number
	UpgradeItemId int `json:"upgradeItemId,omitempty"`

	// HardwareGenericComponentModel - The generic hardware component that this item represents.
	HardwareGenericComponentModel *SoftLayer_Hardware_Component_Model_Generic `json:"hardwareGenericComponentModel,omitempty"`

	// AttributeCount - A count of the attribute values for a product item. These are additional properties
	// that give extra information about the product being sold.
	AttributeCount uint64 `json:"attributeCount,omitempty"`

	// AvailabilityAttributeCount - A count of attributes that govern when an item may no longer be
	// available.
	AvailabilityAttributeCount uint64 `json:"availabilityAttributeCount,omitempty"`

	// SoftwareDescription - The SoftLayer_Software_Description tied to this item. This will only be
	// populated for software items.
	SoftwareDescription *SoftLayer_Software_Description `json:"softwareDescription,omitempty"`

	// TotalProcessorCapacity - no documentation
	TotalProcessorCapacity int `json:"totalProcessorCapacity,omitempty"`

	// ConflictCount - A count of an item's conflicts. For example, McAfee LinuxShield cannot be ordered
	// with Windows. It was not meant for that operating system and as such is a conflict.
	ConflictCount uint64 `json:"conflictCount,omitempty"`

	// PresaleEventCount - no documentation
	PresaleEventCount uint64 `json:"presaleEventCount,omitempty"`

	// UpgradeItemCount - A count of some product items have an upgrade path. These are those upgrade
	// product items.
	UpgradeItemCount uint64 `json:"upgradeItemCount,omitempty"`

	// CoreRestrictedItemFlag - This flag indicates that this product is restricted by the number of cores
	// on the compute instance.
	CoreRestrictedItemFlag bool `json:"coreRestrictedItemFlag,omitempty"`

	// DowngradeItems - Some product items have a downgrade path. These are those product items.
	DowngradeItems []*SoftLayer_Product_Item `json:"downgradeItems,omitempty"`

	// PhysicalCoreCapacity - no documentation
	PhysicalCoreCapacity string `json:"physicalCoreCapacity,omitempty"`

	// ConfigurationTemplateCount - A count of some product items have configuration templates which can be
	// used to during provisioning of that product.
	ConfigurationTemplateCount uint64 `json:"configurationTemplateCount,omitempty"`

	// AvailabilityAttributes - Attributes that govern when an item may no longer be available.
	AvailabilityAttributes []*SoftLayer_Product_Item_Attribute `json:"availabilityAttributes,omitempty"`

	// PriceCount - no documentation
	PriceCount uint64 `json:"priceCount,omitempty"`

	// HideFromPortalFlag - <nil>
	HideFromPortalFlag bool `json:"hideFromPortalFlag,omitempty"`

	// Packages - A collection of all the SoftLayer_Product_Package(s) in which this item exists.
	Packages []*SoftLayer_Product_Package `json:"packages,omitempty"`

	// UpgradeItem - Some product items have an upgrade path. This is the next product item in the upgrade
	// path.
	UpgradeItem *SoftLayer_Product_Item `json:"upgradeItem,omitempty"`

	// BillingType - no documentation
	BillingType string `json:"billingType,omitempty"`

	// Bundle - An item's included products. Some items have other items included in them that we
	// specifically detail. They are here called Bundled Items. An example is Plesk unlimited. It as a
	// bundled item labeled 'SiteBuilder'. These are the SoftLayer_Product_Item_Bundles objects.
	Bundle []*SoftLayer_Product_Item_Bundles `json:"bundle,omitempty"`

	// ThirdPartySupportVendor - The 3rd party vendor for a support subscription item. (Deprecated)
	ThirdPartySupportVendor string `json:"thirdPartySupportVendor,omitempty"`

	// TotalPhysicalCoreCapacity - The total number of physical processing cores (excluding virtual cores /
	// hyperthreads) for this server.
	TotalPhysicalCoreCapacity int `json:"totalPhysicalCoreCapacity,omitempty"`

	// BundleCount - A count of an item's included products. Some items have other items included in them
	// that we specifically detail. They are here called Bundled Items. An example is Plesk unlimited. It
	// as a bundled item labeled 'SiteBuilder'. These are the SoftLayer_Product_Item_Bundles objects.
	BundleCount uint64 `json:"bundleCount,omitempty"`

	// DowngradeItemCount - A count of some product items have a downgrade path. These are those product
	// items.
	DowngradeItemCount uint64 `json:"downgradeItemCount,omitempty"`

	// PackageCount - A count of a collection of all the SoftLayer_Product_Package(s) in which this item
	// exists.
	PackageCount uint64 `json:"packageCount,omitempty"`

	// LocationConflicts - An item's location conflicts. For example, Dual Path network functionality
	// cannot be ordered in WDC and as such is a conflict.
	LocationConflicts []*SoftLayer_Product_Item_Resource_Conflict `json:"locationConflicts,omitempty"`

	// ActivePresaleEventCount - no documentation
	ActivePresaleEventCount uint64 `json:"activePresaleEventCount,omitempty"`

	// ActiveUsagePriceCount - no documentation
	ActiveUsagePriceCount uint64 `json:"activeUsagePriceCount,omitempty"`

	// CategoryCount - no documentation
	CategoryCount uint64 `json:"categoryCount,omitempty"`

	// InventoryCount - A count of an item's inventory status per datacenter.
	InventoryCount uint64 `json:"inventoryCount,omitempty"`

	// ActivePresaleEvents - <nil>
	ActivePresaleEvents []*SoftLayer_Sales_Presale_Event `json:"activePresaleEvents,omitempty"`

	// IsEngineeredServerProduct - Flag to indicate the server product is engineered for a multi-server
	// solution. (Deprecated)
	IsEngineeredServerProduct bool `json:"isEngineeredServerProduct,omitempty"`

	// Requirements - If an item must be ordered with another item, it will have a requirement item here.
	Requirements []*SoftLayer_Product_Item_Requirement `json:"requirements,omitempty"`

	// ConfigurationTemplates - Some product items have configuration templates which can be used to during
	// provisioning of that product.
	ConfigurationTemplates []*SoftLayer_Configuration_Template `json:"configurationTemplates,omitempty"`

	// Inventory - no documentation
	Inventory []*SoftLayer_Product_Package_Inventory `json:"inventory,omitempty"`

	// Prices - no documentation
	Prices []*SoftLayer_Product_Item_Price `json:"prices,omitempty"`

	// TotalPhysicalCoreCount - <nil>
	TotalPhysicalCoreCount string `json:"totalPhysicalCoreCount,omitempty"`

	// Conflicts - An item's conflicts. For example, McAfee LinuxShield cannot be ordered with Windows. It
	// was not meant for that operating system and as such is a conflict.
	Conflicts []*SoftLayer_Product_Item_Resource_Conflict `json:"conflicts,omitempty"`

	// UpgradeItems - Some product items have an upgrade path. These are those upgrade product items.
	UpgradeItems []*SoftLayer_Product_Item `json:"upgradeItems,omitempty"`

	// ItemCategory - no documentation
	ItemCategory *SoftLayer_Product_Item_Category `json:"itemCategory,omitempty"`

	// PresaleEvents - <nil>
	PresaleEvents []*SoftLayer_Sales_Presale_Event `json:"presaleEvents,omitempty"`

	// Attributes - The attribute values for a product item. These are additional properties that give
	// extra information about the product being sold.
	Attributes []*SoftLayer_Product_Item_Attribute `json:"attributes,omitempty"`

	// Categories - no documentation
	Categories []*SoftLayer_Product_Item_Category `json:"categories,omitempty"`

	// TaxCategory - no documentation
	TaxCategory *SoftLayer_Product_Item_Tax_Category `json:"taxCategory,omitempty"`

	// RequirementCount - A count of if an item must be ordered with another item, it will have a
	// requirement item here.
	RequirementCount uint64 `json:"requirementCount,omitempty"`

	// ObjectStorageItemFlag - <nil>
	ObjectStorageItemFlag bool `json:"objectStorageItemFlag,omitempty"`

	// GlobalCategoryConflictCount - A count of an item's category conflicts. For example, 10 Gbps
	// redundant network functionality cannot be ordered with a secondary GPU and as such is a conflict.
	GlobalCategoryConflictCount uint64 `json:"globalCategoryConflictCount,omitempty"`

	// ActiveUsagePrices - no documentation
	ActiveUsagePrices []*SoftLayer_Product_Item_Price `json:"activeUsagePrices,omitempty"`

	// DowngradeItem - Some product items have a downgrade path. This is the first product item in the
	// downgrade path.
	DowngradeItem *SoftLayer_Product_Item `json:"downgradeItem,omitempty"`

	// GlobalCategoryConflicts - An item's category conflicts. For example, 10 Gbps redundant network
	// functionality cannot be ordered with a secondary GPU and as such is a conflict.
	GlobalCategoryConflicts []*SoftLayer_Product_Item_Resource_Conflict `json:"globalCategoryConflicts,omitempty"`

	// LocationConflictCount - A count of an item's location conflicts. For example, Dual Path network
	// functionality cannot be ordered in WDC and as such is a conflict.
	LocationConflictCount uint64 `json:"locationConflictCount,omitempty"`
}

func (softlayer_product_item *SoftLayer_Product_Item) String() string {
	return "SoftLayer_Product_Item"
}
