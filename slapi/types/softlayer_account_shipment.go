package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Shipment - The SoftLayer_Account_Shipment data type contains information relating
// to a shipment. Basic information such as addresses, the shipment courier, and any tracking
// information for as shipment is accessible with this data type.
type SoftLayer_Account_Shipment struct {

	// OriginationAddressId - no documentation
	OriginationAddressId int `json:"originationAddressId"`

	// StatusId - no documentation
	StatusId int `json:"statusId"`

	// AccountId - no documentation
	AccountId int `json:"accountId"`

	// OriginationDate - no documentation
	OriginationDate *time.Time `json:"originationDate"`

	// Id - no documentation
	Id int `json:"id"`

	// TypeId - no documentation
	TypeId int `json:"typeId"`

	// DestinationAddressId - no documentation
	DestinationAddressId int `json:"destinationAddressId"`

	// CourierName - no documentation
	CourierName string `json:"courierName"`

	// ModifyUserId - no documentation
	ModifyUserId int `json:"modifyUserId"`

	// CourierId - no documentation
	CourierId int `json:"courierId"`

	// CreateUserId - no documentation
	CreateUserId int `json:"createUserId"`

	// DestinationDate - no documentation
	DestinationDate *time.Time `json:"destinationDate"`

	// Note - no documentation
	Note string `json:"note"`
}

func (softlayer_account_shipment *SoftLayer_Account_Shipment) String() string {
	return "SoftLayer_Account_Shipment"
}

// SoftLayer_Account_Shipment_Extended is SoftLayer_Account_Shipment with all maskable types.
type SoftLayer_Account_Shipment_Extended struct {
	SoftLayer_Account_Shipment

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// DestinationAddress - no documentation
	DestinationAddress *SoftLayer_Account_Address `json:"destinationAddress"`

	// CreateUser - no documentation
	CreateUser *SoftLayer_User_Customer `json:"createUser"`

	// OriginationAddress - no documentation
	OriginationAddress *SoftLayer_Account_Address `json:"originationAddress"`

	// TrackingDataCount - no documentation
	TrackingDataCount uint64 `json:"trackingDataCount"`

	// Status - no documentation
	Status *SoftLayer_Account_Shipment_Status `json:"status"`

	// TrackingData - no documentation
	TrackingData []*SoftLayer_Account_Shipment_Tracking_Data `json:"trackingData"`

	// ShipmentItemCount - no documentation
	ShipmentItemCount uint64 `json:"shipmentItemCount"`

	// Courier - no documentation
	Courier *SoftLayer_Auxiliary_Shipping_Courier `json:"courier"`

	// ModifyEmployee - no documentation
	ModifyEmployee *SoftLayer_User_Employee `json:"modifyEmployee"`

	// ModifyUser - no documentation
	ModifyUser *SoftLayer_User_Customer `json:"modifyUser"`

	// Type - The type of shipment (e.g. for Data Transfer Service or Colocation Service).
	Type *SoftLayer_Account_Shipment_Type `json:"type"`

	// CreateEmployee - no documentation
	CreateEmployee *SoftLayer_User_Employee `json:"createEmployee"`

	// ShipmentItems - no documentation
	ShipmentItems []*SoftLayer_Account_Shipment_Item `json:"shipmentItems"`
}

func (softlayer_account_shipment *SoftLayer_Account_Shipment_Extended) String() string {
	return "SoftLayer_Account_Shipment"
}
