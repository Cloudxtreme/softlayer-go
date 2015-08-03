package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Billing_Order - The SoftLayer_Billing_Order data type contains general information
// relating to an individual order applied to a SoftLayer customer account or to a new customer.
// Personal information in this type such as names, addresses, and phone numbers are taken from the
// account's contact information at the time the order is generated for existing SoftLayer customer.
type SoftLayer_Billing_Order struct {

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// AccountId - no documentation
	AccountId int `json:"accountId"`

	// Brand - <nil>
	Brand *SoftLayer_Brand `json:"brand"`

	// Cart - A cart is similar to a quote, except that it can be continually modified by the customer and
	// does not have locked-in prices. Not all orders will have a cart associated with them. See
	// [[SoftLayer_Billing_Order_Cart]] for more information.
	Cart *SoftLayer_Billing_Order_Cart `json:"cart"`

	// CoreRestrictedItemCount - A count of the SoftLayer_Billing_Order_items that are core restricted
	CoreRestrictedItemCount uint64 `json:"coreRestrictedItemCount"`

	// CoreRestrictedItems - The SoftLayer_Billing_Order_items that are core restricted
	CoreRestrictedItems []*SoftLayer_Billing_Order_Item `json:"coreRestrictedItems"`

	// CreateDate - The point in time at which a billing item was created.
	CreateDate *time.Time `json:"createDate"`

	// CreditCardTransactionCount - A count of all credit card transactions associated with this order. If
	// this order was not placed with a credit card, this will be empty.
	CreditCardTransactionCount uint64 `json:"creditCardTransactionCount"`

	// CreditCardTransactions - All credit card transactions associated with this order. If this order was
	// not placed with a credit card, this will be empty.
	CreditCardTransactions []*SoftLayer_Billing_Payment_Card_Transaction `json:"creditCardTransactions"`

	// ExchangeRate - <nil>
	ExchangeRate *SoftLayer_Billing_Currency_ExchangeRate `json:"exchangeRate"`

	// Id - <nil>
	Id int `json:"id"`

	// ImpersonatingUserRecordId - The SoftLayer_User_Customer id of the portal or API user who
	// impersonated the user which submitted an order.
	ImpersonatingUserRecordId int `json:"impersonatingUserRecordId"`

	// InitialInvoice - <nil>
	InitialInvoice *SoftLayer_Billing_Invoice `json:"initialInvoice"`

	// ItemCount - A count of the SoftLayer_Billing_Order_items included in an order.
	ItemCount uint64 `json:"itemCount"`

	// Items - The SoftLayer_Billing_Order_items included in an order.
	Items []*SoftLayer_Billing_Order_Item `json:"items"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// OrderApprovalDate - <nil>
	OrderApprovalDate *time.Time `json:"orderApprovalDate"`

	// OrderNonServerMonthlyAmount - no documentation
	OrderNonServerMonthlyAmount float64 `json:"orderNonServerMonthlyAmount"`

	// OrderQuoteId - The SoftLayer_Billing_Order_Quote id of the quote's user who finalized an order.
	OrderQuoteId int `json:"orderQuoteId"`

	// OrderServerMonthlyAmount - no documentation
	OrderServerMonthlyAmount float64 `json:"orderServerMonthlyAmount"`

	// OrderTopLevelItemCount - A count of an order's top level items. This normally includes the server
	// line item and any non-server additional services such as NAS or
	OrderTopLevelItemCount uint64 `json:"orderTopLevelItemCount"`

	// OrderTopLevelItems - An order's top level items. This normally includes the server line item and any
	// non-server additional services such as NAS or
	OrderTopLevelItems []*SoftLayer_Billing_Order_Item `json:"orderTopLevelItems"`

	// OrderTotalAmount - This amount represents the order's initial charge including set up fee and taxes.
	OrderTotalAmount float64 `json:"orderTotalAmount"`

	// OrderTotalOneTime - An order's total one time amount summing all the set up fees, the labor fees and
	// the one time fees. Taxes will be applied for non-tax-exempt. This amount represents the initial fees
	// that will be charged.
	OrderTotalOneTime float64 `json:"orderTotalOneTime"`

	// OrderTotalOneTimeAmount - An order's total one time amount. This amount represents the initial fees
	// before tax.
	OrderTotalOneTimeAmount float64 `json:"orderTotalOneTimeAmount"`

	// OrderTotalOneTimeTaxAmount - An order's total one time tax amount. This amount represents the tax
	// that will be applied to the total charge, if the SoftLayer_Account tied to a SoftLayer_Billing_Order
	// is a taxable account.
	OrderTotalOneTimeTaxAmount float64 `json:"orderTotalOneTimeTaxAmount"`

	// OrderTotalRecurring - An order's total recurring amount. Taxes will be applied for non-tax-exempt.
	// This amount represents the fees that will be charged on a recurring (usually monthly) basis.
	OrderTotalRecurring float64 `json:"orderTotalRecurring"`

	// OrderTotalRecurringAmount - An order's total recurring amount. This amount represents the fees that
	// will be charged on a recurring (usually monthly) basis.
	OrderTotalRecurringAmount float64 `json:"orderTotalRecurringAmount"`

	// OrderTotalRecurringTaxAmount - The total tax amount of the recurring fees, if the SoftLayer_Account
	// tied to a SoftLayer_Billing_Order is a taxable account.
	OrderTotalRecurringTaxAmount float64 `json:"orderTotalRecurringTaxAmount"`

	// OrderTotalSetupAmount - no documentation
	OrderTotalSetupAmount float64 `json:"orderTotalSetupAmount"`

	// OrderType - The type of an order. This lets you know where this order was generated from.
	OrderType *SoftLayer_Billing_Order_Type `json:"orderType"`

	// OrderTypeId - no documentation
	OrderTypeId int `json:"orderTypeId"`

	// PaypalTransactionCount - A count of all PayPal transactions associated with this order. If this
	// order was not placed with PayPal, this will be empty.
	PaypalTransactionCount uint64 `json:"paypalTransactionCount"`

	// PaypalTransactions - All PayPal transactions associated with this order. If this order was not
	// placed with PayPal, this will be empty.
	PaypalTransactions []*SoftLayer_Billing_Payment_PayPal_Transaction `json:"paypalTransactions"`

	// PresaleEvent - <nil>
	PresaleEvent *SoftLayer_Sales_Presale_Event `json:"presaleEvent"`

	// PresaleEventId - <nil>
	PresaleEventId int `json:"presaleEventId"`

	// PrivateCloudOrderFlag - Flag indicating a private cloud solution order (Deprecated)
	PrivateCloudOrderFlag bool `json:"privateCloudOrderFlag"`

	// Quote - The quote of an order. This quote holds information about its expiration date, creation
	// date, name and status. This information is tied to an order having the status
	Quote *SoftLayer_Billing_Order_Quote `json:"quote"`

	// ReferralPartner - The Referral Partner who referred this order. (Only necessary for new customer
	// orders)
	ReferralPartner *SoftLayer_Account `json:"referralPartner"`

	// Status - Purchaser current status e.i. Approved, Pending_Approval
	Status string `json:"status"`

	// UpgradeRequestFlag - no documentation
	UpgradeRequestFlag bool `json:"upgradeRequestFlag"`

	// UserRecord - The SoftLayer_User_Customer object tied to an order.
	UserRecord *SoftLayer_User_Customer `json:"userRecord"`

	// UserRecordId - The SoftLayer_User_Customer id of the portal or API user who submitted an order.
	UserRecordId int `json:"userRecordId"`
}

func (softlayer_billing_order *SoftLayer_Billing_Order) String() string {
	return "SoftLayer_Billing_Order"
}

// ApproveModifiedOrder - When an order has been modified, the customer will need to approve the
// changes. This method will allow the customer to approve the changes.
func (softlayer_billing_order *SoftLayer_Billing_Order) ApproveModifiedOrder(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetAllObjects - This will get all billing orders for your account.
func (softlayer_billing_order *SoftLayer_Billing_Order) GetAllObjects(ctx *slapi.RequestContext) ([]*SoftLayer_Billing_Order, error) {
	var returnValue []*SoftLayer_Billing_Order
	return returnValue, nil
}

// GetObject - getObject retrieves the SoftLayer_Billing_Order object whose ID number corresponds to
// the ID number of the init parameter passed to the SoftLayer_Billing_Order service. You can only
// retrieve orders that are assigned to your portal user's account.
func (softlayer_billing_order *SoftLayer_Billing_Order) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Billing_Order, error) {
	var returnValue *SoftLayer_Billing_Order
	return returnValue, nil
}

// GetOrderStatuses - Get a list of [[SoftLayer_Container_Billing_Order_Status]] objects.
func (softlayer_billing_order *SoftLayer_Billing_Order) GetOrderStatuses(ctx *slapi.RequestContext) ([]*SoftLayer_Container_Billing_Order_Status, error) {
	var returnValue []*SoftLayer_Container_Billing_Order_Status
	return returnValue, nil
}

// GetPdf - Retrieve a PDF record of a SoftLayer quote. If the order is not a quote, an error will be
// thrown.
func (softlayer_billing_order *SoftLayer_Billing_Order) GetPdf(ctx *slapi.RequestContext) (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetPdfFilename - no documentation
func (softlayer_billing_order *SoftLayer_Billing_Order) GetPdfFilename(ctx *slapi.RequestContext) (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetRecalculatedOrderContainer - Get a SoftLayer_Container_Product_Order with all the recalculated
// total with considerations for promotions, reseller status and taxes.
func (softlayer_billing_order *SoftLayer_Billing_Order) GetRecalculatedOrderContainer(ctx *slapi.RequestContext, message string, ignoreDiscountsFlag bool) (*SoftLayer_Container_Product_Order, error) {
	var returnValue *SoftLayer_Container_Product_Order
	return returnValue, nil
}

// IsPendingEditApproval - When an order has been modified, it will contain a status indicating so.
// This method checks that status and also verifies that the active user's account is the same as the
// account on the order.
func (softlayer_billing_order *SoftLayer_Billing_Order) IsPendingEditApproval(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}
