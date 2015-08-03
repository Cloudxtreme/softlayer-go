package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Utility_Message - <nil>
type SoftLayer_Container_Utility_Message struct {

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// Id - <nil>
	Id int `json:"id"`

	// Message - <nil>
	Message string `json:"message"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate"`

	// Summary - <nil>
	Summary string `json:"summary"`
}

func (softlayer_container_utility_message *SoftLayer_Container_Utility_Message) String() string {
	return "SoftLayer_Container_Utility_Message"
}
