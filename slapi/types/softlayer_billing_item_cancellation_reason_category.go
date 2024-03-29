package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Billing_Item_Cancellation_Reason_Category - The
// SoftLayer_Billing_Item_Cancellation_Reason_Category data type contains cancellation reason
// categories.
type SoftLayer_Billing_Item_Cancellation_Reason_Category struct {

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// BillingCancellationReasonCount - A count of the corresponding billing cancellation reasons having
	// the specific billing cancellation reason category.
	BillingCancellationReasonCount uint64 `json:"billingCancellationReasonCount,omitempty"`

	// BillingCancellationReasons - The corresponding billing cancellation reasons having the specific
	// billing cancellation reason category.
	BillingCancellationReasons []*SoftLayer_Billing_Item_Cancellation_Reason `json:"billingCancellationReasons,omitempty"`
}

func (softlayer_billing_item_cancellation_reason_category *SoftLayer_Billing_Item_Cancellation_Reason_Category) String() string {
	return "SoftLayer_Billing_Item_Cancellation_Reason_Category"
}
