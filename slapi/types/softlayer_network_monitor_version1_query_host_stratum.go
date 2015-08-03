package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

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

	// Hardware - The hardware object that these monitoring permissions applies to.
	Hardware *SoftLayer_Hardware `json:"hardware"`

	// MonitorLevel - The highest level of a monitoring query type allowed on this server
	MonitorLevel int `json:"monitorLevel"`

	// ResponseLevel - The highest level of a monitoring response type allowed on this server
	ResponseLevel int `json:"responseLevel"`
}

func (softlayer_network_monitor_version1_query_host_stratum *SoftLayer_Network_Monitor_Version1_Query_Host_Stratum) String() string {
	return "SoftLayer_Network_Monitor_Version1_Query_Host_Stratum"
}

// GetAllQueryTypes - Calling this function returns all possible query type objects. These objects are
// to be used to set the values on the SoftLayer_Network_Monitor_Version1_Query_Host when creating new
// monitoring instances.
func (softlayer_network_monitor_version1_query_host_stratum *SoftLayer_Network_Monitor_Version1_Query_Host_Stratum) GetAllQueryTypes(ctx *slapi.RequestContext) ([]*SoftLayer_Network_Monitor_Version1_Query_Type, error) {
	var returnValue []*SoftLayer_Network_Monitor_Version1_Query_Type
	return returnValue, nil
}

// GetAllResponseTypes - Calling this function returns all possible response type objects. These
// objects are to be used to set the values on the SoftLayer_Network_Monitor_Version1_Query_Host when
// creating new monitoring instances.
func (softlayer_network_monitor_version1_query_host_stratum *SoftLayer_Network_Monitor_Version1_Query_Host_Stratum) GetAllResponseTypes(ctx *slapi.RequestContext) ([]*SoftLayer_Network_Monitor_Version1_Query_ResponseType, error) {
	var returnValue []*SoftLayer_Network_Monitor_Version1_Query_ResponseType
	return returnValue, nil
}

// GetObject - getObject retrieves the SoftLayer_Network_Monitor_Version1_Query_Host_Stratum object
// whose ID number corresponds to the ID number of the init parameter passed to the
// SoftLayer_Network_Monitor_Version1_Query_Host_Stratum service. You can only retrieve strata attached
// to hardware that belong to your account.
func (softlayer_network_monitor_version1_query_host_stratum *SoftLayer_Network_Monitor_Version1_Query_Host_Stratum) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Network_Monitor_Version1_Query_Host_Stratum, error) {
	var returnValue *SoftLayer_Network_Monitor_Version1_Query_Host_Stratum
	return returnValue, nil
}
