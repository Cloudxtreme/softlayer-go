package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Location_Reservation - <nil>
type SoftLayer_Location_Reservation struct {

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// Allotment - The bandwidth allotment that the reservation belongs to.
	Allotment *SoftLayer_Network_Bandwidth_Version1_Allotment `json:"allotment"`

	// AllotmentId - <nil>
	AllotmentId int `json:"allotmentId"`

	// BillingItem - The bandwidth allotment that the reservation belongs to.
	BillingItem *SoftLayer_Billing_Item `json:"billingItem"`

	// Id - <nil>
	Id int `json:"id"`

	// Location - The datacenter location that the reservation belongs to.
	Location *SoftLayer_Location `json:"location"`

	// LocationId - <nil>
	LocationId int `json:"locationId"`

	// LocationReservationRack - no documentation
	LocationReservationRack *SoftLayer_Location_Reservation_Rack `json:"locationReservationRack"`

	// Name - <nil>
	Name string `json:"name"`

	// Notes - <nil>
	Notes string `json:"notes"`
}

// GetAccountReservations - <nil>
func (softlayer_location_reservation *SoftLayer_Location_Reservation) GetAccountReservations(ctx *slapi.RequestContext) ([]*SoftLayer_Location_Reservation, error) {
	var returnValue []*SoftLayer_Location_Reservation
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_location_reservation *SoftLayer_Location_Reservation) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Location_Reservation, error) {
	var returnValue *SoftLayer_Location_Reservation
	return returnValue, nil
}
