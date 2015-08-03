package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Ticket_Update - The SoftLayer_Ticket_Update type relates to a single update to a ticket,
// either by a customer or an employee.
type SoftLayer_Ticket_Update struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// EditorId - The internal identifier of the SoftLayer portal or API user who created a ticket update.
	// This is only used if a ticket update's ''editorType'' property is
	EditorId int `json:"editorId"`

	// EditorType - The type user who created a ticket update. This is either for an update created by a
	// SoftLayer portal or API user, for an update created by a SoftLayer employee, or if a ticket update
	// was generated automatically by SoftLayer's backend systems.
	EditorType string `json:"editorType"`

	// Id - no documentation
	Id int `json:"id"`

	// TicketId - The internal identifier of the ticket that a ticket update belongs to.
	TicketId int `json:"ticketId"`

	// Entry - no documentation
	Entry string `json:"entry"`
}

// SoftLayer_Ticket_Update_Extended is SoftLayer_Ticket_Update with all maskable types.
type SoftLayer_Ticket_Update_Extended struct {
	SoftLayer_Ticket_Update

	// FileAttachmentCount - no documentation
	FileAttachmentCount uint64 `json:"fileAttachmentCount"`

	// Editor - The user or SoftLayer employee who created a ticket update.
	Editor *SoftLayer_User_Interface `json:"editor"`

	// Type - no documentation
	Type *SoftLayer_Ticket_Update_Type `json:"type"`

	// ChangeOwnerActivity - <nil>
	ChangeOwnerActivity string `json:"changeOwnerActivity"`

	// FileAttachment - no documentation
	FileAttachment []*SoftLayer_Ticket_Attachment_File `json:"fileAttachment"`

	// Ticket - no documentation
	Ticket *SoftLayer_Ticket `json:"ticket"`
}

func (softlayer_ticket_update *SoftLayer_Ticket_Update) String() string {
	return "SoftLayer_Ticket_Update"
}
