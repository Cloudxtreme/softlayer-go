package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Network_Media_Transcode_Preset_Element - no documentation
type SoftLayer_Container_Network_Media_Transcode_Preset_Element struct {

	// ExtendedDescription - no documentation
	ExtendedDescription string `json:"extendedDescription"`

	// MaximumValue - no documentation
	MaximumValue int `json:"maximumValue"`

	// MinimumValue - no documentation
	MinimumValue int `json:"minimumValue"`

	// Name - no documentation
	Name string `json:"name"`

	// Type - no documentation
	Type string `json:"type"`

	// AdditionalElements - no documentation
	AdditionalElements []*SoftLayer_Container_Network_Media_Transcode_Preset_Element_Option `json:"additionalElements"`

	// DefaultValue - no documentation
	DefaultValue string `json:"defaultValue"`

	// Hidden - The flag that indicates whether an element is hidden or not
	Hidden bool `json:"hidden"`

	// ParentName - no documentation
	ParentName string `json:"parentName"`

	// Description - no documentation
	Description string `json:"description"`

	// Enabled - The flag that indicates whether an element is enabled or not
	Enabled bool `json:"enabled"`
}

func (softlayer_container_network_media_transcode_preset_element *SoftLayer_Container_Network_Media_Transcode_Preset_Element) String() string {
	return "SoftLayer_Container_Network_Media_Transcode_Preset_Element"
}
