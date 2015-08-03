package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Virtual_Guest - The virtual guest data type presents the structure in which all virtual
// guests will be presented. Internally, the structure supports various virtualization platforms with
// no change to external interaction. A guest, also known as a virtual server, represents an allocation
// of resources on a virtual host.
type SoftLayer_Virtual_Guest struct {

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// AccountId - A computing instance's associated [[SoftLayer_Account|account]] id
	AccountId int `json:"accountId"`

	// AccountOwnedPoolFlag - <nil>
	AccountOwnedPoolFlag bool `json:"accountOwnedPoolFlag"`

	// ActiveNetworkMonitorIncident - A virtual guest's currently active network monitoring incidents.
	ActiveNetworkMonitorIncident []*SoftLayer_Network_Monitor_Version1_Incident `json:"activeNetworkMonitorIncident"`

	// ActiveNetworkMonitorIncidentCount - A count of a virtual guest's currently active network monitoring
	// incidents.
	ActiveNetworkMonitorIncidentCount uint64 `json:"activeNetworkMonitorIncidentCount"`

	// ActiveTicketCount - no documentation
	ActiveTicketCount uint64 `json:"activeTicketCount"`

	// ActiveTickets - <nil>
	ActiveTickets []*SoftLayer_Ticket `json:"activeTickets"`

	// ActiveTransaction - A transaction that is still be performed on a cloud server.
	ActiveTransaction *SoftLayer_Provisioning_Version1_Transaction `json:"activeTransaction"`

	// ActiveTransactionCount - A count of any active transaction(s) that are currently running for the
	// server (example: os reload).
	ActiveTransactionCount uint64 `json:"activeTransactionCount"`

	// ActiveTransactions - Any active transaction(s) that are currently running for the server (example:
	// os reload).
	ActiveTransactions []*SoftLayer_Provisioning_Version1_Transaction `json:"activeTransactions"`

	// AllowedHost - The SoftLayer_Network_Storage_Allowed_Host information to connect this Virtual Guest
	// to Network Storage volumes that require access control lists.
	AllowedHost *SoftLayer_Network_Storage_Allowed_Host `json:"allowedHost"`

	// AllowedNetworkStorage - The SoftLayer_Network_Storage objects that this SoftLayer_Virtual_Guest has
	// access to.
	AllowedNetworkStorage []*SoftLayer_Network_Storage `json:"allowedNetworkStorage"`

	// AllowedNetworkStorageCount - A count of the SoftLayer_Network_Storage objects that this
	// SoftLayer_Virtual_Guest has access to.
	AllowedNetworkStorageCount uint64 `json:"allowedNetworkStorageCount"`

	// AllowedNetworkStorageReplicaCount - A count of the SoftLayer_Network_Storage objects whose Replica
	// that this SoftLayer_Virtual_Guest has access to.
	AllowedNetworkStorageReplicaCount uint64 `json:"allowedNetworkStorageReplicaCount"`

	// AllowedNetworkStorageReplicas - The SoftLayer_Network_Storage objects whose Replica that this
	// SoftLayer_Virtual_Guest has access to.
	AllowedNetworkStorageReplicas []*SoftLayer_Network_Storage `json:"allowedNetworkStorageReplicas"`

	// AntivirusSpywareSoftwareComponent - no documentation
	AntivirusSpywareSoftwareComponent *SoftLayer_Software_Component `json:"antivirusSpywareSoftwareComponent"`

	// ApplicationDeliveryController - <nil>
	ApplicationDeliveryController *SoftLayer_Network_Application_Delivery_Controller `json:"applicationDeliveryController"`

	// AttributeCount - no documentation
	AttributeCount uint64 `json:"attributeCount"`

	// Attributes - <nil>
	Attributes []*SoftLayer_Virtual_Guest_Attribute `json:"attributes"`

	// AvailableMonitoring - An object that stores the maximum level for the monitoring query types and
	// response types.
	AvailableMonitoring []*SoftLayer_Network_Monitor_Version1_Query_Host_Stratum `json:"availableMonitoring"`

	// AvailableMonitoringCount - A count of an object that stores the maximum level for the monitoring
	// query types and response types.
	AvailableMonitoringCount uint64 `json:"availableMonitoringCount"`

	// AverageDailyPrivateBandwidthUsage - The average daily private bandwidth usage for the current
	// billing cycle.
	AverageDailyPrivateBandwidthUsage float32 `json:"averageDailyPrivateBandwidthUsage"`

	// AverageDailyPublicBandwidthUsage - The average daily public bandwidth usage for the current billing
	// cycle.
	AverageDailyPublicBandwidthUsage float32 `json:"averageDailyPublicBandwidthUsage"`

	// BackendNetworkComponentCount - no documentation
	BackendNetworkComponentCount uint64 `json:"backendNetworkComponentCount"`

	// BackendNetworkComponents - no documentation
	BackendNetworkComponents []*SoftLayer_Virtual_Guest_Network_Component `json:"backendNetworkComponents"`

	// BackendRouterCount - no documentation
	BackendRouterCount uint64 `json:"backendRouterCount"`

	// BackendRouters - no documentation
	BackendRouters []*SoftLayer_Hardware `json:"backendRouters"`

	// BandwidthAllocation - A computing instance's allotted bandwidth (measured in
	BandwidthAllocation float64 `json:"bandwidthAllocation"`

	// BandwidthAllotmentDetail - A computing instance's allotted detail record. Allotment details link
	// bandwidth allocation with allotments.
	BandwidthAllotmentDetail *SoftLayer_Network_Bandwidth_Version1_Allotment_Detail `json:"bandwidthAllotmentDetail"`

	// BillingCycleBandwidthUsage - The raw bandwidth usage data for the current billing cycle. One object
	// will be returned for each network this server is attached to.
	BillingCycleBandwidthUsage []*SoftLayer_Network_Bandwidth_Usage `json:"billingCycleBandwidthUsage"`

	// BillingCycleBandwidthUsageCount - A count of the raw bandwidth usage data for the current billing
	// cycle. One object will be returned for each network this server is attached to.
	BillingCycleBandwidthUsageCount uint64 `json:"billingCycleBandwidthUsageCount"`

	// BillingCyclePrivateBandwidthUsage - The raw private bandwidth usage data for the current billing
	// cycle.
	BillingCyclePrivateBandwidthUsage *SoftLayer_Network_Bandwidth_Usage `json:"billingCyclePrivateBandwidthUsage"`

	// BillingCyclePublicBandwidthUsage - The raw public bandwidth usage data for the current billing
	// cycle.
	BillingCyclePublicBandwidthUsage *SoftLayer_Network_Bandwidth_Usage `json:"billingCyclePublicBandwidthUsage"`

	// BillingItem - The billing item for a CloudLayer Compute Instance.
	BillingItem *SoftLayer_Billing_Item_Virtual_Guest `json:"billingItem"`

	// BlockCancelBecauseDisconnectedFlag - Determines whether the instance is ineligible for cancellation
	// because it is disconnected.
	BlockCancelBecauseDisconnectedFlag bool `json:"blockCancelBecauseDisconnectedFlag"`

	// BlockDeviceCount - A count of a computing instance's block devices. Block devices link
	// [[SoftLayer_Virtual_Disk_Image|disk images]] to computing instances.
	BlockDeviceCount uint64 `json:"blockDeviceCount"`

	// BlockDeviceTemplateGroup - The global identifier for the image template that was used to provision a
	// guest.
	BlockDeviceTemplateGroup *SoftLayer_Virtual_Guest_Block_Device_Template_Group `json:"blockDeviceTemplateGroup"`

	// BlockDevices - A computing instance's block devices. Block devices link
	// [[SoftLayer_Virtual_Disk_Image|disk images]] to computing instances.
	BlockDevices []*SoftLayer_Virtual_Guest_Block_Device `json:"blockDevices"`

	// ConsoleIpAddressFlag - A flag indicating a computing instance's console IP address is assigned.
	ConsoleIpAddressFlag bool `json:"consoleIpAddressFlag"`

	// ConsoleIpAddressRecord - A record containing information about a computing instance's console IP and
	// port number.
	ConsoleIpAddressRecord *SoftLayer_Virtual_Guest_Network_Component_IpAddress `json:"consoleIpAddressRecord"`

	// ContinuousDataProtectionSoftwareComponent - A continuous data protection software component object.
	ContinuousDataProtectionSoftwareComponent *SoftLayer_Software_Component `json:"continuousDataProtectionSoftwareComponent"`

	// ControlPanel - no documentation
	ControlPanel *SoftLayer_Software_Component `json:"controlPanel"`

	// CreateDate - The date a virtual computing instance was created.
	CreateDate *time.Time `json:"createDate"`

	// CurrentBandwidthSummary - An object that provides commonly used bandwidth summary components for the
	// current billing cycle.
	CurrentBandwidthSummary *SoftLayer_Metric_Tracking_Object_Bandwidth_Summary `json:"currentBandwidthSummary"`

	// Datacenter - no documentation
	Datacenter *SoftLayer_Location `json:"datacenter"`

	// DedicatedAccountHostOnlyFlag - When true this flag specifies that a compute instance is to run on
	// hosts that only have guests from the same account.
	DedicatedAccountHostOnlyFlag bool `json:"dedicatedAccountHostOnlyFlag"`

	// Domain - no documentation
	Domain string `json:"domain"`

	// EvaultNetworkStorage - A guest's associated EVault network storage service account.
	EvaultNetworkStorage []*SoftLayer_Network_Storage `json:"evaultNetworkStorage"`

	// EvaultNetworkStorageCount - A count of a guest's associated EVault network storage service account.
	EvaultNetworkStorageCount uint64 `json:"evaultNetworkStorageCount"`

	// FirewallServiceComponent - no documentation
	FirewallServiceComponent *SoftLayer_Network_Component_Firewall `json:"firewallServiceComponent"`

	// FrontendNetworkComponentCount - no documentation
	FrontendNetworkComponentCount uint64 `json:"frontendNetworkComponentCount"`

	// FrontendNetworkComponents - no documentation
	FrontendNetworkComponents []*SoftLayer_Virtual_Guest_Network_Component `json:"frontendNetworkComponents"`

	// FrontendRouters - no documentation
	FrontendRouters *SoftLayer_Hardware `json:"frontendRouters"`

	// FullyQualifiedDomainName - A name reflecting the hostname and domain of the computing instance.
	FullyQualifiedDomainName string `json:"fullyQualifiedDomainName"`

	// GlobalIdentifier - no documentation
	GlobalIdentifier string `json:"globalIdentifier"`

	// GuestBootParameter - <nil>
	GuestBootParameter *SoftLayer_Virtual_Guest_Boot_Parameter `json:"guestBootParameter"`

	// Host - The virtual host on which a virtual guest resides (available only on private clouds).
	Host *SoftLayer_Virtual_Host `json:"host"`

	// HostIpsSoftwareComponent - no documentation
	HostIpsSoftwareComponent *SoftLayer_Software_Component `json:"hostIpsSoftwareComponent"`

	// Hostname - no documentation
	Hostname string `json:"hostname"`

	// HourlyBillingFlag - Whether or not a computing instance is billed hourly instead of monthly.
	HourlyBillingFlag bool `json:"hourlyBillingFlag"`

	// Id - no documentation
	Id int `json:"id"`

	// InboundPrivateBandwidthUsage - The total private inbound bandwidth for this computing instance for
	// the current billing cycle.
	InboundPrivateBandwidthUsage float64 `json:"inboundPrivateBandwidthUsage"`

	// InboundPublicBandwidthUsage - The total public inbound bandwidth for this computing instance for the
	// current billing cycle.
	InboundPublicBandwidthUsage float64 `json:"inboundPublicBandwidthUsage"`

	// InternalTagReferenceCount - no documentation
	InternalTagReferenceCount uint64 `json:"internalTagReferenceCount"`

	// InternalTagReferences - <nil>
	InternalTagReferences []*SoftLayer_Tag_Reference `json:"internalTagReferences"`

	// LastKnownPowerState - The last known power state of a virtual guest in the event the guest is turned
	// off outside of IMS or has gone offline.
	LastKnownPowerState *SoftLayer_Virtual_Guest_Power_State `json:"lastKnownPowerState"`

	// LastOperatingSystemReload - The last transaction that a cloud server's operating system was loaded.
	LastOperatingSystemReload *SoftLayer_Provisioning_Version1_Transaction `json:"lastOperatingSystemReload"`

	// LastPowerStateId - <nil>
	LastPowerStateId int `json:"lastPowerStateId"`

	// LastTransaction - no documentation
	LastTransaction *SoftLayer_Provisioning_Version1_Transaction `json:"lastTransaction"`

	// LastVerifiedDate - The last timestamp of when the guest was verified as a resident virtual machine
	// on the host's hypervisor platform.
	LastVerifiedDate *time.Time `json:"lastVerifiedDate"`

	// LatestNetworkMonitorIncident - A virtual guest's latest network monitoring incident.
	LatestNetworkMonitorIncident *SoftLayer_Network_Monitor_Version1_Incident `json:"latestNetworkMonitorIncident"`

	// LocalDiskFlag - A flag indicating that the virtual guest has at least one disk which is local to the
	// host it runs on. This does not include a device.
	LocalDiskFlag bool `json:"localDiskFlag"`

	// Location - Where guest is located within SoftLayer's location hierarchy.
	Location *SoftLayer_Location `json:"location"`

	// ManagedResourceFlag - A flag indicating that the virtual guest is a managed resource.
	ManagedResourceFlag bool `json:"managedResourceFlag"`

	// MaxCpu - The maximum amount of CPU resources a computing instance may utilize.
	MaxCpu int `json:"maxCpu"`

	// MaxCpuUnits - The unit of the maximum amount of CPU resources a computing instance may utilize.
	MaxCpuUnits string `json:"maxCpuUnits"`

	// MaxMemory - The maximum amount of memory a computing instance may utilize.
	MaxMemory int `json:"maxMemory"`

	// MetricPollDate - The date of the most recent metric tracking poll performed.
	MetricPollDate *time.Time `json:"metricPollDate"`

	// MetricTrackingObject - no documentation
	MetricTrackingObject *SoftLayer_Metric_Tracking_Object `json:"metricTrackingObject"`

	// MetricTrackingObjectId - no documentation
	MetricTrackingObjectId int `json:"metricTrackingObjectId"`

	// ModifyDate - The date a virtual computing instance was last modified.
	ModifyDate *time.Time `json:"modifyDate"`

	// MonitoringAgentCount - no documentation
	MonitoringAgentCount uint64 `json:"monitoringAgentCount"`

	// MonitoringAgents - <nil>
	MonitoringAgents []*SoftLayer_Monitoring_Agent `json:"monitoringAgents"`

	// MonitoringRobot - <nil>
	MonitoringRobot *SoftLayer_Monitoring_Robot `json:"monitoringRobot"`

	// MonitoringServiceComponent - no documentation
	MonitoringServiceComponent *SoftLayer_Network_Monitor_Version1_Query_Host_Stratum `json:"monitoringServiceComponent"`

	// MonitoringServiceEligibilityFlag - <nil>
	MonitoringServiceEligibilityFlag bool `json:"monitoringServiceEligibilityFlag"`

	// MonitoringServiceFlag - <nil>
	MonitoringServiceFlag bool `json:"monitoringServiceFlag"`

	// MonitoringUserNotification - The monitoring notification objects for this guest. Each object links
	// this guest instance to a user account that will be notified if monitoring on this guest object fails
	MonitoringUserNotification []*SoftLayer_User_Customer_Notification_Virtual_Guest `json:"monitoringUserNotification"`

	// MonitoringUserNotificationCount - A count of the monitoring notification objects for this guest.
	// Each object links this guest instance to a user account that will be notified if monitoring on this
	// guest object fails
	MonitoringUserNotificationCount uint64 `json:"monitoringUserNotificationCount"`

	// NetworkComponentCount - no documentation
	NetworkComponentCount uint64 `json:"networkComponentCount"`

	// NetworkComponents - no documentation
	NetworkComponents []*SoftLayer_Virtual_Guest_Network_Component `json:"networkComponents"`

	// NetworkMonitorCount - no documentation
	NetworkMonitorCount uint64 `json:"networkMonitorCount"`

	// NetworkMonitorIncidentCount - A count of all of a virtual guest's network monitoring incidents.
	NetworkMonitorIncidentCount uint64 `json:"networkMonitorIncidentCount"`

	// NetworkMonitorIncidents - All of a virtual guest's network monitoring incidents.
	NetworkMonitorIncidents []*SoftLayer_Network_Monitor_Version1_Incident `json:"networkMonitorIncidents"`

	// NetworkMonitors - no documentation
	NetworkMonitors []*SoftLayer_Network_Monitor_Version1_Query_Host `json:"networkMonitors"`

	// NetworkStorage - no documentation
	NetworkStorage []*SoftLayer_Network_Storage `json:"networkStorage"`

	// NetworkStorageCount - A count of a guest's associated network storage accounts.
	NetworkStorageCount uint64 `json:"networkStorageCount"`

	// NetworkVlanCount - A count of the network Vlans that a guest's network components are associated
	// with.
	NetworkVlanCount uint64 `json:"networkVlanCount"`

	// NetworkVlans - The network Vlans that a guest's network components are associated with.
	NetworkVlans []*SoftLayer_Network_Vlan `json:"networkVlans"`

	// Notes - A note of up to 1,000 characters about a virtual server.
	Notes string `json:"notes"`

	// OpenCancellationTicket - An open ticket requesting cancellation of this server, if one exists.
	OpenCancellationTicket *SoftLayer_Ticket `json:"openCancellationTicket"`

	// OperatingSystem - no documentation
	OperatingSystem *SoftLayer_Software_Component_OperatingSystem `json:"operatingSystem"`

	// OperatingSystemReferenceCode - no documentation
	OperatingSystemReferenceCode string `json:"operatingSystemReferenceCode"`

	// OrderedPackageId - The original package id provided with the order for a Cloud Computing Instance.
	OrderedPackageId string `json:"orderedPackageId"`

	// OutboundPrivateBandwidthUsage - The total private outbound bandwidth for this computing instance for
	// the current billing cycle.
	OutboundPrivateBandwidthUsage float64 `json:"outboundPrivateBandwidthUsage"`

	// OutboundPublicBandwidthUsage - The total public outbound bandwidth for this computing instance for
	// the current billing cycle.
	OutboundPublicBandwidthUsage float64 `json:"outboundPublicBandwidthUsage"`

	// OverBandwidthAllocationFlag - Whether the bandwidth usage for this computing instance for the
	// current billing cycle exceeds the allocation.
	OverBandwidthAllocationFlag int `json:"overBandwidthAllocationFlag"`

	// PostInstallScriptUri - URI of the script to be downloaded and executed after installation is
	// complete. This is deprecated in favor of supplementalCreateObjectOptions' postInstallScriptUri.
	PostInstallScriptUri string `json:"postInstallScriptUri"`

	// PowerState - no documentation
	PowerState *SoftLayer_Virtual_Guest_Power_State `json:"powerState"`

	// PrimaryBackendIpAddress - no documentation
	PrimaryBackendIpAddress string `json:"primaryBackendIpAddress"`

	// PrimaryBackendNetworkComponent - no documentation
	PrimaryBackendNetworkComponent *SoftLayer_Virtual_Guest_Network_Component `json:"primaryBackendNetworkComponent"`

	// PrimaryIpAddress - no documentation
	PrimaryIpAddress string `json:"primaryIpAddress"`

	// PrimaryNetworkComponent - no documentation
	PrimaryNetworkComponent *SoftLayer_Virtual_Guest_Network_Component `json:"primaryNetworkComponent"`

	// PrivateNetworkOnlyFlag - Whether the computing instance only has access to the private network.
	PrivateNetworkOnlyFlag bool `json:"privateNetworkOnlyFlag"`

	// ProjectedOverBandwidthAllocationFlag - Whether the bandwidth usage for this computing instance for
	// the current billing cycle is projected to exceed the allocation.
	ProjectedOverBandwidthAllocationFlag int `json:"projectedOverBandwidthAllocationFlag"`

	// ProjectedPublicBandwidthUsage - The projected public outbound bandwidth for this computing instance
	// for the current billing cycle.
	ProjectedPublicBandwidthUsage float32 `json:"projectedPublicBandwidthUsage"`

	// ProvisionDate - <nil>
	ProvisionDate *time.Time `json:"provisionDate"`

	// RecentEventCount - A count of recent events that impact this computing instance.
	RecentEventCount uint64 `json:"recentEventCount"`

	// RecentEvents - no documentation
	RecentEvents []*SoftLayer_Notification_Occurrence_Event `json:"recentEvents"`

	// RegionalGroup - no documentation
	RegionalGroup *SoftLayer_Location_Group_Regional `json:"regionalGroup"`

	// RegionalInternetRegistry - <nil>
	RegionalInternetRegistry *SoftLayer_Network_Regional_Internet_Registry `json:"regionalInternetRegistry"`

	// ScaleAssetCount - A count of collection of scale assets this guest corresponds to.
	ScaleAssetCount uint64 `json:"scaleAssetCount"`

	// ScaleAssets - Collection of scale assets this guest corresponds to.
	ScaleAssets []*SoftLayer_Scale_Asset `json:"scaleAssets"`

	// ScaleMember - no documentation
	ScaleMember *SoftLayer_Scale_Member_Virtual_Guest `json:"scaleMember"`

	// ScaledFlag - Whether or not this guest is a member of a scale group and was automatically created as
	// part of a scale group action.
	ScaledFlag bool `json:"scaledFlag"`

	// SecurityScanRequestCount - no documentation
	SecurityScanRequestCount uint64 `json:"securityScanRequestCount"`

	// SecurityScanRequests - no documentation
	SecurityScanRequests []*SoftLayer_Network_Security_Scanner_Request `json:"securityScanRequests"`

	// ServerRoom - The server room that a guest is located at. There may be more than one server room for
	// every data center.
	ServerRoom *SoftLayer_Location `json:"serverRoom"`

	// SoftwareComponentCount - no documentation
	SoftwareComponentCount uint64 `json:"softwareComponentCount"`

	// SoftwareComponents - no documentation
	SoftwareComponents []*SoftLayer_Software_Component `json:"softwareComponents"`

	// SshKeyCount - A count of sSH keys to be installed on the server during provisioning or an OS reload.
	SshKeyCount uint64 `json:"sshKeyCount"`

	// SshKeys - SSH keys to be installed on the server during provisioning or an OS reload.
	SshKeys []*SoftLayer_Security_Ssh_Key `json:"sshKeys"`

	// StartCpus - The number of CPUs available to a computing instance upon startup.
	StartCpus int `json:"startCpus"`

	// Status - no documentation
	Status *SoftLayer_Virtual_Guest_Status `json:"status"`

	// StatusId - A computing instances [[SoftLayer_Virtual_Guest_Status|status]] ID
	StatusId int `json:"statusId"`

	// SupplementalCreateObjectOptions - Extra options needed for
	// [[SoftLayer_Virtual_Guest/createObject|createObject]] and
	// [[SoftLayer_Virtual_Guest/createObjects|createObjects]].
	SupplementalCreateObjectOptions *SoftLayer_Virtual_Guest_SupplementalCreateObjectOptions `json:"supplementalCreateObjectOptions"`

	// TagReferenceCount - no documentation
	TagReferenceCount uint64 `json:"tagReferenceCount"`

	// TagReferences - <nil>
	TagReferences []*SoftLayer_Tag_Reference `json:"tagReferences"`

	// UpgradeRequest - A computing instance's associated upgrade request object if any.
	UpgradeRequest *SoftLayer_Product_Upgrade_Request `json:"upgradeRequest"`

	// UserCount - A count of a list of users that have access to this computing instance.
	UserCount uint64 `json:"userCount"`

	// UserData - A base64 encoded string containing custom user data for a Cloud Computing Instance order.
	UserData []*SoftLayer_Virtual_Guest_Attribute `json:"userData"`

	// UserDataCount - A count of a base64 encoded string containing custom user data for a Cloud Computing
	// Instance order.
	UserDataCount uint64 `json:"userDataCount"`

	// Users - A list of users that have access to this computing instance.
	Users []*SoftLayer_User_Customer `json:"users"`

	// Uuid - Unique ID for a computing instance's record on a virtualization platform.
	Uuid string `json:"uuid"`

	// VirtualRack - The name of the bandwidth allotment that a hardware belongs too.
	VirtualRack *SoftLayer_Network_Bandwidth_Version1_Allotment `json:"virtualRack"`

	// VirtualRackId - The id of the bandwidth allotment that a computing instance belongs too.
	VirtualRackId int `json:"virtualRackId"`

	// VirtualRackName - The name of the bandwidth allotment that a computing instance belongs too.
	VirtualRackName string `json:"virtualRackName"`
}

