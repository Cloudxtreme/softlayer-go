package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Billing_Item_Network_Message_Queue - The SoftLayer_Billing_Item_Network_Message_Queue data
// describes the related billing item.
type SoftLayer_Billing_Item_Network_Message_Queue struct {
}

// SoftLayer_Billing_Item_Network_Message_Queue_Extended is SoftLayer_Billing_Item_Network_Message_Queue with all maskable types.
type SoftLayer_Billing_Item_Network_Message_Queue_Extended struct {
	SoftLayer_Billing_Item_Network_Message_Queue

	// Resource - no documentation
	Resource *SoftLayer_Network_Message_Queue `json:"resource"`
}

func (softlayer_billing_item_network_message_queue *SoftLayer_Billing_Item_Network_Message_Queue) String() string {
	return "SoftLayer_Billing_Item_Network_Message_Queue"
}
