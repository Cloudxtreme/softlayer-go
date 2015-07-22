package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Product_Upgrade_Request - The SoftLayer_Product_Upgrade_Request data type contains general
// information relating to a hardware, virtual server, or service upgrade. It also relates a
// [[SoftLayer_Billing_Order]] to a [[SoftLayer_Ticket]].
type SoftLayer_Product_Upgrade_Request struct {

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// AccountId - no documentation
	AccountId int `json:"accountId"`

	// CompletedFlag - Indicates that the upgrade request has completed or has been cancelled.
	CompletedFlag bool `json:"completedFlag"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// EmployeeId - no documentation
	EmployeeId int `json:"employeeId"`

	// GuestId - The unique internal id of the virtual server that an upgrade will be done
	GuestId int `json:"guestId"`

	// HardwareId - The unique internal id of the hardware that an upgrade will be done
	HardwareId int `json:"hardwareId"`

	// Id - no documentation
	Id int `json:"id"`

	// Invoice - This is the invoice associated with the upgrade request. For hourly servers or services,
	// an invoice will not be available.
	Invoice *SoftLayer_Billing_Invoice `json:"invoice"`

	// MaintenanceStartTimeUtc - The time that system admin starts working on the order item. This is used
	// for upgrade orders.
	MaintenanceStartTimeUtc *time.Time `json:"maintenanceStartTimeUtc"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// Order - no documentation
	Order *SoftLayer_Billing_Order `json:"order"`

	// OrderId - The unique internal id of the order that an upgrade request is related to
	OrderId int `json:"orderId"`

	// OrderTotal - no documentation
	OrderTotal float64 `json:"orderTotal"`

	// ProratedTotal - no documentation
	ProratedTotal float64 `json:"proratedTotal"`

	// Server - A server object associated with the upgrade request if any.
	Server *SoftLayer_Hardware `json:"server"`

	// Status - no documentation
	Status *SoftLayer_Product_Upgrade_Request_Status `json:"status"`

	// StatusId - no documentation
	StatusId int `json:"statusId"`

	// Ticket - The ticket that is used to coordinate the upgrade process.
	Ticket *SoftLayer_Ticket `json:"ticket"`

	// TicketId - The unique internal id of the ticket related to an upgrade request
	TicketId int `json:"ticketId"`

	// User - no documentation
	User *SoftLayer_User_Customer `json:"user"`

	// UserId - The unique internal id of the customer who place the order
	UserId int `json:"userId"`

	// VirtualGuest - A virtual server object associated with the upgrade request if any.
	VirtualGuest *SoftLayer_Virtual_Guest `json:"virtualGuest"`
}

// ApproveChanges - When a change is made to an upgrade by Sales, this method will approve the changes
// that were made. A customer must acknowledge the change and approve it so that the upgrade request
// can proceed.
func (softlayer_product_upgrade_request *SoftLayer_Product_Upgrade_Request) ApproveChanges() (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - getObject retrieves a SoftLayer_Product_Upgrade_Request object on your account whose ID
// corresponds to the ID of the init parameter passed to the SoftLayer_Product_Upgrade_Request service.
func (softlayer_product_upgrade_request *SoftLayer_Product_Upgrade_Request) GetObject() (*SoftLayer_Product_Upgrade_Request, error) {
	var returnValue *SoftLayer_Product_Upgrade_Request
	return returnValue, nil
}

// UpdateMaintenanceWindow - In case an upgrade cannot be performed, the maintenance window needs to be
// updated to a future date.
func (softlayer_product_upgrade_request *SoftLayer_Product_Upgrade_Request) UpdateMaintenanceWindow(maintenanceStartTime time.Time, maintenanceWindowId int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}
