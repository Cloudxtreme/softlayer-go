package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Virtual_Guest_Block_Device_Template_Group_Accounts - The
// SoftLayer_Virtual_Guest_Block_Device_Template_Group_Accounts data type represents the SoftLayer
// customer accounts which have access to provision CloudLayer Computing Instances from an image
// template group. All accounts other than the image template group owner have read-only access to that
// image template group. It is important to note that this data type should only exist to give accounts
// access to the parent template group object, not the child. All image template sharing between
// accounts should occur on the parent object.
type SoftLayer_Virtual_Guest_Block_Device_Template_Group_Accounts struct {

	// AccountId - The [[SoftLayer_Account|account]] ID which will have access to an image.
	AccountId int `json:"accountId,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// GroupId - The [[SoftLayer_Virtual_Guest_Block_Device_Template_Group|group]] ID which access will be
	// granted to.
	GroupId int `json:"groupId,omitempty"`
}

func (softlayer_virtual_guest_block_device_template_group_accounts *SoftLayer_Virtual_Guest_Block_Device_Template_Group_Accounts) String() string {
	return "SoftLayer_Virtual_Guest_Block_Device_Template_Group_Accounts"
}

// SoftLayer_Virtual_Guest_Block_Device_Template_Group_Accounts_Extended is SoftLayer_Virtual_Guest_Block_Device_Template_Group_Accounts with all maskable types.
type SoftLayer_Virtual_Guest_Block_Device_Template_Group_Accounts_Extended struct {
	SoftLayer_Virtual_Guest_Block_Device_Template_Group_Accounts

	// Account - The [[SoftLayer_Account|account]] that an image template group is shared with.
	Account *SoftLayer_Account `json:"account,omitempty"`

	// Group - The [[SoftLayer_Virtual_Guest_Block_Device_Template_Group|image template group]] that is
	// shared with an account.
	Group *SoftLayer_Virtual_Guest_Block_Device_Template_Group `json:"group,omitempty"`
}

func (softlayer_virtual_guest_block_device_template_group_accounts *SoftLayer_Virtual_Guest_Block_Device_Template_Group_Accounts_Extended) String() string {
	return "SoftLayer_Virtual_Guest_Block_Device_Template_Group_Accounts"
}
