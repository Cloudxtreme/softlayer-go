package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Item_Category_Question - The SoftLayer_Product_Item_Category_Question data type
// represents a single question to be answered by an end user. The question may or may not be required
// which can be located by looking at the 'required' property on the item category references. The
// answerValueExpression property is a regular expression that is used to validate the answer to the
// question. The description and valueExample properties can be used to get an idea of the type of
// answer that should be provided.
type SoftLayer_Product_Item_Category_Question struct {

	// AnswerValueExpression - no documentation
	AnswerValueExpression string `json:"answerValueExpression,omitempty"`

	// Description - no documentation
	Description string `json:"description,omitempty"`

	// ValueExample - An example and/or explanation of what the answer for the question is expected to look
	// like.
	ValueExample string `json:"valueExample,omitempty"`

	// FieldTypeId - no documentation
	FieldTypeId int `json:"fieldTypeId,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// KeyName - no documentation
	KeyName string `json:"keyName,omitempty"`

	// Question - no documentation
	Question string `json:"question,omitempty"`
}

func (softlayer_product_item_category_question *SoftLayer_Product_Item_Category_Question) String() string {
	return "SoftLayer_Product_Item_Category_Question"
}

// SoftLayer_Product_Item_Category_Question_Extended is SoftLayer_Product_Item_Category_Question with all maskable types.
type SoftLayer_Product_Item_Category_Question_Extended struct {
	SoftLayer_Product_Item_Category_Question

	// ItemCategoryReferenceCount - A count of the link between an item category and an item category
	// question.
	ItemCategoryReferenceCount uint64 `json:"itemCategoryReferenceCount,omitempty"`

	// FieldType - The type of field that should be used in an form to accept an answer from an end user.
	FieldType *SoftLayer_Product_Item_Category_Question_Field_Type `json:"fieldType,omitempty"`

	// ItemCategoryReferences - The link between an item category and an item category question.
	ItemCategoryReferences []*SoftLayer_Product_Item_Category_Question_Xref `json:"itemCategoryReferences,omitempty"`
}

func (softlayer_product_item_category_question *SoftLayer_Product_Item_Category_Question_Extended) String() string {
	return "SoftLayer_Product_Item_Category_Question"
}
