package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"

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

	// ParentId - <nil>
	ParentId int `json:"parentId,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// LaborFeeTaxRate - The rate at which labor fees are taxed if you are a taxable customer.
	LaborFeeTaxRate slapi.Float64 `json:"laborFeeTaxRate,omitempty"`

	// OneTimeAfterTaxAmount - An order item's one-time fee total after taxes. This does not include any
	// child invoice items.
	OneTimeAfterTaxAmount slapi.Float64 `json:"oneTimeAfterTaxAmount,omitempty"`

	// SetupFeeTaxRate - The rate at which setup fees are taxed if you are a taxable customer.
	SetupFeeTaxRate slapi.Float64 `json:"setupFeeTaxRate,omitempty"`

	// HourlyRecurringFee - The amount of money charged per hourly for an order item, if applicable, and
	// only if it was ordered this day. hourlyRecurringFee is measured in US Dollars
	HourlyRecurringFee slapi.Float64 `json:"hourlyRecurringFee,omitempty"`

	// ItemId - no documentation
	ItemId int `json:"itemId,omitempty"`

	// OneTimeFee - The amount of money charged as a one-time charge for an order item, if applicable.
	// oneTimeFee is measured in US Dollars
	OneTimeFee slapi.Float64 `json:"oneTimeFee,omitempty"`

	// ItemPriceId - the item price id (SoftLayer_Product_Item_Price->id) of the ordered item.
	ItemPriceId slapi.Float64 `json:"itemPriceId,omitempty"`

	// Quantity - no documentation
	Quantity int `json:"quantity,omitempty"`

	// SetupTaxAmount - An order item's setup tax amount. This does not include any child invoice items.
	SetupTaxAmount slapi.Float64 `json:"setupTaxAmount,omitempty"`

	// OneTimeFeeTaxRate - The rate at which one time fees are taxed if you are a taxable customer.
	OneTimeFeeTaxRate slapi.Float64 `json:"oneTimeFeeTaxRate,omitempty"`

	// RecurringAfterTaxAmount - An order item's recurring fee total after taxes. This does not include any
	// child invoice items.
	RecurringAfterTaxAmount slapi.Float64 `json:"recurringAfterTaxAmount,omitempty"`

	// PromoCodeId - <nil>
	PromoCodeId int `json:"promoCodeId,omitempty"`

	// RecurringFee - The amount of money charged per month for an order item, if applicable. recurringFee
	// is measured in US Dollars
	RecurringFee slapi.Float64 `json:"recurringFee,omitempty"`

	// HostName - The hostname of the server as designated by the purchaser at the time of order placement.
	HostName string `json:"hostName,omitempty"`

	// LaborAfterTaxAmount - An order item's labor fee total after taxes. This does not include any child
	// invoice items.
	LaborAfterTaxAmount slapi.Float64 `json:"laborAfterTaxAmount,omitempty"`

	// LaborTaxAmount - An order item's labor tax amount. This does not include any child invoice items.
	LaborTaxAmount slapi.Float64 `json:"laborTaxAmount,omitempty"`

	// RecurringTaxAmount - An order item's recurring tax amount. This does not include any child invoice
	// items.
	RecurringTaxAmount slapi.Float64 `json:"recurringTaxAmount,omitempty"`

	// SetupAfterTaxAmount - An order item's setup fee total after taxes. This does not include any child
	// invoice items.
	SetupAfterTaxAmount slapi.Float64 `json:"setupAfterTaxAmount,omitempty"`

	// DomainName - The domain name of the server as designated by the purchaser at the time of order
	// placement.
	DomainName string `json:"domainName,omitempty"`

	// SetupFee - no documentation
	SetupFee slapi.Float64 `json:"setupFee,omitempty"`

	// CategoryCode - no documentation
	CategoryCode string `json:"categoryCode,omitempty"`

	// OneTimeTaxAmount - An order item's one-time tax amount. This does not include any child invoice
	// items.
	OneTimeTaxAmount slapi.Float64 `json:"oneTimeTaxAmount,omitempty"`

	// SetupFeeDeferralMonths - no documentation
	SetupFeeDeferralMonths int `json:"setupFeeDeferralMonths,omitempty"`

	// Description - no documentation
	Description string `json:"description,omitempty"`

	// LaborFee - no documentation
	LaborFee slapi.Float64 `json:"laborFee,omitempty"`

	// Children - The child order items for an order item. All server order items should have children.
	// These children are considered a part of the server.
	Children []*SoftLayer_Billing_Order_Item `json:"children,omitempty"`

	// TotalRecurringAmount - The recurring fee of an ordered item. This amount represents the fees that
	// will be charged on a recurring (usually monthly) basis.
	TotalRecurringAmount slapi.Float64 `json:"totalRecurringAmount,omitempty"`

	// ItemCategoryAnswers - no documentation
	ItemCategoryAnswers []*SoftLayer_Billing_Order_Item_Category_Answer `json:"itemCategoryAnswers,omitempty"`

	// BillingItem - no documentation
	BillingItem *SoftLayer_Billing_Item `json:"billingItem,omitempty"`

	// GlobalIdentifier - no documentation
	GlobalIdentifier string `json:"globalIdentifier,omitempty"`

	// SoftwareDescription - For ordered items that are software items, a full description of that software
	// can be found with this property.
	SoftwareDescription *SoftLayer_Software_Description `json:"softwareDescription,omitempty"`

	// ItemCategoryAnswerCount - no documentation
	ItemCategoryAnswerCount uint64 `json:"itemCategoryAnswerCount,omitempty"`

	// OrderApprovalDate - <nil>
	OrderApprovalDate *time.Time `json:"orderApprovalDate,omitempty"`

	// BundledItems - no documentation
	BundledItems []*SoftLayer_Billing_Order_Item `json:"bundledItems,omitempty"`

	// UpgradeItem - The next SoftLayer_Product_Item in the upgrade path for this order item.
	UpgradeItem *SoftLayer_Product_Item `json:"upgradeItem,omitempty"`

	// Parent - The parent order item ID for an item. Items that are associated with a server will have a
	// parent. The parent will be the server item itself.
	Parent *SoftLayer_Billing_Order_Item `json:"parent,omitempty"`

	// BundledItemCount - A count of the other items included with an ordered item.
	BundledItemCount uint64 `json:"bundledItemCount,omitempty"`

	// HardwareGenericComponent - The component type tied to an order item. All hardware-specific items
	// should have a generic hardware component.
	HardwareGenericComponent *SoftLayer_Hardware_Component_Model_Generic `json:"hardwareGenericComponent,omitempty"`

	// OldBillingItem - This is only populated when an upgrade order is placed. The old billing item
	// represents what the billing was before the upgrade happened.
	OldBillingItem *SoftLayer_Billing_Item `json:"oldBillingItem,omitempty"`

	// Order - The order to which this item belongs. The order contains all the information related to the
	// items included in an order
	Order *SoftLayer_Billing_Order `json:"order,omitempty"`

	// Package - The SoftLayer_Product_Package an order item is a part of.
	Package *SoftLayer_Product_Package `json:"package,omitempty"`

	// Category - no documentation
	Category *SoftLayer_Product_Item_Category `json:"category,omitempty"`

	// RedundantPowerSupplyCount - A count of power supplies contained within this SoftLayer_Billing_Order
	RedundantPowerSupplyCount uint `json:"redundantPowerSupplyCount,omitempty"`

	// NextOrderChildrenCount - no documentation
	NextOrderChildrenCount uint64 `json:"nextOrderChildrenCount,omitempty"`

	// ItemPrice - The SoftLayer_Product_Item_Price tied to an order item. The item price object describes
	// the cost of an item.
	ItemPrice *SoftLayer_Product_Item_Price `json:"itemPrice,omitempty"`

	// Location - The location of an ordered item. This is usually the same as the server it is being
	// ordered with. Otherwise it describes the location of the additional service being ordered.
	Location *SoftLayer_Location `json:"location,omitempty"`

	// StorageGroups - The drive storage groups that are attached to this billing order item.
	StorageGroups []*SoftLayer_Configuration_Storage_Group_Order `json:"storageGroups,omitempty"`

	// Item - The SoftLayer_Product_Item tied to an order item. The item is the actual definition of the
	// product being sold.
	Item *SoftLayer_Product_Item `json:"item,omitempty"`

	// StorageGroupCount - A count of the drive storage groups that are attached to this billing order
	// item.
	StorageGroupCount uint64 `json:"storageGroupCount,omitempty"`

	// ChildrenCount - A count of the child order items for an order item. All server order items should
	// have children. These children are considered a part of the server.
	ChildrenCount uint64 `json:"childrenCount,omitempty"`

	// NextOrderChildren - <nil>
	NextOrderChildren []*SoftLayer_Billing_Order_Item `json:"nextOrderChildren,omitempty"`
}

func (softlayer_billing_order_item *SoftLayer_Billing_Order_Item) String() string {
	return "SoftLayer_Billing_Order_Item"
}
