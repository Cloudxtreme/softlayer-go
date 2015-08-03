package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Virtual_Storage_Repository - The SoftLayer_Virtual_Storage_Repository represents a web
// based storage system that can be accessed through many types of devices, interfaces, and other
// resources.
type SoftLayer_Virtual_Storage_Repository struct {

	// Id - no documentation
	Id int `json:"id"`

	// Description - A storage repositories description that describes its purpose or contents
	Description string `json:"description"`

	// TypeId - A storage repositories [[SoftLayer_Virtual_Storage_Repository_Type|type]] ID
	TypeId int `json:"typeId"`

	// PublicFlag - <nil>
	PublicFlag int `json:"publicFlag"`

	// Capacity - A storage repositories capacity measured in Giga-Bytes
	Capacity float32 `json:"capacity"`

	// Name - A storage repositories name that describes its purpose or contents
	Name string `json:"name"`
}

func (softlayer_virtual_storage_repository *SoftLayer_Virtual_Storage_Repository) String() string {
	return "SoftLayer_Virtual_Storage_Repository"
}

// SoftLayer_Virtual_Storage_Repository_Extended is SoftLayer_Virtual_Storage_Repository with all maskable types.
type SoftLayer_Virtual_Storage_Repository_Extended struct {
	SoftLayer_Virtual_Storage_Repository

	// Account - The [[SoftLayer_Account|account]] that a storage repository belongs to.
	Account *SoftLayer_Account `json:"account"`

	// BillingItem - no documentation
	BillingItem *SoftLayer_Billing_Item `json:"billingItem"`

	// Datacenter - The datacenter that a virtual storage repository resides in.
	Datacenter *SoftLayer_Location `json:"datacenter"`

	// DiskImageCount - A count of the [[SoftLayer_Virtual_Disk_Image|disk images]] that are in a storage
	// repository. Disk images are the virtual hard drives for a virtual guest.
	DiskImageCount uint64 `json:"diskImageCount"`

	// GuestCount - A count of the computing instances that have disk images in a storage repository.
	GuestCount uint64 `json:"guestCount"`

	// Guests - The computing instances that have disk images in a storage repository.
	Guests []*SoftLayer_Virtual_Guest `json:"guests"`

	// PublicImageBillingItem - The current billing item for a public storage repository.
	PublicImageBillingItem *SoftLayer_Billing_Item `json:"publicImageBillingItem"`

	// DiskImages - The [[SoftLayer_Virtual_Disk_Image|disk images]] that are in a storage repository. Disk
	// images are the virtual hard drives for a virtual guest.
	DiskImages []*SoftLayer_Virtual_Disk_Image `json:"diskImages"`

	// MetricTrackingObject - <nil>
	MetricTrackingObject *SoftLayer_Metric_Tracking_Object_Virtual_Storage_Repository `json:"metricTrackingObject"`

	// Type - A storage repository's [[SoftLayer_Virtual_Storage_Repository_Type|type]].
	Type *SoftLayer_Virtual_Storage_Repository_Type `json:"type"`
}

func (softlayer_virtual_storage_repository *SoftLayer_Virtual_Storage_Repository_Extended) String() string {
	return "SoftLayer_Virtual_Storage_Repository"
}
