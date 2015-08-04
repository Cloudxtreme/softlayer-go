package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Monitor_Version1_Incident - The SoftLayer_Network_Monitor_Version1_Incident data
// type models a single virtual server or physical hardware network monitoring event.
// SoftLayer_Network_Monitor_Version1_Incidents are created when the SoftLayer monitoring system
// detects a service down on your hardware of virtual server. As the incident is resolved it's status
// changes from to
type SoftLayer_Network_Monitor_Version1_Incident struct {

	// Status - A network monitoring incident's status, either the string denoting an ongoing incident or
	// meaning the incident has been resolved.
	Status string `json:"status,omitempty"`
}

func (softlayer_network_monitor_version1_incident *SoftLayer_Network_Monitor_Version1_Incident) String() string {
	return "SoftLayer_Network_Monitor_Version1_Incident"
}
