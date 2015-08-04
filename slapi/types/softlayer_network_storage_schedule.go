package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Storage_Schedule - Schedules can be created for select Storage services, such as
// iscsi. These schedules are used to perform various tasks such as scheduling snapshots or
// synchronizing replicants.
type SoftLayer_Network_Storage_Schedule struct {

	// VolumeId - no documentation
	VolumeId int `json:"volumeId,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// Active - no documentation
	Active int `json:"active,omitempty"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// PartnershipId - The partnership id which a schedule is associated with.
	PartnershipId int `json:"partnershipId,omitempty"`

	// TypeId - no documentation
	TypeId int `json:"typeId,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`
}

func (softlayer_network_storage_schedule *SoftLayer_Network_Storage_Schedule) String() string {
	return "SoftLayer_Network_Storage_Schedule"
}

// SoftLayer_Network_Storage_Schedule_Extended is SoftLayer_Network_Storage_Schedule with all maskable types.
type SoftLayer_Network_Storage_Schedule_Extended struct {
	SoftLayer_Network_Storage_Schedule

	// EventCount - A count of events which have been created as the result of a schedule execution.
	EventCount uint64 `json:"eventCount,omitempty"`

	// PropertyCount - A count of properties used for configuration of a schedule.
	PropertyCount uint64 `json:"propertyCount,omitempty"`

	// DayOfWeek - no documentation
	DayOfWeek string `json:"dayOfWeek,omitempty"`

	// RetentionCount - The number of snapshots this schedule is configured to retain.
	RetentionCount string `json:"retentionCount,omitempty"`

	// SnapshotCount - A count of snapshots which have been created as the result of this schedule's
	// execution.
	SnapshotCount uint64 `json:"snapshotCount,omitempty"`

	// Events - Events which have been created as the result of a schedule execution.
	Events []*SoftLayer_Network_Storage_Event `json:"events,omitempty"`

	// Properties - no documentation
	Properties []*SoftLayer_Network_Storage_Schedule_Property `json:"properties,omitempty"`

	// Type - The type provides a standardized definition for a schedule.
	Type *SoftLayer_Network_Storage_Schedule_Type `json:"type,omitempty"`

	// DayOfMonth - no documentation
	DayOfMonth string `json:"dayOfMonth,omitempty"`

	// Minute - no documentation
	Minute string `json:"minute,omitempty"`

	// Volume - no documentation
	Volume *SoftLayer_Network_Storage `json:"volume,omitempty"`

	// Hour - no documentation
	Hour string `json:"hour,omitempty"`

	// MonthOfYear - no documentation
	MonthOfYear string `json:"monthOfYear,omitempty"`

	// Partnership - no documentation
	Partnership *SoftLayer_Network_Storage_Partnership `json:"partnership,omitempty"`

	// Snapshots - Snapshots which have been created as the result of this schedule's execution.
	Snapshots []*SoftLayer_Network_Storage `json:"snapshots,omitempty"`
}

func (softlayer_network_storage_schedule *SoftLayer_Network_Storage_Schedule_Extended) String() string {
	return "SoftLayer_Network_Storage_Schedule"
}
