package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Hardware_Server - The SoftLayer_Hardware_Server data type contains general information
// relating to a single SoftLayer server.
type SoftLayer_Hardware_Server struct {

	// ActiveNetworkFirewallBillingItem - The billing item for a server's attached network firewall.
	ActiveNetworkFirewallBillingItem *SoftLayer_Billing_Item `json:"activeNetworkFirewallBillingItem"`

	// ActiveTicketCount - no documentation
	ActiveTicketCount uint64 `json:"activeTicketCount"`

	// ActiveTickets - <nil>
	ActiveTickets []*SoftLayer_Ticket `json:"activeTickets"`

	// ActiveTransaction - no documentation
	ActiveTransaction *SoftLayer_Provisioning_Version1_Transaction `json:"activeTransaction"`

	// ActiveTransactionCount - A count of any active transaction(s) that are currently running for the
	// server (example: os reload).
	ActiveTransactionCount uint64 `json:"activeTransactionCount"`

	// ActiveTransactions - Any active transaction(s) that are currently running for the server (example:
	// os reload).
	ActiveTransactions []*SoftLayer_Provisioning_Version1_Transaction `json:"activeTransactions"`

	// AvailableMonitoring - An object that stores the maximum level for the monitoring query types and
	// response types.
	AvailableMonitoring []*SoftLayer_Network_Monitor_Version1_Query_Host_Stratum `json:"availableMonitoring"`

	// AvailableMonitoringCount - A count of an object that stores the maximum level for the monitoring
	// query types and response types.
	AvailableMonitoringCount uint64 `json:"availableMonitoringCount"`

	// AverageDailyBandwidthUsage - The average daily total bandwidth usage for the current billing cycle.
	AverageDailyBandwidthUsage float32 `json:"averageDailyBandwidthUsage"`

	// AverageDailyPrivateBandwidthUsage - The average daily private bandwidth usage for the current
	// billing cycle.
	AverageDailyPrivateBandwidthUsage float32 `json:"averageDailyPrivateBandwidthUsage"`

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

	// ContainsSolidStateDrivesFlag - <nil>
	ContainsSolidStateDrivesFlag bool `json:"containsSolidStateDrivesFlag"`

	// ControlPanel - no documentation
	ControlPanel *SoftLayer_Software_Component_ControlPanel `json:"controlPanel"`

	// Cost - The total cost of a server, measured in US Dollars
	Cost float32 `json:"cost"`

	// CurrentBandwidthSummary - An object that provides commonly used bandwidth summary components for the
	// current billing cycle.
	CurrentBandwidthSummary *SoftLayer_Metric_Tracking_Object_Bandwidth_Summary `json:"currentBandwidthSummary"`

	// CustomerInstalledOperatingSystemFlag - no documentation
	CustomerInstalledOperatingSystemFlag bool `json:"customerInstalledOperatingSystemFlag"`

	// CustomerOwnedFlag - no documentation
	CustomerOwnedFlag bool `json:"customerOwnedFlag"`

	// InboundPrivateBandwidthUsage - The total private inbound bandwidth for this hardware for the current
	// billing cycle.
	InboundPrivateBandwidthUsage float64 `json:"inboundPrivateBandwidthUsage"`

	// LastOperatingSystemReload - The last transaction that a server's operating system was loaded.
	LastOperatingSystemReload *SoftLayer_Provisioning_Version1_Transaction `json:"lastOperatingSystemReload"`

	// MetricTrackingObjectId - no documentation
	MetricTrackingObjectId int `json:"metricTrackingObjectId"`

	// MonitoringUserNotification - The monitoring notification objects for this hardware. Each object
	// links this hardware instance to a user account that will be notified if monitoring on this hardware
	// object fails
	MonitoringUserNotification []*SoftLayer_User_Customer_Notification_Hardware `json:"monitoringUserNotification"`

	// MonitoringUserNotificationCount - A count of the monitoring notification objects for this hardware.
	// Each object links this hardware instance to a user account that will be notified if monitoring on
	// this hardware object fails
	MonitoringUserNotificationCount uint64 `json:"monitoringUserNotificationCount"`

	// OpenCancellationTicket - An open ticket requesting cancellation of this server, if one exists.
	OpenCancellationTicket *SoftLayer_Ticket `json:"openCancellationTicket"`

	// OutboundPrivateBandwidthUsage - The total private outbound bandwidth for this hardware for the
	// current billing cycle.
	OutboundPrivateBandwidthUsage float64 `json:"outboundPrivateBandwidthUsage"`

	// OverBandwidthAllocationFlag - Whether the bandwidth usage for this hardware for the current billing
	// cycle exceeds the allocation.
	OverBandwidthAllocationFlag int `json:"overBandwidthAllocationFlag"`

	// PrivateIpAddress - no documentation
	PrivateIpAddress string `json:"privateIpAddress"`

	// ProjectedOverBandwidthAllocationFlag - Whether the bandwidth usage for this hardware for the current
	// billing cycle is projected to exceed the allocation.
	ProjectedOverBandwidthAllocationFlag int `json:"projectedOverBandwidthAllocationFlag"`

	// ProjectedPublicBandwidthUsage - The projected public outbound bandwidth for this hardware for the
	// current billing cycle.
	ProjectedPublicBandwidthUsage float32 `json:"projectedPublicBandwidthUsage"`

	// RecentRemoteManagementCommandCount - A count of the last five commands issued to the server's remote
	// management card.
	RecentRemoteManagementCommandCount uint64 `json:"recentRemoteManagementCommandCount"`

	// RecentRemoteManagementCommands - The last five commands issued to the server's remote management
	// card.
	RecentRemoteManagementCommands []*SoftLayer_Hardware_Component_RemoteManagement_Command_Request `json:"recentRemoteManagementCommands"`

	// RegionalInternetRegistry - <nil>
	RegionalInternetRegistry *SoftLayer_Network_Regional_Internet_Registry `json:"regionalInternetRegistry"`

	// RemoteManagement - no documentation
	RemoteManagement *SoftLayer_Hardware_Component_RemoteManagement `json:"remoteManagement"`

	// RemoteManagementUserCount - A count of user(s) who have access to issue commands and/or interact
	// with the server's remote management card.
	RemoteManagementUserCount uint64 `json:"remoteManagementUserCount"`

	// RemoteManagementUsers - User(s) who have access to issue commands and/or interact with the server's
	// remote management card.
	RemoteManagementUsers []*SoftLayer_Hardware_Component_RemoteManagement_User `json:"remoteManagementUsers"`

	// StatisticsRemoteManagement - A server's remote management card used for statistics.
	StatisticsRemoteManagement *SoftLayer_Hardware_Component_RemoteManagement `json:"statisticsRemoteManagement"`

	// UserCount - A count of a list of users that have access to this computing instance.
	UserCount uint64 `json:"userCount"`

	// Users - A list of users that have access to this computing instance.
	Users []*SoftLayer_User_Customer `json:"users"`

	// VirtualGuestCount - no documentation
	VirtualGuestCount uint64 `json:"virtualGuestCount"`

	// VirtualGuests - no documentation
	VirtualGuests []*SoftLayer_Virtual_Guest `json:"virtualGuests"`
}

