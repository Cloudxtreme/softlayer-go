package types

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

	// TypeId - no documentation
	TypeId int `json:"typeId"`

	// Value - no documentation
	Value string `json:"value"`
}

// SoftLayer_Network_Storage_Schedule_Property_Extended is SoftLayer_Network_Storage_Schedule_Property with all maskable types.
type SoftLayer_Network_Storage_Schedule_Property_Extended struct {
	SoftLayer_Network_Storage_Schedule_Property

	// Schedule - no documentation
	Schedule *SoftLayer_Network_Storage_Schedule `json:"schedule"`

	// Type - The type provides a standardized definition for a property.
	Type *SoftLayer_Network_Storage_Schedule_Property_Type `json:"type"`
}

func (softlayer_network_storage_schedule_property *SoftLayer_Network_Storage_Schedule_Property) String() string {
	return "SoftLayer_Network_Storage_Schedule_Property"
}
