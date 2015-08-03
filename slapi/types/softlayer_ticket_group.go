package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Ticket_Group - SoftLayer tickets have the ability to be assigned to one of SoftLayer's
// internal departments. The department that a ticket is assigned to is modeled by the
// SoftLayer_Ticket_Group data type. Ticket groups help to ensure that the proper department is
// handling a ticket. Standard support tickets are created from a number of pre-determined subjects.
// These subjects help determine which group a standard ticket is assigned to.
type SoftLayer_Ticket_Group struct {

	// AssignedBrandCount - no documentation
	AssignedBrandCount uint64 `json:"assignedBrandCount"`

	// AssignedBrands - <nil>
	AssignedBrands []*SoftLayer_Brand `json:"assignedBrands"`

	// Category - no documentation
	Category *SoftLayer_Ticket_Group_Category `json:"category"`

	// Id - no documentation
	Id int `json:"id"`

	// Name - no documentation
	Name string `json:"name"`

	// TicketGroupCategoryId - The internal identifier for the category that a ticket group belongs to..
	TicketGroupCategoryId int `json:"ticketGroupCategoryId"`
}

func (softlayer_ticket_group *SoftLayer_Ticket_Group) String() string {
	return "SoftLayer_Ticket_Group"
}
