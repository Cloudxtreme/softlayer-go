package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Package_Preset_Configuration - <nil>
type SoftLayer_Product_Package_Preset_Configuration struct {

	// PackagePreset - <nil>
	PackagePreset *SoftLayer_Product_Package_Preset `json:"packagePreset,omitempty"`

	// Price - <nil>
	Price *SoftLayer_Product_Item_Price `json:"price,omitempty"`

	// Category - <nil>
	Category *SoftLayer_Product_Item_Category `json:"category,omitempty"`
}

func (softlayer_product_package_preset_configuration *SoftLayer_Product_Package_Preset_Configuration) String() string {
	return "SoftLayer_Product_Package_Preset_Configuration"
}
