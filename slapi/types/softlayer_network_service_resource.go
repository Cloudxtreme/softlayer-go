package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Service_Resource - The SoftLayer_Network_Service_Resource is used to store
// information related to a service. It is used for determining the correct resource to connect to for
// a given service, like Evault, etc.
type SoftLayer_Network_Service_Resource struct {

	// FrontendIpAddress - no documentation
	FrontendIpAddress string `json:"frontendIpAddress,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// BackendIpAddress - no documentation
	BackendIpAddress string `json:"backendIpAddress,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// ApiPassword - <nil>
	ApiPassword string `json:"apiPassword,omitempty"`

	// ApiPath - <nil>
	ApiPath string `json:"apiPath,omitempty"`

	// ApiPort - <nil>
	ApiPort string `json:"apiPort,omitempty"`

	// Attributes - <nil>
	Attributes []*SoftLayer_Network_Service_Resource_Attribute `json:"attributes,omitempty"`

	// Type - The network information associated with this resource.
	Type *SoftLayer_Network_Service_Resource_Type `json:"type,omitempty"`

	// SshUsername - <nil>
	SshUsername string `json:"sshUsername,omitempty"`

	// ApiProtocol - <nil>
	ApiProtocol string `json:"apiProtocol,omitempty"`

	// NetworkDevice - The hardware information associated with this resource.
	NetworkDevice *SoftLayer_Hardware `json:"networkDevice,omitempty"`

	// ApiHost - <nil>
	ApiHost string `json:"apiHost,omitempty"`

	// ApiUsername - <nil>
	ApiUsername string `json:"apiUsername,omitempty"`

	// ApiVersion - <nil>
	ApiVersion string `json:"apiVersion,omitempty"`

	// Datacenter - <nil>
	Datacenter *SoftLayer_Location `json:"datacenter,omitempty"`

	// AttributeCount - no documentation
	AttributeCount uint64 `json:"attributeCount,omitempty"`
}

func (softlayer_network_service_resource *SoftLayer_Network_Service_Resource) String() string {
	return "SoftLayer_Network_Service_Resource"
}
