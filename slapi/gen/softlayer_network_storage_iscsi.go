package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Storage_Iscsi - The iscsi data type provides access to additional information
// about an iscsi volume such as the snapshot capacity limit and replication partners.
type SoftLayer_Network_Storage_Iscsi struct {
}

// AllowAccessFromHardware - This method is used to modify the access control list for this Storage
// volume. The SoftLayer_Hardware objects which have been allowed access to this storage will be listed
// in the allowedHardware property of this storage volume.
func (softlayer_network_storage_iscsi *SoftLayer_Network_Storage_Iscsi) AllowAccessFromHardware(hardwareObjectTemplate SoftLayer_Hardware) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// AllowAccessFromIpAddress - <nil>
func (softlayer_network_storage_iscsi *SoftLayer_Network_Storage_Iscsi) AllowAccessFromIpAddress(ipAddressObjectTemplate SoftLayer_Network_Subnet_IpAddress) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// AllowAccessFromVirtualGuest - This method is used to modify the access control list for this Storage
// volume. The SoftLayer_Virtual_Guest objects which have been allowed access to this storage will be
// listed in the allowedVirtualGuests property of this storage volume.
func (softlayer_network_storage_iscsi *SoftLayer_Network_Storage_Iscsi) AllowAccessFromVirtualGuest(virtualGuestObjectTemplate SoftLayer_Virtual_Guest) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// AllowAccessToReplicantFromHardwareList - This method is used to modify the access control list for
// this Storage replica volume. The SoftLayer_Hardware objects which have been allowed access to this
// storage will be listed in the allowedHardware property of this storage replica volume.
func (softlayer_network_storage_iscsi *SoftLayer_Network_Storage_Iscsi) AllowAccessToReplicantFromHardwareList(hardwareObjectTemplates []SoftLayer_Hardware) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// AllowAccessToReplicantFromIpAddressList - This method is used to modify the access control list for
// this Storage volume. The SoftLayer_Network_Subnet_IpAddress objects which have been allowed access
// to this storage will be listed in the allowedVirtualGuests property of this storage volume.
func (softlayer_network_storage_iscsi *SoftLayer_Network_Storage_Iscsi) AllowAccessToReplicantFromIpAddressList(ipAddressObjectTemplates []SoftLayer_Network_Subnet_IpAddress) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// AllowAccessToReplicantFromVirtualGuestList - This method is used to modify the access control list
// for this Storage volume. The SoftLayer_Virtual_Guest objects which have been allowed access to this
// storage will be listed in the allowedVirtualGuests property of this storage volume.
func (softlayer_network_storage_iscsi *SoftLayer_Network_Storage_Iscsi) AllowAccessToReplicantFromVirtualGuestList(virtualGuestObjectTemplates []SoftLayer_Virtual_Guest) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_network_storage_iscsi *SoftLayer_Network_Storage_Iscsi) GetObject() (*SoftLayer_Network_Storage_Iscsi, error) {
	var returnValue *SoftLayer_Network_Storage_Iscsi
	return returnValue, nil
}

// GetSnapshotsForVolume - Retrieves a list of snapshots for this SoftLayer_Network_Storage volume.
// This method works with the result limits and offset to support pagination.
func (softlayer_network_storage_iscsi *SoftLayer_Network_Storage_Iscsi) GetSnapshotsForVolume() ([]*SoftLayer_Network_Storage, error) {
	var returnValue []*SoftLayer_Network_Storage
	return returnValue, nil
}

// RemoveAccessFromHardware - This method is used to modify the access control list for this Storage
// volume. The SoftLayer_Hardware objects which have been allowed access to this storage will be listed
// in the allowedHardware property of this storage volume.
func (softlayer_network_storage_iscsi *SoftLayer_Network_Storage_Iscsi) RemoveAccessFromHardware(hardwareObjectTemplate SoftLayer_Hardware) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RemoveAccessFromIpAddress - <nil>
func (softlayer_network_storage_iscsi *SoftLayer_Network_Storage_Iscsi) RemoveAccessFromIpAddress(ipAddressObjectTemplate SoftLayer_Network_Subnet_IpAddress) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RemoveAccessFromVirtualGuest - This method is used to modify the access control list for this
// Storage volume. The SoftLayer_Virtual_Guest objects which have been allowed access to this storage
// will be listed in the allowedVirtualGuests property of this storage volume.
func (softlayer_network_storage_iscsi *SoftLayer_Network_Storage_Iscsi) RemoveAccessFromVirtualGuest(virtualGuestObjectTemplate SoftLayer_Virtual_Guest) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RemoveAccessToReplicantFromHardwareList - This method is used to modify the access control list for
// this Storage replica volume. The SoftLayer_Hardware objects which have been allowed access to this
// storage will be listed in the allowedHardware property of this storage replica volume.
func (softlayer_network_storage_iscsi *SoftLayer_Network_Storage_Iscsi) RemoveAccessToReplicantFromHardwareList(hardwareObjectTemplates []SoftLayer_Hardware) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RemoveAccessToReplicantFromIpAddressList - This method is used to modify the access control list for
// this Storage replica volume. The SoftLayer_Network_Subnet_IpAddress objects which have been allowed
// access to this storage will be listed in the allowedIpAddresses property of this storage replica
// volume.
func (softlayer_network_storage_iscsi *SoftLayer_Network_Storage_Iscsi) RemoveAccessToReplicantFromIpAddressList(ipAddressObjectTemplates []SoftLayer_Network_Subnet_IpAddress) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RemoveAccessToReplicantFromVirtualGuestList - This method is used to modify the access control list
// for this Storage replica volume. The SoftLayer_Virtual_Guest objects which have been allowed access
// to this storage will be listed in the allowedVirtualGuests property of this storage replica volume.
func (softlayer_network_storage_iscsi *SoftLayer_Network_Storage_Iscsi) RemoveAccessToReplicantFromVirtualGuestList(virtualGuestObjectTemplates []SoftLayer_Virtual_Guest) (bool, error) {
	var returnValue bool
	return returnValue, nil
}
