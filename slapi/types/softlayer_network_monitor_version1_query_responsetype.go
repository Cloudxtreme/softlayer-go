package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Monitor_Version1_Query_ResponseType - The ResponseType type stores only an ID and
// a description of the response type. The only use for this object is in reference. The user chooses a
// response action that would be appropriate for a monitoring instance, and sets the ResponseTypeId to
// the SoftLayer_Network_Monitor_Version1_Query_Host->responseActionId value. The user can retrieve all
// available ResponseTypes with the getAllObjects method on this service.
type SoftLayer_Network_Monitor_Version1_Query_ResponseType struct {

	// ActionDescription - The description of the action the monitoring system will take on failure
	ActionDescription string `json:"actionDescription"`

	// Id - no documentation
	Id int `json:"id"`

	// Level - The level of this response. The level the customer has access to is determined by values in
	// SoftLayer_Network_Monitor_Version1_Query_Host_Stratum
	Level int `json:"level"`
}

func (softlayer_network_monitor_version1_query_responsetype *SoftLayer_Network_Monitor_Version1_Query_ResponseType) String() string {
	return "SoftLayer_Network_Monitor_Version1_Query_ResponseType"
}
