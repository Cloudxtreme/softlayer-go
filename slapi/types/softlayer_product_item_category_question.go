package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Item_Category_Question - The SoftLayer_Product_Item_Category_Question data type
// represents a single question to be answered by an end user. The question may or may not be required
// which can be located by looking at the 'required' property on the item category references. The
// answerValueExpression property is a regular expression that is used to validate the answer to the
// question. The description and valueExample properties can be used to get an idea of the type of
// answer that should be provided.
type SoftLayer_Product_Item_Category_Question struct {

	// AnswerValueExpression - no documentation
	AnswerValueExpression string `json:"answerValueExpression"`

	// Description - no documentation
	Description string `json:"description"`

	// FieldType - The type of field that should be used in an form to accept an answer from an end user.
	FieldType *SoftLayer_Product_Item_Category_Question_Field_Type `json:"fieldType"`

	// FieldTypeId - no documentation
	FieldTypeId int `json:"fieldTypeId"`

	// Id - no documentation
	Id int `json:"id"`

	// ItemCategoryReferenceCount - A count of the link between an item category and an item category
	// question.
	ItemCategoryReferenceCount uint64 `json:"itemCategoryReferenceCount"`

	// ItemCategoryReferences - The link between an item category and an item category question.
	ItemCategoryReferences []*SoftLayer_Product_Item_Category_Question_Xref `json:"itemCategoryReferences"`

	// KeyName - no documentation
	KeyName string `json:"keyName"`

	// Question - no documentation
	Question string `json:"question"`

	// ValueExample - An example and/or explanation of what the answer for the question is expected to look
	// like.
	ValueExample string `json:"valueExample"`
}
