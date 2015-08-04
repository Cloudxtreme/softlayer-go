package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Virtual_Host - The virtual host represents the platform on which virtual guests reside. At
// times a virtual host has no allocations on the physical server, however with many modern platforms
// it is a virtual machine with small CPU and Memory allocations that runs in the Control Domain.
type SoftLayer_Virtual_Host struct {

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// EnabledFlag - The enabled flag specifies whether a virtual host can run guests.
	EnabledFlag int `json:"enabledFlag,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// PhysicalMemoryCapacity - The amount of memory physically available for a virtual host.
	PhysicalMemoryCapacity int `json:"physicalMemoryCapacity,omitempty"`

	// Uuid - Unique ID for a virtual host's record on a virtualization platform.
	Uuid string `json:"uuid,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Description - no documentation
	Description string `json:"description,omitempty"`

	// HardwareId - no documentation
	HardwareId int `json:"hardwareId,omitempty"`

	// AccountId - no documentation
	AccountId int `json:"accountId,omitempty"`
}

func (softlayer_virtual_host *SoftLayer_Virtual_Host) String() string {
	return "SoftLayer_Virtual_Host"
}

// SoftLayer_Virtual_Host_Extended is SoftLayer_Virtual_Host with all maskable types.
type SoftLayer_Virtual_Host_Extended struct {
	SoftLayer_Virtual_Host

	// MetricTrackingObject - no documentation
	MetricTrackingObject *SoftLayer_Metric_Tracking_Object `json:"metricTrackingObject,omitempty"`

	// GuestCount - A count of the guests associated with a virtual host.
	GuestCount uint64 `json:"guestCount,omitempty"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account,omitempty"`

	// BilledPerGuestFlag - Boolean flag indicating whether this virtualization platform gets billed per
	// guest rather than at a fixed rate.
	BilledPerGuestFlag bool `json:"billedPerGuestFlag,omitempty"`

	// BilledPerMemoryUsageFlag - Boolean flag indicating whether this virtualization platform gets billed
	// per memory usage rather than at a fixed rate.
	BilledPerMemoryUsageFlag bool `json:"billedPerMemoryUsageFlag,omitempty"`

	// Guests - no documentation
	Guests []*SoftLayer_Virtual_Guest `json:"guests,omitempty"`

	// Hardware - The hardware record which a virtual host resides on.
	Hardware *SoftLayer_Hardware_Server `json:"hardware,omitempty"`
}

func (softlayer_virtual_host *SoftLayer_Virtual_Host_Extended) String() string {
	return "SoftLayer_Virtual_Host"
}
