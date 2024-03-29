package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Message_Queue - The SoftLayer_Network_Message_Queue data type contains general
// information relating to Message Queue account
type SoftLayer_Network_Message_Queue struct {

	// AccountId - A message queue's associated [[SoftLayer_Account|account]] id.
	AccountId int `json:"accountId,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// MessageQueueStatusId - no documentation
	MessageQueueStatusId int `json:"messageQueueStatusId,omitempty"`

	// Notes - no documentation
	Notes string `json:"notes,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account,omitempty"`

	// BillingItem - The current billing item for this message queue account.
	BillingItem *SoftLayer_Billing_Item `json:"billingItem,omitempty"`

	// Nodes - no documentation
	Nodes []*SoftLayer_Network_Message_Queue_Node `json:"nodes,omitempty"`

	// NodeCount - no documentation
	NodeCount uint64 `json:"nodeCount,omitempty"`

	// Status - no documentation
	Status *SoftLayer_Network_Message_Queue_Status `json:"status,omitempty"`
}

func (softlayer_network_message_queue *SoftLayer_Network_Message_Queue) String() string {
	return "SoftLayer_Network_Message_Queue"
}
