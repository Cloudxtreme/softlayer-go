package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Component_Network_Vlan_Trunk - Represents the association between a
// Network_Component and Network_Vlan in the manner of a 'trunk'. Trunking a to a port allows that
// ports to receive and send packets tagged with the corresponding number.
type SoftLayer_Network_Component_Network_Vlan_Trunk struct {

	// NetworkComponentId - no documentation
	NetworkComponentId int `json:"networkComponentId"`

	// NetworkVlanId - The identifier of the network that is a trunk on the network component.
	NetworkVlanId int `json:"networkVlanId"`
}

// SoftLayer_Network_Component_Network_Vlan_Trunk_Extended is SoftLayer_Network_Component_Network_Vlan_Trunk with all maskable types.
type SoftLayer_Network_Component_Network_Vlan_Trunk_Extended struct {
	SoftLayer_Network_Component_Network_Vlan_Trunk

	// NetworkComponent - The network component that the is being trunked to.
	NetworkComponent *SoftLayer_Network_Component `json:"networkComponent"`

	// NetworkVlan - The that is being trunked to the network component.
	NetworkVlan *SoftLayer_Network_Vlan `json:"networkVlan"`
}

func (softlayer_network_component_network_vlan_trunk *SoftLayer_Network_Component_Network_Vlan_Trunk) String() string {
	return "SoftLayer_Network_Component_Network_Vlan_Trunk"
}
