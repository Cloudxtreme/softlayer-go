package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Order_Hardware_Server - This is the datatype that needs to be populated
// and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to place an
// order with SoftLayer.
type SoftLayer_Container_Product_Order_Hardware_Server struct {

	// MonitoringAgentConfigurationTemplateGroupId - Id of the
	// [[SoftLayer_Monitoring_Agent_Configuration_Template_Group]] to be used with the monitoring package
	MonitoringAgentConfigurationTemplateGroupId int `json:"monitoringAgentConfigurationTemplateGroupId"`

	// PrivateCloudServerRole - When ordering Virtual Server (Private Node), this variable specifies the
	// role of the server configuration. (Deprecated)
	PrivateCloudServerRole string `json:"privateCloudServerRole"`

	// Tags - no documentation
	Tags []*SoftLayer_Container_Product_Order_Property `json:"tags"`
}

func (softlayer_container_product_order_hardware_server *SoftLayer_Container_Product_Order_Hardware_Server) String() string {
	return "SoftLayer_Container_Product_Order_Hardware_Server"
}
