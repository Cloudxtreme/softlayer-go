package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"

	time "time"
)

// SoftLayer_Billing_Item_Virtual_Dedicated_Rack - A SoftLayer_Billing_Item_Virtual_Dedicated_Rack data
// type models the billing information for a single bandwidth pooling. Bandwidth pooling members share
// their public bandwidth allocations, and incur overage charges instead of the overages on individual
// rack members. Virtual rack billing items are the parent items for all of it's rack membership
// billing items.
type SoftLayer_Billing_Item_Virtual_Dedicated_Rack struct {

	// CancellationDate - A billing item's cancellation date. A billing item with a cancellation date in
	// the past is not charged on your SoftLayer invoice. Cancellation dates in the future indicate the
	// current billing item is active, but will be cancelled and not charged for in the future. A billing
	// item with a null cancellation date is also considered an active billing item and is charged once
	// every billing cycle.
	CancellationDate *time.Time `json:"cancellationDate,omitempty"`

	// LaborFeeTaxRate - The rate at which labor fees are taxed if you are a taxable customer.
	LaborFeeTaxRate slapi.Float64 `json:"laborFeeTaxRate,omitempty"`

	// HostName - no documentation
	HostName string `json:"hostName,omitempty"`

	// LaborFee - no documentation
	LaborFee slapi.Float64 `json:"laborFee,omitempty"`

	// OrderItemId - the SoftLayer_Billing_Order_Item ID. This is a reference to the original order item
	// from which this billing item was originally created.
	OrderItemId int `json:"orderItemId,omitempty"`

	// AssociatedBillingItemId - This is sometimes populated for orphan billing items that are not attached
	// to servers. Billing items like secondary portable IP addresses fit into this category. A user may
	// set an association by calling [[SoftLayer_Billing_Item::setAssociationId]]. This will cause this
	// orphan item to appear under its associated server billing item on future invoices. You may only
	// attach orphaned billing items to server billing items without cancellation dates set.
	AssociatedBillingItemId string `json:"associatedBillingItemId,omitempty"`

	// SetupFee - no documentation
	SetupFee slapi.Float64 `json:"setupFee,omitempty"`

	// CreateDate - The date the billing item was created. You can see this date on the invoice.
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Description - no documentation
	Description string `json:"description,omitempty"`

	// CategoryCode - The category code of this billing item. It is used to tell us the difference between
	// a primary disk and a secondary disk, for instance.
	CategoryCode string `json:"categoryCode,omitempty"`

	// SetupFeeTaxRate - The rate at which setup fees are taxed if you are a taxable customer.
	SetupFeeTaxRate slapi.Float64 `json:"setupFeeTaxRate,omitempty"`

	// ServiceProviderId - This is the service provider for this billing item.
	ServiceProviderId int `json:"serviceProviderId,omitempty"`

	// HoursUsed - This is the number of hours the hourly billing item has been in use this billing period.
	// For virtual servers, this means running, paused or stopped.
	HoursUsed string `json:"hoursUsed,omitempty"`

	// RecurringFeeTaxRate - The rate at which recurring fees are taxed if you are a taxable customer.
	RecurringFeeTaxRate slapi.Float64 `json:"recurringFeeTaxRate,omitempty"`

	// OneTimeFee - The amount of money charged as a one-time charge for a billing item, if applicable.
	// oneTimeFee is measured in US Dollars
	OneTimeFee slapi.Float64 `json:"oneTimeFee,omitempty"`

	// LastBillDate - no documentation
	LastBillDate *time.Time `json:"lastBillDate,omitempty"`

	// OneTimeFeeTaxRate - The rate at which one time fees are taxed if you are a taxable customer.
	OneTimeFeeTaxRate slapi.Float64 `json:"oneTimeFeeTaxRate,omitempty"`

	// CurrentHourlyCharge - This is the total charge for the billing item for this billing item. It is
	// calculated based on the hourlyRecurringFee * hoursUsed.
	CurrentHourlyCharge string `json:"currentHourlyCharge,omitempty"`

	// NextBillDate - The date on which your account will be charged for this billing item.
	NextBillDate *time.Time `json:"nextBillDate,omitempty"`

	// HourlyRecurringFee - The amount of money charged per hour for a billing item, if applicable.
	// hourlyRecurringFee is measured in US Dollars
	HourlyRecurringFee slapi.Float64 `json:"hourlyRecurringFee,omitempty"`

	// RecurringMonths - The number of months in which the recurring fees will be incurred.
	RecurringMonths int `json:"recurringMonths,omitempty"`

	// ParentId - The unique identifier of the parent of this billing item.
	ParentId int `json:"parentId,omitempty"`

	// RecurringFee - The amount of money charged per month for a billing item, if applicable. recurringFee
	// is measured in US Dollars
	RecurringFee slapi.Float64 `json:"recurringFee,omitempty"`

	// AllowCancellationFlag - Flag to check if a billing item can be cancelled. 1 = yes. 0 = no.
	AllowCancellationFlag int `json:"allowCancellationFlag,omitempty"`

	// Notes - Extra information provided to help you identify this billing item. This is often a username
	// or something to help identify items that customers have more than one of.
	Notes string `json:"notes,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// DomainName - The domain name is provided for server billing items.
	DomainName string `json:"domainName,omitempty"`

	// CycleStartDate - no documentation
	CycleStartDate *time.Time `json:"cycleStartDate,omitempty"`
}

func (softlayer_billing_item_virtual_dedicated_rack *SoftLayer_Billing_Item_Virtual_Dedicated_Rack) String() string {
	return "SoftLayer_Billing_Item_Virtual_Dedicated_Rack"
}

// SoftLayer_Billing_Item_Virtual_Dedicated_Rack_Extended is SoftLayer_Billing_Item_Virtual_Dedicated_Rack with all maskable types.
type SoftLayer_Billing_Item_Virtual_Dedicated_Rack_Extended struct {
	SoftLayer_Billing_Item_Virtual_Dedicated_Rack

	// AssociatedParentCount - A count of a billing item's associated parent billing item. This object will
	// be the same as the parent billing item if parentId is set.
	AssociatedParentCount uint64 `json:"associatedParentCount,omitempty"`

	// ActiveAgreement - <nil>
	ActiveAgreement *SoftLayer_Account_Agreement `json:"activeAgreement,omitempty"`

	// BillableChildren - A billing item's recurring child items that have once been billed and are
	// scheduled to be billed in the future.
	BillableChildren []*SoftLayer_Billing_Item `json:"billableChildren,omitempty"`

	// ActiveBundledItemCount - A count of a Billing Item's active bundled billing items.
	ActiveBundledItemCount uint64 `json:"activeBundledItemCount,omitempty"`

	// BillingCyclePublicUsageIn - The total public inbound bandwidth for this virtual rack for the current
	// billing cycle.
	BillingCyclePublicUsageIn slapi.Float64 `json:"billingCyclePublicUsageIn,omitempty"`

	// NextInvoiceTotalOneTimeAmount - A Billing Item's total, including any child billing items if they
	// exist.'
	NextInvoiceTotalOneTimeAmount slapi.Float64 `json:"nextInvoiceTotalOneTimeAmount,omitempty"`

	// ActiveSparePoolBundledItems - no documentation
	ActiveSparePoolBundledItems []*SoftLayer_Billing_Item `json:"activeSparePoolBundledItems,omitempty"`

	// FilteredNextInvoiceChildrenCount - A count of a Billing Item's associated child billing items,
	// excluding some items with a $0.00 recurring fee.
	FilteredNextInvoiceChildrenCount uint64 `json:"filteredNextInvoiceChildrenCount,omitempty"`

	// BillingCyclePublicBandwidthUsage - The raw public bandwidth usage data for the current billing
	// cycle.
	BillingCyclePublicBandwidthUsage []*SoftLayer_Network_Bandwidth_Usage `json:"billingCyclePublicBandwidthUsage,omitempty"`

	// PendingOrderItem - The new order item that will replace this billing item.
	PendingOrderItem *SoftLayer_Billing_Order_Item `json:"pendingOrderItem,omitempty"`

	// DowngradeItems - For product items which have a downgrade path defined, this will return those
	// product items.
	DowngradeItems []*SoftLayer_Product_Item `json:"downgradeItems,omitempty"`

	// AvailableMatchingVlanCount - no documentation
	AvailableMatchingVlanCount uint64 `json:"availableMatchingVlanCount,omitempty"`

	// ChildrenCount - no documentation
	ChildrenCount uint64 `json:"childrenCount,omitempty"`

	// InvoiceItemCount - A count of all invoice items associated with the billing item
	InvoiceItemCount uint64 `json:"invoiceItemCount,omitempty"`

	// PendingCancellationFlag - This flag indicates whether a billing item is scheduled to be canceled or
	// not.
	PendingCancellationFlag bool `json:"pendingCancellationFlag,omitempty"`

	// ActiveAssociatedGuestDiskBillingItems - <nil>
	ActiveAssociatedGuestDiskBillingItems []*SoftLayer_Billing_Item `json:"activeAssociatedGuestDiskBillingItems,omitempty"`

	// InvoiceItems - no documentation
	InvoiceItems []*SoftLayer_Billing_Invoice_Item `json:"invoiceItems,omitempty"`

	// SoftwareDescription - no documentation
	SoftwareDescription *SoftLayer_Software_Description `json:"softwareDescription,omitempty"`

	// CancellationRequests - This will return any cancellation requests that are associated with this
	// billing item.
	CancellationRequests []*SoftLayer_Billing_Item_Cancellation_Request `json:"cancellationRequests,omitempty"`

	// ActiveSparePoolAssociatedGuestDiskBillingItems - <nil>
	ActiveSparePoolAssociatedGuestDiskBillingItems []*SoftLayer_Billing_Item `json:"activeSparePoolAssociatedGuestDiskBillingItems,omitempty"`

	// BundleItems - no documentation
	BundleItems []*SoftLayer_Product_Item_Bundles `json:"bundleItems,omitempty"`

	// NextInvoiceChildren - A Billing Item's child billing items and associated items'
	NextInvoiceChildren []*SoftLayer_Billing_Item `json:"nextInvoiceChildren,omitempty"`

	// UpgradeItemCount - A count of billing items whose product item has an upgrade path defined in our
	// system will return all the product items in the upgrade path.
	UpgradeItemCount uint64 `json:"upgradeItemCount,omitempty"`

	// BillingCyclePrivateUsageOut - The total private network outbound bandwidth for this virtual rack for
	// the current billing cycle.
	BillingCyclePrivateUsageOut slapi.Float64 `json:"billingCyclePrivateUsageOut,omitempty"`

	// BillingCyclePrivateUsageIn - The total private network inbound bandwidth for this virtual rack for
	// the current billing cycle.
	BillingCyclePrivateUsageIn slapi.Float64 `json:"billingCyclePrivateUsageIn,omitempty"`

	// Category - The item category to which the billing item's item belongs.
	Category *SoftLayer_Product_Item_Category `json:"category,omitempty"`

	// Location - The location of the billing item. Some billing items have physical properties such as the
	// server itself. For items such as these, we provide location information.
	Location *SoftLayer_Location `json:"location,omitempty"`

	// ActiveAssociatedGuestDiskBillingItemCount - no documentation
	ActiveAssociatedGuestDiskBillingItemCount uint64 `json:"activeAssociatedGuestDiskBillingItemCount,omitempty"`

	// UpgradeItem - Billing items whose product item has an upgrade path defined in our system will return
	// the next product item in the upgrade path.
	UpgradeItem *SoftLayer_Product_Item `json:"upgradeItem,omitempty"`

	// AssociatedBillingItem - A billing item's associated parent. This is to be used for billing items
	// that are "floating", and therefore are not child items of any parent billing item. If it is desired
	// to associate an item to another, populate this with the SoftLayer_Billing_Item ID of that associated
	// parent item.
	AssociatedBillingItem *SoftLayer_Billing_Item `json:"associatedBillingItem,omitempty"`

	// ActiveAssociatedChildren - A billing item's active associated child billing items. This includes
	// "floating" items that are not necessarily child items of this billing item.
	ActiveAssociatedChildren []*SoftLayer_Billing_Item `json:"activeAssociatedChildren,omitempty"`

	// BundledItemCount - no documentation
	BundledItemCount uint64 `json:"bundledItemCount,omitempty"`

	// AvailableMatchingVlans - <nil>
	AvailableMatchingVlans []*SoftLayer_Network_Vlan `json:"availableMatchingVlans,omitempty"`

	// BillingCycleBandwidthUsageCount - A count of the raw bandwidth usage data for the current billing
	// cycle. One object is returned for each network a virtual rack is attached to.
	BillingCycleBandwidthUsageCount uint64 `json:"billingCycleBandwidthUsageCount,omitempty"`

	// AssociatedBillingItemHistory - A history of billing items which a billing item has been associated
	// with.
	AssociatedBillingItemHistory []*SoftLayer_Billing_Item_Association_History `json:"associatedBillingItemHistory,omitempty"`

	// ActiveCancellationItem - A service cancellation request item that corresponds to the billing item.
	ActiveCancellationItem *SoftLayer_Billing_Item_Cancellation_Request_Item `json:"activeCancellationItem,omitempty"`

	// BillingCyclePrivateUsageTotal - The total private network bandwidth for this virtual rack for the
	// current billing cycle.
	BillingCyclePrivateUsageTotal uint `json:"billingCyclePrivateUsageTotal,omitempty"`

	// NextInvoiceTotalOneTimeTaxAmount - A Billing Item's total, including any child billing items if they
	// exist.'
	NextInvoiceTotalOneTimeTaxAmount slapi.Float64 `json:"nextInvoiceTotalOneTimeTaxAmount,omitempty"`

	// BillableChildrenCount - A count of a billing item's recurring child items that have once been billed
	// and are scheduled to be billed in the future.
	BillableChildrenCount uint64 `json:"billableChildrenCount,omitempty"`

	// ProvisionTransaction - no documentation
	ProvisionTransaction *SoftLayer_Provisioning_Version1_Transaction `json:"provisionTransaction,omitempty"`

	// ActiveChildrenCount - A count of a Billing Item's active child billing items.
	ActiveChildrenCount uint64 `json:"activeChildrenCount,omitempty"`

	// ActiveSparePoolBundledItemCount - A count of a Billing Item's spare pool bundled billing items.
	ActiveSparePoolBundledItemCount uint64 `json:"activeSparePoolBundledItemCount,omitempty"`

	// AssociatedParent - A billing item's associated parent billing item. This object will be the same as
	// the parent billing item if parentId is set.
	AssociatedParent []*SoftLayer_Billing_Item `json:"associatedParent,omitempty"`

	// BandwidthAllocation - no documentation
	BandwidthAllocation *SoftLayer_Network_Bandwidth_Version1_Allocation `json:"bandwidthAllocation,omitempty"`

	// ActiveBundledItems - no documentation
	ActiveBundledItems []*SoftLayer_Billing_Item `json:"activeBundledItems,omitempty"`

	// CancellationReason - no documentation
	CancellationReason *SoftLayer_Billing_Item_Cancellation_Reason `json:"cancellationReason,omitempty"`

	// BillingCyclePrivateBandwidthUsage - The raw private bandwidth usage data for the current billing
	// cycle.
	BillingCyclePrivateBandwidthUsage []*SoftLayer_Network_Bandwidth_Usage `json:"billingCyclePrivateBandwidthUsage,omitempty"`

	// NextInvoiceTotalRecurringAmount - A Billing Item's total, including any child billing items and
	// associated billing items if they exist.'
	NextInvoiceTotalRecurringAmount slapi.Float64 `json:"nextInvoiceTotalRecurringAmount,omitempty"`

	// OriginalLocation - The original physical location for this billing item--may differ from current.
	OriginalLocation *SoftLayer_Location `json:"originalLocation,omitempty"`

	// Resource - The virtual rack that a virtual rack billing item is associated with.
	Resource *SoftLayer_Network_Bandwidth_Version1_Allotment `json:"resource,omitempty"`

	// AssociatedChildren - A Billing Item's associated child billing items. This includes "floating" items
	// that are not necessarily child billing items of this billing item.
	AssociatedChildren []*SoftLayer_Billing_Item `json:"associatedChildren,omitempty"`

	// ActiveFlag - <nil>
	ActiveFlag bool `json:"activeFlag,omitempty"`

	// NextInvoiceChildrenCount - A count of a Billing Item's child billing items and associated items'
	NextInvoiceChildrenCount uint64 `json:"nextInvoiceChildrenCount,omitempty"`

	// BillingCyclePublicBandwidthUsageCount - A count of the raw public bandwidth usage data for the
	// current billing cycle.
	BillingCyclePublicBandwidthUsageCount uint64 `json:"billingCyclePublicBandwidthUsageCount,omitempty"`

	// ChildrenWithActiveAgreement - no documentation
	ChildrenWithActiveAgreement []*SoftLayer_Billing_Item `json:"childrenWithActiveAgreement,omitempty"`

	// BillingCyclePublicUsageOut - The total public outbound bandwidth for this virtual rack for the
	// current billing cycle.
	BillingCyclePublicUsageOut slapi.Float64 `json:"billingCyclePublicUsageOut,omitempty"`

	// NextInvoiceTotalRecurringTaxAmount - This is deprecated and will always be zero. Because tax is
	// calculated in real-time, previewing the next recurring invoice is pre-tax only.
	NextInvoiceTotalRecurringTaxAmount slapi.Float64 `json:"nextInvoiceTotalRecurringTaxAmount,omitempty"`

	// CanceledChildren - no documentation
	CanceledChildren []*SoftLayer_Billing_Item `json:"canceledChildren,omitempty"`

	// ActiveChildren - no documentation
	ActiveChildren []*SoftLayer_Billing_Item `json:"activeChildren,omitempty"`

	// DowngradeItemCount - A count of for product items which have a downgrade path defined, this will
	// return those product items.
	DowngradeItemCount uint64 `json:"downgradeItemCount,omitempty"`

	// NonZeroNextInvoiceChildren - A Billing Item's associated child billing items, excluding ALL items
	// with a $0.00 recurring fee.
	NonZeroNextInvoiceChildren []*SoftLayer_Billing_Item `json:"nonZeroNextInvoiceChildren,omitempty"`

	// ActiveAssociatedChildrenCount - A count of a billing item's active associated child billing items.
	// This includes "floating" items that are not necessarily child items of this billing item.
	ActiveAssociatedChildrenCount uint64 `json:"activeAssociatedChildrenCount,omitempty"`

	// ChildrenWithActiveAgreementCount - A count of a Billing Item's active child billing items.
	ChildrenWithActiveAgreementCount uint64 `json:"childrenWithActiveAgreementCount,omitempty"`

	// OrderItem - A billing item's original order item. Simply a reference to the original order from
	// which this billing item was created.
	OrderItem *SoftLayer_Billing_Order_Item `json:"orderItem,omitempty"`

	// NonZeroNextInvoiceChildrenCount - A count of a Billing Item's associated child billing items,
	// excluding ALL items with a $0.00 recurring fee.
	NonZeroNextInvoiceChildrenCount uint64 `json:"nonZeroNextInvoiceChildrenCount,omitempty"`

	// Parent - A billing item's parent item. If a billing item has no parent item then this value is null.
	Parent *SoftLayer_Billing_Item `json:"parent,omitempty"`

	// HourlyFlag - A flag that will reflect whether this billing item is billed on an hourly basis or not.
	HourlyFlag bool `json:"hourlyFlag,omitempty"`

	// BundledItems - no documentation
	BundledItems []*SoftLayer_Billing_Item `json:"bundledItems,omitempty"`

	// FilteredNextInvoiceChildren - A Billing Item's associated child billing items, excluding some items
	// with a $0.00 recurring fee.
	FilteredNextInvoiceChildren []*SoftLayer_Billing_Item `json:"filteredNextInvoiceChildren,omitempty"`

	// BillingCyclePublicUsageTotal - The total public bandwidth for this virtual rack for the current
	// billing cycle.
	BillingCyclePublicUsageTotal uint `json:"billingCyclePublicUsageTotal,omitempty"`

	// Item - The entry in the SoftLayer product catalog that a billing item is based upon.
	Item *SoftLayer_Product_Item `json:"item,omitempty"`

	// AssociatedBillingItemHistoryCount - A count of a history of billing items which a billing item has
	// been associated with.
	AssociatedBillingItemHistoryCount uint64 `json:"associatedBillingItemHistoryCount,omitempty"`

	// InvoiceItem - no documentation
	InvoiceItem *SoftLayer_Billing_Invoice_Item `json:"invoiceItem,omitempty"`

	// ParentVirtualGuestBillingItem - A billing item's parent item. If a billing item has no parent item
	// then this value is null.
	ParentVirtualGuestBillingItem *SoftLayer_Billing_Item_Virtual_Guest `json:"parentVirtualGuestBillingItem,omitempty"`

	// BillingCyclePrivateBandwidthUsageCount - A count of the raw private bandwidth usage data for the
	// current billing cycle.
	BillingCyclePrivateBandwidthUsageCount uint64 `json:"billingCyclePrivateBandwidthUsageCount,omitempty"`

	// Children - no documentation
	Children []*SoftLayer_Billing_Item `json:"children,omitempty"`

	// BundleItemCount - no documentation
	BundleItemCount uint64 `json:"bundleItemCount,omitempty"`

	// BillingCycleBandwidthUsage - The raw bandwidth usage data for the current billing cycle. One object
	// is returned for each network a virtual rack is attached to.
	BillingCycleBandwidthUsage []*SoftLayer_Network_Bandwidth_Usage `json:"billingCycleBandwidthUsage,omitempty"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account,omitempty"`

	// CanceledChildrenCount - A count of a Billing Item's active child billing items.
	CanceledChildrenCount uint64 `json:"canceledChildrenCount,omitempty"`

	// CancellationRequestCount - A count of this will return any cancellation requests that are associated
	// with this billing item.
	CancellationRequestCount uint64 `json:"cancellationRequestCount,omitempty"`

	// Package - The package under which this billing item was sold. A Package is the general grouping of
	// products as seen on our order forms.
	Package *SoftLayer_Product_Package `json:"package,omitempty"`

	// ActiveSparePoolAssociatedGuestDiskBillingItemCount - no documentation
	ActiveSparePoolAssociatedGuestDiskBillingItemCount uint64 `json:"activeSparePoolAssociatedGuestDiskBillingItemCount,omitempty"`

	// AssociatedChildrenCount - A count of a Billing Item's associated child billing items. This includes
	// "floating" items that are not necessarily child billing items of this billing item.
	AssociatedChildrenCount uint64 `json:"associatedChildrenCount,omitempty"`

	// ActiveAgreementFlag - A flag indicating that the billing item is under an active agreement.
	ActiveAgreementFlag *SoftLayer_Account_Agreement `json:"activeAgreementFlag,omitempty"`

	// UpgradeItems - Billing items whose product item has an upgrade path defined in our system will
	// return all the product items in the upgrade path.
	UpgradeItems []*SoftLayer_Product_Item `json:"upgradeItems,omitempty"`
}

func (softlayer_billing_item_virtual_dedicated_rack *SoftLayer_Billing_Item_Virtual_Dedicated_Rack_Extended) String() string {
	return "SoftLayer_Billing_Item_Virtual_Dedicated_Rack"
}
