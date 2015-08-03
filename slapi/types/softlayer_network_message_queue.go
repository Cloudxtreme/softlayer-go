package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Network_Message_Queue - The SoftLayer_Network_Message_Queue data type contains general
// information relating to Message Queue account
type SoftLayer_Network_Message_Queue struct {

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// AccountId - A message queue's associated [[SoftLayer_Account|account]] id.
	AccountId int `json:"accountId"`

	// BillingItem - The current billing item for this message queue account.
	BillingItem *SoftLayer_Billing_Item `json:"billingItem"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Id - no documentation
	Id int `json:"id"`

	// MessageQueueStatusId - no documentation
	MessageQueueStatusId int `json:"messageQueueStatusId"`

	// Name - no documentation
	Name string `json:"name"`

	// NodeCount - no documentation
	NodeCount uint64 `json:"nodeCount"`

	// Nodes - no documentation
	Nodes []*SoftLayer_Network_Message_Queue_Node `json:"nodes"`

	// Notes - no documentation
	Notes string `json:"notes"`

	// Status - no documentation
	Status *SoftLayer_Network_Message_Queue_Status `json:"status"`
}

func (softlayer_network_message_queue *SoftLayer_Network_Message_Queue) String() string {
	return "SoftLayer_Network_Message_Queue"
}

// GetObject - <nil>
func (softlayer_network_message_queue *SoftLayer_Network_Message_Queue) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Network_Message_Queue, error) {
	var returnValue *SoftLayer_Network_Message_Queue
	return returnValue, nil
}
