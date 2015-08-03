package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Hardware_Server - The SoftLayer_Hardware_Server data type contains general information
// relating to a single SoftLayer server.
type SoftLayer_Hardware_Server struct {
}

// SoftLayer_Hardware_Server_Extended is SoftLayer_Hardware_Server with all maskable types.
type SoftLayer_Hardware_Server_Extended struct {
	SoftLayer_Hardware_Server

	// MetricTrackingObjectId - no documentation
	MetricTrackingObjectId int `json:"metricTrackingObjectId"`

	// PrivateIpAddress - no documentation
	PrivateIpAddress string `json:"privateIpAddress"`

	// ProjectedPublicBandwidthUsage - The projected public outbound bandwidth for this hardware for the
	// current billing cycle.
	ProjectedPublicBandwidthUsage float32 `json:"projectedPublicBandwidthUsage"`

	// RegionalInternetRegistry - <nil>
	RegionalInternetRegistry *SoftLayer_Network_Regional_Internet_Registry `json:"regionalInternetRegistry"`

	// RemoteManagement - no documentation
	RemoteManagement *SoftLayer_Hardware_Component_RemoteManagement `json:"remoteManagement"`

	// ActiveTransactionCount - A count of any active transaction(s) that are currently running for the
	// server (example: os reload).
	ActiveTransactionCount uint64 `json:"activeTransactionCount"`

	// ActiveTransactions - Any active transaction(s) that are currently running for the server (example:
	// os reload).
	ActiveTransactions []*SoftLayer_Provisioning_Version1_Transaction `json:"activeTransactions"`

	// BillingCyclePublicBandwidthUsage - The raw public bandwidth usage data for the current billing
	// cycle.
	BillingCyclePublicBandwidthUsage *SoftLayer_Network_Bandwidth_Usage `json:"billingCyclePublicBandwidthUsage"`

	// AvailableMonitoringCount - A count of an object that stores the maximum level for the monitoring
	// query types and response types.
	AvailableMonitoringCount uint64 `json:"availableMonitoringCount"`

	// OpenCancellationTicket - An open ticket requesting cancellation of this server, if one exists.
	OpenCancellationTicket *SoftLayer_Ticket `json:"openCancellationTicket"`

	// RecentRemoteManagementCommands - The last five commands issued to the server's remote management
	// card.
	RecentRemoteManagementCommands []*SoftLayer_Hardware_Component_RemoteManagement_Command_Request `json:"recentRemoteManagementCommands"`

	// Users - A list of users that have access to this computing instance.
	Users []*SoftLayer_User_Customer `json:"users"`

	// ActiveTicketCount - no documentation
	ActiveTicketCount uint64 `json:"activeTicketCount"`

	// RecentRemoteManagementCommandCount - A count of the last five commands issued to the server's remote
	// management card.
	RecentRemoteManagementCommandCount uint64 `json:"recentRemoteManagementCommandCount"`

	// ActiveNetworkFirewallBillingItem - The billing item for a server's attached network firewall.
	ActiveNetworkFirewallBillingItem *SoftLayer_Billing_Item `json:"activeNetworkFirewallBillingItem"`

	// ActiveTickets - <nil>
	ActiveTickets []*SoftLayer_Ticket `json:"activeTickets"`

	// BillingCycleBandwidthUsageCount - A count of the raw bandwidth usage data for the current billing
	// cycle. One object will be returned for each network this server is attached to.
	BillingCycleBandwidthUsageCount uint64 `json:"billingCycleBandwidthUsageCount"`

	// LastOperatingSystemReload - The last transaction that a server's operating system was loaded.
	LastOperatingSystemReload *SoftLayer_Provisioning_Version1_Transaction `json:"lastOperatingSystemReload"`

	// ProjectedOverBandwidthAllocationFlag - Whether the bandwidth usage for this hardware for the current
	// billing cycle is projected to exceed the allocation.
	ProjectedOverBandwidthAllocationFlag int `json:"projectedOverBandwidthAllocationFlag"`

	// CustomerOwnedFlag - no documentation
	CustomerOwnedFlag bool `json:"customerOwnedFlag"`

	// OverBandwidthAllocationFlag - Whether the bandwidth usage for this hardware for the current billing
	// cycle exceeds the allocation.
	OverBandwidthAllocationFlag int `json:"overBandwidthAllocationFlag"`

	// AverageDailyBandwidthUsage - The average daily total bandwidth usage for the current billing cycle.
	AverageDailyBandwidthUsage float32 `json:"averageDailyBandwidthUsage"`

	// ControlPanel - no documentation
	ControlPanel *SoftLayer_Software_Component_ControlPanel `json:"controlPanel"`

	// InboundPrivateBandwidthUsage - The total private inbound bandwidth for this hardware for the current
	// billing cycle.
	InboundPrivateBandwidthUsage float64 `json:"inboundPrivateBandwidthUsage"`

	// AvailableMonitoring - An object that stores the maximum level for the monitoring query types and
	// response types.
	AvailableMonitoring []*SoftLayer_Network_Monitor_Version1_Query_Host_Stratum `json:"availableMonitoring"`

	// CustomerInstalledOperatingSystemFlag - no documentation
	CustomerInstalledOperatingSystemFlag bool `json:"customerInstalledOperatingSystemFlag"`

	// CurrentBandwidthSummary - An object that provides commonly used bandwidth summary components for the
	// current billing cycle.
	CurrentBandwidthSummary *SoftLayer_Metric_Tracking_Object_Bandwidth_Summary `json:"currentBandwidthSummary"`

	// UserCount - A count of a list of users that have access to this computing instance.
	UserCount uint64 `json:"userCount"`

	// BillingCyclePrivateBandwidthUsage - The raw private bandwidth usage data for the current billing
	// cycle.
	BillingCyclePrivateBandwidthUsage *SoftLayer_Network_Bandwidth_Usage `json:"billingCyclePrivateBandwidthUsage"`

	// ContainsSolidStateDrivesFlag - <nil>
	ContainsSolidStateDrivesFlag bool `json:"containsSolidStateDrivesFlag"`

	// OutboundPrivateBandwidthUsage - The total private outbound bandwidth for this hardware for the
	// current billing cycle.
	OutboundPrivateBandwidthUsage float64 `json:"outboundPrivateBandwidthUsage"`

	// VirtualGuests - no documentation
	VirtualGuests []*SoftLayer_Virtual_Guest `json:"virtualGuests"`

	// MonitoringUserNotificationCount - A count of the monitoring notification objects for this hardware.
	// Each object links this hardware instance to a user account that will be notified if monitoring on
	// this hardware object fails
	MonitoringUserNotificationCount uint64 `json:"monitoringUserNotificationCount"`

	// VirtualGuestCount - no documentation
	VirtualGuestCount uint64 `json:"virtualGuestCount"`

	// AverageDailyPrivateBandwidthUsage - The average daily private bandwidth usage for the current
	// billing cycle.
	AverageDailyPrivateBandwidthUsage float32 `json:"averageDailyPrivateBandwidthUsage"`

	// BillingCycleBandwidthUsage - The raw bandwidth usage data for the current billing cycle. One object
	// will be returned for each network this server is attached to.
	BillingCycleBandwidthUsage []*SoftLayer_Network_Bandwidth_Usage `json:"billingCycleBandwidthUsage"`

	// MonitoringUserNotification - The monitoring notification objects for this hardware. Each object
	// links this hardware instance to a user account that will be notified if monitoring on this hardware
	// object fails
	MonitoringUserNotification []*SoftLayer_User_Customer_Notification_Hardware `json:"monitoringUserNotification"`

	// RemoteManagementUsers - User(s) who have access to issue commands and/or interact with the server's
	// remote management card.
	RemoteManagementUsers []*SoftLayer_Hardware_Component_RemoteManagement_User `json:"remoteManagementUsers"`

	// StatisticsRemoteManagement - A server's remote management card used for statistics.
	StatisticsRemoteManagement *SoftLayer_Hardware_Component_RemoteManagement `json:"statisticsRemoteManagement"`

	// RemoteManagementUserCount - A count of user(s) who have access to issue commands and/or interact
	// with the server's remote management card.
	RemoteManagementUserCount uint64 `json:"remoteManagementUserCount"`

	// ActiveTransaction - no documentation
	ActiveTransaction *SoftLayer_Provisioning_Version1_Transaction `json:"activeTransaction"`

	// Cost - The total cost of a server, measured in US Dollars
	Cost float32 `json:"cost"`
}

func (softlayer_hardware_server *SoftLayer_Hardware_Server) String() string {
	return "SoftLayer_Hardware_Server"
}
