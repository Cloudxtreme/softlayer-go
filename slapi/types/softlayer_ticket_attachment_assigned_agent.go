package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Ticket_Attachment_Assigned_Agent - <nil>
type SoftLayer_Ticket_Attachment_Assigned_Agent struct {

	// AssignedAgentId - The internal identifier of an assigned Agent that is attached to a ticket.
	AssignedAgentId int `json:"assignedAgentId"`

	// Resource - <nil>
	Resource *SoftLayer_User_Customer `json:"resource"`
}

func (softlayer_ticket_attachment_assigned_agent *SoftLayer_Ticket_Attachment_Assigned_Agent) String() string {
	return "SoftLayer_Ticket_Attachment_Assigned_Agent"
}
