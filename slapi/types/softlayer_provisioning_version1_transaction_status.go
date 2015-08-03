package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Provisioning_Version1_Transaction_Status - The
// SoftLayer_Provisioning_Version1_Transaction_Status data type contains general information relating
// to a single SoftLayer hardware transaction status. SoftLayer customers are unable to change their
// hardware transaction status.
type SoftLayer_Provisioning_Version1_Transaction_Status struct {

	// AverageDuration - no documentation
	AverageDuration float64 `json:"averageDuration"`

	// FriendlyName - no documentation
	FriendlyName string `json:"friendlyName"`

	// Name - no documentation
	Name string `json:"name"`

	// NonCompletedTransactionCount - no documentation
	NonCompletedTransactionCount uint64 `json:"nonCompletedTransactionCount"`

	// NonCompletedTransactions - <nil>
	NonCompletedTransactions []*SoftLayer_Provisioning_Version1_Transaction `json:"nonCompletedTransactions"`
}

func (softlayer_provisioning_version1_transaction_status *SoftLayer_Provisioning_Version1_Transaction_Status) String() string {
	return "SoftLayer_Provisioning_Version1_Transaction_Status"
}
