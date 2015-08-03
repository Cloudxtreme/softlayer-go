package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Storage_History - The SoftLayer_Network_Storage_History contains the
// username/password past history for Storage services except Evault. Information such as the username,
// passwords, notes and the date of the password change may be retrieved.
type SoftLayer_Network_Storage_History struct {

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// NasVolume - The Storage service that the password history belongs to.
	NasVolume *SoftLayer_Network_Storage `json:"nasVolume"`

	// Notes - no documentation
	Notes string `json:"notes"`

	// Password - Password for the Storage service that was used in the past.
	Password string `json:"password"`

	// Username - no documentation
	Username string `json:"username"`
}

func (softlayer_network_storage_history *SoftLayer_Network_Storage_History) String() string {
	return "SoftLayer_Network_Storage_History"
}
