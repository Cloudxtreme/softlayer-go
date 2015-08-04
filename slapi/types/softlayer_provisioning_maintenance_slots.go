package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Provisioning_Maintenance_Slots - The SoftLayer_Provisioning_Maintenance_Slots represent
// the available slots for a given maintenance window at a SoftLayer data center.
type SoftLayer_Provisioning_Maintenance_Slots struct {

	// AvailableSlots - no documentation
	AvailableSlots int `json:"availableSlots,omitempty"`
}

func (softlayer_provisioning_maintenance_slots *SoftLayer_Provisioning_Maintenance_Slots) String() string {
	return "SoftLayer_Provisioning_Maintenance_Slots"
}
