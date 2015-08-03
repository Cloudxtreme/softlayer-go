package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Security_Ssh_Key - <nil>
type SoftLayer_Security_Ssh_Key struct {

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`

	// BlockDeviceTemplateGroupCount - A count of the image template groups that are linked to an SSH key.
	BlockDeviceTemplateGroupCount uint64 `json:"blockDeviceTemplateGroupCount"`

	// BlockDeviceTemplateGroups - The image template groups that are linked to an SSH key.
	BlockDeviceTemplateGroups []*SoftLayer_Virtual_Guest_Block_Device_Template_Group `json:"blockDeviceTemplateGroups"`

	// CreateDate - The date a ssh key was added. This property is read only. Changes made will be silently
	// ignored.
	CreateDate *time.Time `json:"createDate"`

	// Fingerprint - A short sequence of bytes used to authenticate or lookup a longer ssh key. This will
	// automatically be generated upon adding or modifying the ssh key. This property is read only. Changes
	// made will be silently ignored.
	Fingerprint string `json:"fingerprint"`

	// Id - no documentation
	Id int `json:"id"`

	// Key - no documentation
	Key string `json:"key"`

	// Label - no documentation
	Label string `json:"label"`

	// ModifyDate - The date a ssh key was last modified. This property is read only. Changes made will be
	// silently ignored.
	ModifyDate *time.Time `json:"modifyDate"`

	// Notes - A small note about a ssh key to use at your discretion.
	Notes string `json:"notes"`

	// SoftwarePasswordCount - A count of the OS root users that are linked to an SSH key.
	SoftwarePasswordCount uint64 `json:"softwarePasswordCount"`

	// SoftwarePasswords - no documentation
	SoftwarePasswords []*SoftLayer_Software_Component_Password `json:"softwarePasswords"`
}

func (softlayer_security_ssh_key *SoftLayer_Security_Ssh_Key) String() string {
	return "SoftLayer_Security_Ssh_Key"
}
