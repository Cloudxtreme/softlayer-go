package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Provisioning_Maintenance_Slots - The SoftLayer_Provisioning_Maintenance_Slots represent
// the available slots for a given maintenance window at a SoftLayer data center.
type SoftLayer_Provisioning_Maintenance_Slots struct {

	// AvailableSlots - no documentation
	AvailableSlots int `json:"availableSlots"`
}

func (softlayer_provisioning_maintenance_slots *SoftLayer_Provisioning_Maintenance_Slots) String() string {
	return "SoftLayer_Provisioning_Maintenance_Slots"
}

// GetObject - <nil>
func (softlayer_provisioning_maintenance_slots *SoftLayer_Provisioning_Maintenance_Slots) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Provisioning_Maintenance_Slots, error) {
	var returnValue *SoftLayer_Provisioning_Maintenance_Slots
	return returnValue, nil
}
