package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Billing_Order_Item_Category_Answer - The SoftLayer_Billing_Order_Item_Category_Answer data
// type represents a single answer to an item category question.
type SoftLayer_Billing_Order_Item_Category_Answer struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// QuestionId - The identifier for the question that the answer belongs to.
	QuestionId int `json:"questionId,omitempty"`

	// Answer - no documentation
	Answer string `json:"answer,omitempty"`

	// OrderItem - no documentation
	OrderItem *SoftLayer_Billing_Order_Item `json:"orderItem,omitempty"`

	// Question - no documentation
	Question *SoftLayer_Product_Item_Category_Question `json:"question,omitempty"`
}

func (softlayer_billing_order_item_category_answer *SoftLayer_Billing_Order_Item_Category_Answer) String() string {
	return "SoftLayer_Billing_Order_Item_Category_Answer"
}
