package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Component_Group - <nil>
type SoftLayer_Network_Component_Group struct {

	// GroupTypeId - <nil>
	GroupTypeId int `json:"groupTypeId,omitempty"`
}

func (softlayer_network_component_group *SoftLayer_Network_Component_Group) String() string {
	return "SoftLayer_Network_Component_Group"
}

// SoftLayer_Network_Component_Group_Extended is SoftLayer_Network_Component_Group with all maskable types.
type SoftLayer_Network_Component_Group_Extended struct {
	SoftLayer_Network_Component_Group

	// NetworkComponents - A network component group's associated network components.
	NetworkComponents []*SoftLayer_Network_Component `json:"networkComponents,omitempty"`

	// NetworkComponentCount - A count of a network component group's associated network components.
	NetworkComponentCount uint64 `json:"networkComponentCount,omitempty"`
}

func (softlayer_network_component_group *SoftLayer_Network_Component_Group_Extended) String() string {
	return "SoftLayer_Network_Component_Group"
}