func (softlayer_hardware_server *SoftLayer_Hardware_Server) String() string {
	return "SoftLayer_Hardware_Server"
}

// ActivatePrivatePort - no documentation
func (softlayer_hardware_server *SoftLayer_Hardware_Server) ActivatePrivatePort(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// ActivatePublicPort - no documentation
func (softlayer_hardware_server *SoftLayer_Hardware_Server) ActivatePublicPort(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// BootToRescueLayer - The Rescue Kernel is designed to provide you with the ability to bring a server
// online in order to troubleshoot system problems that would normally only be resolved by an OS
// Reload. The correct Rescue Kernel will be selected based upon the currently installed operating
// system. When the rescue kernel process is initiated, the server will shutdown and reboot on to the
// public network with the same IP's assigned to the server to allow for remote connections. It will
// bring your server offline for approximately 10 minutes while the rescue is in progress. The
// root/administrator password will be the same as what is listed in the portal for the server.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) BootToRescueLayer(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// CreateFirmwareUpdateTransaction - You can launch firmware updates by selecting from your server
// list. It will bring your server offline for approximately 20 minutes while the updates are in
// progress. In the event of a hardware failure during this test our datacenter engineers will be
// notified of the problem automatically. They will then replace any failed components to bring your
// server back online, and will be contacting you to ensure that impact on your server is minimal.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) CreateFirmwareUpdateTransaction(ctx *slapi.RequestContext, ipmi int, raidController int, bios int, harddrive int) (bool, error) {
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
func (softlayer_hardware_server *SoftLayer_Hardware_Server) CreateObject(ctx *slapi.RequestContext, templateObject SoftLayer_Hardware_Server) (*SoftLayer_Hardware_Server, error) {
	var returnValue *SoftLayer_Hardware_Server
	return returnValue, nil
}

// CreatePostSoftwareInstallTransaction - <nil>
func (softlayer_hardware_server *SoftLayer_Hardware_Server) CreatePostSoftwareInstallTransaction(ctx *slapi.RequestContext, installCodes []string, returnBoolean bool) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObject - no documentation
func (softlayer_hardware_server *SoftLayer_Hardware_Server) EditObject(ctx *slapi.RequestContext, templateObject SoftLayer_Hardware_Server) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetBackendBandwidthUsage - Use this method to return an array of private bandwidth utilization
// records between a given date range. This method represents the NEW version of
// getFrontendBandwidthUse
func (softlayer_hardware_server *SoftLayer_Hardware_Server) GetBackendBandwidthUsage(ctx *slapi.RequestContext, startDate time.Time, endDate time.Time) ([]*SoftLayer_Metric_Tracking_Object_Data, error) {
	var returnValue []*SoftLayer_Metric_Tracking_Object_Data
	return returnValue, nil
}

// GetBackendBandwidthUse - Use this method to return an array of private bandwidth utilization records
// between a given date range.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) GetBackendBandwidthUse(ctx *slapi.RequestContext, startDate time.Time, endDate time.Time) ([]*SoftLayer_Network_Bandwidth_Version1_Usage_Detail, error) {
	var returnValue []*SoftLayer_Network_Bandwidth_Version1_Usage_Detail
	return returnValue, nil
}

// GetBandwidthForDateRange - Retrieve a collection of bandwidth data from an individual public or
// private network tracking object. Data is ideal if you with to employ your own traffic storage and
// graphing systems.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) GetBandwidthForDateRange(ctx *slapi.RequestContext, startDate time.Time, endDate time.Time) ([]*SoftLayer_Metric_Tracking_Object_Data, error) {
	var returnValue []*SoftLayer_Metric_Tracking_Object_Data
	return returnValue, nil
}

