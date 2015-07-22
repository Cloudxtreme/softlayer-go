package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Provisioning_Version1_Transaction_Group - The
// SoftLayer_Provisioning_Version1_Transaction_Group data type contains general information relating to
// a single SoftLayer hardware transaction group. SoftLayer customers are unable to change their
// hardware transactions or the hardware transaction group.
type SoftLayer_Provisioning_Version1_Transaction_Group struct {

	// AverageTimeToComplete - Average time, in minutes, for this type of transaction to complete. Please
	// note that this is only an estimate.
	AverageTimeToComplete float64 `json:"averageTimeToComplete"`

	// Name - no documentation
	Name string `json:"name"`
}

// GetAllObjects - <nil>
func (softlayer_provisioning_version1_transaction_group *SoftLayer_Provisioning_Version1_Transaction_Group) GetAllObjects(commonOptions *slapi.CommonOptions) ([]*SoftLayer_Provisioning_Version1_Transaction_Group, error) {
	var returnValue []*SoftLayer_Provisioning_Version1_Transaction_Group
	return returnValue, nil
}

// GetObject - getObject retrieves the SoftLayer_Provisioning_Version1_Transaction_Group object whose
// ID number corresponds to the ID number of the init parameter passed to the
// SoftLayer_Provisioning_Version1_Transaction_Group service.
func (softlayer_provisioning_version1_transaction_group *SoftLayer_Provisioning_Version1_Transaction_Group) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Provisioning_Version1_Transaction_Group, error) {
	var returnValue *SoftLayer_Provisioning_Version1_Transaction_Group
	return returnValue, nil
}
