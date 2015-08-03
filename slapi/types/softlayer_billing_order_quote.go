package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
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