// GetBandwidthImage - Use this method when needing a bandwidth image for a single server. It will
// gather the correct input parameters for the generic graphing utility automatically based on the
// snapshot specified. Use the $draw flag to suppress the generation of the actual binary PNG image.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) GetBandwidthImage(ctx *slapi.RequestContext, networkType string, snapshotRange string, draw bool, dateSpecified time.Time) (*SoftLayer_Container_Bandwidth_GraphOutputs, error) {
	var returnValue *SoftLayer_Container_Bandwidth_GraphOutputs
	return returnValue, nil
}

// GetCurrentBenchmarkCertificationResultFile - Attempt to retrieve the file associated with the
// current benchmark certification result, if such a file exists. If there is no file for this
// benchmark certification result, calling this method throws an exception.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) GetCurrentBenchmarkCertificationResultFile(ctx *slapi.RequestContext) (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetCustomBandwidthDataByDate - no documentation
func (softlayer_hardware_server *SoftLayer_Hardware_Server) GetCustomBandwidthDataByDate(ctx *slapi.RequestContext, graphData SoftLayer_Container_Graph) (*SoftLayer_Container_Graph, error) {
	var returnValue *SoftLayer_Container_Graph
	return returnValue, nil
}

// GetFirewallProtectableSubnets - Get the subnets associated with this server that are protectable by
// a network component firewall.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) GetFirewallProtectableSubnets(ctx *slapi.RequestContext) ([]*SoftLayer_Network_Subnet, error) {
	var returnValue []*SoftLayer_Network_Subnet
	return returnValue, nil
}

// GetFrontendBandwidthUsage - Use this method to return an array of public bandwidth utilization
// records between a given date range. This method represents the NEW version of
// getFrontendBandwidthUse
func (softlayer_hardware_server *SoftLayer_Hardware_Server) GetFrontendBandwidthUsage(ctx *slapi.RequestContext, startDate time.Time, endDate time.Time) ([]*SoftLayer_Metric_Tracking_Object_Data, error) {
	var returnValue []*SoftLayer_Metric_Tracking_Object_Data
	return returnValue, nil
}

