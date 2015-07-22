package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Billing_Item_Association_History - The SoftLayer_Billing_Item_Association_History type
// keeps a record of which server billing items an "orphan" item has been associated with. Orphan
// billing items are billable items for secondary portable services (such as secondary subnets and
// StorageLayer accounts) that are not associated with a server and appear at the bottom of a SoftLayer
// invoice. The [[SoftLayer_Billing_Item::setAssociationId]] method allows you to associate these kinds
// of items with servers, making them appear as a child item of the server on your invoice. A
// SoftLayer_Billing_Item_Association_History record is created every time one of these associations
// are set.
type SoftLayer_Billing_Item_Association_History struct {

	// AssociatedBillingItem - The server billing item that an orphaned billing item was associated with.
	AssociatedBillingItem *SoftLayer_Billing_Item `json:"associatedBillingItem"`

	// AssociatedBillingItemId - The internal identifier of the server billing item that an orphaned
	// billing item was associated with.
	AssociatedBillingItemId int `json:"associatedBillingItemId"`

	// BillingItem - The billing item that was associated with a server billing item.
	BillingItem *SoftLayer_Billing_Item `json:"billingItem"`

	// BillingItemId - The internal identifier of the billing item that was associated with a server
	// billing item.
	BillingItemId int `json:"billingItemId"`

	// CreateDate - The date that a billing item association was last changed.
	CreateDate *time.Time `json:"createDate"`

	// Id - A billing item association history's internal identifier.
	Id int `json:"id"`
}