package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"

	time "time"
)

// SoftLayer_Billing_Invoice_Item_Hardware - The SoftLayer_Billing_Invoice_Item_Hardware data type
// contains a "resource". This resource is a link to the hardware tied to a SoftLayer_Billing_item
// whose category code is "server".
type SoftLayer_Billing_Invoice_Item_Hardware struct {

	// OneTimeAfterTaxAmount - An invoice item's one-time fee total after taxes. This does not include any
	// child invoice items.
	OneTimeAfterTaxAmount slapi.Float64 `json:"oneTimeAfterTaxAmount,omitempty"`

	// LaborAfterTaxAmount - An invoice item's labor fee total after taxes. This does not include any child
	// invoice items.
	LaborAfterTaxAmount slapi.Float64 `json:"laborAfterTaxAmount,omitempty"`

	// HostName - The Host name of the invoiced item. This is only used on invoice items whose category is
	// "server".
	HostName string `json:"hostName,omitempty"`

	// ResourceTableId - A unique identifier for a SoftLayer Service that is associated to an invoice item.
	ResourceTableId int `json:"resourceTableId,omitempty"`

	// CategoryCode - The item category of the invoice item being invoiced.
	CategoryCode string `json:"categoryCode,omitempty"`

	// OneTimeFeeTaxRate - no documentation
	OneTimeFeeTaxRate slapi.Float64 `json:"oneTimeFeeTaxRate,omitempty"`

	// SetupAfterTaxAmount - An invoice item's setup fee total after taxes. This does not include any child
	// invoice items.
	SetupAfterTaxAmount slapi.Float64 `json:"setupAfterTaxAmount,omitempty"`

	// InvoiceId - no documentation
	InvoiceId int `json:"invoiceId,omitempty"`

	// ParentId - The parent invoice item, usually the server invoice item.
	ParentId int `json:"parentId,omitempty"`

	// SetupFee - If there were any setup fees they will show up here. These are normally a one-time fee.
	SetupFee slapi.Float64 `json:"setupFee,omitempty"`

	// SetupFeeTaxRate - no documentation
	SetupFeeTaxRate slapi.Float64 `json:"setupFeeTaxRate,omitempty"`

	// LaborTaxAmount - An invoice item's labor tax amount. This does not include any child invoice items.
	LaborTaxAmount slapi.Float64 `json:"laborTaxAmount,omitempty"`

	// OneTimeTaxAmount - An invoice item's one-time tax amount. This does not include any child invoice
	// items.
	OneTimeTaxAmount slapi.Float64 `json:"oneTimeTaxAmount,omitempty"`

	// HourlyRecurringFee - The hourly recurring fee of the invoice item represented by a floating point
	// decimal in US Dollars
	HourlyRecurringFee slapi.Float64 `json:"hourlyRecurringFee,omitempty"`

	// Description - no documentation
	Description string `json:"description,omitempty"`

	// SetupTaxAmount - An invoice item's setup tax amount. This does not include any child invoice items.
	SetupTaxAmount slapi.Float64 `json:"setupTaxAmount,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// DomainName - The domain name of the invoiced item. This is only used on invoice items whose category
	// is "server".
	DomainName string `json:"domainName,omitempty"`

	// LaborFee - no documentation
	LaborFee slapi.Float64 `json:"laborFee,omitempty"`

	// RecurringFee - The recurring fee of the invoice item represented by a floating point decimal in US
	// Dollars
	RecurringFee slapi.Float64 `json:"recurringFee,omitempty"`

	// AssociatedInvoiceItemId - no documentation
	AssociatedInvoiceItemId int `json:"associatedInvoiceItemId,omitempty"`

	// RecurringAfterTaxAmount - An invoice item's recurring fee total after taxes. This does not include
	// any child invoice items.
	RecurringAfterTaxAmount slapi.Float64 `json:"recurringAfterTaxAmount,omitempty"`

	// RecurringFeeTaxRate - no documentation
	RecurringFeeTaxRate slapi.Float64 `json:"recurringFeeTaxRate,omitempty"`

	// LaborFeeTaxRate - no documentation
	LaborFeeTaxRate slapi.Float64 `json:"laborFeeTaxRate,omitempty"`

	// Notes - A note to help describe more about the item. This normally holds usernames, or some other
	// bit of extra information.
	Notes string `json:"notes,omitempty"`

	// BillingItemId - The billing item from which this invoice item was generated.
	BillingItemId int `json:"billingItemId,omitempty"`

	// OneTimeFee - If there are any one-time charges assessed, it will show up here represented by a
	// floating point decimal in US Dollars
	OneTimeFee slapi.Float64 `json:"oneTimeFee,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// RecurringTaxAmount - An invoice item's recurring tax amount. This does not include any child invoice
	// items.
	RecurringTaxAmount slapi.Float64 `json:"recurringTaxAmount,omitempty"`

	// ChildrenCount - A count of an Invoice Item's child invoice items. Only parent invoice items have
	// children. For instance, a server invoice item will have children.
	ChildrenCount uint64 `json:"childrenCount,omitempty"`

	// AssociatedInvoiceItem - An Invoice Item's associated invoice item. If this is populated, it means
	// this is an orphaned invoice item, but logically belongs to the associated invoice item.
	AssociatedInvoiceItem *SoftLayer_Billing_Invoice_Item `json:"associatedInvoiceItem,omitempty"`

	// FilteredAssociatedChildren - An Invoice Item's associated child invoice items, excluding some items
	// with a $0.00 recurring fee. Only parent invoice items have associated children. For instance, a
	// server invoice item may have associated children.
	FilteredAssociatedChildren []*SoftLayer_Billing_Invoice_Item `json:"filteredAssociatedChildren,omitempty"`

	// Location - no documentation
	Location *SoftLayer_Location `json:"location,omitempty"`

	// NonZeroAssociatedChildrenCount - A count of an Invoice Item's associated child invoice items,
	// excluding ALL items with a $0.00 recurring fee. Only parent invoice items have associated children.
	// For instance, a server invoice item may have associated children.
	NonZeroAssociatedChildrenCount uint64 `json:"nonZeroAssociatedChildrenCount,omitempty"`

	// TotalRecurringTaxAmount - A Billing Item's total, including any child billing items if they exist.'
	TotalRecurringTaxAmount slapi.Float64 `json:"totalRecurringTaxAmount,omitempty"`

	// TotalOneTimeAmount - An invoice Item's total, including any child invoice items if they exist.
	TotalOneTimeAmount slapi.Float64 `json:"totalOneTimeAmount,omitempty"`

	// TotalRecurringAmount - An invoice Item's total, including any child invoice items if they exist.
	TotalRecurringAmount slapi.Float64 `json:"totalRecurringAmount,omitempty"`

	// BillingItem - An Invoice Item's billing item, from which this item was generated.
	BillingItem *SoftLayer_Billing_Item `json:"billingItem,omitempty"`

	// Parent - Every item tied to a server should have a parent invoice item which is the server line
	// item. This is how we associate items to a server.
	Parent *SoftLayer_Billing_Invoice_Item `json:"parent,omitempty"`

	// AssociatedChildrenCount - A count of an Invoice Item's associated child invoice items. Only parent
	// invoice items have associated children. For instance, a server invoice item may have associated
	// children.
	AssociatedChildrenCount uint64 `json:"associatedChildrenCount,omitempty"`

	// Category - no documentation
	Category *SoftLayer_Product_Item_Category `json:"category,omitempty"`

	// Children - An Invoice Item's child invoice items. Only parent invoice items have children. For
	// instance, a server invoice item will have children.
	Children []*SoftLayer_Billing_Invoice_Item `json:"children,omitempty"`

	// AssociatedChildren - An Invoice Item's associated child invoice items. Only parent invoice items
	// have associated children. For instance, a server invoice item may have associated children.
	AssociatedChildren []*SoftLayer_Billing_Invoice_Item `json:"associatedChildren,omitempty"`

	// NonZeroAssociatedChildren - An Invoice Item's associated child invoice items, excluding ALL items
	// with a $0.00 recurring fee. Only parent invoice items have associated children. For instance, a
	// server invoice item may have associated children.
	NonZeroAssociatedChildren []*SoftLayer_Billing_Invoice_Item `json:"nonZeroAssociatedChildren,omitempty"`

	// Resource - no documentation
	Resource *SoftLayer_Hardware `json:"resource,omitempty"`

	// Invoice - no documentation
	Invoice *SoftLayer_Billing_Invoice `json:"invoice,omitempty"`

	// TotalOneTimeTaxAmount - An invoice Item's total, including any child invoice items if they exist.
	TotalOneTimeTaxAmount slapi.Float64 `json:"totalOneTimeTaxAmount,omitempty"`

	// FilteredAssociatedChildrenCount - A count of an Invoice Item's associated child invoice items,
	// excluding some items with a $0.00 recurring fee. Only parent invoice items have associated children.
	// For instance, a server invoice item may have associated children.
	FilteredAssociatedChildrenCount uint64 `json:"filteredAssociatedChildrenCount,omitempty"`
}

func (softlayer_billing_invoice_item_hardware *SoftLayer_Billing_Invoice_Item_Hardware) String() string {
	return "SoftLayer_Billing_Invoice_Item_Hardware"
}
