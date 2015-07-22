package sl

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

	// OrderItem - no documentation
	OrderItem *SoftLayer_Billing_Order_Item `json:"orderItem"`

	// Question - no documentation
	Question *SoftLayer_Product_Item_Category_Question `json:"question"`

	// QuestionId - The identifier for the question that the answer belongs to.
	QuestionId int `json:"questionId"`
}
