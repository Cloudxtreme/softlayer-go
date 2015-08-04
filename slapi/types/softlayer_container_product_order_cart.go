package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Order_Cart - no documentation
type SoftLayer_Container_Product_Order_Cart struct {

	// SavedCartKey - no documentation
	SavedCartKey string `json:"savedCartKey,omitempty"`

	// SavedCartName - no documentation
	SavedCartName string `json:"savedCartName,omitempty"`

	// CurrencyShortName - <nil>
	CurrencyShortName string `json:"currencyShortName,omitempty"`

	// Items - no documentation
	Items []*SoftLayer_Container_Product_Order_Cart_Item `json:"items,omitempty"`

	// Receipt - The receipt generated from ordering the items in this cart
	Receipt *SoftLayer_Container_Product_Order_Receipt `json:"receipt,omitempty"`

	// SavedCartId - no documentation
	SavedCartId int `json:"savedCartId,omitempty"`
}

func (softlayer_container_product_order_cart *SoftLayer_Container_Product_Order_Cart) String() string {
	return "SoftLayer_Container_Product_Order_Cart"
}
