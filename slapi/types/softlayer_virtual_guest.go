package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"

	time "time"
)

// SoftLayer_Virtual_Guest - The virtual guest data type presents the structure in which all virtual
// guests will be presented. Internally, the structure supports various virtualization platforms with
// no change to external interaction. A guest, also known as a virtual server, represents an allocation
// of resources on a virtual host.
type SoftLayer_Virtual_Guest struct {

	// StartCpus - The number of CPUs available to a computing instance upon startup.
	StartCpus int `json:"startCpus,omitempty"`

	// CreateDate - The date a virtual computing instance was created.
	CreateDate *time.Time `json:"createDate,omitempty"`

	// MaxCpu - The maximum amount of CPU resources a computing instance may utilize.
	MaxCpu int `json:"maxCpu,omitempty"`

	// Hostname - no documentation
	Hostname string `json:"hostname,omitempty"`

	// MaxMemory - The maximum amount of memory a computing instance may utilize.
	MaxMemory int `json:"maxMemory,omitempty"`

	// ProvisionDate - <nil>
	ProvisionDate *time.Time `json:"provisionDate,omitempty"`

	// PostInstallScriptUri - URI of the script to be downloaded and executed after installation is
	// complete. This is deprecated in favor of supplementalCreateObjectOptions' postInstallScriptUri.
	PostInstallScriptUri string `json:"postInstallScriptUri,omitempty"`

	// StatusId - A computing instances [[SoftLayer_Virtual_Guest_Status|status]] ID
	StatusId int `json:"statusId,omitempty"`

	// Uuid - Unique ID for a computing instance's record on a virtualization platform.
	Uuid string `json:"uuid,omitempty"`

	// Notes - A note of up to 1,000 characters about a virtual server.
	Notes string `json:"notes,omitempty"`

	// MaxCpuUnits - The unit of the maximum amount of CPU resources a computing instance may utilize.
	MaxCpuUnits string `json:"maxCpuUnits,omitempty"`

	// DedicatedAccountHostOnlyFlag - When true this flag specifies that a compute instance is to run on
	// hosts that only have guests from the same account.
	DedicatedAccountHostOnlyFlag bool `json:"dedicatedAccountHostOnlyFlag,omitempty"`

	// AccountId - A computing instance's associated [[SoftLayer_Account|account]] id
	AccountId int `json:"accountId,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// SupplementalCreateObjectOptions - Extra options needed for
	// [[SoftLayer_Virtual_Guest/createObject|createObject]] and
	// [[SoftLayer_Virtual_Guest/createObjects|createObjects]].
	SupplementalCreateObjectOptions *SoftLayer_Virtual_Guest_SupplementalCreateObjectOptions `json:"supplementalCreateObjectOptions,omitempty"`

	// FullyQualifiedDomainName - A name reflecting the hostname and domain of the computing instance.
	FullyQualifiedDomainName string `json:"fullyQualifiedDomainName,omitempty"`

	// Domain - no documentation
	Domain string `json:"domain,omitempty"`

	// LastPowerStateId - <nil>
	LastPowerStateId int `json:"lastPowerStateId,omitempty"`

	// ModifyDate - The date a virtual computing instance was last modified.
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// MetricPollDate - The date of the most recent metric tracking poll performed.
	MetricPollDate *time.Time `json:"metricPollDate,omitempty"`

	// LastVerifiedDate - The last timestamp of when the guest was verified as a resident virtual machine
	// on the host's hypervisor platform.
	LastVerifiedDate *time.Time `json:"lastVerifiedDate,omitempty"`

	// AllowedNetworkStorage - The SoftLayer_Network_Storage objects that this SoftLayer_Virtual_Guest has
	// access to.
	AllowedNetworkStorage []*SoftLayer_Network_Storage `json:"allowedNetworkStorage,omitempty"`

	// BlockCancelBecauseDisconnectedFlag - Determines whether the instance is ineligible for cancellation
	// because it is disconnected.
	BlockCancelBecauseDisconnectedFlag bool `json:"blockCancelBecauseDisconnectedFlag,omitempty"`

	// PrimaryBackendIpAddress - no documentation
	PrimaryBackendIpAddress string `json:"primaryBackendIpAddress,omitempty"`

	// BillingCycleBandwidthUsageCount - A count of the raw bandwidth usage data for the current billing
	// cycle. One object will be returned for each network this server is attached to.
	BillingCycleBandwidthUsageCount uint64 `json:"billingCycleBandwidthUsageCount,omitempty"`

	// OperatingSystemReferenceCode - no documentation
	OperatingSystemReferenceCode string `json:"operatingSystemReferenceCode,omitempty"`

	// PowerState - no documentation
	PowerState *SoftLayer_Virtual_Guest_Power_State `json:"powerState,omitempty"`

	// ActiveTransactions - Any active transaction(s) that are currently running for the server (example:
	// os reload).
	ActiveTransactions []*SoftLayer_Provisioning_Version1_Transaction `json:"activeTransactions,omitempty"`

	// BandwidthAllotmentDetail - A computing instance's allotted detail record. Allotment details link
	// bandwidth allocation with allotments.
	BandwidthAllotmentDetail *SoftLayer_Network_Bandwidth_Version1_Allotment_Detail `json:"bandwidthAllotmentDetail,omitempty"`

	// GuestBootParameter - <nil>
	GuestBootParameter *SoftLayer_Virtual_Guest_Boot_Parameter `json:"guestBootParameter,omitempty"`

	// NetworkMonitorIncidentCount - A count of all of a virtual guest's network monitoring incidents.
	NetworkMonitorIncidentCount uint64 `json:"networkMonitorIncidentCount,omitempty"`

	// BillingCyclePrivateBandwidthUsage - The raw private bandwidth usage data for the current billing
	// cycle.
	BillingCyclePrivateBandwidthUsage *SoftLayer_Network_Bandwidth_Usage `json:"billingCyclePrivateBandwidthUsage,omitempty"`

	// BillingCyclePublicBandwidthUsage - The raw public bandwidth usage data for the current billing
	// cycle.
	BillingCyclePublicBandwidthUsage *SoftLayer_Network_Bandwidth_Usage `json:"billingCyclePublicBandwidthUsage,omitempty"`

	// MonitoringRobot - <nil>
	MonitoringRobot *SoftLayer_Monitoring_Robot `json:"monitoringRobot,omitempty"`

	// RecentEvents - no documentation
	RecentEvents []*SoftLayer_Notification_Occurrence_Event `json:"recentEvents,omitempty"`

	// EvaultNetworkStorageCount - A count of a guest's associated EVault network storage service account.
	EvaultNetworkStorageCount uint64 `json:"evaultNetworkStorageCount,omitempty"`

	// MonitoringAgents - <nil>
	MonitoringAgents []*SoftLayer_Monitoring_Agent `json:"monitoringAgents,omitempty"`

	// ProjectedPublicBandwidthUsage - The projected public outbound bandwidth for this computing instance
	// for the current billing cycle.
	ProjectedPublicBandwidthUsage slapi.Float64 `json:"projectedPublicBandwidthUsage,omitempty"`

	// ActiveTicketCount - no documentation
	ActiveTicketCount uint64 `json:"activeTicketCount,omitempty"`

	// AllowedNetworkStorageCount - A count of the SoftLayer_Network_Storage objects that this
	// SoftLayer_Virtual_Guest has access to.
	AllowedNetworkStorageCount uint64 `json:"allowedNetworkStorageCount,omitempty"`

	// ActiveTickets - <nil>
	ActiveTickets []*SoftLayer_Ticket `json:"activeTickets,omitempty"`

	// MetricTrackingObjectId - no documentation
	MetricTrackingObjectId int `json:"metricTrackingObjectId,omitempty"`

	// PrimaryIpAddress - no documentation
	PrimaryIpAddress string `json:"primaryIpAddress,omitempty"`

	// VirtualRackName - The name of the bandwidth allotment that a computing instance belongs too.
	VirtualRackName string `json:"virtualRackName,omitempty"`

	// HourlyBillingFlag - Whether or not a computing instance is billed hourly instead of monthly.
	HourlyBillingFlag bool `json:"hourlyBillingFlag,omitempty"`

	// Location - Where guest is located within SoftLayer's location hierarchy.
	Location *SoftLayer_Location `json:"location,omitempty"`

	// NetworkComponentCount - no documentation
	NetworkComponentCount uint64 `json:"networkComponentCount,omitempty"`

	// SshKeyCount - A count of sSH keys to be installed on the server during provisioning or an OS reload.
	SshKeyCount uint64 `json:"sshKeyCount,omitempty"`

	// OperatingSystem - no documentation
	OperatingSystem *SoftLayer_Software_Component_OperatingSystem `json:"operatingSystem,omitempty"`

	// RegionalInternetRegistry - <nil>
	RegionalInternetRegistry *SoftLayer_Network_Regional_Internet_Registry `json:"regionalInternetRegistry,omitempty"`

	// VirtualRackId - The id of the bandwidth allotment that a computing instance belongs too.
	VirtualRackId int `json:"virtualRackId,omitempty"`

	// AllowedNetworkStorageReplicaCount - A count of the SoftLayer_Network_Storage objects whose Replica
	// that this SoftLayer_Virtual_Guest has access to.
	AllowedNetworkStorageReplicaCount uint64 `json:"allowedNetworkStorageReplicaCount,omitempty"`

	// Attributes - <nil>
	Attributes []*SoftLayer_Virtual_Guest_Attribute `json:"attributes,omitempty"`

	// ContinuousDataProtectionSoftwareComponent - A continuous data protection software component object.
	ContinuousDataProtectionSoftwareComponent *SoftLayer_Software_Component `json:"continuousDataProtectionSoftwareComponent,omitempty"`

	// GlobalIdentifier - no documentation
	GlobalIdentifier string `json:"globalIdentifier,omitempty"`

	// LastOperatingSystemReload - The last transaction that a cloud server's operating system was loaded.
	LastOperatingSystemReload *SoftLayer_Provisioning_Version1_Transaction `json:"lastOperatingSystemReload,omitempty"`

	// NetworkComponents - no documentation
	NetworkComponents []*SoftLayer_Virtual_Guest_Network_Component `json:"networkComponents,omitempty"`

	// BackendNetworkComponents - no documentation
	BackendNetworkComponents []*SoftLayer_Virtual_Guest_Network_Component `json:"backendNetworkComponents,omitempty"`

	// Datacenter - no documentation
	Datacenter *SoftLayer_Location `json:"datacenter,omitempty"`

	// OrderedPackageId - The original package id provided with the order for a Cloud Computing Instance.
	OrderedPackageId string `json:"orderedPackageId,omitempty"`

	// BandwidthAllocation - A computing instance's allotted bandwidth (measured in
	BandwidthAllocation slapi.Float64 `json:"bandwidthAllocation,omitempty"`

	// HostIpsSoftwareComponent - no documentation
	HostIpsSoftwareComponent *SoftLayer_Software_Component `json:"hostIpsSoftwareComponent,omitempty"`

	// OpenCancellationTicket - An open ticket requesting cancellation of this server, if one exists.
	OpenCancellationTicket *SoftLayer_Ticket `json:"openCancellationTicket,omitempty"`

	// RecentEventCount - A count of recent events that impact this computing instance.
	RecentEventCount uint64 `json:"recentEventCount,omitempty"`

	// ConsoleIpAddressRecord - A record containing information about a computing instance's console IP and
	// port number.
	ConsoleIpAddressRecord *SoftLayer_Virtual_Guest_Network_Component_IpAddress `json:"consoleIpAddressRecord,omitempty"`

	// OutboundPrivateBandwidthUsage - The total private outbound bandwidth for this computing instance for
	// the current billing cycle.
	OutboundPrivateBandwidthUsage slapi.Float64 `json:"outboundPrivateBandwidthUsage,omitempty"`

	// PrimaryNetworkComponent - no documentation
	PrimaryNetworkComponent *SoftLayer_Virtual_Guest_Network_Component `json:"primaryNetworkComponent,omitempty"`

	// ActiveTransactionCount - A count of any active transaction(s) that are currently running for the
	// server (example: os reload).
	ActiveTransactionCount uint64 `json:"activeTransactionCount,omitempty"`

	// NetworkStorageCount - A count of a guest's associated network storage accounts.
	NetworkStorageCount uint64 `json:"networkStorageCount,omitempty"`

	// ActiveNetworkMonitorIncident - A virtual guest's currently active network monitoring incidents.
	ActiveNetworkMonitorIncident []*SoftLayer_Network_Monitor_Version1_Incident `json:"activeNetworkMonitorIncident,omitempty"`

	// LastTransaction - no documentation
	LastTransaction *SoftLayer_Provisioning_Version1_Transaction `json:"lastTransaction,omitempty"`

	// PrivateNetworkOnlyFlag - Whether the computing instance only has access to the private network.
	PrivateNetworkOnlyFlag bool `json:"privateNetworkOnlyFlag,omitempty"`

	// SshKeys - SSH keys to be installed on the server during provisioning or an OS reload.
	SshKeys []*SoftLayer_Security_Ssh_Key `json:"sshKeys,omitempty"`

	// AllowedNetworkStorageReplicas - The SoftLayer_Network_Storage objects whose Replica that this
	// SoftLayer_Virtual_Guest has access to.
	AllowedNetworkStorageReplicas []*SoftLayer_Network_Storage `json:"allowedNetworkStorageReplicas,omitempty"`

	// InboundPublicBandwidthUsage - The total public inbound bandwidth for this computing instance for the
	// current billing cycle.
	InboundPublicBandwidthUsage slapi.Float64 `json:"inboundPublicBandwidthUsage,omitempty"`

	// ScaleMember - no documentation
	ScaleMember *SoftLayer_Scale_Member_Virtual_Guest `json:"scaleMember,omitempty"`

	// CurrentBandwidthSummary - An object that provides commonly used bandwidth summary components for the
	// current billing cycle.
	CurrentBandwidthSummary *SoftLayer_Metric_Tracking_Object_Bandwidth_Summary `json:"currentBandwidthSummary,omitempty"`

	// Users - A list of users that have access to this computing instance.
	Users []*SoftLayer_User_Customer `json:"users,omitempty"`

	// BackendNetworkComponentCount - no documentation
	BackendNetworkComponentCount uint64 `json:"backendNetworkComponentCount,omitempty"`

	// AllowedHost - The SoftLayer_Network_Storage_Allowed_Host information to connect this Virtual Guest
	// to Network Storage volumes that require access control lists.
	AllowedHost *SoftLayer_Network_Storage_Allowed_Host `json:"allowedHost,omitempty"`

	// OutboundPublicBandwidthUsage - The total public outbound bandwidth for this computing instance for
	// the current billing cycle.
	OutboundPublicBandwidthUsage slapi.Float64 `json:"outboundPublicBandwidthUsage,omitempty"`

	// BlockDeviceCount - A count of a computing instance's block devices. Block devices link
	// [[SoftLayer_Virtual_Disk_Image|disk images]] to computing instances.
	BlockDeviceCount uint64 `json:"blockDeviceCount,omitempty"`

	// UserCount - A count of a list of users that have access to this computing instance.
	UserCount uint64 `json:"userCount,omitempty"`

	// InboundPrivateBandwidthUsage - The total private inbound bandwidth for this computing instance for
	// the current billing cycle.
	InboundPrivateBandwidthUsage slapi.Float64 `json:"inboundPrivateBandwidthUsage,omitempty"`

	// NetworkVlans - The network Vlans that a guest's network components are associated with.
	NetworkVlans []*SoftLayer_Network_Vlan `json:"networkVlans,omitempty"`

	// TagReferenceCount - no documentation
	TagReferenceCount uint64 `json:"tagReferenceCount,omitempty"`

	// NetworkMonitorCount - no documentation
	NetworkMonitorCount uint64 `json:"networkMonitorCount,omitempty"`

	// UserDataCount - A count of a base64 encoded string containing custom user data for a Cloud Computing
	// Instance order.
	UserDataCount uint64 `json:"userDataCount,omitempty"`

	// EvaultNetworkStorage - A guest's associated EVault network storage service account.
	EvaultNetworkStorage []*SoftLayer_Network_Storage `json:"evaultNetworkStorage,omitempty"`

	// LastKnownPowerState - The last known power state of a virtual guest in the event the guest is turned
	// off outside of IMS or has gone offline.
	LastKnownPowerState *SoftLayer_Virtual_Guest_Power_State `json:"lastKnownPowerState,omitempty"`

	// LocalDiskFlag - A flag indicating that the virtual guest has at least one disk which is local to the
	// host it runs on. This does not include a device.
	LocalDiskFlag bool `json:"localDiskFlag,omitempty"`

	// ManagedResourceFlag - A flag indicating that the virtual guest is a managed resource.
	ManagedResourceFlag bool `json:"managedResourceFlag,omitempty"`

	// ControlPanel - no documentation
	ControlPanel *SoftLayer_Software_Component `json:"controlPanel,omitempty"`

	// NetworkMonitors - no documentation
	NetworkMonitors []*SoftLayer_Network_Monitor_Version1_Query_Host `json:"networkMonitors,omitempty"`

	// NetworkStorage - no documentation
	NetworkStorage []*SoftLayer_Network_Storage `json:"networkStorage,omitempty"`

	// PrimaryBackendNetworkComponent - no documentation
	PrimaryBackendNetworkComponent *SoftLayer_Virtual_Guest_Network_Component `json:"primaryBackendNetworkComponent,omitempty"`

	// ScaleAssets - Collection of scale assets this guest corresponds to.
	ScaleAssets []*SoftLayer_Scale_Asset `json:"scaleAssets,omitempty"`

	// ScaledFlag - Whether or not this guest is a member of a scale group and was automatically created as
	// part of a scale group action.
	ScaledFlag bool `json:"scaledFlag,omitempty"`

	// InternalTagReferenceCount - no documentation
	InternalTagReferenceCount uint64 `json:"internalTagReferenceCount,omitempty"`

	// MetricTrackingObject - no documentation
	MetricTrackingObject *SoftLayer_Metric_Tracking_Object `json:"metricTrackingObject,omitempty"`

	// MonitoringServiceEligibilityFlag - <nil>
	MonitoringServiceEligibilityFlag bool `json:"monitoringServiceEligibilityFlag,omitempty"`

	// SoftwareComponents - no documentation
	SoftwareComponents []*SoftLayer_Software_Component `json:"softwareComponents,omitempty"`

	// FrontendNetworkComponentCount - no documentation
	FrontendNetworkComponentCount uint64 `json:"frontendNetworkComponentCount,omitempty"`

	// AverageDailyPrivateBandwidthUsage - The average daily private bandwidth usage for the current
	// billing cycle.
	AverageDailyPrivateBandwidthUsage slapi.Float64 `json:"averageDailyPrivateBandwidthUsage,omitempty"`

	// FirewallServiceComponent - no documentation
	FirewallServiceComponent *SoftLayer_Network_Component_Firewall `json:"firewallServiceComponent,omitempty"`

	// VirtualRack - The name of the bandwidth allotment that a hardware belongs too.
	VirtualRack *SoftLayer_Network_Bandwidth_Version1_Allotment `json:"virtualRack,omitempty"`

	// ScaleAssetCount - A count of collection of scale assets this guest corresponds to.
	ScaleAssetCount uint64 `json:"scaleAssetCount,omitempty"`

	// ConsoleIpAddressFlag - A flag indicating a computing instance's console IP address is assigned.
	ConsoleIpAddressFlag bool `json:"consoleIpAddressFlag,omitempty"`

	// OverBandwidthAllocationFlag - Whether the bandwidth usage for this computing instance for the
	// current billing cycle exceeds the allocation.
	OverBandwidthAllocationFlag int `json:"overBandwidthAllocationFlag,omitempty"`

	// AttributeCount - no documentation
	AttributeCount uint64 `json:"attributeCount,omitempty"`

	// MonitoringAgentCount - no documentation
	MonitoringAgentCount uint64 `json:"monitoringAgentCount,omitempty"`

	// SecurityScanRequestCount - no documentation
	SecurityScanRequestCount uint64 `json:"securityScanRequestCount,omitempty"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account,omitempty"`

	// AntivirusSpywareSoftwareComponent - no documentation
	AntivirusSpywareSoftwareComponent *SoftLayer_Software_Component `json:"antivirusSpywareSoftwareComponent,omitempty"`

	// BlockDevices - A computing instance's block devices. Block devices link
	// [[SoftLayer_Virtual_Disk_Image|disk images]] to computing instances.
	BlockDevices []*SoftLayer_Virtual_Guest_Block_Device `json:"blockDevices,omitempty"`

	// UserData - A base64 encoded string containing custom user data for a Cloud Computing Instance order.
	UserData []*SoftLayer_Virtual_Guest_Attribute `json:"userData,omitempty"`

	// UpgradeRequest - A computing instance's associated upgrade request object if any.
	UpgradeRequest *SoftLayer_Product_Upgrade_Request `json:"upgradeRequest,omitempty"`

	// BillingCycleBandwidthUsage - The raw bandwidth usage data for the current billing cycle. One object
	// will be returned for each network this server is attached to.
	BillingCycleBandwidthUsage []*SoftLayer_Network_Bandwidth_Usage `json:"billingCycleBandwidthUsage,omitempty"`

	// BillingItem - The billing item for a CloudLayer Compute Instance.
	BillingItem *SoftLayer_Billing_Item_Virtual_Guest `json:"billingItem,omitempty"`

	// InternalTagReferences - <nil>
	InternalTagReferences []*SoftLayer_Tag_Reference `json:"internalTagReferences,omitempty"`

	// LatestNetworkMonitorIncident - A virtual guest's latest network monitoring incident.
	LatestNetworkMonitorIncident *SoftLayer_Network_Monitor_Version1_Incident `json:"latestNetworkMonitorIncident,omitempty"`

	// SecurityScanRequests - no documentation
	SecurityScanRequests []*SoftLayer_Network_Security_Scanner_Request `json:"securityScanRequests,omitempty"`

	// ActiveNetworkMonitorIncidentCount - A count of a virtual guest's currently active network monitoring
	// incidents.
	ActiveNetworkMonitorIncidentCount uint64 `json:"activeNetworkMonitorIncidentCount,omitempty"`

	// AvailableMonitoringCount - A count of an object that stores the maximum level for the monitoring
	// query types and response types.
	AvailableMonitoringCount uint64 `json:"availableMonitoringCount,omitempty"`

	// AccountOwnedPoolFlag - <nil>
	AccountOwnedPoolFlag bool `json:"accountOwnedPoolFlag,omitempty"`

	// AvailableMonitoring - An object that stores the maximum level for the monitoring query types and
	// response types.
	AvailableMonitoring []*SoftLayer_Network_Monitor_Version1_Query_Host_Stratum `json:"availableMonitoring,omitempty"`

	// FrontendNetworkComponents - no documentation
	FrontendNetworkComponents []*SoftLayer_Virtual_Guest_Network_Component `json:"frontendNetworkComponents,omitempty"`

	// Host - The virtual host on which a virtual guest resides (available only on private clouds).
	Host *SoftLayer_Virtual_Host `json:"host,omitempty"`

	// MonitoringServiceComponent - no documentation
	MonitoringServiceComponent *SoftLayer_Network_Monitor_Version1_Query_Host_Stratum `json:"monitoringServiceComponent,omitempty"`

	// SoftwareComponentCount - no documentation
	SoftwareComponentCount uint64 `json:"softwareComponentCount,omitempty"`

	// BackendRouters - no documentation
	BackendRouters []*SoftLayer_Hardware `json:"backendRouters,omitempty"`

	// ServerRoom - The server room that a guest is located at. There may be more than one server room for
	// every data center.
	ServerRoom *SoftLayer_Location `json:"serverRoom,omitempty"`

	// BackendRouterCount - no documentation
	BackendRouterCount uint64 `json:"backendRouterCount,omitempty"`

	// TagReferences - <nil>
	TagReferences []*SoftLayer_Tag_Reference `json:"tagReferences,omitempty"`

	// NetworkVlanCount - A count of the network Vlans that a guest's network components are associated
	// with.
	NetworkVlanCount uint64 `json:"networkVlanCount,omitempty"`

	// ActiveTransaction - A transaction that is still be performed on a cloud server.
	ActiveTransaction *SoftLayer_Provisioning_Version1_Transaction `json:"activeTransaction,omitempty"`

	// FrontendRouters - no documentation
	FrontendRouters *SoftLayer_Hardware `json:"frontendRouters,omitempty"`

	// MonitoringUserNotification - The monitoring notification objects for this guest. Each object links
	// this guest instance to a user account that will be notified if monitoring on this guest object fails
	MonitoringUserNotification []*SoftLayer_User_Customer_Notification_Virtual_Guest `json:"monitoringUserNotification,omitempty"`

	// RegionalGroup - no documentation
	RegionalGroup *SoftLayer_Location_Group_Regional `json:"regionalGroup,omitempty"`

	// Status - no documentation
	Status *SoftLayer_Virtual_Guest_Status `json:"status,omitempty"`

	// AverageDailyPublicBandwidthUsage - The average daily public bandwidth usage for the current billing
	// cycle.
	AverageDailyPublicBandwidthUsage slapi.Float64 `json:"averageDailyPublicBandwidthUsage,omitempty"`

	// BlockDeviceTemplateGroup - The global identifier for the image template that was used to provision a
	// guest.
	BlockDeviceTemplateGroup *SoftLayer_Virtual_Guest_Block_Device_Template_Group `json:"blockDeviceTemplateGroup,omitempty"`

	// ProjectedOverBandwidthAllocationFlag - Whether the bandwidth usage for this computing instance for
	// the current billing cycle is projected to exceed the allocation.
	ProjectedOverBandwidthAllocationFlag int `json:"projectedOverBandwidthAllocationFlag,omitempty"`

	// MonitoringUserNotificationCount - A count of the monitoring notification objects for this guest.
	// Each object links this guest instance to a user account that will be notified if monitoring on this
	// guest object fails
	MonitoringUserNotificationCount uint64 `json:"monitoringUserNotificationCount,omitempty"`

	// ApplicationDeliveryController - <nil>
	ApplicationDeliveryController *SoftLayer_Network_Application_Delivery_Controller `json:"applicationDeliveryController,omitempty"`

	// MonitoringServiceFlag - <nil>
	MonitoringServiceFlag bool `json:"monitoringServiceFlag,omitempty"`

	// NetworkMonitorIncidents - All of a virtual guest's network monitoring incidents.
	NetworkMonitorIncidents []*SoftLayer_Network_Monitor_Version1_Incident `json:"networkMonitorIncidents,omitempty"`
}

func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) String() string {
	return "SoftLayer_Virtual_Guest"
}
