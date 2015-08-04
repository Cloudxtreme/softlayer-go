package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Layout_Container - The SoftLayer_Layout_Container contains definitions for default page
// layouts
type SoftLayer_Layout_Container struct {

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// Keyname - The unique key name of the layout container, used primarily for programmatic purposes
	Keyname string `json:"keyname,omitempty"`

	// LayoutContainerTypeId - The internal identifier of the related [[SoftLayer_Layout_Container_Type]]
	LayoutContainerTypeId int `json:"layoutContainerTypeId,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// LayoutItemCount - A count of the layout items assigned to this layout container
	LayoutItemCount uint64 `json:"layoutItemCount,omitempty"`

	// LayoutContainerType - no documentation
	LayoutContainerType *SoftLayer_Layout_Container_Type `json:"layoutContainerType,omitempty"`

	// LayoutItems - no documentation
	LayoutItems []*SoftLayer_Layout_Item `json:"layoutItems,omitempty"`
}

func (softlayer_layout_container *SoftLayer_Layout_Container) String() string {
	return "SoftLayer_Layout_Container"
}
