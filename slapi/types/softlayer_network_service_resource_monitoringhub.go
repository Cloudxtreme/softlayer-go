package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Service_Resource_MonitoringHub - <nil>
type SoftLayer_Network_Service_Resource_MonitoringHub struct {

	// BackendIpAddress - no documentation
	BackendIpAddress string `json:"backendIpAddress,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// FrontendIpAddress - no documentation
	FrontendIpAddress string `json:"frontendIpAddress,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// ApiProtocol - <nil>
	ApiProtocol string `json:"apiProtocol,omitempty"`

	// ApiHost - <nil>
	ApiHost string `json:"apiHost,omitempty"`

	// HubAddress - <nil>
	HubAddress string `json:"hubAddress,omitempty"`

	// RobotsCount - <nil>
	RobotsCount string `json:"robotsCount,omitempty"`

	// ApiPort - <nil>
	ApiPort string `json:"apiPort,omitempty"`

	// Attributes - <nil>
	Attributes []*SoftLayer_Network_Service_Resource_Attribute `json:"attributes,omitempty"`

	// SshUsername - <nil>
	SshUsername string `json:"sshUsername,omitempty"`

	// ApiVersion - <nil>
	ApiVersion string `json:"apiVersion,omitempty"`

	// ApiPassword - <nil>
	ApiPassword string `json:"apiPassword,omitempty"`

	// NetworkDevice - The hardware information associated with this resource.
	NetworkDevice *SoftLayer_Hardware `json:"networkDevice,omitempty"`

	// AttributeCount - no documentation
	AttributeCount uint64 `json:"attributeCount,omitempty"`

	// Type - The network information associated with this resource.
	Type *SoftLayer_Network_Service_Resource_Type `json:"type,omitempty"`

	// ApiPath - <nil>
	ApiPath string `json:"apiPath,omitempty"`

	// Datacenter - <nil>
	Datacenter *SoftLayer_Location `json:"datacenter,omitempty"`

	// HubConnectionTimeout - <nil>
	HubConnectionTimeout string `json:"hubConnectionTimeout,omitempty"`

	// RobotsMax - <nil>
	RobotsMax string `json:"robotsMax,omitempty"`

	// ApiUsername - <nil>
	ApiUsername string `json:"apiUsername,omitempty"`
}

func (softlayer_network_service_resource_monitoringhub *SoftLayer_Network_Service_Resource_MonitoringHub) String() string {
	return "SoftLayer_Network_Service_Resource_MonitoringHub"
}
