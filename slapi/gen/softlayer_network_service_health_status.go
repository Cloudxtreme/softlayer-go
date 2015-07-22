package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Service_Health_Status - Many general services that SoftLayer provides are marked
// by a status message. These health messages give portal users a quick way of determining the state of
// a SoftLayer service. Services range from backbones to VPN endpoints and routers. Generally a health
// status is either "Up" or "Down".
type SoftLayer_Network_Service_Health_Status struct {

	// Name - The status of a SoftLayer service. This is typically "Up" or "Down".
	Name string `json:"name"`
}
