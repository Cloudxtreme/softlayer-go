package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Item_Price_Account_Restriction - The SoftLayer_Product_Item_Price data type gives
// more information about the item price restrictions. An item price may be restricted to one or more
// accounts. If the item price is restricted to an account, only that account will see the restriction
// details.
type SoftLayer_Product_Item_Price_Account_Restriction struct {

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// AccountId - The account id for the item price account restriction.
	AccountId int `json:"accountId"`

	// Id - The unique identifier for the item price account restriction.
	Id int `json:"id"`

	// ItemPrice - no documentation
	ItemPrice *SoftLayer_Product_Item_Price `json:"itemPrice"`

	// ItemPriceId - The item price id for the item price account restriction.
	ItemPriceId int `json:"itemPriceId"`
}

func (softlayer_product_item_price_account_restriction *SoftLayer_Product_Item_Price_Account_Restriction) String() string {
	return "SoftLayer_Product_Item_Price_Account_Restriction"
}
