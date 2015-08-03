package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Media_Data_Transfer_Request - The SoftLayer_Account_Media_Data_Transfer_Request
// data type contains information on a single Data Transfer Service request. Creation of these requests
// is limited to SoftLayer customers through the SoftLayer Customer Portal.
type SoftLayer_Account_Media_Data_Transfer_Request struct {

	// StatusId - no documentation
	StatusId int `json:"statusId"`

	// AccountId - no documentation
	AccountId int `json:"accountId"`

	// Id - no documentation
	Id int `json:"id"`

	// CreateUserId - no documentation
	CreateUserId int `json:"createUserId"`

	// EndDate - no documentation
	EndDate *time.Time `json:"endDate"`

	// StartDate - no documentation
	StartDate *time.Time `json:"startDate"`

	// ModifyUserId - no documentation
	ModifyUserId int `json:"modifyUserId"`
}

// SoftLayer_Account_Media_Data_Transfer_Request_Extended is SoftLayer_Account_Media_Data_Transfer_Request with all maskable types.
type SoftLayer_Account_Media_Data_Transfer_Request_Extended struct {
	SoftLayer_Account_Media_Data_Transfer_Request

	// BillingItem - no documentation
	BillingItem *SoftLayer_Billing_Item `json:"billingItem"`

	// CreateUser - no documentation
	CreateUser *SoftLayer_User_Customer `json:"createUser"`

	// TicketCount - A count of all tickets that are attached to the data transfer request.
	TicketCount uint64 `json:"ticketCount"`

	// ActiveTicketCount - A count of the active tickets that are attached to the data transfer request.
	ActiveTicketCount uint64 `json:"activeTicketCount"`

	// ShipmentCount - no documentation
	ShipmentCount uint64 `json:"shipmentCount"`

	// Media - no documentation
	Media *SoftLayer_Account_Media `json:"media"`

	// ModifyEmployee - no documentation
	ModifyEmployee *SoftLayer_User_Employee `json:"modifyEmployee"`

	// Status - no documentation
	Status *SoftLayer_Account_Media_Data_Transfer_Request_Status `json:"status"`

	// ActiveTickets - The active tickets that are attached to the data transfer request.
	ActiveTickets []*SoftLayer_Ticket `json:"activeTickets"`

	// Shipments - no documentation
	Shipments []*SoftLayer_Account_Shipment `json:"shipments"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// ModifyUser - no documentation
	ModifyUser *SoftLayer_User_Customer `json:"modifyUser"`

	// Tickets - All tickets that are attached to the data transfer request.
	Tickets []*SoftLayer_Ticket `json:"tickets"`
}

func (softlayer_account_media_data_transfer_request *SoftLayer_Account_Media_Data_Transfer_Request) String() string {
	return "SoftLayer_Account_Media_Data_Transfer_Request"
}
