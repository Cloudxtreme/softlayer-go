package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Account_Shipment_Tracking_Data - The SoftLayer_Account_Shipment_Tracking_Data data type
// contains information on a single piece of tracking information pertaining to a shipment. This
// tracking information tracking numbers by which the shipment may be tracked through the shipping
// courier.
type SoftLayer_Account_Shipment_Tracking_Data struct {

	// CreateUserId - no documentation
	CreateUserId int `json:"createUserId,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// ShipmentId - no documentation
	ShipmentId int `json:"shipmentId,omitempty"`

	// ModifyUserId - no documentation
	ModifyUserId int `json:"modifyUserId,omitempty"`

	// PackageId - no documentation
	PackageId int `json:"packageId,omitempty"`

	// Sequence - no documentation
	Sequence int `json:"sequence,omitempty"`

	// TrackingData - The tracking data (tracking number/reference number).
	TrackingData string `json:"trackingData,omitempty"`
}

func (softlayer_account_shipment_tracking_data *SoftLayer_Account_Shipment_Tracking_Data) String() string {
	return "SoftLayer_Account_Shipment_Tracking_Data"
}

// SoftLayer_Account_Shipment_Tracking_Data_Extended is SoftLayer_Account_Shipment_Tracking_Data with all maskable types.
type SoftLayer_Account_Shipment_Tracking_Data_Extended struct {
	SoftLayer_Account_Shipment_Tracking_Data

	// CreateEmployee - no documentation
	CreateEmployee *SoftLayer_User_Employee `json:"createEmployee,omitempty"`

	// CreateUser - no documentation
	CreateUser *SoftLayer_User_Customer `json:"createUser,omitempty"`

	// ModifyEmployee - no documentation
	ModifyEmployee *SoftLayer_User_Employee `json:"modifyEmployee,omitempty"`

	// ModifyUser - The customer user who last modified the tracking datum.
	ModifyUser *SoftLayer_User_Customer `json:"modifyUser,omitempty"`

	// Shipment - no documentation
	Shipment *SoftLayer_Account_Shipment `json:"shipment,omitempty"`
}

func (softlayer_account_shipment_tracking_data *SoftLayer_Account_Shipment_Tracking_Data_Extended) String() string {
	return "SoftLayer_Account_Shipment_Tracking_Data"
}
