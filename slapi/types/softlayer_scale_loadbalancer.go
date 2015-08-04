package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Scale_LoadBalancer - <nil>
type SoftLayer_Scale_LoadBalancer struct {

	// Id - The load balancer configuration's internal identifier.
	Id int `json:"id,omitempty"`

	// DeleteFlag - When set and true any edit that happens on this object, be it calling edit on this
	// directly or setting as a child while editing a parent object, will end up being a deletion.
	DeleteFlag bool `json:"deleteFlag,omitempty"`

	// ModifyDate - When this load balancer configuration was last modified.
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// VirtualServerId - The identifier of the virtual server this load balancer configuration uses.
	VirtualServerId int `json:"virtualServerId,omitempty"`

	// ScaleGroupId - The identifier of the group this load balancer configuration applies to.
	ScaleGroupId int `json:"scaleGroupId,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// HealthCheckId - The identifier for the health check of this load balancer configuration
	HealthCheckId int `json:"healthCheckId,omitempty"`

	// Port - no documentation
	Port int `json:"port,omitempty"`

	// RoutingMethod - no documentation
	RoutingMethod *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Routing_Method `json:"routingMethod,omitempty"`

	// RoutingType - no documentation
	RoutingType *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Routing_Type `json:"routingType,omitempty"`

	// AllocationPercent - The percentage of connections allocated to this virtual server.
	AllocationPercent int `json:"allocationPercent,omitempty"`

	// VirtualServerPort - no documentation
	VirtualServerPort int `json:"virtualServerPort,omitempty"`

	// VirtualIpAddressId - no documentation
	VirtualIpAddressId int `json:"virtualIpAddressId,omitempty"`

	// VirtualServer - no documentation
	VirtualServer *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualServer `json:"virtualServer,omitempty"`

	// HealthCheck - no documentation
	HealthCheck *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check `json:"healthCheck,omitempty"`

	// ScaleGroup - no documentation
	ScaleGroup *SoftLayer_Scale_Group `json:"scaleGroup,omitempty"`
}

func (softlayer_scale_loadbalancer *SoftLayer_Scale_LoadBalancer) String() string {
	return "SoftLayer_Scale_LoadBalancer"
}
