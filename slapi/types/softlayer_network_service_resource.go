package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Service_Resource - The SoftLayer_Network_Service_Resource is used to store
// information related to a service. It is used for determining the correct resource to connect to for
// a given service, like Evault, etc.
type SoftLayer_Network_Service_Resource struct {

	// BackendIpAddress - no documentation
	BackendIpAddress string `json:"backendIpAddress,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// FrontendIpAddress - no documentation
	FrontendIpAddress string `json:"frontendIpAddress,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`
}

func (softlayer_network_service_resource *SoftLayer_Network_Service_Resource) String() string {
	return "SoftLayer_Network_Service_Resource"
}

// SoftLayer_Network_Service_Resource_Extended is SoftLayer_Network_Service_Resource with all maskable types.
type SoftLayer_Network_Service_Resource_Extended struct {
	SoftLayer_Network_Service_Resource

	// AttributeCount - no documentation
	AttributeCount uint64 `json:"attributeCount,omitempty"`

	// ApiPort - <nil>
	ApiPort string `json:"apiPort,omitempty"`

	// NetworkDevice - The hardware information associated with this resource.
	NetworkDevice *SoftLayer_Hardware `json:"networkDevice,omitempty"`

	// Type - The network information associated with this resource.
	Type *SoftLayer_Network_Service_Resource_Type `json:"type,omitempty"`

	// ApiUsername - <nil>
	ApiUsername string `json:"apiUsername,omitempty"`

	// ApiPassword - <nil>
	ApiPassword string `json:"apiPassword,omitempty"`

	// ApiPath - <nil>
	ApiPath string `json:"apiPath,omitempty"`

	// SshUsername - <nil>
	SshUsername string `json:"sshUsername,omitempty"`

	// ApiHost - <nil>
	ApiHost string `json:"apiHost,omitempty"`

	// ApiVersion - <nil>
	ApiVersion string `json:"apiVersion,omitempty"`

	// Attributes - <nil>
	Attributes []*SoftLayer_Network_Service_Resource_Attribute `json:"attributes,omitempty"`

	// Datacenter - <nil>
	Datacenter *SoftLayer_Location `json:"datacenter,omitempty"`

	// ApiProtocol - <nil>
	ApiProtocol string `json:"apiProtocol,omitempty"`
}

func (softlayer_network_service_resource *SoftLayer_Network_Service_Resource_Extended) String() string {
	return "SoftLayer_Network_Service_Resource"
}
