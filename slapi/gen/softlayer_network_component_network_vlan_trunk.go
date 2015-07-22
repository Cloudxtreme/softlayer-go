package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Component_Network_Vlan_Trunk - Represents the association between a
// Network_Component and Network_Vlan in the manner of a 'trunk'. Trunking a to a port allows that
// ports to receive and send packets tagged with the corresponding number.
type SoftLayer_Network_Component_Network_Vlan_Trunk struct {

	// NetworkComponent - The network component that the is being trunked to.
	NetworkComponent *SoftLayer_Network_Component `json:"networkComponent"`

	// NetworkComponentId - no documentation
	NetworkComponentId int `json:"networkComponentId"`

	// NetworkVlan - The that is being trunked to the network component.
	NetworkVlan *SoftLayer_Network_Vlan `json:"networkVlan"`

	// NetworkVlanId - The identifier of the network that is a trunk on the network component.
	NetworkVlanId int `json:"networkVlanId"`
}
