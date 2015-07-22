package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Item_Category_Question_Xref - The SoftLayer_Product_Item_Category_Question_Xref
// data type represents a link between an item category and an item category question. It also contains
// a 'required' field that designates if the question is required to be answered for the given item
// category.
type SoftLayer_Product_Item_Category_Question_Xref struct {

	// Id - no documentation
	Id int `json:"id"`

	// ItemCategory - The product item category that this reference points to.
	ItemCategory *SoftLayer_Product_Item_Category `json:"itemCategory"`

	// ItemCategoryId - no documentation
	ItemCategoryId int `json:"itemCategoryId"`

	// LocationId - no documentation
	LocationId int `json:"locationId"`

	// Question - The item category question that this reference points to.
	Question *SoftLayer_Product_Item_Category_Question `json:"question"`

	// QuestionId - no documentation
	QuestionId int `json:"questionId"`

	// Required - Flag to indicate whether an answer is required for the question..
	Required bool `json:"required"`
}
