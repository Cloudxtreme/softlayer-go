package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Message_Delivery_Email_Sendgrid - <nil>
type SoftLayer_Network_Message_Delivery_Email_Sendgrid struct {

	// EmailAddress - no documentation
	EmailAddress string `json:"emailAddress"`

	// SmtpAccess - A flag that determines if a SendGrid e-mail delivery account has access to send mail
	// through the SendGrid server.
	SmtpAccess string `json:"smtpAccess"`
}

// AddUnsubscribeEmailAddress - <nil>
func (softlayer_network_message_delivery_email_sendgrid *SoftLayer_Network_Message_Delivery_Email_Sendgrid) AddUnsubscribeEmailAddress(emailAddress string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// DeleteEmailListEntries - <nil>
func (softlayer_network_message_delivery_email_sendgrid *SoftLayer_Network_Message_Delivery_Email_Sendgrid) DeleteEmailListEntries(list string, entries []string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// DisableSmtpAccess - <nil>
func (softlayer_network_message_delivery_email_sendgrid *SoftLayer_Network_Message_Delivery_Email_Sendgrid) DisableSmtpAccess() (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EnableSmtpAccess - <nil>
func (softlayer_network_message_delivery_email_sendgrid *SoftLayer_Network_Message_Delivery_Email_Sendgrid) EnableSmtpAccess() (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetAccountOverview - <nil>
func (softlayer_network_message_delivery_email_sendgrid *SoftLayer_Network_Message_Delivery_Email_Sendgrid) GetAccountOverview() (*SoftLayer_Container_Network_Message_Delivery_Email_Sendgrid_Account_Overview, error) {
	var returnValue *SoftLayer_Container_Network_Message_Delivery_Email_Sendgrid_Account_Overview
	return returnValue, nil
}

// GetCategoryList - <nil>
func (softlayer_network_message_delivery_email_sendgrid *SoftLayer_Network_Message_Delivery_Email_Sendgrid) GetCategoryList() ([]string, error) {
	var returnValue []string
	return returnValue, nil
}

// GetEmailList - <nil>
func (softlayer_network_message_delivery_email_sendgrid *SoftLayer_Network_Message_Delivery_Email_Sendgrid) GetEmailList(list string) ([]*SoftLayer_Container_Network_Message_Delivery_Email_Sendgrid_List_Entry, error) {
	var returnValue []*SoftLayer_Container_Network_Message_Delivery_Email_Sendgrid_List_Entry
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_network_message_delivery_email_sendgrid *SoftLayer_Network_Message_Delivery_Email_Sendgrid) GetObject() (*SoftLayer_Network_Message_Delivery_Email_Sendgrid, error) {
	var returnValue *SoftLayer_Network_Message_Delivery_Email_Sendgrid
	return returnValue, nil
}

// GetStatistics - <nil>
func (softlayer_network_message_delivery_email_sendgrid *SoftLayer_Network_Message_Delivery_Email_Sendgrid) GetStatistics(options SoftLayer_Container_Network_Message_Delivery_Email_Sendgrid_Statistics_Options) ([]*SoftLayer_Container_Network_Message_Delivery_Email_Sendgrid_Statistics, error) {
	var returnValue []*SoftLayer_Container_Network_Message_Delivery_Email_Sendgrid_Statistics
	return returnValue, nil
}

// GetStatisticsGraph - <nil>
func (softlayer_network_message_delivery_email_sendgrid *SoftLayer_Network_Message_Delivery_Email_Sendgrid) GetStatisticsGraph(options SoftLayer_Container_Network_Message_Delivery_Email_Sendgrid_Statistics_Options) (*SoftLayer_Container_Network_Message_Delivery_Email_Sendgrid_Statistics_Graph, error) {
	var returnValue *SoftLayer_Container_Network_Message_Delivery_Email_Sendgrid_Statistics_Graph
	return returnValue, nil
}

// GetVendorPortalUrl - <nil>
func (softlayer_network_message_delivery_email_sendgrid *SoftLayer_Network_Message_Delivery_Email_Sendgrid) GetVendorPortalUrl() (string, error) {
	var returnValue string
	return returnValue, nil
}

// SendEmail - <nil>
func (softlayer_network_message_delivery_email_sendgrid *SoftLayer_Network_Message_Delivery_Email_Sendgrid) SendEmail(emailContainer SoftLayer_Container_Network_Message_Delivery_Email) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// UpdateEmailAddress - <nil>
func (softlayer_network_message_delivery_email_sendgrid *SoftLayer_Network_Message_Delivery_Email_Sendgrid) UpdateEmailAddress(emailAddress string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}
