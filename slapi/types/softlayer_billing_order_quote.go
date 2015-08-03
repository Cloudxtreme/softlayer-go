package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Billing_Order_Quote - The SoftLayer_Billing_Oder_Quote data type contains general
// information relating to an individual order applied to a SoftLayer customer account or to a new
// customer. Personal information in this type such as names, addresses, and phone numbers are taken
// from the account's contact information at the time the quote is generated for existing SoftLayer
// customer.
type SoftLayer_Billing_Order_Quote struct {

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// AccountId - Identification Number of the account record tied to the quote
	AccountId int `json:"accountId"`

	// CompletedPurchaseDataId - Identification Number of the order record tied to the quote.
	CompletedPurchaseDataId int `json:"completedPurchaseDataId"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// ExpirationDate - This property holds the date of expiration of a quote, after that date the quote
	// would be deem expired
	ExpirationDate *time.Time `json:"expirationDate"`

	// Id - no documentation
	Id int `json:"id"`

	// ModifyDate - Holds the date when the quote record was modified with reference to its creation date
	ModifyDate *time.Time `json:"modifyDate"`

	// Name - no documentation
	Name string `json:"name"`

	// Order - no documentation
	Order *SoftLayer_Billing_Order `json:"order"`

	// OrdersFromQuote - no documentation
	OrdersFromQuote []*SoftLayer_Billing_Order `json:"ordersFromQuote"`

	// OrdersFromQuoteCount - no documentation
	OrdersFromQuoteCount uint64 `json:"ordersFromQuoteCount"`

	// PublicNote - This property Holds system generated notes. In our case if a quote is tied to an order
	// where one of the order item has an inactive promotion code, the quote will be considered invalid.
	PublicNote string `json:"publicNote"`

	// QuoteKey - no documentation
	QuoteKey string `json:"quoteKey"`

	// Status - This property Holds the current status of a Quote: pending,expired, saved or deleted
	Status string `json:"status"`
}

func (softlayer_billing_order_quote *SoftLayer_Billing_Order_Quote) String() string {
	return "SoftLayer_Billing_Order_Quote"
}

// Claim - This method is used to transfer an anonymous quote to the active user and associated
// account. An anonymous quote is one that was created by a user without being authenticated. If a
// quote was created anonymously and then the customer attempts to access that anonymous quote via the
// API (which requires authentication), the customer will be unable to retrieve the quote due to the
// security restrictions in place. By providing the ability for a customer to claim a quote, s/he will
// be able to pull the anonymous quote onto his/her account and successfully view the quote. To claim a
// quote, both the quote id and the quote key (the 32-character random string) must be provided.
func (softlayer_billing_order_quote *SoftLayer_Billing_Order_Quote) Claim(ctx *slapi.RequestContext, quoteKey string, quoteId int) (*SoftLayer_Billing_Order_Quote, error) {
	var returnValue *SoftLayer_Billing_Order_Quote
	return returnValue, nil
}

// DeleteQuote - Account master users and sub-users in the SoftLayer customer portal can delete the
// quote of an order.
func (softlayer_billing_order_quote *SoftLayer_Billing_Order_Quote) DeleteQuote(ctx *slapi.RequestContext) (*SoftLayer_Billing_Order_Quote, error) {
	var returnValue *SoftLayer_Billing_Order_Quote
	return returnValue, nil
}

// GetObject - getObject retrieves the SoftLayer_Billing_Order_Quote object whose ID number corresponds
// to the ID number of the init parameter passed to the SoftLayer_Billing_Order_Quote service. You can
// only retrieve quotes that are assigned to your portal user's account.
func (softlayer_billing_order_quote *SoftLayer_Billing_Order_Quote) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Billing_Order_Quote, error) {
	var returnValue *SoftLayer_Billing_Order_Quote
	return returnValue, nil
}

// GetPdf - Retrieve a PDF record of a SoftLayer quoted order. SoftLayer keeps PDF records of all
// quoted orders for customer retrieval from the portal and You must have a PDF reader installed in
// order to view these quoted order files.
func (softlayer_billing_order_quote *SoftLayer_Billing_Order_Quote) GetPdf(ctx *slapi.RequestContext) (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetQuoteByQuoteKey - Retrieve a valid quote record of a SoftLayer order. Quote whose promotion code
// expired or one of the items is no longer available will not be retrieved.
func (softlayer_billing_order_quote *SoftLayer_Billing_Order_Quote) GetQuoteByQuoteKey(ctx *slapi.RequestContext, quoteKey string) (*SoftLayer_Billing_Order_Quote, error) {
	var returnValue *SoftLayer_Billing_Order_Quote
	return returnValue, nil
}

// GetRecalculatedOrderContainer - Get a SoftLayer_Container_Product_Order with all the recalculated
// total with considerations for promotions, reseller status and taxes.
func (softlayer_billing_order_quote *SoftLayer_Billing_Order_Quote) GetRecalculatedOrderContainer(ctx *slapi.RequestContext, orderData SoftLayer_Container_Product_Order, orderBeingPlacedFlag bool) (*SoftLayer_Container_Product_Order, error) {
	var returnValue *SoftLayer_Container_Product_Order
	return returnValue, nil
}

// PlaceOrder - Use this method for placing server orders and additional services orders. The same
// applies for this as with verifyOrder. Send in the SoftLayer_Container_Product_Order_Hardware_Server
// for server orders. In addition to verifying the order, placeOrder() also makes an initial
// authorization on the SoftLayer_Account tied to this order, if a credit card is on file. If the
// account tied to this order is a paypal customer, an URL will also be returned to the customer. After
// placing the order, you must go to this URL to finish the authorization process. This tells paypal
// that you indeed want to place the order. After going to this it will direct you back to a SoftLayer
// webpage that tells us you have finished the process. After this, it will go to sales for final
// approval.
func (softlayer_billing_order_quote *SoftLayer_Billing_Order_Quote) PlaceOrder(ctx *slapi.RequestContext, orderData SoftLayer_Container_Product_Order) (*SoftLayer_Container_Product_Order_Receipt, error) {
	var returnValue *SoftLayer_Container_Product_Order_Receipt
	return returnValue, nil
}

// PlaceQuote - Use this method for placing server quotes and additional services quotes. The same
// applies for this as with verifyOrder. Send in the SoftLayer_Container_Product_Order_Hardware_Server
// for server quotes. In addition to verifying the quote, placeQuote() also makes an initial
// authorization on the SoftLayer_Account tied to this order, if a credit card is on file. If the
// account tied to this order is a paypal customer, an URL will also be returned to the customer. After
// placing the order, you must go to this URL to finish the authorization process. This tells paypal
// that you indeed want to place the order. After going to this it will direct you back to a SoftLayer
// webpage that tells us you have finished the process.
func (softlayer_billing_order_quote *SoftLayer_Billing_Order_Quote) PlaceQuote(ctx *slapi.RequestContext, orderData SoftLayer_Container_Product_Order) (*SoftLayer_Container_Product_Order, error) {
	var returnValue *SoftLayer_Container_Product_Order
	return returnValue, nil
}

// SaveQuote - Account master users and sub-users in the SoftLayer customer portal can save the quote
// of an order to avoid its deletion after 5 days or its expiration after 2 days.
func (softlayer_billing_order_quote *SoftLayer_Billing_Order_Quote) SaveQuote(ctx *slapi.RequestContext) (*SoftLayer_Billing_Order_Quote, error) {
	var returnValue *SoftLayer_Billing_Order_Quote
	return returnValue, nil
}

// VerifyOrder - Use this method for placing server orders and additional services orders. The same
// applies for this as with verifyOrder. Send in the SoftLayer_Container_Product_Order_Hardware_Server
// for server orders. In addition to verifying the order, placeOrder() also makes an initial
// authorization on the SoftLayer_Account tied to this order, if a credit card is on file. If the
// account tied to this order is a paypal customer, an URL will also be returned to the customer. After
// placing the order, you must go to this URL to finish the authorization process. This tells paypal
// that you indeed want to place the order. After going to this it will direct you back to a SoftLayer
// webpage that tells us you have finished the process. After this, it will go to sales for final
// approval.
func (softlayer_billing_order_quote *SoftLayer_Billing_Order_Quote) VerifyOrder(ctx *slapi.RequestContext, orderData SoftLayer_Container_Product_Order) (*SoftLayer_Container_Product_Order, error) {
	var returnValue *SoftLayer_Container_Product_Order
	return returnValue, nil
}
