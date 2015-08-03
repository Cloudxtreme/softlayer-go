package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Storage_Credential - <nil>
type SoftLayer_Network_Storage_Credential struct {

	// AccountId - This is the account id associated with the volume.
	AccountId string `json:"accountId"`

	// ModifyDate - This is the date that the record was last updated in the table.
	ModifyDate *time.Time `json:"modifyDate"`

	// Username - no documentation
	Username string `json:"username"`

	// Password - no documentation
	Password string `json:"password"`

	// CreateDate - This is the data that the record was created in the table.
	CreateDate *time.Time `json:"createDate"`

	// Id - <nil>
	Id int `json:"id"`

	// NasCredentialTypeId - This is the id of the type of credential that this object represents.
	NasCredentialTypeId int `json:"nasCredentialTypeId"`
}

// SoftLayer_Network_Storage_Credential_Extended is SoftLayer_Network_Storage_Credential with all maskable types.
type SoftLayer_Network_Storage_Credential_Extended struct {
	SoftLayer_Network_Storage_Credential

	// NetworkStorageAllowedHosts - These are the SoftLayer_Network_Storage_Allowed_Host entries that this
	// credential is assigned to.
	NetworkStorageAllowedHosts *SoftLayer_Network_Storage_Allowed_Host `json:"networkStorageAllowedHosts"`

	// Type - These are the types of storage that the credential can be assigned to.
	Type *SoftLayer_Network_Storage_Credential_Type `json:"type"`

	// Account - This is the account that the storage credential is tied to.
	Account *SoftLayer_Account `json:"account"`

	// Volumes - These are the SoftLayer_Network_Storage volumes that this credential is assigned to.
	Volumes []*SoftLayer_Network_Storage `json:"volumes"`

	// VolumeCount - A count of these are the SoftLayer_Network_Storage volumes that this credential is
	// assigned to.
	VolumeCount uint64 `json:"volumeCount"`
}

func (softlayer_network_storage_credential *SoftLayer_Network_Storage_Credential) String() string {
	return "SoftLayer_Network_Storage_Credential"
}
