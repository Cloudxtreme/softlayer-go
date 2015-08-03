package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Security_Certificate_Request_Status - SoftLayer_Security_Certificate_Request_Status data
// type represents the status of an SSL certificate request.
type SoftLayer_Security_Certificate_Request_Status struct {

	// Description - The description of a SSL certificate request status
	Description string `json:"description"`

	// Id - The internal identifier of an SSL certificate request status
	Id int `json:"id"`

	// Name - no documentation
	Name string `json:"name"`
}

func (softlayer_security_certificate_request_status *SoftLayer_Security_Certificate_Request_Status) String() string {
	return "SoftLayer_Security_Certificate_Request_Status"
}
