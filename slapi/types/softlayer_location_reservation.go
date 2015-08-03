package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Location_Reservation - <nil>
type SoftLayer_Location_Reservation struct {

	// Name - <nil>
	Name string `json:"name"`

	// Notes - <nil>
	Notes string `json:"notes"`

	// AllotmentId - <nil>
	AllotmentId int `json:"allotmentId"`

	// Id - <nil>
	Id int `json:"id"`

	// LocationId - <nil>
	LocationId int `json:"locationId"`
}

// SoftLayer_Location_Reservation_Extended is SoftLayer_Location_Reservation with all maskable types.
type SoftLayer_Location_Reservation_Extended struct {
	SoftLayer_Location_Reservation

	// LocationReservationRack - no documentation
	LocationReservationRack *SoftLayer_Location_Reservation_Rack `json:"locationReservationRack"`

	// Allotment - The bandwidth allotment that the reservation belongs to.
	Allotment *SoftLayer_Network_Bandwidth_Version1_Allotment `json:"allotment"`

	// BillingItem - The bandwidth allotment that the reservation belongs to.
	BillingItem *SoftLayer_Billing_Item `json:"billingItem"`

	// Location - The datacenter location that the reservation belongs to.
	Location *SoftLayer_Location `json:"location"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`
}

func (softlayer_location_reservation *SoftLayer_Location_Reservation) String() string {
	return "SoftLayer_Location_Reservation"
}
