package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Ticket_Update_Employee - The SoftLayer_Ticket_Update_Employee data type models an update
// to a ticket made by a SoftLayer employee.
type SoftLayer_Ticket_Update_Employee struct {

	// Entry - no documentation
	Entry string `json:"entry,omitempty"`

	// TicketId - The internal identifier of the ticket that a ticket update belongs to.
	TicketId int `json:"ticketId,omitempty"`

	// EditorId - The internal identifier of the SoftLayer portal or API user who created a ticket update.
	// This is only used if a ticket update's ''editorType'' property is
	EditorId int `json:"editorId,omitempty"`

	// ResponseRating - A ticket update's response rating. Ticket updates posted by SoftLayer employees
	// have the option of earning a rating from SoftLayer's customers. Ratings are based on a 1 - 5 scale,
	// with one being a poor rating while 5 is a very high rating. This is only used if a ticket update's
	// ''editorType'' property is
	ResponseRating int `json:"responseRating,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// EditorType - The type user who created a ticket update. This is either for an update created by a
	// SoftLayer portal or API user, for an update created by a SoftLayer employee, or if a ticket update
	// was generated automatically by SoftLayer's backend systems.
	EditorType string `json:"editorType,omitempty"`
}

func (softlayer_ticket_update_employee *SoftLayer_Ticket_Update_Employee) String() string {
	return "SoftLayer_Ticket_Update_Employee"
}

// SoftLayer_Ticket_Update_Employee_Extended is SoftLayer_Ticket_Update_Employee with all maskable types.
type SoftLayer_Ticket_Update_Employee_Extended struct {
	SoftLayer_Ticket_Update_Employee

	// FileAttachment - no documentation
	FileAttachment []*SoftLayer_Ticket_Attachment_File `json:"fileAttachment,omitempty"`

	// Ticket - no documentation
	Ticket *SoftLayer_Ticket `json:"ticket,omitempty"`

	// FileAttachmentCount - no documentation
	FileAttachmentCount uint64 `json:"fileAttachmentCount,omitempty"`

	// Type - no documentation
	Type *SoftLayer_Ticket_Update_Type `json:"type,omitempty"`

	// ChangeOwnerActivity - <nil>
	ChangeOwnerActivity string `json:"changeOwnerActivity,omitempty"`

	// Editor - The user or SoftLayer employee who created a ticket update.
	Editor *SoftLayer_User_Interface `json:"editor,omitempty"`
}

func (softlayer_ticket_update_employee *SoftLayer_Ticket_Update_Employee_Extended) String() string {
	return "SoftLayer_Ticket_Update_Employee"
}
