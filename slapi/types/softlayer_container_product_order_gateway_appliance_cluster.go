package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Order_Gateway_Appliance_Cluster - This is the datatype that needs to be
// populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to
// place a Gateway Appliance Cluster order with SoftLayer.
type SoftLayer_Container_Product_Order_Gateway_Appliance_Cluster struct {

	// ClusterIdentifier - Used to identify which items on an order belong in the same cluster.
	ClusterIdentifier string `json:"clusterIdentifier"`

	// ClusterOrderType - Indicates what type of cluster order is being placed Provision).
	ClusterOrderType string `json:"clusterOrderType"`
}

func (softlayer_container_product_order_gateway_appliance_cluster *SoftLayer_Container_Product_Order_Gateway_Appliance_Cluster) String() string {
	return "SoftLayer_Container_Product_Order_Gateway_Appliance_Cluster"
}
