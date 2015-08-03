package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Hardware_Component_RemoteManagement - This class adds functionality to the base
// SoftLayer_Hardware class for web servers (all server hardware)
type SoftLayer_Hardware_Component_RemoteManagement struct {
}

func (softlayer_hardware_component_remotemanagement *SoftLayer_Hardware_Component_RemoteManagement) String() string {
	return "SoftLayer_Hardware_Component_RemoteManagement"
}

// SoftLayer_Hardware_Component_RemoteManagement_Extended is SoftLayer_Hardware_Component_RemoteManagement with all maskable types.
type SoftLayer_Hardware_Component_RemoteManagement_Extended struct {
	SoftLayer_Hardware_Component_RemoteManagement

	// NetworkComponent - no documentation
	NetworkComponent *SoftLayer_Network_Component `json:"networkComponent"`
}

func (softlayer_hardware_component_remotemanagement *SoftLayer_Hardware_Component_RemoteManagement_Extended) String() string {
	return "SoftLayer_Hardware_Component_RemoteManagement"
}
