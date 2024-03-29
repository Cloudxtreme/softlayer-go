package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"

	time "time"
)

// SoftLayer_Product_Upgrade_Request - The SoftLayer_Product_Upgrade_Request data type contains general
// information relating to a hardware, virtual server, or service upgrade. It also relates a
// [[SoftLayer_Billing_Order]] to a [[SoftLayer_Ticket]].
type SoftLayer_Product_Upgrade_Request struct {

	// GuestId - The unique internal id of the virtual server that an upgrade will be done
	GuestId int `json:"guestId,omitempty"`

	// HardwareId - The unique internal id of the hardware that an upgrade will be done
	HardwareId int `json:"hardwareId,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// UserId - The unique internal id of the customer who place the order
	UserId int `json:"userId,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// OrderId - The unique internal id of the order that an upgrade request is related to
	OrderId int `json:"orderId,omitempty"`

	// ProratedTotal - no documentation
	ProratedTotal slapi.Float64 `json:"proratedTotal,omitempty"`

	// StatusId - no documentation
	StatusId int `json:"statusId,omitempty"`

	// MaintenanceStartTimeUtc - The time that system admin starts working on the order item. This is used
	// for upgrade orders.
	MaintenanceStartTimeUtc *time.Time `json:"maintenanceStartTimeUtc,omitempty"`

	// OrderTotal - no documentation
	OrderTotal slapi.Float64 `json:"orderTotal,omitempty"`

	// TicketId - The unique internal id of the ticket related to an upgrade request
	TicketId int `json:"ticketId,omitempty"`

	// AccountId - no documentation
	AccountId int `json:"accountId,omitempty"`

	// EmployeeId - no documentation
	EmployeeId int `json:"employeeId,omitempty"`

	// Order - no documentation
	Order *SoftLayer_Billing_Order `json:"order,omitempty"`

	// Status - no documentation
	Status *SoftLayer_Product_Upgrade_Request_Status `json:"status,omitempty"`

	// Ticket - The ticket that is used to coordinate the upgrade process.
	Ticket *SoftLayer_Ticket `json:"ticket,omitempty"`

	// User - no documentation
	User *SoftLayer_User_Customer `json:"user,omitempty"`

	// VirtualGuest - A virtual server object associated with the upgrade request if any.
	VirtualGuest *SoftLayer_Virtual_Guest `json:"virtualGuest,omitempty"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account,omitempty"`

	// CompletedFlag - Indicates that the upgrade request has completed or has been cancelled.
	CompletedFlag bool `json:"completedFlag,omitempty"`

	// Invoice - This is the invoice associated with the upgrade request. For hourly servers or services,
	// an invoice will not be available.
	Invoice *SoftLayer_Billing_Invoice `json:"invoice,omitempty"`

	// Server - A server object associated with the upgrade request if any.
	Server *SoftLayer_Hardware `json:"server,omitempty"`
}

func (softlayer_product_upgrade_request *SoftLayer_Product_Upgrade_Request) String() string {
	return "SoftLayer_Product_Upgrade_Request"
}
