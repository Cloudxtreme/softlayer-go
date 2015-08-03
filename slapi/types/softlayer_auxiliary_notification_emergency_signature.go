package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Auxiliary_Notification_Emergency_Signature - Every
// SoftLayer_Auxiliary_Notification_Emergency has a signatureId that references a
// SoftLayer_Auxiliary_Notification_Emergency_Signature data type. The signature is the user or group
// responsible for the current event.
type SoftLayer_Auxiliary_Notification_Emergency_Signature struct {

	// Name - The name or signature for the current Emergency Notification.
	Name string `json:"name"`
}

func (softlayer_auxiliary_notification_emergency_signature *SoftLayer_Auxiliary_Notification_Emergency_Signature) String() string {
	return "SoftLayer_Auxiliary_Notification_Emergency_Signature"
}
