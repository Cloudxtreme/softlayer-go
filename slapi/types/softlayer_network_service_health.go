package sl

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

	// Location - no documentation
	Location *SoftLayer_Location `json:"location"`

	// LocationId - no documentation
	LocationId int `json:"locationId"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// Status - The status portion of a service/status relationship.
	Status *SoftLayer_Network_Service_Health_Status `json:"status"`

	// StatusId - no documentation
	StatusId int `json:"statusId"`
}

func (softlayer_network_service_health *SoftLayer_Network_Service_Health) String() string {
	return "SoftLayer_Network_Service_Health"
}
