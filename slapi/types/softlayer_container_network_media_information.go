package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Network_Media_Information - This container class holds information on a media
// file such as file name, codec, frame rate and so on
type SoftLayer_Container_Network_Media_Information struct {

	// AudioBitRate - no documentation
	AudioBitRate int `json:"audioBitRate"`

	// AudioChannelMode - no documentation
	AudioChannelMode string `json:"audioChannelMode"`

	// AudioChannels - no documentation
	AudioChannels int `json:"audioChannels"`

	// AudioCodec - no documentation
	AudioCodec string `json:"audioCodec"`

	// AudioSampleRate - no documentation
	AudioSampleRate int `json:"audioSampleRate"`

	// Duration - no documentation
	Duration float32 `json:"duration"`

	// ErrorMessage - no documentation
	ErrorMessage string `json:"errorMessage"`

	// File - no documentation
	File string `json:"file"`

	// FileFormat - no documentation
	FileFormat string `json:"fileFormat"`

	// FileSize - no documentation
	FileSize uint64 `json:"fileSize"`

	// FrameRate - no documentation
	FrameRate float32 `json:"frameRate"`

	// SizeX - no documentation
	SizeX int `json:"sizeX"`

	// SizeY - no documentation
	SizeY int `json:"sizeY"`

	// TotalFrames - no documentation
	TotalFrames uint64 `json:"totalFrames"`

	// VideoAspectX - The width in a video's width to height aspect ratio
	VideoAspectX float32 `json:"videoAspectX"`

	// VideoAspectY - The height in a video's width to height aspect ratio
	VideoAspectY int `json:"videoAspectY"`

	// VideoCodec - no documentation
	VideoCodec string `json:"videoCodec"`
}

func (softlayer_container_network_media_information *SoftLayer_Container_Network_Media_Information) String() string {
	return "SoftLayer_Container_Network_Media_Information"
}
