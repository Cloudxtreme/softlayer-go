package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
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

// CreateObject - This method creates a service cancellation request. You need to have "Cancel
// Services" privilege to create a cancellation request. You have to provide at least one
// SoftLayer_Billing_Item_Cancellation_Request_Item in the "items" property. Make sure billing item's
// category code belongs to the cancelable product codes. You can retrieve the cancelable product
// category by the [[SoftLayer_Product_Item_Category::getValidCancelableServiceItemCategories|product
// category]] service.
func (softlayer_billing_item_cancellation_request *SoftLayer_Billing_Item_Cancellation_Request) CreateObject(ctx *slapi.RequestContext, templateObject SoftLayer_Billing_Item_Cancellation_Request) (*SoftLayer_Billing_Item_Cancellation_Request, error) {
	var returnValue *SoftLayer_Billing_Item_Cancellation_Request
	return returnValue, nil
}

// GetAllCancellationRequests - This method returns all service cancellation requests. Make sure to
// include the "resultLimit" in the request header for quicker response. If there is no result limit
// header is passed, it will return the latest 25 results by default.
func (softlayer_billing_item_cancellation_request *SoftLayer_Billing_Item_Cancellation_Request) GetAllCancellationRequests(ctx *slapi.RequestContext) ([]*SoftLayer_Billing_Item_Cancellation_Request, error) {
	var returnValue []*SoftLayer_Billing_Item_Cancellation_Request
	return returnValue, nil
}

// GetCancellationCutoffDate - Services can be canceled 2 or 3 days prior to your next bill date. This
// service returns the time by which a cancellation request submission is permitted in the current
// billing cycle. If the current time falls into the cut off date, this will return next earliest
// cancellation cut off date. Available category codes are: service, server
func (softlayer_billing_item_cancellation_request *SoftLayer_Billing_Item_Cancellation_Request) GetCancellationCutoffDate(ctx *slapi.RequestContext, accountId int, categoryCode string) (*time.Time, error) {
	var returnValue *time.Time
	return returnValue, nil
}

// GetObject - getObject retrieves the SoftLayer_Billing_Item_Cancellation_Request object whose ID
// number corresponds to the ID number of the init parameter passed to the
// SoftLayer_Billing_Item_Cancellation_Request service. You can only retrieve cancellation request
// records that are assigned to your SoftLayer account.
func (softlayer_billing_item_cancellation_request *SoftLayer_Billing_Item_Cancellation_Request) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Billing_Item_Cancellation_Request, error) {
	var returnValue *SoftLayer_Billing_Item_Cancellation_Request
	return returnValue, nil
}

// RemoveCancellationItem - This method removes a cancellation item from a cancellation request that is
// in "Pending" or "Approved" status.
func (softlayer_billing_item_cancellation_request *SoftLayer_Billing_Item_Cancellation_Request) RemoveCancellationItem(ctx *slapi.RequestContext, itemId int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// ValidateBillingItemForCancellation - This method examined if a billing item is eligible for
// cancellation. It checks if the billing item you provided is already in your existing cancellation
// request.
func (softlayer_billing_item_cancellation_request *SoftLayer_Billing_Item_Cancellation_Request) ValidateBillingItemForCancellation(ctx *slapi.RequestContext, billingItemId int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// Void - This method voids a service cancellation request in "Pending" or "Approved" status.
func (softlayer_billing_item_cancellation_request *SoftLayer_Billing_Item_Cancellation_Request) Void(ctx *slapi.RequestContext, closeRelatedTicketFlag bool) (bool, error) {
	var returnValue bool
	return returnValue, nil
}
