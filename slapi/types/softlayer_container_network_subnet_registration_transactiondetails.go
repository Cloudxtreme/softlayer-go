package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Network_Subnet_Registration_TransactionDetails -
// SoftLayer_Container_Subnet_Registration_TransactionDetails is provided to return details of a newly
// created Subnet Registration Transaction.
type SoftLayer_Container_Network_Subnet_Registration_TransactionDetails struct {

	// SubnetReferences - The IDs and Subnets of the [[SoftLayer_Network_Subnet_Registration]] object.
	SubnetReferences []*SoftLayer_Container_Network_Subnet_Registration_SubnetReference `json:"subnetReferences"`

	// TransactionId - no documentation
	TransactionId int `json:"transactionId"`
}

func (softlayer_container_network_subnet_registration_transactiondetails *SoftLayer_Container_Network_Subnet_Registration_TransactionDetails) String() string {
	return "SoftLayer_Container_Network_Subnet_Registration_TransactionDetails"
}
