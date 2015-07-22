package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
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

// AddAllowedHost - Use this method to attach a SoftLayer_Network_Storage_Allowed_Host object to this
// group. This will automatically enable access from this host to any SoftLayer_Network_Storage volumes
// currently attached to this group.
func (softlayer_network_storage_group *SoftLayer_Network_Storage_Group) AddAllowedHost(ctx *slapi.RequestContext, allowedHost SoftLayer_Network_Storage_Allowed_Host) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// AttachToVolume - Use this method to attach a SoftLayer_Network_Storage volume to this group. This
// will automatically enable access to this volume for any SoftLayer_Network_Storage_Allowed_Host
// objects currently attached to this group.
func (softlayer_network_storage_group *SoftLayer_Network_Storage_Group) AttachToVolume(ctx *slapi.RequestContext, volume SoftLayer_Network_Storage) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// CreateObject - <nil>
func (softlayer_network_storage_group *SoftLayer_Network_Storage_Group) CreateObject(ctx *slapi.RequestContext, templateObject SoftLayer_Network_Storage_Group) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// DeleteObject - <nil>
func (softlayer_network_storage_group *SoftLayer_Network_Storage_Group) DeleteObject(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObject - <nil>
func (softlayer_network_storage_group *SoftLayer_Network_Storage_Group) EditObject(ctx *slapi.RequestContext, templateObject SoftLayer_Network_Storage_Group) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetAllObjects - Use this method to retrieve all network storage groups.
func (softlayer_network_storage_group *SoftLayer_Network_Storage_Group) GetAllObjects(ctx *slapi.RequestContext) ([]*SoftLayer_Network_Storage_Group, error) {
	var returnValue []*SoftLayer_Network_Storage_Group
	return returnValue, nil
}

// GetNetworkConnectionDetails - Use this method to retrieve network connection information for
// SoftLayer_Network_Storage_Allowed_Host objects within this group.
func (softlayer_network_storage_group *SoftLayer_Network_Storage_Group) GetNetworkConnectionDetails(ctx *slapi.RequestContext) (*SoftLayer_Container_Network_Storage_NetworkConnectionInformation, error) {
	var returnValue *SoftLayer_Container_Network_Storage_NetworkConnectionInformation
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_network_storage_group *SoftLayer_Network_Storage_Group) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Network_Storage_Group, error) {
	var returnValue *SoftLayer_Network_Storage_Group
	return returnValue, nil
}

// RemoveAllowedHost - Use this method to remove a SoftLayer_Network_Storage_Allowed_Host object from
// this group. This will automatically disable access from this host to any SoftLayer_Network_Storage
// volumes currently attached to this group.
func (softlayer_network_storage_group *SoftLayer_Network_Storage_Group) RemoveAllowedHost(ctx *slapi.RequestContext, allowedHost SoftLayer_Network_Storage_Allowed_Host) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RemoveFromVolume - Use this method to remove a SoftLayer_Network_Storage volume from this group.
// This will automatically disable access to this volume for any SoftLayer_Network_Storage_Allowed_Host
// objects currently attached to this group.
func (softlayer_network_storage_group *SoftLayer_Network_Storage_Group) RemoveFromVolume(ctx *slapi.RequestContext, volume SoftLayer_Network_Storage) (bool, error) {
	var returnValue bool
	return returnValue, nil
}
