package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Storage_Schedule - Schedules can be created for select Storage services, such as
// iscsi. These schedules are used to perform various tasks such as scheduling snapshots or
// synchronizing replicants.
type SoftLayer_Network_Storage_Schedule struct {

	// TypeId - no documentation
	TypeId int `json:"typeId"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Name - no documentation
	Name string `json:"name"`

	// Id - no documentation
	Id int `json:"id"`

	// PartnershipId - The partnership id which a schedule is associated with.
	PartnershipId int `json:"partnershipId"`

	// VolumeId - no documentation
	VolumeId int `json:"volumeId"`

	// Active - no documentation
	Active int `json:"active"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`
}

func (softlayer_network_storage_schedule *SoftLayer_Network_Storage_Schedule) String() string {
	return "SoftLayer_Network_Storage_Schedule"
}

// SoftLayer_Network_Storage_Schedule_Extended is SoftLayer_Network_Storage_Schedule with all maskable types.
type SoftLayer_Network_Storage_Schedule_Extended struct {
	SoftLayer_Network_Storage_Schedule

	// DayOfMonth - no documentation
	DayOfMonth string `json:"dayOfMonth"`

	// Hour - no documentation
	Hour string `json:"hour"`

	// Partnership - no documentation
	Partnership *SoftLayer_Network_Storage_Partnership `json:"partnership"`

	// Properties - no documentation
	Properties []*SoftLayer_Network_Storage_Schedule_Property `json:"properties"`

	// RetentionCount - The number of snapshots this schedule is configured to retain.
	RetentionCount string `json:"retentionCount"`

	// Type - The type provides a standardized definition for a schedule.
	Type *SoftLayer_Network_Storage_Schedule_Type `json:"type"`

	// SnapshotCount - A count of snapshots which have been created as the result of this schedule's
	// execution.
	SnapshotCount uint64 `json:"snapshotCount"`

	// EventCount - A count of events which have been created as the result of a schedule execution.
	EventCount uint64 `json:"eventCount"`

	// PropertyCount - A count of properties used for configuration of a schedule.
	PropertyCount uint64 `json:"propertyCount"`

	// DayOfWeek - no documentation
	DayOfWeek string `json:"dayOfWeek"`

	// Events - Events which have been created as the result of a schedule execution.
	Events []*SoftLayer_Network_Storage_Event `json:"events"`

	// Minute - no documentation
	Minute string `json:"minute"`

	// MonthOfYear - no documentation
	MonthOfYear string `json:"monthOfYear"`

	// Snapshots - Snapshots which have been created as the result of this schedule's execution.
	Snapshots []*SoftLayer_Network_Storage `json:"snapshots"`

	// Volume - no documentation
	Volume *SoftLayer_Network_Storage `json:"volume"`
}

func (softlayer_network_storage_schedule *SoftLayer_Network_Storage_Schedule_Extended) String() string {
	return "SoftLayer_Network_Storage_Schedule"
}
