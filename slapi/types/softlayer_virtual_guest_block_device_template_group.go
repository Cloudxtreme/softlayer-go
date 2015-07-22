package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
)

// SoftLayer_Virtual_Guest_Block_Device_Template_Group - The virtual block device template group data
// type presents the structure in which a group of archived image templates will be presented. The
// structure consists of a parent template group which contain multiple child template group objects.
// Each child template group object represents the image template in a particular location. Unless
// editing/deleting a specific child template group object, it is best to use the parent object. A
// virtual block device template group, also known as an image template group, represents an image of a
// virtual guest instance.
type SoftLayer_Virtual_Guest_Block_Device_Template_Group struct {

	// Account - A block device template group's [[SoftLayer_Account|account]].
	Account *SoftLayer_Account `json:"account"`

	// AccountContactCount - no documentation
	AccountContactCount uint64 `json:"accountContactCount"`

	// AccountContacts - <nil>
	AccountContacts []*SoftLayer_Account_Contact `json:"accountContacts"`

	// AccountId - A block device template group's [[SoftLayer_Account|account]] ID
	AccountId int `json:"accountId"`

	// AccountReferenceCount - A count of the accounts which may have read-only access to an image template
	// group. Will only be populated for parent template group objects.
	AccountReferenceCount uint64 `json:"accountReferenceCount"`

	// AccountReferences - The accounts which may have read-only access to an image template group. Will
	// only be populated for parent template group objects.
	AccountReferences []*SoftLayer_Virtual_Guest_Block_Device_Template_Group_Accounts `json:"accountReferences"`

	// BlockDeviceCount - A count of the block devices that are part of an image template group
	BlockDeviceCount uint64 `json:"blockDeviceCount"`

	// BlockDevices - The block devices that are part of an image template group
	BlockDevices []*SoftLayer_Virtual_Guest_Block_Device_Template `json:"blockDevices"`

	// BlockDevicesDiskSpaceTotal - The total disk space of all images in a image template group.
	BlockDevicesDiskSpaceTotal float32 `json:"blockDevicesDiskSpaceTotal"`

	// Children - The image template groups that are clones of an image template group.
	Children []*SoftLayer_Virtual_Guest_Block_Device_Template_Group `json:"children"`

	// ChildrenCount - A count of the image template groups that are clones of an image template group.
	ChildrenCount uint64 `json:"childrenCount"`

	// CreateDate - The date a block device template group was created.
	CreateDate *time.Time `json:"createDate"`

	// Datacenter - The location containing this image template group. Will only be populated for child
	// template group objects.
	Datacenter *SoftLayer_Location `json:"datacenter"`

	// DatacenterCount - A count of a collection of locations containing a copy of this image template
	// group. Will only be populated for parent template group objects.
	DatacenterCount uint64 `json:"datacenterCount"`

	// Datacenters - A collection of locations containing a copy of this image template group. Will only be
	// populated for parent template group objects.
	Datacenters []*SoftLayer_Location `json:"datacenters"`

	// FlexImageFlag - no documentation
	FlexImageFlag bool `json:"flexImageFlag"`

	// GlobalIdentifier - no documentation
	GlobalIdentifier string `json:"globalIdentifier"`

	// Id - no documentation
	Id int `json:"id"`

	// ImageType - The virtual disk image type of this template. Value will be populated on parent and
	// child, but only supports object filtering on the parent.
	ImageType string `json:"imageType"`

	// ImageTypeKeyName - The virtual disk image type keyname (e.g. etc) of this template. Value will be
	// populated on parent and child, but only supports object filtering on the parent.
	ImageTypeKeyName string `json:"imageTypeKeyName"`

	// Name - A user definable and optional name of a block device template group.
	Name string `json:"name"`

	// Note - A block device template group's user defined note.
	Note string `json:"note"`

	// Parent - The image template group that another image template group was cloned from.
	Parent *SoftLayer_Virtual_Guest_Block_Device_Template_Group `json:"parent"`

	// ParentId - A block device template group's
	// [[SoftLayer_Virtual_Guest_Block_Device_Template_Group|parent]] ID. This will only be set when a
	// template group is created from a previously existing template group
	ParentId int `json:"parentId"`

	// PublicFlag - <nil>
	PublicFlag int `json:"publicFlag"`

	// SshKeyCount - A count of the ssh keys to be implemented on the server when provisioned or reloaded
	// from an image template group.
	SshKeyCount uint64 `json:"sshKeyCount"`

	// SshKeys - The ssh keys to be implemented on the server when provisioned or reloaded from an image
	// template group.
	SshKeys []*SoftLayer_Security_Ssh_Key `json:"sshKeys"`

	// Status - no documentation
	Status *SoftLayer_Virtual_Guest_Block_Device_Template_Group_Status `json:"status"`

	// StatusId - A block device template group's
	// [[SoftLayer_Virtual_Guest_Block_Device_Template_Group_Status|status]] ID
	StatusId int `json:"statusId"`

	// StorageRepository - The storage repository that an image template group resides on.
	StorageRepository *SoftLayer_Virtual_Storage_Repository `json:"storageRepository"`

	// Summary - A block device template group's user defined summary.
	Summary string `json:"summary"`

	// TagReferenceCount - A count of the tags associated with this image template group.
	TagReferenceCount uint64 `json:"tagReferenceCount"`

	// TagReferences - The tags associated with this image template group.
	TagReferences []*SoftLayer_Tag_Reference `json:"tagReferences"`

	// Transaction - A transaction that is being performed on a image template group.
	Transaction *SoftLayer_Provisioning_Version1_Transaction `json:"transaction"`

	// TransactionId - A block device template group's
	// [[SoftLayer_Provisioning_Version1_Transaction|transaction]] ID. This will only be set when there is
	// a transaction being performed on the block device template group.
	TransactionId int `json:"transactionId"`

	// UserRecordId - A block device template group's [[SoftLayer_User|user]] ID
	UserRecordId int `json:"userRecordId"`
}

