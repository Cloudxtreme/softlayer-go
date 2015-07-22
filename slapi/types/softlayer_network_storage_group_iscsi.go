package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Network_Storage_Group_Iscsi - <nil>
type SoftLayer_Network_Storage_Group_Iscsi struct {
}

// AddAllowedHost - Use this method to attach a SoftLayer_Network_Storage_Allowed_Host object to this
// group. This will automatically enable access from this host to any SoftLayer_Network_Storage volumes
// currently attached to this group.
func (softlayer_network_storage_group_iscsi *SoftLayer_Network_Storage_Group_Iscsi) AddAllowedHost(ctx *slapi.RequestContext, allowedHost SoftLayer_Network_Storage_Allowed_Host) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// AttachToVolume - Use this method to attach a SoftLayer_Network_Storage volume to this group. This
// will automatically enable access to this volume for any SoftLayer_Network_Storage_Allowed_Host
// objects currently attached to this group.
func (softlayer_network_storage_group_iscsi *SoftLayer_Network_Storage_Group_Iscsi) AttachToVolume(ctx *slapi.RequestContext, volume SoftLayer_Network_Storage) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_network_storage_group_iscsi *SoftLayer_Network_Storage_Group_Iscsi) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Network_Storage_Group_Iscsi, error) {
	var returnValue *SoftLayer_Network_Storage_Group_Iscsi
	return returnValue, nil
}

// RemoveAllowedHost - Use this method to remove a SoftLayer_Network_Storage_Allowed_Host object from
// this group. This will automatically disable access from this host to any SoftLayer_Network_Storage
// volumes currently attached to this group.
func (softlayer_network_storage_group_iscsi *SoftLayer_Network_Storage_Group_Iscsi) RemoveAllowedHost(ctx *slapi.RequestContext, allowedHost SoftLayer_Network_Storage_Allowed_Host) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RemoveFromVolume - Use this method to remove a SoftLayer_Network_Storage volume from this group.
// This will automatically disable access to this volume for any SoftLayer_Network_Storage_Allowed_Host
// objects currently attached to this group.
func (softlayer_network_storage_group_iscsi *SoftLayer_Network_Storage_Group_Iscsi) RemoveFromVolume(ctx *slapi.RequestContext, volume SoftLayer_Network_Storage) (bool, error) {
	var returnValue bool
	return returnValue, nil
}
