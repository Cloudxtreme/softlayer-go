package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Order_Monitoring_Package - This is the datatype that needs to be
// populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to
// place a Monitoring Package order with SoftLayer.
type SoftLayer_Container_Product_Order_Monitoring_Package struct {

	// ConfigurationTemplateGroups - <nil>
	ConfigurationTemplateGroups []*SoftLayer_Monitoring_Agent_Configuration_Template_Group `json:"configurationTemplateGroups"`

	// ServerType - <nil>
	ServerType string `json:"serverType"`
}

func (softlayer_container_product_order_monitoring_package *SoftLayer_Container_Product_Order_Monitoring_Package) String() string {
	return "SoftLayer_Container_Product_Order_Monitoring_Package"
}
