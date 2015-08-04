package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Location_Region_Location - The SoftLayer_Location_Region_Location is very specific to the
// location where services will actually be provisioned. When accessed through a package, this location
// is the top priority location for a region. All new servers and services are provisioned at this
// location. When a server is ordered and a region is selected, this is the location within that region
// where the server will actually exist and have software/services installed.
type SoftLayer_Location_Region_Location struct {
}

func (softlayer_location_region_location *SoftLayer_Location_Region_Location) String() string {
	return "SoftLayer_Location_Region_Location"
}

// SoftLayer_Location_Region_Location_Extended is SoftLayer_Location_Region_Location with all maskable types.
type SoftLayer_Location_Region_Location_Extended struct {
	SoftLayer_Location_Region_Location

	// Location - The SoftLayer_Location tied to a region's location. This provides more information about
	// the location, including specific datacenter information.
	Location *SoftLayer_Location `json:"location,omitempty"`

	// LocationPackageDetails - A region's location also has delivery information as well as other
	// information to be determined. For now, availability is provided and could weigh into the decision as
	// to where to decide to have a server provisioned.'
	LocationPackageDetails []*SoftLayer_Product_Package_Locations `json:"locationPackageDetails,omitempty"`

	// Region - no documentation
	Region *SoftLayer_Location_Region `json:"region,omitempty"`

	// LocationPackageDetailCount - A count of a region's location also has delivery information as well as
	// other information to be determined. For now, availability is provided and could weigh into the
	// decision as to where to decide to have a server provisioned.'
	LocationPackageDetailCount uint64 `json:"locationPackageDetailCount,omitempty"`
}

func (softlayer_location_region_location *SoftLayer_Location_Region_Location_Extended) String() string {
	return "SoftLayer_Location_Region_Location"
}
