package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Billing_Order_Note - <nil>
type SoftLayer_Billing_Order_Note struct {

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate,omitempty"`
}

func (softlayer_billing_order_note *SoftLayer_Billing_Order_Note) String() string {
	return "SoftLayer_Billing_Order_Note"
}

// SoftLayer_Billing_Order_Note_Extended is SoftLayer_Billing_Order_Note with all maskable types.
type SoftLayer_Billing_Order_Note_Extended struct {
	SoftLayer_Billing_Order_Note

	// Employee - <nil>
	Employee *SoftLayer_User_Employee `json:"employee,omitempty"`

	// Order - <nil>
	Order *SoftLayer_Billing_Order `json:"order,omitempty"`
}

func (softlayer_billing_order_note *SoftLayer_Billing_Order_Note_Extended) String() string {
	return "SoftLayer_Billing_Order_Note"
}
