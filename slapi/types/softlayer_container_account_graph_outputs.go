package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Account_Graph_Outputs - no documentation
type SoftLayer_Container_Account_Graph_Outputs struct {

	// EndDate - no documentation
	EndDate *time.Time `json:"endDate,omitempty"`

	// GraphError - no documentation
	GraphError string `json:"graphError,omitempty"`

	// InboundUsage - no documentation
	InboundUsage string `json:"inboundUsage,omitempty"`

	// PendingCustomerResponseTicketCount - no documentation
	PendingCustomerResponseTicketCount string `json:"pendingCustomerResponseTicketCount,omitempty"`

	// ClosedTickets - The count of closed tickets included in this graph.
	ClosedTickets string `json:"closedTickets,omitempty"`

	// ConflictBackupCount - The count of conflicted backups included in this graph.
	ConflictBackupCount string `json:"conflictBackupCount,omitempty"`

	// OpenTickets - no documentation
	OpenTickets string `json:"openTickets,omitempty"`

	// OutboundUsage - no documentation
	OutboundUsage string `json:"outboundUsage,omitempty"`

	// UrlUptime - no documentation
	UrlUptime string `json:"urlUptime,omitempty"`

	// GraphImage - The raw PNG binary data to be displayed once the graph is drawn.
	GraphImage string `json:"graphImage,omitempty"`

	// HardwareUptime - The average of hardware uptime included in this graph.
	HardwareUptime string `json:"hardwareUptime,omitempty"`

	// WaitingEmployeeResponseTicketCount - no documentation
	WaitingEmployeeResponseTicketCount string `json:"waitingEmployeeResponseTicketCount,omitempty"`

	// CompletedBackupCount - The count of completed backups included in this graph.
	CompletedBackupCount string `json:"completedBackupCount,omitempty"`

	// StartDate - no documentation
	StartDate *time.Time `json:"startDate,omitempty"`

	// FailedBackupCount - The count of failed backups included in this graph.
	FailedBackupCount string `json:"failedBackupCount,omitempty"`
}

func (softlayer_container_account_graph_outputs *SoftLayer_Container_Account_Graph_Outputs) String() string {
	return "SoftLayer_Container_Account_Graph_Outputs"
}
