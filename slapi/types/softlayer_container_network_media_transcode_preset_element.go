package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Network_Media_Transcode_Preset_Element - no documentation
type SoftLayer_Container_Network_Media_Transcode_Preset_Element struct {

	// AdditionalElements - no documentation
	AdditionalElements []*SoftLayer_Container_Network_Media_Transcode_Preset_Element_Option `json:"additionalElements"`

	// DefaultValue - no documentation
	DefaultValue string `json:"defaultValue"`

	// Description - no documentation
	Description string `json:"description"`

	// Enabled - The flag that indicates whether an element is enabled or not
	Enabled bool `json:"enabled"`

	// ExtendedDescription - no documentation
	ExtendedDescription string `json:"extendedDescription"`

	// Hidden - The flag that indicates whether an element is hidden or not
	Hidden bool `json:"hidden"`

	// MaximumValue - no documentation
	MaximumValue int `json:"maximumValue"`

	// MinimumValue - no documentation
	MinimumValue int `json:"minimumValue"`

	// Name - no documentation
	Name string `json:"name"`

	// ParentName - no documentation
	ParentName string `json:"parentName"`

	// Type - no documentation
	Type string `json:"type"`
}