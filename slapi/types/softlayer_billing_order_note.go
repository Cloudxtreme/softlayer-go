package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Billing_Order_Note - <nil>
type SoftLayer_Billing_Order_Note struct {

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Employee - <nil>
	Employee *SoftLayer_User_Employee `json:"employee,omitempty"`

	// Order - <nil>
	Order *SoftLayer_Billing_Order `json:"order,omitempty"`
}

func (softlayer_billing_order_note *SoftLayer_Billing_Order_Note) String() string {
	return "SoftLayer_Billing_Order_Note"
}
