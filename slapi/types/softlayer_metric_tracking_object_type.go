package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Metric_Tracking_Object_Type - SoftLayer [[SoftLayer_Metric_Tracking_Object|tracking
// objects]] can model various kinds of measured data, from server and virtual server bandwidth usage
// to CPU use to remote storage usage. SoftLayer_Metric_Tracking_Object_Type models one of these types
// and is referred to in tracking objects to reflect what type of data they track.
type SoftLayer_Metric_Tracking_Object_Type struct {

	// Keyname - Description A tracking object type's key name. This is a shorter description of what kind
	// of data a tracking object group is polling.
	Keyname string `json:"keyname"`

	// Name - Description A tracking object type's name. This describes what kind of data a tracking object
	// group is polling.
	Name string `json:"name"`
}

func (softlayer_metric_tracking_object_type *SoftLayer_Metric_Tracking_Object_Type) String() string {
	return "SoftLayer_Metric_Tracking_Object_Type"
}
