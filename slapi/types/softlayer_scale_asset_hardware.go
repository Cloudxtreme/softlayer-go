package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Scale_Asset_Hardware - <nil>
type SoftLayer_Scale_Asset_Hardware struct {
}

// SoftLayer_Scale_Asset_Hardware_Extended is SoftLayer_Scale_Asset_Hardware with all maskable types.
type SoftLayer_Scale_Asset_Hardware_Extended struct {
	SoftLayer_Scale_Asset_Hardware

	// Hardware - no documentation
	Hardware *SoftLayer_Hardware `json:"hardware"`

	// HardwareId - no documentation
	HardwareId int `json:"hardwareId"`
}

func (softlayer_scale_asset_hardware *SoftLayer_Scale_Asset_Hardware) String() string {
	return "SoftLayer_Scale_Asset_Hardware"
}