// GetFrontendBandwidthUse - Use this method to return an array of public bandwidth utilization records
// between a given date range.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) GetFrontendBandwidthUse(ctx *slapi.RequestContext, startDate time.Time, endDate time.Time) ([]*SoftLayer_Network_Bandwidth_Version1_Usage_Detail, error) {
	var returnValue []*SoftLayer_Network_Bandwidth_Version1_Usage_Detail
	return returnValue, nil
}

// GetHardwareByIpAddress - Retrieve a server by searching for the primary IP address.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) GetHardwareByIpAddress(ctx *slapi.RequestContext, ipAddress string) (*SoftLayer_Hardware_Server, error) {
	var returnValue *SoftLayer_Hardware_Server
	return returnValue, nil
}

// GetItemPricesFromSoftwareDescriptions - Return a collection of SoftLayer_Item_Price objects from a
// collection of SoftLayer_Software_Description
func (softlayer_hardware_server *SoftLayer_Hardware_Server) GetItemPricesFromSoftwareDescriptions(ctx *slapi.RequestContext, softwareDescriptions []SoftLayer_Software_Description, includeTranslationsFlag bool, returnAllPricesFlag bool) ([]*SoftLayer_Product_Item, error) {
	var returnValue []*SoftLayer_Product_Item
	return returnValue, nil
}

// GetManagementNetworkComponent - Retrieve the remote management network component attached with this
// server.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) GetManagementNetworkComponent(ctx *slapi.RequestContext) (*SoftLayer_Network_Component, error) {
	var returnValue *SoftLayer_Network_Component
	return returnValue, nil
}

// GetNetworkComponentFirewallProtectableIpAddresses - Get the IP addresses associated with this server
// that are protectable by a network component firewall. Note, this may not return all values for IPv6
// subnets for this server. Please use getFirewallProtectableSubnets to get all protectable subnets.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) GetNetworkComponentFirewallProtectableIpAddresses(ctx *slapi.RequestContext) ([]*SoftLayer_Network_Subnet_IpAddress, error) {
	var returnValue []*SoftLayer_Network_Subnet_IpAddress
	return returnValue, nil
}

// GetObject - getObject retrieves the SoftLayer_Hardware_Server object whose ID number corresponds to
// the ID number of the init parameter passed to the SoftLayer_Hardware service. You can only retrieve
// servers from the account that your portal user is assigned to.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Hardware_Server, error) {
	var returnValue *SoftLayer_Hardware_Server
	return returnValue, nil
}

// GetPMInfo - Retrieve a server's hardware state via its internal sensors. Remote sensor data is
// transmitted to the SoftLayer API by way of the server's remote management card. Sensor data measures
// system temperatures, voltages, and other local server settings. Sensor data is cached for 30
// seconds. Calls made to getSensorData for the same server within 30 seconds of each other will return
// the same data. Subsequent calls will return new data once the cache expires.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) GetPMInfo(ctx *slapi.RequestContext) ([]*SoftLayer_Container_RemoteManagement_PmInfo, error) {
	var returnValue []*SoftLayer_Container_RemoteManagement_PmInfo
	return returnValue, nil
}

// GetPrimaryDriveSize - <nil>
func (softlayer_hardware_server *SoftLayer_Hardware_Server) GetPrimaryDriveSize(ctx *slapi.RequestContext) (int, error) {
	var returnValue int
	return returnValue, nil
}

// GetPrivateBandwidthDataSummary - Retrieve a brief summary of a server's private network bandwidth
// usage. getPrivateBandwidthDataSummary retrieves a server's bandwidth allocation for its billing
// period, its estimated usage during its billing period, and an estimation of how much bandwidth it
// will use during its billing period based on its current usage. A server's projected bandwidth usage
// increases in accuracy as it progresses through its billing period.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) GetPrivateBandwidthDataSummary(ctx *slapi.RequestContext) (*SoftLayer_Container_Network_Bandwidth_Data_Summary, error) {
	var returnValue *SoftLayer_Container_Network_Bandwidth_Data_Summary
	return returnValue, nil
}

// GetPrivateBandwidthGraphImage - Retrieve a graph of a server's private network bandwidth usage over
// the specified time frame. If no time frame is specified then getPublicBandwidthGraphImage retrieves
// the last 24 hours of public bandwidth usage. getPublicBandwidthGraphImage returns a PNG image
func (softlayer_hardware_server *SoftLayer_Hardware_Server) GetPrivateBandwidthGraphImage(ctx *slapi.RequestContext, startTime string, endTime string) (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetPrivateNetworkComponent - Retrieve the private network component attached with this server.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) GetPrivateNetworkComponent(ctx *slapi.RequestContext) (*SoftLayer_Network_Component, error) {
	var returnValue *SoftLayer_Network_Component
	return returnValue, nil
}

