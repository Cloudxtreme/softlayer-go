package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Billing_Invoice_Item - Each billing invoice item makes up a record within an invoice. This
// provides you with a detailed record of everything related to an invoice item. When you are billed,
// our system takes active billing items and creates an invoice. These invoice items are a copy of your
// active billing items, and make up the contents of your invoice.
type SoftLayer_Billing_Invoice_Item struct {

	// CategoryCode - The item category of the invoice item being invoiced.
	CategoryCode string `json:"categoryCode"`

	// RecurringFee - The recurring fee of the invoice item represented by a floating point decimal in US
	// Dollars
	RecurringFee float64 `json:"recurringFee"`

	// SetupTaxAmount - An invoice item's setup tax amount. This does not include any child invoice items.
	SetupTaxAmount float64 `json:"setupTaxAmount"`

	// AssociatedInvoiceItemId - no documentation
	AssociatedInvoiceItemId int `json:"associatedInvoiceItemId"`

	// LaborTaxAmount - An invoice item's labor tax amount. This does not include any child invoice items.
	LaborTaxAmount float64 `json:"laborTaxAmount"`

	// Notes - A note to help describe more about the item. This normally holds usernames, or some other
	// bit of extra information.
	Notes string `json:"notes"`

	// ResourceTableId - A unique identifier for a SoftLayer Service that is associated to an invoice item.
	ResourceTableId int `json:"resourceTableId"`

	// InvoiceId - no documentation
	InvoiceId int `json:"invoiceId"`

	// RecurringTaxAmount - An invoice item's recurring tax amount. This does not include any child invoice
	// items.
	RecurringTaxAmount float64 `json:"recurringTaxAmount"`

	// SetupFeeTaxRate - no documentation
	SetupFeeTaxRate float64 `json:"setupFeeTaxRate"`

	// Description - no documentation
	Description string `json:"description"`

	// HostName - The Host name of the invoiced item. This is only used on invoice items whose category is
	// "server".
	HostName string `json:"hostName"`

	// OneTimeAfterTaxAmount - An invoice item's one-time fee total after taxes. This does not include any
	// child invoice items.
	OneTimeAfterTaxAmount float64 `json:"oneTimeAfterTaxAmount"`

	// OneTimeTaxAmount - An invoice item's one-time tax amount. This does not include any child invoice
	// items.
	OneTimeTaxAmount float64 `json:"oneTimeTaxAmount"`

	// ParentId - The parent invoice item, usually the server invoice item.
	ParentId int `json:"parentId"`

	// SetupFee - If there were any setup fees they will show up here. These are normally a one-time fee.
	SetupFee float64 `json:"setupFee"`

	// BillingItemId - The billing item from which this invoice item was generated.
	BillingItemId int `json:"billingItemId"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// DomainName - The domain name of the invoiced item. This is only used on invoice items whose category
	// is "server".
	DomainName string `json:"domainName"`

	// Id - no documentation
	Id int `json:"id"`

	// LaborFee - no documentation
	LaborFee float64 `json:"laborFee"`

	// LaborFeeTaxRate - no documentation
	LaborFeeTaxRate float64 `json:"laborFeeTaxRate"`

	// OneTimeFeeTaxRate - no documentation
	OneTimeFeeTaxRate float64 `json:"oneTimeFeeTaxRate"`

	// RecurringAfterTaxAmount - An invoice item's recurring fee total after taxes. This does not include
	// any child invoice items.
	RecurringAfterTaxAmount float64 `json:"recurringAfterTaxAmount"`

	// HourlyRecurringFee - The hourly recurring fee of the invoice item represented by a floating point
	// decimal in US Dollars
	HourlyRecurringFee float64 `json:"hourlyRecurringFee"`

	// RecurringFeeTaxRate - no documentation
	RecurringFeeTaxRate float64 `json:"recurringFeeTaxRate"`

	// LaborAfterTaxAmount - An invoice item's labor fee total after taxes. This does not include any child
	// invoice items.
	LaborAfterTaxAmount float64 `json:"laborAfterTaxAmount"`

	// OneTimeFee - If there are any one-time charges assessed, it will show up here represented by a
	// floating point decimal in US Dollars
	OneTimeFee float64 `json:"oneTimeFee"`

	// SetupAfterTaxAmount - An invoice item's setup fee total after taxes. This does not include any child
	// invoice items.
	SetupAfterTaxAmount float64 `json:"setupAfterTaxAmount"`
}

func (softlayer_billing_invoice_item *SoftLayer_Billing_Invoice_Item) String() string {
	return "SoftLayer_Billing_Invoice_Item"
}

// SoftLayer_Billing_Invoice_Item_Extended is SoftLayer_Billing_Invoice_Item with all maskable types.
type SoftLayer_Billing_Invoice_Item_Extended struct {
	SoftLayer_Billing_Invoice_Item

	// Invoice - no documentation
	Invoice *SoftLayer_Billing_Invoice `json:"invoice"`

	// NonZeroAssociatedChildren - An Invoice Item's associated child invoice items, excluding ALL items
	// with a $0.00 recurring fee. Only parent invoice items have associated children. For instance, a
	// server invoice item may have associated children.
	NonZeroAssociatedChildren []*SoftLayer_Billing_Invoice_Item `json:"nonZeroAssociatedChildren"`

	// TotalRecurringAmount - An invoice Item's total, including any child invoice items if they exist.
	TotalRecurringAmount float32 `json:"totalRecurringAmount"`

	// TotalOneTimeTaxAmount - An invoice Item's total, including any child invoice items if they exist.
	TotalOneTimeTaxAmount float32 `json:"totalOneTimeTaxAmount"`

	// AssociatedInvoiceItem - An Invoice Item's associated invoice item. If this is populated, it means
	// this is an orphaned invoice item, but logically belongs to the associated invoice item.
	AssociatedInvoiceItem *SoftLayer_Billing_Invoice_Item `json:"associatedInvoiceItem"`

	// Parent - Every item tied to a server should have a parent invoice item which is the server line
	// item. This is how we associate items to a server.
	Parent *SoftLayer_Billing_Invoice_Item `json:"parent"`

	// FilteredAssociatedChildren - An Invoice Item's associated child invoice items, excluding some items
	// with a $0.00 recurring fee. Only parent invoice items have associated children. For instance, a
	// server invoice item may have associated children.
	FilteredAssociatedChildren []*SoftLayer_Billing_Invoice_Item `json:"filteredAssociatedChildren"`

	// Location - no documentation
	Location *SoftLayer_Location `json:"location"`

	// ChildrenCount - A count of an Invoice Item's child invoice items. Only parent invoice items have
	// children. For instance, a server invoice item will have children.
	ChildrenCount uint64 `json:"childrenCount"`

	// Children - An Invoice Item's child invoice items. Only parent invoice items have children. For
	// instance, a server invoice item will have children.
	Children []*SoftLayer_Billing_Invoice_Item `json:"children"`

	// TotalOneTimeAmount - An invoice Item's total, including any child invoice items if they exist.
	TotalOneTimeAmount float32 `json:"totalOneTimeAmount"`

	// AssociatedChildrenCount - A count of an Invoice Item's associated child invoice items. Only parent
	// invoice items have associated children. For instance, a server invoice item may have associated
	// children.
	AssociatedChildrenCount uint64 `json:"associatedChildrenCount"`

	// AssociatedChildren - An Invoice Item's associated child invoice items. Only parent invoice items
	// have associated children. For instance, a server invoice item may have associated children.
	AssociatedChildren []*SoftLayer_Billing_Invoice_Item `json:"associatedChildren"`

	// NonZeroAssociatedChildrenCount - A count of an Invoice Item's associated child invoice items,
	// excluding ALL items with a $0.00 recurring fee. Only parent invoice items have associated children.
	// For instance, a server invoice item may have associated children.
	NonZeroAssociatedChildrenCount uint64 `json:"nonZeroAssociatedChildrenCount"`

	// BillingItem - An Invoice Item's billing item, from which this item was generated.
	BillingItem *SoftLayer_Billing_Item `json:"billingItem"`

	// Category - no documentation
	Category *SoftLayer_Product_Item_Category `json:"category"`

	// TotalRecurringTaxAmount - A Billing Item's total, including any child billing items if they exist.'
	TotalRecurringTaxAmount float32 `json:"totalRecurringTaxAmount"`

	// FilteredAssociatedChildrenCount - A count of an Invoice Item's associated child invoice items,
	// excluding some items with a $0.00 recurring fee. Only parent invoice items have associated children.
	// For instance, a server invoice item may have associated children.
	FilteredAssociatedChildrenCount uint64 `json:"filteredAssociatedChildrenCount"`
}

func (softlayer_billing_invoice_item *SoftLayer_Billing_Invoice_Item_Extended) String() string {
	return "SoftLayer_Billing_Invoice_Item"
}
