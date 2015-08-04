package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Media_Transcode_Job_History - <nil>
type SoftLayer_Network_Media_Transcode_Job_History struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// PublicNotes - no documentation
	PublicNotes string `json:"publicNotes,omitempty"`

	// TranscodeJobId - no documentation
	TranscodeJobId int `json:"transcodeJobId,omitempty"`
}

func (softlayer_network_media_transcode_job_history *SoftLayer_Network_Media_Transcode_Job_History) String() string {
	return "SoftLayer_Network_Media_Transcode_Job_History"
}

// SoftLayer_Network_Media_Transcode_Job_History_Extended is SoftLayer_Network_Media_Transcode_Job_History with all maskable types.
type SoftLayer_Network_Media_Transcode_Job_History_Extended struct {
	SoftLayer_Network_Media_Transcode_Job_History

	// TranscodeStatusName - no documentation
	TranscodeStatusName string `json:"transcodeStatusName,omitempty"`
}

func (softlayer_network_media_transcode_job_history *SoftLayer_Network_Media_Transcode_Job_History_Extended) String() string {
	return "SoftLayer_Network_Media_Transcode_Job_History"
}
