package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"

	time "time"
)

// SoftLayer_Container_Account_Discount_Program - SoftLayer_Container_Account_Discount_Program models a
// single outbound object for a graph of given data sets.
type SoftLayer_Container_Account_Discount_Program struct {

	// AppliedCredit - The credit allowance that has already been applied during the current billing cycle.
	// If the lifetime limit has been or soon will be reached, this amount may included credit applied in
	// previous billing cycles.
	AppliedCredit slapi.Float64 `json:"appliedCredit,omitempty"`

	// IsParticipant - Flag to signify whether the account is a participant in the discount program.
	IsParticipant bool `json:"isParticipant,omitempty"`

	// LifetimeRemainingCredit - Remaining credit allowance available over the remaining duration of the
	// program enrollment. If null, enrollment credit is applied on a strictly monthly basis and there is
	// no lifetime maximum. Enrollments with non-null remaining lifetime credit will receive the lesser of
	// the remaining monthly credit or the remaining lifetime credit.
	LifetimeRemainingCredit slapi.Float64 `json:"lifetimeRemainingCredit,omitempty"`

	// MonthlyCredit - The monthly credit allowance that is available at the beginning of the billing
	// cycle.
	MonthlyCredit slapi.Float64 `json:"monthlyCredit,omitempty"`

	// ProgramName - Name of the Flexible Credit Program the account is enrolled in.
	ProgramName string `json:"programName,omitempty"`

	// RemainingCreditTax - Taxes are calculated in real time and discount amounts are shown pre-tax in all
	// cases. Tax values in the SoftLayer_Container_Account_Discount_Program container are now populated
	// with the related pre-tax values.
	RemainingCreditTax slapi.Float64 `json:"remainingCreditTax,omitempty"`

	// LifetimeAppliedCredit - Credit allowance applied over the course of the entire program enrollment.
	// For enrollments without a lifetime restriction, this property will not be populated as credit will
	// be tracked on a purely monthly basis.
	LifetimeAppliedCredit slapi.Float64 `json:"lifetimeAppliedCredit,omitempty"`

	// LifetimeCredit - Credit allowance available over the course of the entire program enrollment. If
	// null, enrollment credit is applied on a strictly monthly basis and there is no lifetime maximum.
	// Enrollments with non-null lifetime credit will receive the lesser of the remaining monthly credit or
	// the remaining lifetime credit.
	LifetimeCredit slapi.Float64 `json:"lifetimeCredit,omitempty"`

	// MaximumActiveOrders - Maximum number of orders the enrolled account is allowed to have open at one
	// time. If null, then the Flexible Credit Program does not impose an order limit.
	MaximumActiveOrders slapi.Float64 `json:"maximumActiveOrders,omitempty"`

	// PostTaxRemainingCredit - Taxes are calculated in real time and discount amounts are shown pre-tax in
	// all cases. Tax values in the SoftLayer_Container_Account_Discount_Program container are now
	// populated with the related pre-tax values.
	PostTaxRemainingCredit slapi.Float64 `json:"postTaxRemainingCredit,omitempty"`

	// ProgramEndDate - no documentation
	ProgramEndDate *time.Time `json:"programEndDate,omitempty"`

	// RemainingCredit - The credit allowance that is available during the current billing cycle. If the
	// lifetime limit has been or soon will be reached, this amount may be reduced by credit applied in
	// previous billing cycles.
	RemainingCredit slapi.Float64 `json:"remainingCredit,omitempty"`
}

func (softlayer_container_account_discount_program *SoftLayer_Container_Account_Discount_Program) String() string {
	return "SoftLayer_Container_Account_Discount_Program"
}
