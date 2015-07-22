package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Location_Reservation_Rack - <nil>
type SoftLayer_Location_Reservation_Rack struct {

	// Allotment - The bandwidth allotment that the reservation belongs to.
	Allotment *SoftLayer_Network_Bandwidth_Version1_Allotment `json:"allotment"`

	// Children - no documentation
	Children []*SoftLayer_Location_Reservation_Rack_Member `json:"children"`

	// ChildrenCount - no documentation
	ChildrenCount uint64 `json:"childrenCount"`

	// Location - <nil>
	Location *SoftLayer_Location `json:"location"`

	// LocationId - <nil>
	LocationId int `json:"locationId"`

	// LocationReservation - <nil>
	LocationReservation *SoftLayer_Location_Reservation `json:"locationReservation"`

	// LocationReservationId - <nil>
	LocationReservationId int `json:"locationReservationId"`

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

	// SlotReservation - <nil>
	SlotReservation int `json:"slotReservation"`
}

// GetObject - <nil>
func (softlayer_location_reservation_rack *SoftLayer_Location_Reservation_Rack) GetObject() (*SoftLayer_Location_Reservation_Rack, error) {
	var returnValue *SoftLayer_Location_Reservation_Rack
	return returnValue, nil
}
