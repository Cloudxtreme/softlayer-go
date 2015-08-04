package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Shipment - The SoftLayer_Account_Shipment data type contains information relating
// to a shipment. Basic information such as addresses, the shipment courier, and any tracking
// information for as shipment is accessible with this data type.
type SoftLayer_Account_Shipment struct {

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// OriginationDate - no documentation
	OriginationDate *time.Time `json:"originationDate,omitempty"`

	// ModifyUserId - no documentation
	ModifyUserId int `json:"modifyUserId,omitempty"`

	// StatusId - no documentation
	StatusId int `json:"statusId,omitempty"`

	// OriginationAddressId - no documentation
	OriginationAddressId int `json:"originationAddressId,omitempty"`

	// CourierName - no documentation
	CourierName string `json:"courierName,omitempty"`

	// CreateUserId - no documentation
	CreateUserId int `json:"createUserId,omitempty"`

	// DestinationDate - no documentation
	DestinationDate *time.Time `json:"destinationDate,omitempty"`

	// Note - no documentation
	Note string `json:"note,omitempty"`

	// AccountId - no documentation
	AccountId int `json:"accountId,omitempty"`

	// CourierId - no documentation
	CourierId int `json:"courierId,omitempty"`

	// DestinationAddressId - no documentation
	DestinationAddressId int `json:"destinationAddressId,omitempty"`

	// TypeId - no documentation
	TypeId int `json:"typeId,omitempty"`

	// DestinationAddress - no documentation
	DestinationAddress *SoftLayer_Account_Address `json:"destinationAddress,omitempty"`

	// CreateUser - no documentation
	CreateUser *SoftLayer_User_Customer `json:"createUser,omitempty"`

	// ModifyUser - no documentation
	ModifyUser *SoftLayer_User_Customer `json:"modifyUser,omitempty"`

	// TrackingData - no documentation
	TrackingData []*SoftLayer_Account_Shipment_Tracking_Data `json:"trackingData,omitempty"`

	// Type - The type of shipment (e.g. for Data Transfer Service or Colocation Service).
	Type *SoftLayer_Account_Shipment_Type `json:"type,omitempty"`

	// ShipmentItemCount - no documentation
	ShipmentItemCount uint64 `json:"shipmentItemCount,omitempty"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account,omitempty"`

	// ShipmentItems - no documentation
	ShipmentItems []*SoftLayer_Account_Shipment_Item `json:"shipmentItems,omitempty"`

	// Courier - no documentation
	Courier *SoftLayer_Auxiliary_Shipping_Courier `json:"courier,omitempty"`

	// TrackingDataCount - no documentation
	TrackingDataCount uint64 `json:"trackingDataCount,omitempty"`

	// CreateEmployee - no documentation
	CreateEmployee *SoftLayer_User_Employee `json:"createEmployee,omitempty"`

	// ModifyEmployee - no documentation
	ModifyEmployee *SoftLayer_User_Employee `json:"modifyEmployee,omitempty"`

	// OriginationAddress - no documentation
	OriginationAddress *SoftLayer_Account_Address `json:"originationAddress,omitempty"`

	// Status - no documentation
	Status *SoftLayer_Account_Shipment_Status `json:"status,omitempty"`
}

func (softlayer_account_shipment *SoftLayer_Account_Shipment) String() string {
	return "SoftLayer_Account_Shipment"
}
