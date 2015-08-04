package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Subnet_Swip_Transaction - The SoftLayer_Network_Subnet_Swip_Transaction data type
// contains basic information tracked at SoftLayer to allow automation of Swip creation, update, and
// removal requests. A specific transaction is attached to an accountId and a subnetId. This also
// contains a "Status Name" which tells the customer what the transaction is doing: * Request is queued
// up to be sent to * The email request has been sent to * has confirmed that the request is good, and
// should be available in 24 hours * OK: The subnet has been checked with and it the transaction has
// completed correctly * A subnet is queued to be removed from ARIN's systems * The removal email
// request has been sent to * has confirmed that the removal request is good, and the subnet should be
// clear in in 24 hours * This specific Transaction has been removed from and is no longer in effect *
// Sometimes a request doesn't go through correctly and has to be manually processed by SoftLayer. This
// may take some time.
type SoftLayer_Network_Subnet_Swip_Transaction struct {

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// StatusName - A Name describing which state a transaction is in.
	StatusName string `json:"statusName,omitempty"`

	// SubnetId - no documentation
	SubnetId int `json:"subnetId,omitempty"`
}

func (softlayer_network_subnet_swip_transaction *SoftLayer_Network_Subnet_Swip_Transaction) String() string {
	return "SoftLayer_Network_Subnet_Swip_Transaction"
}

// SoftLayer_Network_Subnet_Swip_Transaction_Extended is SoftLayer_Network_Subnet_Swip_Transaction with all maskable types.
type SoftLayer_Network_Subnet_Swip_Transaction_Extended struct {
	SoftLayer_Network_Subnet_Swip_Transaction

	// Account - no documentation
	Account *SoftLayer_Account `json:"account,omitempty"`

	// Subnet - no documentation
	Subnet *SoftLayer_Network_Subnet `json:"subnet,omitempty"`
}

func (softlayer_network_subnet_swip_transaction *SoftLayer_Network_Subnet_Swip_Transaction_Extended) String() string {
	return "SoftLayer_Network_Subnet_Swip_Transaction"
}
