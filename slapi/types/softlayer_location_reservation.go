package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Location_Reservation - <nil>
type SoftLayer_Location_Reservation struct {

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// LocationId - <nil>
	LocationId int `json:"locationId,omitempty"`

	// AllotmentId - <nil>
	AllotmentId int `json:"allotmentId,omitempty"`

	// Name - <nil>
	Name string `json:"name,omitempty"`

	// Notes - <nil>
	Notes string `json:"notes,omitempty"`

	// Location - The datacenter location that the reservation belongs to.
	Location *SoftLayer_Location `json:"location,omitempty"`

	// LocationReservationRack - no documentation
	LocationReservationRack *SoftLayer_Location_Reservation_Rack `json:"locationReservationRack,omitempty"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account,omitempty"`

	// Allotment - The bandwidth allotment that the reservation belongs to.
	Allotment *SoftLayer_Network_Bandwidth_Version1_Allotment `json:"allotment,omitempty"`

	// BillingItem - The bandwidth allotment that the reservation belongs to.
	BillingItem *SoftLayer_Billing_Item `json:"billingItem,omitempty"`
}

func (softlayer_location_reservation *SoftLayer_Location_Reservation) String() string {
	return "SoftLayer_Location_Reservation"
}
