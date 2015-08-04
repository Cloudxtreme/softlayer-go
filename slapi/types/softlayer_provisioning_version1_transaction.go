package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Provisioning_Version1_Transaction - The SoftLayer_Provisioning_Version1_Transaction data
// type contains general information relating to a single SoftLayer hardware transaction. SoftLayer
// customers are unable to change their hardware transactions.
type SoftLayer_Provisioning_Version1_Transaction struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// HardwareId - A transaction's associated hardware identification number.
	HardwareId int `json:"hardwareId,omitempty"`

	// ElapsedSeconds - The amount of seconds that have elapsed since the transaction was last modified.
	ElapsedSeconds int `json:"elapsedSeconds,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// StatusChangeDate - no documentation
	StatusChangeDate *time.Time `json:"statusChangeDate,omitempty"`

	// GuestId - A transaction's associated guest identification number.
	GuestId int `json:"guestId,omitempty"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account,omitempty"`

	// TicketScheduledActionReferenceCount - no documentation
	TicketScheduledActionReferenceCount uint64 `json:"ticketScheduledActionReferenceCount,omitempty"`

	// TransactionGroup - A transaction's group. This group object determines what type of service is being
	// done on the hardware.
	TransactionGroup *SoftLayer_Provisioning_Version1_Transaction_Group `json:"transactionGroup,omitempty"`

	// LoopbackCount - no documentation
	LoopbackCount uint64 `json:"loopbackCount,omitempty"`

	// Guest - no documentation
	Guest *SoftLayer_Virtual_Guest `json:"guest,omitempty"`

	// Loopback - <nil>
	Loopback []*SoftLayer_Provisioning_Version1_Transaction `json:"loopback,omitempty"`

	// TicketScheduledActionReference - <nil>
	TicketScheduledActionReference []*SoftLayer_Ticket_Attachment `json:"ticketScheduledActionReference,omitempty"`

	// PendingTransactionCount - no documentation
	PendingTransactionCount uint64 `json:"pendingTransactionCount,omitempty"`

	// PendingTransactions - <nil>
	PendingTransactions []*SoftLayer_Provisioning_Version1_Transaction `json:"pendingTransactions,omitempty"`

	// TransactionStatus - A transaction's status. This status object determines the state it is in the
	// transaction group.
	TransactionStatus *SoftLayer_Provisioning_Version1_Transaction_Status `json:"transactionStatus,omitempty"`

	// Hardware - no documentation
	Hardware *SoftLayer_Hardware `json:"hardware,omitempty"`
}

func (softlayer_provisioning_version1_transaction *SoftLayer_Provisioning_Version1_Transaction) String() string {
	return "SoftLayer_Provisioning_Version1_Transaction"
}