func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) String() string {
	return "SoftLayer_Virtual_Guest"
}

// ActivatePrivatePort - no documentation
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) ActivatePrivatePort(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// ActivatePublicPort - no documentation
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) ActivatePublicPort(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// AllowAccessToNetworkStorage - This method is used to allow access to a SoftLayer_Network_Storage
// volume that supports host- or network-level access control.
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) AllowAccessToNetworkStorage(ctx *slapi.RequestContext, networkStorageTemplateObject SoftLayer_Network_Storage) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// AllowAccessToNetworkStorageList - This method is used to allow access to multiple
// SoftLayer_Network_Storage volumes that support host- or network-level access control.
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) AllowAccessToNetworkStorageList(ctx *slapi.RequestContext, networkStorageTemplateObjects []SoftLayer_Network_Storage) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// AttachDiskImage - Creates a transaction to attach a guest's disk image. If the disk image is already
// attached it will be ignored. SoftLayer_Virtual_Guest::checkHostDiskAvailability should be called
// before this method. If the SoftLayer_Virtual_Guest::checkHostDiskAvailability method is not called
// before this method, the guest migration will happen automatically.
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) AttachDiskImage(ctx *slapi.RequestContext, imageId int) (*SoftLayer_Provisioning_Version1_Transaction, error) {
	var returnValue *SoftLayer_Provisioning_Version1_Transaction
	return returnValue, nil
}

