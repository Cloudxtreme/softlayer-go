package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"

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

	// CreateDate - The date a block device template group was created.
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Name - A user definable and optional name of a block device template group.
	Name string `json:"name,omitempty"`

	// TransactionId - A block device template group's
	// [[SoftLayer_Provisioning_Version1_Transaction|transaction]] ID. This will only be set when there is
	// a transaction being performed on the block device template group.
	TransactionId int `json:"transactionId,omitempty"`

	// ParentId - A block device template group's
	// [[SoftLayer_Virtual_Guest_Block_Device_Template_Group|parent]] ID. This will only be set when a
	// template group is created from a previously existing template group
	ParentId int `json:"parentId,omitempty"`

	// StatusId - A block device template group's
	// [[SoftLayer_Virtual_Guest_Block_Device_Template_Group_Status|status]] ID
	StatusId int `json:"statusId,omitempty"`

	// PublicFlag - <nil>
	PublicFlag int `json:"publicFlag,omitempty"`

	// Note - A block device template group's user defined note.
	Note string `json:"note,omitempty"`

	// Summary - A block device template group's user defined summary.
	Summary string `json:"summary,omitempty"`

	// UserRecordId - A block device template group's [[SoftLayer_User|user]] ID
	UserRecordId int `json:"userRecordId,omitempty"`

	// AccountId - A block device template group's [[SoftLayer_Account|account]] ID
	AccountId int `json:"accountId,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`
}

func (softlayer_virtual_guest_block_device_template_group *SoftLayer_Virtual_Guest_Block_Device_Template_Group) String() string {
	return "SoftLayer_Virtual_Guest_Block_Device_Template_Group"
}

// SoftLayer_Virtual_Guest_Block_Device_Template_Group_Extended is SoftLayer_Virtual_Guest_Block_Device_Template_Group with all maskable types.
type SoftLayer_Virtual_Guest_Block_Device_Template_Group_Extended struct {
	SoftLayer_Virtual_Guest_Block_Device_Template_Group

	// ChildrenCount - A count of the image template groups that are clones of an image template group.
	ChildrenCount uint64 `json:"childrenCount,omitempty"`

	// DatacenterCount - A count of a collection of locations containing a copy of this image template
	// group. Will only be populated for parent template group objects.
	DatacenterCount uint64 `json:"datacenterCount,omitempty"`

	// AccountContacts - <nil>
	AccountContacts []*SoftLayer_Account_Contact `json:"accountContacts,omitempty"`

	// Children - The image template groups that are clones of an image template group.
	Children []*SoftLayer_Virtual_Guest_Block_Device_Template_Group `json:"children,omitempty"`

	// Datacenter - The location containing this image template group. Will only be populated for child
	// template group objects.
	Datacenter *SoftLayer_Location `json:"datacenter,omitempty"`

	// Status - no documentation
	Status *SoftLayer_Virtual_Guest_Block_Device_Template_Group_Status `json:"status,omitempty"`

	// ImageType - The virtual disk image type of this template. Value will be populated on parent and
	// child, but only supports object filtering on the parent.
	ImageType string `json:"imageType,omitempty"`

	// BlockDeviceCount - A count of the block devices that are part of an image template group
	BlockDeviceCount uint64 `json:"blockDeviceCount,omitempty"`

	// Account - A block device template group's [[SoftLayer_Account|account]].
	Account *SoftLayer_Account `json:"account,omitempty"`

	// AccountReferences - The accounts which may have read-only access to an image template group. Will
	// only be populated for parent template group objects.
	AccountReferences []*SoftLayer_Virtual_Guest_Block_Device_Template_Group_Accounts `json:"accountReferences,omitempty"`

	// SshKeyCount - A count of the ssh keys to be implemented on the server when provisioned or reloaded
	// from an image template group.
	SshKeyCount uint64 `json:"sshKeyCount,omitempty"`

	// AccountReferenceCount - A count of the accounts which may have read-only access to an image template
	// group. Will only be populated for parent template group objects.
	AccountReferenceCount uint64 `json:"accountReferenceCount,omitempty"`

	// BlockDevicesDiskSpaceTotal - The total disk space of all images in a image template group.
	BlockDevicesDiskSpaceTotal slapi.Float64 `json:"blockDevicesDiskSpaceTotal,omitempty"`

	// Datacenters - A collection of locations containing a copy of this image template group. Will only be
	// populated for parent template group objects.
	Datacenters []*SoftLayer_Location `json:"datacenters,omitempty"`

	// FlexImageFlag - no documentation
	FlexImageFlag bool `json:"flexImageFlag,omitempty"`

	// ImageTypeKeyName - The virtual disk image type keyname (e.g. etc) of this template. Value will be
	// populated on parent and child, but only supports object filtering on the parent.
	ImageTypeKeyName string `json:"imageTypeKeyName,omitempty"`

	// SshKeys - The ssh keys to be implemented on the server when provisioned or reloaded from an image
	// template group.
	SshKeys []*SoftLayer_Security_Ssh_Key `json:"sshKeys,omitempty"`

	// TagReferences - The tags associated with this image template group.
	TagReferences []*SoftLayer_Tag_Reference `json:"tagReferences,omitempty"`

	// AccountContactCount - no documentation
	AccountContactCount uint64 `json:"accountContactCount,omitempty"`

	// BlockDevices - The block devices that are part of an image template group
	BlockDevices []*SoftLayer_Virtual_Guest_Block_Device_Template `json:"blockDevices,omitempty"`

	// GlobalIdentifier - no documentation
	GlobalIdentifier string `json:"globalIdentifier,omitempty"`

	// Parent - The image template group that another image template group was cloned from.
	Parent *SoftLayer_Virtual_Guest_Block_Device_Template_Group `json:"parent,omitempty"`

	// StorageRepository - The storage repository that an image template group resides on.
	StorageRepository *SoftLayer_Virtual_Storage_Repository `json:"storageRepository,omitempty"`

	// Transaction - A transaction that is being performed on a image template group.
	Transaction *SoftLayer_Provisioning_Version1_Transaction `json:"transaction,omitempty"`

	// TagReferenceCount - A count of the tags associated with this image template group.
	TagReferenceCount uint64 `json:"tagReferenceCount,omitempty"`
}

func (softlayer_virtual_guest_block_device_template_group *SoftLayer_Virtual_Guest_Block_Device_Template_Group_Extended) String() string {
	return "SoftLayer_Virtual_Guest_Block_Device_Template_Group"
}
