package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Security_Certificate_Request_ServerType - Represents a server type that can be specified
// when ordering an SSL certificate.
type SoftLayer_Security_Certificate_Request_ServerType struct {

	// Description - no documentation
	Description string `json:"description"`

	// Id - The internal identifier of the certificate server type.
	Id int `json:"id"`

	// Name - no documentation
	Name string `json:"name"`

	// Value - no documentation
	Value int `json:"value"`
}

// GetAllObjects - Returns all SSL certificate server types, which passed in on a certificate order.
func (softlayer_security_certificate_request_servertype *SoftLayer_Security_Certificate_Request_ServerType) GetAllObjects(ctx *slapi.RequestContext) ([]*SoftLayer_Security_Certificate_Request, error) {
	var returnValue []*SoftLayer_Security_Certificate_Request
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_security_certificate_request_servertype *SoftLayer_Security_Certificate_Request_ServerType) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Security_Certificate_Request_ServerType, error) {
	var returnValue *SoftLayer_Security_Certificate_Request_ServerType
	return returnValue, nil
}
