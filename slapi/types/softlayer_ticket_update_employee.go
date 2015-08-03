package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Ticket_Update_Employee - The SoftLayer_Ticket_Update_Employee data type models an update
// to a ticket made by a SoftLayer employee.
type SoftLayer_Ticket_Update_Employee struct {

	// ResponseRating - A ticket update's response rating. Ticket updates posted by SoftLayer employees
	// have the option of earning a rating from SoftLayer's customers. Ratings are based on a 1 - 5 scale,
	// with one being a poor rating while 5 is a very high rating. This is only used if a ticket update's
	// ''editorType'' property is
	ResponseRating int `json:"responseRating"`
}

func (softlayer_ticket_update_employee *SoftLayer_Ticket_Update_Employee) String() string {
	return "SoftLayer_Ticket_Update_Employee"
}
