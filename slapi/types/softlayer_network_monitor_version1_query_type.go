package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Monitor_Version1_Query_Type - The MonitorType type stores a name, long
// description, and default arguments for the monitor types. The only use for this object is in
// reference. The user chooses a monitoring type that would be appropriate for their server, and sets
// the id of the Query_Type to SoftLayer_Network_Monitor_Version1_Query_Host->queryTypeId The user can
// retrieve all available Query Types with the getAllObjects method on this service.
type SoftLayer_Network_Monitor_Version1_Query_Type struct {

	// ArgumentDescription - The type of parameter sent to the monitoring command.
	ArgumentDescription string `json:"argumentDescription"`

	// Description - no documentation
	Description string `json:"description"`

	// Id - no documentation
	Id int `json:"id"`

	// MonitorLevel - The level of this monitoring type. The level the customer has access to is determined
	// by values in SoftLayer_Network_Monitor_Version1_Query_Host_Stratum
	MonitorLevel int `json:"monitorLevel"`

	// Name - no documentation
	Name string `json:"name"`
}

func (softlayer_network_monitor_version1_query_type *SoftLayer_Network_Monitor_Version1_Query_Type) String() string {
	return "SoftLayer_Network_Monitor_Version1_Query_Type"
}
