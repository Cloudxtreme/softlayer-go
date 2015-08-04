package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Product_Order_Receipt - When an order is placed
// (SoftLayer_Product_Order::placeOrder), a receipt is returned when the order is created successfully.
// The information in the receipt helps explain information about the order. It's order ID, and all the
// data within the order as well. For PayPal Orders, an URL is also returned to the user so that the
// user can complete the transaction. Users paying with PayPal must continue on to this login and pay.
// When doing this, PayPal will redirect the user back to a SoftLayer page which will then "finalize"
// the authorization process. From here, Sales will verify the order by contacting the user in some
// way, unless sales has already spoken to the user about approving the order. For users paying with a
// credit card, a receipt means the order has gone to sales and is awaiting approval.
type SoftLayer_Container_Product_Order_Receipt struct {

	// OrderDetails - This is a copy of the order container (SoftLayer_Container_Product_Order) which holds
	// all the data related to an order. This will only return when an order is processed successfully. It
	// will contain all the items in an order as well as the order totals.
	OrderDetails *SoftLayer_Container_Product_Order `json:"orderDetails,omitempty"`

	// OrderId - no documentation
	OrderId int `json:"orderId,omitempty"`

	// PaypalCheckoutUrl - The paypal Token, if this order was processed via PayPal. If a token exists,
	// proceed to the address given. This will tell PayPal that you wish to purchase services. Upon
	// completion of the order at PayPal, you will be directed back to SoftLayer. When this happens, We
	// will be alerted and will set the order to a status that alerts the sales team to approve the order
	// in our system. When this happens, it will be immediately provisioned and the server will be online
	// shortly!
	PaypalCheckoutUrl string `json:"paypalCheckoutUrl,omitempty"`

	// PaypalToken - The paypal Token, if this order was processed via PayPal. The paypalCheckoutUrl should
	// also be populated. This will tell you where to go next to continue the order process.
	PaypalToken string `json:"paypalToken,omitempty"`

	// PlacedOrder - This is a copy of the order that was successfully placed (SoftLayer_Billing_Order).
	// This will only return when an order is processed successfully.
	PlacedOrder *SoftLayer_Billing_Order `json:"placedOrder,omitempty"`

	// Quote - This is a copy of the quote container (SoftLayer_Billing_Order_Quote) which holds all the
	// data related to a quote. This will only return when a quote is processed successfully.
	Quote *SoftLayer_Billing_Order_Quote `json:"quote,omitempty"`

	// OrderDate - no documentation
	OrderDate *time.Time `json:"orderDate,omitempty"`
}

func (softlayer_container_product_order_receipt *SoftLayer_Container_Product_Order_Receipt) String() string {
	return "SoftLayer_Container_Product_Order_Receipt"
}
