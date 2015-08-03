package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Scale_Policy_Trigger_ResourceUse - <nil>
type SoftLayer_Scale_Policy_Trigger_ResourceUse struct {

	// WatchCount - no documentation
	WatchCount uint64 `json:"watchCount"`

	// Watches - no documentation
	Watches []*SoftLayer_Scale_Policy_Trigger_ResourceUse_Watch `json:"watches"`
}

func (softlayer_scale_policy_trigger_resourceuse *SoftLayer_Scale_Policy_Trigger_ResourceUse) String() string {
	return "SoftLayer_Scale_Policy_Trigger_ResourceUse"
}
