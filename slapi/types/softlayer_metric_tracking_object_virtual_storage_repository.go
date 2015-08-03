package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Metric_Tracking_Object_Virtual_Storage_Repository - <nil>
type SoftLayer_Metric_Tracking_Object_Virtual_Storage_Repository struct {
}

func (softlayer_metric_tracking_object_virtual_storage_repository *SoftLayer_Metric_Tracking_Object_Virtual_Storage_Repository) String() string {
	return "SoftLayer_Metric_Tracking_Object_Virtual_Storage_Repository"
}

// SoftLayer_Metric_Tracking_Object_Virtual_Storage_Repository_Extended is SoftLayer_Metric_Tracking_Object_Virtual_Storage_Repository with all maskable types.
type SoftLayer_Metric_Tracking_Object_Virtual_Storage_Repository_Extended struct {
	SoftLayer_Metric_Tracking_Object_Virtual_Storage_Repository

	// Resource - The virtual storage repository that this tracking object tracks.
	Resource *SoftLayer_Virtual_Storage_Repository `json:"resource"`
}

func (softlayer_metric_tracking_object_virtual_storage_repository *SoftLayer_Metric_Tracking_Object_Virtual_Storage_Repository_Extended) String() string {
	return "SoftLayer_Metric_Tracking_Object_Virtual_Storage_Repository"
}
