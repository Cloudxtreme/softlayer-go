package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Storage_Schedule - Schedules can be created for select Storage services, such as
// iscsi. These schedules are used to perform various tasks such as scheduling snapshots or
// synchronizing replicants.
type SoftLayer_Network_Storage_Schedule struct {

	// PartnershipId - The partnership id which a schedule is associated with.
	PartnershipId int `json:"partnershipId"`

	// TypeId - no documentation
	TypeId int `json:"typeId"`

	// Name - no documentation
	Name string `json:"name"`

	// VolumeId - no documentation
	VolumeId int `json:"volumeId"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// Id - no documentation
	Id int `json:"id"`

	// Active - no documentation
	Active int `json:"active"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`
}

// SoftLayer_Network_Storage_Schedule_Extended is SoftLayer_Network_Storage_Schedule with all maskable types.
type SoftLayer_Network_Storage_Schedule_Extended struct {
	SoftLayer_Network_Storage_Schedule

	// Minute - no documentation
	Minute string `json:"minute"`

	// MonthOfYear - no documentation
	MonthOfYear string `json:"monthOfYear"`

	// RetentionCount - The number of snapshots this schedule is configured to retain.
	RetentionCount string `json:"retentionCount"`

	// Volume - no documentation
	Volume *SoftLayer_Network_Storage `json:"volume"`

	// EventCount - A count of events which have been created as the result of a schedule execution.
	EventCount uint64 `json:"eventCount"`

	// DayOfWeek - no documentation
	DayOfWeek string `json:"dayOfWeek"`

	// Snapshots - Snapshots which have been created as the result of this schedule's execution.
	Snapshots []*SoftLayer_Network_Storage `json:"snapshots"`

	// Type - The type provides a standardized definition for a schedule.
	Type *SoftLayer_Network_Storage_Schedule_Type `json:"type"`

	// SnapshotCount - A count of snapshots which have been created as the result of this schedule's
	// execution.
	SnapshotCount uint64 `json:"snapshotCount"`

	// Hour - no documentation
	Hour string `json:"hour"`

	// Partnership - no documentation
	Partnership *SoftLayer_Network_Storage_Partnership `json:"partnership"`

	// Properties - no documentation
	Properties []*SoftLayer_Network_Storage_Schedule_Property `json:"properties"`

	// DayOfMonth - no documentation
	DayOfMonth string `json:"dayOfMonth"`

	// Events - Events which have been created as the result of a schedule execution.
	Events []*SoftLayer_Network_Storage_Event `json:"events"`

	// PropertyCount - A count of properties used for configuration of a schedule.
	PropertyCount uint64 `json:"propertyCount"`
}

func (softlayer_network_storage_schedule *SoftLayer_Network_Storage_Schedule) String() string {
	return "SoftLayer_Network_Storage_Schedule"
}
