package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Message_Queue_Node - The SoftLayer_Network_Message_Queue_Node data type contains
// general information relating to Message Queue node
type SoftLayer_Network_Message_Queue_Node struct {

	// Notes - no documentation
	Notes string `json:"notes"`

	// AccountName - no documentation
	AccountName string `json:"accountName"`

	// Id - A message queue node's internal identification number
	Id int `json:"id"`

	// MessageQueueId - A message queue node's associated message queue id.
	MessageQueueId int `json:"messageQueueId"`

	// Name - no documentation
	Name string `json:"name"`
}

// SoftLayer_Network_Message_Queue_Node_Extended is SoftLayer_Network_Message_Queue_Node with all maskable types.
type SoftLayer_Network_Message_Queue_Node_Extended struct {
	SoftLayer_Network_Message_Queue_Node

	// MessageQueue - no documentation
	MessageQueue *SoftLayer_Network_Message_Queue `json:"messageQueue"`

	// MetricTrackingObject - A message queue node's metric tracking object. This object records all
	// request and notification count data for this message queue node.
	MetricTrackingObject *SoftLayer_Metric_Tracking_Object `json:"metricTrackingObject"`

	// ServiceResource - <nil>
	ServiceResource *SoftLayer_Network_Service_Resource `json:"serviceResource"`
}

func (softlayer_network_message_queue_node *SoftLayer_Network_Message_Queue_Node) String() string {
	return "SoftLayer_Network_Message_Queue_Node"
}
