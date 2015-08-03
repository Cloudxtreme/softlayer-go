package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Virtual_Storage_Repository - The SoftLayer_Virtual_Storage_Repository represents a web
// based storage system that can be accessed through many types of devices, interfaces, and other
// resources.
type SoftLayer_Virtual_Storage_Repository struct {

	// Account - The [[SoftLayer_Account|account]] that a storage repository belongs to.
	Account *SoftLayer_Account `json:"account"`

	// BillingItem - no documentation
	BillingItem *SoftLayer_Billing_Item `json:"billingItem"`

	// Capacity - A storage repositories capacity measured in Giga-Bytes
	Capacity float32 `json:"capacity"`

	// Datacenter - The datacenter that a virtual storage repository resides in.
	Datacenter *SoftLayer_Location `json:"datacenter"`

	// Description - A storage repositories description that describes its purpose or contents
	Description string `json:"description"`

	// DiskImageCount - A count of the [[SoftLayer_Virtual_Disk_Image|disk images]] that are in a storage
	// repository. Disk images are the virtual hard drives for a virtual guest.
	DiskImageCount uint64 `json:"diskImageCount"`

	// DiskImages - The [[SoftLayer_Virtual_Disk_Image|disk images]] that are in a storage repository. Disk
	// images are the virtual hard drives for a virtual guest.
	DiskImages []*SoftLayer_Virtual_Disk_Image `json:"diskImages"`

	// GuestCount - A count of the computing instances that have disk images in a storage repository.
	GuestCount uint64 `json:"guestCount"`

	// Guests - The computing instances that have disk images in a storage repository.
	Guests []*SoftLayer_Virtual_Guest `json:"guests"`

	// Id - no documentation
	Id int `json:"id"`

	// MetricTrackingObject - <nil>
	MetricTrackingObject *SoftLayer_Metric_Tracking_Object_Virtual_Storage_Repository `json:"metricTrackingObject"`

	// Name - A storage repositories name that describes its purpose or contents
	Name string `json:"name"`

	// PublicFlag - <nil>
	PublicFlag int `json:"publicFlag"`

	// PublicImageBillingItem - The current billing item for a public storage repository.
	PublicImageBillingItem *SoftLayer_Billing_Item `json:"publicImageBillingItem"`

	// Type - A storage repository's [[SoftLayer_Virtual_Storage_Repository_Type|type]].
	Type *SoftLayer_Virtual_Storage_Repository_Type `json:"type"`

	// TypeId - A storage repositories [[SoftLayer_Virtual_Storage_Repository_Type|type]] ID
	TypeId int `json:"typeId"`
}

func (softlayer_virtual_storage_repository *SoftLayer_Virtual_Storage_Repository) String() string {
	return "SoftLayer_Virtual_Storage_Repository"
}
