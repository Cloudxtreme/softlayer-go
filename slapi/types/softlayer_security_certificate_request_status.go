package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

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

// GetObject - <nil>
func (softlayer_security_certificate_request_status *SoftLayer_Security_Certificate_Request_Status) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Security_Certificate_Request_Status, error) {
	var returnValue *SoftLayer_Security_Certificate_Request_Status
	return returnValue, nil
}

// GetSslRequestStatuses - Returns all SSL certificate request status objects
func (softlayer_security_certificate_request_status *SoftLayer_Security_Certificate_Request_Status) GetSslRequestStatuses(ctx *slapi.RequestContext) ([]*SoftLayer_Security_Certificate_Request_Status, error) {
	var returnValue []*SoftLayer_Security_Certificate_Request_Status
	return returnValue, nil
}
