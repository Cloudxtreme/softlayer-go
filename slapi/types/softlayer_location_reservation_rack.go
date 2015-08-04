package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Location_Reservation_Rack - <nil>
type SoftLayer_Location_Reservation_Rack struct {

	// LocationReservationId - <nil>
	LocationReservationId int `json:"locationReservationId,omitempty"`

	// NetworkConnectionCapacity - <nil>
	NetworkConnectionCapacity int `json:"networkConnectionCapacity,omitempty"`

	// NetworkConnectionReservation - <nil>
	NetworkConnectionReservation int `json:"networkConnectionReservation,omitempty"`

	// SlotCapacity - <nil>
	SlotCapacity int `json:"slotCapacity,omitempty"`

	// LocationId - <nil>
	LocationId int `json:"locationId,omitempty"`

	// PowerConnectionCapacity - <nil>
	PowerConnectionCapacity int `json:"powerConnectionCapacity,omitempty"`

	// PowerConnectionReservation - <nil>
	PowerConnectionReservation int `json:"powerConnectionReservation,omitempty"`

	// SlotReservation - <nil>
	SlotReservation int `json:"slotReservation,omitempty"`
}

func (softlayer_location_reservation_rack *SoftLayer_Location_Reservation_Rack) String() string {
	return "SoftLayer_Location_Reservation_Rack"
}

// SoftLayer_Location_Reservation_Rack_Extended is SoftLayer_Location_Reservation_Rack with all maskable types.
type SoftLayer_Location_Reservation_Rack_Extended struct {
	SoftLayer_Location_Reservation_Rack

	// LocationReservation - <nil>
	LocationReservation *SoftLayer_Location_Reservation `json:"locationReservation,omitempty"`

	// ChildrenCount - no documentation
	ChildrenCount uint64 `json:"childrenCount,omitempty"`

	// Allotment - The bandwidth allotment that the reservation belongs to.
	Allotment *SoftLayer_Network_Bandwidth_Version1_Allotment `json:"allotment,omitempty"`

	// Children - no documentation
	Children []*SoftLayer_Location_Reservation_Rack_Member `json:"children,omitempty"`

	// Location - <nil>
	Location *SoftLayer_Location `json:"location,omitempty"`
}

func (softlayer_location_reservation_rack *SoftLayer_Location_Reservation_Rack_Extended) String() string {
	return "SoftLayer_Location_Reservation_Rack"
}
