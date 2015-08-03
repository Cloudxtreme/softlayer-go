package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Account_Graph_Outputs - no documentation
type SoftLayer_Container_Account_Graph_Outputs struct {

	// FailedBackupCount - The count of failed backups included in this graph.
	FailedBackupCount string `json:"failedBackupCount"`

	// UrlUptime - no documentation
	UrlUptime string `json:"urlUptime"`

	// WaitingEmployeeResponseTicketCount - no documentation
	WaitingEmployeeResponseTicketCount string `json:"waitingEmployeeResponseTicketCount"`

	// PendingCustomerResponseTicketCount - no documentation
	PendingCustomerResponseTicketCount string `json:"pendingCustomerResponseTicketCount"`

	// ClosedTickets - The count of closed tickets included in this graph.
	ClosedTickets string `json:"closedTickets"`

	// EndDate - no documentation
	EndDate *time.Time `json:"endDate"`

	// GraphImage - The raw PNG binary data to be displayed once the graph is drawn.
	GraphImage string `json:"graphImage"`

	// InboundUsage - no documentation
	InboundUsage string `json:"inboundUsage"`

	// OpenTickets - no documentation
	OpenTickets string `json:"openTickets"`

	// CompletedBackupCount - The count of completed backups included in this graph.
	CompletedBackupCount string `json:"completedBackupCount"`

	// ConflictBackupCount - The count of conflicted backups included in this graph.
	ConflictBackupCount string `json:"conflictBackupCount"`

	// HardwareUptime - The average of hardware uptime included in this graph.
	HardwareUptime string `json:"hardwareUptime"`

	// GraphError - no documentation
	GraphError string `json:"graphError"`

	// OutboundUsage - no documentation
	OutboundUsage string `json:"outboundUsage"`

	// StartDate - no documentation
	StartDate *time.Time `json:"startDate"`
}

func (softlayer_container_account_graph_outputs *SoftLayer_Container_Account_Graph_Outputs) String() string {
	return "SoftLayer_Container_Account_Graph_Outputs"
}
