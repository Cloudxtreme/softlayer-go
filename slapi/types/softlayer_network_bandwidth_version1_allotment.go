package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
)

// SoftLayer_Network_Bandwidth_Version1_Allotment - The SoftLayer_Network_Bandwidth_Version1_Allotment
// class provides methods and data structures necessary to work with an array of hardware objects
// associated with a single Bandwidth Pooling.
type SoftLayer_Network_Bandwidth_Version1_Allotment struct {

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// AccountId - The user account identifier associated with this allotment.
	AccountId int `json:"accountId"`

	// ApplicationDeliveryControllerCount - A count of the Application Delivery Controller contained within
	// a virtual rack.
	ApplicationDeliveryControllerCount uint64 `json:"applicationDeliveryControllerCount"`

	// ApplicationDeliveryControllers - The Application Delivery Controller contained within a virtual
	// rack.
	ApplicationDeliveryControllers []*SoftLayer_Network_Application_Delivery_Controller `json:"applicationDeliveryControllers"`

	// AverageDailyPublicBandwidthUsage - The average daily public bandwidth usage for the current billing
	// cycle.
	AverageDailyPublicBandwidthUsage float32 `json:"averageDailyPublicBandwidthUsage"`

	// BandwidthAllotmentTypeId - An identifier marking this allotment as a virtual private rack (1) or a
	// bandwidth pooling(2).
	BandwidthAllotmentTypeId int `json:"bandwidthAllotmentTypeId"`

	// BareMetalInstanceCount - A count of the bare metal server instances contained within a virtual rack.
	BareMetalInstanceCount uint64 `json:"bareMetalInstanceCount"`

	// BareMetalInstances - The bare metal server instances contained within a virtual rack.
	BareMetalInstances []*SoftLayer_Hardware `json:"bareMetalInstances"`

	// BillingCycleBandwidthUsage - A virtual rack's raw bandwidth usage data for an account's current
	// billing cycle. One object is returned for each network this server is attached to.
	BillingCycleBandwidthUsage []*SoftLayer_Network_Bandwidth_Usage `json:"billingCycleBandwidthUsage"`

	// BillingCycleBandwidthUsageCount - A count of a virtual rack's raw bandwidth usage data for an
	// account's current billing cycle. One object is returned for each network this server is attached to.
	BillingCycleBandwidthUsageCount uint64 `json:"billingCycleBandwidthUsageCount"`

	// BillingCyclePrivateBandwidthUsage - A virtual rack's raw private network bandwidth usage data for an
	// account's current billing cycle.
	BillingCyclePrivateBandwidthUsage *SoftLayer_Network_Bandwidth_Usage `json:"billingCyclePrivateBandwidthUsage"`

	// BillingCyclePublicBandwidthUsage - A virtual rack's raw public network bandwidth usage data for an
	// account's current billing cycle.
	BillingCyclePublicBandwidthUsage *SoftLayer_Network_Bandwidth_Usage `json:"billingCyclePublicBandwidthUsage"`

	// BillingCyclePublicUsageTotal - The total public bandwidth used in this virtual rack for an account's
	// current billing cycle.
	BillingCyclePublicUsageTotal uint `json:"billingCyclePublicUsageTotal"`

	// BillingItem - no documentation
	BillingItem *SoftLayer_Billing_Item `json:"billingItem"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// CurrentBandwidthSummary - An object that provides commonly used bandwidth summary components for the
	// current billing cycle.
	CurrentBandwidthSummary *SoftLayer_Metric_Tracking_Object_Bandwidth_Summary `json:"currentBandwidthSummary"`

	// DetailCount - A count of the bandwidth allotment detail records associated with this virtual rack.
	DetailCount uint64 `json:"detailCount"`

	// Details - The bandwidth allotment detail records associated with this virtual rack.
	Details []*SoftLayer_Network_Bandwidth_Version1_Allotment_Detail `json:"details"`

	// EndDate - no documentation
	EndDate *time.Time `json:"endDate"`

	// Hardware - no documentation
	Hardware []*SoftLayer_Hardware `json:"hardware"`

	// HardwareCount - A count of the hardware contained within a virtual rack.
	HardwareCount uint64 `json:"hardwareCount"`

	// Id - no documentation
	Id int `json:"id"`

	// InboundPublicBandwidthUsage - The total public inbound bandwidth used in this virtual rack for an
	// account's current billing cycle.
	InboundPublicBandwidthUsage float64 `json:"inboundPublicBandwidthUsage"`

	// LocationGroup - The location group associated with this virtual rack.
	LocationGroup *SoftLayer_Location_Group `json:"locationGroup"`

	// LocationGroupId - no documentation
	LocationGroupId int `json:"locationGroupId"`

	// ManagedBareMetalInstanceCount - A count of the managed bare metal server instances contained within
	// a virtual rack.
	ManagedBareMetalInstanceCount uint64 `json:"managedBareMetalInstanceCount"`

	// ManagedBareMetalInstances - The managed bare metal server instances contained within a virtual rack.
	ManagedBareMetalInstances []*SoftLayer_Hardware `json:"managedBareMetalInstances"`

	// ManagedHardware - The managed hardware contained within a virtual rack.
	ManagedHardware []*SoftLayer_Hardware `json:"managedHardware"`

	// ManagedHardwareCount - A count of the managed hardware contained within a virtual rack.
	ManagedHardwareCount uint64 `json:"managedHardwareCount"`

	// ManagedVirtualGuestCount - A count of the managed Virtual Server contained within a virtual rack.
	ManagedVirtualGuestCount uint64 `json:"managedVirtualGuestCount"`

	// ManagedVirtualGuests - The managed Virtual Server contained within a virtual rack.
	ManagedVirtualGuests []*SoftLayer_Virtual_Guest `json:"managedVirtualGuests"`

	// MetricTrackingObject - A virtual rack's metric tracking object. This object records all periodic
	// polled data available to this rack.
	MetricTrackingObject *SoftLayer_Metric_Tracking_Object_VirtualDedicatedRack `json:"metricTrackingObject"`

	// MetricTrackingObjectId - no documentation
	MetricTrackingObjectId int `json:"metricTrackingObjectId"`

	// Name - no documentation
	Name string `json:"name"`

	// OutboundPublicBandwidthUsage - The total public outbound bandwidth used in this virtual rack for an
	// account's current billing cycle.
	OutboundPublicBandwidthUsage float64 `json:"outboundPublicBandwidthUsage"`

	// OverBandwidthAllocationFlag - Whether the bandwidth usage for this bandwidth pool for the current
	// billing cycle exceeds the allocation.
	OverBandwidthAllocationFlag int `json:"overBandwidthAllocationFlag"`

	// PrivateNetworkOnlyHardware - The private network only hardware contained within a virtual rack.
	PrivateNetworkOnlyHardware []*SoftLayer_Hardware `json:"privateNetworkOnlyHardware"`

	// PrivateNetworkOnlyHardwareCount - A count of the private network only hardware contained within a
	// virtual rack.
	PrivateNetworkOnlyHardwareCount uint64 `json:"privateNetworkOnlyHardwareCount"`

	// ProjectedOverBandwidthAllocationFlag - Whether the bandwidth usage for this bandwidth pool for the
	// current billing cycle is projected to exceed the allocation.
	ProjectedOverBandwidthAllocationFlag int `json:"projectedOverBandwidthAllocationFlag"`

	// ProjectedPublicBandwidthUsage - The projected public outbound bandwidth for this virtual server for
	// the current billing cycle.
	ProjectedPublicBandwidthUsage float32 `json:"projectedPublicBandwidthUsage"`

	// ServiceProvider - <nil>
	ServiceProvider *SoftLayer_Service_Provider `json:"serviceProvider"`

	// ServiceProviderId - no documentation
	ServiceProviderId int `json:"serviceProviderId"`

	// TotalBandwidthAllocated - The combined allocated bandwidth for all servers in a virtual rack.
	TotalBandwidthAllocated uint64 `json:"totalBandwidthAllocated"`

	// VirtualGuestCount - A count of the Virtual Server contained within a virtual rack.
	VirtualGuestCount uint64 `json:"virtualGuestCount"`

	// VirtualGuests - The Virtual Server contained within a virtual rack.
	VirtualGuests []*SoftLayer_Virtual_Guest `json:"virtualGuests"`
}

// CreateObject - Create a allotment for servers to pool bandwidth and avoid overages in billing if
// they use more than there allocated bandwidth.
func (softlayer_network_bandwidth_version1_allotment *SoftLayer_Network_Bandwidth_Version1_Allotment) CreateObject(commonOptions *slapi.CommonOptions, templateObject SoftLayer_Network_Bandwidth_Version1_Allotment) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObject - Edit a bandwidth allotment's local properties. Currently you may only change an
// allotment's name. Use the
// [[SoftLayer_Network_Bandwidth_Version1_Allotment::reassignServers|reassignServers()]] and
// [[SoftLayer_Network_Bandwidth_Version1_Allotment::unassignServers|unassignServers()]] methods to
// move servers in and out of your allotments.
func (softlayer_network_bandwidth_version1_allotment *SoftLayer_Network_Bandwidth_Version1_Allotment) EditObject(commonOptions *slapi.CommonOptions, templateObject SoftLayer_Network_Bandwidth_Version1_Allotment) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetBackendBandwidthByHour - This method recurses through all servers on a Bandwidth Pool for 24 hour
// time span starting at a given date/time. To get the private data set for all servers on a Bandwidth
// Pool from midnight Feb 1st, 2008 to 23:59 on Feb 1st, you would pass a parameter of '02/01/2008
// 0:00'. The ending date / time is calculated for you to prevent requesting data from the server for
// periods larger than 24 hours as this method requires processing a lot of data records and can get
// slow at times.
func (softlayer_network_bandwidth_version1_allotment *SoftLayer_Network_Bandwidth_Version1_Allotment) GetBackendBandwidthByHour(commonOptions *slapi.CommonOptions, date time.Time) ([]*SoftLayer_Container_Network_Bandwidth_Version1_Usage, error) {
	var returnValue []*SoftLayer_Container_Network_Bandwidth_Version1_Usage
	return returnValue, nil
}

// GetBackendBandwidthUse - This method recurses through all servers on a Bandwidth Pool between the
// given start and end dates to retrieve public bandwidth data.
func (softlayer_network_bandwidth_version1_allotment *SoftLayer_Network_Bandwidth_Version1_Allotment) GetBackendBandwidthUse(commonOptions *slapi.CommonOptions, startDate time.Time, endDate time.Time) ([]*SoftLayer_Network_Bandwidth_Version1_Usage_Detail, error) {
	var returnValue []*SoftLayer_Network_Bandwidth_Version1_Usage_Detail
	return returnValue, nil
}

// GetBandwidthForDateRange - Retrieve a collection of bandwidth data from an individual public or
// private network tracking object. Data is ideal if you with to employ your own traffic storage and
// graphing systems.
func (softlayer_network_bandwidth_version1_allotment *SoftLayer_Network_Bandwidth_Version1_Allotment) GetBandwidthForDateRange(commonOptions *slapi.CommonOptions, startDate time.Time, endDate time.Time) ([]*SoftLayer_Metric_Tracking_Object_Data, error) {
	var returnValue []*SoftLayer_Metric_Tracking_Object_Data
	return returnValue, nil
}

// GetBandwidthImage - This method recurses through all servers on a Bandwidth Pool for a given
// snapshot range, gathers the necessary parameters, and then calls the bandwidth graphing server. The
// return result is a container that includes the min and max dates for all servers to be used in the
// query, as well as an image in PNG format. Depending on the current workload of the graphing server
// this method make take some time to return. As such, it is recommended the page not wait for the
// result without providing the user some sort of visual feedback. To facilitate this, an option to
// gather the parameters but not actually draw the image is included. This allows you to call the
// method twice, once for setting up your page, and then again to actually do the time consuming draw.
func (softlayer_network_bandwidth_version1_allotment *SoftLayer_Network_Bandwidth_Version1_Allotment) GetBandwidthImage(commonOptions *slapi.CommonOptions, networkType string, snapshotRange string, draw bool, dateSpecified time.Time) (*SoftLayer_Container_Bandwidth_GraphOutputs, error) {
	var returnValue *SoftLayer_Container_Bandwidth_GraphOutputs
	return returnValue, nil
}

// GetCustomBandwidthDataByDate - no documentation
func (softlayer_network_bandwidth_version1_allotment *SoftLayer_Network_Bandwidth_Version1_Allotment) GetCustomBandwidthDataByDate(commonOptions *slapi.CommonOptions, graphData SoftLayer_Container_Graph) (*SoftLayer_Container_Graph, error) {
	var returnValue *SoftLayer_Container_Graph
	return returnValue, nil
}

// GetFrontendBandwidthByHour - This method recurses through all servers on a Bandwidth Pool for 24
// hour time span starting at a given date/time. To get the public data set for all servers on a
// Bandwidth Pool from midnight Feb 1st, 2008 to 23:59 on Feb 1st, you would pass a parameter of
// '02/01/2008 0:00'. The ending date / time is calculated for you to prevent requesting data from the
// server for periods larger than 24 hours as this method requires processing a lot of data records and
// can get slow at times.
func (softlayer_network_bandwidth_version1_allotment *SoftLayer_Network_Bandwidth_Version1_Allotment) GetFrontendBandwidthByHour(commonOptions *slapi.CommonOptions, date time.Time) ([]*SoftLayer_Container_Network_Bandwidth_Version1_Usage, error) {
	var returnValue []*SoftLayer_Container_Network_Bandwidth_Version1_Usage
	return returnValue, nil
}

// GetFrontendBandwidthUse - This method recurses through all servers on a Bandwidth Pool between the
// given start and end dates to retrieve private bandwidth data.
func (softlayer_network_bandwidth_version1_allotment *SoftLayer_Network_Bandwidth_Version1_Allotment) GetFrontendBandwidthUse(commonOptions *slapi.CommonOptions, startDate time.Time, endDate time.Time) ([]*SoftLayer_Network_Bandwidth_Version1_Usage_Detail, error) {
	var returnValue []*SoftLayer_Network_Bandwidth_Version1_Usage_Detail
	return returnValue, nil
}

// GetObject - getObject retrieves the SoftLayer_Network_Bandwidth_Version1_Allotment object whose ID
// number corresponds to the ID number of the init parameter passed to the SoftLayer_Hardware service.
// You can only retrieve an allotment associated with the account that your portal user is assigned to.
func (softlayer_network_bandwidth_version1_allotment *SoftLayer_Network_Bandwidth_Version1_Allotment) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Network_Bandwidth_Version1_Allotment, error) {
	var returnValue *SoftLayer_Network_Bandwidth_Version1_Allotment
	return returnValue, nil
}

// New_getBandwidthImage - This method recurses through all servers on a Bandwidth Pool for a given
// snapshot range, gathers the necessary parameters, and then calls the bandwidth graphing server. The
// return result is a container that includes the min and max dates for all servers to be used in the
// query, as well as an image in PNG format. This method uses the new and improved drawing routines
// which should return in a reasonable time frame now that the new backend data warehouse is used.
func (softlayer_network_bandwidth_version1_allotment *SoftLayer_Network_Bandwidth_Version1_Allotment) New_getBandwidthImage(commonOptions *slapi.CommonOptions, networkType string, snapshotRange string, draw bool, dateSpecified time.Time) (*SoftLayer_Container_Bandwidth_GraphOutputs, error) {
	var returnValue *SoftLayer_Container_Bandwidth_GraphOutputs
	return returnValue, nil
}

// ReassignServers - This method will reassign a collection of SoftLayer hardware to a bandwidth
// allotment Bandwidth Pool.
func (softlayer_network_bandwidth_version1_allotment *SoftLayer_Network_Bandwidth_Version1_Allotment) ReassignServers(commonOptions *slapi.CommonOptions, templateObjects []SoftLayer_Hardware, newAllotmentId int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RequestVdrCancellation - This will remove a bandwidth pooling from a customer's allotments by
// cancelling the billing item. All servers in that allotment will get moved to the account's vpr.
func (softlayer_network_bandwidth_version1_allotment *SoftLayer_Network_Bandwidth_Version1_Allotment) RequestVdrCancellation(commonOptions *slapi.CommonOptions) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RequestVdrContentUpdates - This will move servers into a bandwidth pool, removing them from their
// previous bandwidth pool and optionally remove the bandwidth pool on completion.
func (softlayer_network_bandwidth_version1_allotment *SoftLayer_Network_Bandwidth_Version1_Allotment) RequestVdrContentUpdates(commonOptions *slapi.CommonOptions, hardwareToAdd []SoftLayer_Hardware, hardwareToRemove []SoftLayer_Hardware, cloudsToAdd []SoftLayer_Virtual_Guest, cloudsToRemove []SoftLayer_Virtual_Guest, optionalAllotmentId int, adcToAdd []SoftLayer_Network_Application_Delivery_Controller, adcToRemove []SoftLayer_Network_Application_Delivery_Controller) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// SetVdrContent - This will update the bandwith pool to the servers provided. Servers currently in the
// bandwith pool not provided on update will be removed. Servers provided on update not currently in
// the bandwith pool will be added. If all servers are removed, this removes the bandwidth pool on
// completion.
func (softlayer_network_bandwidth_version1_allotment *SoftLayer_Network_Bandwidth_Version1_Allotment) SetVdrContent(commonOptions *slapi.CommonOptions, hardware []SoftLayer_Hardware, bareMetalServers []SoftLayer_Hardware, virtualServerInstance []SoftLayer_Virtual_Guest, adc []SoftLayer_Network_Application_Delivery_Controller, optionalAllotmentId int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// UnassignServers - This method will reassign a collection of SoftLayer hardware to the virtual
// private rack
func (softlayer_network_bandwidth_version1_allotment *SoftLayer_Network_Bandwidth_Version1_Allotment) UnassignServers(commonOptions *slapi.CommonOptions, templateObjects []SoftLayer_Hardware) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// VoidPendingServerMove - This method will void a pending server removal from this bandwidth pooling.
// Pass in the id of the hardware object or virtual guest you wish to update. Assuming that object is
// currently pending removal from the bandwidth pool at the start of the next billing cycle, the
// bandwidth pool member status will be restored and the pending cancellation removed.
func (softlayer_network_bandwidth_version1_allotment *SoftLayer_Network_Bandwidth_Version1_Allotment) VoidPendingServerMove(commonOptions *slapi.CommonOptions, id int, type_ string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// VoidPendingVdrCancellation - This method will void a pending cancellation on a bandwidth pool. Note
// however any servers that belonged to the rack will have to be restored individually using the method
// voidPendingServerMove($id, $type).
func (softlayer_network_bandwidth_version1_allotment *SoftLayer_Network_Bandwidth_Version1_Allotment) VoidPendingVdrCancellation(commonOptions *slapi.CommonOptions) (bool, error) {
	var returnValue bool
	return returnValue, nil
}
