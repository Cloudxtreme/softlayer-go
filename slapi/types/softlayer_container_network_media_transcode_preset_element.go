package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Network_Media_Transcode_Preset_Element - no documentation
type SoftLayer_Container_Network_Media_Transcode_Preset_Element struct {

	// Enabled - The flag that indicates whether an element is enabled or not
	Enabled bool `json:"enabled"`

	// ExtendedDescription - no documentation
	ExtendedDescription string `json:"extendedDescription"`

	// MinimumValue - no documentation
	MinimumValue int `json:"minimumValue"`

	// Name - no documentation
	Name string `json:"name"`

	// ParentName - no documentation
	ParentName string `json:"parentName"`

	// Type - no documentation
	Type string `json:"type"`

	// DefaultValue - no documentation
	DefaultValue string `json:"defaultValue"`

	// Description - no documentation
	Description string `json:"description"`

	// Hidden - The flag that indicates whether an element is hidden or not
	Hidden bool `json:"hidden"`

	// MaximumValue - no documentation
	MaximumValue int `json:"maximumValue"`

	// AdditionalElements - no documentation
	AdditionalElements []*SoftLayer_Container_Network_Media_Transcode_Preset_Element_Option `json:"additionalElements"`
}

func (softlayer_container_network_media_transcode_preset_element *SoftLayer_Container_Network_Media_Transcode_Preset_Element) String() string {
	return "SoftLayer_Container_Network_Media_Transcode_Preset_Element"
}
