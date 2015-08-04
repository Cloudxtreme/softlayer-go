package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Monitor_Version1_Query_Host_Stratum - The monitoring stratum type stores the
// maximum level of the various components of the monitoring system that a particular hardware object
// has access to. This object cannot be accessed by ID, and cannot be modified. The user can access
// this object through Hardware_Server->availableMonitoring. There are two values on this object that
// are important: # monitorLevel determines the highest level of
// SoftLayer_Network_Monitor_Version1_Query_Type object that can be placed in a monitoring instance on
// this server # responseLevel determines the highest level of
// SoftLayer_Network_Monitor_Version1_Query_ResponseType object that can be placed in a monitoring
// instance on this server Also note that the query type and response types are available through
// getAllQueryTypes and getAllResponseTypes, respectively.
type SoftLayer_Network_Monitor_Version1_Query_Host_Stratum struct {

	// MonitorLevel - The highest level of a monitoring query type allowed on this server
	MonitorLevel int `json:"monitorLevel,omitempty"`

	// ResponseLevel - The highest level of a monitoring response type allowed on this server
	ResponseLevel int `json:"responseLevel,omitempty"`
}

func (softlayer_network_monitor_version1_query_host_stratum *SoftLayer_Network_Monitor_Version1_Query_Host_Stratum) String() string {
	return "SoftLayer_Network_Monitor_Version1_Query_Host_Stratum"
}

// SoftLayer_Network_Monitor_Version1_Query_Host_Stratum_Extended is SoftLayer_Network_Monitor_Version1_Query_Host_Stratum with all maskable types.
type SoftLayer_Network_Monitor_Version1_Query_Host_Stratum_Extended struct {
	SoftLayer_Network_Monitor_Version1_Query_Host_Stratum

	// Hardware - The hardware object that these monitoring permissions applies to.
	Hardware *SoftLayer_Hardware `json:"hardware,omitempty"`
}

func (softlayer_network_monitor_version1_query_host_stratum *SoftLayer_Network_Monitor_Version1_Query_Host_Stratum_Extended) String() string {
	return "SoftLayer_Network_Monitor_Version1_Query_Host_Stratum"
}
