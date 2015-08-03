package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Message_Queue - The SoftLayer_Network_Message_Queue data type contains general
// information relating to Message Queue account
type SoftLayer_Network_Message_Queue struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Name - no documentation
	Name string `json:"name"`

	// AccountId - A message queue's associated [[SoftLayer_Account|account]] id.
	AccountId int `json:"accountId"`

	// Id - no documentation
	Id int `json:"id"`

	// MessageQueueStatusId - no documentation
	MessageQueueStatusId int `json:"messageQueueStatusId"`

	// Notes - no documentation
	Notes string `json:"notes"`
}

// SoftLayer_Network_Message_Queue_Extended is SoftLayer_Network_Message_Queue with all maskable types.
type SoftLayer_Network_Message_Queue_Extended struct {
	SoftLayer_Network_Message_Queue

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// Status - no documentation
	Status *SoftLayer_Network_Message_Queue_Status `json:"status"`

	// NodeCount - no documentation
	NodeCount uint64 `json:"nodeCount"`

	// BillingItem - The current billing item for this message queue account.
	BillingItem *SoftLayer_Billing_Item `json:"billingItem"`

	// Nodes - no documentation
	Nodes []*SoftLayer_Network_Message_Queue_Node `json:"nodes"`
}

func (softlayer_network_message_queue *SoftLayer_Network_Message_Queue) String() string {
	return "SoftLayer_Network_Message_Queue"
}
