package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Network_Media_Transcode_Preset - Transcode preset is a set of configuration
// parameters that defines a Transcode output format.
// SoftLayer_Container_Network_Media_Transcode_Preset contains a preset information defined on a
// Transcode server
type SoftLayer_Container_Network_Media_Transcode_Preset struct {

	// GUID - no documentation
	GUID string `json:"GUID,omitempty"`

	// Category - no documentation
	Category string `json:"category,omitempty"`

	// Description - no documentation
	Description string `json:"description,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`
}

func (softlayer_container_network_media_transcode_preset *SoftLayer_Container_Network_Media_Transcode_Preset) String() string {
	return "SoftLayer_Container_Network_Media_Transcode_Preset"
}
