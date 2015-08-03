package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Billing_Order_Item_Category_Answer - The SoftLayer_Billing_Order_Item_Category_Answer data
// type represents a single answer to an item category question.
type SoftLayer_Billing_Order_Item_Category_Answer struct {

	// Answer - no documentation
	Answer string `json:"answer"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// QuestionId - The identifier for the question that the answer belongs to.
	QuestionId int `json:"questionId"`
}

// SoftLayer_Billing_Order_Item_Category_Answer_Extended is SoftLayer_Billing_Order_Item_Category_Answer with all maskable types.
type SoftLayer_Billing_Order_Item_Category_Answer_Extended struct {
	SoftLayer_Billing_Order_Item_Category_Answer

	// OrderItem - no documentation
	OrderItem *SoftLayer_Billing_Order_Item `json:"orderItem"`

	// Question - no documentation
	Question *SoftLayer_Product_Item_Category_Question `json:"question"`
}

func (softlayer_billing_order_item_category_answer *SoftLayer_Billing_Order_Item_Category_Answer) String() string {
	return "SoftLayer_Billing_Order_Item_Category_Answer"
}
