package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Scale_LoadBalancer - <nil>
type SoftLayer_Scale_LoadBalancer struct {

	// DeleteFlag - When set and true any edit that happens on this object, be it calling edit on this
	// directly or setting as a child while editing a parent object, will end up being a deletion.
	DeleteFlag bool `json:"deleteFlag"`

	// ModifyDate - When this load balancer configuration was last modified.
	ModifyDate *time.Time `json:"modifyDate"`

	// Port - no documentation
	Port int `json:"port"`

	// ScaleGroupId - The identifier of the group this load balancer configuration applies to.
	ScaleGroupId int `json:"scaleGroupId"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// HealthCheckId - The identifier for the health check of this load balancer configuration
	HealthCheckId int `json:"healthCheckId"`

	// Id - The load balancer configuration's internal identifier.
	Id int `json:"id"`

	// VirtualServerId - The identifier of the virtual server this load balancer configuration uses.
	VirtualServerId int `json:"virtualServerId"`
}

// SoftLayer_Scale_LoadBalancer_Extended is SoftLayer_Scale_LoadBalancer with all maskable types.
type SoftLayer_Scale_LoadBalancer_Extended struct {
	SoftLayer_Scale_LoadBalancer

	// AllocationPercent - The percentage of connections allocated to this virtual server.
	AllocationPercent int `json:"allocationPercent"`

	// HealthCheck - no documentation
	HealthCheck *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check `json:"healthCheck"`

	// ScaleGroup - no documentation
	ScaleGroup *SoftLayer_Scale_Group `json:"scaleGroup"`

	// VirtualServerPort - no documentation
	VirtualServerPort int `json:"virtualServerPort"`

	// RoutingMethod - no documentation
	RoutingMethod *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Routing_Method `json:"routingMethod"`

	// RoutingType - no documentation
	RoutingType *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Routing_Type `json:"routingType"`

	// VirtualIpAddressId - no documentation
	VirtualIpAddressId int `json:"virtualIpAddressId"`

	// VirtualServer - no documentation
	VirtualServer *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualServer `json:"virtualServer"`
}

func (softlayer_scale_loadbalancer *SoftLayer_Scale_LoadBalancer) String() string {
	return "SoftLayer_Scale_LoadBalancer"
}
