package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Network_ContentDelivery_Report_Usage - <nil>
type SoftLayer_Container_Network_ContentDelivery_Report_Usage struct {

	// EndDate - <nil>
	EndDate *time.Time `json:"endDate"`

	// Flash - <nil>
	Flash float32 `json:"flash"`

	// StandardTotal - <nil>
	StandardTotal float32 `json:"standardTotal"`

	// WindowsMedia - <nil>
	WindowsMedia float32 `json:"windowsMedia"`

	// ApplicationDeliveryNetworkSsl - <nil>
	ApplicationDeliveryNetworkSsl float32 `json:"applicationDeliveryNetworkSsl"`

	// Https - <nil>
	Https float32 `json:"https"`

	// HttpsSmall - <nil>
	HttpsSmall float32 `json:"httpsSmall"`

	// DiskSpace - <nil>
	DiskSpace float32 `json:"diskSpace"`

	// Http - <nil>
	Http float32 `json:"http"`

	// HttpSmall - <nil>
	HttpSmall float32 `json:"httpSmall"`

	// SslTotal - <nil>
	SslTotal float32 `json:"sslTotal"`

	// ApplicationDeliveryNetwork - <nil>
	ApplicationDeliveryNetwork float32 `json:"applicationDeliveryNetwork"`

	// Region - <nil>
	Region string `json:"region"`

	// StartDate - <nil>
	StartDate *time.Time `json:"startDate"`
}

func (softlayer_container_network_contentdelivery_report_usage *SoftLayer_Container_Network_ContentDelivery_Report_Usage) String() string {
	return "SoftLayer_Container_Network_ContentDelivery_Report_Usage"
}
