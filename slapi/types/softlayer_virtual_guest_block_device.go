package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Virtual_Guest_Block_Device - The block device data type presents the structure in which
// all block devices will be presented. A block device attaches a disk image to a guest. Internally,
// the structure supports various virtualization platforms with no change to external interaction. A
// guest, also known as a virtual server, represents an allocation of resources on a virtual host.
type SoftLayer_Virtual_Guest_Block_Device struct {

	// HotPlugFlag - A flag indicating if a block device can be plugged into a computing instance without
	// having to shut down the instance.
	HotPlugFlag int `json:"hotPlugFlag,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// MountType - The type of device that a virtual block device is mounted as, either "Disk" for a
	// directly connected storage disk or for devices that are mounted as optical drives..
	MountType string `json:"mountType,omitempty"`

	// StatusId - The status of the device, either disconnected or connected
	StatusId int `json:"statusId,omitempty"`

	// Device - no documentation
	Device string `json:"device,omitempty"`

	// MountMode - The writing mode that a virtual block device is mounted as, either for read-only mode or
	// for read and write mode.
	MountMode string `json:"mountMode,omitempty"`

	// DiskImageId - A block device [[SoftLayer_Virtual_Disk_Image|disk image]]'s unique ID.
	DiskImageId int `json:"diskImageId,omitempty"`

	// BootableFlag - A flag indicating if a block device can be booted from.
	BootableFlag int `json:"bootableFlag,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// GuestId - The [[SoftLayer_Virtual_Guest|computing instance]] that a block device is associated with.
	GuestId int `json:"guestId,omitempty"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// Uuid - A block device's unique ID on a virtualization platform.
	Uuid string `json:"uuid,omitempty"`

	// Guest - The computing instance that this block device is attached to.
	Guest *SoftLayer_Virtual_Guest `json:"guest,omitempty"`

	// Status - <nil>
	Status *SoftLayer_Virtual_Guest_Block_Device_Status `json:"status,omitempty"`

	// DiskImage - The disk image that a block device connects to in a computing instance.
	DiskImage *SoftLayer_Virtual_Disk_Image `json:"diskImage,omitempty"`
}

func (softlayer_virtual_guest_block_device *SoftLayer_Virtual_Guest_Block_Device) String() string {
	return "SoftLayer_Virtual_Guest_Block_Device"
}
