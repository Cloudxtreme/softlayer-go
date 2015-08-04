package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Provisioning_Version1_Transaction_Status - The
// SoftLayer_Provisioning_Version1_Transaction_Status data type contains general information relating
// to a single SoftLayer hardware transaction status. SoftLayer customers are unable to change their
// hardware transaction status.
type SoftLayer_Provisioning_Version1_Transaction_Status struct {

	// FriendlyName - no documentation
	FriendlyName string `json:"friendlyName,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// AverageDuration - no documentation
	AverageDuration slapi.Float64 `json:"averageDuration,omitempty"`

	// NonCompletedTransactionCount - no documentation
	NonCompletedTransactionCount uint64 `json:"nonCompletedTransactionCount,omitempty"`

	// NonCompletedTransactions - <nil>
	NonCompletedTransactions []*SoftLayer_Provisioning_Version1_Transaction `json:"nonCompletedTransactions,omitempty"`
}

func (softlayer_provisioning_version1_transaction_status *SoftLayer_Provisioning_Version1_Transaction_Status) String() string {
	return "SoftLayer_Provisioning_Version1_Transaction_Status"
}
