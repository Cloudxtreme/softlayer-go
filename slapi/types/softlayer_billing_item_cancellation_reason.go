package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Billing_Item_Cancellation_Reason - The SoftLayer_Billing_Item_Cancellation_Reason data
// type contains cancellation reasons.
type SoftLayer_Billing_Item_Cancellation_Reason struct {

	// BillingCancelReasonCategoryId - no documentation
	BillingCancelReasonCategoryId int `json:"billingCancelReasonCategoryId"`

	// Id - no documentation
	Id int `json:"id"`

	// KeyName - no documentation
	KeyName string `json:"keyName"`

	// Reason - no documentation
	Reason string `json:"reason"`
}

// SoftLayer_Billing_Item_Cancellation_Reason_Extended is SoftLayer_Billing_Item_Cancellation_Reason with all maskable types.
type SoftLayer_Billing_Item_Cancellation_Reason_Extended struct {
	SoftLayer_Billing_Item_Cancellation_Reason

	// BillingItemCount - A count of the corresponding billing items having the specific cancellation
	// reason.
	BillingItemCount uint64 `json:"billingItemCount"`

	// BillingCancellationReasonCategory - no documentation
	BillingCancellationReasonCategory *SoftLayer_Billing_Item_Cancellation_Reason_Category `json:"billingCancellationReasonCategory"`

	// BillingItems - The corresponding billing items having the specific cancellation reason.
	BillingItems []*SoftLayer_Billing_Item `json:"billingItems"`

	// TranslatedReason - <nil>
	TranslatedReason string `json:"translatedReason"`
}

func (softlayer_billing_item_cancellation_reason *SoftLayer_Billing_Item_Cancellation_Reason) String() string {
	return "SoftLayer_Billing_Item_Cancellation_Reason"
}
