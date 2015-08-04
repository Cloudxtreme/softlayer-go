package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Exception - The SoftLayer_Container_Exception data type represents a
// SoftLayer_Exception.
type SoftLayer_Container_Exception struct {

	// ExceptionClass - no documentation
	ExceptionClass string `json:"exceptionClass,omitempty"`

	// ExceptionMessage - no documentation
	ExceptionMessage string `json:"exceptionMessage,omitempty"`
}

func (softlayer_container_exception *SoftLayer_Container_Exception) String() string {
	return "SoftLayer_Container_Exception"
}
