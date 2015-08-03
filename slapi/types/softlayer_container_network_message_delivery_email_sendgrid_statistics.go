package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Network_Message_Delivery_Email_Sendgrid_Statistics - <nil>
type SoftLayer_Container_Network_Message_Delivery_Email_Sendgrid_Statistics struct {

	// InvalidEmail - <nil>
	InvalidEmail int `json:"invalidEmail"`

	// Requests - <nil>
	Requests int `json:"requests"`

	// SpamReports - <nil>
	SpamReports int `json:"spamReports"`

	// UniqueClicks - <nil>
	UniqueClicks int `json:"uniqueClicks"`

	// Date - <nil>
	Date string `json:"date"`

	// Delivered - <nil>
	Delivered int `json:"delivered"`

	// RepeatBounces - <nil>
	RepeatBounces int `json:"repeatBounces"`

	// Bounces - <nil>
	Bounces int `json:"bounces"`

	// Clicks - <nil>
	Clicks int `json:"clicks"`

	// UniqueOpens - <nil>
	UniqueOpens int `json:"uniqueOpens"`

	// Unsubscribes - <nil>
	Unsubscribes int `json:"unsubscribes"`

	// Blocks - <nil>
	Blocks int `json:"blocks"`

	// Opens - <nil>
	Opens int `json:"opens"`

	// RepeatSpamReports - <nil>
	RepeatSpamReports int `json:"repeatSpamReports"`

	// RepeatUnsubscribes - <nil>
	RepeatUnsubscribes int `json:"repeatUnsubscribes"`
}

func (softlayer_container_network_message_delivery_email_sendgrid_statistics *SoftLayer_Container_Network_Message_Delivery_Email_Sendgrid_Statistics) String() string {
	return "SoftLayer_Container_Network_Message_Delivery_Email_Sendgrid_Statistics"
}
