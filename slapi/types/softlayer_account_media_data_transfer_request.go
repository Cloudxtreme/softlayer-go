package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Media_Data_Transfer_Request - The SoftLayer_Account_Media_Data_Transfer_Request
// data type contains information on a single Data Transfer Service request. Creation of these requests
// is limited to SoftLayer customers through the SoftLayer Customer Portal.
type SoftLayer_Account_Media_Data_Transfer_Request struct {

	// ModifyUserId - no documentation
	ModifyUserId int `json:"modifyUserId,omitempty"`

	// AccountId - no documentation
	AccountId int `json:"accountId,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// CreateUserId - no documentation
	CreateUserId int `json:"createUserId,omitempty"`

	// EndDate - no documentation
	EndDate *time.Time `json:"endDate,omitempty"`

	// StartDate - no documentation
	StartDate *time.Time `json:"startDate,omitempty"`

	// StatusId - no documentation
	StatusId int `json:"statusId,omitempty"`
}

func (softlayer_account_media_data_transfer_request *SoftLayer_Account_Media_Data_Transfer_Request) String() string {
	return "SoftLayer_Account_Media_Data_Transfer_Request"
}

// SoftLayer_Account_Media_Data_Transfer_Request_Extended is SoftLayer_Account_Media_Data_Transfer_Request with all maskable types.
type SoftLayer_Account_Media_Data_Transfer_Request_Extended struct {
	SoftLayer_Account_Media_Data_Transfer_Request

	// ActiveTickets - The active tickets that are attached to the data transfer request.
	ActiveTickets []*SoftLayer_Ticket `json:"activeTickets,omitempty"`

	// ModifyUser - no documentation
	ModifyUser *SoftLayer_User_Customer `json:"modifyUser,omitempty"`

	// ActiveTicketCount - A count of the active tickets that are attached to the data transfer request.
	ActiveTicketCount uint64 `json:"activeTicketCount,omitempty"`

	// TicketCount - A count of all tickets that are attached to the data transfer request.
	TicketCount uint64 `json:"ticketCount,omitempty"`

	// CreateUser - no documentation
	CreateUser *SoftLayer_User_Customer `json:"createUser,omitempty"`

	// ModifyEmployee - no documentation
	ModifyEmployee *SoftLayer_User_Employee `json:"modifyEmployee,omitempty"`

	// Shipments - no documentation
	Shipments []*SoftLayer_Account_Shipment `json:"shipments,omitempty"`

	// ShipmentCount - no documentation
	ShipmentCount uint64 `json:"shipmentCount,omitempty"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account,omitempty"`

	// Media - no documentation
	Media *SoftLayer_Account_Media `json:"media,omitempty"`

	// Tickets - All tickets that are attached to the data transfer request.
	Tickets []*SoftLayer_Ticket `json:"tickets,omitempty"`

	// BillingItem - no documentation
	BillingItem *SoftLayer_Billing_Item `json:"billingItem,omitempty"`

	// Status - no documentation
	Status *SoftLayer_Account_Media_Data_Transfer_Request_Status `json:"status,omitempty"`
}

func (softlayer_account_media_data_transfer_request *SoftLayer_Account_Media_Data_Transfer_Request_Extended) String() string {
	return "SoftLayer_Account_Media_Data_Transfer_Request"
}
