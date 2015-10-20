package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Shipment - The SoftLayer_Account_Shipment data type contains information relating
// to a shipment. Basic information such as addresses, the shipment courier, and any tracking
// information for as shipment is accessible with this data type.
type SoftLayer_Account_Shipment struct {

	// OriginationDate - no documentation
	OriginationDate *time.Time `json:"originationDate,omitempty"`

	// DestinationDate - no documentation
	DestinationDate *time.Time `json:"destinationDate,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// CourierId - no documentation
	CourierId int `json:"courierId,omitempty"`

	// CourierName - no documentation
	CourierName string `json:"courierName,omitempty"`

	// ModifyUserId - no documentation
	ModifyUserId int `json:"modifyUserId,omitempty"`

	// Note - no documentation
	Note string `json:"note,omitempty"`

	// StatusId - no documentation
	StatusId int `json:"statusId,omitempty"`

	// DestinationAddressId - no documentation
	DestinationAddressId int `json:"destinationAddressId,omitempty"`

	// OriginationAddressId - no documentation
	OriginationAddressId int `json:"originationAddressId,omitempty"`

	// TypeId - no documentation
	TypeId int `json:"typeId,omitempty"`

	// AccountId - no documentation
	AccountId int `json:"accountId,omitempty"`

	// CreateUserId - no documentation
	CreateUserId int `json:"createUserId,omitempty"`

	// CreateUser - no documentation
	CreateUser *SoftLayer_User_Customer `json:"createUser,omitempty"`

	// ShipmentItems - no documentation
	ShipmentItems []*SoftLayer_Account_Shipment_Item `json:"shipmentItems,omitempty"`

	// ShipmentItemCount - no documentation
	ShipmentItemCount uint64 `json:"shipmentItemCount,omitempty"`

	// Status - no documentation
	Status *SoftLayer_Account_Shipment_Status `json:"status,omitempty"`

	// ModifyUser - no documentation
	ModifyUser *SoftLayer_User_Customer `json:"modifyUser,omitempty"`

	// TrackingData - no documentation
	TrackingData []*SoftLayer_Account_Shipment_Tracking_Data `json:"trackingData,omitempty"`

	// Type - The type of shipment (e.g. for Data Transfer Service or Colocation Service).
	Type *SoftLayer_Account_Shipment_Type `json:"type,omitempty"`

	// TrackingDataCount - no documentation
	TrackingDataCount uint64 `json:"trackingDataCount,omitempty"`

	// Courier - no documentation
	Courier *SoftLayer_Auxiliary_Shipping_Courier `json:"courier,omitempty"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account,omitempty"`

	// DestinationAddress - no documentation
	DestinationAddress *SoftLayer_Account_Address `json:"destinationAddress,omitempty"`

	// ModifyEmployee - no documentation
	ModifyEmployee *SoftLayer_User_Employee `json:"modifyEmployee,omitempty"`

	// CreateEmployee - no documentation
	CreateEmployee *SoftLayer_User_Employee `json:"createEmployee,omitempty"`

	// OriginationAddress - no documentation
	OriginationAddress *SoftLayer_Account_Address `json:"originationAddress,omitempty"`
}

func (softlayer_account_shipment *SoftLayer_Account_Shipment) String() string {
	return "SoftLayer_Account_Shipment"
}
