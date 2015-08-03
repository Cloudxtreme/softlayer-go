package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Account_Shipment - The SoftLayer_Account_Shipment data type contains information relating
// to a shipment. Basic information such as addresses, the shipment courier, and any tracking
// information for as shipment is accessible with this data type.
type SoftLayer_Account_Shipment struct {

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// AccountId - no documentation
	AccountId int `json:"accountId"`

	// Courier - no documentation
	Courier *SoftLayer_Auxiliary_Shipping_Courier `json:"courier"`

	// CourierId - no documentation
	CourierId int `json:"courierId"`

	// CourierName - no documentation
	CourierName string `json:"courierName"`

	// CreateEmployee - no documentation
	CreateEmployee *SoftLayer_User_Employee `json:"createEmployee"`

	// CreateUser - no documentation
	CreateUser *SoftLayer_User_Customer `json:"createUser"`

	// CreateUserId - no documentation
	CreateUserId int `json:"createUserId"`

	// DestinationAddress - no documentation
	DestinationAddress *SoftLayer_Account_Address `json:"destinationAddress"`

	// DestinationAddressId - no documentation
	DestinationAddressId int `json:"destinationAddressId"`

	// DestinationDate - no documentation
	DestinationDate *time.Time `json:"destinationDate"`

	// Id - no documentation
	Id int `json:"id"`

	// ModifyEmployee - no documentation
	ModifyEmployee *SoftLayer_User_Employee `json:"modifyEmployee"`

	// ModifyUser - no documentation
	ModifyUser *SoftLayer_User_Customer `json:"modifyUser"`

	// ModifyUserId - no documentation
	ModifyUserId int `json:"modifyUserId"`

	// Note - no documentation
	Note string `json:"note"`

	// OriginationAddress - no documentation
	OriginationAddress *SoftLayer_Account_Address `json:"originationAddress"`

	// OriginationAddressId - no documentation
	OriginationAddressId int `json:"originationAddressId"`

	// OriginationDate - no documentation
	OriginationDate *time.Time `json:"originationDate"`

	// ShipmentItemCount - no documentation
	ShipmentItemCount uint64 `json:"shipmentItemCount"`

	// ShipmentItems - no documentation
	ShipmentItems []*SoftLayer_Account_Shipment_Item `json:"shipmentItems"`

	// Status - no documentation
	Status *SoftLayer_Account_Shipment_Status `json:"status"`

	// StatusId - no documentation
	StatusId int `json:"statusId"`

	// TrackingData - no documentation
	TrackingData []*SoftLayer_Account_Shipment_Tracking_Data `json:"trackingData"`

	// TrackingDataCount - no documentation
	TrackingDataCount uint64 `json:"trackingDataCount"`

	// Type - The type of shipment (e.g. for Data Transfer Service or Colocation Service).
	Type *SoftLayer_Account_Shipment_Type `json:"type"`

	// TypeId - no documentation
	TypeId int `json:"typeId"`
}

func (softlayer_account_shipment *SoftLayer_Account_Shipment) String() string {
	return "SoftLayer_Account_Shipment"
}

// EditObject - Edit the properties of a shipment record by passing in a modified instance of a
// SoftLayer_Account_Shipment object.
func (softlayer_account_shipment *SoftLayer_Account_Shipment) EditObject(ctx *slapi.RequestContext, templateObject SoftLayer_Account_Shipment) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetAllCouriers - no documentation
func (softlayer_account_shipment *SoftLayer_Account_Shipment) GetAllCouriers(ctx *slapi.RequestContext) ([]*SoftLayer_Auxiliary_Shipping_Courier, error) {
	var returnValue []*SoftLayer_Auxiliary_Shipping_Courier
	return returnValue, nil
}

// GetAllCouriersByType - no documentation
func (softlayer_account_shipment *SoftLayer_Account_Shipment) GetAllCouriersByType(ctx *slapi.RequestContext, courierTypeKeyName string) ([]*SoftLayer_Auxiliary_Shipping_Courier, error) {
	var returnValue []*SoftLayer_Auxiliary_Shipping_Courier
	return returnValue, nil
}

// GetAllShipmentStatuses - no documentation
func (softlayer_account_shipment *SoftLayer_Account_Shipment) GetAllShipmentStatuses(ctx *slapi.RequestContext) ([]*SoftLayer_Account_Shipment_Status, error) {
	var returnValue []*SoftLayer_Account_Shipment_Status
	return returnValue, nil
}

// GetAllShipmentTypes - no documentation
func (softlayer_account_shipment *SoftLayer_Account_Shipment) GetAllShipmentTypes(ctx *slapi.RequestContext) ([]*SoftLayer_Account_Shipment_Type, error) {
	var returnValue []*SoftLayer_Account_Shipment_Type
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_account_shipment *SoftLayer_Account_Shipment) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Account_Shipment, error) {
	var returnValue *SoftLayer_Account_Shipment
	return returnValue, nil
}
