package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Hardware_Power_Component - <nil>
type SoftLayer_Hardware_Power_Component struct {

	// Id - <nil>
	Id int `json:"id"`

	// HardwareId - <nil>
	HardwareId int `json:"hardwareId"`
}

func (softlayer_hardware_power_component *SoftLayer_Hardware_Power_Component) String() string {
	return "SoftLayer_Hardware_Power_Component"
}

// SoftLayer_Hardware_Power_Component_Extended is SoftLayer_Hardware_Power_Component with all maskable types.
type SoftLayer_Hardware_Power_Component_Extended struct {
	SoftLayer_Hardware_Power_Component

	// Hardware - <nil>
	Hardware *SoftLayer_Hardware `json:"hardware"`
}

func (softlayer_hardware_power_component *SoftLayer_Hardware_Power_Component_Extended) String() string {
	return "SoftLayer_Hardware_Power_Component"
}
