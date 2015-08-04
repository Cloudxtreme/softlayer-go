package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Network_Media_Transcode_Job_Watermark - <nil>
type SoftLayer_Container_Network_Media_Transcode_Job_Watermark struct {

	// FileName - Filename of image to use as watermark in transcoding job
	FileName string `json:"fileName,omitempty"`

	// Position - no documentation
	Position *SoftLayer_Container_Network_Media_Transcode_Job_Watermark_Position `json:"position,omitempty"`

	// StartTime - no documentation
	StartTime int `json:"startTime,omitempty"`

	// Text - no documentation
	Text string `json:"text,omitempty"`

	// TransparencyPercentage - no documentation
	TransparencyPercentage int `json:"transparencyPercentage,omitempty"`

	// EndTime - no documentation
	EndTime int `json:"endTime,omitempty"`
}

func (softlayer_container_network_media_transcode_job_watermark *SoftLayer_Container_Network_Media_Transcode_Job_Watermark) String() string {
	return "SoftLayer_Container_Network_Media_Transcode_Job_Watermark"
}
