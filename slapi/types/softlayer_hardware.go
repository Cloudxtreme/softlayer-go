package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
)

// SoftLayer_Hardware - The SoftLayer_Hardware data type contains general information relating to a
// single SoftLayer hardware.
type SoftLayer_Hardware struct {

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// AccountId - A hardware's associated [[SoftLayer_Account|account]] id.
	AccountId int `json:"accountId"`

	// ActiveComponentCount - A count of a piece of hardware's active physical components.
	ActiveComponentCount uint64 `json:"activeComponentCount"`

	// ActiveComponents - no documentation
	ActiveComponents []*SoftLayer_Hardware_Component `json:"activeComponents"`

	// ActiveNetworkMonitorIncident - A piece of hardware's active network monitoring incidents.
	ActiveNetworkMonitorIncident []*SoftLayer_Network_Monitor_Version1_Incident `json:"activeNetworkMonitorIncident"`

	// ActiveNetworkMonitorIncidentCount - A count of a piece of hardware's active network monitoring
	// incidents.
	ActiveNetworkMonitorIncidentCount uint64 `json:"activeNetworkMonitorIncidentCount"`

	// AllPowerComponentCount - no documentation
	AllPowerComponentCount uint64 `json:"allPowerComponentCount"`

	// AllPowerComponents - <nil>
	AllPowerComponents []*SoftLayer_Hardware_Power_Component `json:"allPowerComponents"`

	// AllowedHost - The SoftLayer_Network_Storage_Allowed_Host information to connect this server to
	// Network Storage volumes that require access control lists.
	AllowedHost *SoftLayer_Network_Storage_Allowed_Host `json:"allowedHost"`

	// AllowedNetworkStorage - The SoftLayer_Network_Storage objects that this SoftLayer_Hardware has
	// access to.
	AllowedNetworkStorage []*SoftLayer_Network_Storage `json:"allowedNetworkStorage"`

	// AllowedNetworkStorageCount - A count of the SoftLayer_Network_Storage objects that this
	// SoftLayer_Hardware has access to.
	AllowedNetworkStorageCount uint64 `json:"allowedNetworkStorageCount"`

	// AllowedNetworkStorageReplicaCount - A count of the SoftLayer_Network_Storage objects whose Replica
	// that this SoftLayer_Hardware has access to.
	AllowedNetworkStorageReplicaCount uint64 `json:"allowedNetworkStorageReplicaCount"`

	// AllowedNetworkStorageReplicas - The SoftLayer_Network_Storage objects whose Replica that this
	// SoftLayer_Hardware has access to.
	AllowedNetworkStorageReplicas []*SoftLayer_Network_Storage `json:"allowedNetworkStorageReplicas"`

	// AntivirusSpywareSoftwareComponent - Information regarding an antivirus/spyware software component
	// object.
	AntivirusSpywareSoftwareComponent *SoftLayer_Software_Component `json:"antivirusSpywareSoftwareComponent"`

	// AttributeCount - A count of information regarding a piece of hardware's specific attributes.
	AttributeCount uint64 `json:"attributeCount"`

	// Attributes - Information regarding a piece of hardware's specific attributes.
	Attributes []*SoftLayer_Hardware_Attribute `json:"attributes"`

	// AverageDailyPublicBandwidthUsage - The average daily public bandwidth usage for the current billing
	// cycle.
	AverageDailyPublicBandwidthUsage float32 `json:"averageDailyPublicBandwidthUsage"`

	// BackendNetworkComponentCount - A count of a piece of hardware's back-end or private network
	// components.
	BackendNetworkComponentCount uint64 `json:"backendNetworkComponentCount"`

	// BackendNetworkComponents - A piece of hardware's back-end or private network components.
	BackendNetworkComponents []*SoftLayer_Network_Component `json:"backendNetworkComponents"`

	// BackendRouterCount - no documentation
	BackendRouterCount uint64 `json:"backendRouterCount"`

	// BackendRouters - no documentation
	BackendRouters []*SoftLayer_Hardware `json:"backendRouters"`

	// BandwidthAllocation - no documentation
	BandwidthAllocation float64 `json:"bandwidthAllocation"`

	// BandwidthAllotmentDetail - A hardware's allotted detail record. Allotment details link bandwidth
	// allocation with allotments.
	BandwidthAllotmentDetail *SoftLayer_Network_Bandwidth_Version1_Allotment_Detail `json:"bandwidthAllotmentDetail"`

	// BareMetalInstanceFlag - When true, this flag specifies that a hardware is Bare Metal Server. Bare
	// Metal Servers are physical bare metal servers that are billed with the same options as Virtual
	// Servers, with monthly and hourly rates. Bare Metal instances are ordered based on processor core
	// count and ram amount.
	BareMetalInstanceFlag int `json:"bareMetalInstanceFlag"`

	// BenchmarkCertificationCount - A count of information regarding a piece of hardware's benchmark
	// certifications.
	BenchmarkCertificationCount uint64 `json:"benchmarkCertificationCount"`

	// BenchmarkCertifications - Information regarding a piece of hardware's benchmark certifications.
	BenchmarkCertifications []*SoftLayer_Hardware_Benchmark_Certification `json:"benchmarkCertifications"`

	// BillingItem - Information regarding the billing item for a server.
	BillingItem *SoftLayer_Billing_Item_Hardware `json:"billingItem"`

	// BillingItemFlag - no documentation
	BillingItemFlag bool `json:"billingItemFlag"`

	// BlockCancelBecauseDisconnectedFlag - Determines whether the hardware is ineligible for cancellation
	// because it is disconnected.
	BlockCancelBecauseDisconnectedFlag bool `json:"blockCancelBecauseDisconnectedFlag"`

	// BusinessContinuanceInsuranceFlag - Status indicating whether or not a piece of hardware has business
	// continuance insurance.
	BusinessContinuanceInsuranceFlag bool `json:"businessContinuanceInsuranceFlag"`

	// ComponentCount - no documentation
	ComponentCount uint64 `json:"componentCount"`

	// Components - no documentation
	Components []*SoftLayer_Hardware_Component `json:"components"`

	// ContinuousDataProtectionSoftwareComponent - A continuous data protection/server backup software
	// component object.
	ContinuousDataProtectionSoftwareComponent *SoftLayer_Software_Component `json:"continuousDataProtectionSoftwareComponent"`

	// CurrentBillableBandwidthUsage - The current billable public outbound bandwidth for this hardware for
	// the current billing cycle.
	CurrentBillableBandwidthUsage float64 `json:"currentBillableBandwidthUsage"`

	// Datacenter - Information regarding the datacenter in which a piece of hardware resides.
	Datacenter *SoftLayer_Location `json:"datacenter"`

	// DatacenterName - The name of the datacenter in which a piece of hardware resides.
	DatacenterName string `json:"datacenterName"`

	// Domain - no documentation
	Domain string `json:"domain"`

	// DownlinkHardware - All hardware that has uplink network connections to a piece of hardware.
	DownlinkHardware []*SoftLayer_Hardware `json:"downlinkHardware"`

	// DownlinkHardwareCount - A count of all hardware that has uplink network connections to a piece of
	// hardware.
	DownlinkHardwareCount uint64 `json:"downlinkHardwareCount"`

	// DownlinkNetworkHardware - All hardware that has uplink network connections to a piece of hardware.
	DownlinkNetworkHardware []*SoftLayer_Hardware `json:"downlinkNetworkHardware"`

	// DownlinkNetworkHardwareCount - A count of all hardware that has uplink network connections to a
	// piece of hardware.
	DownlinkNetworkHardwareCount uint64 `json:"downlinkNetworkHardwareCount"`

	// DownlinkServerCount - A count of information regarding all servers attached to a piece of network
	// hardware.
	DownlinkServerCount uint64 `json:"downlinkServerCount"`

	// DownlinkServers - Information regarding all servers attached to a piece of network hardware.
	DownlinkServers []*SoftLayer_Hardware `json:"downlinkServers"`

	// DownlinkVirtualGuestCount - A count of information regarding all virtual guests attached to a piece
	// of network hardware.
	DownlinkVirtualGuestCount uint64 `json:"downlinkVirtualGuestCount"`

	// DownlinkVirtualGuests - Information regarding all virtual guests attached to a piece of network
	// hardware.
	DownlinkVirtualGuests []*SoftLayer_Virtual_Guest `json:"downlinkVirtualGuests"`

	// DownstreamHardwareBindingCount - A count of all hardware downstream from a network device.
	DownstreamHardwareBindingCount uint64 `json:"downstreamHardwareBindingCount"`

	// DownstreamHardwareBindings - no documentation
	DownstreamHardwareBindings []*SoftLayer_Network_Component_Uplink_Hardware `json:"downstreamHardwareBindings"`

	// DownstreamNetworkHardware - All network hardware downstream from the selected piece of hardware.
	DownstreamNetworkHardware []*SoftLayer_Hardware `json:"downstreamNetworkHardware"`

	// DownstreamNetworkHardwareCount - A count of all network hardware downstream from the selected piece
	// of hardware.
	DownstreamNetworkHardwareCount uint64 `json:"downstreamNetworkHardwareCount"`

	// DownstreamNetworkHardwareWithIncidentCount - A count of all network hardware with monitoring
	// warnings or errors that are downstream from the selected piece of hardware.
	DownstreamNetworkHardwareWithIncidentCount uint64 `json:"downstreamNetworkHardwareWithIncidentCount"`

	// DownstreamNetworkHardwareWithIncidents - All network hardware with monitoring warnings or errors
	// that are downstream from the selected piece of hardware.
	DownstreamNetworkHardwareWithIncidents []*SoftLayer_Hardware `json:"downstreamNetworkHardwareWithIncidents"`

	// DownstreamServerCount - A count of information regarding all servers attached downstream to a piece
	// of network hardware.
	DownstreamServerCount uint64 `json:"downstreamServerCount"`

	// DownstreamServers - Information regarding all servers attached downstream to a piece of network
	// hardware.
	DownstreamServers []*SoftLayer_Hardware `json:"downstreamServers"`

	// DownstreamVirtualGuestCount - A count of information regarding all virtual guests attached to a
	// piece of network hardware.
	DownstreamVirtualGuestCount uint64 `json:"downstreamVirtualGuestCount"`

	// DownstreamVirtualGuests - Information regarding all virtual guests attached to a piece of network
	// hardware.
	DownstreamVirtualGuests []*SoftLayer_Virtual_Guest `json:"downstreamVirtualGuests"`

	// DriveControllerCount - A count of the drive controllers contained within a piece of hardware.
	DriveControllerCount uint64 `json:"driveControllerCount"`

	// DriveControllers - The drive controllers contained within a piece of hardware.
	DriveControllers []*SoftLayer_Hardware_Component `json:"driveControllers"`

	// EvaultNetworkStorage - Information regarding a piece of hardware's associated EVault network storage
	// service account.
	EvaultNetworkStorage []*SoftLayer_Network_Storage `json:"evaultNetworkStorage"`

	// EvaultNetworkStorageCount - A count of information regarding a piece of hardware's associated EVault
	// network storage service account.
	EvaultNetworkStorageCount uint64 `json:"evaultNetworkStorageCount"`

	// FirewallServiceComponent - Information regarding a piece of hardware's firewall services.
	FirewallServiceComponent *SoftLayer_Network_Component_Firewall `json:"firewallServiceComponent"`

	// FixedConfigurationPreset - Defines the fixed components in a fixed configuration bare metal server.
	FixedConfigurationPreset *SoftLayer_Product_Package_Preset `json:"fixedConfigurationPreset"`

	// FrontendNetworkComponentCount - A count of a piece of hardware's front-end or public network
	// components.
	FrontendNetworkComponentCount uint64 `json:"frontendNetworkComponentCount"`

	// FrontendNetworkComponents - A piece of hardware's front-end or public network components.
	FrontendNetworkComponents []*SoftLayer_Network_Component `json:"frontendNetworkComponents"`

	// FrontendRouterCount - no documentation
	FrontendRouterCount uint64 `json:"frontendRouterCount"`

	// FrontendRouters - no documentation
	FrontendRouters []*SoftLayer_Hardware `json:"frontendRouters"`

	// FullyQualifiedDomainName - A name reflecting the hostname and domain of the hardware. This is
	// created from the combined values of the hardware's hostname and domain name automatically, and thus
	// should not be edited directly.
	FullyQualifiedDomainName string `json:"fullyQualifiedDomainName"`

	// GlobalIdentifier - no documentation
	GlobalIdentifier string `json:"globalIdentifier"`

	// HardDriveCount - A count of the hard drives contained within a piece of hardware.
	HardDriveCount uint64 `json:"hardDriveCount"`

	// HardDrives - The hard drives contained within a piece of hardware.
	HardDrives []*SoftLayer_Hardware_Component `json:"hardDrives"`

	// HardwareChassis - no documentation
	HardwareChassis *SoftLayer_Hardware_Chassis `json:"hardwareChassis"`

	// HardwareFunction - no documentation
	HardwareFunction *SoftLayer_Hardware_Function `json:"hardwareFunction"`

	// HardwareFunctionDescription - no documentation
	HardwareFunctionDescription string `json:"hardwareFunctionDescription"`

	// HardwareStatus - no documentation
	HardwareStatus *SoftLayer_Hardware_Status `json:"hardwareStatus"`

	// HardwareStatusId - no documentation
	HardwareStatusId int `json:"hardwareStatusId"`

	// HasTrustedPlatformModuleBillingItemFlag - no documentation
	HasTrustedPlatformModuleBillingItemFlag bool `json:"hasTrustedPlatformModuleBillingItemFlag"`

	// HostIpsSoftwareComponent - Information regarding a host IPS software component object.
	HostIpsSoftwareComponent *SoftLayer_Software_Component `json:"hostIpsSoftwareComponent"`

	// Hostname - no documentation
	Hostname string `json:"hostname"`

	// HourlyBillingFlag - no documentation
	HourlyBillingFlag bool `json:"hourlyBillingFlag"`

	// Id - no documentation
	Id int `json:"id"`

	// InboundBandwidthUsage - The sum of all the inbound network traffic data for the last 30 days.
	InboundBandwidthUsage float64 `json:"inboundBandwidthUsage"`

	// InboundPublicBandwidthUsage - The total public inbound bandwidth for this hardware for the current
	// billing cycle.
	InboundPublicBandwidthUsage float64 `json:"inboundPublicBandwidthUsage"`

	// LastTransaction - Information regarding the last transaction a server performed.
	LastTransaction *SoftLayer_Provisioning_Version1_Transaction `json:"lastTransaction"`

	// LatestNetworkMonitorIncident - A piece of hardware's latest network monitoring incident.
	LatestNetworkMonitorIncident *SoftLayer_Network_Monitor_Version1_Incident `json:"latestNetworkMonitorIncident"`

	// Location - Where a piece of hardware is located within SoftLayer's location hierarchy.
	Location *SoftLayer_Location `json:"location"`

	// LocationPathString - <nil>
	LocationPathString string `json:"locationPathString"`

	// LockboxNetworkStorage - Information regarding a lockbox account associated with a server.
	LockboxNetworkStorage *SoftLayer_Network_Storage `json:"lockboxNetworkStorage"`

	// ManagedResourceFlag - A flag indicating that the hardware is a managed resource.
	ManagedResourceFlag bool `json:"managedResourceFlag"`

	// ManufacturerSerialNumber - A hardware's serial number that is supplied by the manufacturer.
	ManufacturerSerialNumber string `json:"manufacturerSerialNumber"`

	// Memory - Information regarding a piece of hardware's memory.
	Memory []*SoftLayer_Hardware_Component `json:"memory"`

	// MemoryCapacity - The amount of memory a piece of hardware has, measured in gigabytes.
	MemoryCapacity uint `json:"memoryCapacity"`

	// MemoryCount - A count of information regarding a piece of hardware's memory.
	MemoryCount uint64 `json:"memoryCount"`

	// MetricTrackingObject - no documentation
	MetricTrackingObject *SoftLayer_Metric_Tracking_Object_HardwareServer `json:"metricTrackingObject"`

	// MonitoringAgentCount - A count of information regarding the monitoring agents associated with a
	// piece of hardware.
	MonitoringAgentCount uint64 `json:"monitoringAgentCount"`

	// MonitoringAgents - Information regarding the monitoring agents associated with a piece of hardware.
	MonitoringAgents []*SoftLayer_Monitoring_Agent `json:"monitoringAgents"`

	// MonitoringRobot - Information regarding the hardware's monitoring robot.
	MonitoringRobot *SoftLayer_Monitoring_Robot `json:"monitoringRobot"`

	// MonitoringServiceComponent - Information regarding a piece of hardware's network monitoring
	// services.
	MonitoringServiceComponent *SoftLayer_Network_Monitor_Version1_Query_Host_Stratum `json:"monitoringServiceComponent"`

	// MonitoringServiceEligibilityFlag - The monitoring service flag eligibility status for a piece of
	// hardware.
	MonitoringServiceEligibilityFlag bool `json:"monitoringServiceEligibilityFlag"`

	// MonitoringServiceFlag - no documentation
	MonitoringServiceFlag bool `json:"monitoringServiceFlag"`

	// Motherboard - Information regarding a piece of hardware's motherboard.
	Motherboard *SoftLayer_Hardware_Component `json:"motherboard"`

	// NetworkCardCount - A count of information regarding a piece of hardware's network cards.
	NetworkCardCount uint64 `json:"networkCardCount"`

	// NetworkCards - Information regarding a piece of hardware's network cards.
	NetworkCards []*SoftLayer_Hardware_Component `json:"networkCards"`

	// NetworkComponentCount - A count of returns a hardware's network components.
	NetworkComponentCount uint64 `json:"networkComponentCount"`

	// NetworkComponents - no documentation
	NetworkComponents []*SoftLayer_Network_Component `json:"networkComponents"`

	// NetworkGatewayMember - The gateway member if this device is part of a network gateway.
	NetworkGatewayMember *SoftLayer_Network_Gateway_Member `json:"networkGatewayMember"`

	// NetworkGatewayMemberFlag - Whether or not this device is part of a network gateway.
	NetworkGatewayMemberFlag bool `json:"networkGatewayMemberFlag"`

	// NetworkManagementIpAddress - A piece of hardware's network management IP address.
	NetworkManagementIpAddress string `json:"networkManagementIpAddress"`

	// NetworkMonitorAttachedDownHardware - All servers with failed monitoring that are attached downstream
	// to a piece of hardware.
	NetworkMonitorAttachedDownHardware []*SoftLayer_Hardware `json:"networkMonitorAttachedDownHardware"`

	// NetworkMonitorAttachedDownHardwareCount - A count of all servers with failed monitoring that are
	// attached downstream to a piece of hardware.
	NetworkMonitorAttachedDownHardwareCount uint64 `json:"networkMonitorAttachedDownHardwareCount"`

	// NetworkMonitorAttachedDownVirtualGuestCount - A count of virtual guests that are attached downstream
	// to a hardware that have failed monitoring
	NetworkMonitorAttachedDownVirtualGuestCount uint64 `json:"networkMonitorAttachedDownVirtualGuestCount"`

	// NetworkMonitorAttachedDownVirtualGuests - Virtual guests that are attached downstream to a hardware
	// that have failed monitoring
	NetworkMonitorAttachedDownVirtualGuests []*SoftLayer_Virtual_Guest `json:"networkMonitorAttachedDownVirtualGuests"`

	// NetworkMonitorCount - A count of information regarding a piece of hardware's network monitors.
	NetworkMonitorCount uint64 `json:"networkMonitorCount"`

	// NetworkMonitorIncidentCount - A count of the status of all of a piece of hardware's network
	// monitoring incidents.
	NetworkMonitorIncidentCount uint64 `json:"networkMonitorIncidentCount"`

	// NetworkMonitorIncidents - The status of all of a piece of hardware's network monitoring incidents.
	NetworkMonitorIncidents []*SoftLayer_Network_Monitor_Version1_Incident `json:"networkMonitorIncidents"`

	// NetworkMonitors - Information regarding a piece of hardware's network monitors.
	NetworkMonitors []*SoftLayer_Network_Monitor_Version1_Query_Host `json:"networkMonitors"`

	// NetworkStatus - The value of a hardware's network status attribute.
	NetworkStatus string `json:"networkStatus"`

	// NetworkStatusAttribute - no documentation
	NetworkStatusAttribute *SoftLayer_Hardware_Attribute `json:"networkStatusAttribute"`

	// NetworkStorage - Information regarding a piece of hardware's associated network storage service
	// account.
	NetworkStorage []*SoftLayer_Network_Storage `json:"networkStorage"`

	// NetworkStorageCount - A count of information regarding a piece of hardware's associated network
	// storage service account.
	NetworkStorageCount uint64 `json:"networkStorageCount"`

	// NetworkVlanCount - A count of the network virtual LANs (VLANs) associated with a piece of hardware's
	// network components.
	NetworkVlanCount uint64 `json:"networkVlanCount"`

	// NetworkVlans - The network virtual LANs (VLANs) associated with a piece of hardware's network
	// components.
	NetworkVlans []*SoftLayer_Network_Vlan `json:"networkVlans"`

	// NextBillingCycleBandwidthAllocation - A hardware's allotted bandwidth for the next billing cycle
	// (measured in
	NextBillingCycleBandwidthAllocation float64 `json:"nextBillingCycleBandwidthAllocation"`

	// Notes - A small note about a piece of hardware to use at your discretion.
	Notes string `json:"notes"`

	// NotesHistory - <nil>
	NotesHistory []*SoftLayer_Hardware_Note `json:"notesHistory"`

	// NotesHistoryCount - no documentation
	NotesHistoryCount uint64 `json:"notesHistoryCount"`

	// OperatingSystem - Information regarding a piece of hardware's operating system.
	OperatingSystem *SoftLayer_Software_Component_OperatingSystem `json:"operatingSystem"`

	// OperatingSystemReferenceCode - A hardware's operating system software description.
	OperatingSystemReferenceCode string `json:"operatingSystemReferenceCode"`

	// OutboundBandwidthUsage - The sum of all the outbound network traffic data for the last 30 days.
	OutboundBandwidthUsage float64 `json:"outboundBandwidthUsage"`

	// OutboundPublicBandwidthUsage - The total public outbound bandwidth for this hardware for the current
	// billing cycle.
	OutboundPublicBandwidthUsage float64 `json:"outboundPublicBandwidthUsage"`

	// PointOfPresenceLocation - Information regarding the Point of Presence (PoP) location in which a
	// piece of hardware resides.
	PointOfPresenceLocation *SoftLayer_Location `json:"pointOfPresenceLocation"`

	// PostInstallScriptUri - URI of the script to be downloaded and executed after installation is
	// complete.
	PostInstallScriptUri string `json:"postInstallScriptUri"`

	// PowerComponentCount - A count of the power components for a hardware object.
	PowerComponentCount uint64 `json:"powerComponentCount"`

	// PowerComponents - no documentation
	PowerComponents []*SoftLayer_Hardware_Power_Component `json:"powerComponents"`

	// PowerSupply - Information regarding a piece of hardware's power supply.
	PowerSupply []*SoftLayer_Hardware_Component `json:"powerSupply"`

	// PowerSupplyCount - A count of information regarding a piece of hardware's power supply.
	PowerSupplyCount uint64 `json:"powerSupplyCount"`

	// PrimaryBackendIpAddress - no documentation
	PrimaryBackendIpAddress string `json:"primaryBackendIpAddress"`

	// PrimaryBackendNetworkComponent - Information regarding the hardware's primary back-end network
	// component.
	PrimaryBackendNetworkComponent *SoftLayer_Network_Component `json:"primaryBackendNetworkComponent"`

	// PrimaryIpAddress - no documentation
	PrimaryIpAddress string `json:"primaryIpAddress"`

	// PrimaryNetworkComponent - Information regarding the hardware's primary public network component.
	PrimaryNetworkComponent *SoftLayer_Network_Component `json:"primaryNetworkComponent"`

	// PrivateNetworkOnlyFlag - Whether the hardware only has access to the private network.
	PrivateNetworkOnlyFlag bool `json:"privateNetworkOnlyFlag"`

	// ProcessorCoreAmount - The total number of processor cores, summed from all processors that are
	// attached to a piece of hardware
	ProcessorCoreAmount uint `json:"processorCoreAmount"`

	// ProcessorCount - A count of information regarding a piece of hardware's processors.
	ProcessorCount uint64 `json:"processorCount"`

	// ProcessorPhysicalCoreAmount - The total number of physical processor cores, summed from all
	// processors that are attached to a piece of hardware
	ProcessorPhysicalCoreAmount uint `json:"processorPhysicalCoreAmount"`

	// Processors - Information regarding a piece of hardware's processors.
	Processors []*SoftLayer_Hardware_Component `json:"processors"`

	// ProvisionDate - <nil>
	ProvisionDate *time.Time `json:"provisionDate"`

	// Rack - <nil>
	Rack *SoftLayer_Location `json:"rack"`

	// RaidControllerCount - A count of the controllers contained within a piece of hardware.
	RaidControllerCount uint64 `json:"raidControllerCount"`

	// RaidControllers - The controllers contained within a piece of hardware.
	RaidControllers []*SoftLayer_Hardware_Component `json:"raidControllers"`

	// RecentEventCount - A count of recent events that impact this hardware.
	RecentEventCount uint64 `json:"recentEventCount"`

	// RecentEvents - no documentation
	RecentEvents []*SoftLayer_Notification_Occurrence_Event `json:"recentEvents"`

	// RemoteManagementAccountCount - A count of user credentials to issue commands and/or interact with
	// the server's remote management card.
	RemoteManagementAccountCount uint64 `json:"remoteManagementAccountCount"`

	// RemoteManagementAccounts - User credentials to issue commands and/or interact with the server's
	// remote management card.
	RemoteManagementAccounts []*SoftLayer_Hardware_Component_RemoteManagement_User `json:"remoteManagementAccounts"`

	// RemoteManagementComponent - A hardware's associated remote management component. This is normally
	RemoteManagementComponent *SoftLayer_Network_Component `json:"remoteManagementComponent"`

	// ResourceGroupCount - A count of the resource groups in which this hardware is a member.
	ResourceGroupCount uint64 `json:"resourceGroupCount"`

	// ResourceGroupMemberReferenceCount - no documentation
	ResourceGroupMemberReferenceCount uint64 `json:"resourceGroupMemberReferenceCount"`

	// ResourceGroupMemberReferences - <nil>
	ResourceGroupMemberReferences []*SoftLayer_Resource_Group_Member `json:"resourceGroupMemberReferences"`

	// ResourceGroupRoleCount - no documentation
	ResourceGroupRoleCount uint64 `json:"resourceGroupRoleCount"`

	// ResourceGroupRoles - <nil>
	ResourceGroupRoles []*SoftLayer_Resource_Group_Role `json:"resourceGroupRoles"`

	// ResourceGroups - The resource groups in which this hardware is a member.
	ResourceGroups []*SoftLayer_Resource_Group `json:"resourceGroups"`

	// RouterCount - no documentation
	RouterCount uint64 `json:"routerCount"`

	// Routers - no documentation
	Routers []*SoftLayer_Hardware `json:"routers"`

	// ScaleAssetCount - A count of collection of scale assets this hardware corresponds to.
	ScaleAssetCount uint64 `json:"scaleAssetCount"`

	// ScaleAssets - Collection of scale assets this hardware corresponds to.
	ScaleAssets []*SoftLayer_Scale_Asset `json:"scaleAssets"`

	// SecurityScanRequestCount - A count of information regarding a piece of hardware's vulnerability scan
	// requests.
	SecurityScanRequestCount uint64 `json:"securityScanRequestCount"`

	// SecurityScanRequests - Information regarding a piece of hardware's vulnerability scan requests.
	SecurityScanRequests []*SoftLayer_Network_Security_Scanner_Request `json:"securityScanRequests"`

	// SerialNumber - A hardware's serial number that is supplied by SoftLayer.
	SerialNumber string `json:"serialNumber"`

	// ServerRoom - Information regarding the server room in which the hardware is located.
	ServerRoom *SoftLayer_Location `json:"serverRoom"`

	// ServiceProvider - Information regarding the piece of hardware's service provider.
	ServiceProvider *SoftLayer_Service_Provider `json:"serviceProvider"`

	// ServiceProviderId - <nil>
	ServiceProviderId int `json:"serviceProviderId"`

	// ServiceProviderResourceId - A hardware's internal identification number at its service provider
	ServiceProviderResourceId int `json:"serviceProviderResourceId"`

	// SoftwareComponentCount - A count of information regarding a piece of hardware's installed software.
	SoftwareComponentCount uint64 `json:"softwareComponentCount"`

	// SoftwareComponents - Information regarding a piece of hardware's installed software.
	SoftwareComponents []*SoftLayer_Software_Component `json:"softwareComponents"`

	// SparePoolBillingItem - Information regarding the billing item for a spare pool server.
	SparePoolBillingItem *SoftLayer_Billing_Item_Hardware `json:"sparePoolBillingItem"`

	// SshKeyCount - A count of sSH keys to be installed on the server during provisioning or an OS reload.
	SshKeyCount uint64 `json:"sshKeyCount"`

	// SshKeys - SSH keys to be installed on the server during provisioning or an OS reload.
	SshKeys []*SoftLayer_Security_Ssh_Key `json:"sshKeys"`

	// StorageNetworkComponentCount - no documentation
	StorageNetworkComponentCount uint64 `json:"storageNetworkComponentCount"`

	// StorageNetworkComponents - <nil>
	StorageNetworkComponents []*SoftLayer_Network_Component `json:"storageNetworkComponents"`

	// TagReferenceCount - no documentation
	TagReferenceCount uint64 `json:"tagReferenceCount"`

	// TagReferences - <nil>
	TagReferences []*SoftLayer_Tag_Reference `json:"tagReferences"`

	// TopLevelLocation - <nil>
	TopLevelLocation *SoftLayer_Location `json:"topLevelLocation"`

	// UpgradeRequest - An account's associated upgrade request object, if any.
	UpgradeRequest *SoftLayer_Product_Upgrade_Request `json:"upgradeRequest"`

	// UplinkHardware - The network device connected to a piece of hardware.
	UplinkHardware *SoftLayer_Hardware `json:"uplinkHardware"`

	// UplinkNetworkComponentCount - A count of information regarding the network component that is one
	// level higher than a piece of hardware on the network infrastructure.
	UplinkNetworkComponentCount uint64 `json:"uplinkNetworkComponentCount"`

	// UplinkNetworkComponents - Information regarding the network component that is one level higher than
	// a piece of hardware on the network infrastructure.
	UplinkNetworkComponents []*SoftLayer_Network_Component `json:"uplinkNetworkComponents"`

	// UserData - A string containing custom user data for a hardware order.
	UserData []*SoftLayer_Hardware_Attribute `json:"userData"`

	// UserDataCount - A count of a string containing custom user data for a hardware order.
	UserDataCount uint64 `json:"userDataCount"`

	// VirtualChassis - Information regarding the virtual chassis for a piece of hardware.
	VirtualChassis *SoftLayer_Hardware_Group `json:"virtualChassis"`

	// VirtualChassisSiblingCount - A count of information regarding the virtual chassis siblings for a
	// piece of hardware.
	VirtualChassisSiblingCount uint64 `json:"virtualChassisSiblingCount"`

	// VirtualChassisSiblings - Information regarding the virtual chassis siblings for a piece of hardware.
	VirtualChassisSiblings []*SoftLayer_Hardware `json:"virtualChassisSiblings"`

	// VirtualHost - no documentation
	VirtualHost *SoftLayer_Virtual_Host `json:"virtualHost"`

	// VirtualLicenseCount - A count of information regarding a piece of hardware's virtual software
	// licenses.
	VirtualLicenseCount uint64 `json:"virtualLicenseCount"`

	// VirtualLicenses - Information regarding a piece of hardware's virtual software licenses.
	VirtualLicenses []*SoftLayer_Software_VirtualLicense `json:"virtualLicenses"`

	// VirtualRack - Information regarding the bandwidth allotment to which a piece of hardware belongs.
	VirtualRack *SoftLayer_Network_Bandwidth_Version1_Allotment `json:"virtualRack"`

	// VirtualRackId - The name of the bandwidth allotment belonging to a piece of hardware.
	VirtualRackId int `json:"virtualRackId"`

	// VirtualRackName - The name of the bandwidth allotment belonging to a piece of hardware.
	VirtualRackName string `json:"virtualRackName"`

	// VirtualizationPlatform - A piece of hardware's virtualization platform software.
	VirtualizationPlatform *SoftLayer_Software_Component `json:"virtualizationPlatform"`
}

