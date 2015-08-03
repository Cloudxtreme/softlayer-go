package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Product_Upgrade_Request - The SoftLayer_Product_Upgrade_Request data type contains general
// information relating to a hardware, virtual server, or service upgrade. It also relates a
// [[SoftLayer_Billing_Order]] to a [[SoftLayer_Ticket]].
type SoftLayer_Product_Upgrade_Request struct {

	// ProratedTotal - no documentation
	ProratedTotal float64 `json:"proratedTotal"`

	// EmployeeId - no documentation
	EmployeeId int `json:"employeeId"`

	// OrderTotal - no documentation
	OrderTotal float64 `json:"orderTotal"`

	// AccountId - no documentation
	AccountId int `json:"accountId"`

	// GuestId - The unique internal id of the virtual server that an upgrade will be done
	GuestId int `json:"guestId"`

	// HardwareId - The unique internal id of the hardware that an upgrade will be done
	HardwareId int `json:"hardwareId"`

	// Id - no documentation
	Id int `json:"id"`

	// UserId - The unique internal id of the customer who place the order
	UserId int `json:"userId"`

	// MaintenanceStartTimeUtc - The time that system admin starts working on the order item. This is used
	// for upgrade orders.
	MaintenanceStartTimeUtc *time.Time `json:"maintenanceStartTimeUtc"`

	// StatusId - no documentation
	StatusId int `json:"statusId"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// OrderId - The unique internal id of the order that an upgrade request is related to
	OrderId int `json:"orderId"`

	// TicketId - The unique internal id of the ticket related to an upgrade request
	TicketId int `json:"ticketId"`
}

func (softlayer_product_upgrade_request *SoftLayer_Product_Upgrade_Request) String() string {
	return "SoftLayer_Product_Upgrade_Request"
}

// SoftLayer_Product_Upgrade_Request_Extended is SoftLayer_Product_Upgrade_Request with all maskable types.
type SoftLayer_Product_Upgrade_Request_Extended struct {
	SoftLayer_Product_Upgrade_Request

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// Ticket - The ticket that is used to coordinate the upgrade process.
	Ticket *SoftLayer_Ticket `json:"ticket"`

	// CompletedFlag - Indicates that the upgrade request has completed or has been cancelled.
	CompletedFlag bool `json:"completedFlag"`

	// Order - no documentation
	Order *SoftLayer_Billing_Order `json:"order"`

	// Server - A server object associated with the upgrade request if any.
	Server *SoftLayer_Hardware `json:"server"`

	// Status - no documentation
	Status *SoftLayer_Product_Upgrade_Request_Status `json:"status"`

	// User - no documentation
	User *SoftLayer_User_Customer `json:"user"`

	// Invoice - This is the invoice associated with the upgrade request. For hourly servers or services,
	// an invoice will not be available.
	Invoice *SoftLayer_Billing_Invoice `json:"invoice"`

	// VirtualGuest - A virtual server object associated with the upgrade request if any.
	VirtualGuest *SoftLayer_Virtual_Guest `json:"virtualGuest"`
}

func (softlayer_product_upgrade_request *SoftLayer_Product_Upgrade_Request_Extended) String() string {
	return "SoftLayer_Product_Upgrade_Request"
}
