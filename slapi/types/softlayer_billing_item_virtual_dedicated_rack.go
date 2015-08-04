package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Billing_Item_Virtual_Dedicated_Rack - A SoftLayer_Billing_Item_Virtual_Dedicated_Rack data
// type models the billing information for a single bandwidth pooling. Bandwidth pooling members share
// their public bandwidth allocations, and incur overage charges instead of the overages on individual
// rack members. Virtual rack billing items are the parent items for all of it's rack membership
// billing items.
type SoftLayer_Billing_Item_Virtual_Dedicated_Rack struct {

	// CategoryCode - The category code of this billing item. It is used to tell us the difference between
	// a primary disk and a secondary disk, for instance.
	CategoryCode string `json:"categoryCode,omitempty"`

	// HoursUsed - This is the number of hours the hourly billing item has been in use this billing period.
	// For virtual servers, this means running, paused or stopped.
	HoursUsed string `json:"hoursUsed,omitempty"`

	// AssociatedBillingItemId - This is sometimes populated for orphan billing items that are not attached
	// to servers. Billing items like secondary portable IP addresses fit into this category. A user may
	// set an association by calling [[SoftLayer_Billing_Item::setAssociationId]]. This will cause this
	// orphan item to appear under its associated server billing item on future invoices. You may only
	// attach orphaned billing items to server billing items without cancellation dates set.
	AssociatedBillingItemId string `json:"associatedBillingItemId,omitempty"`

	// OneTimeFeeTaxRate - The rate at which one time fees are taxed if you are a taxable customer.
	OneTimeFeeTaxRate float64 `json:"oneTimeFeeTaxRate,omitempty"`

	// LastBillDate - no documentation
	LastBillDate *time.Time `json:"lastBillDate,omitempty"`

	// HourlyRecurringFee - The amount of money charged per hour for a billing item, if applicable.
	// hourlyRecurringFee is measured in US Dollars
	HourlyRecurringFee float64 `json:"hourlyRecurringFee,omitempty"`

	// Notes - Extra information provided to help you identify this billing item. This is often a username
	// or something to help identify items that customers have more than one of.
	Notes string `json:"notes,omitempty"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// ServiceProviderId - This is the service provider for this billing item.
	ServiceProviderId int `json:"serviceProviderId,omitempty"`

	// CycleStartDate - no documentation
	CycleStartDate *time.Time `json:"cycleStartDate,omitempty"`

	// ParentId - The unique identifier of the parent of this billing item.
	ParentId int `json:"parentId,omitempty"`

	// OrderItemId - the SoftLayer_Billing_Order_Item ID. This is a reference to the original order item
	// from which this billing item was originally created.
	OrderItemId int `json:"orderItemId,omitempty"`

	// CreateDate - The date the billing item was created. You can see this date on the invoice.
	CreateDate *time.Time `json:"createDate,omitempty"`

	// SetupFee - no documentation
	SetupFee float64 `json:"setupFee,omitempty"`

	// AllowCancellationFlag - Flag to check if a billing item can be cancelled. 1 = yes. 0 = no.
	AllowCancellationFlag int `json:"allowCancellationFlag,omitempty"`

	// LaborFee - no documentation
	LaborFee float64 `json:"laborFee,omitempty"`

	// CurrentHourlyCharge - This is the total charge for the billing item for this billing item. It is
	// calculated based on the hourlyRecurringFee * hoursUsed.
	CurrentHourlyCharge string `json:"currentHourlyCharge,omitempty"`

	// SetupFeeTaxRate - The rate at which setup fees are taxed if you are a taxable customer.
	SetupFeeTaxRate float64 `json:"setupFeeTaxRate,omitempty"`

	// Description - no documentation
	Description string `json:"description,omitempty"`

	// LaborFeeTaxRate - The rate at which labor fees are taxed if you are a taxable customer.
	LaborFeeTaxRate float64 `json:"laborFeeTaxRate,omitempty"`

	// NextBillDate - The date on which your account will be charged for this billing item.
	NextBillDate *time.Time `json:"nextBillDate,omitempty"`

	// DomainName - The domain name is provided for server billing items.
	DomainName string `json:"domainName,omitempty"`

	// CancellationDate - A billing item's cancellation date. A billing item with a cancellation date in
	// the past is not charged on your SoftLayer invoice. Cancellation dates in the future indicate the
	// current billing item is active, but will be cancelled and not charged for in the future. A billing
	// item with a null cancellation date is also considered an active billing item and is charged once
	// every billing cycle.
	CancellationDate *time.Time `json:"cancellationDate,omitempty"`

	// RecurringFee - The amount of money charged per month for a billing item, if applicable. recurringFee
	// is measured in US Dollars
	RecurringFee float64 `json:"recurringFee,omitempty"`

	// RecurringFeeTaxRate - The rate at which recurring fees are taxed if you are a taxable customer.
	RecurringFeeTaxRate float64 `json:"recurringFeeTaxRate,omitempty"`

	// HostName - no documentation
	HostName string `json:"hostName,omitempty"`

	// RecurringMonths - The number of months in which the recurring fees will be incurred.
	RecurringMonths int `json:"recurringMonths,omitempty"`

	// OneTimeFee - The amount of money charged as a one-time charge for a billing item, if applicable.
	// oneTimeFee is measured in US Dollars
	OneTimeFee float64 `json:"oneTimeFee,omitempty"`
}

func (softlayer_billing_item_virtual_dedicated_rack *SoftLayer_Billing_Item_Virtual_Dedicated_Rack) String() string {
	return "SoftLayer_Billing_Item_Virtual_Dedicated_Rack"
}

// SoftLayer_Billing_Item_Virtual_Dedicated_Rack_Extended is SoftLayer_Billing_Item_Virtual_Dedicated_Rack with all maskable types.
type SoftLayer_Billing_Item_Virtual_Dedicated_Rack_Extended struct {
	SoftLayer_Billing_Item_Virtual_Dedicated_Rack

	// ActiveSparePoolBundledItemCount - A count of a Billing Item's spare pool bundled billing items.
	ActiveSparePoolBundledItemCount uint64 `json:"activeSparePoolBundledItemCount,omitempty"`

	// AssociatedBillingItem - A billing item's associated parent. This is to be used for billing items
	// that are "floating", and therefore are not child items of any parent billing item. If it is desired
	// to associate an item to another, populate this with the SoftLayer_Billing_Item ID of that associated
	// parent item.
	AssociatedBillingItem *SoftLayer_Billing_Item `json:"associatedBillingItem,omitempty"`

	// BillingCycleBandwidthUsage - The raw bandwidth usage data for the current billing cycle. One object
	// is returned for each network a virtual rack is attached to.
	BillingCycleBandwidthUsage []*SoftLayer_Network_Bandwidth_Usage `json:"billingCycleBandwidthUsage,omitempty"`

	// AvailableMatchingVlanCount - no documentation
	AvailableMatchingVlanCount uint64 `json:"availableMatchingVlanCount,omitempty"`

	// FilteredNextInvoiceChildrenCount - A count of a Billing Item's associated child billing items,
	// excluding some items with a $0.00 recurring fee.
	FilteredNextInvoiceChildrenCount uint64 `json:"filteredNextInvoiceChildrenCount,omitempty"`

	// AssociatedChildrenCount - A count of a Billing Item's associated child billing items. This includes
	// "floating" items that are not necessarily child billing items of this billing item.
	AssociatedChildrenCount uint64 `json:"associatedChildrenCount,omitempty"`

	// BillingCyclePublicBandwidthUsageCount - A count of the raw public bandwidth usage data for the
	// current billing cycle.
	BillingCyclePublicBandwidthUsageCount uint64 `json:"billingCyclePublicBandwidthUsageCount,omitempty"`

	// ActiveSparePoolBundledItems - no documentation
	ActiveSparePoolBundledItems []*SoftLayer_Billing_Item `json:"activeSparePoolBundledItems,omitempty"`

	// BundleItemCount - no documentation
	BundleItemCount uint64 `json:"bundleItemCount,omitempty"`

	// Parent - A billing item's parent item. If a billing item has no parent item then this value is null.
	Parent *SoftLayer_Billing_Item `json:"parent,omitempty"`

	// NonZeroNextInvoiceChildrenCount - A count of a Billing Item's associated child billing items,
	// excluding ALL items with a $0.00 recurring fee.
	NonZeroNextInvoiceChildrenCount uint64 `json:"nonZeroNextInvoiceChildrenCount,omitempty"`

	// ChildrenWithActiveAgreement - no documentation
	ChildrenWithActiveAgreement []*SoftLayer_Billing_Item `json:"childrenWithActiveAgreement,omitempty"`

	// AssociatedBillingItemHistory - A history of billing items which a billing item has been associated
	// with.
	AssociatedBillingItemHistory []*SoftLayer_Billing_Item_Association_History `json:"associatedBillingItemHistory,omitempty"`

	// Children - no documentation
	Children []*SoftLayer_Billing_Item `json:"children,omitempty"`

	// InvoiceItemCount - A count of all invoice items associated with the billing item
	InvoiceItemCount uint64 `json:"invoiceItemCount,omitempty"`

	// OriginalLocation - The original physical location for this billing item--may differ from current.
	OriginalLocation *SoftLayer_Location `json:"originalLocation,omitempty"`

	// CanceledChildrenCount - A count of a Billing Item's active child billing items.
	CanceledChildrenCount uint64 `json:"canceledChildrenCount,omitempty"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account,omitempty"`

	// NextInvoiceTotalRecurringAmount - A Billing Item's total, including any child billing items and
	// associated billing items if they exist.'
	NextInvoiceTotalRecurringAmount float32 `json:"nextInvoiceTotalRecurringAmount,omitempty"`

	// ChildrenCount - no documentation
	ChildrenCount uint64 `json:"childrenCount,omitempty"`

	// ActiveSparePoolAssociatedGuestDiskBillingItems - <nil>
	ActiveSparePoolAssociatedGuestDiskBillingItems []*SoftLayer_Billing_Item `json:"activeSparePoolAssociatedGuestDiskBillingItems,omitempty"`

	// BillingCyclePublicUsageOut - The total public outbound bandwidth for this virtual rack for the
	// current billing cycle.
	BillingCyclePublicUsageOut float64 `json:"billingCyclePublicUsageOut,omitempty"`

	// BandwidthAllocation - no documentation
	BandwidthAllocation *SoftLayer_Network_Bandwidth_Version1_Allocation `json:"bandwidthAllocation,omitempty"`

	// BillingCyclePublicBandwidthUsage - The raw public bandwidth usage data for the current billing
	// cycle.
	BillingCyclePublicBandwidthUsage []*SoftLayer_Network_Bandwidth_Usage `json:"billingCyclePublicBandwidthUsage,omitempty"`

	// ParentVirtualGuestBillingItem - A billing item's parent item. If a billing item has no parent item
	// then this value is null.
	ParentVirtualGuestBillingItem *SoftLayer_Billing_Item_Virtual_Guest `json:"parentVirtualGuestBillingItem,omitempty"`

	// ActiveChildrenCount - A count of a Billing Item's active child billing items.
	ActiveChildrenCount uint64 `json:"activeChildrenCount,omitempty"`

	// Item - The entry in the SoftLayer product catalog that a billing item is based upon.
	Item *SoftLayer_Product_Item `json:"item,omitempty"`

	// ProvisionTransaction - no documentation
	ProvisionTransaction *SoftLayer_Provisioning_Version1_Transaction `json:"provisionTransaction,omitempty"`

	// ActiveBundledItems - no documentation
	ActiveBundledItems []*SoftLayer_Billing_Item `json:"activeBundledItems,omitempty"`

	// Category - The item category to which the billing item's item belongs.
	Category *SoftLayer_Product_Item_Category `json:"category,omitempty"`

	// NextInvoiceTotalOneTimeAmount - A Billing Item's total, including any child billing items if they
	// exist.'
	NextInvoiceTotalOneTimeAmount float32 `json:"nextInvoiceTotalOneTimeAmount,omitempty"`

	// SoftwareDescription - no documentation
	SoftwareDescription *SoftLayer_Software_Description `json:"softwareDescription,omitempty"`

	// BillingCyclePrivateUsageIn - The total private network inbound bandwidth for this virtual rack for
	// the current billing cycle.
	BillingCyclePrivateUsageIn float64 `json:"billingCyclePrivateUsageIn,omitempty"`

	// AssociatedParentCount - A count of a billing item's associated parent billing item. This object will
	// be the same as the parent billing item if parentId is set.
	AssociatedParentCount uint64 `json:"associatedParentCount,omitempty"`

	// CancellationRequests - This will return any cancellation requests that are associated with this
	// billing item.
	CancellationRequests []*SoftLayer_Billing_Item_Cancellation_Request `json:"cancellationRequests,omitempty"`

	// BillingCyclePublicUsageIn - The total public inbound bandwidth for this virtual rack for the current
	// billing cycle.
	BillingCyclePublicUsageIn float64 `json:"billingCyclePublicUsageIn,omitempty"`

	// BundledItems - no documentation
	BundledItems []*SoftLayer_Billing_Item `json:"bundledItems,omitempty"`

	// BillableChildren - A billing item's recurring child items that have once been billed and are
	// scheduled to be billed in the future.
	BillableChildren []*SoftLayer_Billing_Item `json:"billableChildren,omitempty"`

	// CancellationRequestCount - A count of this will return any cancellation requests that are associated
	// with this billing item.
	CancellationRequestCount uint64 `json:"cancellationRequestCount,omitempty"`

	// BillingCyclePrivateUsageTotal - The total private network bandwidth for this virtual rack for the
	// current billing cycle.
	BillingCyclePrivateUsageTotal uint `json:"billingCyclePrivateUsageTotal,omitempty"`

	// ActiveAssociatedGuestDiskBillingItemCount - no documentation
	ActiveAssociatedGuestDiskBillingItemCount uint64 `json:"activeAssociatedGuestDiskBillingItemCount,omitempty"`

	// BundledItemCount - no documentation
	BundledItemCount uint64 `json:"bundledItemCount,omitempty"`

	// CancellationReason - no documentation
	CancellationReason *SoftLayer_Billing_Item_Cancellation_Reason `json:"cancellationReason,omitempty"`

	// BillingCyclePrivateUsageOut - The total private network outbound bandwidth for this virtual rack for
	// the current billing cycle.
	BillingCyclePrivateUsageOut float64 `json:"billingCyclePrivateUsageOut,omitempty"`

	// ActiveChildren - no documentation
	ActiveChildren []*SoftLayer_Billing_Item `json:"activeChildren,omitempty"`

	// ActiveCancellationItem - A service cancellation request item that corresponds to the billing item.
	ActiveCancellationItem *SoftLayer_Billing_Item_Cancellation_Request_Item `json:"activeCancellationItem,omitempty"`

	// NextInvoiceChildren - A Billing Item's child billing items and associated items'
	NextInvoiceChildren []*SoftLayer_Billing_Item `json:"nextInvoiceChildren,omitempty"`

	// CanceledChildren - no documentation
	CanceledChildren []*SoftLayer_Billing_Item `json:"canceledChildren,omitempty"`

	// ActiveFlag - <nil>
	ActiveFlag bool `json:"activeFlag,omitempty"`

	// UpgradeItemCount - A count of billing items whose product item has an upgrade path defined in our
	// system will return all the product items in the upgrade path.
	UpgradeItemCount uint64 `json:"upgradeItemCount,omitempty"`

	// OrderItem - A billing item's original order item. Simply a reference to the original order from
	// which this billing item was created.
	OrderItem *SoftLayer_Billing_Order_Item `json:"orderItem,omitempty"`

	// Resource - The virtual rack that a virtual rack billing item is associated with.
	Resource *SoftLayer_Network_Bandwidth_Version1_Allotment `json:"resource,omitempty"`

	// Location - The location of the billing item. Some billing items have physical properties such as the
	// server itself. For items such as these, we provide location information.
	Location *SoftLayer_Location `json:"location,omitempty"`

	// HourlyFlag - A flag that will reflect whether this billing item is billed on an hourly basis or not.
	HourlyFlag bool `json:"hourlyFlag,omitempty"`

	// BillingCyclePrivateBandwidthUsageCount - A count of the raw private bandwidth usage data for the
	// current billing cycle.
	BillingCyclePrivateBandwidthUsageCount uint64 `json:"billingCyclePrivateBandwidthUsageCount,omitempty"`

	// UpgradeItems - Billing items whose product item has an upgrade path defined in our system will
	// return all the product items in the upgrade path.
	UpgradeItems []*SoftLayer_Product_Item `json:"upgradeItems,omitempty"`

	// ActiveAssociatedChildren - A billing item's active associated child billing items. This includes
	// "floating" items that are not necessarily child items of this billing item.
	ActiveAssociatedChildren []*SoftLayer_Billing_Item `json:"activeAssociatedChildren,omitempty"`

	// ActiveAgreementFlag - A flag indicating that the billing item is under an active agreement.
	ActiveAgreementFlag *SoftLayer_Account_Agreement `json:"activeAgreementFlag,omitempty"`

	// AvailableMatchingVlans - <nil>
	AvailableMatchingVlans []*SoftLayer_Network_Vlan `json:"availableMatchingVlans,omitempty"`

	// BundleItems - no documentation
	BundleItems []*SoftLayer_Product_Item_Bundles `json:"bundleItems,omitempty"`

	// DowngradeItemCount - A count of for product items which have a downgrade path defined, this will
	// return those product items.
	DowngradeItemCount uint64 `json:"downgradeItemCount,omitempty"`

	// NextInvoiceTotalRecurringTaxAmount - This is deprecated and will always be zero. Because tax is
	// calculated in real-time, previewing the next recurring invoice is pre-tax only.
	NextInvoiceTotalRecurringTaxAmount float32 `json:"nextInvoiceTotalRecurringTaxAmount,omitempty"`

	// UpgradeItem - Billing items whose product item has an upgrade path defined in our system will return
	// the next product item in the upgrade path.
	UpgradeItem *SoftLayer_Product_Item `json:"upgradeItem,omitempty"`

	// FilteredNextInvoiceChildren - A Billing Item's associated child billing items, excluding some items
	// with a $0.00 recurring fee.
	FilteredNextInvoiceChildren []*SoftLayer_Billing_Item `json:"filteredNextInvoiceChildren,omitempty"`

	// InvoiceItems - no documentation
	InvoiceItems []*SoftLayer_Billing_Invoice_Item `json:"invoiceItems,omitempty"`

	// AssociatedParent - A billing item's associated parent billing item. This object will be the same as
	// the parent billing item if parentId is set.
	AssociatedParent []*SoftLayer_Billing_Item `json:"associatedParent,omitempty"`

	// BillableChildrenCount - A count of a billing item's recurring child items that have once been billed
	// and are scheduled to be billed in the future.
	BillableChildrenCount uint64 `json:"billableChildrenCount,omitempty"`

	// InvoiceItem - no documentation
	InvoiceItem *SoftLayer_Billing_Invoice_Item `json:"invoiceItem,omitempty"`

	// AssociatedBillingItemHistoryCount - A count of a history of billing items which a billing item has
	// been associated with.
	AssociatedBillingItemHistoryCount uint64 `json:"associatedBillingItemHistoryCount,omitempty"`

	// PendingOrderItem - The new order item that will replace this billing item.
	PendingOrderItem *SoftLayer_Billing_Order_Item `json:"pendingOrderItem,omitempty"`

	// ActiveAssociatedGuestDiskBillingItems - <nil>
	ActiveAssociatedGuestDiskBillingItems []*SoftLayer_Billing_Item `json:"activeAssociatedGuestDiskBillingItems,omitempty"`

	// BillingCyclePrivateBandwidthUsage - The raw private bandwidth usage data for the current billing
	// cycle.
	BillingCyclePrivateBandwidthUsage []*SoftLayer_Network_Bandwidth_Usage `json:"billingCyclePrivateBandwidthUsage,omitempty"`

	// ChildrenWithActiveAgreementCount - A count of a Billing Item's active child billing items.
	ChildrenWithActiveAgreementCount uint64 `json:"childrenWithActiveAgreementCount,omitempty"`

	// BillingCycleBandwidthUsageCount - A count of the raw bandwidth usage data for the current billing
	// cycle. One object is returned for each network a virtual rack is attached to.
	BillingCycleBandwidthUsageCount uint64 `json:"billingCycleBandwidthUsageCount,omitempty"`

	// AssociatedChildren - A Billing Item's associated child billing items. This includes "floating" items
	// that are not necessarily child billing items of this billing item.
	AssociatedChildren []*SoftLayer_Billing_Item `json:"associatedChildren,omitempty"`

	// ActiveBundledItemCount - A count of a Billing Item's active bundled billing items.
	ActiveBundledItemCount uint64 `json:"activeBundledItemCount,omitempty"`

	// BillingCyclePublicUsageTotal - The total public bandwidth for this virtual rack for the current
	// billing cycle.
	BillingCyclePublicUsageTotal uint `json:"billingCyclePublicUsageTotal,omitempty"`

	// ActiveSparePoolAssociatedGuestDiskBillingItemCount - no documentation
	ActiveSparePoolAssociatedGuestDiskBillingItemCount uint64 `json:"activeSparePoolAssociatedGuestDiskBillingItemCount,omitempty"`

	// NextInvoiceChildrenCount - A count of a Billing Item's child billing items and associated items'
	NextInvoiceChildrenCount uint64 `json:"nextInvoiceChildrenCount,omitempty"`

	// ActiveAgreement - <nil>
	ActiveAgreement *SoftLayer_Account_Agreement `json:"activeAgreement,omitempty"`

	// PendingCancellationFlag - This flag indicates whether a billing item is scheduled to be canceled or
	// not.
	PendingCancellationFlag bool `json:"pendingCancellationFlag,omitempty"`

	// NextInvoiceTotalOneTimeTaxAmount - A Billing Item's total, including any child billing items if they
	// exist.'
	NextInvoiceTotalOneTimeTaxAmount float32 `json:"nextInvoiceTotalOneTimeTaxAmount,omitempty"`

	// ActiveAssociatedChildrenCount - A count of a billing item's active associated child billing items.
	// This includes "floating" items that are not necessarily child items of this billing item.
	ActiveAssociatedChildrenCount uint64 `json:"activeAssociatedChildrenCount,omitempty"`

	// DowngradeItems - For product items which have a downgrade path defined, this will return those
	// product items.
	DowngradeItems []*SoftLayer_Product_Item `json:"downgradeItems,omitempty"`

	// Package - The package under which this billing item was sold. A Package is the general grouping of
	// products as seen on our order forms.
	Package *SoftLayer_Product_Package `json:"package,omitempty"`

	// NonZeroNextInvoiceChildren - A Billing Item's associated child billing items, excluding ALL items
	// with a $0.00 recurring fee.
	NonZeroNextInvoiceChildren []*SoftLayer_Billing_Item `json:"nonZeroNextInvoiceChildren,omitempty"`
}

func (softlayer_billing_item_virtual_dedicated_rack *SoftLayer_Billing_Item_Virtual_Dedicated_Rack_Extended) String() string {
	return "SoftLayer_Billing_Item_Virtual_Dedicated_Rack"
}
