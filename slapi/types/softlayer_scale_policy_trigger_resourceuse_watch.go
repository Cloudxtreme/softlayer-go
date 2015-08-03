package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Scale_Policy_Trigger_ResourceUse_Watch - <nil>
type SoftLayer_Scale_Policy_Trigger_ResourceUse_Watch struct {

	// Algorithm - The algorithm to use when aggregating and comparing. Currently, the only value that is
	// accepted is (Exponential Weighted Moving Average). is the default value if no value is given.
	Algorithm string `json:"algorithm"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// DeleteFlag - When set and true any edit that happens on this object, be it calling edit on this
	// directly or setting as a child while editing a parent object, will end up being a deletion.
	DeleteFlag bool `json:"deleteFlag"`

	// Id - no documentation
	Id int `json:"id"`

	// Metric - The metric to watch. Possible values: * host.cpu.percent - On a scale of 0 to 100, the
	// percent CPU a guest is using. * host.network.backend.in and host.network.frontend.in - The network
	// bytes-per-second incoming on the interface of either the frontend or backend network. *
	// host.network.backend.out and host.network.frontend.out - The network bytes-per-second incoming on
	// the interface of either the frontend or backend network.
	Metric string `json:"metric"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// Operator - The operator to use for comparison. The only two valid values are ">" and
	Operator string `json:"operator"`

	// Period - The number of seconds the values are aggregated for when compared to value. If values are
	// not retrieved steadily and consecutively for the length of this period, nothing is compared.
	Period int `json:"period"`

	// ScalePolicyTrigger - no documentation
	ScalePolicyTrigger *SoftLayer_Scale_Policy_Trigger_ResourceUse `json:"scalePolicyTrigger"`

	// ScalePolicyTriggerId - no documentation
	ScalePolicyTriggerId int `json:"scalePolicyTriggerId"`

	// Value - The value to compare against. Although the value is a string, validation will be done on the
	// value for restrictions (such as numeric-only) based on the metric.
	Value string `json:"value"`
}

func (softlayer_scale_policy_trigger_resourceuse_watch *SoftLayer_Scale_Policy_Trigger_ResourceUse_Watch) String() string {
	return "SoftLayer_Scale_Policy_Trigger_ResourceUse_Watch"
}

// CreateObject - <nil>
func (softlayer_scale_policy_trigger_resourceuse_watch *SoftLayer_Scale_Policy_Trigger_ResourceUse_Watch) CreateObject(ctx *slapi.RequestContext, templateObject SoftLayer_Scale_Policy_Trigger_ResourceUse_Watch) (*SoftLayer_Scale_Policy_Trigger_ResourceUse_Watch, error) {
	var returnValue *SoftLayer_Scale_Policy_Trigger_ResourceUse_Watch
	return returnValue, nil
}

// DeleteObject - <nil>
func (softlayer_scale_policy_trigger_resourceuse_watch *SoftLayer_Scale_Policy_Trigger_ResourceUse_Watch) DeleteObject(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObject - <nil>
func (softlayer_scale_policy_trigger_resourceuse_watch *SoftLayer_Scale_Policy_Trigger_ResourceUse_Watch) EditObject(ctx *slapi.RequestContext, templateObject SoftLayer_Scale_Policy_Trigger_ResourceUse_Watch) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetAllPossibleAlgorithms - <nil>
func (softlayer_scale_policy_trigger_resourceuse_watch *SoftLayer_Scale_Policy_Trigger_ResourceUse_Watch) GetAllPossibleAlgorithms(ctx *slapi.RequestContext) ([]string, error) {
	var returnValue []string
	return returnValue, nil
}

// GetAllPossibleMetrics - <nil>
func (softlayer_scale_policy_trigger_resourceuse_watch *SoftLayer_Scale_Policy_Trigger_ResourceUse_Watch) GetAllPossibleMetrics(ctx *slapi.RequestContext) ([]string, error) {
	var returnValue []string
	return returnValue, nil
}

// GetAllPossibleOperators - <nil>
func (softlayer_scale_policy_trigger_resourceuse_watch *SoftLayer_Scale_Policy_Trigger_ResourceUse_Watch) GetAllPossibleOperators(ctx *slapi.RequestContext) ([]string, error) {
	var returnValue []string
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_scale_policy_trigger_resourceuse_watch *SoftLayer_Scale_Policy_Trigger_ResourceUse_Watch) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Scale_Policy_Trigger_ResourceUse_Watch, error) {
	var returnValue *SoftLayer_Scale_Policy_Trigger_ResourceUse_Watch
	return returnValue, nil
}