// AllowAccessToNetworkStorage - This method is used to allow access to a SoftLayer_Network_Storage
// volume that supports host- or network-level access control.
func (softlayer_hardware *SoftLayer_Hardware) AllowAccessToNetworkStorage(commonOptions *slapi.CommonOptions, networkStorageTemplateObject SoftLayer_Network_Storage) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// AllowAccessToNetworkStorageList - This method is used to allow access to multiple
// SoftLayer_Network_Storage volumes that support host- or network-level access control.
func (softlayer_hardware *SoftLayer_Hardware) AllowAccessToNetworkStorageList(commonOptions *slapi.CommonOptions, networkStorageTemplateObjects []SoftLayer_Network_Storage) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// CaptureImage - Captures a Flex Image of the hard disk on the physical machine, based on the capture
// template parameter. Returns the image template group containing the disk image.
func (softlayer_hardware *SoftLayer_Hardware) CaptureImage(commonOptions *slapi.CommonOptions, captureTemplate SoftLayer_Container_Disk_Image_Capture_Template) (*SoftLayer_Virtual_Guest_Block_Device_Template_Group, error) {
	var returnValue *SoftLayer_Virtual_Guest_Block_Device_Template_Group
	return returnValue, nil
}

// CloseAlarm - no documentation
func (softlayer_hardware *SoftLayer_Hardware) CloseAlarm(commonOptions *slapi.CommonOptions, alarmId string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// CreateObject - createObject() enables the creation of servers on an account. This method is a
// simplified alternative to interacting with the ordering system directly. In order to create a
// server, a template object must be sent in with a few required values. When this method returns an
// order will have been placed for a server of the specified configuration. To determine when the
// server is available you can poll the server via [[SoftLayer_Hardware/getObject|getObject]], checking
// the provisionDate property. When provisionDate is not null, the server will be ready. Be sure to use
// the globalIdentifier as your initialization parameter. Warning: Servers created via this method will
// incur charges on your account. For testing input parameters see
// [[SoftLayer_Hardware/generateOrderTemplate|generateOrderTemplate]]. Input - [[SoftLayer_Hardware
// (type)|SoftLayer_Hardware]] See [[SoftLayer_Hardware/getCreateObjectOptions|getCreateObjectOptions]]
// for available options. See [[SoftLayer_Hardware/getCreateObjectOptions|getCreateObjectOptions]] for
// available options. When true the server will be billed on hourly usage, otherwise it will be billed
// on a monthly basis. An identifier for the operating system to provision the server with. See
// [[SoftLayer_Hardware/getCreateObjectOptions|getCreateObjectOptions]] for available options.
// Specifies which datacenter the server is to be provisioned in. The datacenter property is a
// [[SoftLayer_Location (type)|location]] structure with the name field set. See
// [[SoftLayer_Hardware/getCreateObjectOptions|getCreateObjectOptions]] for available options. {
// "datacenter": { "name": "dal05" } Specifies the connection speed for the server's network
// components. Default - The highest available zero cost port speed will be used. Description - The
// networkComponents property is an array with a single [[SoftLayer_Network_Component (type)|network
// component]] structure. The maxSpeed property must be set to specify the network uplink speed, in
// megabits per second, of the server. See
// [[SoftLayer_Hardware/getCreateObjectOptions|getCreateObjectOptions]] for available options. {
// "networkComponents": [ { "maxSpeed": 1000 } ] Specifies whether or not the server's network
// components should be in redundancy groups. Default - false Description - The networkComponents
// property is an array with a single [[SoftLayer_Network_Component (type)|network component]]
// structure. When the redundancyEnabledFlag property is true the server's network components will be
// in redundancy groups. { "networkComponents": [ { "redundancyEnabledFlag": false } ] Specifies
// whether or not the server only has access to the private network Default - false When true this flag
// specifies that a server is to only have access to the private network. Specifies the network vlan
// which is to be used for the frontend interface of the server. Description - The
// primaryNetworkComponent property is a [[SoftLayer_Network_Component (type)|network component]]
// structure with the networkVlan property populated with a [[SoftLayer_Network_Vlan (type)|vlan]]
// structure. The id property must be set to specify the frontend network vlan of the server. {
// "primaryNetworkComponent": { "networkVlan": { "id": 1 } } Specifies the network vlan which is to be
// used for the backend interface of the server. Description - The primaryBackendNetworkComponent
// property is a [[SoftLayer_Network_Component (type)|network component]] structure with the
// networkVlan property populated with a [[SoftLayer_Network_Vlan (type)|vlan]] structure. The id
// property must be set to specify the backend network vlan of the server. {
// "primaryBackendNetworkComponent": { "networkVlan": { "id": 2 } } Description - The
// fixedConfigurationPreset property is a [[SoftLayer_Product_Package_Preset (type)|fixed configuration
// preset]] structure. The keyName property must be set to specify preset to use. If a fixed
// configuration preset is used processorCoreAmount , memoryCapacity and hardDrives properties must not
// be set. See [[SoftLayer_Hardware/getCreateObjectOptions|getCreateObjectOptions]] for available
// options. { "fixedConfigurationPreset": { "keyName": } Description - The userData property is an
// array with a single [[SoftLayer_Hardware_Attribute (type)|attribute]] structure with the value
// property set to an arbitrary value. This value can be retrieved via the
// [[SoftLayer_Resource_Metadata/getUserMetadata|getUserMetadata]] method from a request originating
// from the server. This is primarily useful for providing data to software that may be on the server
// and configured to execute upon first boot. { "userData": [ { "value": "someValue" } ] Type -
// SoftLayer_Hardware_Component Default - The largest available capacity for a zero cost primary disk
// will be used. Description - The hardDrives property is an array of [[SoftLayer_Hardware_Component
// (type)|hardware component]] structures. Each hard drive must specify the capacity property. See
// [[SoftLayer_Hardware/getCreateObjectOptions|getCreateObjectOptions]] for available options. {
// "hardDrives": [ { "capacity": 500 } ] SSH keys to install on the server upon provisioning. Type -
// array of [[SoftLayer_Security_Ssh_Key (type)|SoftLayer_Security_Ssh_Key]] Description - The sshKeys
// property is an array of [[SoftLayer_Security_Ssh_Key (type)|SSH Key]] structures with the id
// property set to the value of an existing SSH key. To create a new SSH key, call
// [[SoftLayer_Security_Ssh_Key/createObject|createObject]] on the [[SoftLayer_Security_Ssh_Key]]
// service. To obtain a list of existing SSH keys, call [[SoftLayer_Account/getSshKeys|getSshKeys]] on
// the [[SoftLayer_Account]] service. { "sshKeys": [ { "id": 123 } ] Specifies the uri location of the
// script to be downloaded and run after installation is complete. Example curl -X -d '{ "parameters":[
// { "hostname": "host1", "domain": "example.com", "processorCoreAmount": 2, "memoryCapacity": 2,
// "hourlyBillingFlag": true, "operatingSystemReferenceCode": } ] }'
// https://api.softlayer.com/rest/v3/SoftLayer_Hardware.json HTTP/1.1 201 Created Location:
// https://api.softlayer.com/rest/v3/SoftLayer_Hardware/f5a3fcff-db1d-4b7c-9fa0-0349e41c29c5/getObject
// { "accountId": 232298, "bareMetalInstanceFlag": null, "domain": "example.com", "hardwareStatusId":
// null, "hostname": "host1", "id": null, "serviceProviderId": null, "serviceProviderResourceId": null,
// "globalIdentifier": "f5a3fcff-db1d-4b7c-9fa0-0349e41c29c5", "hourlyBillingFlag": true,
// "memoryCapacity": 2, "operatingSystemReferenceCode": "processorCoreAmount": 2 }
func (softlayer_hardware *SoftLayer_Hardware) CreateObject(commonOptions *slapi.CommonOptions, templateObject SoftLayer_Hardware) (*SoftLayer_Hardware, error) {
	var returnValue *SoftLayer_Hardware
	return returnValue, nil
}

// DeleteObject - This method will cancel a server effective immediately. For servers billed hourly,
// the charges will stop immediately after the method returns.
func (softlayer_hardware *SoftLayer_Hardware) DeleteObject(commonOptions *slapi.CommonOptions) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// DeleteSoftwareComponentPasswords - no documentation
func (softlayer_hardware *SoftLayer_Hardware) DeleteSoftwareComponentPasswords(commonOptions *slapi.CommonOptions, softwareComponentPasswords []SoftLayer_Software_Component_Password) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditSoftwareComponentPasswords - Edit the properties of a software component password such as the
// username, password, and notes.
func (softlayer_hardware *SoftLayer_Hardware) EditSoftwareComponentPasswords(commonOptions *slapi.CommonOptions, softwareComponentPasswords []SoftLayer_Software_Component_Password) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// FindByIpAddress - The '''findByIpAddress''' method finds hardware using its primary public or
// private IP address. IP addresses that have a secondary subnet tied to the hardware will not return
// the hardware - alternate means of locating the hardware must be used (see '''Associated Methods''').
// If no hardware is found, no errors are generated and no data is returned.
func (softlayer_hardware *SoftLayer_Hardware) FindByIpAddress(commonOptions *slapi.CommonOptions, ipAddress string) (*SoftLayer_Hardware, error) {
	var returnValue *SoftLayer_Hardware
	return returnValue, nil
}

// GenerateOrderTemplate - Obtain an [[SoftLayer_Container_Product_Order_Hardware_Server (type)|order
// container]] that can be sent to [[SoftLayer_Product_Order/verifyOrder|verifyOrder]] or
// [[SoftLayer_Product_Order/placeOrder|placeOrder]]. This is primarily useful when there is a
// necessity to confirm the price which will be charged for an order. See
// [[SoftLayer_Hardware/createObject|createObject]] for specifics on the requirements of the template
// object parameter.
func (softlayer_hardware *SoftLayer_Hardware) GenerateOrderTemplate(commonOptions *slapi.CommonOptions, templateObject SoftLayer_Hardware) (*SoftLayer_Container_Product_Order, error) {
	var returnValue *SoftLayer_Container_Product_Order
	return returnValue, nil
}

// GetAlarmHistory - The '''getAlarmHistory''' method retrieves a detailed history for the monitoring
// alarm. When calling this method, a start and end date for the history to be retrieved must be
// entered.
func (softlayer_hardware *SoftLayer_Hardware) GetAlarmHistory(commonOptions *slapi.CommonOptions, startDate time.Time, endDate time.Time, alarmId string) ([]*SoftLayer_Container_Monitoring_Alarm_History, error) {
	var returnValue []*SoftLayer_Container_Monitoring_Alarm_History
	return returnValue, nil
}

// GetAttachedNetworkStorages - This method is retrieve a list of SoftLayer_Network_Storage volumes
// that are authorized access to this SoftLayer_Hardware.
func (softlayer_hardware *SoftLayer_Hardware) GetAttachedNetworkStorages(commonOptions *slapi.CommonOptions, nasType string) ([]*SoftLayer_Network_Storage, error) {
	var returnValue []*SoftLayer_Network_Storage
	return returnValue, nil
}

// GetAvailableNetworkStorages - This method retrieves a list of SoftLayer_Network_Storage volumes that
// can be authorized to this SoftLayer_Hardware.
func (softlayer_hardware *SoftLayer_Hardware) GetAvailableNetworkStorages(commonOptions *slapi.CommonOptions, nasType string) ([]*SoftLayer_Network_Storage, error) {
	var returnValue []*SoftLayer_Network_Storage
	return returnValue, nil
}

// GetBackendIncomingBandwidth - The '''getBackendIncomingBandwidth''' method retrieves the amount of
// incoming private network traffic used between the given start date and end date parameters. When
// entering start and end dates, only the month, day and year are used to calculate bandwidth totals -
// the time is ignored and defaults to midnight. The amount of bandwidth retrieved is measured in
// gigabytes.
func (softlayer_hardware *SoftLayer_Hardware) GetBackendIncomingBandwidth(commonOptions *slapi.CommonOptions, startDate time.Time, endDate time.Time) (float32, error) {
	var returnValue float32
	return returnValue, nil
}

// GetBackendOutgoingBandwidth - The '''getBackendOutgoingBandwidth''' method retrieves the amount of
// outgoing private network traffic used between the given start date and end date parameters. When
// entering start and end dates, only the month, day and year are used to calculate bandwidth totals -
// the time is ignored and defaults to midnight. The amount of bandwidth retrieved is measured in
// gigabytes.
func (softlayer_hardware *SoftLayer_Hardware) GetBackendOutgoingBandwidth(commonOptions *slapi.CommonOptions, startDate time.Time, endDate time.Time) (float32, error) {
	var returnValue float32
	return returnValue, nil
}

// GetCreateObjectOptions - There are many options that may be provided while ordering a server, this
// method can be used to determine what these options are. Detailed information on the return value can
// be found on the data type page for [[SoftLayer_Container_Hardware_Configuration (type)]].
func (softlayer_hardware *SoftLayer_Hardware) GetCreateObjectOptions(commonOptions *slapi.CommonOptions) (*SoftLayer_Container_Hardware_Configuration, error) {
	var returnValue *SoftLayer_Container_Hardware_Configuration
	return returnValue, nil
}

// GetCurrentBillingDetail - Get the billing detail for this instance for the current billing period.
// This does not include bandwidth usage.
func (softlayer_hardware *SoftLayer_Hardware) GetCurrentBillingDetail(commonOptions *slapi.CommonOptions) ([]*SoftLayer_Billing_Item, error) {
	var returnValue []*SoftLayer_Billing_Item
	return returnValue, nil
}

// GetCurrentBillingTotal - The '''getCurrentBillingTotal''' method retrieves the total bill amount in
// US Dollars ($) for the current billing period. In addition to the total bill amount, the billing
// detail also includes all bandwidth used up to the point the method is called on the piece of
// hardware.
func (softlayer_hardware *SoftLayer_Hardware) GetCurrentBillingTotal(commonOptions *slapi.CommonOptions) (float64, error) {
	var returnValue float64
	return returnValue, nil
}

// GetDailyAverage - The '''getDailyAverage''' method calculates the average daily network traffic used
// by the selected server. Using the required parameter ''dateTime'' to enter a start and end date, the
// user retrieves this average, measure in gigabytes for the specified date range. When entering
// parameters, only the month, day and year are required - time entries are omitted as this method
// defaults the time to midnight in order to account for the entire day.
func (softlayer_hardware *SoftLayer_Hardware) GetDailyAverage(commonOptions *slapi.CommonOptions, startDate time.Time, endDate time.Time) (float32, error) {
	var returnValue float32
	return returnValue, nil
}

// GetFrontendIncomingBandwidth - The '''getFrontendIncomingBandwidth''' method retrieves the amount of
// incoming public network traffic used by a server between the given start and end date parameters.
// When entering the ''dateTime'' parameter, only the month, day and year of the start and end dates
// are required - the time (hour, minute and second) are set to midnight by default and cannot be
// changed. The amount of bandwidth retrieved is measured in gigabytes
func (softlayer_hardware *SoftLayer_Hardware) GetFrontendIncomingBandwidth(commonOptions *slapi.CommonOptions, startDate time.Time, endDate time.Time) (float32, error) {
	var returnValue float32
	return returnValue, nil
}

// GetFrontendOutgoingBandwidth - The '''getFrontendOutgoingBandwidth''' method retrieves the amount of
// outgoing public network traffic used by a server between the given start and end date parameters.
// The ''dateTime'' parameter requires only the day, month and year to be entered - the time (hour,
// minute and second) are set to midnight be default in order to gather the data for the entire start
// and end date indicated in the parameter. The amount of bandwidth retrieved is measured in gigabytes
func (softlayer_hardware *SoftLayer_Hardware) GetFrontendOutgoingBandwidth(commonOptions *slapi.CommonOptions, startDate time.Time, endDate time.Time) (float32, error) {
	var returnValue float32
	return returnValue, nil
}

// GetHourlyBandwidth - The '''getHourlyBandwidth''' method retrieves all bandwidth updates hourly for
// the specified hardware. Because the potential number of data points can become excessive, the method
// limits the user to obtain data in 24-hour intervals. The required ''dateTime'' parameter is used as
// the starting point for the query and will be calculated for the 24-hour period starting with the
// specified date and time. For example, entering a parameter of '02/01/2008 0:00' results in a return
// of all bandwidth data for the entire day of February 1, 2008, as 0:00 specifies a midnight start
// date. Please note that the time entered should be completed using a 24-hour clock (military time,
// astronomical time). For data spanning more than a single 24-hour period, refer to the
// getBandwidthData function on the metricTrackingObject for the piece of hardware.
func (softlayer_hardware *SoftLayer_Hardware) GetHourlyBandwidth(commonOptions *slapi.CommonOptions, mode string, day time.Time) ([]*SoftLayer_Metric_Tracking_Object_Data, error) {
	var returnValue []*SoftLayer_Metric_Tracking_Object_Data
	return returnValue, nil
}

// GetMonitoringActiveAlarms - Returns open monitoring alarms for a given time period
func (softlayer_hardware *SoftLayer_Hardware) GetMonitoringActiveAlarms(commonOptions *slapi.CommonOptions, startDate time.Time, endDate time.Time) ([]*SoftLayer_Container_Monitoring_Alarm_History, error) {
	var returnValue []*SoftLayer_Container_Monitoring_Alarm_History
	return returnValue, nil
}

// GetMonitoringClosedAlarms - Returns closed monitoring alarms for a given time period
func (softlayer_hardware *SoftLayer_Hardware) GetMonitoringClosedAlarms(commonOptions *slapi.CommonOptions, startDate time.Time, endDate time.Time) ([]*SoftLayer_Container_Monitoring_Alarm_History, error) {
	var returnValue []*SoftLayer_Container_Monitoring_Alarm_History
	return returnValue, nil
}

// GetObject - getObject retrieves the SoftLayer_Hardware object whose ID number corresponds to the ID
// number of the init parameter passed to the SoftLayer_Hardware service. You can only retrieve the
// account that your portal user is assigned to.
func (softlayer_hardware *SoftLayer_Hardware) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Hardware, error) {
	var returnValue *SoftLayer_Hardware
	return returnValue, nil
}

