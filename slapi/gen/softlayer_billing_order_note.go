package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Billing_Order_Note - <nil>
type SoftLayer_Billing_Order_Note struct {

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// Employee - <nil>
	Employee *SoftLayer_User_Employee `json:"employee"`

	// Order - <nil>
	Order *SoftLayer_Billing_Order `json:"order"`
}
