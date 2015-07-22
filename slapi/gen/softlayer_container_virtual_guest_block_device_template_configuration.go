package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Virtual_Guest_Block_Device_Template_Configuration - The
// SoftLayer_Container_Virtual_Guest_Block_Device_Template_Configuration data type contains information
// relating to a template's external location for importing and exporting
type SoftLayer_Container_Virtual_Guest_Block_Device_Template_Configuration struct {

	// Name - The group name to be applied to the imported template
	Name string `json:"name"`

	// Note - no documentation
	Note string `json:"note"`

	// OperatingSystemReferenceCode - The referenceCode of the operating system software description for
	// the imported VHD
	OperatingSystemReferenceCode string `json:"operatingSystemReferenceCode"`

	// Uri - The URI for an object storage object (.vhd/.iso file) swift://
	Uri string `json:"uri"`
}
