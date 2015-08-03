package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Message_Delivery_Email_Sendgrid - <nil>
type SoftLayer_Network_Message_Delivery_Email_Sendgrid struct {
}

func (softlayer_network_message_delivery_email_sendgrid *SoftLayer_Network_Message_Delivery_Email_Sendgrid) String() string {
	return "SoftLayer_Network_Message_Delivery_Email_Sendgrid"
}

// SoftLayer_Network_Message_Delivery_Email_Sendgrid_Extended is SoftLayer_Network_Message_Delivery_Email_Sendgrid with all maskable types.
type SoftLayer_Network_Message_Delivery_Email_Sendgrid_Extended struct {
	SoftLayer_Network_Message_Delivery_Email_Sendgrid

	// EmailAddress - no documentation
	EmailAddress string `json:"emailAddress"`

	// SmtpAccess - A flag that determines if a SendGrid e-mail delivery account has access to send mail
	// through the SendGrid server.
	SmtpAccess string `json:"smtpAccess"`
}

func (softlayer_network_message_delivery_email_sendgrid *SoftLayer_Network_Message_Delivery_Email_Sendgrid_Extended) String() string {
	return "SoftLayer_Network_Message_Delivery_Email_Sendgrid"
}
