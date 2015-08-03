package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Package_Type - The [[SoftLayer_Product_Package_Type]] object indicates the type
// for a service offering (package). The type can be used to filter packages. For example, if you are
// looking for the package representing virtual servers, you can filter on the type's key name of For
// bare metal servers by core or filter on or respectively.
type SoftLayer_Product_Package_Type struct {

	// Id - no documentation
	Id int `json:"id"`

	// KeyName - The unique key name of the package type. Use this value when filtering.
	KeyName string `json:"keyName"`

	// Name - no documentation
	Name string `json:"name"`
}

func (softlayer_product_package_type *SoftLayer_Product_Package_Type) String() string {
	return "SoftLayer_Product_Package_Type"
}

// SoftLayer_Product_Package_Type_Extended is SoftLayer_Product_Package_Type with all maskable types.
type SoftLayer_Product_Package_Type_Extended struct {
	SoftLayer_Product_Package_Type

	// PackageCount - A count of all the packages associated with the given package type.
	PackageCount uint64 `json:"packageCount"`

	// Packages - All the packages associated with the given package type.
	Packages []*SoftLayer_Product_Package `json:"packages"`
}

func (softlayer_product_package_type *SoftLayer_Product_Package_Type_Extended) String() string {
	return "SoftLayer_Product_Package_Type"
}
