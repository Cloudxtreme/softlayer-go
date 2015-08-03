package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Security_Scanner_Request - The SoftLayer_Network_Security_Scanner_Request data
// type represents a single vulnerability scan request. It provides information on when the scan was
// created, last updated, and the current status. The status messages are as follows: *Scan Pending
// *Scan Processing *Scan Complete *Scan Cancelled *Generating Report.
type SoftLayer_Network_Security_Scanner_Request struct {

	// Account - The account associated with a security scan request.
	Account *SoftLayer_Account `json:"account"`

	// AccountId - A request's associated customer account identifier.
	AccountId int `json:"accountId"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Guest - no documentation
	Guest *SoftLayer_Virtual_Guest `json:"guest"`

	// GuestId - Virtual Guest Identification Number for the guest this security scanner request belongs
	// to.
	GuestId int `json:"guestId"`

	// Hardware - no documentation
	Hardware *SoftLayer_Hardware `json:"hardware"`

	// HardwareId - The identifier of the hardware item a scan is run on.
	HardwareId int `json:"hardwareId"`

	// HostId - Identification Number for the host this security scanner request belongs to.
	HostId int `json:"hostId"`

	// Id - no documentation
	Id int `json:"id"`

	// IpAddress - no documentation
	IpAddress string `json:"ipAddress"`

	// ModifyDate - The date and time that the request was last modified.
	ModifyDate *time.Time `json:"modifyDate"`

	// RequestorOwnedFlag - Flag whether the requestor owns the hardware the scan was run on. This flag
	// will return for hardware servers only, virtual servers will result in a null return even if you have
	// a request out for them.
	RequestorOwnedFlag bool `json:"requestorOwnedFlag"`

	// Status - no documentation
	Status *SoftLayer_Network_Security_Scanner_Request_Status `json:"status"`

	// StatusId - no documentation
	StatusId int `json:"statusId"`
}

func (softlayer_network_security_scanner_request *SoftLayer_Network_Security_Scanner_Request) String() string {
	return "SoftLayer_Network_Security_Scanner_Request"
}
