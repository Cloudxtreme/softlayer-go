package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Storage_Schedule_Property - Schedule properties provide attributes such as start
// date, end date, interval, and other properties to a storage schedule.
type SoftLayer_Network_Storage_Schedule_Property struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Id - no documentation
	Id int `json:"id"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// Schedule - no documentation
	Schedule *SoftLayer_Network_Storage_Schedule `json:"schedule"`

	// Type - The type provides a standardized definition for a property.
	Type *SoftLayer_Network_Storage_Schedule_Property_Type `json:"type"`

	// TypeId - no documentation
	TypeId int `json:"typeId"`

	// Value - no documentation
	Value string `json:"value"`
}

func (softlayer_network_storage_schedule_property *SoftLayer_Network_Storage_Schedule_Property) String() string {
	return "SoftLayer_Network_Storage_Schedule_Property"
}
