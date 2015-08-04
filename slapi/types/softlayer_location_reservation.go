package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Location_Reservation - <nil>
type SoftLayer_Location_Reservation struct {

	// LocationId - <nil>
	LocationId int `json:"locationId,omitempty"`

	// Name - <nil>
	Name string `json:"name,omitempty"`

	// AllotmentId - <nil>
	AllotmentId int `json:"allotmentId,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// Notes - <nil>
	Notes string `json:"notes,omitempty"`
}

func (softlayer_location_reservation *SoftLayer_Location_Reservation) String() string {
	return "SoftLayer_Location_Reservation"
}

// SoftLayer_Location_Reservation_Extended is SoftLayer_Location_Reservation with all maskable types.
type SoftLayer_Location_Reservation_Extended struct {
	SoftLayer_Location_Reservation

	// Allotment - The bandwidth allotment that the reservation belongs to.
	Allotment *SoftLayer_Network_Bandwidth_Version1_Allotment `json:"allotment,omitempty"`

	// Location - The datacenter location that the reservation belongs to.
	Location *SoftLayer_Location `json:"location,omitempty"`

	// LocationReservationRack - no documentation
	LocationReservationRack *SoftLayer_Location_Reservation_Rack `json:"locationReservationRack,omitempty"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account,omitempty"`

	// BillingItem - The bandwidth allotment that the reservation belongs to.
	BillingItem *SoftLayer_Billing_Item `json:"billingItem,omitempty"`
}

func (softlayer_location_reservation *SoftLayer_Location_Reservation_Extended) String() string {
	return "SoftLayer_Location_Reservation"
}
