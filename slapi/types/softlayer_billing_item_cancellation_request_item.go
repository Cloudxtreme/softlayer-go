package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Billing_Item_Cancellation_Request_Item - SoftLayer_Billing_Item_Cancellation_Request_Item
// data type contains a billing item for cancellation. This data type is used to harness billing items
// to the associated service.
type SoftLayer_Billing_Item_Cancellation_Request_Item struct {

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// ImmediateCancellationFlag - This flag indicated if a billing item should be canceled immediately or
	// not. Set this flag to true when creating a cancellation request.
	ImmediateCancellationFlag bool `json:"immediateCancellationFlag,omitempty"`

	// ScheduledCancellationDate - no documentation
	ScheduledCancellationDate *time.Time `json:"scheduledCancellationDate,omitempty"`

	// ServiceReclaimStatusCode - no documentation
	ServiceReclaimStatusCode string `json:"serviceReclaimStatusCode,omitempty"`

	// BillingItemId - no documentation
	BillingItemId int `json:"billingItemId,omitempty"`

	// CancellationRequestId - no documentation
	CancellationRequestId int `json:"cancellationRequestId,omitempty"`

	// BillingItem - no documentation
	BillingItem *SoftLayer_Billing_Item `json:"billingItem,omitempty"`

	// CancellationRequest - The service cancellation request that a cancellation item belongs to.
	CancellationRequest *SoftLayer_Billing_Item_Cancellation_Request `json:"cancellationRequest,omitempty"`
}

func (softlayer_billing_item_cancellation_request_item *SoftLayer_Billing_Item_Cancellation_Request_Item) String() string {
	return "SoftLayer_Billing_Item_Cancellation_Request_Item"
}