// CancelIsolationForDestructiveAction - Reopens the public and/or private ports to reverse the changes
// made when the server was isolated for a destructive action.
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) CancelIsolationForDestructiveAction(ctx *slapi.RequestContext) error {
	return nil
}

// CaptureImage - Captures a Flex Image of the hard disk on the virtual machine, based on the capture
// template parameter. Returns the image template group containing the disk image.
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) CaptureImage(ctx *slapi.RequestContext, captureTemplate SoftLayer_Container_Disk_Image_Capture_Template) (*SoftLayer_Virtual_Guest_Block_Device_Template_Group, error) {
	var returnValue *SoftLayer_Virtual_Guest_Block_Device_Template_Group
	return returnValue, nil
}

// CheckHostDiskAvailability - Checks the associated host for available disk space to determine if
// guest migration is necessary. This method is only used with local disks. If this method returns
// false, calling attachDiskImage($imageId) will automatically migrate the destination guest to a new
// host before attaching the portable volume.
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) CheckHostDiskAvailability(ctx *slapi.RequestContext, diskCapacity int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// CloseAlarm - no documentation
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) CloseAlarm(ctx *slapi.RequestContext, alarmId string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// ConfigureMetadataDisk - Creates a transaction to configure the guest's metadata disk. If the guest
// has user data associated with it, the transaction will create a small virtual drive and write the
// metadata to a file on the drive; if the drive already exists, the metadata will be rewritten. If the
// guest has no user data associated with it, the transaction will remove the virtual drive if it
// exists. The transaction created by this service will shut down the guest while the metadata disk is
// configured. The guest will be turned back on once this process is complete.
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) ConfigureMetadataDisk(ctx *slapi.RequestContext) (*SoftLayer_Provisioning_Version1_Transaction, error) {
	var returnValue *SoftLayer_Provisioning_Version1_Transaction
	return returnValue, nil
}

