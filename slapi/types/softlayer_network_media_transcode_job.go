package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Media_Transcode_Job - The SoftLayer_Network_Media_Transcode_Job contains
// information regarding a transcode job such as input file, output format, user id and so on.
type SoftLayer_Network_Media_Transcode_Job struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// TranscodeJobGuid - no documentation
	TranscodeJobGuid string `json:"transcodeJobGuid"`

	// Id - no documentation
	Id int `json:"id"`

	// InputFile - no documentation
	InputFile string `json:"inputFile"`

	// TranscodeAccountId - no documentation
	TranscodeAccountId int `json:"transcodeAccountId"`

	// TranscodePresetName - no documentation
	TranscodePresetName string `json:"transcodePresetName"`

	// TranscodeStatusId - no documentation
	TranscodeStatusId int `json:"transcodeStatusId"`

	// UserId - The internal identifier of the user who created a transcode job
	UserId int `json:"userId"`

	// ByteIn - no documentation
	ByteIn int `json:"byteIn"`

	// OutputFile - no documentation
	OutputFile string `json:"outputFile"`

	// Watermark - no documentation
	Watermark *SoftLayer_Container_Network_Media_Transcode_Job_Watermark `json:"watermark"`

	// AutoDeleteDuration - The auto-deletion duration in seconds. This value determines how long the input
	// file will be kept on the storage.
	AutoDeleteDuration int `json:"autoDeleteDuration"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// Name - no documentation
	Name string `json:"name"`

	// TranscodePresetGuid - no documentation
	TranscodePresetGuid string `json:"transcodePresetGuid"`
}

func (softlayer_network_media_transcode_job *SoftLayer_Network_Media_Transcode_Job) String() string {
	return "SoftLayer_Network_Media_Transcode_Job"
}

// SoftLayer_Network_Media_Transcode_Job_Extended is SoftLayer_Network_Media_Transcode_Job with all maskable types.
type SoftLayer_Network_Media_Transcode_Job_Extended struct {
	SoftLayer_Network_Media_Transcode_Job

	// HistoryCount - no documentation
	HistoryCount uint64 `json:"historyCount"`

	// History - <nil>
	History []*SoftLayer_Network_Media_Transcode_Job_History `json:"history"`

	// TranscodeAccount - no documentation
	TranscodeAccount *SoftLayer_Network_Media_Transcode_Account `json:"transcodeAccount"`

	// TranscodeStatusName - no documentation
	TranscodeStatusName string `json:"transcodeStatusName"`

	// TranscodeStatus - no documentation
	TranscodeStatus *SoftLayer_Network_Media_Transcode_Job_Status `json:"transcodeStatus"`

	// User - no documentation
	User *SoftLayer_User_Customer `json:"user"`
}

func (softlayer_network_media_transcode_job *SoftLayer_Network_Media_Transcode_Job_Extended) String() string {
	return "SoftLayer_Network_Media_Transcode_Job"
}
