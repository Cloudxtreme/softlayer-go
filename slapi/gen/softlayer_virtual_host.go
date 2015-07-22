package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Virtual_Host - The virtual host represents the platform on which virtual guests reside. At
// times a virtual host has no allocations on the physical server, however with many modern platforms
// it is a virtual machine with small CPU and Memory allocations that runs in the Control Domain.
type SoftLayer_Virtual_Host struct {

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// AccountId - no documentation
	AccountId int `json:"accountId"`

	// BilledPerGuestFlag - Boolean flag indicating whether this virtualization platform gets billed per
	// guest rather than at a fixed rate.
	BilledPerGuestFlag bool `json:"billedPerGuestFlag"`

	// BilledPerMemoryUsageFlag - Boolean flag indicating whether this virtualization platform gets billed
	// per memory usage rather than at a fixed rate.
	BilledPerMemoryUsageFlag bool `json:"billedPerMemoryUsageFlag"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Description - no documentation
	Description string `json:"description"`

	// EnabledFlag - The enabled flag specifies whether a virtual host can run guests.
	EnabledFlag int `json:"enabledFlag"`

	// GuestCount - A count of the guests associated with a virtual host.
	GuestCount uint64 `json:"guestCount"`

	// Guests - no documentation
	Guests []*SoftLayer_Virtual_Guest `json:"guests"`

	// Hardware - The hardware record which a virtual host resides on.
	Hardware *SoftLayer_Hardware_Server `json:"hardware"`

	// HardwareId - no documentation
	HardwareId int `json:"hardwareId"`

	// Id - no documentation
	Id int `json:"id"`

	// MetricTrackingObject - no documentation
	MetricTrackingObject *SoftLayer_Metric_Tracking_Object `json:"metricTrackingObject"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// Name - no documentation
	Name string `json:"name"`

	// PhysicalMemoryCapacity - The amount of memory physically available for a virtual host.
	PhysicalMemoryCapacity int `json:"physicalMemoryCapacity"`

	// Uuid - Unique ID for a virtual host's record on a virtualization platform.
	Uuid string `json:"uuid"`
}

// GetLiveGuestByUuid - Query a virtualization platform directly to retrieve details regarding a guest.
func (softlayer_virtual_host *SoftLayer_Virtual_Host) GetLiveGuestByUuid(uuid string) (*SoftLayer_Virtual_Guest, error) {
	var returnValue *SoftLayer_Virtual_Guest
	return returnValue, nil
}

// GetLiveGuestList - Query a virtualization platform directly to retrieve a list of known guests.
func (softlayer_virtual_host *SoftLayer_Virtual_Host) GetLiveGuestList() ([]*SoftLayer_Virtual_Guest, error) {
	var returnValue []*SoftLayer_Virtual_Guest
	return returnValue, nil
}

// GetLiveGuestRecentMetricData - Query a virtualization platform directly to retrieve recent metric
// data for a guest.
func (softlayer_virtual_host *SoftLayer_Virtual_Host) GetLiveGuestRecentMetricData(uuid string, time int, limit int, interval int) ([]*SoftLayer_Metric_Tracking_Object, error) {
	var returnValue []*SoftLayer_Metric_Tracking_Object
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_virtual_host *SoftLayer_Virtual_Host) GetObject() (*SoftLayer_Virtual_Host, error) {
	var returnValue *SoftLayer_Virtual_Host
	return returnValue, nil
}

// PauseLiveGuest - no documentation
func (softlayer_virtual_host *SoftLayer_Virtual_Host) PauseLiveGuest(uuid string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// PowerCycleLiveGuest - no documentation
func (softlayer_virtual_host *SoftLayer_Virtual_Host) PowerCycleLiveGuest(uuid string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// PowerOffLiveGuest - no documentation
func (softlayer_virtual_host *SoftLayer_Virtual_Host) PowerOffLiveGuest(uuid string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// PowerOnLiveGuest - no documentation
func (softlayer_virtual_host *SoftLayer_Virtual_Host) PowerOnLiveGuest(uuid string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RebootSoftLiveGuest - Attempt to complete a soft reboot of a guest by shutting down the operating
// system.
func (softlayer_virtual_host *SoftLayer_Virtual_Host) RebootSoftLiveGuest(uuid string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// ResumeLiveGuest - no documentation
func (softlayer_virtual_host *SoftLayer_Virtual_Host) ResumeLiveGuest(uuid string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}
