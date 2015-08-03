package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Billing_Order_Cart - <nil>
type SoftLayer_Billing_Order_Cart struct {
}

func (softlayer_billing_order_cart *SoftLayer_Billing_Order_Cart) String() string {
	return "SoftLayer_Billing_Order_Cart"
}

// CreateCart - When creating a new cart, the order data is sent through
// SoftLayer_Product_Order::verifyOrder to make sure that the cart contains valid data. If an issue is
// found with the order, an exception will be thrown and you will receive the same response as if
// SoftLayer_Product_Order::verifyOrder were called directly. Once the order verification is complete,
// the cart will be created. The response is the new cart id.
func (softlayer_billing_order_cart *SoftLayer_Billing_Order_Cart) CreateCart(ctx *slapi.RequestContext, orderData SoftLayer_Container_Product_Order) (int, error) {
	var returnValue int
	return returnValue, nil
}

// DeleteCart - If a cart is no longer needed, it can be deleted using this service. Once a cart has
// been deleted, it cannot be retrieved again.
func (softlayer_billing_order_cart *SoftLayer_Billing_Order_Cart) DeleteCart(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetCartByCartKey - Retrieve a valid cart record of a SoftLayer order.
func (softlayer_billing_order_cart *SoftLayer_Billing_Order_Cart) GetCartByCartKey(ctx *slapi.RequestContext, cartKey string) (*SoftLayer_Billing_Order_Cart, error) {
	var returnValue *SoftLayer_Billing_Order_Cart
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_billing_order_cart *SoftLayer_Billing_Order_Cart) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Billing_Order_Cart, error) {
	var returnValue *SoftLayer_Billing_Order_Cart
	return returnValue, nil
}

// GetPdf - no documentation
func (softlayer_billing_order_cart *SoftLayer_Billing_Order_Cart) GetPdf(ctx *slapi.RequestContext) (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetRecalculatedOrderContainer - This method allows the customer to retrieve a saved cart and put it
// in a format that's suitable to be sent to SoftLayer_Billing_Order_Cart::createCart to create a new
// cart or to SoftLayer_Billing_Order_Cart::updateCart to update an existing cart.
func (softlayer_billing_order_cart *SoftLayer_Billing_Order_Cart) GetRecalculatedOrderContainer(ctx *slapi.RequestContext, orderData SoftLayer_Container_Product_Order, orderBeingPlacedFlag bool) (*SoftLayer_Container_Product_Order, error) {
	var returnValue *SoftLayer_Container_Product_Order
	return returnValue, nil
}

// UpdateCart - Like SoftLayer_Billing_Order_Cart::createCart, the order data will be sent through
// SoftLayer_Product_Order::verifyOrder to make sure that the updated cart information is valid. Once
// it has been verified, the new order data will be saved. This will return the cart id.
func (softlayer_billing_order_cart *SoftLayer_Billing_Order_Cart) UpdateCart(ctx *slapi.RequestContext, orderData SoftLayer_Container_Product_Order) (int, error) {
	var returnValue int
	return returnValue, nil
}
