package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Hardware_Server_Details - The SoftLayer_Container_Hardware_Server_Details data
// type contains information relating to a server's component information, network information, and
// software information.
type SoftLayer_Container_Hardware_Server_Details struct {

	// Components - no documentation
	Components []*SoftLayer_Hardware_Component `json:"components"`

	// NetworkComponents - The network components that belong to a piece of hardware.
	NetworkComponents []*SoftLayer_Network_Component `json:"networkComponents"`

	// Software - no documentation
	Software []*SoftLayer_Software_Component `json:"software"`
}
