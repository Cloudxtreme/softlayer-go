package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Network_IntrusionProtection_Statistic - The IntrusionProtection_Statistic is
// used exclusively by the getMainStatistics method on the TippingPointReporting service, and serves
// mainly as a pair object, storing a name and an attack count. Name is usually the name of an attack,
// but it can also be an attacking IP Address
type SoftLayer_Container_Network_IntrusionProtection_Statistic struct {

	// AttackCount - The number of attacks effecting this name over the time period
	AttackCount int `json:"attackCount"`

	// Name - Either the name of the attack in question, or the attacking IP Address
	Name string `json:"name"`
}

func (softlayer_container_network_intrusionprotection_statistic *SoftLayer_Container_Network_IntrusionProtection_Statistic) String() string {
	return "SoftLayer_Container_Network_IntrusionProtection_Statistic"
}
