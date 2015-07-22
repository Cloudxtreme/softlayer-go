package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
)

// SoftLayer_Network_Message_Queue_Node - The SoftLayer_Network_Message_Queue_Node data type contains
// general information relating to Message Queue node
type SoftLayer_Network_Message_Queue_Node struct {

	// AccountName - no documentation
	AccountName string `json:"accountName"`

	// Id - A message queue node's internal identification number
	Id int `json:"id"`

	// MessageQueue - no documentation
	MessageQueue *SoftLayer_Network_Message_Queue `json:"messageQueue"`

	// MessageQueueId - A message queue node's associated message queue id.
	MessageQueueId int `json:"messageQueueId"`

	// MetricTrackingObject - A message queue node's metric tracking object. This object records all
	// request and notification count data for this message queue node.
	MetricTrackingObject *SoftLayer_Metric_Tracking_Object `json:"metricTrackingObject"`

	// Name - no documentation
	Name string `json:"name"`

	// Notes - no documentation
	Notes string `json:"notes"`

	// ServiceResource - <nil>
	ServiceResource *SoftLayer_Network_Service_Resource `json:"serviceResource"`
}

// AddUser - <nil>
func (softlayer_network_message_queue_node *SoftLayer_Network_Message_Queue_Node) AddUser(ctx *slapi.RequestContext, username string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// DeleteUser - <nil>
func (softlayer_network_message_queue_node *SoftLayer_Network_Message_Queue_Node) DeleteUser(ctx *slapi.RequestContext, username string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetAllUsers - <nil>
func (softlayer_network_message_queue_node *SoftLayer_Network_Message_Queue_Node) GetAllUsers(ctx *slapi.RequestContext) ([]string, error) {
	var returnValue []string
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_network_message_queue_node *SoftLayer_Network_Message_Queue_Node) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Network_Message_Queue_Node, error) {
	var returnValue *SoftLayer_Network_Message_Queue_Node
	return returnValue, nil
}

// GetUsage - no documentation
func (softlayer_network_message_queue_node *SoftLayer_Network_Message_Queue_Node) GetUsage(ctx *slapi.RequestContext, startDate time.Time, endDate time.Time) ([]*SoftLayer_Metric_Tracking_Object_Data, error) {
	var returnValue []*SoftLayer_Metric_Tracking_Object_Data
	return returnValue, nil
}

// GetUsageGraph - no documentation
func (softlayer_network_message_queue_node *SoftLayer_Network_Message_Queue_Node) GetUsageGraph(ctx *slapi.RequestContext, graphData SoftLayer_Container_Graph) (*SoftLayer_Container_Graph, error) {
	var returnValue *SoftLayer_Container_Graph
	return returnValue, nil
}
