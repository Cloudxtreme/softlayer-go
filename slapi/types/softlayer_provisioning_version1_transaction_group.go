package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

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

func (softlayer_provisioning_version1_transaction_group *SoftLayer_Provisioning_Version1_Transaction_Group) String() string {
	return "SoftLayer_Provisioning_Version1_Transaction_Group"
}