// GetPrivateBandwidthData - Retrieve a graph of a server's private network bandwidth usage over the
// specified timeframe. If no timeframe is specified then getPublicBandwidthGraphImage retrieves the
// last 24 hours of public bandwidth usage. getPrivateBandwidthGraphImage returns a PNG image measuring
// 827 pixels by 293 pixels.
func (softlayer_hardware *SoftLayer_Hardware) GetPrivateBandwidthData(commonOptions *slapi.CommonOptions, startTime int, endTime int) ([]*SoftLayer_Metric_Tracking_Object_Data, error) {
	var returnValue []*SoftLayer_Metric_Tracking_Object_Data
	return returnValue, nil
}

// GetPublicBandwidthData - Retrieve a graph of a server's public network bandwidth usage over the
// specified timeframe. If no timeframe is specified then getPublicBandwidthGraphImage retrieves the
// last 24 hours of public bandwidth usage. getPublicBandwidthGraphImage returns a PNG image measuring
// 827 pixels by 293 pixels.
func (softlayer_hardware *SoftLayer_Hardware) GetPublicBandwidthData(commonOptions *slapi.CommonOptions, startTime int, endTime int) ([]*SoftLayer_Metric_Tracking_Object_Data, error) {
	var returnValue []*SoftLayer_Metric_Tracking_Object_Data
	return returnValue, nil
}

