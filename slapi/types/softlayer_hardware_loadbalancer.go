package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"

	time "time"
)

// SoftLayer_Hardware_LoadBalancer - <nil>
type SoftLayer_Hardware_LoadBalancer struct {

	// ServiceProviderResourceId - A hardware's internal identification number at its service provider
	ServiceProviderResourceId int `json:"serviceProviderResourceId,omitempty"`

	// BareMetalInstanceFlag - When true, this flag specifies that a hardware is Bare Metal Server. Bare
	// Metal Servers are physical bare metal servers that are billed with the same options as Virtual
	// Servers, with monthly and hourly rates. Bare Metal instances are ordered based on processor core
	// count and ram amount.
	BareMetalInstanceFlag int `json:"bareMetalInstanceFlag,omitempty"`

	// PostInstallScriptUri - URI of the script to be downloaded and executed after installation is
	// complete.
	PostInstallScriptUri string `json:"postInstallScriptUri,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// Domain - no documentation
	Domain string `json:"domain,omitempty"`

	// FullyQualifiedDomainName - A name reflecting the hostname and domain of the hardware. This is
	// created from the combined values of the hardware's hostname and domain name automatically, and thus
	// should not be edited directly.
	FullyQualifiedDomainName string `json:"fullyQualifiedDomainName,omitempty"`

	// SerialNumber - A hardware's serial number that is supplied by SoftLayer.
	SerialNumber string `json:"serialNumber,omitempty"`

	// Notes - A small note about a piece of hardware to use at your discretion.
	Notes string `json:"notes,omitempty"`

	// AccountId - A hardware's associated [[SoftLayer_Account|account]] id.
	AccountId int `json:"accountId,omitempty"`

	// Hostname - no documentation
	Hostname string `json:"hostname,omitempty"`

	// ManufacturerSerialNumber - A hardware's serial number that is supplied by the manufacturer.
	ManufacturerSerialNumber string `json:"manufacturerSerialNumber,omitempty"`

	// ProvisionDate - <nil>
	ProvisionDate *time.Time `json:"provisionDate,omitempty"`

	// HardwareStatusId - no documentation
	HardwareStatusId int `json:"hardwareStatusId,omitempty"`

	// ServiceProviderId - <nil>
	ServiceProviderId int `json:"serviceProviderId,omitempty"`
}

func (softlayer_hardware_loadbalancer *SoftLayer_Hardware_LoadBalancer) String() string {
	return "SoftLayer_Hardware_LoadBalancer"
}

// SoftLayer_Hardware_LoadBalancer_Extended is SoftLayer_Hardware_LoadBalancer with all maskable types.
type SoftLayer_Hardware_LoadBalancer_Extended struct {
	SoftLayer_Hardware_LoadBalancer

	// NetworkMonitorAttachedDownVirtualGuests - Virtual guests that are attached downstream to a hardware
	// that have failed monitoring
	NetworkMonitorAttachedDownVirtualGuests []*SoftLayer_Virtual_Guest `json:"networkMonitorAttachedDownVirtualGuests,omitempty"`

	// UserCount - A count of a list of users that have access to this hardware load balancer.
	UserCount uint64 `json:"userCount,omitempty"`

	// GlobalIdentifier - no documentation
	GlobalIdentifier string `json:"globalIdentifier,omitempty"`

	// ServiceProvider - Information regarding the piece of hardware's service provider.
	ServiceProvider *SoftLayer_Service_Provider `json:"serviceProvider,omitempty"`

	// ProcessorCount - A count of information regarding a piece of hardware's processors.
	ProcessorCount uint64 `json:"processorCount,omitempty"`

	// NetworkCardCount - A count of information regarding a piece of hardware's network cards.
	NetworkCardCount uint64 `json:"networkCardCount,omitempty"`

	// BillingItem - Information regarding the billing item for a server.
	BillingItem *SoftLayer_Billing_Item_Hardware `json:"billingItem,omitempty"`

	// PowerComponents - no documentation
	PowerComponents []*SoftLayer_Hardware_Power_Component `json:"powerComponents,omitempty"`

	// DownstreamHardwareBindings - no documentation
	DownstreamHardwareBindings []*SoftLayer_Network_Component_Uplink_Hardware `json:"downstreamHardwareBindings,omitempty"`

	// StorageNetworkComponentCount - no documentation
	StorageNetworkComponentCount uint64 `json:"storageNetworkComponentCount,omitempty"`

	// PrivateNetworkOnlyFlag - Whether the hardware only has access to the private network.
	PrivateNetworkOnlyFlag bool `json:"privateNetworkOnlyFlag,omitempty"`

	// SoftwareComponents - Information regarding a piece of hardware's installed software.
	SoftwareComponents []*SoftLayer_Software_Component `json:"softwareComponents,omitempty"`

	// DownstreamVirtualGuestCount - A count of information regarding all virtual guests attached to a
	// piece of network hardware.
	DownstreamVirtualGuestCount uint64 `json:"downstreamVirtualGuestCount,omitempty"`

	// ScaleAssets - Collection of scale assets this hardware corresponds to.
	ScaleAssets []*SoftLayer_Scale_Asset `json:"scaleAssets,omitempty"`

	// Datacenter - Information regarding the datacenter in which a piece of hardware resides.
	Datacenter *SoftLayer_Location `json:"datacenter,omitempty"`

	// DownstreamNetworkHardwareWithIncidents - All network hardware with monitoring warnings or errors
	// that are downstream from the selected piece of hardware.
	DownstreamNetworkHardwareWithIncidents []*SoftLayer_Hardware `json:"downstreamNetworkHardwareWithIncidents,omitempty"`

	// UpgradeRequest - An account's associated upgrade request object, if any.
	UpgradeRequest *SoftLayer_Product_Upgrade_Request `json:"upgradeRequest,omitempty"`

	// DownlinkNetworkHardware - All hardware that has uplink network connections to a piece of hardware.
	DownlinkNetworkHardware []*SoftLayer_Hardware `json:"downlinkNetworkHardware,omitempty"`

	// PrimaryNetworkComponent - Information regarding the hardware's primary public network component.
	PrimaryNetworkComponent *SoftLayer_Network_Component `json:"primaryNetworkComponent,omitempty"`

	// RouterCount - no documentation
	RouterCount uint64 `json:"routerCount,omitempty"`

	// Processors - Information regarding a piece of hardware's processors.
	Processors []*SoftLayer_Hardware_Component `json:"processors,omitempty"`

	// HostIpsSoftwareComponent - Information regarding a host IPS software component object.
	HostIpsSoftwareComponent *SoftLayer_Software_Component `json:"hostIpsSoftwareComponent,omitempty"`

	// NetworkMonitorIncidentCount - A count of the status of all of a piece of hardware's network
	// monitoring incidents.
	NetworkMonitorIncidentCount uint64 `json:"networkMonitorIncidentCount,omitempty"`

	// LockboxNetworkStorage - Information regarding a lockbox account associated with a server.
	LockboxNetworkStorage *SoftLayer_Network_Storage `json:"lockboxNetworkStorage,omitempty"`

	// BackendRouterCount - no documentation
	BackendRouterCount uint64 `json:"backendRouterCount,omitempty"`

	// HasTrustedPlatformModuleBillingItemFlag - no documentation
	HasTrustedPlatformModuleBillingItemFlag bool `json:"hasTrustedPlatformModuleBillingItemFlag,omitempty"`

	// PrimaryIpAddress - no documentation
	PrimaryIpAddress string `json:"primaryIpAddress,omitempty"`

	// NetworkManagementIpAddress - A piece of hardware's network management IP address.
	NetworkManagementIpAddress string `json:"networkManagementIpAddress,omitempty"`

	// VirtualizationPlatform - A piece of hardware's virtualization platform software.
	VirtualizationPlatform *SoftLayer_Software_Component `json:"virtualizationPlatform,omitempty"`

	// NetworkMonitorIncidents - The status of all of a piece of hardware's network monitoring incidents.
	NetworkMonitorIncidents []*SoftLayer_Network_Monitor_Version1_Incident `json:"networkMonitorIncidents,omitempty"`

	// Motherboard - Information regarding a piece of hardware's motherboard.
	Motherboard *SoftLayer_Hardware_Component `json:"motherboard,omitempty"`

	// RecentEvents - no documentation
	RecentEvents []*SoftLayer_Notification_Occurrence_Event `json:"recentEvents,omitempty"`

	// NetworkMonitorAttachedDownHardwareCount - A count of all servers with failed monitoring that are
	// attached downstream to a piece of hardware.
	NetworkMonitorAttachedDownHardwareCount uint64 `json:"networkMonitorAttachedDownHardwareCount,omitempty"`

	// LocationPathString - <nil>
	LocationPathString string `json:"locationPathString,omitempty"`

	// DownstreamNetworkHardwareWithIncidentCount - A count of all network hardware with monitoring
	// warnings or errors that are downstream from the selected piece of hardware.
	DownstreamNetworkHardwareWithIncidentCount uint64 `json:"downstreamNetworkHardwareWithIncidentCount,omitempty"`

	// AllowedNetworkStorageReplicas - The SoftLayer_Network_Storage objects whose Replica that this
	// SoftLayer_Hardware has access to.
	AllowedNetworkStorageReplicas []*SoftLayer_Network_Storage `json:"allowedNetworkStorageReplicas,omitempty"`

	// NetworkMonitors - Information regarding a piece of hardware's network monitors.
	NetworkMonitors []*SoftLayer_Network_Monitor_Version1_Query_Host `json:"networkMonitors,omitempty"`

	// SshKeys - SSH keys to be installed on the server during provisioning or an OS reload.
	SshKeys []*SoftLayer_Security_Ssh_Key `json:"sshKeys,omitempty"`

	// AllowedHost - The SoftLayer_Network_Storage_Allowed_Host information to connect this server to
	// Network Storage volumes that require access control lists.
	AllowedHost *SoftLayer_Network_Storage_Allowed_Host `json:"allowedHost,omitempty"`

	// BenchmarkCertificationCount - A count of information regarding a piece of hardware's benchmark
	// certifications.
	BenchmarkCertificationCount uint64 `json:"benchmarkCertificationCount,omitempty"`

	// DownstreamServerCount - A count of information regarding all servers attached downstream to a piece
	// of network hardware.
	DownstreamServerCount uint64 `json:"downstreamServerCount,omitempty"`

	// NetworkStatus - The value of a hardware's network status attribute.
	NetworkStatus string `json:"networkStatus,omitempty"`

	// PowerComponentCount - A count of the power components for a hardware object.
	PowerComponentCount uint64 `json:"powerComponentCount,omitempty"`

	// CurrentBillableBandwidthUsage - The current billable public outbound bandwidth for this hardware for
	// the current billing cycle.
	CurrentBillableBandwidthUsage slapi.Float64 `json:"currentBillableBandwidthUsage,omitempty"`

	// AllowedNetworkStorageCount - A count of the SoftLayer_Network_Storage objects that this
	// SoftLayer_Hardware has access to.
	AllowedNetworkStorageCount uint64 `json:"allowedNetworkStorageCount,omitempty"`

	// VirtualLicenses - Information regarding a piece of hardware's virtual software licenses.
	VirtualLicenses []*SoftLayer_Software_VirtualLicense `json:"virtualLicenses,omitempty"`

	// ResourceGroupRoleCount - no documentation
	ResourceGroupRoleCount uint64 `json:"resourceGroupRoleCount,omitempty"`

	// BillingItemFlag - no documentation
	BillingItemFlag bool `json:"billingItemFlag,omitempty"`

	// StorageNetworkComponents - <nil>
	StorageNetworkComponents []*SoftLayer_Network_Component `json:"storageNetworkComponents,omitempty"`

	// NetworkStorage - Information regarding a piece of hardware's associated network storage service
	// account.
	NetworkStorage []*SoftLayer_Network_Storage `json:"networkStorage,omitempty"`

	// SparePoolBillingItem - Information regarding the billing item for a spare pool server.
	SparePoolBillingItem *SoftLayer_Billing_Item_Hardware `json:"sparePoolBillingItem,omitempty"`

	// VirtualChassis - Information regarding the virtual chassis for a piece of hardware.
	VirtualChassis *SoftLayer_Hardware_Group `json:"virtualChassis,omitempty"`

	// Components - no documentation
	Components []*SoftLayer_Hardware_Component `json:"components,omitempty"`

	// OutboundBandwidthUsage - The sum of all the outbound network traffic data for the last 30 days.
	OutboundBandwidthUsage slapi.Float64 `json:"outboundBandwidthUsage,omitempty"`

	// NetworkComponentCount - A count of returns a hardware's network components.
	NetworkComponentCount uint64 `json:"networkComponentCount,omitempty"`

	// TagReferences - <nil>
	TagReferences []*SoftLayer_Tag_Reference `json:"tagReferences,omitempty"`

	// PowerSupplyCount - A count of information regarding a piece of hardware's power supply.
	PowerSupplyCount uint64 `json:"powerSupplyCount,omitempty"`

	// PointOfPresenceLocation - Information regarding the Point of Presence (PoP) location in which a
	// piece of hardware resides.
	PointOfPresenceLocation *SoftLayer_Location `json:"pointOfPresenceLocation,omitempty"`

	// OperatingSystem - Information regarding a piece of hardware's operating system.
	OperatingSystem *SoftLayer_Software_Component_OperatingSystem `json:"operatingSystem,omitempty"`

	// DownstreamVirtualGuests - Information regarding all virtual guests attached to a piece of network
	// hardware.
	DownstreamVirtualGuests []*SoftLayer_Virtual_Guest `json:"downstreamVirtualGuests,omitempty"`

	// RemoteManagementAccountCount - A count of user credentials to issue commands and/or interact with
	// the server's remote management card.
	RemoteManagementAccountCount uint64 `json:"remoteManagementAccountCount,omitempty"`

	// UplinkNetworkComponentCount - A count of information regarding the network component that is one
	// level higher than a piece of hardware on the network infrastructure.
	UplinkNetworkComponentCount uint64 `json:"uplinkNetworkComponentCount,omitempty"`

	// ModelFamily - <nil>
	ModelFamily string `json:"modelFamily,omitempty"`

	// SoftwareComponentCount - A count of information regarding a piece of hardware's installed software.
	SoftwareComponentCount uint64 `json:"softwareComponentCount,omitempty"`

	// FrontendRouters - no documentation
	FrontendRouters []*SoftLayer_Hardware `json:"frontendRouters,omitempty"`

	// ContinuousDataProtectionSoftwareComponent - A continuous data protection/server backup software
	// component object.
	ContinuousDataProtectionSoftwareComponent *SoftLayer_Software_Component `json:"continuousDataProtectionSoftwareComponent,omitempty"`

	// BenchmarkCertifications - Information regarding a piece of hardware's benchmark certifications.
	BenchmarkCertifications []*SoftLayer_Hardware_Benchmark_Certification `json:"benchmarkCertifications,omitempty"`

	// AllPowerComponentCount - no documentation
	AllPowerComponentCount uint64 `json:"allPowerComponentCount,omitempty"`

	// MonitoringRobot - Information regarding the hardware's monitoring robot.
	MonitoringRobot *SoftLayer_Monitoring_Robot `json:"monitoringRobot,omitempty"`

	// ResourceGroupRoles - <nil>
	ResourceGroupRoles []*SoftLayer_Resource_Group_Role `json:"resourceGroupRoles,omitempty"`

	// Routers - no documentation
	Routers []*SoftLayer_Hardware `json:"routers,omitempty"`

	// EvaultNetworkStorageCount - A count of information regarding a piece of hardware's associated EVault
	// network storage service account.
	EvaultNetworkStorageCount uint64 `json:"evaultNetworkStorageCount,omitempty"`

	// UplinkNetworkComponents - Information regarding the network component that is one level higher than
	// a piece of hardware on the network infrastructure.
	UplinkNetworkComponents []*SoftLayer_Network_Component `json:"uplinkNetworkComponents,omitempty"`

	// BandwidthAllocation - no documentation
	BandwidthAllocation slapi.Float64 `json:"bandwidthAllocation,omitempty"`

	// RaidControllers - The controllers contained within a piece of hardware.
	RaidControllers []*SoftLayer_Hardware_Component `json:"raidControllers,omitempty"`

	// SecurityScanRequests - Information regarding a piece of hardware's vulnerability scan requests.
	SecurityScanRequests []*SoftLayer_Network_Security_Scanner_Request `json:"securityScanRequests,omitempty"`

	// MetricTrackingObject - no documentation
	MetricTrackingObject *SoftLayer_Metric_Tracking_Object_HardwareServer `json:"metricTrackingObject,omitempty"`

	// NetworkGatewayMemberFlag - Whether or not this device is part of a network gateway.
	NetworkGatewayMemberFlag bool `json:"networkGatewayMemberFlag,omitempty"`

	// HardDriveCount - A count of the hard drives contained within a piece of hardware.
	HardDriveCount uint64 `json:"hardDriveCount,omitempty"`

	// AllowedNetworkStorage - The SoftLayer_Network_Storage objects that this SoftLayer_Hardware has
	// access to.
	AllowedNetworkStorage []*SoftLayer_Network_Storage `json:"allowedNetworkStorage,omitempty"`

	// ManagedResourceFlag - A flag indicating that the hardware is a managed resource.
	ManagedResourceFlag bool `json:"managedResourceFlag,omitempty"`

	// NotesHistory - <nil>
	NotesHistory []*SoftLayer_Hardware_Note `json:"notesHistory,omitempty"`

	// NetworkStatusAttribute - no documentation
	NetworkStatusAttribute *SoftLayer_Hardware_Attribute `json:"networkStatusAttribute,omitempty"`

	// HourlyBillingFlag - no documentation
	HourlyBillingFlag bool `json:"hourlyBillingFlag,omitempty"`

	// AverageDailyPublicBandwidthUsage - The average daily public bandwidth usage for the current billing
	// cycle.
	AverageDailyPublicBandwidthUsage slapi.Float64 `json:"averageDailyPublicBandwidthUsage,omitempty"`

	// DownlinkNetworkHardwareCount - A count of all hardware that has uplink network connections to a
	// piece of hardware.
	DownlinkNetworkHardwareCount uint64 `json:"downlinkNetworkHardwareCount,omitempty"`

	// DriveControllerCount - A count of the drive controllers contained within a piece of hardware.
	DriveControllerCount uint64 `json:"driveControllerCount,omitempty"`

	// DownstreamNetworkHardwareCount - A count of all network hardware downstream from the selected piece
	// of hardware.
	DownstreamNetworkHardwareCount uint64 `json:"downstreamNetworkHardwareCount,omitempty"`

	// AllowedNetworkStorageReplicaCount - A count of the SoftLayer_Network_Storage objects whose Replica
	// that this SoftLayer_Hardware has access to.
	AllowedNetworkStorageReplicaCount uint64 `json:"allowedNetworkStorageReplicaCount,omitempty"`

	// DownstreamNetworkHardware - All network hardware downstream from the selected piece of hardware.
	DownstreamNetworkHardware []*SoftLayer_Hardware `json:"downstreamNetworkHardware,omitempty"`

	// MonitoringAgents - Information regarding the monitoring agents associated with a piece of hardware.
	MonitoringAgents []*SoftLayer_Monitoring_Agent `json:"monitoringAgents,omitempty"`

	// VirtualRackName - The name of the bandwidth allotment belonging to a piece of hardware.
	VirtualRackName string `json:"virtualRackName,omitempty"`

	// DownstreamServers - Information regarding all servers attached downstream to a piece of network
	// hardware.
	DownstreamServers []*SoftLayer_Hardware `json:"downstreamServers,omitempty"`

	// Memory - Information regarding a piece of hardware's memory.
	Memory []*SoftLayer_Hardware_Component `json:"memory,omitempty"`

	// FrontendRouterCount - no documentation
	FrontendRouterCount uint64 `json:"frontendRouterCount,omitempty"`

	// DownlinkVirtualGuests - Information regarding all virtual guests attached to a piece of network
	// hardware.
	DownlinkVirtualGuests []*SoftLayer_Virtual_Guest `json:"downlinkVirtualGuests,omitempty"`

	// VirtualLicenseCount - A count of information regarding a piece of hardware's virtual software
	// licenses.
	VirtualLicenseCount uint64 `json:"virtualLicenseCount,omitempty"`

	// RemoteManagementComponent - A hardware's associated remote management component. This is normally
	RemoteManagementComponent *SoftLayer_Network_Component `json:"remoteManagementComponent,omitempty"`

	// DownlinkServers - Information regarding all servers attached to a piece of network hardware.
	DownlinkServers []*SoftLayer_Hardware `json:"downlinkServers,omitempty"`

	// MemoryCount - A count of information regarding a piece of hardware's memory.
	MemoryCount uint64 `json:"memoryCount,omitempty"`

	// ResourceGroups - The resource groups in which this hardware is a member.
	ResourceGroups []*SoftLayer_Resource_Group `json:"resourceGroups,omitempty"`

	// Attributes - Information regarding a piece of hardware's specific attributes.
	Attributes []*SoftLayer_Hardware_Attribute `json:"attributes,omitempty"`

	// ResourceGroupMemberReferenceCount - no documentation
	ResourceGroupMemberReferenceCount uint64 `json:"resourceGroupMemberReferenceCount,omitempty"`

	// NetworkMonitorAttachedDownHardware - All servers with failed monitoring that are attached downstream
	// to a piece of hardware.
	NetworkMonitorAttachedDownHardware []*SoftLayer_Hardware `json:"networkMonitorAttachedDownHardware,omitempty"`

	// EvaultNetworkStorage - Information regarding a piece of hardware's associated EVault network storage
	// service account.
	EvaultNetworkStorage []*SoftLayer_Network_Storage `json:"evaultNetworkStorage,omitempty"`

	// FrontendNetworkComponents - A piece of hardware's front-end or public network components.
	FrontendNetworkComponents []*SoftLayer_Network_Component `json:"frontendNetworkComponents,omitempty"`

	// ActiveComponents - no documentation
	ActiveComponents []*SoftLayer_Hardware_Component `json:"activeComponents,omitempty"`

	// ResourceGroupCount - A count of the resource groups in which this hardware is a member.
	ResourceGroupCount uint64 `json:"resourceGroupCount,omitempty"`

	// Rack - <nil>
	Rack *SoftLayer_Location `json:"rack,omitempty"`

	// NetworkMonitorAttachedDownVirtualGuestCount - A count of virtual guests that are attached downstream
	// to a hardware that have failed monitoring
	NetworkMonitorAttachedDownVirtualGuestCount uint64 `json:"networkMonitorAttachedDownVirtualGuestCount,omitempty"`

	// NetworkCards - Information regarding a piece of hardware's network cards.
	NetworkCards []*SoftLayer_Hardware_Component `json:"networkCards,omitempty"`

	// OperatingSystemReferenceCode - A hardware's operating system software description.
	OperatingSystemReferenceCode string `json:"operatingSystemReferenceCode,omitempty"`

	// NotesHistoryCount - no documentation
	NotesHistoryCount uint64 `json:"notesHistoryCount,omitempty"`

	// DownlinkVirtualGuestCount - A count of information regarding all virtual guests attached to a piece
	// of network hardware.
	DownlinkVirtualGuestCount uint64 `json:"downlinkVirtualGuestCount,omitempty"`

	// RemoteManagementAccounts - User credentials to issue commands and/or interact with the server's
	// remote management card.
	RemoteManagementAccounts []*SoftLayer_Hardware_Component_RemoteManagement_User `json:"remoteManagementAccounts,omitempty"`

	// BackendRouters - no documentation
	BackendRouters []*SoftLayer_Hardware `json:"backendRouters,omitempty"`

	// FirewallServiceComponent - Information regarding a piece of hardware's firewall services.
	FirewallServiceComponent *SoftLayer_Network_Component_Firewall `json:"firewallServiceComponent,omitempty"`

	// TagReferenceCount - no documentation
	TagReferenceCount uint64 `json:"tagReferenceCount,omitempty"`

	// SecurityScanRequestCount - A count of information regarding a piece of hardware's vulnerability scan
	// requests.
	SecurityScanRequestCount uint64 `json:"securityScanRequestCount,omitempty"`

	// TopLevelLocation - <nil>
	TopLevelLocation *SoftLayer_Location `json:"topLevelLocation,omitempty"`

	// NetworkComponents - no documentation
	NetworkComponents []*SoftLayer_Network_Component `json:"networkComponents,omitempty"`

	// NetworkVlanCount - A count of the network virtual LANs (VLANs) associated with a piece of hardware's
	// network components.
	NetworkVlanCount uint64 `json:"networkVlanCount,omitempty"`

	// OutboundPublicBandwidthUsage - The total public outbound bandwidth for this hardware for the current
	// billing cycle.
	OutboundPublicBandwidthUsage slapi.Float64 `json:"outboundPublicBandwidthUsage,omitempty"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account,omitempty"`

	// SshKeyCount - A count of sSH keys to be installed on the server during provisioning or an OS reload.
	SshKeyCount uint64 `json:"sshKeyCount,omitempty"`

	// FixedConfigurationPreset - Defines the fixed components in a fixed configuration bare metal server.
	FixedConfigurationPreset *SoftLayer_Product_Package_Preset `json:"fixedConfigurationPreset,omitempty"`

	// ResourceGroupMemberReferences - <nil>
	ResourceGroupMemberReferences []*SoftLayer_Resource_Group_Member `json:"resourceGroupMemberReferences,omitempty"`

	// ScaleAssetCount - A count of collection of scale assets this hardware corresponds to.
	ScaleAssetCount uint64 `json:"scaleAssetCount,omitempty"`

	// PrimaryBackendIpAddress - no documentation
	PrimaryBackendIpAddress string `json:"primaryBackendIpAddress,omitempty"`

	// AntivirusSpywareSoftwareComponent - Information regarding an antivirus/spyware software component
	// object.
	AntivirusSpywareSoftwareComponent *SoftLayer_Software_Component `json:"antivirusSpywareSoftwareComponent,omitempty"`

	// RaidControllerCount - A count of the controllers contained within a piece of hardware.
	RaidControllerCount uint64 `json:"raidControllerCount,omitempty"`

	// VirtualHost - no documentation
	VirtualHost *SoftLayer_Virtual_Host `json:"virtualHost,omitempty"`

	// PowerSupply - Information regarding a piece of hardware's power supply.
	PowerSupply []*SoftLayer_Hardware_Component `json:"powerSupply,omitempty"`

	// Users - A list of users that have access to this hardware load balancer.
	Users []*SoftLayer_User_Customer `json:"users,omitempty"`

	// DriveControllers - The drive controllers contained within a piece of hardware.
	DriveControllers []*SoftLayer_Hardware_Component `json:"driveControllers,omitempty"`

	// BusinessContinuanceInsuranceFlag - Status indicating whether or not a piece of hardware has business
	// continuance insurance.
	BusinessContinuanceInsuranceFlag bool `json:"businessContinuanceInsuranceFlag,omitempty"`

	// ServerRoom - Information regarding the server room in which the hardware is located.
	ServerRoom *SoftLayer_Location `json:"serverRoom,omitempty"`

	// DownlinkHardware - All hardware that has uplink network connections to a piece of hardware.
	DownlinkHardware []*SoftLayer_Hardware `json:"downlinkHardware,omitempty"`

	// MonitoringAgentCount - A count of information regarding the monitoring agents associated with a
	// piece of hardware.
	MonitoringAgentCount uint64 `json:"monitoringAgentCount,omitempty"`

	// Location - Where a piece of hardware is located within SoftLayer's location hierarchy.
	Location *SoftLayer_Location `json:"location,omitempty"`

	// NetworkGatewayMember - The gateway member if this device is part of a network gateway.
	NetworkGatewayMember *SoftLayer_Network_Gateway_Member `json:"networkGatewayMember,omitempty"`

	// AllPowerComponents - <nil>
	AllPowerComponents []*SoftLayer_Hardware_Power_Component `json:"allPowerComponents,omitempty"`

	// ProcessorCoreAmount - The total number of processor cores, summed from all processors that are
	// attached to a piece of hardware
	ProcessorCoreAmount uint `json:"processorCoreAmount,omitempty"`

	// LastTransaction - Information regarding the last transaction a server performed.
	LastTransaction *SoftLayer_Provisioning_Version1_Transaction `json:"lastTransaction,omitempty"`

	// MonitoringServiceComponent - Information regarding a piece of hardware's network monitoring
	// services.
	MonitoringServiceComponent *SoftLayer_Network_Monitor_Version1_Query_Host_Stratum `json:"monitoringServiceComponent,omitempty"`

	// MonitoringServiceEligibilityFlag - The monitoring service flag eligibility status for a piece of
	// hardware.
	MonitoringServiceEligibilityFlag bool `json:"monitoringServiceEligibilityFlag,omitempty"`

	// NextBillingCycleBandwidthAllocation - A hardware's allotted bandwidth for the next billing cycle
	// (measured in
	NextBillingCycleBandwidthAllocation slapi.Float64 `json:"nextBillingCycleBandwidthAllocation,omitempty"`

	// NetworkVlans - The network virtual LANs (VLANs) associated with a piece of hardware's network
	// components.
	NetworkVlans []*SoftLayer_Network_Vlan `json:"networkVlans,omitempty"`

	// AttributeCount - A count of information regarding a piece of hardware's specific attributes.
	AttributeCount uint64 `json:"attributeCount,omitempty"`

	// RecentEventCount - A count of recent events that impact this hardware.
	RecentEventCount uint64 `json:"recentEventCount,omitempty"`

	// ActiveComponentCount - A count of a piece of hardware's active physical components.
	ActiveComponentCount uint64 `json:"activeComponentCount,omitempty"`

	// PrimaryBackendNetworkComponent - Information regarding the hardware's primary back-end network
	// component.
	PrimaryBackendNetworkComponent *SoftLayer_Network_Component `json:"primaryBackendNetworkComponent,omitempty"`

	// VirtualChassisSiblingCount - A count of information regarding the virtual chassis siblings for a
	// piece of hardware.
	VirtualChassisSiblingCount uint64 `json:"virtualChassisSiblingCount,omitempty"`

	// ActiveNetworkMonitorIncidentCount - A count of a piece of hardware's active network monitoring
	// incidents.
	ActiveNetworkMonitorIncidentCount uint64 `json:"activeNetworkMonitorIncidentCount,omitempty"`

	// VirtualRackId - The name of the bandwidth allotment belonging to a piece of hardware.
	VirtualRackId int `json:"virtualRackId,omitempty"`

	// UplinkHardware - The network device connected to a piece of hardware.
	UplinkHardware *SoftLayer_Hardware `json:"uplinkHardware,omitempty"`

	// UserData - A string containing custom user data for a hardware order.
	UserData []*SoftLayer_Hardware_Attribute `json:"userData,omitempty"`

	// HardwareStatus - no documentation
	HardwareStatus *SoftLayer_Hardware_Status `json:"hardwareStatus,omitempty"`

	// HardwareFunction - no documentation
	HardwareFunction *SoftLayer_Hardware_Function `json:"hardwareFunction,omitempty"`

	// BackendNetworkComponentCount - A count of a piece of hardware's back-end or private network
	// components.
	BackendNetworkComponentCount uint64 `json:"backendNetworkComponentCount,omitempty"`

	// MemoryCapacity - The amount of memory a piece of hardware has, measured in gigabytes.
	MemoryCapacity uint `json:"memoryCapacity,omitempty"`

	// InboundPublicBandwidthUsage - The total public inbound bandwidth for this hardware for the current
	// billing cycle.
	InboundPublicBandwidthUsage slapi.Float64 `json:"inboundPublicBandwidthUsage,omitempty"`

	// ProcessorPhysicalCoreAmount - The total number of physical processor cores, summed from all
	// processors that are attached to a piece of hardware
	ProcessorPhysicalCoreAmount uint `json:"processorPhysicalCoreAmount,omitempty"`

	// VirtualChassisSiblings - Information regarding the virtual chassis siblings for a piece of hardware.
	VirtualChassisSiblings []*SoftLayer_Hardware `json:"virtualChassisSiblings,omitempty"`

	// FrontendNetworkComponentCount - A count of a piece of hardware's front-end or public network
	// components.
	FrontendNetworkComponentCount uint64 `json:"frontendNetworkComponentCount,omitempty"`

	// VirtualRack - Information regarding the bandwidth allotment to which a piece of hardware belongs.
	VirtualRack *SoftLayer_Network_Bandwidth_Version1_Allotment `json:"virtualRack,omitempty"`

	// LatestNetworkMonitorIncident - A piece of hardware's latest network monitoring incident.
	LatestNetworkMonitorIncident *SoftLayer_Network_Monitor_Version1_Incident `json:"latestNetworkMonitorIncident,omitempty"`

	// NetworkMonitorCount - A count of information regarding a piece of hardware's network monitors.
	NetworkMonitorCount uint64 `json:"networkMonitorCount,omitempty"`

	// DownlinkServerCount - A count of information regarding all servers attached to a piece of network
	// hardware.
	DownlinkServerCount uint64 `json:"downlinkServerCount,omitempty"`

	// BackendNetworkComponents - A piece of hardware's back-end or private network components.
	BackendNetworkComponents []*SoftLayer_Network_Component `json:"backendNetworkComponents,omitempty"`

	// ComponentCount - no documentation
	ComponentCount uint64 `json:"componentCount,omitempty"`

	// HardwareChassis - no documentation
	HardwareChassis *SoftLayer_Hardware_Chassis `json:"hardwareChassis,omitempty"`

	// HardDrives - The hard drives contained within a piece of hardware.
	HardDrives []*SoftLayer_Hardware_Component `json:"hardDrives,omitempty"`

	// DatacenterName - The name of the datacenter in which a piece of hardware resides.
	DatacenterName string `json:"datacenterName,omitempty"`

	// InboundBandwidthUsage - The sum of all the inbound network traffic data for the last 30 days.
	InboundBandwidthUsage slapi.Float64 `json:"inboundBandwidthUsage,omitempty"`

	// MonitoringServiceFlag - no documentation
	MonitoringServiceFlag bool `json:"monitoringServiceFlag,omitempty"`

	// DownstreamHardwareBindingCount - A count of all hardware downstream from a network device.
	DownstreamHardwareBindingCount uint64 `json:"downstreamHardwareBindingCount,omitempty"`

	// BlockCancelBecauseDisconnectedFlag - Determines whether the hardware is ineligible for cancellation
	// because it is disconnected.
	BlockCancelBecauseDisconnectedFlag bool `json:"blockCancelBecauseDisconnectedFlag,omitempty"`

	// BandwidthAllotmentDetail - A hardware's allotted detail record. Allotment details link bandwidth
	// allocation with allotments.
	BandwidthAllotmentDetail *SoftLayer_Network_Bandwidth_Version1_Allotment_Detail `json:"bandwidthAllotmentDetail,omitempty"`

	// ActiveNetworkMonitorIncident - A piece of hardware's active network monitoring incidents.
	ActiveNetworkMonitorIncident []*SoftLayer_Network_Monitor_Version1_Incident `json:"activeNetworkMonitorIncident,omitempty"`

	// UserDataCount - A count of a string containing custom user data for a hardware order.
	UserDataCount uint64 `json:"userDataCount,omitempty"`

	// DownlinkHardwareCount - A count of all hardware that has uplink network connections to a piece of
	// hardware.
	DownlinkHardwareCount uint64 `json:"downlinkHardwareCount,omitempty"`

	// NetworkStorageCount - A count of information regarding a piece of hardware's associated network
	// storage service account.
	NetworkStorageCount uint64 `json:"networkStorageCount,omitempty"`

	// HardwareFunctionDescription - no documentation
	HardwareFunctionDescription string `json:"hardwareFunctionDescription,omitempty"`
}

func (softlayer_hardware_loadbalancer *SoftLayer_Hardware_LoadBalancer_Extended) String() string {
	return "SoftLayer_Hardware_LoadBalancer"
}
