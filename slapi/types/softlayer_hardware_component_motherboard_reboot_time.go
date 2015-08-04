package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Hardware_Component_Motherboard_Reboot_Time - The
// SoftLayer_Hardware_Component_Motherboard_Reboot_Time contains the average reboot times for
// motherboards. There are two types of average times. One is for motherboards without raid, and the
// other is for motherboards with raid. These times are based on averages and have been gathered
// through numerous test cases.
type SoftLayer_Hardware_Component_Motherboard_Reboot_Time struct {

	// WithRaid - Average reboot time in seconds for the motherboard when raid is installed.
	WithRaid int `json:"withRaid,omitempty"`

	// WithoutRaid - Average reboot time in seconds for the motherboard when NO raid is installed.
	WithoutRaid int `json:"withoutRaid,omitempty"`
}

func (softlayer_hardware_component_motherboard_reboot_time *SoftLayer_Hardware_Component_Motherboard_Reboot_Time) String() string {
	return "SoftLayer_Hardware_Component_Motherboard_Reboot_Time"
}

// SoftLayer_Hardware_Component_Motherboard_Reboot_Time_Extended is SoftLayer_Hardware_Component_Motherboard_Reboot_Time with all maskable types.
type SoftLayer_Hardware_Component_Motherboard_Reboot_Time_Extended struct {
	SoftLayer_Hardware_Component_Motherboard_Reboot_Time

	// HardwareComponentModel - Motherboard's specifications (manufacturer, version, etc....)
	HardwareComponentModel *SoftLayer_Hardware_Component_Model `json:"hardwareComponentModel,omitempty"`
}

func (softlayer_hardware_component_motherboard_reboot_time *SoftLayer_Hardware_Component_Motherboard_Reboot_Time_Extended) String() string {
	return "SoftLayer_Hardware_Component_Motherboard_Reboot_Time"
}
