package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Network_Media_Information - This container class holds information on a media
// file such as file name, codec, frame rate and so on
type SoftLayer_Container_Network_Media_Information struct {

	// AudioSampleRate - no documentation
	AudioSampleRate int `json:"audioSampleRate"`

	// ErrorMessage - no documentation
	ErrorMessage string `json:"errorMessage"`

	// FileFormat - no documentation
	FileFormat string `json:"fileFormat"`

	// VideoCodec - no documentation
	VideoCodec string `json:"videoCodec"`

	// VideoAspectX - The width in a video's width to height aspect ratio
	VideoAspectX float32 `json:"videoAspectX"`

	// VideoAspectY - The height in a video's width to height aspect ratio
	VideoAspectY int `json:"videoAspectY"`

	// AudioChannelMode - no documentation
	AudioChannelMode string `json:"audioChannelMode"`

	// AudioCodec - no documentation
	AudioCodec string `json:"audioCodec"`

	// FileSize - no documentation
	FileSize uint64 `json:"fileSize"`

	// TotalFrames - no documentation
	TotalFrames uint64 `json:"totalFrames"`

	// AudioBitRate - no documentation
	AudioBitRate int `json:"audioBitRate"`

	// Duration - no documentation
	Duration float32 `json:"duration"`

	// SizeY - no documentation
	SizeY int `json:"sizeY"`

	// AudioChannels - no documentation
	AudioChannels int `json:"audioChannels"`

	// File - no documentation
	File string `json:"file"`

	// FrameRate - no documentation
	FrameRate float32 `json:"frameRate"`

	// SizeX - no documentation
	SizeX int `json:"sizeX"`
}

func (softlayer_container_network_media_information *SoftLayer_Container_Network_Media_Information) String() string {
	return "SoftLayer_Container_Network_Media_Information"
}
