package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Ticket_Update_Chat - A SoftLayer_Ticket_Update_Chat is a chat between a customer and a
// customer service representative relating to a ticket.
type SoftLayer_Ticket_Update_Chat struct {

	// Chat - no documentation
	Chat *SoftLayer_Ticket_Chat_Liveperson `json:"chat"`
}

func (softlayer_ticket_update_chat *SoftLayer_Ticket_Update_Chat) String() string {
	return "SoftLayer_Ticket_Update_Chat"
}
