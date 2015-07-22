package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
)

// SoftLayer_Network_Storage_Schedule - Schedules can be created for select Storage services, such as
// iscsi. These schedules are used to perform various tasks such as scheduling snapshots or
// synchronizing replicants.
type SoftLayer_Network_Storage_Schedule struct {

	// Active - no documentation
	Active int `json:"active"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// DayOfMonth - no documentation
	DayOfMonth string `json:"dayOfMonth"`

	// DayOfWeek - no documentation
	DayOfWeek string `json:"dayOfWeek"`

	// EventCount - A count of events which have been created as the result of a schedule execution.
	EventCount uint64 `json:"eventCount"`

	// Events - Events which have been created as the result of a schedule execution.
	Events []*SoftLayer_Network_Storage_Event `json:"events"`

	// Hour - no documentation
	Hour string `json:"hour"`

	// Id - no documentation
	Id int `json:"id"`

	// Minute - no documentation
	Minute string `json:"minute"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// MonthOfYear - no documentation
	MonthOfYear string `json:"monthOfYear"`

	// Name - no documentation
	Name string `json:"name"`

	// Partnership - no documentation
	Partnership *SoftLayer_Network_Storage_Partnership `json:"partnership"`

	// PartnershipId - The partnership id which a schedule is associated with.
	PartnershipId int `json:"partnershipId"`

	// Properties - no documentation
	Properties []*SoftLayer_Network_Storage_Schedule_Property `json:"properties"`

	// PropertyCount - A count of properties used for configuration of a schedule.
	PropertyCount uint64 `json:"propertyCount"`

	// RetentionCount - The number of snapshots this schedule is configured to retain.
	RetentionCount string `json:"retentionCount"`

	// SnapshotCount - A count of snapshots which have been created as the result of this schedule's
	// execution.
	SnapshotCount uint64 `json:"snapshotCount"`

	// Snapshots - Snapshots which have been created as the result of this schedule's execution.
	Snapshots []*SoftLayer_Network_Storage `json:"snapshots"`

	// Type - The type provides a standardized definition for a schedule.
	Type *SoftLayer_Network_Storage_Schedule_Type `json:"type"`

	// TypeId - no documentation
	TypeId int `json:"typeId"`

	// Volume - no documentation
	Volume *SoftLayer_Network_Storage `json:"volume"`

	// VolumeId - no documentation
	VolumeId int `json:"volumeId"`
}

// CreateObject - no documentation
func (softlayer_network_storage_schedule *SoftLayer_Network_Storage_Schedule) CreateObject(commonOptions *slapi.CommonOptions, templateObject SoftLayer_Network_Storage_Schedule) (*SoftLayer_Network_Storage_Schedule, error) {
	var returnValue *SoftLayer_Network_Storage_Schedule
	return returnValue, nil
}

// DeleteObject - Delete a network storage schedule. '''This cannot be undone.''' ''deleteObject''
// returns Boolean ''true'' on successful deletion or ''false'' if it was unable to remove a schedule;
func (softlayer_network_storage_schedule *SoftLayer_Network_Storage_Schedule) DeleteObject(commonOptions *slapi.CommonOptions) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObject - no documentation
func (softlayer_network_storage_schedule *SoftLayer_Network_Storage_Schedule) EditObject(commonOptions *slapi.CommonOptions, templateObject SoftLayer_Network_Storage_Schedule) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_network_storage_schedule *SoftLayer_Network_Storage_Schedule) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Network_Storage_Schedule, error) {
	var returnValue *SoftLayer_Network_Storage_Schedule
	return returnValue, nil
}
