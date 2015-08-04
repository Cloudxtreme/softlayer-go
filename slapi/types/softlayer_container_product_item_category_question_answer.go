package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Item_Category_Question_Answer - The
// SoftLayer_Container_Product_Item_Category_Question_Answer data type represents an answer to an item
// category question. It contains the category, the question being answered, and the answer.
type SoftLayer_Container_Product_Item_Category_Question_Answer struct {

	// Answer - no documentation
	Answer string `json:"answer,omitempty"`

	// CategoryCode - no documentation
	CategoryCode string `json:"categoryCode,omitempty"`

	// CategoryId - no documentation
	CategoryId int `json:"categoryId,omitempty"`

	// QuestionId - no documentation
	QuestionId int `json:"questionId,omitempty"`
}

func (softlayer_container_product_item_category_question_answer *SoftLayer_Container_Product_Item_Category_Question_Answer) String() string {
	return "SoftLayer_Container_Product_Item_Category_Question_Answer"
}
