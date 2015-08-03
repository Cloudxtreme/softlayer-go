package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Ticket_Survey - <nil>
type SoftLayer_Ticket_Survey struct {
}

func (softlayer_ticket_survey *SoftLayer_Ticket_Survey) String() string {
	return "SoftLayer_Ticket_Survey"
}

// GetPreference - Use this method to retrieve the ticket survey preferences. It will return your
// [[SoftLayer_Container_Ticket_Survey_Preference|survey preference]] which indicates if your account
// is applicable to receive a survey and if you're opted in. You can control the survey opt via the
// [[SoftLayer_Ticket_Survey::optIn|opt-in]] or [[SoftLayer_Ticket_Survey::optOut|opt-out]] method.
func (softlayer_ticket_survey *SoftLayer_Ticket_Survey) GetPreference(ctx *slapi.RequestContext) (*SoftLayer_Container_Ticket_Survey_Preference, error) {
	var returnValue *SoftLayer_Container_Ticket_Survey_Preference
	return returnValue, nil
}

// OptIn - You will not receive a ticket survey if you are opted out. Use this method to opt back in if
// you wish to provide feedback to our support team. You may use the
// [[SoftLayer_Ticket_Survey::getPreference|getPreference]] method to check your current opt status.
func (softlayer_ticket_survey *SoftLayer_Ticket_Survey) OptIn(ctx *slapi.RequestContext) (*SoftLayer_Container_Ticket_Survey_Preference, error) {
	var returnValue *SoftLayer_Container_Ticket_Survey_Preference
	return returnValue, nil
}

// OptOut - By default, customers will occasionally receive a ticket survey upon closing of a ticket.
// Use this method to opt out of it for the next 90 days. Ticket surveys may not be applicable for some
// customers. Use the [[SoftLayer_Ticket_Survey::getPreference|getPreference]] method to retrieve your
// survey preference. The "applicable" property of the
// [[SoftLayer_Container_Ticket_Survey_Preference|survey preference]] indicates if the survey is
// relevant to your account or not.
func (softlayer_ticket_survey *SoftLayer_Ticket_Survey) OptOut(ctx *slapi.RequestContext) (*SoftLayer_Container_Ticket_Survey_Preference, error) {
	var returnValue *SoftLayer_Container_Ticket_Survey_Preference
	return returnValue, nil
}
