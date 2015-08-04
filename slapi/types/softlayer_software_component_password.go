package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Software_Component_Password - This SoftLayer_Software_Component_Password data type
// contains a password for a specific software component instance.
type SoftLayer_Software_Component_Password struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Id - An id number for this specific username/password pair.
	Id int `json:"id,omitempty"`

	// Password - no documentation
	Password string `json:"password,omitempty"`

	// Port - The application access port for the Software Component.
	Port int `json:"port,omitempty"`

	// Username - no documentation
	Username string `json:"username,omitempty"`

	// ModifyDate - The date of the last modification to this username/password pair.
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// Notes - A note string stored for this username/password pair.
	Notes string `json:"notes,omitempty"`

	// SoftwareId - An id number for the software component this username/password pair is valid for.
	SoftwareId int `json:"softwareId,omitempty"`
}

func (softlayer_software_component_password *SoftLayer_Software_Component_Password) String() string {
	return "SoftLayer_Software_Component_Password"
}

// SoftLayer_Software_Component_Password_Extended is SoftLayer_Software_Component_Password with all maskable types.
type SoftLayer_Software_Component_Password_Extended struct {
	SoftLayer_Software_Component_Password

	// SshKeyCount - A count of sSH keys to be installed on the server during provisioning or an OS reload.
	SshKeyCount uint64 `json:"sshKeyCount,omitempty"`

	// Software - The SoftLayer_Software_Component instance that this username/password pair is valid for.
	Software *SoftLayer_Software_Component `json:"software,omitempty"`

	// SshKeys - SSH keys to be installed on the server during provisioning or an OS reload.
	SshKeys []*SoftLayer_Security_Ssh_Key `json:"sshKeys,omitempty"`
}

func (softlayer_software_component_password *SoftLayer_Software_Component_Password_Extended) String() string {
	return "SoftLayer_Software_Component_Password"
}
