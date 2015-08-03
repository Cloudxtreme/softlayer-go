package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Survey_Answer - The SoftLayer_Survey_Answer data type contains general information
// relating to a single SoftLayer survey answer.
type SoftLayer_Survey_Answer struct {

	// SurveyQuestionId - A survey answer's associated [[SoftLayer_Survey_Question|Survey Question]] Id.
	SurveyQuestionId int `json:"surveyQuestionId"`

	// Answer - A survey answer's answer that a user can response too.
	Answer string `json:"answer"`

	// AnswerOrder - A value indicating the order in when a survey answer will be displayed to a user.
	AnswerOrder int `json:"answerOrder"`

	// Id - no documentation
	Id int `json:"id"`
}

func (softlayer_survey_answer *SoftLayer_Survey_Answer) String() string {
	return "SoftLayer_Survey_Answer"
}

// SoftLayer_Survey_Answer_Extended is SoftLayer_Survey_Answer with all maskable types.
type SoftLayer_Survey_Answer_Extended struct {
	SoftLayer_Survey_Answer

	// SurveyQuestion - no documentation
	SurveyQuestion *SoftLayer_Survey_Question `json:"surveyQuestion"`
}

func (softlayer_survey_answer *SoftLayer_Survey_Answer_Extended) String() string {
	return "SoftLayer_Survey_Answer"
}