// CreateArchiveTransaction - Create a transaction to archive a computing instance's block devices
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) CreateArchiveTransaction(ctx *slapi.RequestContext, groupName string, blockDevices []SoftLayer_Virtual_Guest_Block_Device, note string) (*SoftLayer_Provisioning_Version1_Transaction, error) {
	var returnValue *SoftLayer_Provisioning_Version1_Transaction
	return returnValue, nil
}

// CreateObject - createObject() enables the creation of computing instances on an account. This method
// is a simplified alternative to interacting with the ordering system directly. In order to create a
// computing instance, a template object must be sent in with a few required values. When this method
// returns an order will have been placed for a computing instance of the specified configuration. To
// determine when the instance is available you can poll the instance via
// [[SoftLayer_Virtual_Guest/getObject|getObject]], with an [[Extended-Object-Masks|object mask]]
// requesting the provisionDate relational property. When provisionDate is not null, the instance will
// be ready. Warning: Computing instances created via this method will incur charges on your account.
// For testing input parameters see
// [[SoftLayer_Virtual_Guest/generateOrderTemplate|generateOrderTemplate]]. Input -
// [[SoftLayer_Virtual_Guest (type)|SoftLayer_Virtual_Guest]] See
// [[SoftLayer_Virtual_Guest/getCreateObjectOptions|getCreateObjectOptions]] for available options. See
// [[SoftLayer_Virtual_Guest/getCreateObjectOptions|getCreateObjectOptions]] for available options.
// Specifies which datacenter the instance is to be provisioned in. The datacenter property is a
// [[SoftLayer_Location (type)|location]] structure with the name field set. See
// [[SoftLayer_Virtual_Guest/getCreateObjectOptions|getCreateObjectOptions]] for available options. {
// "datacenter": { "name": "dal05" } When true the computing instance will be billed on hourly usage,
// otherwise it will be billed on a monthly basis. When true the disks for the computing instance will
// be provisioned on the host which it runs, otherwise SAN disks will be provisioned. Specifies whether
// or not the instance must only run on hosts with instances from the same account Default - false When
// true this flag specifies that a compute instance is to run on hosts that only have guests from the
// same account. An identifier for the operating system to provision the computing instance with.
// Conditionally required - Disallowed when blockDeviceTemplateGroup.globalIdentifier is provided, as
// the template will specify the operating system. Notice - Some operating systems are charged based on
// the value specified in startCpus . The price which is used can be determined by calling
// [[SoftLayer_Virtual_Guest/generateOrderTemplate|generateOrderTemplate]] with your desired device
// specifications. See [[SoftLayer_Virtual_Guest/getCreateObjectOptions|getCreateObjectOptions]] for
// available options. A global identifier for the template to be used to provision the computing
// instance. Conditionally required - Disallowed when operatingSystemReferenceCode is provided, as the
// template will specify the operating system. Notice - Some operating systems are charged based on the
// value specified in startCpus . The price which is used can be determined by calling
// [[SoftLayer_Virtual_Guest/generateOrderTemplate|generateOrderTemplate]] with your desired device
// specifications. Both public and non-public images may be specified. A list of public images may be
// obtained via a request to
// [[SoftLayer_Virtual_Guest_Block_Device_Template_Group/getPublicImages|getPublicImages]]. A list of
// non-public images, images owned by an account or specifically shared with an account, may be
// obtained via a request to
// [[SoftLayer_Account/getBlockDeviceTemplateGroups|getBlockDeviceTemplateGroups]]. {
// "blockDeviceTemplateGroup": { "globalIdentifier": "07beadaa-1e11-476e-a188-3f7795feb9fb" } Specifies
// the connection speed for the instance's network components. Description - The networkComponents
// property is an array with a single [[SoftLayer_Virtual_Guest_Network_Component (type)|network
// component]] structure. The maxSpeed property must be set to specify the network uplink speed, in
// megabits per second, of the computing instance. See
// [[SoftLayer_Virtual_Guest/getCreateObjectOptions|getCreateObjectOptions]] for available options. {
// "networkComponents": [ { "maxSpeed": 1000 } ] Specifies whether or not the instance only has access
// to the private network Default - false When true this flag specifies that a compute instance is to
// only have access to the private network. Specifies the network vlan which is to be used for the
// frontend interface of the computing instance. Description - The primaryNetworkComponent property is
// a [[SoftLayer_Virtual_Guest_Network_Component (type)|network component]] structure with the
// networkVlan property populated with a [[SoftLayer_Network_Vlan (type)|vlan]] structure. The id
// property must be set to specify the frontend network vlan of the computing instance. {
// "primaryNetworkComponent": { "networkVlan": { "id": 1 } } Specifies the network vlan which is to be
// used for the backend interface of the computing instance. Description - The
// primaryBackendNetworkComponent property is a [[SoftLayer_Virtual_Guest_Network_Component
// (type)|network component]] structure with the networkVlan property populated with a
// [[SoftLayer_Network_Vlan (type)|vlan]] structure. The id property must be set to specify the backend
// network vlan of the computing instance. { "primaryBackendNetworkComponent": { "networkVlan": { "id":
// 2 } } Block device and disk image settings for the computing instance Type - array of
// [[SoftLayer_Virtual_Guest_Block_Device (type)|SoftLayer_Virtual_Guest_Block_Device] Default - The
// smallest available capacity for the primary disk will be used. If an image template is specified the
// disk capacity will be be provided by the template. Description - The blockDevices property is an
// array of [[SoftLayer_Virtual_Guest_Block_Device (type)|block device]] structures. Each block device
// must specify the device property along with the diskImage property, which is a
// [[SoftLayer_Virtual_Disk_Image (type)|disk image]] structure with the capacity property set. The
// device number '1' is reserved for the disk attached to the computing instance. See
// [[SoftLayer_Virtual_Guest/getCreateObjectOptions|getCreateObjectOptions]] for available options. {
// "blockDevices": [ { "device": "0", "diskImage": { "capacity": 100 } } ], "localDiskFlag": true
// Arbitrary data to be made available to the computing instance. Description - The userData property
// is an array with a single [[SoftLayer_Virtual_Guest_Attribute (type)|attribute]] structure with the
// value property set to an arbitrary value. This value can be retrieved via the
// [[SoftLayer_Resource_Metadata/getUserMetadata|getUserMetadata]] method from a request originating
// from the computing instance. This is primarily useful for providing data to software that may be on
// the instance and configured to execute upon first boot. { "userData": [ { "value": "someValue" } ]
// SSH keys to install on the computing instance upon provisioning. Type - array of
// [[SoftLayer_Security_Ssh_Key (type)|SoftLayer_Security_Ssh_Key]] Description - The sshKeys property
// is an array of [[SoftLayer_Security_Ssh_Key (type)|SSH Key]] structures with the id property set to
// the value of an existing SSH key. To create a new SSH key, call
// [[SoftLayer_Security_Ssh_Key/createObject|createObject]] on the [[SoftLayer_Security_Ssh_Key]]
// service. To obtain a list of existing SSH keys, call [[SoftLayer_Account/getSshKeys|getSshKeys]] on
// the [[SoftLayer_Account]] service. { "sshKeys": [ { "id": 123 } ] Specifies the uri location of the
// script to be downloaded and run after installation is complete. Example curl -X -d '{ "parameters":[
// { "hostname": "host1", "domain": "example.com", "startCpus": 1, "maxMemory": 1024,
// "hourlyBillingFlag": true, "localDiskFlag": true, "operatingSystemReferenceCode": } ] }'
// https://api.softlayer.com/rest/v3/SoftLayer_Virtual_Guest.json HTTP/1.1 201 Created Location:
// https://api.softlayer.com/rest/v3/SoftLayer_Virtual_Guest/1301396/getObject { "accountId": 232298,
// "createDate": "2012-11-30T16:28:17-06:00", "dedicatedAccountHostOnlyFlag": false, "domain":
// "example.com", "hostname": "host1", "id": 1301396, "lastPowerStateId": null, "lastVerifiedDate":
// null, "maxCpu": 1, "maxCpuUnits": "maxMemory": 1024, "metricPollDate": null, "modifyDate": null,
// "privateNetworkOnlyFlag": false, "startCpus": 1, "statusId": 1001, "globalIdentifier":
// "2d203774-0ee1-49f5-9599-6ef67358dd31" }
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) CreateObject(ctx *slapi.RequestContext, templateObject SoftLayer_Virtual_Guest) (*SoftLayer_Virtual_Guest, error) {
	var returnValue *SoftLayer_Virtual_Guest
	return returnValue, nil
}

