package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Hardware_Component_RemoteManagement_Command - The
// SoftLayer_Network_Storage_Evault_Version6 contains the names of the remote management commands.
// Currently, only the reboot and power commands for the remote management card exist.
type SoftLayer_Hardware_Component_RemoteManagement_Command struct {

	// KeyName - no documentation
	KeyName string `json:"keyName"`
}

// SoftLayer_Hardware_Component_RemoteManagement_Command_Extended is SoftLayer_Hardware_Component_RemoteManagement_Command with all maskable types.
type SoftLayer_Hardware_Component_RemoteManagement_Command_Extended struct {
	SoftLayer_Hardware_Component_RemoteManagement_Command

	// Requests - All requests issued for the remote management command.
	Requests []*SoftLayer_Hardware_Component_RemoteManagement_Command_Request `json:"requests"`

	// RequestCount - A count of all requests issued for the remote management command.
	RequestCount uint64 `json:"requestCount"`
}

func (softlayer_hardware_component_remotemanagement_command *SoftLayer_Hardware_Component_RemoteManagement_Command) String() string {
	return "SoftLayer_Hardware_Component_RemoteManagement_Command"
}
