package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Service_Health - Many general services that SoftLayer provides are tracked on the
// customer portal with a quick status message. These status message provide users with a quick
// reference to the health of a service, whether it's up or down. These services include SoftLayer's
// Internet backbone connections, VPN entry points, and router networks. The
// SoftLayer_Network_Service_Health data type provides the relationship between these services and
// their health status.
type SoftLayer_Network_Service_Health struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// LocationId - no documentation
	LocationId int `json:"locationId"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// StatusId - no documentation
	StatusId int `json:"statusId"`
}

// SoftLayer_Network_Service_Health_Extended is SoftLayer_Network_Service_Health with all maskable types.
type SoftLayer_Network_Service_Health_Extended struct {
	SoftLayer_Network_Service_Health

	// Location - no documentation
	Location *SoftLayer_Location `json:"location"`

	// Status - The status portion of a service/status relationship.
	Status *SoftLayer_Network_Service_Health_Status `json:"status"`
}

func (softlayer_network_service_health *SoftLayer_Network_Service_Health) String() string {
	return "SoftLayer_Network_Service_Health"
}
