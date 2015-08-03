package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Network_Message_Queue_Status - The SoftLayer_Network_Message_Queue_Status data type
// contains general information relating to Message Queue account status.
type SoftLayer_Network_Message_Queue_Status struct {

	// Description - A brief description on a message queue account status
	Description string `json:"description"`

	// Id - A message queue status's internal identification number
	Id int `json:"id"`

	// Name - A user-friendly name of a message queue account status
	Name string `json:"name"`
}

func (softlayer_network_message_queue_status *SoftLayer_Network_Message_Queue_Status) String() string {
	return "SoftLayer_Network_Message_Queue_Status"
}

// GetObject - <nil>
func (softlayer_network_message_queue_status *SoftLayer_Network_Message_Queue_Status) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Network_Message_Queue_Status, error) {
	var returnValue *SoftLayer_Network_Message_Queue_Status
	return returnValue, nil
}
