package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"

	time "time"
)

// SoftLayer_Network_Bandwidth_Version1_Allotment - The SoftLayer_Network_Bandwidth_Version1_Allotment
// class provides methods and data structures necessary to work with an array of hardware objects
// associated with a single Bandwidth Pooling.
type SoftLayer_Network_Bandwidth_Version1_Allotment struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// ServiceProviderId - no documentation
	ServiceProviderId int `json:"serviceProviderId,omitempty"`

	// EndDate - no documentation
	EndDate *time.Time `json:"endDate,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// AccountId - The user account identifier associated with this allotment.
	AccountId int `json:"accountId,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// BandwidthAllotmentTypeId - An identifier marking this allotment as a virtual private rack (1) or a
	// bandwidth pooling(2).
	BandwidthAllotmentTypeId int `json:"bandwidthAllotmentTypeId,omitempty"`

	// LocationGroupId - no documentation
	LocationGroupId int `json:"locationGroupId,omitempty"`

	// ManagedVirtualGuestCount - A count of the managed Virtual Server contained within a virtual rack.
	ManagedVirtualGuestCount uint64 `json:"managedVirtualGuestCount,omitempty"`

	// BillingCyclePublicBandwidthUsage - A virtual rack's raw public network bandwidth usage data for an
	// account's current billing cycle.
	BillingCyclePublicBandwidthUsage *SoftLayer_Network_Bandwidth_Usage `json:"billingCyclePublicBandwidthUsage,omitempty"`

	// BillingCyclePublicUsageTotal - The total public bandwidth used in this virtual rack for an account's
	// current billing cycle.
	BillingCyclePublicUsageTotal uint `json:"billingCyclePublicUsageTotal,omitempty"`

	// VirtualGuests - The Virtual Server contained within a virtual rack.
	VirtualGuests []*SoftLayer_Virtual_Guest `json:"virtualGuests,omitempty"`

	// BareMetalInstanceCount - A count of the bare metal server instances contained within a virtual rack.
	BareMetalInstanceCount uint64 `json:"bareMetalInstanceCount,omitempty"`

	// DetailCount - A count of the bandwidth allotment detail records associated with this virtual rack.
	DetailCount uint64 `json:"detailCount,omitempty"`

	// BillingCyclePrivateBandwidthUsage - A virtual rack's raw private network bandwidth usage data for an
	// account's current billing cycle.
	BillingCyclePrivateBandwidthUsage *SoftLayer_Network_Bandwidth_Usage `json:"billingCyclePrivateBandwidthUsage,omitempty"`

	// ManagedHardware - The managed hardware contained within a virtual rack.
	ManagedHardware []*SoftLayer_Hardware `json:"managedHardware,omitempty"`

	// PrivateNetworkOnlyHardware - The private network only hardware contained within a virtual rack.
	PrivateNetworkOnlyHardware []*SoftLayer_Hardware `json:"privateNetworkOnlyHardware,omitempty"`

	// ProjectedOverBandwidthAllocationFlag - Whether the bandwidth usage for this bandwidth pool for the
	// current billing cycle is projected to exceed the allocation.
	ProjectedOverBandwidthAllocationFlag int `json:"projectedOverBandwidthAllocationFlag,omitempty"`

	// ManagedHardwareCount - A count of the managed hardware contained within a virtual rack.
	ManagedHardwareCount uint64 `json:"managedHardwareCount,omitempty"`

	// Details - The bandwidth allotment detail records associated with this virtual rack.
	Details []*SoftLayer_Network_Bandwidth_Version1_Allotment_Detail `json:"details,omitempty"`

	// ProjectedPublicBandwidthUsage - The projected public outbound bandwidth for this virtual server for
	// the current billing cycle.
	ProjectedPublicBandwidthUsage slapi.Float64 `json:"projectedPublicBandwidthUsage,omitempty"`

	// ServiceProvider - <nil>
	ServiceProvider *SoftLayer_Service_Provider `json:"serviceProvider,omitempty"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account,omitempty"`

	// AverageDailyPublicBandwidthUsage - The average daily public bandwidth usage for the current billing
	// cycle.
	AverageDailyPublicBandwidthUsage slapi.Float64 `json:"averageDailyPublicBandwidthUsage,omitempty"`

	// LocationGroup - The location group associated with this virtual rack.
	LocationGroup *SoftLayer_Location_Group `json:"locationGroup,omitempty"`

	// MetricTrackingObject - A virtual rack's metric tracking object. This object records all periodic
	// polled data available to this rack.
	MetricTrackingObject *SoftLayer_Metric_Tracking_Object_VirtualDedicatedRack `json:"metricTrackingObject,omitempty"`

	// TotalBandwidthAllocated - The combined allocated bandwidth for all servers in a virtual rack.
	TotalBandwidthAllocated uint64 `json:"totalBandwidthAllocated,omitempty"`

	// ApplicationDeliveryControllers - The Application Delivery Controller contained within a virtual
	// rack.
	ApplicationDeliveryControllers []*SoftLayer_Network_Application_Delivery_Controller `json:"applicationDeliveryControllers,omitempty"`

	// BillingCycleBandwidthUsage - A virtual rack's raw bandwidth usage data for an account's current
	// billing cycle. One object is returned for each network this server is attached to.
	BillingCycleBandwidthUsage []*SoftLayer_Network_Bandwidth_Usage `json:"billingCycleBandwidthUsage,omitempty"`

	// InboundPublicBandwidthUsage - The total public inbound bandwidth used in this virtual rack for an
	// account's current billing cycle.
	InboundPublicBandwidthUsage slapi.Float64 `json:"inboundPublicBandwidthUsage,omitempty"`

	// ManagedVirtualGuests - The managed Virtual Server contained within a virtual rack.
	ManagedVirtualGuests []*SoftLayer_Virtual_Guest `json:"managedVirtualGuests,omitempty"`

	// ManagedBareMetalInstanceCount - A count of the managed bare metal server instances contained within
	// a virtual rack.
	ManagedBareMetalInstanceCount uint64 `json:"managedBareMetalInstanceCount,omitempty"`

	// PrivateNetworkOnlyHardwareCount - A count of the private network only hardware contained within a
	// virtual rack.
	PrivateNetworkOnlyHardwareCount uint64 `json:"privateNetworkOnlyHardwareCount,omitempty"`

	// ApplicationDeliveryControllerCount - A count of the Application Delivery Controller contained within
	// a virtual rack.
	ApplicationDeliveryControllerCount uint64 `json:"applicationDeliveryControllerCount,omitempty"`

	// BareMetalInstances - The bare metal server instances contained within a virtual rack.
	BareMetalInstances []*SoftLayer_Hardware `json:"bareMetalInstances,omitempty"`

	// BillingItem - no documentation
	BillingItem *SoftLayer_Billing_Item `json:"billingItem,omitempty"`

	// CurrentBandwidthSummary - An object that provides commonly used bandwidth summary components for the
	// current billing cycle.
	CurrentBandwidthSummary *SoftLayer_Metric_Tracking_Object_Bandwidth_Summary `json:"currentBandwidthSummary,omitempty"`

	// OverBandwidthAllocationFlag - Whether the bandwidth usage for this bandwidth pool for the current
	// billing cycle exceeds the allocation.
	OverBandwidthAllocationFlag int `json:"overBandwidthAllocationFlag,omitempty"`

	// HardwareCount - A count of the hardware contained within a virtual rack.
	HardwareCount uint64 `json:"hardwareCount,omitempty"`

	// VirtualGuestCount - A count of the Virtual Server contained within a virtual rack.
	VirtualGuestCount uint64 `json:"virtualGuestCount,omitempty"`

	// ManagedBareMetalInstances - The managed bare metal server instances contained within a virtual rack.
	ManagedBareMetalInstances []*SoftLayer_Hardware `json:"managedBareMetalInstances,omitempty"`

	// OutboundPublicBandwidthUsage - The total public outbound bandwidth used in this virtual rack for an
	// account's current billing cycle.
	OutboundPublicBandwidthUsage slapi.Float64 `json:"outboundPublicBandwidthUsage,omitempty"`

	// Hardware - no documentation
	Hardware []*SoftLayer_Hardware `json:"hardware,omitempty"`

	// MetricTrackingObjectId - no documentation
	MetricTrackingObjectId int `json:"metricTrackingObjectId,omitempty"`

	// BillingCycleBandwidthUsageCount - A count of a virtual rack's raw bandwidth usage data for an
	// account's current billing cycle. One object is returned for each network this server is attached to.
	BillingCycleBandwidthUsageCount uint64 `json:"billingCycleBandwidthUsageCount,omitempty"`
}

func (softlayer_network_bandwidth_version1_allotment *SoftLayer_Network_Bandwidth_Version1_Allotment) String() string {
	return "SoftLayer_Network_Bandwidth_Version1_Allotment"
}
