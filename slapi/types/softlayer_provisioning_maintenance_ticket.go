package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Provisioning_Maintenance_Ticket - <nil>
type SoftLayer_Provisioning_Maintenance_Ticket struct {

	// MaintClassId - <nil>
	MaintClassId int `json:"maintClassId,omitempty"`

	// MaintWindowId - <nil>
	MaintWindowId int `json:"maintWindowId,omitempty"`

	// MaintenanceDate - <nil>
	MaintenanceDate *time.Time `json:"maintenanceDate,omitempty"`

	// TicketId - <nil>
	TicketId int `json:"ticketId,omitempty"`
}

func (softlayer_provisioning_maintenance_ticket *SoftLayer_Provisioning_Maintenance_Ticket) String() string {
	return "SoftLayer_Provisioning_Maintenance_Ticket"
}

// SoftLayer_Provisioning_Maintenance_Ticket_Extended is SoftLayer_Provisioning_Maintenance_Ticket with all maskable types.
type SoftLayer_Provisioning_Maintenance_Ticket_Extended struct {
	SoftLayer_Provisioning_Maintenance_Ticket

	// AvailableSlots - <nil>
	AvailableSlots *SoftLayer_Provisioning_Maintenance_Slots `json:"availableSlots,omitempty"`

	// MaintenanceClass - <nil>
	MaintenanceClass *SoftLayer_Provisioning_Maintenance_Classification `json:"maintenanceClass,omitempty"`

	// Ticket - <nil>
	Ticket *SoftLayer_Ticket `json:"ticket,omitempty"`
}

func (softlayer_provisioning_maintenance_ticket *SoftLayer_Provisioning_Maintenance_Ticket_Extended) String() string {
	return "SoftLayer_Provisioning_Maintenance_Ticket"
}
