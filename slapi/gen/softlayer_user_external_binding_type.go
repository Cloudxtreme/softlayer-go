package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_User_External_Binding_Type - The SoftLayer_User_External_Binding_Type data type contains
// information relating to a type of external authentication binding. It contains a user friendly name
// as well as a unique key name.
type SoftLayer_User_External_Binding_Type struct {

	// KeyName - The unique name used to identify a type of external authentication binding.
	KeyName string `json:"keyName"`

	// Name - The user friendly name of a type of external authentication binding.
	Name string `json:"name"`
}