// AddLocations - This method will create transaction(s) to add available locations to an archive image
// template.
func (softlayer_virtual_guest_block_device_template_group *SoftLayer_Virtual_Guest_Block_Device_Template_Group) AddLocations(ctx *slapi.RequestContext, locations []SoftLayer_Location) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// CopyToExternalSource - Create a transaction to export/copy a template to an external source.
func (softlayer_virtual_guest_block_device_template_group *SoftLayer_Virtual_Guest_Block_Device_Template_Group) CopyToExternalSource(ctx *slapi.RequestContext, configuration SoftLayer_Container_Virtual_Guest_Block_Device_Template_Configuration) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// CreateFromExternalSource - Create a transaction to import a disk image from an external source and
// create a standard image template.
func (softlayer_virtual_guest_block_device_template_group *SoftLayer_Virtual_Guest_Block_Device_Template_Group) CreateFromExternalSource(ctx *slapi.RequestContext, configuration SoftLayer_Container_Virtual_Guest_Block_Device_Template_Configuration) (*SoftLayer_Virtual_Guest_Block_Device_Template_Group, error) {
	var returnValue *SoftLayer_Virtual_Guest_Block_Device_Template_Group
	return returnValue, nil
}

// CreatePublicArchiveTransaction - Create a transaction to copy archived block devices into public
// repository
func (softlayer_virtual_guest_block_device_template_group *SoftLayer_Virtual_Guest_Block_Device_Template_Group) CreatePublicArchiveTransaction(ctx *slapi.RequestContext, groupName string, summary string, note string, locations []SoftLayer_Location) (int, error) {
	var returnValue int
	return returnValue, nil
}

