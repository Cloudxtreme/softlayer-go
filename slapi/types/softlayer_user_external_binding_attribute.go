package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_User_External_Binding_Attribute - The SoftLayer_User_External_Binding_Attribute data type
// contains the value for a single attribute associated with an external binding. External binding
// attributes contain additional information about an external binding. An attribute can be generic or
// specific to a 3rd party vendor. For example these attributes relate to Verisign: *Credential Type
// *Credential State *Credential Expiration Date *Credential Last Update Date
type SoftLayer_User_External_Binding_Attribute struct {

	// Value - no documentation
	Value string `json:"value"`
}

// SoftLayer_User_External_Binding_Attribute_Extended is SoftLayer_User_External_Binding_Attribute with all maskable types.
type SoftLayer_User_External_Binding_Attribute_Extended struct {
	SoftLayer_User_External_Binding_Attribute

	// ExternalBinding - The external authentication binding an attribute belongs to.
	ExternalBinding *SoftLayer_User_External_Binding `json:"externalBinding"`
}

func (softlayer_user_external_binding_attribute *SoftLayer_User_External_Binding_Attribute) String() string {
	return "SoftLayer_User_External_Binding_Attribute"
}
