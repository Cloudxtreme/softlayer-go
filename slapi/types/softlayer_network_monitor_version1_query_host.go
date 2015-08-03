package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

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

func (softlayer_network_monitor_version1_query_host *SoftLayer_Network_Monitor_Version1_Query_Host) String() string {
	return "SoftLayer_Network_Monitor_Version1_Query_Host"
}
