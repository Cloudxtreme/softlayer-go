package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Network_Message_Delivery_Email_Sendgrid_Statistics - <nil>
type SoftLayer_Container_Network_Message_Delivery_Email_Sendgrid_Statistics struct {

	// Requests - <nil>
	Requests int `json:"requests,omitempty"`

	// Delivered - <nil>
	Delivered int `json:"delivered,omitempty"`

	// RepeatBounces - <nil>
	RepeatBounces int `json:"repeatBounces,omitempty"`

	// InvalidEmail - <nil>
	InvalidEmail int `json:"invalidEmail,omitempty"`

	// RepeatSpamReports - <nil>
	RepeatSpamReports int `json:"repeatSpamReports,omitempty"`

	// RepeatUnsubscribes - <nil>
	RepeatUnsubscribes int `json:"repeatUnsubscribes,omitempty"`

	// UniqueOpens - <nil>
	UniqueOpens int `json:"uniqueOpens,omitempty"`

	// Bounces - <nil>
	Bounces int `json:"bounces,omitempty"`

	// Date - <nil>
	Date string `json:"date,omitempty"`

	// SpamReports - <nil>
	SpamReports int `json:"spamReports,omitempty"`

	// Clicks - <nil>
	Clicks int `json:"clicks,omitempty"`

	// Opens - <nil>
	Opens int `json:"opens,omitempty"`

	// Unsubscribes - <nil>
	Unsubscribes int `json:"unsubscribes,omitempty"`

	// Blocks - <nil>
	Blocks int `json:"blocks,omitempty"`

	// UniqueClicks - <nil>
	UniqueClicks int `json:"uniqueClicks,omitempty"`
}

func (softlayer_container_network_message_delivery_email_sendgrid_statistics *SoftLayer_Container_Network_Message_Delivery_Email_Sendgrid_Statistics) String() string {
	return "SoftLayer_Container_Network_Message_Delivery_Email_Sendgrid_Statistics"
}
