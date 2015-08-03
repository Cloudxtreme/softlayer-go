package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Package_Preset_Configuration - <nil>
type SoftLayer_Product_Package_Preset_Configuration struct {
}

func (softlayer_product_package_preset_configuration *SoftLayer_Product_Package_Preset_Configuration) String() string {
	return "SoftLayer_Product_Package_Preset_Configuration"
}

// SoftLayer_Product_Package_Preset_Configuration_Extended is SoftLayer_Product_Package_Preset_Configuration with all maskable types.
type SoftLayer_Product_Package_Preset_Configuration_Extended struct {
	SoftLayer_Product_Package_Preset_Configuration

	// Category - <nil>
	Category *SoftLayer_Product_Item_Category `json:"category"`

	// PackagePreset - <nil>
	PackagePreset *SoftLayer_Product_Package_Preset `json:"packagePreset"`

	// Price - <nil>
	Price *SoftLayer_Product_Item_Price `json:"price"`
}

func (softlayer_product_package_preset_configuration *SoftLayer_Product_Package_Preset_Configuration_Extended) String() string {
	return "SoftLayer_Product_Package_Preset_Configuration"
}
