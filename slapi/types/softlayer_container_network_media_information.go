package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Network_Media_Information - This container class holds information on a media
// file such as file name, codec, frame rate and so on
type SoftLayer_Container_Network_Media_Information struct {

	// AudioChannelMode - no documentation
	AudioChannelMode string `json:"audioChannelMode"`

	// AudioChannels - no documentation
	AudioChannels int `json:"audioChannels"`

	// AudioCodec - no documentation
	AudioCodec string `json:"audioCodec"`

	// FileFormat - no documentation
	FileFormat string `json:"fileFormat"`

	// FileSize - no documentation
	FileSize uint64 `json:"fileSize"`

	// SizeY - no documentation
	SizeY int `json:"sizeY"`

	// TotalFrames - no documentation
	TotalFrames uint64 `json:"totalFrames"`

	// VideoAspectX - The width in a video's width to height aspect ratio
	VideoAspectX float32 `json:"videoAspectX"`

	// AudioSampleRate - no documentation
	AudioSampleRate int `json:"audioSampleRate"`

	// File - no documentation
	File string `json:"file"`

	// VideoCodec - no documentation
	VideoCodec string `json:"videoCodec"`

	// Duration - no documentation
	Duration float32 `json:"duration"`

	// ErrorMessage - no documentation
	ErrorMessage string `json:"errorMessage"`

	// SizeX - no documentation
	SizeX int `json:"sizeX"`

	// VideoAspectY - The height in a video's width to height aspect ratio
	VideoAspectY int `json:"videoAspectY"`

	// AudioBitRate - no documentation
	AudioBitRate int `json:"audioBitRate"`

	// FrameRate - no documentation
	FrameRate float32 `json:"frameRate"`
}

func (softlayer_container_network_media_information *SoftLayer_Container_Network_Media_Information) String() string {
	return "SoftLayer_Container_Network_Media_Information"
}
