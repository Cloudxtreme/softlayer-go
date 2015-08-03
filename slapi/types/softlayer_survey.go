package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Survey - The SoftLayer_Survey data type contains general information relating to a single
// SoftLayer survey.
type SoftLayer_Survey struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Active - no documentation
	Active int `json:"active"`

	// Id - no documentation
	Id int `json:"id"`

	// Name - no documentation
	Name string `json:"name"`

	// StatusId - no documentation
	StatusId int `json:"statusId"`

	// TypeId - no documentation
	TypeId int `json:"typeId"`
}

func (softlayer_survey *SoftLayer_Survey) String() string {
	return "SoftLayer_Survey"
}

// SoftLayer_Survey_Extended is SoftLayer_Survey with all maskable types.
type SoftLayer_Survey_Extended struct {
	SoftLayer_Survey

	// QuestionCount - no documentation
	QuestionCount uint64 `json:"questionCount"`

	// Questions - no documentation
	Questions []*SoftLayer_Survey_Question `json:"questions"`

	// Status - no documentation
	Status *SoftLayer_Survey_Status `json:"status"`

	// Type - no documentation
	Type *SoftLayer_Survey_Type `json:"type"`
}

func (softlayer_survey *SoftLayer_Survey_Extended) String() string {
	return "SoftLayer_Survey"
}
