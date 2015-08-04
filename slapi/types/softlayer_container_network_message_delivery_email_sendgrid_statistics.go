package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Network_Message_Delivery_Email_Sendgrid_Statistics - <nil>
type SoftLayer_Container_Network_Message_Delivery_Email_Sendgrid_Statistics struct {

	// Clicks - <nil>
	Clicks int `json:"clicks,omitempty"`

	// RepeatUnsubscribes - <nil>
	RepeatUnsubscribes int `json:"repeatUnsubscribes,omitempty"`

	// UniqueOpens - <nil>
	UniqueOpens int `json:"uniqueOpens,omitempty"`

	// UniqueClicks - <nil>
	UniqueClicks int `json:"uniqueClicks,omitempty"`

	// Unsubscribes - <nil>
	Unsubscribes int `json:"unsubscribes,omitempty"`

	// Blocks - <nil>
	Blocks int `json:"blocks,omitempty"`

	// Delivered - <nil>
	Delivered int `json:"delivered,omitempty"`

	// SpamReports - <nil>
	SpamReports int `json:"spamReports,omitempty"`

	// Bounces - <nil>
	Bounces int `json:"bounces,omitempty"`

	// RepeatBounces - <nil>
	RepeatBounces int `json:"repeatBounces,omitempty"`

	// Requests - <nil>
	Requests int `json:"requests,omitempty"`

	// RepeatSpamReports - <nil>
	RepeatSpamReports int `json:"repeatSpamReports,omitempty"`

	// Date - <nil>
	Date string `json:"date,omitempty"`

	// InvalidEmail - <nil>
	InvalidEmail int `json:"invalidEmail,omitempty"`

	// Opens - <nil>
	Opens int `json:"opens,omitempty"`
}

func (softlayer_container_network_message_delivery_email_sendgrid_statistics *SoftLayer_Container_Network_Message_Delivery_Email_Sendgrid_Statistics) String() string {
	return "SoftLayer_Container_Network_Message_Delivery_Email_Sendgrid_Statistics"
}
