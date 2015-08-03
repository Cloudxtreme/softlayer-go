package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Hardware_Function - The SoftLayer_Hardware_Function data type contains a generic object
// type for a piece of hardware, like switch, firewall, server, etc..
type SoftLayer_Hardware_Function struct {

	// Code - no documentation
	Code string `json:"code"`

	// Description - no documentation
	Description string `json:"description"`

	// Id - no documentation
	Id int `json:"id"`
}

func (softlayer_hardware_function *SoftLayer_Hardware_Function) String() string {
	return "SoftLayer_Hardware_Function"
}
