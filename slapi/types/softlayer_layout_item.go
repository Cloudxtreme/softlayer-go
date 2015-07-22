package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Layout_Item - The SoftLayer_Layout_Item contains definitions for default layout items
type SoftLayer_Layout_Item struct {

	// Id - no documentation
	Id int `json:"id"`

	// Keyname - The unique key name of the layout item, used primarily for programmatic purposes
	Keyname string `json:"keyname"`

	// LayoutItemPreferenceCount - A count of the layout preferences assigned to this layout item
	LayoutItemPreferenceCount uint64 `json:"layoutItemPreferenceCount"`

	// LayoutItemPreferences - The layout preferences assigned to this layout item
	LayoutItemPreferences []*SoftLayer_Layout_Preference `json:"layoutItemPreferences"`

	// LayoutItemType - no documentation
	LayoutItemType *SoftLayer_Layout_Item_Type `json:"layoutItemType"`

	// LayoutItemTypeId - The internal identifier of the related [[SoftLayer_Layout_Item_Type]]
	LayoutItemTypeId int `json:"layoutItemTypeId"`

	// Name - no documentation
	Name string `json:"name"`
}

// GetObject - <nil>
func (softlayer_layout_item *SoftLayer_Layout_Item) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Layout_Item, error) {
	var returnValue *SoftLayer_Layout_Item
	return returnValue, nil
}
