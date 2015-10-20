package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Network_Message_Delivery_Email_Sendgrid_Account_Overview - <nil>
type SoftLayer_Container_Network_Message_Delivery_Email_Sendgrid_Account_Overview struct {

	// Reputation - <nil>
	Reputation int `json:"reputation,omitempty"`

	// Requests - <nil>
	Requests int `json:"requests,omitempty"`

	// CreditsAllowed - <nil>
	CreditsAllowed int `json:"creditsAllowed,omitempty"`

	// CreditsOverage - <nil>
	CreditsOverage int `json:"creditsOverage,omitempty"`

	// CreditsRemain - <nil>
	CreditsRemain int `json:"creditsRemain,omitempty"`

	// CreditsUsed - <nil>
	CreditsUsed int `json:"creditsUsed,omitempty"`

	// Package - <nil>
	Package string `json:"package,omitempty"`
}

func (softlayer_container_network_message_delivery_email_sendgrid_account_overview *SoftLayer_Container_Network_Message_Delivery_Email_Sendgrid_Account_Overview) String() string {
	return "SoftLayer_Container_Network_Message_Delivery_Email_Sendgrid_Account_Overview"
}
