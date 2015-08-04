package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Item_Link_ThePlanet - <nil>
type SoftLayer_Product_Item_Link_ThePlanet struct {
}

func (softlayer_product_item_link_theplanet *SoftLayer_Product_Item_Link_ThePlanet) String() string {
	return "SoftLayer_Product_Item_Link_ThePlanet"
}

// SoftLayer_Product_Item_Link_ThePlanet_Extended is SoftLayer_Product_Item_Link_ThePlanet with all maskable types.
type SoftLayer_Product_Item_Link_ThePlanet_Extended struct {
	SoftLayer_Product_Item_Link_ThePlanet

	// Item - <nil>
	Item *SoftLayer_Product_Item `json:"item,omitempty"`

	// ServiceProvider - <nil>
	ServiceProvider *SoftLayer_Service_Provider `json:"serviceProvider,omitempty"`
}

func (softlayer_product_item_link_theplanet *SoftLayer_Product_Item_Link_ThePlanet_Extended) String() string {
	return "SoftLayer_Product_Item_Link_ThePlanet"
}
