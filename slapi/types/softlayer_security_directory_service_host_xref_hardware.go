package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Security_Directory_Service_Host_Xref_Hardware -
// SoftLayer_Security_Directory_Service_Host_Xref_Hardware extends the
// [[SoftLayer_Security_Directory_Service_Host_Xref]] data type to include hardware specific
// properties.
type SoftLayer_Security_Directory_Service_Host_Xref_Hardware struct {
}

func (softlayer_security_directory_service_host_xref_hardware *SoftLayer_Security_Directory_Service_Host_Xref_Hardware) String() string {
	return "SoftLayer_Security_Directory_Service_Host_Xref_Hardware"
}

// SoftLayer_Security_Directory_Service_Host_Xref_Hardware_Extended is SoftLayer_Security_Directory_Service_Host_Xref_Hardware with all maskable types.
type SoftLayer_Security_Directory_Service_Host_Xref_Hardware_Extended struct {
	SoftLayer_Security_Directory_Service_Host_Xref_Hardware

	// Host - no documentation
	Host *SoftLayer_Hardware `json:"host,omitempty"`
}

func (softlayer_security_directory_service_host_xref_hardware *SoftLayer_Security_Directory_Service_Host_Xref_Hardware_Extended) String() string {
	return "SoftLayer_Security_Directory_Service_Host_Xref_Hardware"
}
