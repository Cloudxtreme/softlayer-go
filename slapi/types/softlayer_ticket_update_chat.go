package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Ticket_Update_Chat - A SoftLayer_Ticket_Update_Chat is a chat between a customer and a
// customer service representative relating to a ticket.
type SoftLayer_Ticket_Update_Chat struct {
}

func (softlayer_ticket_update_chat *SoftLayer_Ticket_Update_Chat) String() string {
	return "SoftLayer_Ticket_Update_Chat"
}

// SoftLayer_Ticket_Update_Chat_Extended is SoftLayer_Ticket_Update_Chat with all maskable types.
type SoftLayer_Ticket_Update_Chat_Extended struct {
	SoftLayer_Ticket_Update_Chat

	// Chat - no documentation
	Chat *SoftLayer_Ticket_Chat_Liveperson `json:"chat"`
}

func (softlayer_ticket_update_chat *SoftLayer_Ticket_Update_Chat_Extended) String() string {
	return "SoftLayer_Ticket_Update_Chat"
}
