package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Billing_Order_Item - Every individual item that a SoftLayer customer is billed for is
// recorded in the SoftLayer_Billing_Item data type. Billing items range from server chassis to hard
// drives to control panels, bandwidth quota upgrades and port upgrade charges. Softlayer
// [[SoftLayer_Billing_Invoice|invoices]] are generated from the cost of a customer's billing items.
// Billing items are copied from the product catalog as they're ordered by customers to create a
// reference between an account and the billable items they own. Billing items exist in a tree
// relationship. Items are associated with each other by parent/child relationships. Component items
// such as CPU's, and software each have a parent billing item for the server chassis they're
// associated with. Billing Items with a null parent item do not have an associated parent item.
type SoftLayer_Billing_Order_Item struct {

	// LaborFee - no documentation
	LaborFee float64 `json:"laborFee"`

	// RecurringTaxAmount - An order item's recurring tax amount. This does not include any child invoice
	// items.
	RecurringTaxAmount float64 `json:"recurringTaxAmount"`

	// DomainName - The domain name of the server as designated by the purchaser at the time of order
	// placement.
	DomainName string `json:"domainName"`

	// OneTimeAfterTaxAmount - An order item's one-time fee total after taxes. This does not include any
	// child invoice items.
	OneTimeAfterTaxAmount float64 `json:"oneTimeAfterTaxAmount"`

	// Description - no documentation
	Description string `json:"description"`

	// Id - <nil>
	Id int `json:"id"`

	// LaborFeeTaxRate - The rate at which labor fees are taxed if you are a taxable customer.
	LaborFeeTaxRate float64 `json:"laborFeeTaxRate"`

	// LaborTaxAmount - An order item's labor tax amount. This does not include any child invoice items.
	LaborTaxAmount float64 `json:"laborTaxAmount"`

	// PromoCodeId - <nil>
	PromoCodeId int `json:"promoCodeId"`

	// Quantity - no documentation
	Quantity int `json:"quantity"`

	// RecurringAfterTaxAmount - An order item's recurring fee total after taxes. This does not include any
	// child invoice items.
	RecurringAfterTaxAmount float64 `json:"recurringAfterTaxAmount"`

	// SetupFee - no documentation
	SetupFee float64 `json:"setupFee"`

	// OneTimeTaxAmount - An order item's one-time tax amount. This does not include any child invoice
	// items.
	OneTimeTaxAmount float64 `json:"oneTimeTaxAmount"`

	// SetupAfterTaxAmount - An order item's setup fee total after taxes. This does not include any child
	// invoice items.
	SetupAfterTaxAmount float64 `json:"setupAfterTaxAmount"`

	// ItemPriceId - the item price id (SoftLayer_Product_Item_Price->id) of the ordered item.
	ItemPriceId float64 `json:"itemPriceId"`

	// OneTimeFee - The amount of money charged as a one-time charge for an order item, if applicable.
	// oneTimeFee is measured in US Dollars
	OneTimeFee float64 `json:"oneTimeFee"`

	// SetupFeeTaxRate - The rate at which setup fees are taxed if you are a taxable customer.
	SetupFeeTaxRate float64 `json:"setupFeeTaxRate"`

	// LaborAfterTaxAmount - An order item's labor fee total after taxes. This does not include any child
	// invoice items.
	LaborAfterTaxAmount float64 `json:"laborAfterTaxAmount"`

	// ParentId - <nil>
	ParentId int `json:"parentId"`

	// CategoryCode - no documentation
	CategoryCode string `json:"categoryCode"`

	// HostName - The hostname of the server as designated by the purchaser at the time of order placement.
	HostName string `json:"hostName"`

	// HourlyRecurringFee - The amount of money charged per hourly for an order item, if applicable, and
	// only if it was ordered this day. hourlyRecurringFee is measured in US Dollars
	HourlyRecurringFee float64 `json:"hourlyRecurringFee"`

	// ItemId - no documentation
	ItemId int `json:"itemId"`

	// OneTimeFeeTaxRate - The rate at which one time fees are taxed if you are a taxable customer.
	OneTimeFeeTaxRate float64 `json:"oneTimeFeeTaxRate"`

	// SetupFeeDeferralMonths - no documentation
	SetupFeeDeferralMonths int `json:"setupFeeDeferralMonths"`

	// SetupTaxAmount - An order item's setup tax amount. This does not include any child invoice items.
	SetupTaxAmount float64 `json:"setupTaxAmount"`

	// RecurringFee - The amount of money charged per month for an order item, if applicable. recurringFee
	// is measured in US Dollars
	RecurringFee float64 `json:"recurringFee"`
}

