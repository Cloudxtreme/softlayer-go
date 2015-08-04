package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Storage_History - The SoftLayer_Network_Storage_History contains the
// username/password past history for Storage services except Evault. Information such as the username,
// passwords, notes and the date of the password change may be retrieved.
type SoftLayer_Network_Storage_History struct {

	// Notes - no documentation
	Notes string `json:"notes,omitempty"`

	// Password - Password for the Storage service that was used in the past.
	Password string `json:"password,omitempty"`

	// Username - no documentation
	Username string `json:"username,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`
}

func (softlayer_network_storage_history *SoftLayer_Network_Storage_History) String() string {
	return "SoftLayer_Network_Storage_History"
}

// SoftLayer_Network_Storage_History_Extended is SoftLayer_Network_Storage_History with all maskable types.
type SoftLayer_Network_Storage_History_Extended struct {
	SoftLayer_Network_Storage_History

	// Account - no documentation
	Account *SoftLayer_Account `json:"account,omitempty"`

	// NasVolume - The Storage service that the password history belongs to.
	NasVolume *SoftLayer_Network_Storage `json:"nasVolume,omitempty"`
}

func (softlayer_network_storage_history *SoftLayer_Network_Storage_History_Extended) String() string {
	return "SoftLayer_Network_Storage_History"
}
