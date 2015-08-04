package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Metric_Tracking_Object_Virtual_Storage_Repository - <nil>
type SoftLayer_Metric_Tracking_Object_Virtual_Storage_Repository struct {

	// Data - no documentation
	Data []*SoftLayer_Metric_Tracking_Object_Data `json:"data,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// Label - no documentation
	Label string `json:"label,omitempty"`

	// ResourceTableId - The identifier of the existing resource this object is attempting to track.
	ResourceTableId int `json:"resourceTableId,omitempty"`

	// StartDate - The date this tracker began tracking this particular resource.
	StartDate *time.Time `json:"startDate,omitempty"`
}

func (softlayer_metric_tracking_object_virtual_storage_repository *SoftLayer_Metric_Tracking_Object_Virtual_Storage_Repository) String() string {
	return "SoftLayer_Metric_Tracking_Object_Virtual_Storage_Repository"
}

// SoftLayer_Metric_Tracking_Object_Virtual_Storage_Repository_Extended is SoftLayer_Metric_Tracking_Object_Virtual_Storage_Repository with all maskable types.
type SoftLayer_Metric_Tracking_Object_Virtual_Storage_Repository_Extended struct {
	SoftLayer_Metric_Tracking_Object_Virtual_Storage_Repository

	// Type - no documentation
	Type *SoftLayer_Metric_Tracking_Object_Type `json:"type,omitempty"`

	// Resource - The virtual storage repository that this tracking object tracks.
	Resource *SoftLayer_Virtual_Storage_Repository `json:"resource,omitempty"`
}

func (softlayer_metric_tracking_object_virtual_storage_repository *SoftLayer_Metric_Tracking_Object_Virtual_Storage_Repository_Extended) String() string {
	return "SoftLayer_Metric_Tracking_Object_Virtual_Storage_Repository"
}
