package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Media_Transcode_Job - The SoftLayer_Network_Media_Transcode_Job contains
// information regarding a transcode job such as input file, output format, user id and so on.
type SoftLayer_Network_Media_Transcode_Job struct {

	// TranscodeAccountId - no documentation
	TranscodeAccountId int `json:"transcodeAccountId,omitempty"`

	// TranscodePresetName - no documentation
	TranscodePresetName string `json:"transcodePresetName,omitempty"`

	// Watermark - no documentation
	Watermark *SoftLayer_Container_Network_Media_Transcode_Job_Watermark `json:"watermark,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// TranscodeStatusId - no documentation
	TranscodeStatusId int `json:"transcodeStatusId,omitempty"`

	// UserId - The internal identifier of the user who created a transcode job
	UserId int `json:"userId,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// TranscodePresetGuid - no documentation
	TranscodePresetGuid string `json:"transcodePresetGuid,omitempty"`

	// TranscodeJobGuid - no documentation
	TranscodeJobGuid string `json:"transcodeJobGuid,omitempty"`

	// AutoDeleteDuration - The auto-deletion duration in seconds. This value determines how long the input
	// file will be kept on the storage.
	AutoDeleteDuration int `json:"autoDeleteDuration,omitempty"`

	// ByteIn - no documentation
	ByteIn int `json:"byteIn,omitempty"`

	// InputFile - no documentation
	InputFile string `json:"inputFile,omitempty"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// OutputFile - no documentation
	OutputFile string `json:"outputFile,omitempty"`
}

func (softlayer_network_media_transcode_job *SoftLayer_Network_Media_Transcode_Job) String() string {
	return "SoftLayer_Network_Media_Transcode_Job"
}

// SoftLayer_Network_Media_Transcode_Job_Extended is SoftLayer_Network_Media_Transcode_Job with all maskable types.
type SoftLayer_Network_Media_Transcode_Job_Extended struct {
	SoftLayer_Network_Media_Transcode_Job

	// TranscodeStatus - no documentation
	TranscodeStatus *SoftLayer_Network_Media_Transcode_Job_Status `json:"transcodeStatus,omitempty"`

	// TranscodeAccount - no documentation
	TranscodeAccount *SoftLayer_Network_Media_Transcode_Account `json:"transcodeAccount,omitempty"`

	// TranscodeStatusName - no documentation
	TranscodeStatusName string `json:"transcodeStatusName,omitempty"`

	// User - no documentation
	User *SoftLayer_User_Customer `json:"user,omitempty"`

	// HistoryCount - no documentation
	HistoryCount uint64 `json:"historyCount,omitempty"`

	// History - <nil>
	History []*SoftLayer_Network_Media_Transcode_Job_History `json:"history,omitempty"`
}

func (softlayer_network_media_transcode_job *SoftLayer_Network_Media_Transcode_Job_Extended) String() string {
	return "SoftLayer_Network_Media_Transcode_Job"
}
