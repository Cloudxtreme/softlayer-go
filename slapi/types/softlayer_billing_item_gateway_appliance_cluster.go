package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Billing_Item_Gateway_Appliance_Cluster - The SoftLayer_Billing_Item_Big_Data_Cluster data
// type contains general information relating to a single SoftLayer billing item for a big data
// cluster.
type SoftLayer_Billing_Item_Gateway_Appliance_Cluster struct {
}

// SoftLayer_Billing_Item_Gateway_Appliance_Cluster_Extended is SoftLayer_Billing_Item_Gateway_Appliance_Cluster with all maskable types.
type SoftLayer_Billing_Item_Gateway_Appliance_Cluster_Extended struct {
	SoftLayer_Billing_Item_Gateway_Appliance_Cluster

	// Resource - no documentation
	Resource *SoftLayer_Resource_Group `json:"resource"`
}

func (softlayer_billing_item_gateway_appliance_cluster *SoftLayer_Billing_Item_Gateway_Appliance_Cluster) String() string {
	return "SoftLayer_Billing_Item_Gateway_Appliance_Cluster"
}
