package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Location_Region - A region is made up of a keyname and a description of that region. A
// region keyname can be used as part of an order. Check the SoftLayer_Product_Order service for more
// details.
type SoftLayer_Location_Region struct {

	// Description - a short description of a region's name. This description is seen on the order forms.
	Description string `json:"description"`

	// Keyname - A unique key name for a region. Provided for easy debugging. This is to be sent in with an
	// order.
	Keyname string `json:"keyname"`

	// Location - Each region can have many datacenter locations tied to it. However, this is the location
	// we currently provision to for a region. This location is the current valid location for a region.
	Location *SoftLayer_Location_Region_Location `json:"location"`

	// SortOrder - An integer representing the order in which this element is displayed.
	SortOrder int `json:"sortOrder"`
}

func (softlayer_location_region *SoftLayer_Location_Region) String() string {
	return "SoftLayer_Location_Region"
}
