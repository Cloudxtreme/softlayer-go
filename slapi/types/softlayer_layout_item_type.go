package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Layout_Item_Type - The SoftLayer_Layout_Item_Type contains definitions for item types
type SoftLayer_Layout_Item_Type struct {

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// Keyname - The unique key name of the item type, used primarily for programmatic purposes
	Keyname string `json:"keyname,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`
}

func (softlayer_layout_item_type *SoftLayer_Layout_Item_Type) String() string {
	return "SoftLayer_Layout_Item_Type"
}
