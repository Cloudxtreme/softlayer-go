package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Ticket_Update_Employee - The SoftLayer_Ticket_Update_Employee data type models an update
// to a ticket made by a SoftLayer employee.
type SoftLayer_Ticket_Update_Employee struct {

	// ResponseRating - A ticket update's response rating. Ticket updates posted by SoftLayer employees
	// have the option of earning a rating from SoftLayer's customers. Ratings are based on a 1 - 5 scale,
	// with one being a poor rating while 5 is a very high rating. This is only used if a ticket update's
	// ''editorType'' property is
	ResponseRating int `json:"responseRating"`
}

// AddResponseRating - As part of the customer service process SoftLayer has provided a quick feedback
// mechanism for its customers to rate the responses that its employees give on tickets.
// addResponseRating() sets the rating for a single ticket update made by a SoftLayer employee. Ticket
// ratings have the integer values 1 through 5, with 1 being the worst and 5 being the best. Once the
// rating is set ''addResponseRating()'' returns a boolean true.
func (softlayer_ticket_update_employee *SoftLayer_Ticket_Update_Employee) AddResponseRating(ctx *slapi.RequestContext, responseRating int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - getObject retrieves the SoftLayer_Ticket_Update_Employee object whose ID number
// corresponds to the ID number of the init parameter passed to the SoftLayer_Ticket_Update_Employee
// service. You can only retrieve employee updates to tickets that your API account has access to.
func (softlayer_ticket_update_employee *SoftLayer_Ticket_Update_Employee) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Ticket_Update_Employee, error) {
	var returnValue *SoftLayer_Ticket_Update_Employee
	return returnValue, nil
}
