package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Security_Certificate_Request_ServerType - Represents a server type that can be specified
// when ordering an SSL certificate.
type SoftLayer_Security_Certificate_Request_ServerType struct {

	// Id - The internal identifier of the certificate server type.
	Id int `json:"id,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// Value - no documentation
	Value int `json:"value,omitempty"`

	// Description - no documentation
	Description string `json:"description,omitempty"`
}

func (softlayer_security_certificate_request_servertype *SoftLayer_Security_Certificate_Request_ServerType) String() string {
	return "SoftLayer_Security_Certificate_Request_ServerType"
}
