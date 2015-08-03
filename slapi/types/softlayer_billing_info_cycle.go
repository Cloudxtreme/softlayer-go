package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Billing_Info_Cycle - The SoftLayer_Billing_Info_Cycle data type models basic information
// concerning a SoftLayer account's previous and current billing cycles. The information in this class
// is only populated for SoftLayer customers who are billed monthly.
type SoftLayer_Billing_Info_Cycle struct {

	// Account - The account that a current billing cycle is associated with.
	Account *SoftLayer_Account `json:"account"`

	// CurrentCycleEndDate - The ending date of an account's current billing cycle.
	CurrentCycleEndDate *time.Time `json:"currentCycleEndDate"`

	// CurrentCycleStartDate - The starting date of an account's current billing cycle.
	CurrentCycleStartDate *time.Time `json:"currentCycleStartDate"`

	// NextCycleStartDate - no documentation
	NextCycleStartDate *time.Time `json:"nextCycleStartDate"`

	// PreviousCycleEndDate - The ending date of an account's previous billing cycle.
	PreviousCycleEndDate *time.Time `json:"previousCycleEndDate"`

	// PreviousCycleStartDate - The starting date of an account's previous billing cycle.
	PreviousCycleStartDate *time.Time `json:"previousCycleStartDate"`
}

func (softlayer_billing_info_cycle *SoftLayer_Billing_Info_Cycle) String() string {
	return "SoftLayer_Billing_Info_Cycle"
}