// GetPrivateVlan - Retrieve the backend for the primary IP address of the server
func (softlayer_hardware_server *SoftLayer_Hardware_Server) GetPrivateVlan(ctx *slapi.RequestContext) (*SoftLayer_Network_Vlan, error) {
	var returnValue *SoftLayer_Network_Vlan
	return returnValue, nil
}

// GetPrivateVlanByIpAddress - Retrieve a backend network by searching for an IP address
func (softlayer_hardware_server *SoftLayer_Hardware_Server) GetPrivateVlanByIpAddress(ctx *slapi.RequestContext, ipAddress string) (*SoftLayer_Network_Vlan, error) {
	var returnValue *SoftLayer_Network_Vlan
	return returnValue, nil
}

// GetProvisionDate - <nil>
func (softlayer_hardware_server *SoftLayer_Hardware_Server) GetProvisionDate(ctx *slapi.RequestContext) (*time.Time, error) {
	var returnValue *time.Time
	return returnValue, nil
}

// GetPublicBandwidthDataSummary - Retrieve a brief summary of a server's public network bandwidth
// usage. getPublicBandwidthDataSummary retrieves a server's bandwidth allocation for its billing
// period, its estimated usage during its billing period, and an estimation of how much bandwidth it
// will use during its billing period based on its current usage. A server's projected bandwidth usage
// increases in accuracy as it progresses through its billing period.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) GetPublicBandwidthDataSummary(ctx *slapi.RequestContext) (*SoftLayer_Container_Network_Bandwidth_Data_Summary, error) {
	var returnValue *SoftLayer_Container_Network_Bandwidth_Data_Summary
	return returnValue, nil
}

// GetPublicBandwidthGraphImage - Retrieve a graph of a server's public network bandwidth usage over
// the specified time frame. If no time frame is specified then getPublicBandwidthGraphImage retrieves
// the last 24 hours of public bandwidth usage. getPublicBandwidthGraphImage returns a PNG image
// measuring 827 pixels by 293 pixels. ON THE NEW
func (softlayer_hardware_server *SoftLayer_Hardware_Server) GetPublicBandwidthGraphImage(ctx *slapi.RequestContext, startTime time.Time, endTime time.Time) (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetPublicBandwidthTotal - Retrieve the total number of bytes used by a server over a specified time
// period via the data warehouse tracking objects for this hardware.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) GetPublicBandwidthTotal(ctx *slapi.RequestContext, startTime int, endTime int) (uint64, error) {
	var returnValue uint64
	return returnValue, nil
}

// GetPublicNetworkComponent - Retrieve a SoftLayer server's public network component. Some servers are
// only connected to the private network and may not have a public network component. In that case
// getPublicNetworkComponent returns a null object.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) GetPublicNetworkComponent(ctx *slapi.RequestContext) (*SoftLayer_Network_Component, error) {
	var returnValue *SoftLayer_Network_Component
	return returnValue, nil
}

// GetPublicVlan - Retrieve the frontend for the primary IP address of the server
func (softlayer_hardware_server *SoftLayer_Hardware_Server) GetPublicVlan(ctx *slapi.RequestContext) (*SoftLayer_Network_Vlan, error) {
	var returnValue *SoftLayer_Network_Vlan
	return returnValue, nil
}

// GetPublicVlanByHostname - Retrieve the frontend network Vlan by searching the hostname of a server
func (softlayer_hardware_server *SoftLayer_Hardware_Server) GetPublicVlanByHostname(ctx *slapi.RequestContext, hostname string) (*SoftLayer_Network_Vlan, error) {
	var returnValue *SoftLayer_Network_Vlan
	return returnValue, nil
}

// GetReverseDomainRecords - Retrieve the reverse domain records associated with this server.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) GetReverseDomainRecords(ctx *slapi.RequestContext) ([]*SoftLayer_Dns_Domain, error) {
	var returnValue []*SoftLayer_Dns_Domain
	return returnValue, nil
}

