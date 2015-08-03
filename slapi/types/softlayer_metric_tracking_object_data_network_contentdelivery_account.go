package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Metric_Tracking_Object_Data_Network_ContentDelivery_Account -
// SoftLayer_Metric_Tracking_Object_Data_Network_ContentDelivery_Account models usage data polled from
// the CDN system.
type SoftLayer_Metric_Tracking_Object_Data_Network_ContentDelivery_Account struct {

	// FileName - The name of a file. This value is only populated in file-based bandwidth reports.
	FileName string `json:"fileName"`

	// PopId - The internal identifier of a CDN POP (Points of Presence).
	PopId int `json:"popId"`
}

func (softlayer_metric_tracking_object_data_network_contentdelivery_account *SoftLayer_Metric_Tracking_Object_Data_Network_ContentDelivery_Account) String() string {
	return "SoftLayer_Metric_Tracking_Object_Data_Network_ContentDelivery_Account"
}
