package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Monitoring_Agent_Status - no documentation
type SoftLayer_Monitoring_Agent_Status struct {

	// Description - no documentation
	Description string `json:"description"`

	// Id - no documentation
	Id int `json:"id"`

	// Name - no documentation
	Name string `json:"name"`
}

func (softlayer_monitoring_agent_status *SoftLayer_Monitoring_Agent_Status) String() string {
	return "SoftLayer_Monitoring_Agent_Status"
}

// GetObject - <nil>
func (softlayer_monitoring_agent_status *SoftLayer_Monitoring_Agent_Status) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Monitoring_Agent_Status, error) {
	var returnValue *SoftLayer_Monitoring_Agent_Status
	return returnValue, nil
}
