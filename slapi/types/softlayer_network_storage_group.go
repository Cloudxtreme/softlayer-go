package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Storage_Group - <nil>
type SoftLayer_Network_Storage_Group struct {

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// AccountId - no documentation
	AccountId int `json:"accountId"`

	// Alias - no documentation
	Alias string `json:"alias"`

	// AllowedHostCount - no documentation
	AllowedHostCount uint64 `json:"allowedHostCount"`

	// AllowedHosts - no documentation
	AllowedHosts []*SoftLayer_Network_Storage_Allowed_Host `json:"allowedHosts"`

	// AttachedVolumeCount - A count of the network storage volumes this group is attached to.
	AttachedVolumeCount uint64 `json:"attachedVolumeCount"`

	// AttachedVolumes - The network storage volumes this group is attached to.
	AttachedVolumes []*SoftLayer_Network_Storage `json:"attachedVolumes"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// GroupType - no documentation
	GroupType *SoftLayer_Network_Storage_Group_Type `json:"groupType"`

	// GroupTypeId - The SoftLayer_Network_Storage_Group_Type which describes this group.
	GroupTypeId int `json:"groupTypeId"`

	// Id - no documentation
	Id int `json:"id"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate"`

	// OsType - no documentation
	OsType *SoftLayer_Network_Storage_Iscsi_OS_Type `json:"osType"`

	// OsTypeId - A SoftLayer_Network_Storage_OS_Type Operating System designation that this group was
	// created for.
	OsTypeId int `json:"osTypeId"`

	// ServiceResource - no documentation
	ServiceResource *SoftLayer_Network_Service_Resource `json:"serviceResource"`

	// ServiceResourceId - A SoftLayer_Network_Service_Resource that this group was created on.
	ServiceResourceId int `json:"serviceResourceId"`
}

func (softlayer_network_storage_group *SoftLayer_Network_Storage_Group) String() string {
	return "SoftLayer_Network_Storage_Group"
}
