package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Order_SshKeys - This object holds all of the ssh key ids that will allow
// authentication to a single server.
type SoftLayer_Container_Product_Order_SshKeys struct {

	// SshKeyIds - An array of SoftLayer_Security_Ssh_Key IDs to assign to a server.
	SshKeyIds []int `json:"sshKeyIds,omitempty"`
}

func (softlayer_container_product_order_sshkeys *SoftLayer_Container_Product_Order_SshKeys) String() string {
	return "SoftLayer_Container_Product_Order_SshKeys"
}
