package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Network_Authentication_Data - This object holds authentication data to a server.
type SoftLayer_Container_Network_Authentication_Data struct {

	// Host - no documentation
	Host string `json:"host,omitempty"`

	// Password - no documentation
	Password string `json:"password,omitempty"`

	// Port - no documentation
	Port int `json:"port,omitempty"`

	// Type - The type of network protocol. This can be ftp, ssh and so on.
	Type string `json:"type,omitempty"`

	// Username - no documentation
	Username string `json:"username,omitempty"`
}

func (softlayer_container_network_authentication_data *SoftLayer_Container_Network_Authentication_Data) String() string {
	return "SoftLayer_Container_Network_Authentication_Data"
}
