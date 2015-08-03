package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Ticket_Group_Category - SoftLayer's support ticket groups represent the department at
// SoftLayer that is assigned to work one of your support tickets. Many departments are responsible for
// handling different types of tickets. These types of tickets are modeled in the
// SoftLayer_Ticket_Group_Category data type. Ticket group categories also help separate differentiate
// your tickets' issues in the SoftLayer customer portal.
type SoftLayer_Ticket_Group_Category struct {

	// Id - no documentation
	Id int `json:"id"`

	// Name - no documentation
	Name string `json:"name"`
}

func (softlayer_ticket_group_category *SoftLayer_Ticket_Group_Category) String() string {
	return "SoftLayer_Ticket_Group_Category"
}
