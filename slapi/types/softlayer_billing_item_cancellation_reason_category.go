package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Billing_Item_Cancellation_Reason_Category - The
// SoftLayer_Billing_Item_Cancellation_Reason_Category data type contains cancellation reason
// categories.
type SoftLayer_Billing_Item_Cancellation_Reason_Category struct {

	// Name - no documentation
	Name string `json:"name"`

	// Id - no documentation
	Id int `json:"id"`
}

// SoftLayer_Billing_Item_Cancellation_Reason_Category_Extended is SoftLayer_Billing_Item_Cancellation_Reason_Category with all maskable types.
type SoftLayer_Billing_Item_Cancellation_Reason_Category_Extended struct {
	SoftLayer_Billing_Item_Cancellation_Reason_Category

	// BillingCancellationReasons - The corresponding billing cancellation reasons having the specific
	// billing cancellation reason category.
	BillingCancellationReasons []*SoftLayer_Billing_Item_Cancellation_Reason `json:"billingCancellationReasons"`

	// BillingCancellationReasonCount - A count of the corresponding billing cancellation reasons having
	// the specific billing cancellation reason category.
	BillingCancellationReasonCount uint64 `json:"billingCancellationReasonCount"`
}

func (softlayer_billing_item_cancellation_reason_category *SoftLayer_Billing_Item_Cancellation_Reason_Category) String() string {
	return "SoftLayer_Billing_Item_Cancellation_Reason_Category"
}
