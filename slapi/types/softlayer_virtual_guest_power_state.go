package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Virtual_Guest_Power_State - The power state class provides a common set of values for
// which a guest's power state will be presented in the SoftLayer
type SoftLayer_Virtual_Guest_Power_State struct {

	// Description - no documentation
	Description string `json:"description,omitempty"`

	// KeyName - no documentation
	KeyName string `json:"keyName,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`
}

func (softlayer_virtual_guest_power_state *SoftLayer_Virtual_Guest_Power_State) String() string {
	return "SoftLayer_Virtual_Guest_Power_State"
}
