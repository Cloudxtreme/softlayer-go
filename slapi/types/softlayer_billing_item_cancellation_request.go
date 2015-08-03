package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Billing_Item_Cancellation_Request - SoftLayer_Billing_Item_Cancellation_Request data type
// is used to cancel service billing items.
type SoftLayer_Billing_Item_Cancellation_Request struct {

	// StatusId - An internal identifier of the service cancellation status that this request is associated
	// with. When a service cancellation is submitted, it will be in "Pending" status until SoftLayer Sales
	// team reviews it. The status of a cancellation request will be updated to "Approved" or "Voided" by
	// SoftLayer Sales. It will be updated to "Complete" when all services are reclaimed.
	StatusId int `json:"statusId"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Id - no documentation
	Id int `json:"id"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// BillingCancelReasonId - no documentation
	BillingCancelReasonId int `json:"billingCancelReasonId"`

	// Notes - no documentation
	Notes string `json:"notes"`

	// AccountId - The internal identifier of the customer account that a service cancellation record
	// belongs to.
	AccountId int `json:"accountId"`

	// TicketId - An internal identifier of the ticket that is associated with a service cancellation
	// request. When a service cancellation is submitted, a support ticket will be created. This ticket
	// contains details on your service cancellation details and SoftLayer Sales team will use it to
	// further communicate with you.
	TicketId int `json:"ticketId"`
}

// SoftLayer_Billing_Item_Cancellation_Request_Extended is SoftLayer_Billing_Item_Cancellation_Request with all maskable types.
type SoftLayer_Billing_Item_Cancellation_Request_Extended struct {
	SoftLayer_Billing_Item_Cancellation_Request

	// Status - no documentation
	Status *SoftLayer_Billing_Item_Cancellation_Request_Status `json:"status"`

	// ItemCount - A count of a collection of service cancellation items.
	ItemCount uint64 `json:"itemCount"`

	// Account - The SoftLayer account that a service cancellation request belongs to.
	Account *SoftLayer_Account `json:"account"`

	// User - The user that initiated a service cancellation request.
	User *SoftLayer_User_Customer `json:"user"`

	// Items - no documentation
	Items []*SoftLayer_Billing_Item_Cancellation_Request_Item `json:"items"`

	// Ticket - The ticket that is associated with the service cancellation request.
	Ticket *SoftLayer_Ticket `json:"ticket"`
}

func (softlayer_billing_item_cancellation_request *SoftLayer_Billing_Item_Cancellation_Request) String() string {
	return "SoftLayer_Billing_Item_Cancellation_Request"
}
