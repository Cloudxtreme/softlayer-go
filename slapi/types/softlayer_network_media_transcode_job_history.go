package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Media_Transcode_Job_History - <nil>
type SoftLayer_Network_Media_Transcode_Job_History struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// PublicNotes - no documentation
	PublicNotes string `json:"publicNotes"`

	// TranscodeJobId - no documentation
	TranscodeJobId int `json:"transcodeJobId"`

	// TranscodeStatusName - no documentation
	TranscodeStatusName string `json:"transcodeStatusName"`
}

func (softlayer_network_media_transcode_job_history *SoftLayer_Network_Media_Transcode_Job_History) String() string {
	return "SoftLayer_Network_Media_Transcode_Job_History"
}