// DeleteObject - Deleting a block device template group is different from the deletion of other
// objects. A block device template group can contain several gigabytes of data in its disk images.
// This may take some time to delete and requires a transaction to be created. This method creates a
// transaction that will delete all resources associated with the block device template group.
func (softlayer_virtual_guest_block_device_template_group *SoftLayer_Virtual_Guest_Block_Device_Template_Group) DeleteObject(ctx *slapi.RequestContext) (*SoftLayer_Provisioning_Version1_Transaction, error) {
	var returnValue *SoftLayer_Provisioning_Version1_Transaction
	return returnValue, nil
}

// DenySharingAccess - This method will deny another SoftLayer customer account's previously given
// access to provision CloudLayer Computing Instances from an image template group. Template access
// should only be removed from the parent template group object, not the child.
func (softlayer_virtual_guest_block_device_template_group *SoftLayer_Virtual_Guest_Block_Device_Template_Group) DenySharingAccess(ctx *slapi.RequestContext, accountId int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObject - Edit an image template group's associated name and note. All other properties in the
// SoftLayer_Virtual_Guest_Block_Device_Template_Group data type are read-only.
func (softlayer_virtual_guest_block_device_template_group *SoftLayer_Virtual_Guest_Block_Device_Template_Group) EditObject(ctx *slapi.RequestContext, templateObject SoftLayer_Virtual_Guest_Block_Device_Template_Group) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_virtual_guest_block_device_template_group *SoftLayer_Virtual_Guest_Block_Device_Template_Group) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Virtual_Guest_Block_Device_Template_Group, error) {
	var returnValue *SoftLayer_Virtual_Guest_Block_Device_Template_Group
	return returnValue, nil
}

// GetPublicImages - <nil>
func (softlayer_virtual_guest_block_device_template_group *SoftLayer_Virtual_Guest_Block_Device_Template_Group) GetPublicImages(ctx *slapi.RequestContext) ([]*SoftLayer_Virtual_Guest_Block_Device_Template_Group, error) {
	var returnValue []*SoftLayer_Virtual_Guest_Block_Device_Template_Group
	return returnValue, nil
}

// GetStorageLocations - no documentation
func (softlayer_virtual_guest_block_device_template_group *SoftLayer_Virtual_Guest_Block_Device_Template_Group) GetStorageLocations(ctx *slapi.RequestContext) ([]*SoftLayer_Location, error) {
	var returnValue []*SoftLayer_Location
	return returnValue, nil
}

// GetVhdImportSoftwareDescriptions - Returns an array of SoftLayer_Software_Description that are
// supported for VHD imports.
func (softlayer_virtual_guest_block_device_template_group *SoftLayer_Virtual_Guest_Block_Device_Template_Group) GetVhdImportSoftwareDescriptions(ctx *slapi.RequestContext) ([]*SoftLayer_Software_Description, error) {
	var returnValue []*SoftLayer_Software_Description
	return returnValue, nil
}

// PermitSharingAccess - This method will permit another SoftLayer customer account access to provision
// CloudLayer Computing Instances from an image template group. Template access should only be given to
// the parent template group object, not the child.
func (softlayer_virtual_guest_block_device_template_group *SoftLayer_Virtual_Guest_Block_Device_Template_Group) PermitSharingAccess(ctx *slapi.RequestContext, accountId int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RemoveLocations - This method will create transaction(s) to remove available locations from an
// archive image template.
func (softlayer_virtual_guest_block_device_template_group *SoftLayer_Virtual_Guest_Block_Device_Template_Group) RemoveLocations(ctx *slapi.RequestContext, locations []SoftLayer_Location) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// SetAvailableLocations - Create transaction(s) to set the archived block device available locations
func (softlayer_virtual_guest_block_device_template_group *SoftLayer_Virtual_Guest_Block_Device_Template_Group) SetAvailableLocations(ctx *slapi.RequestContext, locations []SoftLayer_Location) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// SetTags - no documentation
func (softlayer_virtual_guest_block_device_template_group *SoftLayer_Virtual_Guest_Block_Device_Template_Group) SetTags(ctx *slapi.RequestContext, tags string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}
