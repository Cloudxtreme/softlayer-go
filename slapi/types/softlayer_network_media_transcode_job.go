package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Media_Transcode_Job - The SoftLayer_Network_Media_Transcode_Job contains
// information regarding a transcode job such as input file, output format, user id and so on.
type SoftLayer_Network_Media_Transcode_Job struct {

	// OutputFile - no documentation
	OutputFile string `json:"outputFile,omitempty"`

	// AutoDeleteDuration - The auto-deletion duration in seconds. This value determines how long the input
	// file will be kept on the storage.
	AutoDeleteDuration int `json:"autoDeleteDuration,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// TranscodeStatusId - no documentation
	TranscodeStatusId int `json:"transcodeStatusId,omitempty"`

	// UserId - The internal identifier of the user who created a transcode job
	UserId int `json:"userId,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// TranscodeAccountId - no documentation
	TranscodeAccountId int `json:"transcodeAccountId,omitempty"`

	// TranscodePresetName - no documentation
	TranscodePresetName string `json:"transcodePresetName,omitempty"`

	// TranscodeJobGuid - no documentation
	TranscodeJobGuid string `json:"transcodeJobGuid,omitempty"`

	// TranscodePresetGuid - no documentation
	TranscodePresetGuid string `json:"transcodePresetGuid,omitempty"`

	// Watermark - no documentation
	Watermark *SoftLayer_Container_Network_Media_Transcode_Job_Watermark `json:"watermark,omitempty"`

	// ByteIn - no documentation
	ByteIn int `json:"byteIn,omitempty"`

	// InputFile - no documentation
	InputFile string `json:"inputFile,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// User - no documentation
	User *SoftLayer_User_Customer `json:"user,omitempty"`

	// History - <nil>
	History []*SoftLayer_Network_Media_Transcode_Job_History `json:"history,omitempty"`

	// TranscodeAccount - no documentation
	TranscodeAccount *SoftLayer_Network_Media_Transcode_Account `json:"transcodeAccount,omitempty"`

	// TranscodeStatus - no documentation
	TranscodeStatus *SoftLayer_Network_Media_Transcode_Job_Status `json:"transcodeStatus,omitempty"`

	// TranscodeStatusName - no documentation
	TranscodeStatusName string `json:"transcodeStatusName,omitempty"`

	// HistoryCount - no documentation
	HistoryCount uint64 `json:"historyCount,omitempty"`
}

func (softlayer_network_media_transcode_job *SoftLayer_Network_Media_Transcode_Job) String() string {
	return "SoftLayer_Network_Media_Transcode_Job"
}
