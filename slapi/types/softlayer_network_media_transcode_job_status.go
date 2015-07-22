package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Network_Media_Transcode_Job_Status - The SoftLayer_Network_Media_Transcode_Job_Status
// contains information on a transcode job status.
type SoftLayer_Network_Media_Transcode_Job_Status struct {

	// Description - no documentation
	Description string `json:"description"`

	// Id - no documentation
	Id int `json:"id"`

	// Name - no documentation
	Name string `json:"name"`
}

// GetAllStatuses - no documentation
func (softlayer_network_media_transcode_job_status *SoftLayer_Network_Media_Transcode_Job_Status) GetAllStatuses(ctx *slapi.RequestContext) ([]*SoftLayer_Network_Media_Transcode_Job_Status, error) {
	var returnValue []*SoftLayer_Network_Media_Transcode_Job_Status
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_network_media_transcode_job_status *SoftLayer_Network_Media_Transcode_Job_Status) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Network_Media_Transcode_Job_Status, error) {
	var returnValue *SoftLayer_Network_Media_Transcode_Job_Status
	return returnValue, nil
}
