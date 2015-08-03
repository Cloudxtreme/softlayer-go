package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Software_Component_Password - This SoftLayer_Software_Component_Password data type
// contains a password for a specific software component instance.
type SoftLayer_Software_Component_Password struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Id - An id number for this specific username/password pair.
	Id int `json:"id"`

	// ModifyDate - The date of the last modification to this username/password pair.
	ModifyDate *time.Time `json:"modifyDate"`

	// Notes - A note string stored for this username/password pair.
	Notes string `json:"notes"`

	// Password - no documentation
	Password string `json:"password"`

	// Port - The application access port for the Software Component.
	Port int `json:"port"`

	// Software - The SoftLayer_Software_Component instance that this username/password pair is valid for.
	Software *SoftLayer_Software_Component `json:"software"`

	// SoftwareId - An id number for the software component this username/password pair is valid for.
	SoftwareId int `json:"softwareId"`

	// SshKeyCount - A count of sSH keys to be installed on the server during provisioning or an OS reload.
	SshKeyCount uint64 `json:"sshKeyCount"`

	// SshKeys - SSH keys to be installed on the server during provisioning or an OS reload.
	SshKeys []*SoftLayer_Security_Ssh_Key `json:"sshKeys"`

	// Username - no documentation
	Username string `json:"username"`
}

func (softlayer_software_component_password *SoftLayer_Software_Component_Password) String() string {
	return "SoftLayer_Software_Component_Password"
}
