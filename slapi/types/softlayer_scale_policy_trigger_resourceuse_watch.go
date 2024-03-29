package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Scale_Policy_Trigger_ResourceUse_Watch - <nil>
type SoftLayer_Scale_Policy_Trigger_ResourceUse_Watch struct {

	// Algorithm - The algorithm to use when aggregating and comparing. Currently, the only value that is
	// accepted is (Exponential Weighted Moving Average). is the default value if no value is given.
	Algorithm string `json:"algorithm,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// Metric - The metric to watch. Possible values: * host.cpu.percent - On a scale of 0 to 100, the
	// percent CPU a guest is using. * host.network.backend.in and host.network.frontend.in - The network
	// bytes-per-second incoming on the interface of either the frontend or backend network. *
	// host.network.backend.out and host.network.frontend.out - The network bytes-per-second incoming on
	// the interface of either the frontend or backend network.
	Metric string `json:"metric,omitempty"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// Operator - The operator to use for comparison. The only two valid values are ">" and
	Operator string `json:"operator,omitempty"`

	// Period - The number of seconds the values are aggregated for when compared to value. If values are
	// not retrieved steadily and consecutively for the length of this period, nothing is compared.
	Period int `json:"period,omitempty"`

	// DeleteFlag - When set and true any edit that happens on this object, be it calling edit on this
	// directly or setting as a child while editing a parent object, will end up being a deletion.
	DeleteFlag bool `json:"deleteFlag,omitempty"`

	// ScalePolicyTriggerId - no documentation
	ScalePolicyTriggerId int `json:"scalePolicyTriggerId,omitempty"`

	// Value - The value to compare against. Although the value is a string, validation will be done on the
	// value for restrictions (such as numeric-only) based on the metric.
	Value string `json:"value,omitempty"`

	// ScalePolicyTrigger - no documentation
	ScalePolicyTrigger *SoftLayer_Scale_Policy_Trigger_ResourceUse `json:"scalePolicyTrigger,omitempty"`
}

func (softlayer_scale_policy_trigger_resourceuse_watch *SoftLayer_Scale_Policy_Trigger_ResourceUse_Watch) String() string {
	return "SoftLayer_Scale_Policy_Trigger_ResourceUse_Watch"
}
