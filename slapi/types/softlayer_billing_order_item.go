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

	// HourlyRecurringFee - The amount of money charged per hourly for an order item, if applicable, and
	// only if it was ordered this day. hourlyRecurringFee is measured in US Dollars
	HourlyRecurringFee float64 `json:"hourlyRecurringFee"`

	// OneTimeFeeTaxRate - The rate at which one time fees are taxed if you are a taxable customer.
	OneTimeFeeTaxRate float64 `json:"oneTimeFeeTaxRate"`

	// SetupAfterTaxAmount - An order item's setup fee total after taxes. This does not include any child
	// invoice items.
	SetupAfterTaxAmount float64 `json:"setupAfterTaxAmount"`

	// Id - <nil>
	Id int `json:"id"`

	// LaborAfterTaxAmount - An order item's labor fee total after taxes. This does not include any child
	// invoice items.
	LaborAfterTaxAmount float64 `json:"laborAfterTaxAmount"`

	// SetupFeeTaxRate - The rate at which setup fees are taxed if you are a taxable customer.
	SetupFeeTaxRate float64 `json:"setupFeeTaxRate"`

	// Quantity - no documentation
	Quantity int `json:"quantity"`

	// RecurringTaxAmount - An order item's recurring tax amount. This does not include any child invoice
	// items.
	RecurringTaxAmount float64 `json:"recurringTaxAmount"`

	// OneTimeAfterTaxAmount - An order item's one-time fee total after taxes. This does not include any
	// child invoice items.
	OneTimeAfterTaxAmount float64 `json:"oneTimeAfterTaxAmount"`

	// ParentId - <nil>
	ParentId int `json:"parentId"`

	// RecurringAfterTaxAmount - An order item's recurring fee total after taxes. This does not include any
	// child invoice items.
	RecurringAfterTaxAmount float64 `json:"recurringAfterTaxAmount"`

	// CategoryCode - no documentation
	CategoryCode string `json:"categoryCode"`

	// LaborFeeTaxRate - The rate at which labor fees are taxed if you are a taxable customer.
	LaborFeeTaxRate float64 `json:"laborFeeTaxRate"`

	// PromoCodeId - <nil>
	PromoCodeId int `json:"promoCodeId"`

	// DomainName - The domain name of the server as designated by the purchaser at the time of order
	// placement.
	DomainName string `json:"domainName"`

	// ItemId - no documentation
	ItemId int `json:"itemId"`

	// LaborTaxAmount - An order item's labor tax amount. This does not include any child invoice items.
	LaborTaxAmount float64 `json:"laborTaxAmount"`

	// OneTimeTaxAmount - An order item's one-time tax amount. This does not include any child invoice
	// items.
	OneTimeTaxAmount float64 `json:"oneTimeTaxAmount"`

	// SetupTaxAmount - An order item's setup tax amount. This does not include any child invoice items.
	SetupTaxAmount float64 `json:"setupTaxAmount"`

	// OneTimeFee - The amount of money charged as a one-time charge for an order item, if applicable.
	// oneTimeFee is measured in US Dollars
	OneTimeFee float64 `json:"oneTimeFee"`

	// ItemPriceId - the item price id (SoftLayer_Product_Item_Price->id) of the ordered item.
	ItemPriceId float64 `json:"itemPriceId"`

	// SetupFee - no documentation
	SetupFee float64 `json:"setupFee"`

	// SetupFeeDeferralMonths - no documentation
	SetupFeeDeferralMonths int `json:"setupFeeDeferralMonths"`

	// HostName - The hostname of the server as designated by the purchaser at the time of order placement.
	HostName string `json:"hostName"`

	// LaborFee - no documentation
	LaborFee float64 `json:"laborFee"`

	// Description - no documentation
	Description string `json:"description"`

	// RecurringFee - The amount of money charged per month for an order item, if applicable. recurringFee
	// is measured in US Dollars
	RecurringFee float64 `json:"recurringFee"`
}

func (softlayer_billing_order_item *SoftLayer_Billing_Order_Item) String() string {
	return "SoftLayer_Billing_Order_Item"
}

