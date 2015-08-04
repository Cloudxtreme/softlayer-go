package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Account_Shipment_Tracking_Data - The SoftLayer_Account_Shipment_Tracking_Data data type
// contains information on a single piece of tracking information pertaining to a shipment. This
// tracking information tracking numbers by which the shipment may be tracked through the shipping
// courier.
type SoftLayer_Account_Shipment_Tracking_Data struct {

	// ModifyUserId - no documentation
	ModifyUserId int `json:"modifyUserId,omitempty"`

	// PackageId - no documentation
	PackageId int `json:"packageId,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// Sequence - no documentation
	Sequence int `json:"sequence,omitempty"`

	// ShipmentId - no documentation
	ShipmentId int `json:"shipmentId,omitempty"`

	// TrackingData - The tracking data (tracking number/reference number).
	TrackingData string `json:"trackingData,omitempty"`

	// CreateUserId - no documentation
	CreateUserId int `json:"createUserId,omitempty"`

	// ModifyUser - The customer user who last modified the tracking datum.
	ModifyUser *SoftLayer_User_Customer `json:"modifyUser,omitempty"`

	// Shipment - no documentation
	Shipment *SoftLayer_Account_Shipment `json:"shipment,omitempty"`

	// CreateUser - no documentation
	CreateUser *SoftLayer_User_Customer `json:"createUser,omitempty"`

	// ModifyEmployee - no documentation
	ModifyEmployee *SoftLayer_User_Employee `json:"modifyEmployee,omitempty"`

	// CreateEmployee - no documentation
	CreateEmployee *SoftLayer_User_Employee `json:"createEmployee,omitempty"`
}

func (softlayer_account_shipment_tracking_data *SoftLayer_Account_Shipment_Tracking_Data) String() string {
	return "SoftLayer_Account_Shipment_Tracking_Data"
}
