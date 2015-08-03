package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Package_Locations - Most packages are available in many locations. This object
// describes that availability for each package.
type SoftLayer_Product_Package_Locations struct {

	// DeliveryTimeInformation - This describes the availability of the package tied to this location.
	DeliveryTimeInformation string `json:"deliveryTimeInformation"`

	// IsAvailable - A simple flag which describes whether or not this location is available for this
	// package.
	IsAvailable int `json:"isAvailable"`

	// Location - no documentation
	Location *SoftLayer_Location `json:"location"`

	// LocationId - no documentation
	LocationId int `json:"locationId"`

	// Package - no documentation
	Package *SoftLayer_Product_Package `json:"package"`

	// PackageId - The SoftLayer_Product_Package ID tied to this object.
	PackageId int `json:"packageId"`
}

func (softlayer_product_package_locations *SoftLayer_Product_Package_Locations) String() string {
	return "SoftLayer_Product_Package_Locations"
}