// CreateObjects - createObjects() enables the creation of multiple computing instances on an account
// in a single call. This method is a simplified alternative to interacting with the ordering system
// directly. In order to create a computing instance a set of template objects must be sent in with a
// few required values. Warning: Computing instances created via this method will incur charges on your
// account. See [[SoftLayer_Virtual_Guest/createObject|createObject]] for specifics on the requirements
// of each template object. Example curl -X -d '{ "parameters":[ [ { "hostname": "host1", "domain":
// "example.com", "startCpus": 1, "maxMemory": 1024, "hourlyBillingFlag": true, "localDiskFlag": true,
// "operatingSystemReferenceCode": }, { "hostname": "host2", "domain": "example.com", "startCpus": 1,
// "maxMemory": 1024, "hourlyBillingFlag": true, "localDiskFlag": true, "operatingSystemReferenceCode":
// } ] ] }' https://api.softlayer.com/rest/v3/SoftLayer_Virtual_Guest/createObjects.json HTTP/1.1 200
// OK [ { "accountId": 232298, "createDate": "2012-11-30T23:56:48-06:00",
// "dedicatedAccountHostOnlyFlag": false, "domain": "softlayer.com", "hostname": "ubuntu1", "id":
// 1301456, "lastPowerStateId": null, "lastVerifiedDate": null, "maxCpu": 1, "maxCpuUnits":
// "maxMemory": 1024, "metricPollDate": null, "modifyDate": null, "privateNetworkOnlyFlag": false,
// "startCpus": 1, "statusId": 1001, "globalIdentifier": "fed4c822-48c0-45d0-85e2-90476aa0c542" }, {
// "accountId": 232298, "createDate": "2012-11-30T23:56:49-06:00", "dedicatedAccountHostOnlyFlag":
// false, "domain": "softlayer.com", "hostname": "ubuntu2", "id": 1301457, "lastPowerStateId": null,
// "lastVerifiedDate": null, "maxCpu": 1, "maxCpuUnits": "maxMemory": 1024, "metricPollDate": null,
// "modifyDate": null, "privateNetworkOnlyFlag": false, "startCpus": 1, "statusId": 1001,
// "globalIdentifier": "bed4c686-9562-4ade-9049-dc4d5b6b200c" } ]
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) CreateObjects(ctx *slapi.RequestContext, templateObjects []SoftLayer_Virtual_Guest) ([]*SoftLayer_Virtual_Guest, error) {
	var returnValue []*SoftLayer_Virtual_Guest
	return returnValue, nil
}

// CreatePostSoftwareInstallTransaction - <nil>
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) CreatePostSoftwareInstallTransaction(ctx *slapi.RequestContext, data string, returnBoolean bool) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// DeleteObject - This method will cancel a computing instance effective immediately. For instances
// billed hourly, the charges will stop immediately after the method returns.
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) DeleteObject(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// DetachDiskImage - Creates a transaction to detach a guest's disk image. If the disk image is already
// detached it will be ignored. The transaction created by this service will shut down the guest while
// the disk image is attached. The guest will be turned back on once this process is complete.
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) DetachDiskImage(ctx *slapi.RequestContext, imageId int) (*SoftLayer_Provisioning_Version1_Transaction, error) {
	var returnValue *SoftLayer_Provisioning_Version1_Transaction
	return returnValue, nil
}

