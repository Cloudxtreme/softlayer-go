package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Billing_Item_Cancellation_Reason - The SoftLayer_Billing_Item_Cancellation_Reason data
// type contains cancellation reasons.
type SoftLayer_Billing_Item_Cancellation_Reason struct {

	// BillingCancelReasonCategoryId - no documentation
	BillingCancelReasonCategoryId int `json:"billingCancelReasonCategoryId,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// KeyName - no documentation
	KeyName string `json:"keyName,omitempty"`

	// Reason - no documentation
	Reason string `json:"reason,omitempty"`
}

func (softlayer_billing_item_cancellation_reason *SoftLayer_Billing_Item_Cancellation_Reason) String() string {
	return "SoftLayer_Billing_Item_Cancellation_Reason"
}

// SoftLayer_Billing_Item_Cancellation_Reason_Extended is SoftLayer_Billing_Item_Cancellation_Reason with all maskable types.
type SoftLayer_Billing_Item_Cancellation_Reason_Extended struct {
	SoftLayer_Billing_Item_Cancellation_Reason

	// BillingItems - The corresponding billing items having the specific cancellation reason.
	BillingItems []*SoftLayer_Billing_Item `json:"billingItems,omitempty"`

	// TranslatedReason - <nil>
	TranslatedReason string `json:"translatedReason,omitempty"`

	// BillingItemCount - A count of the corresponding billing items having the specific cancellation
	// reason.
	BillingItemCount uint64 `json:"billingItemCount,omitempty"`

	// BillingCancellationReasonCategory - no documentation
	BillingCancellationReasonCategory *SoftLayer_Billing_Item_Cancellation_Reason_Category `json:"billingCancellationReasonCategory,omitempty"`
}

func (softlayer_billing_item_cancellation_reason *SoftLayer_Billing_Item_Cancellation_Reason_Extended) String() string {
	return "SoftLayer_Billing_Item_Cancellation_Reason"
}
