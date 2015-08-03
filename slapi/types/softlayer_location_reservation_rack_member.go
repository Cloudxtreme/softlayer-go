package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Location_Reservation_Rack_Member - <nil>
type SoftLayer_Location_Reservation_Rack_Member struct {

	// Id - <nil>
	Id int `json:"id"`

	// LocationId - <nil>
	LocationId int `json:"locationId"`
}

func (softlayer_location_reservation_rack_member *SoftLayer_Location_Reservation_Rack_Member) String() string {
	return "SoftLayer_Location_Reservation_Rack_Member"
}

// SoftLayer_Location_Reservation_Rack_Member_Extended is SoftLayer_Location_Reservation_Rack_Member with all maskable types.
type SoftLayer_Location_Reservation_Rack_Member_Extended struct {
	SoftLayer_Location_Reservation_Rack_Member

	// LocationReservationRack - <nil>
	LocationReservationRack *SoftLayer_Location_Reservation `json:"locationReservationRack"`

	// Location - no documentation
	Location *SoftLayer_Location `json:"location"`
}

func (softlayer_location_reservation_rack_member *SoftLayer_Location_Reservation_Rack_Member_Extended) String() string {
	return "SoftLayer_Location_Reservation_Rack_Member"
}
