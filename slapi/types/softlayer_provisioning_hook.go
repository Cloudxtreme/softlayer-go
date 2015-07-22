package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
)

// SoftLayer_Provisioning_Hook - The SoftLayer_Provisioning_Hook contains all the information needed to
// add a hook into a server/Virtual provision and os reload.
type SoftLayer_Provisioning_Hook struct {

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`

	// AccountId - no documentation
	AccountId int `json:"accountId"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// HookType - <nil>
	HookType *SoftLayer_Provisioning_Hook_Type `json:"hookType"`

	// Id - <nil>
	Id int `json:"id"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate"`

	// Name - no documentation
	Name string `json:"name"`

	// TypeId - The ID of the type of hook the script is identified as. Currently only has been
	// implemented.
	TypeId int `json:"typeId"`

	// Uri - The endpoint that the script will be downloaded from AND BE If the endpoint is the script will
	// only be downloaded. If the endpoint is the script will be downloaded and executed.
	Uri string `json:"uri"`
}

// CreateObject - <nil>
func (softlayer_provisioning_hook *SoftLayer_Provisioning_Hook) CreateObject(commonOptions *slapi.CommonOptions, templateObject SoftLayer_Provisioning_Hook) (*SoftLayer_Provisioning_Hook, error) {
	var returnValue *SoftLayer_Provisioning_Hook
	return returnValue, nil
}

// DeleteObject - <nil>
func (softlayer_provisioning_hook *SoftLayer_Provisioning_Hook) DeleteObject(commonOptions *slapi.CommonOptions) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObject - <nil>
func (softlayer_provisioning_hook *SoftLayer_Provisioning_Hook) EditObject(commonOptions *slapi.CommonOptions, templateObject SoftLayer_Provisioning_Hook) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_provisioning_hook *SoftLayer_Provisioning_Hook) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Provisioning_Hook, error) {
	var returnValue *SoftLayer_Provisioning_Hook
	return returnValue, nil
}