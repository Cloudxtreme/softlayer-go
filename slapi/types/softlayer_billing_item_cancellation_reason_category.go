package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Billing_Item_Cancellation_Reason_Category - The
// SoftLayer_Billing_Item_Cancellation_Reason_Category data type contains cancellation reason
// categories.
type SoftLayer_Billing_Item_Cancellation_Reason_Category struct {

	// BillingCancellationReasonCount - A count of the corresponding billing cancellation reasons having
	// the specific billing cancellation reason category.
	BillingCancellationReasonCount uint64 `json:"billingCancellationReasonCount"`

	// BillingCancellationReasons - The corresponding billing cancellation reasons having the specific
	// billing cancellation reason category.
	BillingCancellationReasons []*SoftLayer_Billing_Item_Cancellation_Reason `json:"billingCancellationReasons"`

	// Id - no documentation
	Id int `json:"id"`

	// Name - no documentation
	Name string `json:"name"`
}

// GetAllCancellationReasonCategories - getAllCancellationReasonCategories() retrieves a list of all
// cancellation reason categories
func (softlayer_billing_item_cancellation_reason_category *SoftLayer_Billing_Item_Cancellation_Reason_Category) GetAllCancellationReasonCategories(commonOptions *slapi.CommonOptions) ([]*SoftLayer_Billing_Item_Cancellation_Reason_Category, error) {
	var returnValue []*SoftLayer_Billing_Item_Cancellation_Reason_Category
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_billing_item_cancellation_reason_category *SoftLayer_Billing_Item_Cancellation_Reason_Category) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Billing_Item_Cancellation_Reason_Category, error) {
	var returnValue *SoftLayer_Billing_Item_Cancellation_Reason_Category
	return returnValue, nil
}
