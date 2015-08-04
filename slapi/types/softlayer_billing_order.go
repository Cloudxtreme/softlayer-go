package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Billing_Order - The SoftLayer_Billing_Order data type contains general information
// relating to an individual order applied to a SoftLayer customer account or to a new customer.
// Personal information in this type such as names, addresses, and phone numbers are taken from the
// account's contact information at the time the order is generated for existing SoftLayer customer.
type SoftLayer_Billing_Order struct {

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// PrivateCloudOrderFlag - Flag indicating a private cloud solution order (Deprecated)
	PrivateCloudOrderFlag bool `json:"privateCloudOrderFlag,omitempty"`

	// CreateDate - The point in time at which a billing item was created.
	CreateDate *time.Time `json:"createDate,omitempty"`

	// OrderTypeId - no documentation
	OrderTypeId int `json:"orderTypeId,omitempty"`

	// OrderQuoteId - The SoftLayer_Billing_Order_Quote id of the quote's user who finalized an order.
	OrderQuoteId int `json:"orderQuoteId,omitempty"`

	// AccountId - no documentation
	AccountId int `json:"accountId,omitempty"`

	// ImpersonatingUserRecordId - The SoftLayer_User_Customer id of the portal or API user who
	// impersonated the user which submitted an order.
	ImpersonatingUserRecordId int `json:"impersonatingUserRecordId,omitempty"`

	// UserRecordId - The SoftLayer_User_Customer id of the portal or API user who submitted an order.
	UserRecordId int `json:"userRecordId,omitempty"`

	// PresaleEventId - <nil>
	PresaleEventId int `json:"presaleEventId,omitempty"`

	// Status - Purchaser current status e.i. Approved, Pending_Approval
	Status string `json:"status,omitempty"`
}

func (softlayer_billing_order *SoftLayer_Billing_Order) String() string {
	return "SoftLayer_Billing_Order"
}

