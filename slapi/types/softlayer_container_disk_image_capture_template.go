package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Disk_Image_Capture_Template - <nil>
type SoftLayer_Container_Disk_Image_Capture_Template struct {

	// Description - <nil>
	Description string `json:"description"`

	// Name - <nil>
	Name string `json:"name"`

	// Summary - <nil>
	Summary string `json:"summary"`

	// Volumes - <nil>
	Volumes []*SoftLayer_Container_Disk_Image_Capture_Template_Volume `json:"volumes"`
}

func (softlayer_container_disk_image_capture_template *SoftLayer_Container_Disk_Image_Capture_Template) String() string {
	return "SoftLayer_Container_Disk_Image_Capture_Template"
}
