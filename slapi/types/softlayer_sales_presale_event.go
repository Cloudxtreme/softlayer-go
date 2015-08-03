package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Sales_Presale_Event - The presale event data types indicate the information regarding an
// individual presale event. The '''locationId''' will indicate the datacenter associated with the
// presale event. The '''itemId''' will indicate the product item associated with a particular presale
// event - however these are more rare. The '''startDate''' and '''endDate''' will provide information
// regarding when the presale event is available for use. At the end of the presale event, the server
// or services purchased will be available once approved and provisioned.
type SoftLayer_Sales_Presale_Event struct {

	// ItemId - [[SoftLayer_Product_Item]] id associated with the presale event.
	ItemId int `json:"itemId"`

	// LocationId - no documentation
	LocationId int `json:"locationId"`

	// StartDate - Start date of the presale event. Orders cannot be approved before this date.
	StartDate *time.Time `json:"startDate"`

	// Description - no documentation
	Description string `json:"description"`

	// Id - no documentation
	Id int `json:"id"`

	// EndDate - End date of the presale event. Orders can be approved and provisioned after this date.
	EndDate *time.Time `json:"endDate"`
}

// SoftLayer_Sales_Presale_Event_Extended is SoftLayer_Sales_Presale_Event with all maskable types.
type SoftLayer_Sales_Presale_Event_Extended struct {
	SoftLayer_Sales_Presale_Event

	// ExpiredFlag - A flag to indicate that the presale event is expired. A presale event is expired if
	// the current time is after the end date.
	ExpiredFlag bool `json:"expiredFlag"`

	// Item - The [[SoftLayer_Product_Item]] associated with the presale event.
	Item *SoftLayer_Product_Item `json:"item"`

	// Location - The [[SoftLayer_Location]] associated with the presale event.
	Location *SoftLayer_Location `json:"location"`

	// Orders - The orders ([[SoftLayer_Billing_Order]]) associated with this presale event that were
	// created for the customer's account.
	Orders []*SoftLayer_Billing_Order `json:"orders"`

	// OrderCount - A count of the orders ([[SoftLayer_Billing_Order]]) associated with this presale event
	// that were created for the customer's account.
	OrderCount uint64 `json:"orderCount"`

	// ActiveFlag - A flag to indicate that the presale event is currently active. A presale event is
	// active if the current time is between the start and end dates.
	ActiveFlag bool `json:"activeFlag"`
}

func (softlayer_sales_presale_event *SoftLayer_Sales_Presale_Event) String() string {
	return "SoftLayer_Sales_Presale_Event"
}
