package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Service_Resource - The SoftLayer_Network_Service_Resource is used to store
// information related to a service. It is used for determining the correct resource to connect to for
// a given service, like Evault, etc.
type SoftLayer_Network_Service_Resource struct {

	// Name - no documentation
	Name string `json:"name"`

	// Id - <nil>
	Id int `json:"id"`

	// BackendIpAddress - no documentation
	BackendIpAddress string `json:"backendIpAddress"`

	// FrontendIpAddress - no documentation
	FrontendIpAddress string `json:"frontendIpAddress"`
}

// SoftLayer_Network_Service_Resource_Extended is SoftLayer_Network_Service_Resource with all maskable types.
type SoftLayer_Network_Service_Resource_Extended struct {
	SoftLayer_Network_Service_Resource

	// Type - The network information associated with this resource.
	Type *SoftLayer_Network_Service_Resource_Type `json:"type"`

	// AttributeCount - no documentation
	AttributeCount uint64 `json:"attributeCount"`

	// ApiUsername - <nil>
	ApiUsername string `json:"apiUsername"`

	// ApiHost - <nil>
	ApiHost string `json:"apiHost"`

	// ApiPassword - <nil>
	ApiPassword string `json:"apiPassword"`

	// ApiPath - <nil>
	ApiPath string `json:"apiPath"`

	// Attributes - <nil>
	Attributes []*SoftLayer_Network_Service_Resource_Attribute `json:"attributes"`

	// ApiPort - <nil>
	ApiPort string `json:"apiPort"`

	// ApiProtocol - <nil>
	ApiProtocol string `json:"apiProtocol"`

	// ApiVersion - <nil>
	ApiVersion string `json:"apiVersion"`

	// Datacenter - <nil>
	Datacenter *SoftLayer_Location `json:"datacenter"`

	// NetworkDevice - The hardware information associated with this resource.
	NetworkDevice *SoftLayer_Hardware `json:"networkDevice"`

	// SshUsername - <nil>
	SshUsername string `json:"sshUsername"`
}

func (softlayer_network_service_resource *SoftLayer_Network_Service_Resource) String() string {
	return "SoftLayer_Network_Service_Resource"
}
