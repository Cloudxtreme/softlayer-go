package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Survey_Response - The SoftLayer_Survey_Response data type contains general information
// relating to a single SoftLayer survey response.
type SoftLayer_Survey_Response struct {

	// OtherAnswer - The user typed response for the [[SoftLayer_Survey_Answer|Survey Answer]] that a
	// response is associated with.
	OtherAnswer string `json:"otherAnswer,omitempty"`

	// SurveyAnswerId - The Id of the [[SoftLayer_Survey_Answer|Survey Answer]] that a response was made
	// for.
	SurveyAnswerId int `json:"surveyAnswerId,omitempty"`

	// SurveyAnswer - no documentation
	SurveyAnswer *SoftLayer_Survey_Answer `json:"surveyAnswer,omitempty"`
}

func (softlayer_survey_response *SoftLayer_Survey_Response) String() string {
	return "SoftLayer_Survey_Response"
}
