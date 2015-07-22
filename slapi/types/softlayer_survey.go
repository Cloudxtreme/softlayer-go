package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
)

// SoftLayer_Survey - The SoftLayer_Survey data type contains general information relating to a single
// SoftLayer survey.
type SoftLayer_Survey struct {

	// Active - no documentation
	Active int `json:"active"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Id - no documentation
	Id int `json:"id"`

	// Name - no documentation
	Name string `json:"name"`

	// QuestionCount - no documentation
	QuestionCount uint64 `json:"questionCount"`

	// Questions - no documentation
	Questions []*SoftLayer_Survey_Question `json:"questions"`

	// Status - no documentation
	Status *SoftLayer_Survey_Status `json:"status"`

	// StatusId - no documentation
	StatusId int `json:"statusId"`

	// Type - no documentation
	Type *SoftLayer_Survey_Type `json:"type"`

	// TypeId - no documentation
	TypeId int `json:"typeId"`
}

// GetActiveSurveyByType - <nil>
func (softlayer_survey *SoftLayer_Survey) GetActiveSurveyByType(ctx *slapi.RequestContext, type_ string) (*SoftLayer_Survey, error) {
	var returnValue *SoftLayer_Survey
	return returnValue, nil
}

// GetObject - getObject retrieves the SoftLayer_Survey object whose ID number corresponds to the ID
// number of the init parameter passed to the SoftLayer_Survey service. You can only retrieve the
// survey that your portal user has taken.
func (softlayer_survey *SoftLayer_Survey) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Survey, error) {
	var returnValue *SoftLayer_Survey
	return returnValue, nil
}

// TakeSurvey - no documentation
func (softlayer_survey *SoftLayer_Survey) TakeSurvey(ctx *slapi.RequestContext, responses []SoftLayer_Survey_Response) (bool, error) {
	var returnValue bool
	return returnValue, nil
}
