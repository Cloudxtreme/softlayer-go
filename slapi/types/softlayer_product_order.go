package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Product_Order - <nil>
type SoftLayer_Product_Order struct {
}

// CheckItemAvailability - <nil>
func (softlayer_product_order *SoftLayer_Product_Order) CheckItemAvailability(ctx *slapi.RequestContext, itemPrices []SoftLayer_Product_Item_Price, accountId int, availabilityTypeKeyNames []string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// CheckItemAvailabilityForImageTemplate - <nil>
func (softlayer_product_order *SoftLayer_Product_Order) CheckItemAvailabilityForImageTemplate(ctx *slapi.RequestContext, imageTemplateId int, accountId int, packageId int, availabilityTypeKeyNames []string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// CheckItemConflicts - no documentation
func (softlayer_product_order *SoftLayer_Product_Order) CheckItemConflicts(ctx *slapi.RequestContext, itemPrices []SoftLayer_Product_Item_Price) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetExternalPaymentAuthorizationReceipt - This method simply returns a receipt for a previously
// finalized payment authorization from PayPal. The response matches the response returned from
// placeOrder when the order was originally placed with PayPal as the payment type.
func (softlayer_product_order *SoftLayer_Product_Order) GetExternalPaymentAuthorizationReceipt(ctx *slapi.RequestContext, token string, payerId string) (*SoftLayer_Container_Product_Order_Receipt, error) {
	var returnValue *SoftLayer_Container_Product_Order_Receipt
	return returnValue, nil
}

// GetResellerOrder - When the account is on an external reseller brand, this service will provide a
// SoftLayer_Product_Order with the the pricing adjusted by the external reseller.
func (softlayer_product_order *SoftLayer_Product_Order) GetResellerOrder(ctx *slapi.RequestContext, orderContainer SoftLayer_Container_Product_Order) (*SoftLayer_Container_Product_Order, error) {
	var returnValue *SoftLayer_Container_Product_Order
	return returnValue, nil
}

// GetTaxCalculationResult - Sometimes taxes cannot be calculated immediately, so we start the
// calculations and let them run in the background. This method will return the current progress and
// information related to a specific tax calculation, which allows real-time progress updates on tax
// calculations.
func (softlayer_product_order *SoftLayer_Product_Order) GetTaxCalculationResult(ctx *slapi.RequestContext, orderHash string) (*SoftLayer_Container_Tax_Cache, error) {
	var returnValue *SoftLayer_Container_Tax_Cache
	return returnValue, nil
}

// GetVlans - Return collections of public and private VLANs that are available during ordering. If a
// location ID is provided, the resulting VLANs will be limited to that location. If the Virtual Server
// package id (46) is provided, the VLANs will be narrowed down to those locations that contain routers
// with the data attribute. For the selectedItems parameter, this is a comma-separated string of
// category codes and item values. For example: This parameter is used to narrow the available results
// down even further. It's not necessary when selecting a but it will help avoid errors when attempting
// to place an order. The only acceptable category codes are: A disk category, such as guest_disk0 or
// disk0 , with values of either or For most customers, it's sufficient to only provide the first 2
// parameters.
func (softlayer_product_order *SoftLayer_Product_Order) GetVlans(ctx *slapi.RequestContext, locationId int, packageId int, selectedItems string, vlanIds []int, subnetIds []int, accountId int) (*SoftLayer_Container_Product_Order_Network_Vlans, error) {
	var returnValue *SoftLayer_Container_Product_Order_Network_Vlans
	return returnValue, nil
}

// PlaceOrder - Use this method to place bare metal server, virtual server and additional service
// orders with SoftLayer. Upon success, your credit card or PayPal account will incur charges for the
// monthly order total (or prorated value if ordered mid billing cycle). If all products on the order
// are only billed hourly, you will be charged on your billing anniversary date, which occurs monthly
// on the day you ordered your first service with SoftLayer. For new customers, you are required to
// provide billing information when you place an order. For existing customers, the credit card on file
// will be charged. If you're a PayPal customer, a URL will be returned from the call to
// [[SoftLayer_Product_Order/placeOrder|placeOrder]] which is to be used to finish the authorization
// process. This authorization tells PayPal that you indeed want to place an order with SoftLayer. From
// PayPal's web site, you will be redirected back to SoftLayer for your order receipt. When an order is
// placed, your order will be in a "pending approval" state. When all internal checks pass, your order
// will be automatically approved. For orders that may need extra attention, a Sales representative
// will review the order and contact you if necessary. Once the order is approved, your server or
// service will be provisioned and available to you shortly thereafter. Depending on the type of server
// or service ordered, provisioning times will vary. Order Containers When placing API orders, it's
// important to order your server and services on the appropriate [[SoftLayer_Container_Product_Order
// (type)|order container]]. Failing to provide the correct container may delay your server or service
// from being provisioned in a timely manner. Some common order containers are included below. Note:
// SoftLayer_Container_Product_Order_ has been removed from the containers in the table below for
// readability. Product Order container Package type Bare metal server by
// [[SoftLayer_Container_Product_Order_Hardware_Server (type)|Hardware_Server]] Bare metal server by
// core[[SoftLayer_Container_Product_Order_Hardware_Server (type)|Hardware_Server]] Virtual
// server[[SoftLayer_Container_Product_Order_Virtual_Guest (type)|Virtual_Guest]] DNS domain
// registration[[SoftLayer_Container_Product_Order_Dns_Domain_Registration
// (type)|Dns_Domain_Registration]] Local & dedicated load
// balancers[[SoftLayer_Container_Product_Order_Network_LoadBalancer (type)|Network_LoadBalancer]]
// Content delivery network[[SoftLayer_Container_Product_Order_Network_ContentDelivery_Account
// (type)|Network_ContentDelivery_Account]] Message
// queue[[SoftLayer_Container_Product_Order_Network_Message_Queue (type)|Network_Message_Queue]]
// Hardware & software firewalls[[SoftLayer_Container_Product_Order_Network_Protection_Firewall
// (type)|Network_Protection_Firewall]] Dedicated
// firewall[[SoftLayer_Container_Product_Order_Network_Protection_Firewall_Dedicated
// (type)|Network_Protection_Firewall_Dedicated]] Object
// storage[[SoftLayer_Container_Product_Order_Network_Storage_Hub (type)|Network_Storage_Hub]] Network
// attached storage[[SoftLayer_Container_Product_Order_Network_Storage_Nas (type)|Network_Storage_Nas]]
// Iscsi storage[[SoftLayer_Container_Product_Order_Network_Storage_Iscsi
// (type)|Network_Storage_Iscsi]]
// Evault[[SoftLayer_Container_Product_Order_Network_Storage_Backup_Evault_Vault
// (type)|Network_Storage_Backup_Evault_Vault]] Evault
// Plugin[[SoftLayer_Container_Product_Order_Network_Storage_Backup_Evault_Plugin
// (type)|Network_Storage_Backup_Evault_Plugin]] Application delivery
// appliance[[SoftLayer_Container_Product_Order_Network_Application_Delivery_Controller
// (type)|Network_Application_Delivery_Controller]] Network
// subnet[[SoftLayer_Container_Product_Order_Network_Subnet (type)|Network_Subnet]] Network
// [[SoftLayer_Container_Product_Order_Network_Vlan (type)|Network_Vlan]] Portable
// storage[[SoftLayer_Container_Product_Order_Virtual_Disk_Image (type)|Virtual_Disk_Image]] SSL
// certificate[[SoftLayer_Container_Product_Order_Security_Certificate (type)|Security_Certificate]]
// External authentication[[SoftLayer_Container_Product_Order_User_Customer_External_Binding
// (type)|User_Customer_External_Binding]] Server example This example includes a single bare metal
// server being ordered with monthly billing. Warning: the price ids provided below may be outdated or
// unavailable, so you will need to determine the available prices from the bare metal server
// [[SoftLayer_Product_Package/getAllObjects|packages]], which have a [[SoftLayer_Product_Package_Type
// (type)|package type]] of or You can get a full list of
// [[SoftLayer_Product_Package_Type/getAllObjects|package types]] to see other potentially available
// server packages. Virtual server example This example includes 2 identical virtual servers (except
// for hostname) being ordered for hourly billing. It includes an optional image template id and data
// specified on the virtualGuest objects - primaryBackendNetworkComponent and primaryNetworkComponent
// Warning: the price ids provided below may be outdated or unavailable, so you will need to determine
// the available prices from the virtual server [[SoftLayer_Product_Package/getAllObjects|package]],
// which has a [[SoftLayer_Product_Package_Type (type)|package type]] of example Warning: the price ids
// provided below may be outdated or unavailable, so you will need to determine the available prices
// from the additional services [[SoftLayer_Product_Package/getAllObjects|package]], which has a
// [[SoftLayer_Product_Package_Type (type)|package type]] of You can get a full list of
// [[SoftLayer_Product_Package_Type/getAllObjects|package types]] to find other available additional
// service packages. Multiple products example This example includes a combination of the above
// examples in a single order. Note that all the configuration options for each individual order
// container are the same as above, except now we encapsulate each one within the orderContainers
// property on the base [[SoftLayer_Container_Product_Order (type)|order container]]. Warning: not all
// products are available to be ordered with other products. For example, since SSL certificates
// require validation from a 3rd party, the approval process may take days or even weeks, and this
// would not be acceptable when you need your hourly virtual server right now. To better accommodate
// customers, we restrict several products to be ordered individually.
func (softlayer_product_order *SoftLayer_Product_Order) PlaceOrder(ctx *slapi.RequestContext, orderData SoftLayer_Container_Product_Order, saveAsQuote bool) (*SoftLayer_Container_Product_Order_Receipt, error) {
	var returnValue *SoftLayer_Container_Product_Order_Receipt
	return returnValue, nil
}

// PlaceQuote - Use this method for placing server quotes and additional services quotes. The same
// applies for this as with verifyOrder. Send in the SoftLayer_Container_Product_Order_Hardware_Server
// for server quotes. After placing the quote, you must go to this URL to finish the order process.
// After going to this it will direct you back to a SoftLayer webpage that tells us you have finished
// the process. After this, it will go to sales for final approval.
func (softlayer_product_order *SoftLayer_Product_Order) PlaceQuote(ctx *slapi.RequestContext, orderData SoftLayer_Container_Product_Order) (*SoftLayer_Container_Product_Order_Receipt, error) {
	var returnValue *SoftLayer_Container_Product_Order_Receipt
	return returnValue, nil
}

// ProcessExternalPaymentAuthorization - This method simply finalizes an authorization from PayPal. It
// tells SoftLayer that the customer has completed the PayPal process. This is needed if you, the
// customer, have your own API into PayPal and wish to automate authorizations from PayPal and our
// system. For most, this method will not be needed. Once an order is placed using placeOrder() for
// PayPal customers, a URL is given back to the customer. In it is the token and PayerID. If you want
// to systematically pay with PayPal, do so then call this method with the token and PayerID.
func (softlayer_product_order *SoftLayer_Product_Order) ProcessExternalPaymentAuthorization(ctx *slapi.RequestContext, token string, payerId string) (*SoftLayer_Container_Product_Order, error) {
	var returnValue *SoftLayer_Container_Product_Order
	return returnValue, nil
}

// RequiredItems - Get list of items that are required with the item prices provided
func (softlayer_product_order *SoftLayer_Product_Order) RequiredItems(ctx *slapi.RequestContext, itemPrices []SoftLayer_Product_Item_Price) ([]*SoftLayer_Product_Item, error) {
	var returnValue []*SoftLayer_Product_Item
	return returnValue, nil
}

// VerifyOrder - This service is used to verify that an order meets all the necessary requirements to
// purchase a server, virtual server or service from SoftLayer. It will verify that the products
// requested do not conflict. For example, you cannot order a Windows firewall with a Linux operating
// system. It will also check to make sure you have provided all the products that are required for the
// [[SoftLayer_Product_Package_Order_Configuration (type)|package configuration]] associated with the
// [[SoftLayer_Product_Package|package id]] on each of the [[SoftLayer_Container_Product_Order
// (type)|order containers]] specified. This service returns the same container that was provided, but
// with additional information that can be used for debugging or validation. It will also contain
// pricing information (prorated if applicable) for each of the products on the order. If an exception
// occurs during verification, a container with the SoftLayer_Exception_Order exception type will be
// specified in the result. verifyOrder accepts the same [[SoftLayer_Container_Product_Order
// (type)|container types]] as placeOrder , so see [[SoftLayer_Product_Order/placeOrder|placeOrder]]
// for more details.
func (softlayer_product_order *SoftLayer_Product_Order) VerifyOrder(ctx *slapi.RequestContext, orderData SoftLayer_Container_Product_Order) (*SoftLayer_Container_Product_Order, error) {
	var returnValue *SoftLayer_Container_Product_Order
	return returnValue, nil
}
