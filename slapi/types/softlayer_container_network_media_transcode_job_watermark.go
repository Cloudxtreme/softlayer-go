package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Network_Media_Transcode_Job_Watermark - <nil>
type SoftLayer_Container_Network_Media_Transcode_Job_Watermark struct {

	// TransparencyPercentage - no documentation
	TransparencyPercentage int `json:"transparencyPercentage"`

	// EndTime - no documentation
	EndTime int `json:"endTime"`

	// FileName - Filename of image to use as watermark in transcoding job
	FileName string `json:"fileName"`

	// Position - no documentation
	Position *SoftLayer_Container_Network_Media_Transcode_Job_Watermark_Position `json:"position"`

	// StartTime - no documentation
	StartTime int `json:"startTime"`

	// Text - no documentation
	Text string `json:"text"`
}

func (softlayer_container_network_media_transcode_job_watermark *SoftLayer_Container_Network_Media_Transcode_Job_Watermark) String() string {
	return "SoftLayer_Container_Network_Media_Transcode_Job_Watermark"
}
