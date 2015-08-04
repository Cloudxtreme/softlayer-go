package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Message_Queue_Status - The SoftLayer_Network_Message_Queue_Status data type
// contains general information relating to Message Queue account status.
type SoftLayer_Network_Message_Queue_Status struct {

	// Id - A message queue status's internal identification number
	Id int `json:"id,omitempty"`

	// Name - A user-friendly name of a message queue account status
	Name string `json:"name,omitempty"`

	// Description - A brief description on a message queue account status
	Description string `json:"description,omitempty"`
}

func (softlayer_network_message_queue_status *SoftLayer_Network_Message_Queue_Status) String() string {
	return "SoftLayer_Network_Message_Queue_Status"
}