// EditObject - no documentation
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) EditObject(ctx *slapi.RequestContext, templateObject SoftLayer_Virtual_Guest) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// ExecuteIderaBareMetalRestore - Reboot a guest into the Idera Bare Metal Restore image.
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) ExecuteIderaBareMetalRestore(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// ExecuteR1SoftBareMetalRestore - Reboot a guest into the R1Soft Bare Metal Restore image.
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) ExecuteR1SoftBareMetalRestore(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// ExecuteRescueLayer - no documentation
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) ExecuteRescueLayer(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// FindByIpAddress - Find CCI by only its primary public or private IP address. IP addresses within
// secondary subnets tied to the CCI will not return the If no CCI is found, no errors are generated
// and no data is returned.
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) FindByIpAddress(ctx *slapi.RequestContext, ipAddress string) (*SoftLayer_Virtual_Guest, error) {
	var returnValue *SoftLayer_Virtual_Guest
	return returnValue, nil
}

// GenerateOrderTemplate - Obtain an [[SoftLayer_Container_Product_Order_Virtual_Guest (type)|order
// container]] that can be sent to [[SoftLayer_Product_Order/verifyOrder|verifyOrder]] or
// [[SoftLayer_Product_Order/placeOrder|placeOrder]]. This is primarily useful when there is a
// necessity to confirm the price which will be charged for an order. See
// [[SoftLayer_Virtual_Guest/createObject|createObject]] for specifics on the requirements of the
// template object parameter.
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) GenerateOrderTemplate(ctx *slapi.RequestContext, templateObject SoftLayer_Virtual_Guest) (*SoftLayer_Container_Product_Order, error) {
	var returnValue *SoftLayer_Container_Product_Order
	return returnValue, nil
}

// GetAdditionalRequiredPricesForOsReload - Return a collection of SoftLayer_Item_Price objects for an
// OS reload
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) GetAdditionalRequiredPricesForOsReload(ctx *slapi.RequestContext, config SoftLayer_Container_Hardware_Server_Configuration) ([]*SoftLayer_Product_Item_Price, error) {
	var returnValue []*SoftLayer_Product_Item_Price
	return returnValue, nil
}

// GetAlarmHistory - no documentation
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) GetAlarmHistory(ctx *slapi.RequestContext, startDate time.Time, endDate time.Time, alarmId string) ([]*SoftLayer_Container_Monitoring_Alarm_History, error) {
	var returnValue []*SoftLayer_Container_Monitoring_Alarm_History
	return returnValue, nil
}

// GetAttachedNetworkStorages - This method is retrieve a list of SoftLayer_Network_Storage volumes
// that are authorized access to this SoftLayer_Virtual_Guest.
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) GetAttachedNetworkStorages(ctx *slapi.RequestContext, nasType string) ([]*SoftLayer_Network_Storage, error) {
	var returnValue []*SoftLayer_Network_Storage
	return returnValue, nil
}

// GetAvailableBlockDevicePositions - <nil>
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) GetAvailableBlockDevicePositions(ctx *slapi.RequestContext) ([]string, error) {
	var returnValue []string
	return returnValue, nil
}

// GetAvailableNetworkStorages - This method retrieves a list of SoftLayer_Network_Storage volumes that
// can be authorized to this SoftLayer_Virtual_Guest.
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) GetAvailableNetworkStorages(ctx *slapi.RequestContext, nasType string) ([]*SoftLayer_Network_Storage, error) {
	var returnValue []*SoftLayer_Network_Storage
	return returnValue, nil
}

// GetBandwidthDataByDate - Use this method when needing the metric data for bandwidth for a single
// guest. It will gather the correct input parameters based on the date ranges
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) GetBandwidthDataByDate(ctx *slapi.RequestContext, startDateTime time.Time, endDateTime time.Time, networkType string) ([]*SoftLayer_Metric_Tracking_Object_Data, error) {
	var returnValue []*SoftLayer_Metric_Tracking_Object_Data
	return returnValue, nil
}

// GetBandwidthForDateRange - Retrieve a collection of bandwidth data from an individual public or
// private network tracking object. Data is ideal if you with to employ your own traffic storage and
// graphing systems.
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) GetBandwidthForDateRange(ctx *slapi.RequestContext, startDate time.Time, endDate time.Time) ([]*SoftLayer_Metric_Tracking_Object_Data, error) {
	var returnValue []*SoftLayer_Metric_Tracking_Object_Data
	return returnValue, nil
}

// GetBandwidthImage - Use this method when needing a bandwidth image for a single guest. It will
// gather the correct input parameters for the generic graphing utility automatically based on the
// snapshot specified.
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) GetBandwidthImage(ctx *slapi.RequestContext, networkType string, snapshotRange string, dateSpecified time.Time) (*SoftLayer_Container_Bandwidth_GraphOutputs, error) {
	var returnValue *SoftLayer_Container_Bandwidth_GraphOutputs
	return returnValue, nil
}

// GetBandwidthImageByDate - Use this method when needing a bandwidth image for a single guest. It will
// gather the correct input parameters for the generic graphing utility based on the date ranges
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) GetBandwidthImageByDate(ctx *slapi.RequestContext, startDateTime time.Time, endDateTime time.Time, networkType string) (*SoftLayer_Container_Bandwidth_GraphOutputs, error) {
	var returnValue *SoftLayer_Container_Bandwidth_GraphOutputs
	return returnValue, nil
}

// GetBandwidthTotal - Returns the total amount of bandwidth used during the time specified for a
// computing instance.
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) GetBandwidthTotal(ctx *slapi.RequestContext, startDateTime time.Time, endDateTime time.Time, direction string, side string) (uint64, error) {
	var returnValue uint64
	return returnValue, nil
}

// GetBootOrder - <nil>
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) GetBootOrder(ctx *slapi.RequestContext) (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetConsoleAccessLog - Gets the console access logs for a computing instance
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) GetConsoleAccessLog(ctx *slapi.RequestContext) ([]*SoftLayer_Network_Logging_Syslog, error) {
	var returnValue []*SoftLayer_Network_Logging_Syslog
	return returnValue, nil
}

// GetCoreRestrictedOperatingSystemPrice - If the virtual server currently has an operating system that
// has a core capacity restriction, return the associated core-restricted operating system item price.
// Some operating systems (e.g., Red Hat Enterprise Linux) may be billed by the number of processor
// cores, so therefore require that a certain number of cores be present on the server.
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) GetCoreRestrictedOperatingSystemPrice(ctx *slapi.RequestContext) (*SoftLayer_Product_Item_Price, error) {
	var returnValue *SoftLayer_Product_Item_Price
	return returnValue, nil
}

// GetCpuMetricDataByDate - Use this method when needing the metric data for a single guest's CPUs. It
// will gather the correct input parameters based on the date ranges
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) GetCpuMetricDataByDate(ctx *slapi.RequestContext, startDateTime time.Time, endDateTime time.Time, cpuIndexes []int) ([]*SoftLayer_Metric_Tracking_Object_Data, error) {
	var returnValue []*SoftLayer_Metric_Tracking_Object_Data
	return returnValue, nil
}

// GetCpuMetricImage - Use this method when needing a cpu usage image for a single guest. It will
// gather the correct input parameters for the generic graphing utility automatically based on the
// snapshot specified.
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) GetCpuMetricImage(ctx *slapi.RequestContext, snapshotRange string, dateSpecified time.Time) (*SoftLayer_Container_Bandwidth_GraphOutputs, error) {
	var returnValue *SoftLayer_Container_Bandwidth_GraphOutputs
	return returnValue, nil
}

