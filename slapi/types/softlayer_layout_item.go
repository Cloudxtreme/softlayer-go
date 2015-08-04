package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Layout_Item - The SoftLayer_Layout_Item contains definitions for default layout items
type SoftLayer_Layout_Item struct {

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// Keyname - The unique key name of the layout item, used primarily for programmatic purposes
	Keyname string `json:"keyname,omitempty"`

	// LayoutItemTypeId - The internal identifier of the related [[SoftLayer_Layout_Item_Type]]
	LayoutItemTypeId int `json:"layoutItemTypeId,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// LayoutItemPreferenceCount - A count of the layout preferences assigned to this layout item
	LayoutItemPreferenceCount uint64 `json:"layoutItemPreferenceCount,omitempty"`

	// LayoutItemPreferences - The layout preferences assigned to this layout item
	LayoutItemPreferences []*SoftLayer_Layout_Preference `json:"layoutItemPreferences,omitempty"`

	// LayoutItemType - no documentation
	LayoutItemType *SoftLayer_Layout_Item_Type `json:"layoutItemType,omitempty"`
}

func (softlayer_layout_item *SoftLayer_Layout_Item) String() string {
	return "SoftLayer_Layout_Item"
}
