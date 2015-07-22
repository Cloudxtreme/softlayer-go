package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Billing_Item_Cancellation_Reason - The SoftLayer_Billing_Item_Cancellation_Reason data
// type contains cancellation reasons.
type SoftLayer_Billing_Item_Cancellation_Reason struct {

	// BillingCancelReasonCategoryId - no documentation
	BillingCancelReasonCategoryId int `json:"billingCancelReasonCategoryId"`

	// BillingCancellationReasonCategory - no documentation
	BillingCancellationReasonCategory *SoftLayer_Billing_Item_Cancellation_Reason_Category `json:"billingCancellationReasonCategory"`

	// BillingItemCount - A count of the corresponding billing items having the specific cancellation
	// reason.
	BillingItemCount uint64 `json:"billingItemCount"`

	// BillingItems - The corresponding billing items having the specific cancellation reason.
	BillingItems []*SoftLayer_Billing_Item `json:"billingItems"`

	// Id - no documentation
	Id int `json:"id"`

	// KeyName - no documentation
	KeyName string `json:"keyName"`

	// Reason - no documentation
	Reason string `json:"reason"`

	// TranslatedReason - <nil>
	TranslatedReason string `json:"translatedReason"`
}

// GetAllCancellationReasons - getAllCancellationReasons() retrieves a list of all cancellation reasons
// that a server/service may be assigned to.
func (softlayer_billing_item_cancellation_reason *SoftLayer_Billing_Item_Cancellation_Reason) GetAllCancellationReasons(commonOptions *slapi.CommonOptions) ([]*SoftLayer_Billing_Item_Cancellation_Reason, error) {
	var returnValue []*SoftLayer_Billing_Item_Cancellation_Reason
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_billing_item_cancellation_reason *SoftLayer_Billing_Item_Cancellation_Reason) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Billing_Item_Cancellation_Reason, error) {
	var returnValue *SoftLayer_Billing_Item_Cancellation_Reason
	return returnValue, nil
}
