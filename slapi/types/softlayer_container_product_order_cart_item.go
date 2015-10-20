package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Order_Cart_Item - This object represents an item in the cart, an entire
// configuration (based on a package ID) is considered a cart item.
type SoftLayer_Container_Product_Order_Cart_Item struct {

	// DatacenterDescription - used to identify description for datacenter selected by the user.
	DatacenterDescription string `json:"datacenterDescription,omitempty"`

	// ImageTemplateId - Optional Virtual Disk Image ID to be used as an installation base
	ImageTemplateId int `json:"imageTemplateId,omitempty"`

	// PackageType - no documentation
	PackageType string `json:"packageType,omitempty"`

	// Quantity - no documentation
	Quantity int `json:"quantity,omitempty"`

	// PriceIds - no documentation
	PriceIds []int `json:"priceIds,omitempty"`

	// PrimaryDiskPartitionId - The id of the [[SoftLayer_Hardware_Component_Partition_Template]].
	PrimaryDiskPartitionId int `json:"primaryDiskPartitionId,omitempty"`

	// ResourceType - Optional type for the resource an additional service is being order for.
	ResourceType string `json:"resourceType,omitempty"`

	// SourceVirtualGuestId - Optional Existing Virtual Guest ID used when ordering a copy of it
	SourceVirtualGuestId int `json:"sourceVirtualGuestId,omitempty"`

	// PromotionCode - no documentation
	PromotionCode string `json:"promotionCode,omitempty"`

	// QuantityLimit - The maximum number of servers that can be ordered for this configuration
	QuantityLimit int `json:"quantityLimit,omitempty"`

	// VerifiedConfig - <nil>
	VerifiedConfig *SoftLayer_Product_Order `json:"verifiedConfig,omitempty"`

	// ConfigurationUrl - Contains the URL to be used when redirecting back to the configuration page
	ConfigurationUrl string `json:"configurationUrl,omitempty"`

	// DatacenterCountryLongName - used to identify the country long Name of the datacenter selected by the
	// user.
	DatacenterCountryLongName string `json:"datacenterCountryLongName,omitempty"`

	// Hash - String representing a hash to identify this item within a cart
	Hash string `json:"hash,omitempty"`

	// Location - no documentation
	Location string `json:"location,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// IscsiOsFormatTypeKeyName - no documentation
	IscsiOsFormatTypeKeyName string `json:"iscsiOsFormatTypeKeyName,omitempty"`

	// Properties - no documentation
	Properties []string `json:"properties,omitempty"`

	// TaxCacheHash - Hash used to identify the corresponding tax container
	TaxCacheHash string `json:"taxCacheHash,omitempty"`

	// BillingOrderItemId - The billing order item id that is sent in with quotes
	BillingOrderItemId int `json:"billingOrderItemId,omitempty"`

	// PresetId - no documentation
	PresetId int `json:"presetId,omitempty"`

	// HourlyPricesFlag - If set, it indicates that this is an hourly billed configuration
	HourlyPricesFlag bool `json:"hourlyPricesFlag,omitempty"`

	// OrderContainerType - The order container type to be used for this cart item.
	OrderContainerType string `json:"orderContainerType,omitempty"`

	// PackageId - no documentation
	PackageId int `json:"packageId,omitempty"`

	// ResourceId - Optional Resource Id for the object an additional service is being ordered for
	ResourceId int `json:"resourceId,omitempty"`

	// ClusterResourceId - Used to identify which gateway is being upgraded to
	ClusterResourceId int `json:"clusterResourceId,omitempty"`

	// DatacenterCountry - used to identify the country code of the datacenter selected by the user.
	DatacenterCountry string `json:"datacenterCountry,omitempty"`

	// StorageGroups - Array of storage groups to be used for this cart item.
	StorageGroups []*SoftLayer_Container_Product_Order_Storage_Group `json:"storageGroups,omitempty"`
}

func (softlayer_container_product_order_cart_item *SoftLayer_Container_Product_Order_Cart_Item) String() string {
	return "SoftLayer_Container_Product_Order_Cart_Item"
}
