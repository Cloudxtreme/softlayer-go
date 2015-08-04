package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"

	time "time"
)

// SoftLayer_Container_Network_ContentDelivery_Report_Usage - <nil>
type SoftLayer_Container_Network_ContentDelivery_Report_Usage struct {

	// SslTotal - <nil>
	SslTotal slapi.Float64 `json:"sslTotal,omitempty"`

	// StandardTotal - <nil>
	StandardTotal slapi.Float64 `json:"standardTotal,omitempty"`

	// StartDate - <nil>
	StartDate *time.Time `json:"startDate,omitempty"`

	// WindowsMedia - <nil>
	WindowsMedia slapi.Float64 `json:"windowsMedia,omitempty"`

	// EndDate - <nil>
	EndDate *time.Time `json:"endDate,omitempty"`

	// Flash - <nil>
	Flash slapi.Float64 `json:"flash,omitempty"`

	// HttpSmall - <nil>
	HttpSmall slapi.Float64 `json:"httpSmall,omitempty"`

	// Https - <nil>
	Https slapi.Float64 `json:"https,omitempty"`

	// ApplicationDeliveryNetwork - <nil>
	ApplicationDeliveryNetwork slapi.Float64 `json:"applicationDeliveryNetwork,omitempty"`

	// DiskSpace - <nil>
	DiskSpace slapi.Float64 `json:"diskSpace,omitempty"`

	// Http - <nil>
	Http slapi.Float64 `json:"http,omitempty"`

	// ApplicationDeliveryNetworkSsl - <nil>
	ApplicationDeliveryNetworkSsl slapi.Float64 `json:"applicationDeliveryNetworkSsl,omitempty"`

	// HttpsSmall - <nil>
	HttpsSmall slapi.Float64 `json:"httpsSmall,omitempty"`

	// Region - <nil>
	Region string `json:"region,omitempty"`
}

func (softlayer_container_network_contentdelivery_report_usage *SoftLayer_Container_Network_ContentDelivery_Report_Usage) String() string {
	return "SoftLayer_Container_Network_ContentDelivery_Report_Usage"
}
