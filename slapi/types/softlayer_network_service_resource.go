package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Service_Resource - The SoftLayer_Network_Service_Resource is used to store
// information related to a service. It is used for determining the correct resource to connect to for
// a given service, like Evault, etc.
type SoftLayer_Network_Service_Resource struct {

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// BackendIpAddress - no documentation
	BackendIpAddress string `json:"backendIpAddress,omitempty"`

	// FrontendIpAddress - no documentation
	FrontendIpAddress string `json:"frontendIpAddress,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`
}

func (softlayer_network_service_resource *SoftLayer_Network_Service_Resource) String() string {
	return "SoftLayer_Network_Service_Resource"
}

// SoftLayer_Network_Service_Resource_Extended is SoftLayer_Network_Service_Resource with all maskable types.
type SoftLayer_Network_Service_Resource_Extended struct {
	SoftLayer_Network_Service_Resource

	// ApiPassword - <nil>
	ApiPassword string `json:"apiPassword,omitempty"`

	// ApiPort - <nil>
	ApiPort string `json:"apiPort,omitempty"`

	// ApiUsername - <nil>
	ApiUsername string `json:"apiUsername,omitempty"`

	// ApiVersion - <nil>
	ApiVersion string `json:"apiVersion,omitempty"`

	// ApiHost - <nil>
	ApiHost string `json:"apiHost,omitempty"`

	// ApiPath - <nil>
	ApiPath string `json:"apiPath,omitempty"`

	// ApiProtocol - <nil>
	ApiProtocol string `json:"apiProtocol,omitempty"`

	// Datacenter - <nil>
	Datacenter *SoftLayer_Location `json:"datacenter,omitempty"`

	// NetworkDevice - The hardware information associated with this resource.
	NetworkDevice *SoftLayer_Hardware `json:"networkDevice,omitempty"`

	// AttributeCount - no documentation
	AttributeCount uint64 `json:"attributeCount,omitempty"`

	// Attributes - <nil>
	Attributes []*SoftLayer_Network_Service_Resource_Attribute `json:"attributes,omitempty"`

	// SshUsername - <nil>
	SshUsername string `json:"sshUsername,omitempty"`

	// Type - The network information associated with this resource.
	Type *SoftLayer_Network_Service_Resource_Type `json:"type,omitempty"`
}

func (softlayer_network_service_resource *SoftLayer_Network_Service_Resource_Extended) String() string {
	return "SoftLayer_Network_Service_Resource"
}
