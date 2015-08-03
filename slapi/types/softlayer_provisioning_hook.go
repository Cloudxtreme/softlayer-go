package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Provisioning_Hook - The SoftLayer_Provisioning_Hook contains all the information needed to
// add a hook into a server/Virtual provision and os reload.
type SoftLayer_Provisioning_Hook struct {

	// AccountId - no documentation
	AccountId int `json:"accountId"`

	// Name - no documentation
	Name string `json:"name"`

	// Uri - The endpoint that the script will be downloaded from AND BE If the endpoint is the script will
	// only be downloaded. If the endpoint is the script will be downloaded and executed.
	Uri string `json:"uri"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// Id - <nil>
	Id int `json:"id"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate"`

	// TypeId - The ID of the type of hook the script is identified as. Currently only has been
	// implemented.
	TypeId int `json:"typeId"`
}

func (softlayer_provisioning_hook *SoftLayer_Provisioning_Hook) String() string {
	return "SoftLayer_Provisioning_Hook"
}

// SoftLayer_Provisioning_Hook_Extended is SoftLayer_Provisioning_Hook with all maskable types.
type SoftLayer_Provisioning_Hook_Extended struct {
	SoftLayer_Provisioning_Hook

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`

	// HookType - <nil>
	HookType *SoftLayer_Provisioning_Hook_Type `json:"hookType"`
}

func (softlayer_provisioning_hook *SoftLayer_Provisioning_Hook_Extended) String() string {
	return "SoftLayer_Provisioning_Hook"
}
