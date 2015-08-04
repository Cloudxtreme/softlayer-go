package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Billing_Item_Link_ThePlanet - <nil>
type SoftLayer_Billing_Item_Link_ThePlanet struct {
}

func (softlayer_billing_item_link_theplanet *SoftLayer_Billing_Item_Link_ThePlanet) String() string {
	return "SoftLayer_Billing_Item_Link_ThePlanet"
}

// SoftLayer_Billing_Item_Link_ThePlanet_Extended is SoftLayer_Billing_Item_Link_ThePlanet with all maskable types.
type SoftLayer_Billing_Item_Link_ThePlanet_Extended struct {
	SoftLayer_Billing_Item_Link_ThePlanet

	// BillingItem - <nil>
	BillingItem *SoftLayer_Billing_Item `json:"billingItem,omitempty"`

	// ServiceProvider - <nil>
	ServiceProvider *SoftLayer_Service_Provider `json:"serviceProvider,omitempty"`
}

func (softlayer_billing_item_link_theplanet *SoftLayer_Billing_Item_Link_ThePlanet_Extended) String() string {
	return "SoftLayer_Billing_Item_Link_ThePlanet"
}
