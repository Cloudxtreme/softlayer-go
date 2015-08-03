package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Network_Message_Delivery_Email_Sendgrid_Statistics - <nil>
type SoftLayer_Container_Network_Message_Delivery_Email_Sendgrid_Statistics struct {

	// Blocks - <nil>
	Blocks int `json:"blocks"`

	// Bounces - <nil>
	Bounces int `json:"bounces"`

	// Clicks - <nil>
	Clicks int `json:"clicks"`

	// Date - <nil>
	Date string `json:"date"`

	// Delivered - <nil>
	Delivered int `json:"delivered"`

	// InvalidEmail - <nil>
	InvalidEmail int `json:"invalidEmail"`

	// Opens - <nil>
	Opens int `json:"opens"`

	// RepeatBounces - <nil>
	RepeatBounces int `json:"repeatBounces"`

	// RepeatSpamReports - <nil>
	RepeatSpamReports int `json:"repeatSpamReports"`

	// RepeatUnsubscribes - <nil>
	RepeatUnsubscribes int `json:"repeatUnsubscribes"`

	// Requests - <nil>
	Requests int `json:"requests"`

	// SpamReports - <nil>
	SpamReports int `json:"spamReports"`

	// UniqueClicks - <nil>
	UniqueClicks int `json:"uniqueClicks"`

	// UniqueOpens - <nil>
	UniqueOpens int `json:"uniqueOpens"`

	// Unsubscribes - <nil>
	Unsubscribes int `json:"unsubscribes"`
}

func (softlayer_container_network_message_delivery_email_sendgrid_statistics *SoftLayer_Container_Network_Message_Delivery_Email_Sendgrid_Statistics) String() string {
	return "SoftLayer_Container_Network_Message_Delivery_Email_Sendgrid_Statistics"
}
