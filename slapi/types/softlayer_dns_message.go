package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Dns_Message - The SoftLayer_Dns_Message data type contains information for a single
// message generated by the SoftLayer DNS system. SoftLayer_Dns_Messages are typically created during
// the secondary DNS transfer process.
type SoftLayer_Dns_Message struct {

	// Priority - The priority level for a DNS message. The possible levels are 'notice' and 'error'.
	Priority string `json:"priority,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// Message - no documentation
	Message string `json:"message,omitempty"`
}

func (softlayer_dns_message *SoftLayer_Dns_Message) String() string {
	return "SoftLayer_Dns_Message"
}

// SoftLayer_Dns_Message_Extended is SoftLayer_Dns_Message with all maskable types.
type SoftLayer_Dns_Message_Extended struct {
	SoftLayer_Dns_Message

	// ResourceRecord - The resource record that is associated with a message.
	ResourceRecord *SoftLayer_Dns_Domain_ResourceRecord `json:"resourceRecord,omitempty"`

	// Secondary - The secondary DNS record that a message belongs to.
	Secondary *SoftLayer_Dns_Secondary `json:"secondary,omitempty"`

	// Domain - no documentation
	Domain *SoftLayer_Dns_Domain `json:"domain,omitempty"`
}

func (softlayer_dns_message *SoftLayer_Dns_Message_Extended) String() string {
	return "SoftLayer_Dns_Message"
}
