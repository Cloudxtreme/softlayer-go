package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Network_ContentDelivery_Report_Usage - <nil>
type SoftLayer_Container_Network_ContentDelivery_Report_Usage struct {

	// Region - <nil>
	Region string `json:"region"`

	// StartDate - <nil>
	StartDate *time.Time `json:"startDate"`

	// DiskSpace - <nil>
	DiskSpace float32 `json:"diskSpace"`

	// Flash - <nil>
	Flash float32 `json:"flash"`

	// Https - <nil>
	Https float32 `json:"https"`

	// ApplicationDeliveryNetworkSsl - <nil>
	ApplicationDeliveryNetworkSsl float32 `json:"applicationDeliveryNetworkSsl"`

	// Http - <nil>
	Http float32 `json:"http"`

	// StandardTotal - <nil>
	StandardTotal float32 `json:"standardTotal"`

	// HttpSmall - <nil>
	HttpSmall float32 `json:"httpSmall"`

	// HttpsSmall - <nil>
	HttpsSmall float32 `json:"httpsSmall"`

	// SslTotal - <nil>
	SslTotal float32 `json:"sslTotal"`

	// ApplicationDeliveryNetwork - <nil>
	ApplicationDeliveryNetwork float32 `json:"applicationDeliveryNetwork"`

	// EndDate - <nil>
	EndDate *time.Time `json:"endDate"`

	// WindowsMedia - <nil>
	WindowsMedia float32 `json:"windowsMedia"`
}

func (softlayer_container_network_contentdelivery_report_usage *SoftLayer_Container_Network_ContentDelivery_Report_Usage) String() string {
	return "SoftLayer_Container_Network_ContentDelivery_Report_Usage"
}
