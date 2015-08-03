package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Location_Reservation_Rack - <nil>
type SoftLayer_Location_Reservation_Rack struct {

	// LocationReservationId - <nil>
	LocationReservationId int `json:"locationReservationId"`

	// SlotReservation - <nil>
	SlotReservation int `json:"slotReservation"`

	// LocationId - <nil>
	LocationId int `json:"locationId"`

	// NetworkConnectionCapacity - <nil>
	NetworkConnectionCapacity int `json:"networkConnectionCapacity"`

	// NetworkConnectionReservation - <nil>
	NetworkConnectionReservation int `json:"networkConnectionReservation"`

	// PowerConnectionCapacity - <nil>
	PowerConnectionCapacity int `json:"powerConnectionCapacity"`

	// PowerConnectionReservation - <nil>
	PowerConnectionReservation int `json:"powerConnectionReservation"`

	// SlotCapacity - <nil>
	SlotCapacity int `json:"slotCapacity"`
}

// SoftLayer_Location_Reservation_Rack_Extended is SoftLayer_Location_Reservation_Rack with all maskable types.
type SoftLayer_Location_Reservation_Rack_Extended struct {
	SoftLayer_Location_Reservation_Rack

	// Allotment - The bandwidth allotment that the reservation belongs to.
	Allotment *SoftLayer_Network_Bandwidth_Version1_Allotment `json:"allotment"`

	// Children - no documentation
	Children []*SoftLayer_Location_Reservation_Rack_Member `json:"children"`

	// Location - <nil>
	Location *SoftLayer_Location `json:"location"`

	// LocationReservation - <nil>
	LocationReservation *SoftLayer_Location_Reservation `json:"locationReservation"`

	// ChildrenCount - no documentation
	ChildrenCount uint64 `json:"childrenCount"`
}

func (softlayer_location_reservation_rack *SoftLayer_Location_Reservation_Rack) String() string {
	return "SoftLayer_Location_Reservation_Rack"
}
