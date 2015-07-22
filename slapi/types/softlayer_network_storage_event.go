package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Storage_Event - Storage volumes can create various events to keep track of what
// has occurred to the volume. Events provide an audit trail that can be used to verify that various
// tasks have occurred, such as snapshots to be created by a schedule or remote replication
// synchronization.
type SoftLayer_Network_Storage_Event struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Message - no documentation
	Message string `json:"message"`

	// Schedule - A schedule that is associated with an event. Not all events will have a schedule.
	Schedule *SoftLayer_Network_Storage_Schedule `json:"schedule"`

	// ScheduleId - An identifier for the schedule which is associated with an event.
	ScheduleId int `json:"scheduleId"`

	// TypeId - no documentation
	TypeId int `json:"typeId"`

	// Volume - no documentation
	Volume *SoftLayer_Network_Storage `json:"volume"`

	// VolumeId - no documentation
	VolumeId int `json:"volumeId"`
}