// SoftLayer_Billing_Order_Item_Extended is SoftLayer_Billing_Order_Item with all maskable types.
type SoftLayer_Billing_Order_Item_Extended struct {
	SoftLayer_Billing_Order_Item

	// GlobalIdentifier - no documentation
	GlobalIdentifier string `json:"globalIdentifier"`

	// Parent - The parent order item ID for an item. Items that are associated with a server will have a
	// parent. The parent will be the server item itself.
	Parent *SoftLayer_Billing_Order_Item `json:"parent"`

	// UpgradeItem - The next SoftLayer_Product_Item in the upgrade path for this order item.
	UpgradeItem *SoftLayer_Product_Item `json:"upgradeItem"`

	// Children - The child order items for an order item. All server order items should have children.
	// These children are considered a part of the server.
	Children []*SoftLayer_Billing_Order_Item `json:"children"`

	// SoftwareDescription - For ordered items that are software items, a full description of that software
	// can be found with this property.
	SoftwareDescription *SoftLayer_Software_Description `json:"softwareDescription"`

	// ChildrenCount - A count of the child order items for an order item. All server order items should
	// have children. These children are considered a part of the server.
	ChildrenCount uint64 `json:"childrenCount"`

	// OldBillingItem - This is only populated when an upgrade order is placed. The old billing item
	// represents what the billing was before the upgrade happened.
	OldBillingItem *SoftLayer_Billing_Item `json:"oldBillingItem"`

	// Package - The SoftLayer_Product_Package an order item is a part of.
	Package *SoftLayer_Product_Package `json:"package"`

	// BundledItemCount - A count of the other items included with an ordered item.
	BundledItemCount uint64 `json:"bundledItemCount"`

	// ItemCategoryAnswerCount - no documentation
	ItemCategoryAnswerCount uint64 `json:"itemCategoryAnswerCount"`

	// StorageGroups - The drive storage groups that are attached to this billing order item.
	StorageGroups []*SoftLayer_Configuration_Storage_Group_Order `json:"storageGroups"`

	// StorageGroupCount - A count of the drive storage groups that are attached to this billing order
	// item.
	StorageGroupCount uint64 `json:"storageGroupCount"`

	// NextOrderChildren - <nil>
	NextOrderChildren []*SoftLayer_Billing_Order_Item `json:"nextOrderChildren"`

	// RedundantPowerSupplyCount - A count of power supplies contained within this SoftLayer_Billing_Order
	RedundantPowerSupplyCount uint `json:"redundantPowerSupplyCount"`

	// Location - The location of an ordered item. This is usually the same as the server it is being
	// ordered with. Otherwise it describes the location of the additional service being ordered.
	Location *SoftLayer_Location `json:"location"`

	// HardwareGenericComponent - The component type tied to an order item. All hardware-specific items
	// should have a generic hardware component.
	HardwareGenericComponent *SoftLayer_Hardware_Component_Model_Generic `json:"hardwareGenericComponent"`

	// Item - The SoftLayer_Product_Item tied to an order item. The item is the actual definition of the
	// product being sold.
	Item *SoftLayer_Product_Item `json:"item"`

	// NextOrderChildrenCount - no documentation
	NextOrderChildrenCount uint64 `json:"nextOrderChildrenCount"`

	// Category - no documentation
	Category *SoftLayer_Product_Item_Category `json:"category"`

	// ItemPrice - The SoftLayer_Product_Item_Price tied to an order item. The item price object describes
	// the cost of an item.
	ItemPrice *SoftLayer_Product_Item_Price `json:"itemPrice"`

	// Order - The order to which this item belongs. The order contains all the information related to the
	// items included in an order
	Order *SoftLayer_Billing_Order `json:"order"`

	// BillingItem - no documentation
	BillingItem *SoftLayer_Billing_Item `json:"billingItem"`

	// OrderApprovalDate - <nil>
	OrderApprovalDate *time.Time `json:"orderApprovalDate"`

	// TotalRecurringAmount - The recurring fee of an ordered item. This amount represents the fees that
	// will be charged on a recurring (usually monthly) basis.
	TotalRecurringAmount float32 `json:"totalRecurringAmount"`

	// BundledItems - no documentation
	BundledItems []*SoftLayer_Billing_Order_Item `json:"bundledItems"`

	// ItemCategoryAnswers - no documentation
	ItemCategoryAnswers []*SoftLayer_Billing_Order_Item_Category_Answer `json:"itemCategoryAnswers"`
}

func (softlayer_billing_order_item *SoftLayer_Billing_Order_Item) String() string {
	return "SoftLayer_Billing_Order_Item"
}
