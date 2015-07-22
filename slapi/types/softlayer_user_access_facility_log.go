package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_User_Access_Facility_Log - This class represents a login/logout sheet for facility
// visitors.
type SoftLayer_User_Access_Facility_Log struct {

	// Account - This is the account associated with the log entry. For users under a customer's account,
	// it is the customer's account. For contractors and others visiting a colocation area, it is the
	// account associated with the area they visited.
	Account *SoftLayer_Account `json:"account"`

	// AccountId - This is the account associated with a log record. For a customer logging into a
	// datacenter, this is the customer's account. For a contractor or any other guest logging into a
	// customer's cabinet or colocation cage, this is the customer's account.
	AccountId int `json:"accountId"`

	// Datacenter - no documentation
	Datacenter *SoftLayer_Location `json:"datacenter"`

	// Description - This is a short description of why the person is at the location.
	Description string `json:"description"`

	// Hardware - no documentation
	Hardware *SoftLayer_Hardware `json:"hardware"`

	// HardwareId - <nil>
	HardwareId int `json:"hardwareId"`

	// Id - <nil>
	Id int `json:"id"`

	// LocationId - <nil>
	LocationId int `json:"locationId"`

	// LogType - no documentation
	LogType *SoftLayer_User_Access_Facility_Log_Type `json:"logType"`

	// TimeIn - no documentation
	TimeIn *time.Time `json:"timeIn"`

	// TimeOut - <nil>
	TimeOut *time.Time `json:"timeOut"`

	// Visitor - <nil>
	Visitor *SoftLayer_Entity `json:"visitor"`
}
