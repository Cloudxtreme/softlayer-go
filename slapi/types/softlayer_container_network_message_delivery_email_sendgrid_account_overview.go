package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Network_Message_Delivery_Email_Sendgrid_Account_Overview - <nil>
type SoftLayer_Container_Network_Message_Delivery_Email_Sendgrid_Account_Overview struct {

	// CreditsAllowed - <nil>
	CreditsAllowed int `json:"creditsAllowed"`

	// CreditsOverage - <nil>
	CreditsOverage int `json:"creditsOverage"`

	// CreditsRemain - <nil>
	CreditsRemain int `json:"creditsRemain"`

	// CreditsUsed - <nil>
	CreditsUsed int `json:"creditsUsed"`

	// Package - <nil>
	Package string `json:"package"`

	// Reputation - <nil>
	Reputation int `json:"reputation"`

	// Requests - <nil>
	Requests int `json:"requests"`
}

func (softlayer_container_network_message_delivery_email_sendgrid_account_overview *SoftLayer_Container_Network_Message_Delivery_Email_Sendgrid_Account_Overview) String() string {
	return "SoftLayer_Container_Network_Message_Delivery_Email_Sendgrid_Account_Overview"
}