// GetSensorData - The '''getSensorData''' method retrieves a server's hardware state via its internal
// sensors. Remote sensor data is transmitted to the SoftLayer API by way of the server's remote
// management card. Sensor data measures various information, including system temperatures, voltages
// and other local server settings. Sensor data is cached for 30 second; calls made to this method for
// the same server within 30 seconds of each other will result in the same data being returned. To
// ensure that the data retrieved retrieves snapshot of varied data, make calls greater than 30 seconds
// apart.
func (softlayer_hardware *SoftLayer_Hardware) GetSensorData(commonOptions *slapi.CommonOptions) ([]*SoftLayer_Container_RemoteManagement_SensorReading, error) {
	var returnValue []*SoftLayer_Container_RemoteManagement_SensorReading
	return returnValue, nil
}

// GetSensorDataWithGraphs - The '''getSensorDataWithGraphs''' method retrieves the raw data returned
// from the server's remote management card. Along with raw data, graphs for the CPU and system
// temperatures and fan speeds are also returned. For more details on what information is returned,
// refer to the ''getSensorData'' method.
func (softlayer_hardware *SoftLayer_Hardware) GetSensorDataWithGraphs(commonOptions *slapi.CommonOptions) (*SoftLayer_Container_RemoteManagement_SensorReadingsWithGraphs, error) {
	var returnValue *SoftLayer_Container_RemoteManagement_SensorReadingsWithGraphs
	return returnValue, nil
}

