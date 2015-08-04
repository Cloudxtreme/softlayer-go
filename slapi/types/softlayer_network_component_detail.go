package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Component_Detail - <nil>
type SoftLayer_Network_Component_Detail struct {

	// PollingInterfaceIndex - <nil>
	PollingInterfaceIndex int `json:"pollingInterfaceIndex,omitempty"`
}

func (softlayer_network_component_detail *SoftLayer_Network_Component_Detail) String() string {
	return "SoftLayer_Network_Component_Detail"
}

// SoftLayer_Network_Component_Detail_Extended is SoftLayer_Network_Component_Detail with all maskable types.
type SoftLayer_Network_Component_Detail_Extended struct {
	SoftLayer_Network_Component_Detail

	// NetworkComponent - <nil>
	NetworkComponent *SoftLayer_Network_Component `json:"networkComponent,omitempty"`
}

func (softlayer_network_component_detail *SoftLayer_Network_Component_Detail_Extended) String() string {
	return "SoftLayer_Network_Component_Detail"
}
