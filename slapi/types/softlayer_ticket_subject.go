package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Ticket_Subject - The SoftLayer_Ticket_Subject data type models one of the possible
// subjects that a standard support ticket may belong to. A basic support ticket's title matches it's
// corresponding subject's name.
type SoftLayer_Ticket_Subject struct {

	// Group - <nil>
	Group *SoftLayer_Ticket_Group `json:"group"`

	// Id - no documentation
	Id int `json:"id"`

	// Name - A ticket subject's name. This name is used for a standard support ticket's title.
	Name string `json:"name"`
}

// GetAllObjects - Retrieve all possible ticket subjects. The SoftLayer customer portal uses this
// method in the add standard support ticket form.
func (softlayer_ticket_subject *SoftLayer_Ticket_Subject) GetAllObjects(ctx *slapi.RequestContext) ([]*SoftLayer_Ticket_Subject, error) {
	var returnValue []*SoftLayer_Ticket_Subject
	return returnValue, nil
}

// GetObject - getObject retrieves the SoftLayer_Ticket_Subject object whose ID number corresponds to
// the ID number of the init parameter passed to the SoftLayer_Ticket_Subject service.
func (softlayer_ticket_subject *SoftLayer_Ticket_Subject) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Ticket_Subject, error) {
	var returnValue *SoftLayer_Ticket_Subject
	return returnValue, nil
}

// GetTopFiveKnowledgeLayerQuestions - SoftLayer maintains relationships between the generic subjects
// for standard administration and the top five commonly asked questions about these subjects.
// getTopFileKnowledgeLayerQuestions() retrieves the top five questions and answers from the SoftLayer
// KnowledgeLayer related to the given ticket subject.
func (softlayer_ticket_subject *SoftLayer_Ticket_Subject) GetTopFiveKnowledgeLayerQuestions(ctx *slapi.RequestContext) ([]*SoftLayer_Container_KnowledgeLayer_QuestionAnswer, error) {
	var returnValue []*SoftLayer_Container_KnowledgeLayer_QuestionAnswer
	return returnValue, nil
}
