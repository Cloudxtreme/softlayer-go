package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Location_Group_Regional - <nil>
type SoftLayer_Location_Group_Regional struct {

	// DatacenterCount - no documentation
	DatacenterCount uint64 `json:"datacenterCount"`

	// Datacenters - no documentation
	Datacenters []*SoftLayer_Location `json:"datacenters"`

	// PreferredDatacenter - no documentation
	PreferredDatacenter *SoftLayer_Location_Datacenter `json:"preferredDatacenter"`
}

// GetAllObjects - <nil>
func (softlayer_location_group_regional *SoftLayer_Location_Group_Regional) GetAllObjects(ctx *slapi.RequestContext) ([]*SoftLayer_Location_Group, error) {
	var returnValue []*SoftLayer_Location_Group
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_location_group_regional *SoftLayer_Location_Group_Regional) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Location_Group_Regional, error) {
	var returnValue *SoftLayer_Location_Group_Regional
	return returnValue, nil
}
