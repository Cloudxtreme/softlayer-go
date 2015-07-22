package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Storage_Allowed_Host - <nil>
type SoftLayer_Network_Storage_Allowed_Host struct {

	// AssignedGroupCount - A count of the SoftLayer_Network_Storage_Group objects this
	// SoftLayer_Network_Storage_Allowed_Host is present in.
	AssignedGroupCount uint64 `json:"assignedGroupCount"`

	// AssignedGroups - The SoftLayer_Network_Storage_Group objects this
	// SoftLayer_Network_Storage_Allowed_Host is present in.
	AssignedGroups []*SoftLayer_Network_Storage_Group `json:"assignedGroups"`

	// AssignedReplicationVolumeCount - A count of the SoftLayer_Network_Storage primary volumes whose
	// replicas are allowed access.
	AssignedReplicationVolumeCount uint64 `json:"assignedReplicationVolumeCount"`

	// AssignedReplicationVolumes - The SoftLayer_Network_Storage primary volumes whose replicas are
	// allowed access.
	AssignedReplicationVolumes []*SoftLayer_Network_Storage `json:"assignedReplicationVolumes"`

	// AssignedVolumeCount - A count of the SoftLayer_Network_Storage volumes to which this
	// SoftLayer_Network_Storage_Allowed_Host is allowed access.
	AssignedVolumeCount uint64 `json:"assignedVolumeCount"`

	// AssignedVolumes - The SoftLayer_Network_Storage volumes to which this
	// SoftLayer_Network_Storage_Allowed_Host is allowed access.
	AssignedVolumes []*SoftLayer_Network_Storage `json:"assignedVolumes"`

	// Credential - The SoftLayer_Network_Storage_Credential this allowed host uses.
	Credential *SoftLayer_Network_Storage_Credential `json:"credential"`

	// CredentialId - no documentation
	CredentialId int `json:"credentialId"`

	// Id - no documentation
	Id int `json:"id"`

	// Name - The name of allowed host, usually an IQN or other identifier
	Name string `json:"name"`

	// ResourceTableId - <nil>
	ResourceTableId int `json:"resourceTableId"`

	// ResourceTableName - <nil>
	ResourceTableName string `json:"resourceTableName"`
}

// CreateObject - <nil>
func (softlayer_network_storage_allowed_host *SoftLayer_Network_Storage_Allowed_Host) CreateObject(templateObject SoftLayer_Network_Storage_Allowed_Host) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// DeleteObject - <nil>
func (softlayer_network_storage_allowed_host *SoftLayer_Network_Storage_Allowed_Host) DeleteObject() (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObject - <nil>
func (softlayer_network_storage_allowed_host *SoftLayer_Network_Storage_Allowed_Host) EditObject(templateObject SoftLayer_Network_Storage_Allowed_Host) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_network_storage_allowed_host *SoftLayer_Network_Storage_Allowed_Host) GetObject() (*SoftLayer_Network_Storage_Allowed_Host, error) {
	var returnValue *SoftLayer_Network_Storage_Allowed_Host
	return returnValue, nil
}

// SetCredentialPassword - Use this method to modify the credential password for a
// SoftLayer_Network_Storage_Allowed_Host object.
func (softlayer_network_storage_allowed_host *SoftLayer_Network_Storage_Allowed_Host) SetCredentialPassword(password string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}