// SoftLayer_Billing_Order_Item_Extended is SoftLayer_Billing_Order_Item with all maskable types.
type SoftLayer_Billing_Order_Item_Extended struct {
	SoftLayer_Billing_Order_Item

	// HardwareGenericComponent - The component type tied to an order item. All hardware-specific items
	// should have a generic hardware component.
	HardwareGenericComponent *SoftLayer_Hardware_Component_Model_Generic `json:"hardwareGenericComponent"`

	// RedundantPowerSupplyCount - A count of power supplies contained within this SoftLayer_Billing_Order
	RedundantPowerSupplyCount uint `json:"redundantPowerSupplyCount"`

	// Parent - The parent order item ID for an item. Items that are associated with a server will have a
	// parent. The parent will be the server item itself.
	Parent *SoftLayer_Billing_Order_Item `json:"parent"`

	// ItemCategoryAnswerCount - no documentation
	ItemCategoryAnswerCount uint64 `json:"itemCategoryAnswerCount"`

	// Location - The location of an ordered item. This is usually the same as the server it is being
	// ordered with. Otherwise it describes the location of the additional service being ordered.
	Location *SoftLayer_Location `json:"location"`

	// Category - no documentation
	Category *SoftLayer_Product_Item_Category `json:"category"`

	// StorageGroups - The drive storage groups that are attached to this billing order item.
	StorageGroups []*SoftLayer_Configuration_Storage_Group_Order `json:"storageGroups"`

	// Order - The order to which this item belongs. The order contains all the information related to the
	// items included in an order
	Order *SoftLayer_Billing_Order `json:"order"`

	// BundledItemCount - A count of the other items included with an ordered item.
	BundledItemCount uint64 `json:"bundledItemCount"`

	// OldBillingItem - This is only populated when an upgrade order is placed. The old billing item
	// represents what the billing was before the upgrade happened.
	OldBillingItem *SoftLayer_Billing_Item `json:"oldBillingItem"`

	// BundledItems - no documentation
	BundledItems []*SoftLayer_Billing_Order_Item `json:"bundledItems"`

	// OrderApprovalDate - <nil>
	OrderApprovalDate *time.Time `json:"orderApprovalDate"`

	// TotalRecurringAmount - The recurring fee of an ordered item. This amount represents the fees that
	// will be charged on a recurring (usually monthly) basis.
	TotalRecurringAmount float32 `json:"totalRecurringAmount"`

	// GlobalIdentifier - no documentation
	GlobalIdentifier string `json:"globalIdentifier"`

	// ItemPrice - The SoftLayer_Product_Item_Price tied to an order item. The item price object describes
	// the cost of an item.
	ItemPrice *SoftLayer_Product_Item_Price `json:"itemPrice"`

	// ChildrenCount - A count of the child order items for an order item. All server order items should
	// have children. These children are considered a part of the server.
	ChildrenCount uint64 `json:"childrenCount"`

	// SoftwareDescription - For ordered items that are software items, a full description of that software
	// can be found with this property.
	SoftwareDescription *SoftLayer_Software_Description `json:"softwareDescription"`

	// NextOrderChildrenCount - no documentation
	NextOrderChildrenCount uint64 `json:"nextOrderChildrenCount"`

	// ItemCategoryAnswers - no documentation
	ItemCategoryAnswers []*SoftLayer_Billing_Order_Item_Category_Answer `json:"itemCategoryAnswers"`

	// Item - The SoftLayer_Product_Item tied to an order item. The item is the actual definition of the
	// product being sold.
	Item *SoftLayer_Product_Item `json:"item"`

	// BillingItem - no documentation
	BillingItem *SoftLayer_Billing_Item `json:"billingItem"`

	// Children - The child order items for an order item. All server order items should have children.
	// These children are considered a part of the server.
	Children []*SoftLayer_Billing_Order_Item `json:"children"`

	// NextOrderChildren - <nil>
	NextOrderChildren []*SoftLayer_Billing_Order_Item `json:"nextOrderChildren"`

	// StorageGroupCount - A count of the drive storage groups that are attached to this billing order
	// item.
	StorageGroupCount uint64 `json:"storageGroupCount"`

	// Package - The SoftLayer_Product_Package an order item is a part of.
	Package *SoftLayer_Product_Package `json:"package"`

	// UpgradeItem - The next SoftLayer_Product_Item in the upgrade path for this order item.
	UpgradeItem *SoftLayer_Product_Item `json:"upgradeItem"`
}

func (softlayer_billing_order_item *SoftLayer_Billing_Order_Item_Extended) String() string {
	return "SoftLayer_Billing_Order_Item"
}
