package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Network_ContentDelivery_Report_Usage - <nil>
type SoftLayer_Container_Network_ContentDelivery_Report_Usage struct {

	// ApplicationDeliveryNetworkSsl - <nil>
	ApplicationDeliveryNetworkSsl float32 `json:"applicationDeliveryNetworkSsl,omitempty"`

	// Flash - <nil>
	Flash float32 `json:"flash,omitempty"`

	// HttpsSmall - <nil>
	HttpsSmall float32 `json:"httpsSmall,omitempty"`

	// Region - <nil>
	Region string `json:"region,omitempty"`

	// SslTotal - <nil>
	SslTotal float32 `json:"sslTotal,omitempty"`

	// ApplicationDeliveryNetwork - <nil>
	ApplicationDeliveryNetwork float32 `json:"applicationDeliveryNetwork,omitempty"`

	// WindowsMedia - <nil>
	WindowsMedia float32 `json:"windowsMedia,omitempty"`

	// Http - <nil>
	Http float32 `json:"http,omitempty"`

	// HttpSmall - <nil>
	HttpSmall float32 `json:"httpSmall,omitempty"`

	// StartDate - <nil>
	StartDate *time.Time `json:"startDate,omitempty"`

	// DiskSpace - <nil>
	DiskSpace float32 `json:"diskSpace,omitempty"`

	// Https - <nil>
	Https float32 `json:"https,omitempty"`

	// StandardTotal - <nil>
	StandardTotal float32 `json:"standardTotal,omitempty"`

	// EndDate - <nil>
	EndDate *time.Time `json:"endDate,omitempty"`
}

func (softlayer_container_network_contentdelivery_report_usage *SoftLayer_Container_Network_ContentDelivery_Report_Usage) String() string {
	return "SoftLayer_Container_Network_ContentDelivery_Report_Usage"
}
