package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Service_Resource - The SoftLayer_Network_Service_Resource is used to store
// information related to a service. It is used for determining the correct resource to connect to for
// a given service, like Evault, etc.
type SoftLayer_Network_Service_Resource struct {

	// ApiHost - <nil>
	ApiHost string `json:"apiHost"`

	// ApiPassword - <nil>
	ApiPassword string `json:"apiPassword"`

	// ApiPath - <nil>
	ApiPath string `json:"apiPath"`

	// ApiPort - <nil>
	ApiPort string `json:"apiPort"`

	// ApiProtocol - <nil>
	ApiProtocol string `json:"apiProtocol"`

	// ApiUsername - <nil>
	ApiUsername string `json:"apiUsername"`

	// ApiVersion - <nil>
	ApiVersion string `json:"apiVersion"`

	// AttributeCount - no documentation
	AttributeCount uint64 `json:"attributeCount"`

	// Attributes - <nil>
	Attributes []*SoftLayer_Network_Service_Resource_Attribute `json:"attributes"`

	// BackendIpAddress - no documentation
	BackendIpAddress string `json:"backendIpAddress"`

	// Datacenter - <nil>
	Datacenter *SoftLayer_Location `json:"datacenter"`

	// FrontendIpAddress - no documentation
	FrontendIpAddress string `json:"frontendIpAddress"`

	// Id - <nil>
	Id int `json:"id"`

	// Name - no documentation
	Name string `json:"name"`

	// NetworkDevice - The hardware information associated with this resource.
	NetworkDevice *SoftLayer_Hardware `json:"networkDevice"`

	// SshUsername - <nil>
	SshUsername string `json:"sshUsername"`

	// Type - The network information associated with this resource.
	Type *SoftLayer_Network_Service_Resource_Type `json:"type"`
}
