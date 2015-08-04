package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Network_Message_Delivery_Email - <nil>
type SoftLayer_Container_Network_Message_Delivery_Email struct {

	// ContainsHtml - <nil>
	ContainsHtml bool `json:"containsHtml,omitempty"`

	// From - <nil>
	From string `json:"from,omitempty"`

	// Subject - <nil>
	Subject string `json:"subject,omitempty"`

	// To - <nil>
	To string `json:"to,omitempty"`

	// Body - <nil>
	Body string `json:"body,omitempty"`
}

func (softlayer_container_network_message_delivery_email *SoftLayer_Container_Network_Message_Delivery_Email) String() string {
	return "SoftLayer_Container_Network_Message_Delivery_Email"
}
