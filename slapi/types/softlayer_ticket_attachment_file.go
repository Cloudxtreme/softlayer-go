package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Ticket_Attachment_File - SoftLayer tickets can have have files attached to them. Attaching
// a file to a ticket is a good way to report issues, provide documentation, and give examples of an
// issue. Both SoftLayer customers and employees have the ability to attach files to a ticket. The
// SoftLayer_Ticket_Attachment_File data type models a single file attached to a ticket.
type SoftLayer_Ticket_Attachment_File struct {

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// ModifyDate - The date that a file attachment record was last modified.
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// CreateDate - The date a file was originally attached to a ticket.
	CreateDate *time.Time `json:"createDate,omitempty"`

	// FileName - no documentation
	FileName string `json:"fileName,omitempty"`

	// FileSize - The size of a file attached to a ticket, measured in bytes.
	FileSize string `json:"fileSize,omitempty"`

	// TicketId - The internal identifier of the ticket that a file is attached to.
	TicketId int `json:"ticketId,omitempty"`

	// UpdateId - The internal identifier of the ticket update the attached file is associated with.
	UpdateId int `json:"updateId,omitempty"`

	// UploaderId - The internal identifier of the user that uploaded a ticket file attachment. This is
	// only used when A file attachment's ''uploaderType'' is set to
	UploaderId string `json:"uploaderId,omitempty"`

	// UploaderType - The type of user that attached a file to a ticket. This is either if the file was
	// uploaded by a portal or API user or if the file was uploaded by a SoftLayer employee.
	UploaderType string `json:"uploaderType,omitempty"`

	// Ticket - <nil>
	Ticket *SoftLayer_Ticket `json:"ticket,omitempty"`

	// Update - no documentation
	Update *SoftLayer_Ticket_Update `json:"update,omitempty"`
}

func (softlayer_ticket_attachment_file *SoftLayer_Ticket_Attachment_File) String() string {
	return "SoftLayer_Ticket_Attachment_File"
}
