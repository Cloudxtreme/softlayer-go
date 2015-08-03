package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Media_Transcode_Account - The SoftLayer_Network_Media_Transcode_Account contains
// information regarding a transcode account.
type SoftLayer_Network_Media_Transcode_Account struct {

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// AccountId - no documentation
	AccountId int `json:"accountId"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Id - no documentation
	Id int `json:"id"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// TranscodeJobCount - no documentation
	TranscodeJobCount uint64 `json:"transcodeJobCount"`

	// TranscodeJobs - no documentation
	TranscodeJobs []*SoftLayer_Network_Media_Transcode_Job `json:"transcodeJobs"`
}

func (softlayer_network_media_transcode_account *SoftLayer_Network_Media_Transcode_Account) String() string {
	return "SoftLayer_Network_Media_Transcode_Account"
}
