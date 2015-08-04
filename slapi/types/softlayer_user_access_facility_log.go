package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_User_Access_Facility_Log - This class represents a login/logout sheet for facility
// visitors.
type SoftLayer_User_Access_Facility_Log struct {

	// Description - This is a short description of why the person is at the location.
	Description string `json:"description,omitempty"`

	// TimeOut - <nil>
	TimeOut *time.Time `json:"timeOut,omitempty"`

	// TimeIn - no documentation
	TimeIn *time.Time `json:"timeIn,omitempty"`

	// AccountId - This is the account associated with a log record. For a customer logging into a
	// datacenter, this is the customer's account. For a contractor or any other guest logging into a
	// customer's cabinet or colocation cage, this is the customer's account.
	AccountId int `json:"accountId,omitempty"`

	// HardwareId - <nil>
	HardwareId int `json:"hardwareId,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// LocationId - <nil>
	LocationId int `json:"locationId,omitempty"`

	// Account - This is the account associated with the log entry. For users under a customer's account,
	// it is the customer's account. For contractors and others visiting a colocation area, it is the
	// account associated with the area they visited.
	Account *SoftLayer_Account `json:"account,omitempty"`

	// Hardware - no documentation
	Hardware *SoftLayer_Hardware `json:"hardware,omitempty"`

	// Datacenter - no documentation
	Datacenter *SoftLayer_Location `json:"datacenter,omitempty"`

	// LogType - no documentation
	LogType *SoftLayer_User_Access_Facility_Log_Type `json:"logType,omitempty"`

	// Visitor - <nil>
	Visitor *SoftLayer_Entity `json:"visitor,omitempty"`
}

func (softlayer_user_access_facility_log *SoftLayer_User_Access_Facility_Log) String() string {
	return "SoftLayer_User_Access_Facility_Log"
}
