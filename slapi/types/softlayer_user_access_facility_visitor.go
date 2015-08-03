package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_User_Access_Facility_Visitor - This class represents a facility visitor that is not an
// active employee or customer.
type SoftLayer_User_Access_Facility_Visitor struct {

	// LastName - <nil>
	LastName string `json:"lastName"`

	// TypeId - <nil>
	TypeId int `json:"typeId"`

	// CompanyName - <nil>
	CompanyName string `json:"companyName"`

	// FirstName - <nil>
	FirstName string `json:"firstName"`
}

// SoftLayer_User_Access_Facility_Visitor_Extended is SoftLayer_User_Access_Facility_Visitor with all maskable types.
type SoftLayer_User_Access_Facility_Visitor_Extended struct {
	SoftLayer_User_Access_Facility_Visitor

	// VisitorType - <nil>
	VisitorType *SoftLayer_User_Access_Facility_Visitor_Type `json:"visitorType"`
}

func (softlayer_user_access_facility_visitor *SoftLayer_User_Access_Facility_Visitor) String() string {
	return "SoftLayer_User_Access_Facility_Visitor"
}
