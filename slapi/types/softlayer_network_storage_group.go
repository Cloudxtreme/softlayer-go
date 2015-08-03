package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Storage_Group - <nil>
type SoftLayer_Network_Storage_Group struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Alias - no documentation
	Alias string `json:"alias"`

	// GroupTypeId - The SoftLayer_Network_Storage_Group_Type which describes this group.
	GroupTypeId int `json:"groupTypeId"`

	// Id - no documentation
	Id int `json:"id"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate"`

	// OsTypeId - A SoftLayer_Network_Storage_OS_Type Operating System designation that this group was
	// created for.
	OsTypeId int `json:"osTypeId"`

	// ServiceResourceId - A SoftLayer_Network_Service_Resource that this group was created on.
	ServiceResourceId int `json:"serviceResourceId"`

	// AccountId - no documentation
	AccountId int `json:"accountId"`
}

// SoftLayer_Network_Storage_Group_Extended is SoftLayer_Network_Storage_Group with all maskable types.
type SoftLayer_Network_Storage_Group_Extended struct {
	SoftLayer_Network_Storage_Group

	// AttachedVolumeCount - A count of the network storage volumes this group is attached to.
	AttachedVolumeCount uint64 `json:"attachedVolumeCount"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// AttachedVolumes - The network storage volumes this group is attached to.
	AttachedVolumes []*SoftLayer_Network_Storage `json:"attachedVolumes"`

	// OsType - no documentation
	OsType *SoftLayer_Network_Storage_Iscsi_OS_Type `json:"osType"`

	// GroupType - no documentation
	GroupType *SoftLayer_Network_Storage_Group_Type `json:"groupType"`

	// AllowedHostCount - no documentation
	AllowedHostCount uint64 `json:"allowedHostCount"`

	// AllowedHosts - no documentation
	AllowedHosts []*SoftLayer_Network_Storage_Allowed_Host `json:"allowedHosts"`

	// ServiceResource - no documentation
	ServiceResource *SoftLayer_Network_Service_Resource `json:"serviceResource"`
}

func (softlayer_network_storage_group *SoftLayer_Network_Storage_Group) String() string {
	return "SoftLayer_Network_Storage_Group"
}
