package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Service_Resource - The SoftLayer_Network_Service_Resource is used to store
// information related to a service. It is used for determining the correct resource to connect to for
// a given service, like Evault, etc.
type SoftLayer_Network_Service_Resource struct {

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// FrontendIpAddress - no documentation
	FrontendIpAddress string `json:"frontendIpAddress,omitempty"`

	// BackendIpAddress - no documentation
	BackendIpAddress string `json:"backendIpAddress,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// NetworkDevice - The hardware information associated with this resource.
	NetworkDevice *SoftLayer_Hardware `json:"networkDevice,omitempty"`

	// ApiPassword - <nil>
	ApiPassword string `json:"apiPassword,omitempty"`

	// ApiPort - <nil>
	ApiPort string `json:"apiPort,omitempty"`

	// ApiProtocol - <nil>
	ApiProtocol string `json:"apiProtocol,omitempty"`

	// ApiUsername - <nil>
	ApiUsername string `json:"apiUsername,omitempty"`

	// Attributes - <nil>
	Attributes []*SoftLayer_Network_Service_Resource_Attribute `json:"attributes,omitempty"`

	// ApiHost - <nil>
	ApiHost string `json:"apiHost,omitempty"`

	// ApiPath - <nil>
	ApiPath string `json:"apiPath,omitempty"`

	// SshUsername - <nil>
	SshUsername string `json:"sshUsername,omitempty"`

	// AttributeCount - no documentation
	AttributeCount uint64 `json:"attributeCount,omitempty"`

	// ApiVersion - <nil>
	ApiVersion string `json:"apiVersion,omitempty"`

	// Datacenter - <nil>
	Datacenter *SoftLayer_Location `json:"datacenter,omitempty"`

	// Type - The network information associated with this resource.
	Type *SoftLayer_Network_Service_Resource_Type `json:"type,omitempty"`
}

func (softlayer_network_service_resource *SoftLayer_Network_Service_Resource) String() string {
	return "SoftLayer_Network_Service_Resource"
}
