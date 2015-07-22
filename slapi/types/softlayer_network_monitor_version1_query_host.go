package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Network_Monitor_Version1_Query_Host - The Monitoring_Query_Host type represents a
// monitoring instance. It consists of a hardware ID to monitor, an IP address attached to that
// hardware ID, a method of monitoring, and what to do in the instance that the monitor ever fails.
type SoftLayer_Network_Monitor_Version1_Query_Host struct {

	// Arg1Value - The argument to be used for this monitor, if necessary. The lowest monitoring levels
	// (like ping) ignore this setting, but higher levels like custom use it.
	Arg1Value string `json:"arg1Value"`

	// GuestId - Virtual Guest Identification Number for the guest being monitored.
	GuestId int `json:"guestId"`

	// Hardware - The hardware that is being monitored by this monitoring instance
	Hardware *SoftLayer_Hardware `json:"hardware"`

	// HardwareId - no documentation
	HardwareId int `json:"hardwareId"`

	// HostId - Identification Number for the host being monitored.
	HostId int `json:"hostId"`

	// Id - no documentation
	Id int `json:"id"`

	// IpAddress - The IP address to be monitored. Must be attached to the hardware on this object
	IpAddress string `json:"ipAddress"`

	// LastResult - The most recent result for this particular monitoring instance.
	LastResult *SoftLayer_Network_Monitor_Version1_Query_Result `json:"lastResult"`

	// QueryType - The type of monitoring query that is executed when this hardware is monitored.
	QueryType *SoftLayer_Network_Monitor_Version1_Query_Type `json:"queryType"`

	// QueryTypeId - no documentation
	QueryTypeId int `json:"queryTypeId"`

	// ResponseAction - no documentation
	ResponseAction *SoftLayer_Network_Monitor_Version1_Query_ResponseType `json:"responseAction"`

	// ResponseActionId - The ID of the response action to take when the monitor fails
	ResponseActionId int `json:"responseActionId"`

	// Status - The status of this monitoring instance. Anything other than means that the monitor has been
	// disabled
	Status string `json:"status"`

	// Subnet - The subnet object that holds the IP address associated with this query host.
	Subnet *SoftLayer_Network_Subnet `json:"subnet"`

	// WaitCycles - The number of 5-minute cycles to wait before the "responseAction" is taken. If set to
	// 0, the response action will be taken immediately
	WaitCycles int `json:"waitCycles"`
}

// CreateObject - Passing in an unsaved instances of a Query_Host object into this function will create
// the object and return the results to the user.
func (softlayer_network_monitor_version1_query_host *SoftLayer_Network_Monitor_Version1_Query_Host) CreateObject(commonOptions *slapi.CommonOptions, templateObject SoftLayer_Network_Monitor_Version1_Query_Host) (*SoftLayer_Network_Monitor_Version1_Query_Host, error) {
	var returnValue *SoftLayer_Network_Monitor_Version1_Query_Host
	return returnValue, nil
}

// CreateObjects - Passing in a collection of unsaved instances of Query_Host objects into this
// function will create all objects and return the results to the user.
func (softlayer_network_monitor_version1_query_host *SoftLayer_Network_Monitor_Version1_Query_Host) CreateObjects(commonOptions *slapi.CommonOptions, templateObjects []SoftLayer_Network_Monitor_Version1_Query_Host) ([]*SoftLayer_Network_Monitor_Version1_Query_Host, error) {
	var returnValue []*SoftLayer_Network_Monitor_Version1_Query_Host
	return returnValue, nil
}

// DeleteObject - Like any other API object, the monitoring objects can be deleted by passing an
// instance of them into this function. The ID on the object must be set.
func (softlayer_network_monitor_version1_query_host *SoftLayer_Network_Monitor_Version1_Query_Host) DeleteObject(commonOptions *slapi.CommonOptions) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// DeleteObjects - Like any other API object, the monitoring objects can be deleted by passing an
// instance of them into this function. The ID on the object must be set.
func (softlayer_network_monitor_version1_query_host *SoftLayer_Network_Monitor_Version1_Query_Host) DeleteObjects(commonOptions *slapi.CommonOptions, templateObjects []SoftLayer_Network_Monitor_Version1_Query_Host) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObject - Like any other API object, the monitoring objects can have their exposed properties
// edited by passing in a modified version of the object.
func (softlayer_network_monitor_version1_query_host *SoftLayer_Network_Monitor_Version1_Query_Host) EditObject(commonOptions *slapi.CommonOptions, templateObject SoftLayer_Network_Monitor_Version1_Query_Host) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObjects - Like any other API object, the monitoring objects can have their exposed properties
// edited by passing in a modified version of the object.
func (softlayer_network_monitor_version1_query_host *SoftLayer_Network_Monitor_Version1_Query_Host) EditObjects(commonOptions *slapi.CommonOptions, templateObjects []SoftLayer_Network_Monitor_Version1_Query_Host) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// FindByHardwareId - This method returns all Query_Host objects associated with the passed in hardware
// ID as long as that hardware ID is owned by the current user's account. This behavior can also be
// accomplished by simply tapping networkMonitors on the Hardware_Server object.
func (softlayer_network_monitor_version1_query_host *SoftLayer_Network_Monitor_Version1_Query_Host) FindByHardwareId(commonOptions *slapi.CommonOptions, hardwareId int) ([]*SoftLayer_Network_Monitor_Version1_Query_Host, error) {
	var returnValue []*SoftLayer_Network_Monitor_Version1_Query_Host
	return returnValue, nil
}

// GetObject - getObject retrieves the SoftLayer_Network_Monitor_Version1_Query_Host object whose ID
// number corresponds to the ID number of the init parameter passed to the
// SoftLayer_Network_Monitor_Version1_Query_Host service. You can only retrieve query hosts attached to
// hardware that belong to your account.
func (softlayer_network_monitor_version1_query_host *SoftLayer_Network_Monitor_Version1_Query_Host) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Network_Monitor_Version1_Query_Host, error) {
	var returnValue *SoftLayer_Network_Monitor_Version1_Query_Host
	return returnValue, nil
}