// GetSensorData - Retrieve a server's hardware state via its internal sensors. Remote sensor data is
// transmitted to the SoftLayer API by way of the server's remote management card. Sensor data measures
// system temperatures, voltages, and other local server settings. Sensor data is cached for 30
// seconds. Calls made to getSensorData for the same server within 30 seconds of each other will return
// the same data. Subsequent calls will return new data once the cache expires.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) GetSensorData(ctx *slapi.RequestContext) ([]*SoftLayer_Container_RemoteManagement_SensorReading, error) {
	var returnValue []*SoftLayer_Container_RemoteManagement_SensorReading
	return returnValue, nil
}

// GetSensorDataWithGraphs - Retrieves the raw data returned from the server's remote management card.
// For more details of what is returned please refer to the getSensorData method. Along with the raw
// data, graphs for the cpu and system temperatures and fan speeds are also returned.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) GetSensorDataWithGraphs(ctx *slapi.RequestContext) (*SoftLayer_Container_RemoteManagement_SensorReadingsWithGraphs, error) {
	var returnValue *SoftLayer_Container_RemoteManagement_SensorReadingsWithGraphs
	return returnValue, nil
}

// GetServerDetails - Retrieve a server's hardware components, software, and network components.
// getServerDetails is an aggregation function that combines the results of
// [[SoftLayer_Hardware_Server::getComponents]], [[SoftLayer_Hardware_Server::getSoftware]], and
// [[SoftLayer_Hardware_Server::getNetworkComponents]] in a single container.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) GetServerDetails(ctx *slapi.RequestContext) (*SoftLayer_Container_Hardware_Server_Details, error) {
	var returnValue *SoftLayer_Container_Hardware_Server_Details
	return returnValue, nil
}

// GetServerFanSpeedGraphs - Retrieve the server's fan speeds and displays them using tachometer
// graphs. Data used to construct graphs is retrieved from the server's remote management card. All
// graphs returned will have a title associated with it.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) GetServerFanSpeedGraphs(ctx *slapi.RequestContext) ([]*SoftLayer_Container_RemoteManagement_Graphs_SensorSpeed, error) {
	var returnValue []*SoftLayer_Container_RemoteManagement_Graphs_SensorSpeed
	return returnValue, nil
}

// GetServerPowerState - Retrieves the power state for the server. The server's power status is
// retrieved from its remote management card. This will return 'on' or 'off'.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) GetServerPowerState(ctx *slapi.RequestContext) (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetServerTemperatureGraphs - Retrieve the server's temperature and displays them using thermometer
// graphs. Temperatures retrieved are CPU(s) and system temperatures. Data used to construct graphs is
// retrieved from the server's remote management card. All graphs returned will have a title associated
// with it.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) GetServerTemperatureGraphs(ctx *slapi.RequestContext) ([]*SoftLayer_Container_RemoteManagement_Graphs_SensorTemperature, error) {
	var returnValue []*SoftLayer_Container_RemoteManagement_Graphs_SensorTemperature
	return returnValue, nil
}

// GetValidBlockDeviceTemplateGroups - This method will return the list of block device template groups
// that are valid to the host. For instance, it will only retrieve images.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) GetValidBlockDeviceTemplateGroups(ctx *slapi.RequestContext, visibility string) ([]*SoftLayer_Virtual_Guest_Block_Device_Template_Group, error) {
	var returnValue []*SoftLayer_Virtual_Guest_Block_Device_Template_Group
	return returnValue, nil
}

// GetWindowsUpdateAvailableUpdates - Retrieve a list of Windows updates available for a server from
// the local SoftLayer Windows Server Update Services server. Windows servers provisioned by SoftLayer
// are configured to use the local server via the private network by default.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) GetWindowsUpdateAvailableUpdates(ctx *slapi.RequestContext) ([]*SoftLayer_Container_Utility_Microsoft_Windows_UpdateServices_UpdateItem, error) {
	var returnValue []*SoftLayer_Container_Utility_Microsoft_Windows_UpdateServices_UpdateItem
	return returnValue, nil
}

// GetWindowsUpdateInstalledUpdates - Retrieve a list of Windows updates installed on a server as
// reported by the local SoftLayer Windows Server Update Services server. Windows servers provisioned
// by SoftLayer are configured to use the local server via the private network by default.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) GetWindowsUpdateInstalledUpdates(ctx *slapi.RequestContext) ([]*SoftLayer_Container_Utility_Microsoft_Windows_UpdateServices_UpdateItem, error) {
	var returnValue []*SoftLayer_Container_Utility_Microsoft_Windows_UpdateServices_UpdateItem
	return returnValue, nil
}

