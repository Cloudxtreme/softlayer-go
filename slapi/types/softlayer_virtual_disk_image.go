package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Virtual_Disk_Image - The virtual disk image data type presents the structure in which a
// virtual disk image will be presented. Virtual block devices are assigned to disk images.
type SoftLayer_Virtual_Disk_Image struct {

	// ParentId - The ID of the the disk image that this disk image is based on, if applicable.
	ParentId int `json:"parentId"`

	// TypeId - A disk image's [[SoftLayer_Virtual_Disk_Image_Type|type]] ID
	TypeId int `json:"typeId"`

	// StorageRepositoryId - The [[SoftLayer_Virtual_Storage_Repository|storage repository]] that a disk
	// image is in.
	StorageRepositoryId int `json:"storageRepositoryId"`

	// Description - no documentation
	Description string `json:"description"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// Name - A descriptive name used to identify a disk image to a user.
	Name string `json:"name"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Units - The unit of storage in which the size of the image is measured. Defaults to for gigabytes.
	Units string `json:"units"`

	// Checksum - no documentation
	Checksum string `json:"checksum"`

	// Capacity - no documentation
	Capacity int `json:"capacity"`

	// Uuid - A disk image's unique ID on a virtualization platform.
	Uuid string `json:"uuid"`

	// Id - no documentation
	Id int `json:"id"`
}

func (softlayer_virtual_disk_image *SoftLayer_Virtual_Disk_Image) String() string {
	return "SoftLayer_Virtual_Disk_Image"
}

// SoftLayer_Virtual_Disk_Image_Extended is SoftLayer_Virtual_Disk_Image with all maskable types.
type SoftLayer_Virtual_Disk_Image_Extended struct {
	SoftLayer_Virtual_Disk_Image

	// LocalDiskFlag - <nil>
	LocalDiskFlag bool `json:"localDiskFlag"`

	// SoftwareReferences - References to the software that resides on a disk image.
	SoftwareReferences []*SoftLayer_Virtual_Disk_Image_Software `json:"softwareReferences"`

	// StorageRepository - The storage repository that a disk image resides in.
	StorageRepository *SoftLayer_Virtual_Storage_Repository `json:"storageRepository"`

	// CoalescedDiskImageCount - no documentation
	CoalescedDiskImageCount uint64 `json:"coalescedDiskImageCount"`

	// SoftwareReferenceCount - A count of references to the software that resides on a disk image.
	SoftwareReferenceCount uint64 `json:"softwareReferenceCount"`

	// StorageRepositoryType - The type of storage repository that a disk image resides in.
	StorageRepositoryType *SoftLayer_Virtual_Storage_Repository_Type `json:"storageRepositoryType"`

	// SourceDiskImage - The original disk image that the current disk image was cloned from.
	SourceDiskImage *SoftLayer_Virtual_Disk_Image `json:"sourceDiskImage"`

	// Type - no documentation
	Type *SoftLayer_Virtual_Disk_Image_Type `json:"type"`

	// BillingItem - no documentation
	BillingItem *SoftLayer_Billing_Item_Virtual_Disk_Image `json:"billingItem"`

	// CoalescedDiskImages - <nil>
	CoalescedDiskImages []*SoftLayer_Virtual_Disk_Image `json:"coalescedDiskImages"`

	// CopyOnWriteFlag - <nil>
	CopyOnWriteFlag bool `json:"copyOnWriteFlag"`

	// MetadataFlag - Whether this disk image is meant for storage of custom user data supplied with a
	// Cloud Computing Instance order.
	MetadataFlag bool `json:"metadataFlag"`

	// BlockDeviceCount - A count of the block devices that a disk image is attached to. Block devices
	// connect computing instances to disk images.
	BlockDeviceCount uint64 `json:"blockDeviceCount"`

	// BlockDevices - The block devices that a disk image is attached to. Block devices connect computing
	// instances to disk images.
	BlockDevices []*SoftLayer_Virtual_Guest_Block_Device `json:"blockDevices"`

	// TemplateBlockDevice - The template that attaches a disk image to a
	// [[SoftLayer_Virtual_Guest_Block_Device_Template_Group|archive]].
	TemplateBlockDevice *SoftLayer_Virtual_Guest_Block_Device_Template `json:"templateBlockDevice"`
}

func (softlayer_virtual_disk_image *SoftLayer_Virtual_Disk_Image_Extended) String() string {
	return "SoftLayer_Virtual_Disk_Image"
}
