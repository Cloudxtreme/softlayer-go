package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Network_Message_Delivery_Email_Sendgrid_Statistics - <nil>
type SoftLayer_Container_Network_Message_Delivery_Email_Sendgrid_Statistics struct {

	// SpamReports - <nil>
	SpamReports int `json:"spamReports"`

	// UniqueOpens - <nil>
	UniqueOpens int `json:"uniqueOpens"`

	// Unsubscribes - <nil>
	Unsubscribes int `json:"unsubscribes"`

	// Date - <nil>
	Date string `json:"date"`

	// InvalidEmail - <nil>
	InvalidEmail int `json:"invalidEmail"`

	// UniqueClicks - <nil>
	UniqueClicks int `json:"uniqueClicks"`

	// Blocks - <nil>
	Blocks int `json:"blocks"`

	// Delivered - <nil>
	Delivered int `json:"delivered"`

	// Opens - <nil>
	Opens int `json:"opens"`

	// RepeatSpamReports - <nil>
	RepeatSpamReports int `json:"repeatSpamReports"`

	// Requests - <nil>
	Requests int `json:"requests"`

	// Bounces - <nil>
	Bounces int `json:"bounces"`

	// RepeatBounces - <nil>
	RepeatBounces int `json:"repeatBounces"`

	// RepeatUnsubscribes - <nil>
	RepeatUnsubscribes int `json:"repeatUnsubscribes"`

	// Clicks - <nil>
	Clicks int `json:"clicks"`
}

func (softlayer_container_network_message_delivery_email_sendgrid_statistics *SoftLayer_Container_Network_Message_Delivery_Email_Sendgrid_Statistics) String() string {
	return "SoftLayer_Container_Network_Message_Delivery_Email_Sendgrid_Statistics"
}
