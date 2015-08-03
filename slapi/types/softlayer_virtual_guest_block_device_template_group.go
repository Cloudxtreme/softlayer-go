package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Virtual_Guest_Block_Device_Template_Group - The virtual block device template group data
// type presents the structure in which a group of archived image templates will be presented. The
// structure consists of a parent template group which contain multiple child template group objects.
// Each child template group object represents the image template in a particular location. Unless
// editing/deleting a specific child template group object, it is best to use the parent object. A
// virtual block device template group, also known as an image template group, represents an image of a
// virtual guest instance.
type SoftLayer_Virtual_Guest_Block_Device_Template_Group struct {

	// StatusId - A block device template group's
	// [[SoftLayer_Virtual_Guest_Block_Device_Template_Group_Status|status]] ID
	StatusId int `json:"statusId"`

	// PublicFlag - <nil>
	PublicFlag int `json:"publicFlag"`

	// Summary - A block device template group's user defined summary.
	Summary string `json:"summary"`

	// Note - A block device template group's user defined note.
	Note string `json:"note"`

	// ParentId - A block device template group's
	// [[SoftLayer_Virtual_Guest_Block_Device_Template_Group|parent]] ID. This will only be set when a
	// template group is created from a previously existing template group
	ParentId int `json:"parentId"`

	// Name - A user definable and optional name of a block device template group.
	Name string `json:"name"`

	// Id - no documentation
	Id int `json:"id"`

	// TransactionId - A block device template group's
	// [[SoftLayer_Provisioning_Version1_Transaction|transaction]] ID. This will only be set when there is
	// a transaction being performed on the block device template group.
	TransactionId int `json:"transactionId"`

	// UserRecordId - A block device template group's [[SoftLayer_User|user]] ID
	UserRecordId int `json:"userRecordId"`

	// AccountId - A block device template group's [[SoftLayer_Account|account]] ID
	AccountId int `json:"accountId"`

	// CreateDate - The date a block device template group was created.
	CreateDate *time.Time `json:"createDate"`
}

// SoftLayer_Virtual_Guest_Block_Device_Template_Group_Extended is SoftLayer_Virtual_Guest_Block_Device_Template_Group with all maskable types.
type SoftLayer_Virtual_Guest_Block_Device_Template_Group_Extended struct {
	SoftLayer_Virtual_Guest_Block_Device_Template_Group

	// SshKeys - The ssh keys to be implemented on the server when provisioned or reloaded from an image
	// template group.
	SshKeys []*SoftLayer_Security_Ssh_Key `json:"sshKeys"`

	// AccountContacts - <nil>
	AccountContacts []*SoftLayer_Account_Contact `json:"accountContacts"`

	// FlexImageFlag - no documentation
	FlexImageFlag bool `json:"flexImageFlag"`

	// GlobalIdentifier - no documentation
	GlobalIdentifier string `json:"globalIdentifier"`

	// Transaction - A transaction that is being performed on a image template group.
	Transaction *SoftLayer_Provisioning_Version1_Transaction `json:"transaction"`

	// BlockDeviceCount - A count of the block devices that are part of an image template group
	BlockDeviceCount uint64 `json:"blockDeviceCount"`

	// BlockDevicesDiskSpaceTotal - The total disk space of all images in a image template group.
	BlockDevicesDiskSpaceTotal float32 `json:"blockDevicesDiskSpaceTotal"`

	// Parent - The image template group that another image template group was cloned from.
	Parent *SoftLayer_Virtual_Guest_Block_Device_Template_Group `json:"parent"`

	// TagReferences - The tags associated with this image template group.
	TagReferences []*SoftLayer_Tag_Reference `json:"tagReferences"`

	// TagReferenceCount - A count of the tags associated with this image template group.
	TagReferenceCount uint64 `json:"tagReferenceCount"`

	// Children - The image template groups that are clones of an image template group.
	Children []*SoftLayer_Virtual_Guest_Block_Device_Template_Group `json:"children"`

	// ImageType - The virtual disk image type of this template. Value will be populated on parent and
	// child, but only supports object filtering on the parent.
	ImageType string `json:"imageType"`

	// DatacenterCount - A count of a collection of locations containing a copy of this image template
	// group. Will only be populated for parent template group objects.
	DatacenterCount uint64 `json:"datacenterCount"`

	// Account - A block device template group's [[SoftLayer_Account|account]].
	Account *SoftLayer_Account `json:"account"`

	// SshKeyCount - A count of the ssh keys to be implemented on the server when provisioned or reloaded
	// from an image template group.
	SshKeyCount uint64 `json:"sshKeyCount"`

	// Datacenters - A collection of locations containing a copy of this image template group. Will only be
	// populated for parent template group objects.
	Datacenters []*SoftLayer_Location `json:"datacenters"`

	// StorageRepository - The storage repository that an image template group resides on.
	StorageRepository *SoftLayer_Virtual_Storage_Repository `json:"storageRepository"`

	// AccountReferenceCount - A count of the accounts which may have read-only access to an image template
	// group. Will only be populated for parent template group objects.
	AccountReferenceCount uint64 `json:"accountReferenceCount"`

	// BlockDevices - The block devices that are part of an image template group
	BlockDevices []*SoftLayer_Virtual_Guest_Block_Device_Template `json:"blockDevices"`

	// Datacenter - The location containing this image template group. Will only be populated for child
	// template group objects.
	Datacenter *SoftLayer_Location `json:"datacenter"`

	// ImageTypeKeyName - The virtual disk image type keyname (e.g. etc) of this template. Value will be
	// populated on parent and child, but only supports object filtering on the parent.
	ImageTypeKeyName string `json:"imageTypeKeyName"`

	// AccountContactCount - no documentation
	AccountContactCount uint64 `json:"accountContactCount"`

	// ChildrenCount - A count of the image template groups that are clones of an image template group.
	ChildrenCount uint64 `json:"childrenCount"`

	// AccountReferences - The accounts which may have read-only access to an image template group. Will
	// only be populated for parent template group objects.
	AccountReferences []*SoftLayer_Virtual_Guest_Block_Device_Template_Group_Accounts `json:"accountReferences"`

	// Status - no documentation
	Status *SoftLayer_Virtual_Guest_Block_Device_Template_Group_Status `json:"status"`
}

func (softlayer_virtual_guest_block_device_template_group *SoftLayer_Virtual_Guest_Block_Device_Template_Group) String() string {
	return "SoftLayer_Virtual_Guest_Block_Device_Template_Group"
}
