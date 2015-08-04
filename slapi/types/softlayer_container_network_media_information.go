package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Network_Media_Information - This container class holds information on a media
// file such as file name, codec, frame rate and so on
type SoftLayer_Container_Network_Media_Information struct {

	// ErrorMessage - no documentation
	ErrorMessage string `json:"errorMessage,omitempty"`

	// FrameRate - no documentation
	FrameRate float32 `json:"frameRate,omitempty"`

	// Duration - no documentation
	Duration float32 `json:"duration,omitempty"`

	// AudioChannelMode - no documentation
	AudioChannelMode string `json:"audioChannelMode,omitempty"`

	// AudioCodec - no documentation
	AudioCodec string `json:"audioCodec,omitempty"`

	// FileSize - no documentation
	FileSize uint64 `json:"fileSize,omitempty"`

	// SizeY - no documentation
	SizeY int `json:"sizeY,omitempty"`

	// TotalFrames - no documentation
	TotalFrames uint64 `json:"totalFrames,omitempty"`

	// VideoAspectX - The width in a video's width to height aspect ratio
	VideoAspectX float32 `json:"videoAspectX,omitempty"`

	// AudioBitRate - no documentation
	AudioBitRate int `json:"audioBitRate,omitempty"`

	// File - no documentation
	File string `json:"file,omitempty"`

	// VideoCodec - no documentation
	VideoCodec string `json:"videoCodec,omitempty"`

	// AudioChannels - no documentation
	AudioChannels int `json:"audioChannels,omitempty"`

	// FileFormat - no documentation
	FileFormat string `json:"fileFormat,omitempty"`

	// SizeX - no documentation
	SizeX int `json:"sizeX,omitempty"`

	// VideoAspectY - The height in a video's width to height aspect ratio
	VideoAspectY int `json:"videoAspectY,omitempty"`

	// AudioSampleRate - no documentation
	AudioSampleRate int `json:"audioSampleRate,omitempty"`
}

func (softlayer_container_network_media_information *SoftLayer_Container_Network_Media_Information) String() string {
	return "SoftLayer_Container_Network_Media_Information"
}