// GetServerFanSpeedGraphs - The '''getServerFanSpeedGraphs''' method retrieves the server's fan speeds
// and displays the speeds using tachometer graphs. data used to construct these graphs is retrieved
// from the server's remote management card. Each graph returned will have an associated title.
func (softlayer_hardware *SoftLayer_Hardware) GetServerFanSpeedGraphs(commonOptions *slapi.CommonOptions) ([]*SoftLayer_Container_RemoteManagement_Graphs_SensorSpeed, error) {
	var returnValue []*SoftLayer_Container_RemoteManagement_Graphs_SensorSpeed
	return returnValue, nil
}

// GetServerPowerState - The '''getPowerState''' method retrieves the power state for the selected
// server. The server's power status is retrieved from its remote management card. This method returns
// "on", for a server that has been powered on, or "off" for servers powered off.
func (softlayer_hardware *SoftLayer_Hardware) GetServerPowerState(commonOptions *slapi.CommonOptions) (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetServerTemperatureGraphs - The '''getServerTemperatureGraphs''' retrieves the server's
// temperatures and displays the various temperatures using thermometer graphs. Temperatures retrieved
// are CPU temperature(s) and system temperatures. Data used to construct the graphs is retrieved from
// the server's remote management card. All graphs returned will have an associated title.
func (softlayer_hardware *SoftLayer_Hardware) GetServerTemperatureGraphs(commonOptions *slapi.CommonOptions) ([]*SoftLayer_Container_RemoteManagement_Graphs_SensorTemperature, error) {
	var returnValue []*SoftLayer_Container_RemoteManagement_Graphs_SensorTemperature
	return returnValue, nil
}

// GetUpgradeItemPrices - Retrieve a list of upgradeable items available to this piece of hardware.
// Currently, getUpgradeItemPrices retrieves upgrades available for a server's memory, hard drives,
// network port speed, bandwidth allocation and GPUs.
func (softlayer_hardware *SoftLayer_Hardware) GetUpgradeItemPrices(commonOptions *slapi.CommonOptions) ([]*SoftLayer_Product_Item_Price, error) {
	var returnValue []*SoftLayer_Product_Item_Price
	return returnValue, nil
}

// ImportVirtualHost - The '''importVirtualHost''' method attempts to import the host record for the
// virtualization platform running on a server.
func (softlayer_hardware *SoftLayer_Hardware) ImportVirtualHost(commonOptions *slapi.CommonOptions) (*SoftLayer_Virtual_Host, error) {
	var returnValue *SoftLayer_Virtual_Host
	return returnValue, nil
}

// IsPingable - The '''isPingable''' method issues a ping command to the selected server and returns
// the result of the ping command. This boolean return value displays ''true'' upon successful ping or
// ''false'' for a failed ping.
func (softlayer_hardware *SoftLayer_Hardware) IsPingable(commonOptions *slapi.CommonOptions) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// Ping - Issues a ping command to the server and returns the ping response.
func (softlayer_hardware *SoftLayer_Hardware) Ping(commonOptions *slapi.CommonOptions) (string, error) {
	var returnValue string
	return returnValue, nil
}

// PowerCycle - The '''powerCycle''' method completes a power off and power on of the server
// successively in one command. The power cycle command is equivalent to unplugging the server from the
// power strip and then plugging the server back in. '''This method should only be used when all other
// options have been exhausted'''. Additional remote management commands may not be executed if this
// command was successfully issued within the last 20 minutes to avoid server failure. Remote
// management commands include: rebootSoft rebootHard powerOn powerOff powerCycle
func (softlayer_hardware *SoftLayer_Hardware) PowerCycle(commonOptions *slapi.CommonOptions) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// PowerOff - This method will power off the server via the server's remote management card.
func (softlayer_hardware *SoftLayer_Hardware) PowerOff(commonOptions *slapi.CommonOptions) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// PowerOn - The '''powerOn''' method powers on a server via its remote management card. This boolean
// return value returns ''true'' upon successful execution and ''false'' if unsuccessful. Other remote
// management commands may not be issued in this command was successfully completed within the last 20
// minutes to avoid server failure. Remote management commands include: rebootSoft rebootHard powerOn
// powerOff powerCycle
func (softlayer_hardware *SoftLayer_Hardware) PowerOn(commonOptions *slapi.CommonOptions) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RebootDefault - The '''rebootDefault''' method attempts to reboot the server by issuing a soft
// reboot, or reset, command to the server's remote management card. if the reset attempt is
// unsuccessful, a power cycle command will be issued via the power strip. The power cycle command is
// equivalent to unplugging the server from the power strip and then plugging the server back in. If
// the reset was successful within the last 20 minutes, another remote management command cannot be
// completed to avoid server failure. Remote management commands include: rebootSoft rebootHard powerOn
// powerOff powerCycle
func (softlayer_hardware *SoftLayer_Hardware) RebootDefault(commonOptions *slapi.CommonOptions) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RebootHard - The '''rebootHard''' method reboots the server by issuing a cycle command to the
// server's remote management card. A hard reboot is equivalent to pressing the ''Reset'' button on a
// server - it is issued immediately and will not allow processes to shut down prior to the reboot.
// Completing a hard reboot may initiate system disk checks upon server reboot, causing the boot up to
// take longer than normally expected. Remote management commands are unable to be executed if a reboot
// has been issued successfully within the last 20 minutes to avoid server failure. Remote management
// commands include: rebootSoft rebootHard powerOn powerOff powerCycle
func (softlayer_hardware *SoftLayer_Hardware) RebootHard(commonOptions *slapi.CommonOptions) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RebootSoft - The '''rebootSoft''' method reboots the server by issuing a reset command to the
// server's remote management card via soft reboot. When executing a soft reboot, servers allow all
// processes to shut down completely before rebooting. Remote management commands are unable to be
// issued within 20 minutes of issuing a successful soft reboot in order to avoid server failure.
// Remote management commands include: rebootSoft rebootHard powerOn powerOff powerCycle
func (softlayer_hardware *SoftLayer_Hardware) RebootSoft(commonOptions *slapi.CommonOptions) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RemoveAccessToNetworkStorage - This method is used to remove access to s SoftLayer_Network_Storage
// volumes that supports host- or network-level access control.
func (softlayer_hardware *SoftLayer_Hardware) RemoveAccessToNetworkStorage(commonOptions *slapi.CommonOptions, networkStorageTemplateObject SoftLayer_Network_Storage) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RemoveAccessToNetworkStorageList - This method is used to allow access to multiple
// SoftLayer_Network_Storage volumes that support host- or network-level access control.
func (softlayer_hardware *SoftLayer_Hardware) RemoveAccessToNetworkStorageList(commonOptions *slapi.CommonOptions, networkStorageTemplateObjects []SoftLayer_Network_Storage) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// SetTags - <nil>
func (softlayer_hardware *SoftLayer_Hardware) SetTags(commonOptions *slapi.CommonOptions, tags string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}
