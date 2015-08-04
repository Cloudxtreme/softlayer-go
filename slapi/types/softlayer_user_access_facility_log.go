package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_User_Access_Facility_Log - This class represents a login/logout sheet for facility
// visitors.
type SoftLayer_User_Access_Facility_Log struct {

	// TimeOut - <nil>
	TimeOut *time.Time `json:"timeOut,omitempty"`

	// AccountId - This is the account associated with a log record. For a customer logging into a
	// datacenter, this is the customer's account. For a contractor or any other guest logging into a
	// customer's cabinet or colocation cage, this is the customer's account.
	AccountId int `json:"accountId,omitempty"`

	// Description - This is a short description of why the person is at the location.
	Description string `json:"description,omitempty"`

	// HardwareId - <nil>
	HardwareId int `json:"hardwareId,omitempty"`

	// LocationId - <nil>
	LocationId int `json:"locationId,omitempty"`

	// TimeIn - no documentation
	TimeIn *time.Time `json:"timeIn,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`
}

func (softlayer_user_access_facility_log *SoftLayer_User_Access_Facility_Log) String() string {
	return "SoftLayer_User_Access_Facility_Log"
}

// SoftLayer_User_Access_Facility_Log_Extended is SoftLayer_User_Access_Facility_Log with all maskable types.
type SoftLayer_User_Access_Facility_Log_Extended struct {
	SoftLayer_User_Access_Facility_Log

	// Account - This is the account associated with the log entry. For users under a customer's account,
	// it is the customer's account. For contractors and others visiting a colocation area, it is the
	// account associated with the area they visited.
	Account *SoftLayer_Account `json:"account,omitempty"`

	// Hardware - no documentation
	Hardware *SoftLayer_Hardware `json:"hardware,omitempty"`

	// LogType - no documentation
	LogType *SoftLayer_User_Access_Facility_Log_Type `json:"logType,omitempty"`

	// Visitor - <nil>
	Visitor *SoftLayer_Entity `json:"visitor,omitempty"`

	// Datacenter - no documentation
	Datacenter *SoftLayer_Location `json:"datacenter,omitempty"`
}

func (softlayer_user_access_facility_log *SoftLayer_User_Access_Facility_Log_Extended) String() string {
	return "SoftLayer_User_Access_Facility_Log"
}
