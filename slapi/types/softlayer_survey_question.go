package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Survey_Question - The SoftLayer_Survey_Question data type contains general information
// relating to a single SoftLayer survey question.
type SoftLayer_Survey_Question struct {

	// IsRequired - A flag indicating that a survey question requires a response.
	IsRequired int `json:"isRequired"`

	// QuestionOrder - A value indicating the order in when a survey question will be asked.
	QuestionOrder int `json:"questionOrder"`

	// Id - no documentation
	Id int `json:"id"`

	// MultiAnswer - A flag indicating that a survey question can have multiple answers responded to.
	MultiAnswer int `json:"multiAnswer"`

	// Question - no documentation
	Question string `json:"question"`

	// SurveyId - A survey question's associated [[SoftLayer_Survey|Survey]] Id.
	SurveyId int `json:"surveyId"`
}

// SoftLayer_Survey_Question_Extended is SoftLayer_Survey_Question with all maskable types.
type SoftLayer_Survey_Question_Extended struct {
	SoftLayer_Survey_Question

	// Survey - no documentation
	Survey *SoftLayer_Survey `json:"survey"`

	// AnswerCount - A count of the possible answers for a survey question.
	AnswerCount uint64 `json:"answerCount"`

	// Answers - no documentation
	Answers []*SoftLayer_Survey_Answer `json:"answers"`
}

func (softlayer_survey_question *SoftLayer_Survey_Question) String() string {
	return "SoftLayer_Survey_Question"
}
