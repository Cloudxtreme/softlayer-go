package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Disk_Image_Capture_Template_Volume - <nil>
type SoftLayer_Container_Disk_Image_Capture_Template_Volume struct {

	// Name - <nil>
	Name string `json:"name,omitempty"`

	// Partitions - <nil>
	Partitions []*SoftLayer_Container_Disk_Image_Capture_Template_Volume_Partition `json:"partitions,omitempty"`
}

func (softlayer_container_disk_image_capture_template_volume *SoftLayer_Container_Disk_Image_Capture_Template_Volume) String() string {
	return "SoftLayer_Container_Disk_Image_Capture_Template_Volume"
}
