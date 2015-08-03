package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Ticket_Subject - The SoftLayer_Ticket_Subject data type models one of the possible
// subjects that a standard support ticket may belong to. A basic support ticket's title matches it's
// corresponding subject's name.
type SoftLayer_Ticket_Subject struct {

	// Group - <nil>
	Group *SoftLayer_Ticket_Group `json:"group"`

	// Id - no documentation
	Id int `json:"id"`

	// Name - A ticket subject's name. This name is used for a standard support ticket's title.
	Name string `json:"name"`
}

func (softlayer_ticket_subject *SoftLayer_Ticket_Subject) String() string {
	return "SoftLayer_Ticket_Subject"
}