// SoftLayer_Billing_Order_Extended is SoftLayer_Billing_Order with all maskable types.
type SoftLayer_Billing_Order_Extended struct {
	SoftLayer_Billing_Order

	// OrderTopLevelItemCount - A count of an order's top level items. This normally includes the server
	// line item and any non-server additional services such as NAS or
	OrderTopLevelItemCount uint64 `json:"orderTopLevelItemCount,omitempty"`

	// Brand - <nil>
	Brand *SoftLayer_Brand `json:"brand,omitempty"`

	// CoreRestrictedItems - The SoftLayer_Billing_Order_items that are core restricted
	CoreRestrictedItems []*SoftLayer_Billing_Order_Item `json:"coreRestrictedItems,omitempty"`

	// ExchangeRate - <nil>
	ExchangeRate *SoftLayer_Billing_Currency_ExchangeRate `json:"exchangeRate,omitempty"`

	// OrderTotalAmount - This amount represents the order's initial charge including set up fee and taxes.
	OrderTotalAmount float64 `json:"orderTotalAmount,omitempty"`

	// Quote - The quote of an order. This quote holds information about its expiration date, creation
	// date, name and status. This information is tied to an order having the status
	Quote *SoftLayer_Billing_Order_Quote `json:"quote,omitempty"`

	// InitialInvoice - <nil>
	InitialInvoice *SoftLayer_Billing_Invoice `json:"initialInvoice,omitempty"`

	// OrderNonServerMonthlyAmount - no documentation
	OrderNonServerMonthlyAmount float64 `json:"orderNonServerMonthlyAmount,omitempty"`

	// OrderTotalRecurringTaxAmount - The total tax amount of the recurring fees, if the SoftLayer_Account
	// tied to a SoftLayer_Billing_Order is a taxable account.
	OrderTotalRecurringTaxAmount float64 `json:"orderTotalRecurringTaxAmount,omitempty"`

	// PaypalTransactionCount - A count of all PayPal transactions associated with this order. If this
	// order was not placed with PayPal, this will be empty.
	PaypalTransactionCount uint64 `json:"paypalTransactionCount,omitempty"`

	// OrderTotalRecurring - An order's total recurring amount. Taxes will be applied for non-tax-exempt.
	// This amount represents the fees that will be charged on a recurring (usually monthly) basis.
	OrderTotalRecurring float64 `json:"orderTotalRecurring,omitempty"`

	// ReferralPartner - The Referral Partner who referred this order. (Only necessary for new customer
	// orders)
	ReferralPartner *SoftLayer_Account `json:"referralPartner,omitempty"`

	// UserRecord - The SoftLayer_User_Customer object tied to an order.
	UserRecord *SoftLayer_User_Customer `json:"userRecord,omitempty"`

	// Items - The SoftLayer_Billing_Order_items included in an order.
	Items []*SoftLayer_Billing_Order_Item `json:"items,omitempty"`

	// OrderServerMonthlyAmount - no documentation
	OrderServerMonthlyAmount float64 `json:"orderServerMonthlyAmount,omitempty"`

	// OrderTotalRecurringAmount - An order's total recurring amount. This amount represents the fees that
	// will be charged on a recurring (usually monthly) basis.
	OrderTotalRecurringAmount float64 `json:"orderTotalRecurringAmount,omitempty"`

	// ItemCount - A count of the SoftLayer_Billing_Order_items included in an order.
	ItemCount uint64 `json:"itemCount,omitempty"`

	// OrderTopLevelItems - An order's top level items. This normally includes the server line item and any
	// non-server additional services such as NAS or
	OrderTopLevelItems []*SoftLayer_Billing_Order_Item `json:"orderTopLevelItems,omitempty"`

	// OrderTotalOneTimeTaxAmount - An order's total one time tax amount. This amount represents the tax
	// that will be applied to the total charge, if the SoftLayer_Account tied to a SoftLayer_Billing_Order
	// is a taxable account.
	OrderTotalOneTimeTaxAmount float64 `json:"orderTotalOneTimeTaxAmount,omitempty"`

	// CreditCardTransactionCount - A count of all credit card transactions associated with this order. If
	// this order was not placed with a credit card, this will be empty.
	CreditCardTransactionCount uint64 `json:"creditCardTransactionCount,omitempty"`

	// CreditCardTransactions - All credit card transactions associated with this order. If this order was
	// not placed with a credit card, this will be empty.
	CreditCardTransactions []*SoftLayer_Billing_Payment_Card_Transaction `json:"creditCardTransactions,omitempty"`

	// UpgradeRequestFlag - no documentation
	UpgradeRequestFlag bool `json:"upgradeRequestFlag,omitempty"`

	// Cart - A cart is similar to a quote, except that it can be continually modified by the customer and
	// does not have locked-in prices. Not all orders will have a cart associated with them. See
	// [[SoftLayer_Billing_Order_Cart]] for more information.
	Cart *SoftLayer_Billing_Order_Cart `json:"cart,omitempty"`

	// OrderType - The type of an order. This lets you know where this order was generated from.
	OrderType *SoftLayer_Billing_Order_Type `json:"orderType,omitempty"`

	// PresaleEvent - <nil>
	PresaleEvent *SoftLayer_Sales_Presale_Event `json:"presaleEvent,omitempty"`

	// PaypalTransactions - All PayPal transactions associated with this order. If this order was not
	// placed with PayPal, this will be empty.
	PaypalTransactions []*SoftLayer_Billing_Payment_PayPal_Transaction `json:"paypalTransactions,omitempty"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account,omitempty"`

	// OrderApprovalDate - <nil>
	OrderApprovalDate *time.Time `json:"orderApprovalDate,omitempty"`

	// OrderTotalOneTime - An order's total one time amount summing all the set up fees, the labor fees and
	// the one time fees. Taxes will be applied for non-tax-exempt. This amount represents the initial fees
	// that will be charged.
	OrderTotalOneTime float64 `json:"orderTotalOneTime,omitempty"`

	// OrderTotalOneTimeAmount - An order's total one time amount. This amount represents the initial fees
	// before tax.
	OrderTotalOneTimeAmount float64 `json:"orderTotalOneTimeAmount,omitempty"`

	// OrderTotalSetupAmount - no documentation
	OrderTotalSetupAmount float64 `json:"orderTotalSetupAmount,omitempty"`

	// CoreRestrictedItemCount - A count of the SoftLayer_Billing_Order_items that are core restricted
	CoreRestrictedItemCount uint64 `json:"coreRestrictedItemCount,omitempty"`
}

func (softlayer_billing_order *SoftLayer_Billing_Order_Extended) String() string {
	return "SoftLayer_Billing_Order"
}
