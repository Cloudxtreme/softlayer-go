package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Hardware_Component_RemoteManagement_Command - The
// SoftLayer_Network_Storage_Evault_Version6 contains the names of the remote management commands.
// Currently, only the reboot and power commands for the remote management card exist.
type SoftLayer_Hardware_Component_RemoteManagement_Command struct {

	// KeyName - no documentation
	KeyName string `json:"keyName"`

	// RequestCount - A count of all requests issued for the remote management command.
	RequestCount uint64 `json:"requestCount"`

	// Requests - All requests issued for the remote management command.
	Requests []*SoftLayer_Hardware_Component_RemoteManagement_Command_Request `json:"requests"`
}
