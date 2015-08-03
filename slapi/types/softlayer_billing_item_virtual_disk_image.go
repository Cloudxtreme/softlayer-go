package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Billing_Item_Virtual_Disk_Image - The SoftLayer_Billing_Item_Virtual_Disk_Image data type
// contains general information relating to a single SoftLayer billing item for disk images.
type SoftLayer_Billing_Item_Virtual_Disk_Image struct {

	// ResourceTableId - The resource (unique identifier) for a disk image billing item.
	ResourceTableId int `json:"resourceTableId"`
}

// SoftLayer_Billing_Item_Virtual_Disk_Image_Extended is SoftLayer_Billing_Item_Virtual_Disk_Image with all maskable types.
type SoftLayer_Billing_Item_Virtual_Disk_Image_Extended struct {
	SoftLayer_Billing_Item_Virtual_Disk_Image

	// Resource - no documentation
	Resource *SoftLayer_Virtual_Disk_Image `json:"resource"`
}

func (softlayer_billing_item_virtual_disk_image *SoftLayer_Billing_Item_Virtual_Disk_Image) String() string {
	return "SoftLayer_Billing_Item_Virtual_Disk_Image"
}
