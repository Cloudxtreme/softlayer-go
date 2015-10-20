package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_User_External_Binding_Type - The SoftLayer_User_External_Binding_Type data type contains
// information relating to a type of external authentication binding. It contains a user friendly name
// as well as a unique key name.
type SoftLayer_User_External_Binding_Type struct {

	// Name - The user friendly name of a type of external authentication binding.
	Name string `json:"name,omitempty"`

	// KeyName - The unique name used to identify a type of external authentication binding.
	KeyName string `json:"keyName,omitempty"`
}

func (softlayer_user_external_binding_type *SoftLayer_User_External_Binding_Type) String() string {
	return "SoftLayer_User_External_Binding_Type"
}
