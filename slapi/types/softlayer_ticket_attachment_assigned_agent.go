package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Ticket_Attachment_Assigned_Agent - <nil>
type SoftLayer_Ticket_Attachment_Assigned_Agent struct {

	// AssignedAgentId - The internal identifier of an assigned Agent that is attached to a ticket.
	AssignedAgentId int `json:"assignedAgentId,omitempty"`

	// AttachmentId - The internal identifier of an item that is attached to a ticket.
	AttachmentId int `json:"attachmentId,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// TicketId - The internal identifier of the ticket that an item is attached to.
	TicketId int `json:"ticketId,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// AssignedAgent - <nil>
	AssignedAgent *SoftLayer_User_Customer `json:"assignedAgent,omitempty"`

	// ScheduledAction - <nil>
	ScheduledAction *SoftLayer_Provisioning_Version1_Transaction `json:"scheduledAction,omitempty"`

	// Ticket - no documentation
	Ticket *SoftLayer_Ticket `json:"ticket,omitempty"`

	// Resource - <nil>
	Resource *SoftLayer_User_Customer `json:"resource,omitempty"`
}

func (softlayer_ticket_attachment_assigned_agent *SoftLayer_Ticket_Attachment_Assigned_Agent) String() string {
	return "SoftLayer_Ticket_Attachment_Assigned_Agent"
}
