package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Virtual_Host - The virtual host represents the platform on which virtual guests reside. At
// times a virtual host has no allocations on the physical server, however with many modern platforms
// it is a virtual machine with small CPU and Memory allocations that runs in the Control Domain.
type SoftLayer_Virtual_Host struct {

	// HardwareId - no documentation
	HardwareId int `json:"hardwareId"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// PhysicalMemoryCapacity - The amount of memory physically available for a virtual host.
	PhysicalMemoryCapacity int `json:"physicalMemoryCapacity"`

	// AccountId - no documentation
	AccountId int `json:"accountId"`

	// Uuid - Unique ID for a virtual host's record on a virtualization platform.
	Uuid string `json:"uuid"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Description - no documentation
	Description string `json:"description"`

	// EnabledFlag - The enabled flag specifies whether a virtual host can run guests.
	EnabledFlag int `json:"enabledFlag"`

	// Id - no documentation
	Id int `json:"id"`

	// Name - no documentation
	Name string `json:"name"`
}

func (softlayer_virtual_host *SoftLayer_Virtual_Host) String() string {
	return "SoftLayer_Virtual_Host"
}

// SoftLayer_Virtual_Host_Extended is SoftLayer_Virtual_Host with all maskable types.
type SoftLayer_Virtual_Host_Extended struct {
	SoftLayer_Virtual_Host

	// BilledPerMemoryUsageFlag - Boolean flag indicating whether this virtualization platform gets billed
	// per memory usage rather than at a fixed rate.
	BilledPerMemoryUsageFlag bool `json:"billedPerMemoryUsageFlag"`

	// Guests - no documentation
	Guests []*SoftLayer_Virtual_Guest `json:"guests"`

	// Hardware - The hardware record which a virtual host resides on.
	Hardware *SoftLayer_Hardware_Server `json:"hardware"`

	// GuestCount - A count of the guests associated with a virtual host.
	GuestCount uint64 `json:"guestCount"`

	// BilledPerGuestFlag - Boolean flag indicating whether this virtualization platform gets billed per
	// guest rather than at a fixed rate.
	BilledPerGuestFlag bool `json:"billedPerGuestFlag"`

	// MetricTrackingObject - no documentation
	MetricTrackingObject *SoftLayer_Metric_Tracking_Object `json:"metricTrackingObject"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`
}

func (softlayer_virtual_host *SoftLayer_Virtual_Host_Extended) String() string {
	return "SoftLayer_Virtual_Host"
}
