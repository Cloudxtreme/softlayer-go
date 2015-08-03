package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Security_Directory_Service_Host_Xref_Hardware -
// SoftLayer_Security_Directory_Service_Host_Xref_Hardware extends the
// [[SoftLayer_Security_Directory_Service_Host_Xref]] data type to include hardware specific
// properties.
type SoftLayer_Security_Directory_Service_Host_Xref_Hardware struct {

	// Host - no documentation
	Host *SoftLayer_Hardware `json:"host"`
}

func (softlayer_security_directory_service_host_xref_hardware *SoftLayer_Security_Directory_Service_Host_Xref_Hardware) String() string {
	return "SoftLayer_Security_Directory_Service_Host_Xref_Hardware"
}
