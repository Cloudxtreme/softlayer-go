package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Order_Hardware_Server_Gateway_Appliance - This is the datatype that
// needs to be populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything
// required to place a Gateway Appliance order.
type SoftLayer_Container_Product_Order_Hardware_Server_Gateway_Appliance struct {

	// ClusterIdentifier - Used to identify which items on an order belong in the same cluster.
	ClusterIdentifier string `json:"clusterIdentifier"`

	// ClusterOrderType - Indicates what type of cluster order is being placed Provision).
	ClusterOrderType string `json:"clusterOrderType"`

	// ClusterResourceId - Used to identify which gateway is being upgraded to
	ClusterResourceId int `json:"clusterResourceId"`

	// RequiredUpstreamDeviceId - Used to identify which device the new server should be attached to.
	RequiredUpstreamDeviceId int `json:"requiredUpstreamDeviceId"`
}
