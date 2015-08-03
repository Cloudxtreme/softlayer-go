package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_LoadBalancer_Service - The SoftLayer_Network_LoadBalancer_Service data type
// contains all the information relating to a specific service (destination) on a particular load
// balancer. Information retained on the object itself is the the source and destination of the
// service, routing type, weight, and whether or not the service is currently enabled.
type SoftLayer_Network_LoadBalancer_Service struct {

	// Name - no documentation
	Name string `json:"name"`

	// Type - The connection type of this service. Valid values are and The value of this variable affects
	// available values of healthCheck, listed in that variable's description
	Type string `json:"type"`

	// VipId - Unique ID for this object's parent. Probably not useful in the as this object will always be
	// a child of a VirtualIpAddress anyway.
	VipId int `json:"vipId"`

	// Weight - Weight affects the choices the load balancer makes between your services. The weight of
	// each service is expressed as a percentage of the on the virtual IP Address. All services draw from
	// the same pool of connections, so if you expect to have 4 times as much traffic as your weights for
	// the above example routes would be 40%, 40%, 10%, 10% respectively. The weights should add up to 100%
	// If you go over 100%, an exception will be thrown. Weights must be whole numbers, no fractions or
	// decimals are accepted.
	Weight int `json:"weight"`

	// DestinationIpAddress - The IP Address of the real server you wish to direct traffic to. Your account
	// must own this
	DestinationIpAddress string `json:"destinationIpAddress"`

	// DestinationPort - The port on the real server to direct the traffic. This can be different than the
	// source port. If you wish to obfuscate your traffic, you can accept requests on port 80 on the load
	// balancer, then redirect them to port 932 on your real server.
	DestinationPort int `json:"destinationPort"`

	// HealthCheck - The health check type for this service. If one is supplied, the load balancer will
	// occasionally ping your server to determine if it is still up. Servers that are down are removed from
	// the queue and will not be used to handle requests until their status returns to "up". The value of
	// the health check is determined directly by what option you have selected for the routing type. {| |-
	// ! Type ! Valid Health Checks |- | | |- | TCP | |- | FTP | |- | DNS | |- | UDP | None |}
	HealthCheck string `json:"healthCheck"`

	// HealthResponse - The expected response from the custom health check. If the requested page contains
	// this response, the check succeeds.
	HealthResponse string `json:"healthResponse"`

	// Id - Unique ID for this object, used for the getObject method, and must be set if you are editing
	// this object.
	Id int `json:"id"`

	// PeakConnections - Peak historical connections since the creation of this service. Is reset any time
	// you make a configuration change
	PeakConnections int `json:"peakConnections"`

	// Notes - Holds whether this server is up or down. Does not affect load balancer configuration at all,
	// just for the customer's informational purposes
	Notes string `json:"notes"`

	// SourcePort - The port on the load balancer that this service maps to. This is the port for incoming
	// traffic, it needs to be shared with other services to form a group.
	SourcePort int `json:"sourcePort"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Enabled - A flag (either true or false) that determines if this particular service should be enabled
	// on the load balancer. Set to false to bring the server out of rotation without losing your
	// configuration
	Enabled bool `json:"enabled"`

	// ConnectionLimit - no documentation
	ConnectionLimit int `json:"connectionLimit"`

	// HealthCheckURL - The URL provided here (starting with /) is what the load balancer will request in
	// order to perform a custom health check. You must specify either /location/of/file.html" or
	// /location/of/file.php"
	HealthCheckURL string `json:"healthCheckURL"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`
}

func (softlayer_network_loadbalancer_service *SoftLayer_Network_LoadBalancer_Service) String() string {
	return "SoftLayer_Network_LoadBalancer_Service"
}

// SoftLayer_Network_LoadBalancer_Service_Extended is SoftLayer_Network_LoadBalancer_Service with all maskable types.
type SoftLayer_Network_LoadBalancer_Service_Extended struct {
	SoftLayer_Network_LoadBalancer_Service

	// Vip - no documentation
	Vip *SoftLayer_Network_LoadBalancer_VirtualIpAddress `json:"vip"`
}

func (softlayer_network_loadbalancer_service *SoftLayer_Network_LoadBalancer_Service_Extended) String() string {
	return "SoftLayer_Network_LoadBalancer_Service"
}