// GetWindowsUpdateStatus - This method returns an update status record for this server. That record
// will specify if the server is missing updates, or has updates that must be reinstalled or require a
// reboot to go into affect.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) GetWindowsUpdateStatus(ctx *slapi.RequestContext) (*SoftLayer_Container_Utility_Microsoft_Windows_UpdateServices_Status, error) {
	var returnValue *SoftLayer_Container_Utility_Microsoft_Windows_UpdateServices_Status
	return returnValue, nil
}

// InitiateIderaBareMetalRestore - Idera Bare Metal Server Restore is a backup agent designed
// specifically for making full system restores made with Idera Server Backup.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) InitiateIderaBareMetalRestore(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// InitiateR1SoftBareMetalRestore - R1Soft Bare Metal Server Restore is an R1Soft disk agent designed
// specifically for making full system restores made with R1Soft CDP Server backup.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) InitiateR1SoftBareMetalRestore(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// IsBackendPingable - Issues a ping command and returns the success (true) or failure (false) of the
// ping command.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) IsBackendPingable(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// IsPingable - Issues a ping command and returns the success (true) or failure (false) of the ping
// command.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) IsPingable(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// IsWindowsServer - Determine if the server runs any version of the Microsoft Windows operating
// systems. Return ''true'' if it does and ''false if otherwise.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) IsWindowsServer(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// Ping - Issues a ping command to the server and returns the ping response.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) Ping(ctx *slapi.RequestContext) (string, error) {
	var returnValue string
	return returnValue, nil
}

// PowerCycle - Power off then power on the server via powerstrip. The power cycle command is
// equivalent to unplugging the server from the powerstrip and then plugging the server back into the
// powerstrip. This should only be used as a last resort. If a reboot command has been issued
// successfully in the past 20 minutes, another remote management command (rebootSoft, rebootHard,
// powerOn, powerOff and powerCycle) will not be allowed. This is to avoid any type of server failures.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) PowerCycle(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// PowerOff - This method will power off the server via the server's remote management card.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) PowerOff(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// PowerOn - Power on server via its remote management card. If a reboot command has been issued
// successfully in the past 20 minutes, another remote management command (rebootSoft, rebootHard,
// powerOn, powerOff and powerCycle) will not be allowed. This is to avoid any type of server failures.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) PowerOn(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RebootDefault - Attempts to reboot the server by issuing a reset (soft reboot) command to the
// server's remote management card. If the reset (soft reboot) attempt is unsuccessful, a power cycle
// command will be issued via the powerstrip. The power cycle command is equivalent to unplugging the
// server from the powerstrip and then plugging the server back into the powerstrip. If a reboot
// command has been issued successfully in the past 20 minutes, another remote management command
// (rebootSoft, rebootHard, powerOn, powerOff and powerCycle) will not be allowed. This is to avoid any
// type of server failures.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) RebootDefault(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RebootHard - Reboot the server by issuing a cycle command to the server's remote management card.
// This is equivalent to pressing the 'Reset' button on the server. This command is issued immediately
// and will not wait for processes to shutdown. After this command is issued, the server may take a few
// moments to boot up as server may run system disks checks. If a reboot command has been issued
// successfully in the past 20 minutes, another remote management command (rebootSoft, rebootHard,
// powerOn, powerOff and powerCycle) will not be allowed. This is to avoid any type of server failures.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) RebootHard(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RebootSoft - Reboot the server by issuing a reset command to the server's remote management card.
// This is a graceful reboot. The servers will allow all process to shutdown gracefully before
// rebooting. If a reboot command has been issued successfully in the past 20 minutes, another remote
// management command (rebootSoft, rebootHard, powerOn, powerOff and powerCycle) will not be allowed.
// This is to avoid any type of server failures.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) RebootSoft(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// ReloadCurrentOperatingSystemConfiguration - Reloads current operating system configuration. This
// service has a confirmation protocol for proceeding with the reload. To proceed with the reload
// without confirmation, simply pass in as the token parameter. To proceed with the reload with
// confirmation, simply call the service with no parameter. A token string will be returned by this
// service. The token will remain active for 10 minutes. Use this token as the parameter to confirm
// that a reload is to be performed for the server. As a precaution, we strongly recommend backing up
// all data before reloading the operating system. The reload will format the primary disk and will
// reconfigure the server to the current specifications on record. The reload will take AT 66 minutes.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) ReloadCurrentOperatingSystemConfiguration(ctx *slapi.RequestContext, token string) (string, error) {
	var returnValue string
	return returnValue, nil
}

