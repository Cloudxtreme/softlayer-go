package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Layout_Container - The SoftLayer_Layout_Container contains definitions for default page
// layouts
type SoftLayer_Layout_Container struct {

	// Keyname - The unique key name of the layout container, used primarily for programmatic purposes
	Keyname string `json:"keyname"`

	// LayoutContainerTypeId - The internal identifier of the related [[SoftLayer_Layout_Container_Type]]
	LayoutContainerTypeId int `json:"layoutContainerTypeId"`

	// Name - no documentation
	Name string `json:"name"`

	// Id - no documentation
	Id int `json:"id"`
}

// SoftLayer_Layout_Container_Extended is SoftLayer_Layout_Container with all maskable types.
type SoftLayer_Layout_Container_Extended struct {
	SoftLayer_Layout_Container

	// LayoutItemCount - A count of the layout items assigned to this layout container
	LayoutItemCount uint64 `json:"layoutItemCount"`

	// LayoutContainerType - no documentation
	LayoutContainerType *SoftLayer_Layout_Container_Type `json:"layoutContainerType"`

	// LayoutItems - no documentation
	LayoutItems []*SoftLayer_Layout_Item `json:"layoutItems"`
}

func (softlayer_layout_container *SoftLayer_Layout_Container) String() string {
	return "SoftLayer_Layout_Container"
}