// GetCpuMetricImageByDate - Use this method when needing a CPU usage image for a single guest. It will
// gather the correct input parameters for the generic graphing utility based on the date ranges
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) GetCpuMetricImageByDate(ctx *slapi.RequestContext, startDateTime time.Time, endDateTime time.Time, cpuIndexes []int) (*SoftLayer_Container_Bandwidth_GraphOutputs, error) {
	var returnValue *SoftLayer_Container_Bandwidth_GraphOutputs
	return returnValue, nil
}

// GetCreateObjectOptions - There are many options that may be provided while ordering a computing
// instance, this method can be used to determine what these options are. Detailed information on the
// return value can be found on the data type page for
// [[SoftLayer_Container_Virtual_Guest_Configuration (type)]].
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) GetCreateObjectOptions(ctx *slapi.RequestContext) (*SoftLayer_Container_Virtual_Guest_Configuration, error) {
	var returnValue *SoftLayer_Container_Virtual_Guest_Configuration
	return returnValue, nil
}

// GetCurrentBillingDetail - getUpgradeItemPrices() retrieves a list of all upgrades available to a
// CloudLayer Computing Instance. Upgradeable items include, but are not limited to, number of cores,
// amount of storage configuration, and network port speed.
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) GetCurrentBillingDetail(ctx *slapi.RequestContext) ([]*SoftLayer_Billing_Item, error) {
	var returnValue []*SoftLayer_Billing_Item
	return returnValue, nil
}

// GetCurrentBillingTotal - Get the total billing price in US Dollars ($) for this instance. This
// includes all bandwidth used up to this point for this instance.
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) GetCurrentBillingTotal(ctx *slapi.RequestContext) (float64, error) {
	var returnValue float64
	return returnValue, nil
}

// GetCustomBandwidthDataByDate - no documentation
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) GetCustomBandwidthDataByDate(ctx *slapi.RequestContext, graphData SoftLayer_Container_Graph) (*SoftLayer_Container_Graph, error) {
	var returnValue *SoftLayer_Container_Graph
	return returnValue, nil
}

// GetCustomMetricDataByDate - no documentation
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) GetCustomMetricDataByDate(ctx *slapi.RequestContext, graphData SoftLayer_Container_Graph) (*SoftLayer_Container_Graph, error) {
	var returnValue *SoftLayer_Container_Graph
	return returnValue, nil
}

// GetDataMinMax - Returns 2 tracking object data records. The maximum and minimum dates that are
// available to received metric tracking data for a computing instance.
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) GetDataMinMax(ctx *slapi.RequestContext) ([]*SoftLayer_Metric_Tracking_Object_Data, error) {
	var returnValue []*SoftLayer_Metric_Tracking_Object_Data
	return returnValue, nil
}

// GetDriveRetentionItemPrice - Return a drive retention SoftLayer_Item_Price object for a guest.
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) GetDriveRetentionItemPrice(ctx *slapi.RequestContext) (*SoftLayer_Product_Item_Price, error) {
	var returnValue *SoftLayer_Product_Item_Price
	return returnValue, nil
}

// GetFirewallProtectableSubnets - Get the subnets associated with this CloudLayer computing instance
// that are protectable by a network component firewall.
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) GetFirewallProtectableSubnets(ctx *slapi.RequestContext) ([]*SoftLayer_Network_Subnet, error) {
	var returnValue []*SoftLayer_Network_Subnet
	return returnValue, nil
}

// GetFirstAvailableBlockDevicePosition - <nil>
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) GetFirstAvailableBlockDevicePosition(ctx *slapi.RequestContext) (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetIsoBootImage - <nil>
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) GetIsoBootImage(ctx *slapi.RequestContext) (*SoftLayer_Virtual_Disk_Image, error) {
	var returnValue *SoftLayer_Virtual_Disk_Image
	return returnValue, nil
}

// GetItemPricesFromSoftwareDescriptions - Return a collection of SoftLayer_Item_Price objects from a
// collection of SoftLayer_Software_Description
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) GetItemPricesFromSoftwareDescriptions(ctx *slapi.RequestContext, softwareDescriptions []SoftLayer_Software_Description, includeTranslationsFlag bool, returnAllPricesFlag bool) ([]*SoftLayer_Product_Item, error) {
	var returnValue []*SoftLayer_Product_Item
	return returnValue, nil
}

// GetMemoryMetricDataByDate - Use this method when needing the metric data for memory for a single
// computing instance.
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) GetMemoryMetricDataByDate(ctx *slapi.RequestContext, startDateTime time.Time, endDateTime time.Time) ([]*SoftLayer_Metric_Tracking_Object_Data, error) {
	var returnValue []*SoftLayer_Metric_Tracking_Object_Data
	return returnValue, nil
}

// GetMemoryMetricImage - Use this method when needing a memory usage image for a single guest. It will
// gather the correct input parameters for the generic graphing utility automatically based on the
// snapshot specified.
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) GetMemoryMetricImage(ctx *slapi.RequestContext, snapshotRange string, dateSpecified time.Time) (*SoftLayer_Container_Bandwidth_GraphOutputs, error) {
	var returnValue *SoftLayer_Container_Bandwidth_GraphOutputs
	return returnValue, nil
}

// GetMemoryMetricImageByDate - Use this method when needing a image displaying the amount of memory
// used over time for a single computing instance. It will gather the correct input parameters for the
// generic graphing utility based on the date ranges
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) GetMemoryMetricImageByDate(ctx *slapi.RequestContext, startDateTime time.Time, endDateTime time.Time) (*SoftLayer_Container_Bandwidth_GraphOutputs, error) {
	var returnValue *SoftLayer_Container_Bandwidth_GraphOutputs
	return returnValue, nil
}

// GetMonitoringActiveAlarms - Returns open monitoring alarms for a given time period
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) GetMonitoringActiveAlarms(ctx *slapi.RequestContext, startDate time.Time, endDate time.Time) ([]*SoftLayer_Container_Monitoring_Alarm_History, error) {
	var returnValue []*SoftLayer_Container_Monitoring_Alarm_History
	return returnValue, nil
}

// GetMonitoringClosedAlarms - Returns closed monitoring alarms for a given time period
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) GetMonitoringClosedAlarms(ctx *slapi.RequestContext, startDate time.Time, endDate time.Time) ([]*SoftLayer_Container_Monitoring_Alarm_History, error) {
	var returnValue []*SoftLayer_Container_Monitoring_Alarm_History
	return returnValue, nil
}

// GetNetworkComponentFirewallProtectableIpAddresses - Get the IP addresses associated with this
// CloudLayer computing instance that are protectable by a network component firewall. Note, this may
// not return all values for IPv6 subnets for this CloudLayer computing instance. Please use
// getFirewallProtectableSubnets to get all protectable subnets.
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) GetNetworkComponentFirewallProtectableIpAddresses(ctx *slapi.RequestContext) ([]*SoftLayer_Network_Subnet_IpAddress, error) {
	var returnValue []*SoftLayer_Network_Subnet_IpAddress
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Virtual_Guest, error) {
	var returnValue *SoftLayer_Virtual_Guest
	return returnValue, nil
}

// GetOrderTemplate - Obtain an order container that is ready to be sent to the
// [[SoftLayer_Product_Order#placeOrder|SoftLayer_Product_Order::placeOrder]] method. This container
// will include all services that the selected computing instance has. If desired you may remove prices
// which were returned.
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) GetOrderTemplate(ctx *slapi.RequestContext, billingType string, orderPrices []SoftLayer_Product_Item_Price) (*SoftLayer_Container_Product_Order, error) {
	var returnValue *SoftLayer_Container_Product_Order
	return returnValue, nil
}

// GetProvisionDate - <nil>
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) GetProvisionDate(ctx *slapi.RequestContext) (*time.Time, error) {
	var returnValue *time.Time
	return returnValue, nil
}

// GetRecentMetricData - no documentation
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) GetRecentMetricData(ctx *slapi.RequestContext, time uint) ([]*SoftLayer_Metric_Tracking_Object, error) {
	var returnValue []*SoftLayer_Metric_Tracking_Object
	return returnValue, nil
}

