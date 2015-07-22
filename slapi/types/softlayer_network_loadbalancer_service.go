package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
)

// SoftLayer_Network_LoadBalancer_Service - The SoftLayer_Network_LoadBalancer_Service data type
// contains all the information relating to a specific service (destination) on a particular load
// balancer. Information retained on the object itself is the the source and destination of the
// service, routing type, weight, and whether or not the service is currently enabled.
type SoftLayer_Network_LoadBalancer_Service struct {

	// ConnectionLimit - no documentation
	ConnectionLimit int `json:"connectionLimit"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// DestinationIpAddress - The IP Address of the real server you wish to direct traffic to. Your account
	// must own this
	DestinationIpAddress string `json:"destinationIpAddress"`

	// DestinationPort - The port on the real server to direct the traffic. This can be different than the
	// source port. If you wish to obfuscate your traffic, you can accept requests on port 80 on the load
	// balancer, then redirect them to port 932 on your real server.
	DestinationPort int `json:"destinationPort"`

	// Enabled - A flag (either true or false) that determines if this particular service should be enabled
	// on the load balancer. Set to false to bring the server out of rotation without losing your
	// configuration
	Enabled bool `json:"enabled"`

	// HealthCheck - The health check type for this service. If one is supplied, the load balancer will
	// occasionally ping your server to determine if it is still up. Servers that are down are removed from
	// the queue and will not be used to handle requests until their status returns to "up". The value of
	// the health check is determined directly by what option you have selected for the routing type. {| |-
	// ! Type ! Valid Health Checks |- | | |- | TCP | |- | FTP | |- | DNS | |- | UDP | None |}
	HealthCheck string `json:"healthCheck"`

	// HealthCheckURL - The URL provided here (starting with /) is what the load balancer will request in
	// order to perform a custom health check. You must specify either /location/of/file.html" or
	// /location/of/file.php"
	HealthCheckURL string `json:"healthCheckURL"`

	// HealthResponse - The expected response from the custom health check. If the requested page contains
	// this response, the check succeeds.
	HealthResponse string `json:"healthResponse"`

	// Id - Unique ID for this object, used for the getObject method, and must be set if you are editing
	// this object.
	Id int `json:"id"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// Name - no documentation
	Name string `json:"name"`

	// Notes - Holds whether this server is up or down. Does not affect load balancer configuration at all,
	// just for the customer's informational purposes
	Notes string `json:"notes"`

	// PeakConnections - Peak historical connections since the creation of this service. Is reset any time
	// you make a configuration change
	PeakConnections int `json:"peakConnections"`

	// SourcePort - The port on the load balancer that this service maps to. This is the port for incoming
	// traffic, it needs to be shared with other services to form a group.
	SourcePort int `json:"sourcePort"`

	// Type - The connection type of this service. Valid values are and The value of this variable affects
	// available values of healthCheck, listed in that variable's description
	Type string `json:"type"`

	// Vip - no documentation
	Vip *SoftLayer_Network_LoadBalancer_VirtualIpAddress `json:"vip"`

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
}

// DeleteObject - Calling deleteObject on a particular server will remove it from the load balancer.
// This is the only way to remove a service from your load balancer. If you wish to remove a server,
// first call this function, then reload the virtualIpAddress object and edit the remaining services to
// reflect the other changes that you wish to make.
func (softlayer_network_loadbalancer_service *SoftLayer_Network_LoadBalancer_Service) DeleteObject(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetGraphImage - Get the graph image for a load balancer service based on the supplied graph type and
// metric. The available graph types are: 'connections' and 'status', and the available metrics are:
// 'day', 'week' and 'month'. This method returns the raw binary image data.
func (softlayer_network_loadbalancer_service *SoftLayer_Network_LoadBalancer_Service) GetGraphImage(ctx *slapi.RequestContext, graphType string, metric string) (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetObject - getObject retrieves the SoftLayer_Network_LoadBalancer_Service object whose ID number
// corresponds to the ID number of the init parameter passed to the
// SoftLayer_Network_LoadBalancer_Service service. You can only retrieve services on load balancers
// assigned to your account, and it is recommended that you simply retrieve the entire load balancer,
// as an individual service has no explicit purpose without its "siblings".
func (softlayer_network_loadbalancer_service *SoftLayer_Network_LoadBalancer_Service) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Network_LoadBalancer_Service, error) {
	var returnValue *SoftLayer_Network_LoadBalancer_Service
	return returnValue, nil
}

// GetStatus - Returns an array of SoftLayer_Container_Network_LoadBalancer_StatusEntry objects. A
// SoftLayer_Container_Network_LoadBalancer_StatusEntry object has two variables, "Label" and "Value"
// Calling this function executes a command on the physical load balancer itself, and therefore should
// be called infrequently. For a general idea of the load balancer service, use the "peakConnections"
// variable on the Type Possible values for "Label" are: * IP Address * Port * Server Status * Load
// Status * Current Connections * Total Hits Not all labels are guaranteed to be returned.
func (softlayer_network_loadbalancer_service *SoftLayer_Network_LoadBalancer_Service) GetStatus(ctx *slapi.RequestContext) ([]*SoftLayer_Container_Network_LoadBalancer_StatusEntry, error) {
	var returnValue []*SoftLayer_Container_Network_LoadBalancer_StatusEntry
	return returnValue, nil
}

// ResetPeakConnections - Calling resetPeakConnections will set the peakConnections variable to zero on
// this particular object. Peak connections will continue to increase normally after this method call,
// it will only temporarily reset the statistic to zero, until the next time it is polled.
func (softlayer_network_loadbalancer_service *SoftLayer_Network_LoadBalancer_Service) ResetPeakConnections(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}
