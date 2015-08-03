package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Network_Subnet_Registration_SubnetReference -
// SoftLayer_Container_Network_Subnet_Registration_SubnetReference is provided to reference
// [[SoftLayer_Network_Subnet_Registration]] object and the [[SoftLayer_Network_Subnet]] it references,
// in form.
type SoftLayer_Container_Network_Subnet_Registration_SubnetReference struct {

	// RegistrationId - The ID of the [[SoftLayer_Network_Subnet_Registration]] object.
	RegistrationId int `json:"registrationId"`

	// SubnetCidr - no documentation
	SubnetCidr string `json:"subnetCidr"`
}

func (softlayer_container_network_subnet_registration_subnetreference *SoftLayer_Container_Network_Subnet_Registration_SubnetReference) String() string {
	return "SoftLayer_Container_Network_Subnet_Registration_SubnetReference"
}
