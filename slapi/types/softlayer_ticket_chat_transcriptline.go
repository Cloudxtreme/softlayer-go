package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Ticket_Chat_TranscriptLine - <nil>
type SoftLayer_Ticket_Chat_TranscriptLine struct {
}

// SoftLayer_Ticket_Chat_TranscriptLine_Extended is SoftLayer_Ticket_Chat_TranscriptLine with all maskable types.
type SoftLayer_Ticket_Chat_TranscriptLine_Extended struct {
	SoftLayer_Ticket_Chat_TranscriptLine

	// Speaker - <nil>
	Speaker *SoftLayer_User_Interface `json:"speaker"`
}

func (softlayer_ticket_chat_transcriptline *SoftLayer_Ticket_Chat_TranscriptLine) String() string {
	return "SoftLayer_Ticket_Chat_TranscriptLine"
}
