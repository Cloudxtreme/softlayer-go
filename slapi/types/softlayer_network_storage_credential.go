package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Storage_Credential - <nil>
type SoftLayer_Network_Storage_Credential struct {

	// Account - This is the account that the storage credential is tied to.
	Account *SoftLayer_Account `json:"account"`

	// AccountId - This is the account id associated with the volume.
	AccountId string `json:"accountId"`

	// CreateDate - This is the data that the record was created in the table.
	CreateDate *time.Time `json:"createDate"`

	// Id - <nil>
	Id int `json:"id"`

	// ModifyDate - This is the date that the record was last updated in the table.
	ModifyDate *time.Time `json:"modifyDate"`

	// NasCredentialTypeId - This is the id of the type of credential that this object represents.
	NasCredentialTypeId int `json:"nasCredentialTypeId"`

	// NetworkStorageAllowedHosts - These are the SoftLayer_Network_Storage_Allowed_Host entries that this
	// credential is assigned to.
	NetworkStorageAllowedHosts *SoftLayer_Network_Storage_Allowed_Host `json:"networkStorageAllowedHosts"`

	// Password - no documentation
	Password string `json:"password"`

	// Type - These are the types of storage that the credential can be assigned to.
	Type *SoftLayer_Network_Storage_Credential_Type `json:"type"`

	// Username - no documentation
	Username string `json:"username"`

	// VolumeCount - A count of these are the SoftLayer_Network_Storage volumes that this credential is
	// assigned to.
	VolumeCount uint64 `json:"volumeCount"`

	// Volumes - These are the SoftLayer_Network_Storage volumes that this credential is assigned to.
	Volumes []*SoftLayer_Network_Storage `json:"volumes"`
}
