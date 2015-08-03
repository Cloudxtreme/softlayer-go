package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Component_Uplink_Hardware - The SoftLayer_Network_Component_Uplink_Hardware data
// type abstracts information related to network connections between SoftLayer hardware and SoftLayer
// network components. It is populated via triggers on the network_connection table
// (SoftLayer_Network_Connection), so you shouldn't have to delete or insert records into this table,
// ever.
type SoftLayer_Network_Component_Uplink_Hardware struct {

	// Hardware - A network component uplink's connected [[SoftLayer_Hardware|Hardware]].
	Hardware *SoftLayer_Hardware `json:"hardware"`

	// NetworkComponent - The [[SoftLayer_Network_Component|Network Component]] that a uplink connection
	// belongs to..
	NetworkComponent *SoftLayer_Network_Component `json:"networkComponent"`
}

func (softlayer_network_component_uplink_hardware *SoftLayer_Network_Component_Uplink_Hardware) String() string {
	return "SoftLayer_Network_Component_Uplink_Hardware"
}
