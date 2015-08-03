package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

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

func (softlayer_location_reservation_rack_member *SoftLayer_Location_Reservation_Rack_Member) String() string {
	return "SoftLayer_Location_Reservation_Rack_Member"
}
