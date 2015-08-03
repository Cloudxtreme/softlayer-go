package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Account_Discount_Program - SoftLayer_Container_Account_Discount_Program models a
// single outbound object for a graph of given data sets.
type SoftLayer_Container_Account_Discount_Program struct {

	// LifetimeAppliedCredit - Credit allowance applied over the course of the entire program enrollment.
	// For enrollments without a lifetime restriction, this property will not be populated as credit will
	// be tracked on a purely monthly basis.
	LifetimeAppliedCredit float64 `json:"lifetimeAppliedCredit"`

	// LifetimeCredit - Credit allowance available over the course of the entire program enrollment. If
	// null, enrollment credit is applied on a strictly monthly basis and there is no lifetime maximum.
	// Enrollments with non-null lifetime credit will receive the lesser of the remaining monthly credit or
	// the remaining lifetime credit.
	LifetimeCredit float64 `json:"lifetimeCredit"`

	// LifetimeRemainingCredit - Remaining credit allowance available over the remaining duration of the
	// program enrollment. If null, enrollment credit is applied on a strictly monthly basis and there is
	// no lifetime maximum. Enrollments with non-null remaining lifetime credit will receive the lesser of
	// the remaining monthly credit or the remaining lifetime credit.
	LifetimeRemainingCredit float64 `json:"lifetimeRemainingCredit"`

	// MaximumActiveOrders - Maximum number of orders the enrolled account is allowed to have open at one
	// time. If null, then the Flexible Credit Program does not impose an order limit.
	MaximumActiveOrders float64 `json:"maximumActiveOrders"`

	// MonthlyCredit - The monthly credit allowance that is available at the beginning of the billing
	// cycle.
	MonthlyCredit float64 `json:"monthlyCredit"`

	// PostTaxRemainingCredit - Taxes are calculated in real time and discount amounts are shown pre-tax in
	// all cases. Tax values in the SoftLayer_Container_Account_Discount_Program container are now
	// populated with the related pre-tax values.
	PostTaxRemainingCredit float64 `json:"postTaxRemainingCredit"`

	// ProgramEndDate - no documentation
	ProgramEndDate *time.Time `json:"programEndDate"`

	// IsParticipant - Flag to signify whether the account is a participant in the discount program.
	IsParticipant bool `json:"isParticipant"`

	// RemainingCreditTax - Taxes are calculated in real time and discount amounts are shown pre-tax in all
	// cases. Tax values in the SoftLayer_Container_Account_Discount_Program container are now populated
	// with the related pre-tax values.
	RemainingCreditTax float64 `json:"remainingCreditTax"`

	// RemainingCredit - The credit allowance that is available during the current billing cycle. If the
	// lifetime limit has been or soon will be reached, this amount may be reduced by credit applied in
	// previous billing cycles.
	RemainingCredit float64 `json:"remainingCredit"`

	// ProgramName - Name of the Flexible Credit Program the account is enrolled in.
	ProgramName string `json:"programName"`

	// AppliedCredit - The credit allowance that has already been applied during the current billing cycle.
	// If the lifetime limit has been or soon will be reached, this amount may included credit applied in
	// previous billing cycles.
	AppliedCredit float64 `json:"appliedCredit"`
}

func (softlayer_container_account_discount_program *SoftLayer_Container_Account_Discount_Program) String() string {
	return "SoftLayer_Container_Account_Discount_Program"
}
