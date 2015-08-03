package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Scale_Policy_Trigger_Repeating - <nil>
type SoftLayer_Scale_Policy_Trigger_Repeating struct {

	// Schedule - The cron-formatted schedule. This is run in the UTC timezone.
	Schedule string `json:"schedule"`
}

func (softlayer_scale_policy_trigger_repeating *SoftLayer_Scale_Policy_Trigger_Repeating) String() string {
	return "SoftLayer_Scale_Policy_Trigger_Repeating"
}
