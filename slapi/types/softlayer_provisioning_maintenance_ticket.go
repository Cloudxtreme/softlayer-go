package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Provisioning_Maintenance_Ticket - <nil>
type SoftLayer_Provisioning_Maintenance_Ticket struct {

	// AvailableSlots - <nil>
	AvailableSlots *SoftLayer_Provisioning_Maintenance_Slots `json:"availableSlots"`

	// MaintClassId - <nil>
	MaintClassId int `json:"maintClassId"`

	// MaintWindowId - <nil>
	MaintWindowId int `json:"maintWindowId"`

	// MaintenanceClass - <nil>
	MaintenanceClass *SoftLayer_Provisioning_Maintenance_Classification `json:"maintenanceClass"`

	// MaintenanceDate - <nil>
	MaintenanceDate *time.Time `json:"maintenanceDate"`

	// Ticket - <nil>
	Ticket *SoftLayer_Ticket `json:"ticket"`

	// TicketId - <nil>
	TicketId int `json:"ticketId"`
}

func (softlayer_provisioning_maintenance_ticket *SoftLayer_Provisioning_Maintenance_Ticket) String() string {
	return "SoftLayer_Provisioning_Maintenance_Ticket"
}

// GetObject - <nil>
func (softlayer_provisioning_maintenance_ticket *SoftLayer_Provisioning_Maintenance_Ticket) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Provisioning_Maintenance_Ticket, error) {
	var returnValue *SoftLayer_Provisioning_Maintenance_Ticket
	return returnValue, nil
}
