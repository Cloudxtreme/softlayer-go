package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Security_Ssh_Key - <nil>
type SoftLayer_Security_Ssh_Key struct {

	// Key - no documentation
	Key string `json:"key,omitempty"`

	// Fingerprint - A short sequence of bytes used to authenticate or lookup a longer ssh key. This will
	// automatically be generated upon adding or modifying the ssh key. This property is read only. Changes
	// made will be silently ignored.
	Fingerprint string `json:"fingerprint,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// Label - no documentation
	Label string `json:"label,omitempty"`

	// ModifyDate - The date a ssh key was last modified. This property is read only. Changes made will be
	// silently ignored.
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// Notes - A small note about a ssh key to use at your discretion.
	Notes string `json:"notes,omitempty"`

	// CreateDate - The date a ssh key was added. This property is read only. Changes made will be silently
	// ignored.
	CreateDate *time.Time `json:"createDate,omitempty"`
}

func (softlayer_security_ssh_key *SoftLayer_Security_Ssh_Key) String() string {
	return "SoftLayer_Security_Ssh_Key"
}

// SoftLayer_Security_Ssh_Key_Extended is SoftLayer_Security_Ssh_Key with all maskable types.
type SoftLayer_Security_Ssh_Key_Extended struct {
	SoftLayer_Security_Ssh_Key

	// SoftwarePasswords - no documentation
	SoftwarePasswords []*SoftLayer_Software_Component_Password `json:"softwarePasswords,omitempty"`

	// BlockDeviceTemplateGroupCount - A count of the image template groups that are linked to an SSH key.
	BlockDeviceTemplateGroupCount uint64 `json:"blockDeviceTemplateGroupCount,omitempty"`

	// SoftwarePasswordCount - A count of the OS root users that are linked to an SSH key.
	SoftwarePasswordCount uint64 `json:"softwarePasswordCount,omitempty"`

	// Account - <nil>
	Account *SoftLayer_Account `json:"account,omitempty"`

	// BlockDeviceTemplateGroups - The image template groups that are linked to an SSH key.
	BlockDeviceTemplateGroups []*SoftLayer_Virtual_Guest_Block_Device_Template_Group `json:"blockDeviceTemplateGroups,omitempty"`
}

func (softlayer_security_ssh_key *SoftLayer_Security_Ssh_Key_Extended) String() string {
	return "SoftLayer_Security_Ssh_Key"
}
