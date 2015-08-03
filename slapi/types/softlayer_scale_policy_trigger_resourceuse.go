package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Scale_Policy_Trigger_ResourceUse - <nil>
type SoftLayer_Scale_Policy_Trigger_ResourceUse struct {
}

func (softlayer_scale_policy_trigger_resourceuse *SoftLayer_Scale_Policy_Trigger_ResourceUse) String() string {
	return "SoftLayer_Scale_Policy_Trigger_ResourceUse"
}

// SoftLayer_Scale_Policy_Trigger_ResourceUse_Extended is SoftLayer_Scale_Policy_Trigger_ResourceUse with all maskable types.
type SoftLayer_Scale_Policy_Trigger_ResourceUse_Extended struct {
	SoftLayer_Scale_Policy_Trigger_ResourceUse

	// Watches - no documentation
	Watches []*SoftLayer_Scale_Policy_Trigger_ResourceUse_Watch `json:"watches"`

	// WatchCount - no documentation
	WatchCount uint64 `json:"watchCount"`
}

func (softlayer_scale_policy_trigger_resourceuse *SoftLayer_Scale_Policy_Trigger_ResourceUse_Extended) String() string {
	return "SoftLayer_Scale_Policy_Trigger_ResourceUse"
}
