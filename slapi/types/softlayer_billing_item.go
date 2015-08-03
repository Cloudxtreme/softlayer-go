package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Billing_Item - Every individual item that a SoftLayer customer is billed for is recorded
// in the SoftLayer_Billing_Item data type. Billing items range from server chassis to hard drives to
// control panels, bandwidth quota upgrades and port upgrade charges. Softlayer
// [[SoftLayer_Billing_Invoice|invoices]] are generated from the cost of a customer's billing items.
// Billing items are copied from the product catalog as they're ordered by customers to create a
// reference between an account and the billable items they own. Billing items exist in a tree
// relationship. Items are associated with each other by parent/child relationships. Component items
// such as CPU's, and software each have a parent billing item for the server chassis they're
// associated with. Billing Items with a null parent item do not have an associated parent item.
type SoftLayer_Billing_Item struct {

	// Notes - Extra information provided to help you identify this billing item. This is often a username
	// or something to help identify items that customers have more than one of.
	Notes string `json:"notes"`

	// CreateDate - The date the billing item was created. You can see this date on the invoice.
	CreateDate *time.Time `json:"createDate"`

	// CurrentHourlyCharge - This is the total charge for the billing item for this billing item. It is
	// calculated based on the hourlyRecurringFee * hoursUsed.
	CurrentHourlyCharge string `json:"currentHourlyCharge"`

	// ParentId - The unique identifier of the parent of this billing item.
	ParentId int `json:"parentId"`

	// LaborFeeTaxRate - The rate at which labor fees are taxed if you are a taxable customer.
	LaborFeeTaxRate float64 `json:"laborFeeTaxRate"`

	// OneTimeFeeTaxRate - The rate at which one time fees are taxed if you are a taxable customer.
	OneTimeFeeTaxRate float64 `json:"oneTimeFeeTaxRate"`

	// SetupFeeTaxRate - The rate at which setup fees are taxed if you are a taxable customer.
	SetupFeeTaxRate float64 `json:"setupFeeTaxRate"`

	// AssociatedBillingItemId - This is sometimes populated for orphan billing items that are not attached
	// to servers. Billing items like secondary portable IP addresses fit into this category. A user may
	// set an association by calling [[SoftLayer_Billing_Item::setAssociationId]]. This will cause this
	// orphan item to appear under its associated server billing item on future invoices. You may only
	// attach orphaned billing items to server billing items without cancellation dates set.
	AssociatedBillingItemId string `json:"associatedBillingItemId"`

	// LaborFee - no documentation
	LaborFee float64 `json:"laborFee"`

	// RecurringFee - The amount of money charged per month for a billing item, if applicable. recurringFee
	// is measured in US Dollars
	RecurringFee float64 `json:"recurringFee"`

	// SetupFee - no documentation
	SetupFee float64 `json:"setupFee"`

	// DomainName - The domain name is provided for server billing items.
	DomainName string `json:"domainName"`

	// CancellationDate - A billing item's cancellation date. A billing item with a cancellation date in
	// the past is not charged on your SoftLayer invoice. Cancellation dates in the future indicate the
	// current billing item is active, but will be cancelled and not charged for in the future. A billing
	// item with a null cancellation date is also considered an active billing item and is charged once
	// every billing cycle.
	CancellationDate *time.Time `json:"cancellationDate"`

	// RecurringFeeTaxRate - The rate at which recurring fees are taxed if you are a taxable customer.
	RecurringFeeTaxRate float64 `json:"recurringFeeTaxRate"`

	// CategoryCode - The category code of this billing item. It is used to tell us the difference between
	// a primary disk and a secondary disk, for instance.
	CategoryCode string `json:"categoryCode"`

	// HoursUsed - This is the number of hours the hourly billing item has been in use this billing period.
	// For virtual servers, this means running, paused or stopped.
	HoursUsed string `json:"hoursUsed"`

	// Id - no documentation
	Id int `json:"id"`

	// Description - no documentation
	Description string `json:"description"`

	// NextBillDate - The date on which your account will be charged for this billing item.
	NextBillDate *time.Time `json:"nextBillDate"`

	// CycleStartDate - no documentation
	CycleStartDate *time.Time `json:"cycleStartDate"`

	// OneTimeFee - The amount of money charged as a one-time charge for a billing item, if applicable.
	// oneTimeFee is measured in US Dollars
	OneTimeFee float64 `json:"oneTimeFee"`

	// HostName - no documentation
	HostName string `json:"hostName"`

	// HourlyRecurringFee - The amount of money charged per hour for a billing item, if applicable.
	// hourlyRecurringFee is measured in US Dollars
	HourlyRecurringFee float64 `json:"hourlyRecurringFee"`

	// ServiceProviderId - This is the service provider for this billing item.
	ServiceProviderId int `json:"serviceProviderId"`

	// AllowCancellationFlag - Flag to check if a billing item can be cancelled. 1 = yes. 0 = no.
	AllowCancellationFlag int `json:"allowCancellationFlag"`

	// LastBillDate - no documentation
	LastBillDate *time.Time `json:"lastBillDate"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// OrderItemId - the SoftLayer_Billing_Order_Item ID. This is a reference to the original order item
	// from which this billing item was originally created.
	OrderItemId int `json:"orderItemId"`

	// RecurringMonths - The number of months in which the recurring fees will be incurred.
	RecurringMonths int `json:"recurringMonths"`
}

// SoftLayer_Billing_Item_Extended is SoftLayer_Billing_Item with all maskable types.
type SoftLayer_Billing_Item_Extended struct {
	SoftLayer_Billing_Item

	// ActiveCancellationItem - A service cancellation request item that corresponds to the billing item.
	ActiveCancellationItem *SoftLayer_Billing_Item_Cancellation_Request_Item `json:"activeCancellationItem"`

	// Location - The location of the billing item. Some billing items have physical properties such as the
	// server itself. For items such as these, we provide location information.
	Location *SoftLayer_Location `json:"location"`

	// NextInvoiceChildren - A Billing Item's child billing items and associated items'
	NextInvoiceChildren []*SoftLayer_Billing_Item `json:"nextInvoiceChildren"`

	// ActiveChildrenCount - A count of a Billing Item's active child billing items.
	ActiveChildrenCount uint64 `json:"activeChildrenCount"`

	// ActiveBundledItems - no documentation
	ActiveBundledItems []*SoftLayer_Billing_Item `json:"activeBundledItems"`

	// BundledItems - no documentation
	BundledItems []*SoftLayer_Billing_Item `json:"bundledItems"`

	// HourlyFlag - A flag that will reflect whether this billing item is billed on an hourly basis or not.
	HourlyFlag bool `json:"hourlyFlag"`

	// AvailableMatchingVlans - <nil>
	AvailableMatchingVlans []*SoftLayer_Network_Vlan `json:"availableMatchingVlans"`

	// BundleItems - no documentation
	BundleItems []*SoftLayer_Product_Item_Bundles `json:"bundleItems"`

	// CancellationRequests - This will return any cancellation requests that are associated with this
	// billing item.
	CancellationRequests []*SoftLayer_Billing_Item_Cancellation_Request `json:"cancellationRequests"`

	// InvoiceItems - no documentation
	InvoiceItems []*SoftLayer_Billing_Invoice_Item `json:"invoiceItems"`

	// Item - The entry in the SoftLayer product catalog that a billing item is based upon.
	Item *SoftLayer_Product_Item `json:"item"`

	// ProvisionTransaction - no documentation
	ProvisionTransaction *SoftLayer_Provisioning_Version1_Transaction `json:"provisionTransaction"`

	// UpgradeItem - Billing items whose product item has an upgrade path defined in our system will return
	// the next product item in the upgrade path.
	UpgradeItem *SoftLayer_Product_Item `json:"upgradeItem"`

	// AssociatedChildrenCount - A count of a Billing Item's associated child billing items. This includes
	// "floating" items that are not necessarily child billing items of this billing item.
	AssociatedChildrenCount uint64 `json:"associatedChildrenCount"`

	// NonZeroNextInvoiceChildrenCount - A count of a Billing Item's associated child billing items,
	// excluding ALL items with a $0.00 recurring fee.
	NonZeroNextInvoiceChildrenCount uint64 `json:"nonZeroNextInvoiceChildrenCount"`

	// ActiveAgreementFlag - A flag indicating that the billing item is under an active agreement.
	ActiveAgreementFlag *SoftLayer_Account_Agreement `json:"activeAgreementFlag"`

	// AssociatedBillingItem - A billing item's associated parent. This is to be used for billing items
	// that are "floating", and therefore are not child items of any parent billing item. If it is desired
	// to associate an item to another, populate this with the SoftLayer_Billing_Item ID of that associated
	// parent item.
	AssociatedBillingItem *SoftLayer_Billing_Item `json:"associatedBillingItem"`

	// Category - The item category to which the billing item's item belongs.
	Category *SoftLayer_Product_Item_Category `json:"category"`

	// OriginalLocation - The original physical location for this billing item--may differ from current.
	OriginalLocation *SoftLayer_Location `json:"originalLocation"`

	// ParentVirtualGuestBillingItem - A billing item's parent item. If a billing item has no parent item
	// then this value is null.
	ParentVirtualGuestBillingItem *SoftLayer_Billing_Item_Virtual_Guest `json:"parentVirtualGuestBillingItem"`

	// PendingOrderItem - The new order item that will replace this billing item.
	PendingOrderItem *SoftLayer_Billing_Order_Item `json:"pendingOrderItem"`

	// CancellationReason - no documentation
	CancellationReason *SoftLayer_Billing_Item_Cancellation_Reason `json:"cancellationReason"`

	// NextInvoiceTotalRecurringTaxAmount - This is deprecated and will always be zero. Because tax is
	// calculated in real-time, previewing the next recurring invoice is pre-tax only.
	NextInvoiceTotalRecurringTaxAmount float32 `json:"nextInvoiceTotalRecurringTaxAmount"`

	// NonZeroNextInvoiceChildren - A Billing Item's associated child billing items, excluding ALL items
	// with a $0.00 recurring fee.
	NonZeroNextInvoiceChildren []*SoftLayer_Billing_Item `json:"nonZeroNextInvoiceChildren"`

	// ActiveSparePoolBundledItemCount - A count of a Billing Item's spare pool bundled billing items.
	ActiveSparePoolBundledItemCount uint64 `json:"activeSparePoolBundledItemCount"`

	// AvailableMatchingVlanCount - no documentation
	AvailableMatchingVlanCount uint64 `json:"availableMatchingVlanCount"`

	// FilteredNextInvoiceChildrenCount - A count of a Billing Item's associated child billing items,
	// excluding some items with a $0.00 recurring fee.
	FilteredNextInvoiceChildrenCount uint64 `json:"filteredNextInvoiceChildrenCount"`

	// ActiveChildren - no documentation
	ActiveChildren []*SoftLayer_Billing_Item `json:"activeChildren"`

	// AssociatedParent - A billing item's associated parent billing item. This object will be the same as
	// the parent billing item if parentId is set.
	AssociatedParent []*SoftLayer_Billing_Item `json:"associatedParent"`

	// InvoiceItem - no documentation
	InvoiceItem *SoftLayer_Billing_Invoice_Item `json:"invoiceItem"`

	// UpgradeItems - Billing items whose product item has an upgrade path defined in our system will
	// return all the product items in the upgrade path.
	UpgradeItems []*SoftLayer_Product_Item `json:"upgradeItems"`

	// BandwidthAllocation - no documentation
	BandwidthAllocation *SoftLayer_Network_Bandwidth_Version1_Allocation `json:"bandwidthAllocation"`

	// BillableChildren - A billing item's recurring child items that have once been billed and are
	// scheduled to be billed in the future.
	BillableChildren []*SoftLayer_Billing_Item `json:"billableChildren"`

	// ActiveAssociatedGuestDiskBillingItemCount - no documentation
	ActiveAssociatedGuestDiskBillingItemCount uint64 `json:"activeAssociatedGuestDiskBillingItemCount"`

	// BundledItemCount - no documentation
	BundledItemCount uint64 `json:"bundledItemCount"`

	// CancellationRequestCount - A count of this will return any cancellation requests that are associated
	// with this billing item.
	CancellationRequestCount uint64 `json:"cancellationRequestCount"`

	// NextInvoiceChildrenCount - A count of a Billing Item's child billing items and associated items'
	NextInvoiceChildrenCount uint64 `json:"nextInvoiceChildrenCount"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// ActiveFlag - <nil>
	ActiveFlag bool `json:"activeFlag"`

	// ActiveSparePoolBundledItems - no documentation
	ActiveSparePoolBundledItems []*SoftLayer_Billing_Item `json:"activeSparePoolBundledItems"`

	// CanceledChildren - no documentation
	CanceledChildren []*SoftLayer_Billing_Item `json:"canceledChildren"`

	// ChildrenWithActiveAgreement - no documentation
	ChildrenWithActiveAgreement []*SoftLayer_Billing_Item `json:"childrenWithActiveAgreement"`

	// PendingCancellationFlag - This flag indicates whether a billing item is scheduled to be canceled or
	// not.
	PendingCancellationFlag bool `json:"pendingCancellationFlag"`

	// BundleItemCount - no documentation
	BundleItemCount uint64 `json:"bundleItemCount"`

	// CanceledChildrenCount - A count of a Billing Item's active child billing items.
	CanceledChildrenCount uint64 `json:"canceledChildrenCount"`

	// DowngradeItemCount - A count of for product items which have a downgrade path defined, this will
	// return those product items.
	DowngradeItemCount uint64 `json:"downgradeItemCount"`

	// UpgradeItemCount - A count of billing items whose product item has an upgrade path defined in our
	// system will return all the product items in the upgrade path.
	UpgradeItemCount uint64 `json:"upgradeItemCount"`

	// Package - The package under which this billing item was sold. A Package is the general grouping of
	// products as seen on our order forms.
	Package *SoftLayer_Product_Package `json:"package"`

	// ActiveSparePoolAssociatedGuestDiskBillingItemCount - no documentation
	ActiveSparePoolAssociatedGuestDiskBillingItemCount uint64 `json:"activeSparePoolAssociatedGuestDiskBillingItemCount"`

	// DowngradeItems - For product items which have a downgrade path defined, this will return those
	// product items.
	DowngradeItems []*SoftLayer_Product_Item `json:"downgradeItems"`

	// NextInvoiceTotalRecurringAmount - A Billing Item's total, including any child billing items and
	// associated billing items if they exist.'
	NextInvoiceTotalRecurringAmount float32 `json:"nextInvoiceTotalRecurringAmount"`

	// AssociatedBillingItemHistoryCount - A count of a history of billing items which a billing item has
	// been associated with.
	AssociatedBillingItemHistoryCount uint64 `json:"associatedBillingItemHistoryCount"`

	// ChildrenCount - no documentation
	ChildrenCount uint64 `json:"childrenCount"`

	// AssociatedBillingItemHistory - A history of billing items which a billing item has been associated
	// with.
	AssociatedBillingItemHistory []*SoftLayer_Billing_Item_Association_History `json:"associatedBillingItemHistory"`

	// SoftwareDescription - no documentation
	SoftwareDescription *SoftLayer_Software_Description `json:"softwareDescription"`

	// ChildrenWithActiveAgreementCount - A count of a Billing Item's active child billing items.
	ChildrenWithActiveAgreementCount uint64 `json:"childrenWithActiveAgreementCount"`

	// ActiveAssociatedGuestDiskBillingItems - <nil>
	ActiveAssociatedGuestDiskBillingItems []*SoftLayer_Billing_Item `json:"activeAssociatedGuestDiskBillingItems"`

	// AssociatedChildren - A Billing Item's associated child billing items. This includes "floating" items
	// that are not necessarily child billing items of this billing item.
	AssociatedChildren []*SoftLayer_Billing_Item `json:"associatedChildren"`

	// Children - no documentation
	Children []*SoftLayer_Billing_Item `json:"children"`

	// NextInvoiceTotalOneTimeAmount - A Billing Item's total, including any child billing items if they
	// exist.'
	NextInvoiceTotalOneTimeAmount float32 `json:"nextInvoiceTotalOneTimeAmount"`

	// AssociatedParentCount - A count of a billing item's associated parent billing item. This object will
	// be the same as the parent billing item if parentId is set.
	AssociatedParentCount uint64 `json:"associatedParentCount"`

	// ActiveAssociatedChildren - A billing item's active associated child billing items. This includes
	// "floating" items that are not necessarily child items of this billing item.
	ActiveAssociatedChildren []*SoftLayer_Billing_Item `json:"activeAssociatedChildren"`

	// ActiveSparePoolAssociatedGuestDiskBillingItems - <nil>
	ActiveSparePoolAssociatedGuestDiskBillingItems []*SoftLayer_Billing_Item `json:"activeSparePoolAssociatedGuestDiskBillingItems"`

	// FilteredNextInvoiceChildren - A Billing Item's associated child billing items, excluding some items
	// with a $0.00 recurring fee.
	FilteredNextInvoiceChildren []*SoftLayer_Billing_Item `json:"filteredNextInvoiceChildren"`

	// NextInvoiceTotalOneTimeTaxAmount - A Billing Item's total, including any child billing items if they
	// exist.'
	NextInvoiceTotalOneTimeTaxAmount float32 `json:"nextInvoiceTotalOneTimeTaxAmount"`

	// OrderItem - A billing item's original order item. Simply a reference to the original order from
	// which this billing item was created.
	OrderItem *SoftLayer_Billing_Order_Item `json:"orderItem"`

	// ActiveBundledItemCount - A count of a Billing Item's active bundled billing items.
	ActiveBundledItemCount uint64 `json:"activeBundledItemCount"`

	// InvoiceItemCount - A count of all invoice items associated with the billing item
	InvoiceItemCount uint64 `json:"invoiceItemCount"`

	// ActiveAgreement - <nil>
	ActiveAgreement *SoftLayer_Account_Agreement `json:"activeAgreement"`

	// Parent - A billing item's parent item. If a billing item has no parent item then this value is null.
	Parent *SoftLayer_Billing_Item `json:"parent"`

	// ActiveAssociatedChildrenCount - A count of a billing item's active associated child billing items.
	// This includes "floating" items that are not necessarily child items of this billing item.
	ActiveAssociatedChildrenCount uint64 `json:"activeAssociatedChildrenCount"`

	// BillableChildrenCount - A count of a billing item's recurring child items that have once been billed
	// and are scheduled to be billed in the future.
	BillableChildrenCount uint64 `json:"billableChildrenCount"`
}

func (softlayer_billing_item *SoftLayer_Billing_Item) String() string {
	return "SoftLayer_Billing_Item"
}
