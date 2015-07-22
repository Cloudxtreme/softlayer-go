package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Order_Cart - no documentation
type SoftLayer_Container_Product_Order_Cart struct {

	// CurrencyShortName - <nil>
	CurrencyShortName string `json:"currencyShortName"`

	// Items - no documentation
	Items []*SoftLayer_Container_Product_Order_Cart_Item `json:"items"`

	// Receipt - The receipt generated from ordering the items in this cart
	Receipt *SoftLayer_Container_Product_Order_Receipt `json:"receipt"`

	// SavedCartId - no documentation
	SavedCartId int `json:"savedCartId"`

	// SavedCartKey - no documentation
	SavedCartKey string `json:"savedCartKey"`

	// SavedCartName - no documentation
	SavedCartName string `json:"savedCartName"`
}
