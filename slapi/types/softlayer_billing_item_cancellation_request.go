package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Billing_Item_Cancellation_Request - SoftLayer_Billing_Item_Cancellation_Request data type
// is used to cancel service billing items.
type SoftLayer_Billing_Item_Cancellation_Request struct {

	// Account - The SoftLayer account that a service cancellation request belongs to.
	Account *SoftLayer_Account `json:"account"`

	// AccountId - The internal identifier of the customer account that a service cancellation record
	// belongs to.
	AccountId int `json:"accountId"`

	// BillingCancelReasonId - no documentation
	BillingCancelReasonId int `json:"billingCancelReasonId"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Id - no documentation
	Id int `json:"id"`

	// ItemCount - A count of a collection of service cancellation items.
	ItemCount uint64 `json:"itemCount"`

	// Items - no documentation
	Items []*SoftLayer_Billing_Item_Cancellation_Request_Item `json:"items"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// Notes - no documentation
	Notes string `json:"notes"`

	// Status - no documentation
	Status *SoftLayer_Billing_Item_Cancellation_Request_Status `json:"status"`

	// StatusId - An internal identifier of the service cancellation status that this request is associated
	// with. When a service cancellation is submitted, it will be in "Pending" status until SoftLayer Sales
	// team reviews it. The status of a cancellation request will be updated to "Approved" or "Voided" by
	// SoftLayer Sales. It will be updated to "Complete" when all services are reclaimed.
	StatusId int `json:"statusId"`

	// Ticket - The ticket that is associated with the service cancellation request.
	Ticket *SoftLayer_Ticket `json:"ticket"`

	// TicketId - An internal identifier of the ticket that is associated with a service cancellation
	// request. When a service cancellation is submitted, a support ticket will be created. This ticket
	// contains details on your service cancellation details and SoftLayer Sales team will use it to
	// further communicate with you.
	TicketId int `json:"ticketId"`

	// User - The user that initiated a service cancellation request.
	User *SoftLayer_User_Customer `json:"user"`
}

func (softlayer_billing_item_cancellation_request *SoftLayer_Billing_Item_Cancellation_Request) String() string {
	return "SoftLayer_Billing_Item_Cancellation_Request"
}
