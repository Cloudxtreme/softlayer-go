package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Media_Data_Transfer_Request - The SoftLayer_Account_Media_Data_Transfer_Request
// data type contains information on a single Data Transfer Service request. Creation of these requests
// is limited to SoftLayer customers through the SoftLayer Customer Portal.
type SoftLayer_Account_Media_Data_Transfer_Request struct {

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// AccountId - no documentation
	AccountId int `json:"accountId"`

	// ActiveTicketCount - A count of the active tickets that are attached to the data transfer request.
	ActiveTicketCount uint64 `json:"activeTicketCount"`

	// ActiveTickets - The active tickets that are attached to the data transfer request.
	ActiveTickets []*SoftLayer_Ticket `json:"activeTickets"`

	// BillingItem - no documentation
	BillingItem *SoftLayer_Billing_Item `json:"billingItem"`

	// CreateUser - no documentation
	CreateUser *SoftLayer_User_Customer `json:"createUser"`

	// CreateUserId - no documentation
	CreateUserId int `json:"createUserId"`

	// EndDate - no documentation
	EndDate *time.Time `json:"endDate"`

	// Id - no documentation
	Id int `json:"id"`

	// Media - no documentation
	Media *SoftLayer_Account_Media `json:"media"`

	// ModifyEmployee - no documentation
	ModifyEmployee *SoftLayer_User_Employee `json:"modifyEmployee"`

	// ModifyUser - no documentation
	ModifyUser *SoftLayer_User_Customer `json:"modifyUser"`

	// ModifyUserId - no documentation
	ModifyUserId int `json:"modifyUserId"`

	// ShipmentCount - no documentation
	ShipmentCount uint64 `json:"shipmentCount"`

	// Shipments - no documentation
	Shipments []*SoftLayer_Account_Shipment `json:"shipments"`

	// StartDate - no documentation
	StartDate *time.Time `json:"startDate"`

	// Status - no documentation
	Status *SoftLayer_Account_Media_Data_Transfer_Request_Status `json:"status"`

	// StatusId - no documentation
	StatusId int `json:"statusId"`

	// TicketCount - A count of all tickets that are attached to the data transfer request.
	TicketCount uint64 `json:"ticketCount"`

	// Tickets - All tickets that are attached to the data transfer request.
	Tickets []*SoftLayer_Ticket `json:"tickets"`
}

// EditObject - Edit the properties of a data transfer request record by passing in a modified instance
// of a SoftLayer_Account_Media_Data_Transfer_Request object.
func (softlayer_account_media_data_transfer_request *SoftLayer_Account_Media_Data_Transfer_Request) EditObject(templateObject SoftLayer_Account_Media_Data_Transfer_Request) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetAllRequestStatuses - Retrieves a list of all the possible statuses to which a request may be set.
func (softlayer_account_media_data_transfer_request *SoftLayer_Account_Media_Data_Transfer_Request) GetAllRequestStatuses() ([]*SoftLayer_Account_Media_Data_Transfer_Request_Status, error) {
	var returnValue []*SoftLayer_Account_Media_Data_Transfer_Request_Status
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_account_media_data_transfer_request *SoftLayer_Account_Media_Data_Transfer_Request) GetObject() (*SoftLayer_Account_Media_Data_Transfer_Request, error) {
	var returnValue *SoftLayer_Account_Media_Data_Transfer_Request
	return returnValue, nil
}
