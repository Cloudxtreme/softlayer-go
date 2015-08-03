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

	// HardwareId - The identifier of the hardware item a scan is run on.
	HardwareId int `json:"hardwareId"`

	// IpAddress - no documentation
	IpAddress string `json:"ipAddress"`

	// StatusId - no documentation
	StatusId int `json:"statusId"`

	// AccountId - A request's associated customer account identifier.
	AccountId int `json:"accountId"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Id - no documentation
	Id int `json:"id"`

	// GuestId - Virtual Guest Identification Number for the guest this security scanner request belongs
	// to.
	GuestId int `json:"guestId"`

	// HostId - Identification Number for the host this security scanner request belongs to.
	HostId int `json:"hostId"`

	// ModifyDate - The date and time that the request was last modified.
	ModifyDate *time.Time `json:"modifyDate"`
}

func (softlayer_network_security_scanner_request *SoftLayer_Network_Security_Scanner_Request) String() string {
	return "SoftLayer_Network_Security_Scanner_Request"
}

// SoftLayer_Network_Security_Scanner_Request_Extended is SoftLayer_Network_Security_Scanner_Request with all maskable types.
type SoftLayer_Network_Security_Scanner_Request_Extended struct {
	SoftLayer_Network_Security_Scanner_Request

	// RequestorOwnedFlag - Flag whether the requestor owns the hardware the scan was run on. This flag
	// will return for hardware servers only, virtual servers will result in a null return even if you have
	// a request out for them.
	RequestorOwnedFlag bool `json:"requestorOwnedFlag"`

	// Hardware - no documentation
	Hardware *SoftLayer_Hardware `json:"hardware"`

	// Status - no documentation
	Status *SoftLayer_Network_Security_Scanner_Request_Status `json:"status"`

	// Account - The account associated with a security scan request.
	Account *SoftLayer_Account `json:"account"`

	// Guest - no documentation
	Guest *SoftLayer_Virtual_Guest `json:"guest"`
}

func (softlayer_network_security_scanner_request *SoftLayer_Network_Security_Scanner_Request_Extended) String() string {
	return "SoftLayer_Network_Security_Scanner_Request"
}