// ReloadOperatingSystem - Reloads current operating system configuration. This service has a
// confirmation protocol for proceeding with the reload. To proceed with the reload without
// confirmation, simply pass in as the token parameter. To proceed with the reload with confirmation,
// simply call the service with no parameter. A token string will be returned by this service. The
// token will remain active for 10 minutes. Use this token as the parameter to confirm that a reload is
// to be performed for the server. As a precaution, we strongly recommend backing up all data before
// reloading the operating system. The reload will format the primary disk and will reconfigure the
// server to the current specifications on record. The reload will take AT 66 minutes.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) ReloadOperatingSystem(ctx *slapi.RequestContext, token string, config SoftLayer_Container_Hardware_Server_Configuration) (string, error) {
	var returnValue string
	return returnValue, nil
}

// RunPassmarkCertificationBenchmark - You can launch a new Passmark hardware test by selecting from
// your server list. It will bring your server offline for approximately 20 minutes while the testing
// is in progress, and will publish a certificate with the results to your hardware details page. While
// the hard drives are tested for the initial deployment, the Passmark Certificate utility will not
// test the hard drives on your live server. This is to ensure that no data is overwritten. If you
// would like to test the server's hard drives, you can have the full Passmark suite installed to your
// server free of charge through a new Support ticket. While the test itself does not overwrite any
// data on the server, it is recommended that you make full off-server backups of all data prior to
// launching the test. The Passmark hardware test is designed to force any latent hardware issues to
// the surface, so hardware failure is possible. In the event of a hardware failure during this test
// our datacenter engineers will be notified of the problem automatically. They will then replace any
// failed components to bring your server back online, and will be contacting you to ensure that impact
// on your server is minimal.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) RunPassmarkCertificationBenchmark(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// SetOperatingSystemPassword - Changes the password that we have stored in our database for a servers'
// Operating System
func (softlayer_hardware_server *SoftLayer_Hardware_Server) SetOperatingSystemPassword(ctx *slapi.RequestContext, newPassword string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// SetPrivateNetworkInterfaceSpeed - Sets the private network interface speed to the new speed. Speed
// values can only be 0 (Disconnect), 10, 100, 1000, and 10000. The new speed must be equal to or less
// than the max speed of the interface. It will take less than a minute to update the switch port
// speed. The server uplink will not be operational again until the server interface speed is updated.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) SetPrivateNetworkInterfaceSpeed(ctx *slapi.RequestContext, newSpeed int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// SetPublicNetworkInterfaceSpeed - Sets the public network interface speed to the new speed. Speed
// values can only be 0 (Disconnect), 10, 100, 1000, and 10000. The new speed must be equal to or less
// than the max speed of the interface. It will take less than a minute to update the switch port
// speed. The server uplink will not be operational again until the server interface speed is updated.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) SetPublicNetworkInterfaceSpeed(ctx *slapi.RequestContext, newSpeed int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// SetUserMetadata - Sets the data that will be written to the configuration drive.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) SetUserMetadata(ctx *slapi.RequestContext, metadata []string) ([]*SoftLayer_Hardware_Attribute, error) {
	var returnValue []*SoftLayer_Hardware_Attribute
	return returnValue, nil
}

// ShutdownPrivatePort - no documentation
func (softlayer_hardware_server *SoftLayer_Hardware_Server) ShutdownPrivatePort(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// ShutdownPublicPort - no documentation
func (softlayer_hardware_server *SoftLayer_Hardware_Server) ShutdownPublicPort(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// SparePool - The ability to place bare metal servers in a state where they are powered down, and
// ports closed yet still allocated to the customer as a part of the Spare Pool program.
func (softlayer_hardware_server *SoftLayer_Hardware_Server) SparePool(ctx *slapi.RequestContext, action string, newOrder bool) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// ValidatePartitionsForOperatingSystem - Validates a collection of partitions for an operating system
func (softlayer_hardware_server *SoftLayer_Hardware_Server) ValidatePartitionsForOperatingSystem(ctx *slapi.RequestContext, operatingSystem SoftLayer_Software_Description, partitions []SoftLayer_Hardware_Component_Partition) (bool, error) {
	var returnValue bool
	return returnValue, nil
}
