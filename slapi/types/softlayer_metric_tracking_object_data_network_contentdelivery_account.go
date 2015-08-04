package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"

	time "time"
)

// SoftLayer_Metric_Tracking_Object_Data_Network_ContentDelivery_Account -
// SoftLayer_Metric_Tracking_Object_Data_Network_ContentDelivery_Account models usage data polled from
// the CDN system.
type SoftLayer_Metric_Tracking_Object_Data_Network_ContentDelivery_Account struct {

	// FileName - The name of a file. This value is only populated in file-based bandwidth reports.
	FileName string `json:"fileName,omitempty"`

	// PopId - The internal identifier of a CDN POP (Points of Presence).
	PopId int `json:"popId,omitempty"`

	// DateTime - no documentation
	DateTime *time.Time `json:"dateTime,omitempty"`

	// Type - no documentation
	Type string `json:"type,omitempty"`

	// Counter - no documentation
	Counter slapi.Float64 `json:"counter,omitempty"`
}

func (softlayer_metric_tracking_object_data_network_contentdelivery_account *SoftLayer_Metric_Tracking_Object_Data_Network_ContentDelivery_Account) String() string {
	return "SoftLayer_Metric_Tracking_Object_Data_Network_ContentDelivery_Account"
}
