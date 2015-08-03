package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Provisioning_Version1_Transaction - The SoftLayer_Provisioning_Version1_Transaction data
// type contains general information relating to a single SoftLayer hardware transaction. SoftLayer
// customers are unable to change their hardware transactions.
type SoftLayer_Provisioning_Version1_Transaction struct {

	// Id - no documentation
	Id int `json:"id"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// ElapsedSeconds - The amount of seconds that have elapsed since the transaction was last modified.
	ElapsedSeconds int `json:"elapsedSeconds"`

	// GuestId - A transaction's associated guest identification number.
	GuestId int `json:"guestId"`

	// HardwareId - A transaction's associated hardware identification number.
	HardwareId int `json:"hardwareId"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// StatusChangeDate - no documentation
	StatusChangeDate *time.Time `json:"statusChangeDate"`
}

func (softlayer_provisioning_version1_transaction *SoftLayer_Provisioning_Version1_Transaction) String() string {
	return "SoftLayer_Provisioning_Version1_Transaction"
}

// SoftLayer_Provisioning_Version1_Transaction_Extended is SoftLayer_Provisioning_Version1_Transaction with all maskable types.
type SoftLayer_Provisioning_Version1_Transaction_Extended struct {
	SoftLayer_Provisioning_Version1_Transaction

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// Guest - no documentation
	Guest *SoftLayer_Virtual_Guest `json:"guest"`

	// TransactionGroup - A transaction's group. This group object determines what type of service is being
	// done on the hardware.
	TransactionGroup *SoftLayer_Provisioning_Version1_Transaction_Group `json:"transactionGroup"`

	// PendingTransactionCount - no documentation
	PendingTransactionCount uint64 `json:"pendingTransactionCount"`

	// LoopbackCount - no documentation
	LoopbackCount uint64 `json:"loopbackCount"`

	// Hardware - no documentation
	Hardware *SoftLayer_Hardware `json:"hardware"`

	// Loopback - <nil>
	Loopback []*SoftLayer_Provisioning_Version1_Transaction `json:"loopback"`

	// TransactionStatus - A transaction's status. This status object determines the state it is in the
	// transaction group.
	TransactionStatus *SoftLayer_Provisioning_Version1_Transaction_Status `json:"transactionStatus"`

	// PendingTransactions - <nil>
	PendingTransactions []*SoftLayer_Provisioning_Version1_Transaction `json:"pendingTransactions"`

	// TicketScheduledActionReference - <nil>
	TicketScheduledActionReference []*SoftLayer_Ticket_Attachment `json:"ticketScheduledActionReference"`

	// TicketScheduledActionReferenceCount - no documentation
	TicketScheduledActionReferenceCount uint64 `json:"ticketScheduledActionReferenceCount"`
}

func (softlayer_provisioning_version1_transaction *SoftLayer_Provisioning_Version1_Transaction_Extended) String() string {
	return "SoftLayer_Provisioning_Version1_Transaction"
}
