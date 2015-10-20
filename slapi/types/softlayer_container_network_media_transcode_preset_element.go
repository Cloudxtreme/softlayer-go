package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Network_Media_Transcode_Preset_Element - no documentation
type SoftLayer_Container_Network_Media_Transcode_Preset_Element struct {

	// Enabled - The flag that indicates whether an element is enabled or not
	Enabled bool `json:"enabled,omitempty"`

	// ExtendedDescription - no documentation
	ExtendedDescription string `json:"extendedDescription,omitempty"`

	// ParentName - no documentation
	ParentName string `json:"parentName,omitempty"`

	// DefaultValue - no documentation
	DefaultValue string `json:"defaultValue,omitempty"`

	// Description - no documentation
	Description string `json:"description,omitempty"`

	// MaximumValue - no documentation
	MaximumValue int `json:"maximumValue,omitempty"`

	// MinimumValue - no documentation
	MinimumValue int `json:"minimumValue,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// Type - no documentation
	Type string `json:"type,omitempty"`

	// AdditionalElements - no documentation
	AdditionalElements []*SoftLayer_Container_Network_Media_Transcode_Preset_Element_Option `json:"additionalElements,omitempty"`

	// Hidden - The flag that indicates whether an element is hidden or not
	Hidden bool `json:"hidden,omitempty"`
}

func (softlayer_container_network_media_transcode_preset_element *SoftLayer_Container_Network_Media_Transcode_Preset_Element) String() string {
	return "SoftLayer_Container_Network_Media_Transcode_Preset_Element"
}
