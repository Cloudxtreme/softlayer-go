package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Storage_Allowed_Host_VirtualGuest - <nil>
type SoftLayer_Network_Storage_Allowed_Host_VirtualGuest struct {

	// ResourceTableName - <nil>
	ResourceTableName string `json:"resourceTableName,omitempty"`

	// Name - The name of allowed host, usually an IQN or other identifier
	Name string `json:"name,omitempty"`

	// ResourceTableId - <nil>
	ResourceTableId int `json:"resourceTableId,omitempty"`

	// CredentialId - no documentation
	CredentialId int `json:"credentialId,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`
}

func (softlayer_network_storage_allowed_host_virtualguest *SoftLayer_Network_Storage_Allowed_Host_VirtualGuest) String() string {
	return "SoftLayer_Network_Storage_Allowed_Host_VirtualGuest"
}

// SoftLayer_Network_Storage_Allowed_Host_VirtualGuest_Extended is SoftLayer_Network_Storage_Allowed_Host_VirtualGuest with all maskable types.
type SoftLayer_Network_Storage_Allowed_Host_VirtualGuest_Extended struct {
	SoftLayer_Network_Storage_Allowed_Host_VirtualGuest

	// AssignedReplicationVolumes - The SoftLayer_Network_Storage primary volumes whose replicas are
	// allowed access.
	AssignedReplicationVolumes []*SoftLayer_Network_Storage `json:"assignedReplicationVolumes,omitempty"`

	// AssignedVolumes - The SoftLayer_Network_Storage volumes to which this
	// SoftLayer_Network_Storage_Allowed_Host is allowed access.
	AssignedVolumes []*SoftLayer_Network_Storage `json:"assignedVolumes,omitempty"`

	// AssignedGroupCount - A count of the SoftLayer_Network_Storage_Group objects this
	// SoftLayer_Network_Storage_Allowed_Host is present in.
	AssignedGroupCount uint64 `json:"assignedGroupCount,omitempty"`

	// AssignedGroups - The SoftLayer_Network_Storage_Group objects this
	// SoftLayer_Network_Storage_Allowed_Host is present in.
	AssignedGroups []*SoftLayer_Network_Storage_Group `json:"assignedGroups,omitempty"`

	// Credential - The SoftLayer_Network_Storage_Credential this allowed host uses.
	Credential *SoftLayer_Network_Storage_Credential `json:"credential,omitempty"`

	// Resource - The SoftLayer_Virtual_Guest object which this SoftLayer_Network_Storage_Allowed_Host is
	// referencing.
	Resource *SoftLayer_Virtual_Guest `json:"resource,omitempty"`

	// AssignedVolumeCount - A count of the SoftLayer_Network_Storage volumes to which this
	// SoftLayer_Network_Storage_Allowed_Host is allowed access.
	AssignedVolumeCount uint64 `json:"assignedVolumeCount,omitempty"`

	// AssignedReplicationVolumeCount - A count of the SoftLayer_Network_Storage primary volumes whose
	// replicas are allowed access.
	AssignedReplicationVolumeCount uint64 `json:"assignedReplicationVolumeCount,omitempty"`
}

func (softlayer_network_storage_allowed_host_virtualguest *SoftLayer_Network_Storage_Allowed_Host_VirtualGuest_Extended) String() string {
	return "SoftLayer_Network_Storage_Allowed_Host_VirtualGuest"
}
