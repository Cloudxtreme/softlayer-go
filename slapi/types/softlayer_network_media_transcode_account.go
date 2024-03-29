package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Media_Transcode_Account - The SoftLayer_Network_Media_Transcode_Account contains
// information regarding a transcode account.
type SoftLayer_Network_Media_Transcode_Account struct {

	// AccountId - no documentation
	AccountId int `json:"accountId,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// TranscodeJobCount - no documentation
	TranscodeJobCount uint64 `json:"transcodeJobCount,omitempty"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account,omitempty"`

	// TranscodeJobs - no documentation
	TranscodeJobs []*SoftLayer_Network_Media_Transcode_Job `json:"transcodeJobs,omitempty"`
}

func (softlayer_network_media_transcode_account *SoftLayer_Network_Media_Transcode_Account) String() string {
	return "SoftLayer_Network_Media_Transcode_Account"
}
