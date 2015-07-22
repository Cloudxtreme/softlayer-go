package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_KnowledgeLayer_QuestionAnswer -
// SoftLayer_Container_KnowledgeLayer_QuestionAnswer models a single question and answer pair from
// SoftLayer's KnowledgeLayer knowledge base. SoftLayer's backend network interfaces with the
// KnowledgeLayer to recommend helpful articles when support tickets are created.
type SoftLayer_Container_KnowledgeLayer_QuestionAnswer struct {

	// Answer - The answer to a question asked on the SoftLayer KnowledgeLayer.
	Answer string `json:"answer"`

	// Link - The link to a question asked on the SoftLayer KnowledgeLayer.
	Link string `json:"link"`

	// Question - no documentation
	Question string `json:"question"`
}