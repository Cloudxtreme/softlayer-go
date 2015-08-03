package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Hardware_Component_RemoteManagement - This class adds functionality to the base
// SoftLayer_Hardware class for web servers (all server hardware)
type SoftLayer_Hardware_Component_RemoteManagement struct {

	// NetworkComponent - no documentation
	NetworkComponent *SoftLayer_Network_Component `json:"networkComponent"`
}

func (softlayer_hardware_component_remotemanagement *SoftLayer_Hardware_Component_RemoteManagement) String() string {
	return "SoftLayer_Hardware_Component_RemoteManagement"
}