// GetRemoteMonitoringActiveAlarms - Returns open monitoring alarms generated by monitoring agents that
// reside in the SoftLayer monitoring cluster. A monitoring agent with "remoteMonitoringAgentFlag"
// indicates that it work from SoftLayer monitoring cluster. If a monitoring agent does not have the
// flag, it resides in your cloud instance.
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) GetRemoteMonitoringActiveAlarms(ctx *slapi.RequestContext, startDate time.Time, endDate time.Time) ([]*SoftLayer_Container_Monitoring_Alarm_History, error) {
	var returnValue []*SoftLayer_Container_Monitoring_Alarm_History
	return returnValue, nil
}

// GetRemoteMonitoringClosedAlarms - Returns closed monitoring alarms generated by monitoring agents
// that reside in the SoftLayer monitoring cluster. A monitoring agent with "remoteMonitoringAgentFlag"
// indicates that it work from SoftLayer monitoring cluster. If a monitoring agent does not have the
// flag, it resides in your cloud instance.
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) GetRemoteMonitoringClosedAlarms(ctx *slapi.RequestContext, startDate time.Time, endDate time.Time) ([]*SoftLayer_Container_Monitoring_Alarm_History, error) {
	var returnValue []*SoftLayer_Container_Monitoring_Alarm_History
	return returnValue, nil
}

// GetReverseDomainRecords - Retrieve the reverse domain records associated with this server.
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) GetReverseDomainRecords(ctx *slapi.RequestContext) ([]*SoftLayer_Dns_Domain, error) {
	var returnValue []*SoftLayer_Dns_Domain
	return returnValue, nil
}

// GetUpgradeItemPrices - getUpgradeItemPrices() retrieves a list of all upgrades available to a
// CloudLayer Computing Instance. Upgradeable items include, but are not limited to, number of cores,
// amount of storage configuration, and network port speed. This method exclude downgrade item prices
// by default. You can set the "includeDowngradeItemPrices" parameter to true so that it can include
// downgrade item prices.
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) GetUpgradeItemPrices(ctx *slapi.RequestContext, includeDowngradeItemPrices bool) ([]*SoftLayer_Product_Item_Price, error) {
	var returnValue []*SoftLayer_Product_Item_Price
	return returnValue, nil
}

// GetValidBlockDeviceTemplateGroups - This method will return the list of block device template groups
// that are valid to the host. For instance, it will validate that the template groups returned are
// compatible with the size and number of disks on the host.
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) GetValidBlockDeviceTemplateGroups(ctx *slapi.RequestContext, visibility string) ([]*SoftLayer_Virtual_Guest_Block_Device_Template_Group, error) {
	var returnValue []*SoftLayer_Virtual_Guest_Block_Device_Template_Group
	return returnValue, nil
}

// IsBackendPingable - Issues a ping command and returns the success (true) or failure (false) of the
// ping command.
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) IsBackendPingable(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// IsPingable - Issues a ping command and returns the success (true) or failure (false) of the ping
// command.
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) IsPingable(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// IsolateInstanceForDestructiveAction - Closes the public or private ports to isolate the instance
// before a destructive action.
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) IsolateInstanceForDestructiveAction(ctx *slapi.RequestContext) error {
	return nil
}

// MountIsoImage - <nil>
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) MountIsoImage(ctx *slapi.RequestContext, diskImageId int) (*SoftLayer_Provisioning_Version1_Transaction, error) {
	var returnValue *SoftLayer_Provisioning_Version1_Transaction
	return returnValue, nil
}

// Pause - no documentation
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) Pause(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// PowerCycle - no documentation
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) PowerCycle(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// PowerOff - no documentation
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) PowerOff(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// PowerOffSoft - no documentation
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) PowerOffSoft(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// PowerOn - no documentation
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) PowerOn(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RebootDefault - no documentation
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) RebootDefault(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RebootHard - no documentation
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) RebootHard(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RebootSoft - Attempt to complete a soft reboot of a guest by shutting down the operating system.
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) RebootSoft(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// ReloadCurrentOperatingSystemConfiguration - no documentation
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) ReloadCurrentOperatingSystemConfiguration(ctx *slapi.RequestContext) (*SoftLayer_Provisioning_Version1_Transaction, error) {
	var returnValue *SoftLayer_Provisioning_Version1_Transaction
	return returnValue, nil
}

// ReloadOperatingSystem - Reloads current operating system configuration. This service has a
// confirmation protocol for proceeding with the reload. To proceed with the reload without
// confirmation, simply pass in as the token parameter. To proceed with the reload with confirmation,
// simply call the service with no parameter. A token string will be returned by this service. The
// token will remain active for 10 minutes. Use this token as the parameter to confirm that a reload is
// to be performed for the server. As a precaution, we strongly recommend backing up all data before
// reloading the operating system. The reload will format the primary disk and will reconfigure the
// computing instance to the current specifications on record. If reloading from an image template, we
// recommend first getting the list of valid private block device template groups, by calling the
// getOperatingSystemReloadImages method.
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) ReloadOperatingSystem(ctx *slapi.RequestContext, token string, config SoftLayer_Container_Hardware_Server_Configuration) (string, error) {
	var returnValue string
	return returnValue, nil
}

// RemoveAccessToNetworkStorage - This method is used to remove access to a SoftLayer_Network_Storage
// volume that supports host- or network-level access control.
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) RemoveAccessToNetworkStorage(ctx *slapi.RequestContext, networkStorageTemplateObject SoftLayer_Network_Storage) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RemoveAccessToNetworkStorageList - This method is used to allow access to multiple
// SoftLayer_Network_Storage volumes that support host- or network-level access control.
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) RemoveAccessToNetworkStorageList(ctx *slapi.RequestContext, networkStorageTemplateObjects []SoftLayer_Network_Storage) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// Resume - no documentation
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) Resume(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// SetPrivateNetworkInterfaceSpeed - Sets the private network interface speed to the new speed. Speed
// values can only be 0 (Disconnect), 10, 100, or 1000. The new speed must be equal to or less than the
// max speed of the interface. It will take less than a minute to update the port speed.
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) SetPrivateNetworkInterfaceSpeed(ctx *slapi.RequestContext, newSpeed int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// SetPublicNetworkInterfaceSpeed - Sets the public network interface speed to the new speed. Speed
// values can only be 0 (Disconnect), 10, 100, or 1000. The new speed must be equal to or less than the
// max speed of the interface. It will take less than a minute to update the port speed.
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) SetPublicNetworkInterfaceSpeed(ctx *slapi.RequestContext, newSpeed int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// SetTags - <nil>
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) SetTags(ctx *slapi.RequestContext, tags string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// SetUserMetadata - Sets the data that will be written to the configuration drive.
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) SetUserMetadata(ctx *slapi.RequestContext, metadata []string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// ShutdownPrivatePort - no documentation
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) ShutdownPrivatePort(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// ShutdownPublicPort - no documentation
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) ShutdownPublicPort(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// UnmountIsoImage - <nil>
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) UnmountIsoImage(ctx *slapi.RequestContext) (*SoftLayer_Provisioning_Version1_Transaction, error) {
	var returnValue *SoftLayer_Provisioning_Version1_Transaction
	return returnValue, nil
}

// ValidateImageTemplate - no documentation
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) ValidateImageTemplate(ctx *slapi.RequestContext, imageTemplateId int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// VerifyReloadOperatingSystem - Verify that a virtual server can go through the operating system
// reload process. It may be useful to call this method before attempting to actually reload the
// operating system just to verify that the reload will go smoothly. If the server configuration is not
// setup correctly or there is some other issue, an exception will be thrown indicating the error. If
// there were no issues, this will just return true.
func (softlayer_virtual_guest *SoftLayer_Virtual_Guest) VerifyReloadOperatingSystem(ctx *slapi.RequestContext, config SoftLayer_Container_Hardware_Server_Configuration) (bool, error) {
	var returnValue bool
	return returnValue, nil
}
