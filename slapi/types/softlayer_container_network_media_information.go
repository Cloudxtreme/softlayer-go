package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Container_Network_Media_Information - This container class holds information on a media
// file such as file name, codec, frame rate and so on
type SoftLayer_Container_Network_Media_Information struct {

	// FrameRate - no documentation
	FrameRate slapi.Float64 `json:"frameRate,omitempty"`

	// File - no documentation
	File string `json:"file,omitempty"`

	// AudioCodec - no documentation
	AudioCodec string `json:"audioCodec,omitempty"`

	// Duration - no documentation
	Duration slapi.Float64 `json:"duration,omitempty"`

	// VideoCodec - no documentation
	VideoCodec string `json:"videoCodec,omitempty"`

	// AudioBitRate - no documentation
	AudioBitRate int `json:"audioBitRate,omitempty"`

	// FileSize - no documentation
	FileSize uint64 `json:"fileSize,omitempty"`

	// SizeX - no documentation
	SizeX int `json:"sizeX,omitempty"`

	// SizeY - no documentation
	SizeY int `json:"sizeY,omitempty"`

	// TotalFrames - no documentation
	TotalFrames uint64 `json:"totalFrames,omitempty"`

	// VideoAspectY - The height in a video's width to height aspect ratio
	VideoAspectY int `json:"videoAspectY,omitempty"`

	// AudioSampleRate - no documentation
	AudioSampleRate int `json:"audioSampleRate,omitempty"`

	// AudioChannels - no documentation
	AudioChannels int `json:"audioChannels,omitempty"`

	// ErrorMessage - no documentation
	ErrorMessage string `json:"errorMessage,omitempty"`

	// FileFormat - no documentation
	FileFormat string `json:"fileFormat,omitempty"`

	// VideoAspectX - The width in a video's width to height aspect ratio
	VideoAspectX slapi.Float64 `json:"videoAspectX,omitempty"`

	// AudioChannelMode - no documentation
	AudioChannelMode string `json:"audioChannelMode,omitempty"`
}

func (softlayer_container_network_media_information *SoftLayer_Container_Network_Media_Information) String() string {
	return "SoftLayer_Container_Network_Media_Information"
}
