package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Account_Shipment_Tracking_Data - The SoftLayer_Account_Shipment_Tracking_Data data type
// contains information on a single piece of tracking information pertaining to a shipment. This
// tracking information tracking numbers by which the shipment may be tracked through the shipping
// courier.
type SoftLayer_Account_Shipment_Tracking_Data struct {

	// CreateEmployee - no documentation
	CreateEmployee *SoftLayer_User_Employee `json:"createEmployee"`

	// CreateUser - no documentation
	CreateUser *SoftLayer_User_Customer `json:"createUser"`

	// CreateUserId - no documentation
	CreateUserId int `json:"createUserId"`

	// Id - no documentation
	Id int `json:"id"`

	// ModifyEmployee - no documentation
	ModifyEmployee *SoftLayer_User_Employee `json:"modifyEmployee"`

	// ModifyUser - The customer user who last modified the tracking datum.
	ModifyUser *SoftLayer_User_Customer `json:"modifyUser"`

	// ModifyUserId - no documentation
	ModifyUserId int `json:"modifyUserId"`

	// PackageId - no documentation
	PackageId int `json:"packageId"`

	// Sequence - no documentation
	Sequence int `json:"sequence"`

	// Shipment - no documentation
	Shipment *SoftLayer_Account_Shipment `json:"shipment"`

	// ShipmentId - no documentation
	ShipmentId int `json:"shipmentId"`

	// TrackingData - The tracking data (tracking number/reference number).
	TrackingData string `json:"trackingData"`
}

func (softlayer_account_shipment_tracking_data *SoftLayer_Account_Shipment_Tracking_Data) String() string {
	return "SoftLayer_Account_Shipment_Tracking_Data"
}

// CreateObject - Create a new shipment tracking data. The ''shipmentId'', ''sequence'', and
// ''trackingData'' properties in the templateObject parameter are required parameters to create a
// tracking data record.
func (softlayer_account_shipment_tracking_data *SoftLayer_Account_Shipment_Tracking_Data) CreateObject(ctx *slapi.RequestContext, templateObject SoftLayer_Account_Shipment_Tracking_Data) (*SoftLayer_Account_Shipment_Tracking_Data, error) {
	var returnValue *SoftLayer_Account_Shipment_Tracking_Data
	return returnValue, nil
}

// CreateObjects - Create a new shipment tracking data. The ''shipmentId'', ''sequence'', and
// ''trackingData'' properties of each templateObject in the templateObjects array are required
// parameters to create a tracking data record.
func (softlayer_account_shipment_tracking_data *SoftLayer_Account_Shipment_Tracking_Data) CreateObjects(ctx *slapi.RequestContext, templateObjects []SoftLayer_Account_Shipment_Tracking_Data) ([]*SoftLayer_Account_Shipment_Tracking_Data, error) {
	var returnValue []*SoftLayer_Account_Shipment_Tracking_Data
	return returnValue, nil
}

// DeleteObject - deleteObject permanently removes a shipment tracking datum (number)
func (softlayer_account_shipment_tracking_data *SoftLayer_Account_Shipment_Tracking_Data) DeleteObject(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObject - Edit the properties of a tracking data record by passing in a modified instance of a
// SoftLayer_Account_Shipment_Tracking_Data object.
func (softlayer_account_shipment_tracking_data *SoftLayer_Account_Shipment_Tracking_Data) EditObject(ctx *slapi.RequestContext, templateObject SoftLayer_Account_Shipment_Tracking_Data) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_account_shipment_tracking_data *SoftLayer_Account_Shipment_Tracking_Data) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Account_Shipment_Tracking_Data, error) {
	var returnValue *SoftLayer_Account_Shipment_Tracking_Data
	return returnValue, nil
}
