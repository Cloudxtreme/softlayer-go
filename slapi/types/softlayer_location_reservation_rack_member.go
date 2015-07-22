package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Location_Reservation_Rack_Member - <nil>
type SoftLayer_Location_Reservation_Rack_Member struct {

	// Id - <nil>
	Id int `json:"id"`

	// Location - no documentation
	Location *SoftLayer_Location `json:"location"`

	// LocationId - <nil>
	LocationId int `json:"locationId"`

	// LocationReservationRack - <nil>
	LocationReservationRack *SoftLayer_Location_Reservation `json:"locationReservationRack"`
}

// GetObject - <nil>
func (softlayer_location_reservation_rack_member *SoftLayer_Location_Reservation_Rack_Member) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Location_Reservation_Rack_Member, error) {
	var returnValue *SoftLayer_Location_Reservation_Rack_Member
	return returnValue, nil
}
