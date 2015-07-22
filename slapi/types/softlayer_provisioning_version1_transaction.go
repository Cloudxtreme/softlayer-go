package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Provisioning_Version1_Transaction - The SoftLayer_Provisioning_Version1_Transaction data
// type contains general information relating to a single SoftLayer hardware transaction. SoftLayer
// customers are unable to change their hardware transactions.
type SoftLayer_Provisioning_Version1_Transaction struct {

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// ElapsedSeconds - The amount of seconds that have elapsed since the transaction was last modified.
	ElapsedSeconds int `json:"elapsedSeconds"`

	// Guest - no documentation
	Guest *SoftLayer_Virtual_Guest `json:"guest"`

	// GuestId - A transaction's associated guest identification number.
	GuestId int `json:"guestId"`

	// Hardware - no documentation
	Hardware *SoftLayer_Hardware `json:"hardware"`

	// HardwareId - A transaction's associated hardware identification number.
	HardwareId int `json:"hardwareId"`

	// Id - no documentation
	Id int `json:"id"`

	// Loopback - <nil>
	Loopback []*SoftLayer_Provisioning_Version1_Transaction `json:"loopback"`

	// LoopbackCount - no documentation
	LoopbackCount uint64 `json:"loopbackCount"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// PendingTransactionCount - no documentation
	PendingTransactionCount uint64 `json:"pendingTransactionCount"`

	// PendingTransactions - <nil>
	PendingTransactions []*SoftLayer_Provisioning_Version1_Transaction `json:"pendingTransactions"`

	// StatusChangeDate - no documentation
	StatusChangeDate *time.Time `json:"statusChangeDate"`

	// TicketScheduledActionReference - <nil>
	TicketScheduledActionReference []*SoftLayer_Ticket_Attachment `json:"ticketScheduledActionReference"`

	// TicketScheduledActionReferenceCount - no documentation
	TicketScheduledActionReferenceCount uint64 `json:"ticketScheduledActionReferenceCount"`

	// TransactionGroup - A transaction's group. This group object determines what type of service is being
	// done on the hardware.
	TransactionGroup *SoftLayer_Provisioning_Version1_Transaction_Group `json:"transactionGroup"`

	// TransactionStatus - A transaction's status. This status object determines the state it is in the
	// transaction group.
	TransactionStatus *SoftLayer_Provisioning_Version1_Transaction_Status `json:"transactionStatus"`
}
