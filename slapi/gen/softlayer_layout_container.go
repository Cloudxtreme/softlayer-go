package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Layout_Container - The SoftLayer_Layout_Container contains definitions for default page
// layouts
type SoftLayer_Layout_Container struct {

	// Id - no documentation
	Id int `json:"id"`

	// Keyname - The unique key name of the layout container, used primarily for programmatic purposes
	Keyname string `json:"keyname"`

	// LayoutContainerType - no documentation
	LayoutContainerType *SoftLayer_Layout_Container_Type `json:"layoutContainerType"`

	// LayoutContainerTypeId - The internal identifier of the related [[SoftLayer_Layout_Container_Type]]
	LayoutContainerTypeId int `json:"layoutContainerTypeId"`

	// LayoutItemCount - A count of the layout items assigned to this layout container
	LayoutItemCount uint64 `json:"layoutItemCount"`

	// LayoutItems - no documentation
	LayoutItems []*SoftLayer_Layout_Item `json:"layoutItems"`

	// Name - no documentation
	Name string `json:"name"`
}

// GetAllObjects - Use this method to retrieve all active layout containers that can be customized.
func (softlayer_layout_container *SoftLayer_Layout_Container) GetAllObjects() ([]*SoftLayer_Layout_Container, error) {
	var returnValue []*SoftLayer_Layout_Container
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_layout_container *SoftLayer_Layout_Container) GetObject() (*SoftLayer_Layout_Container, error) {
	var returnValue *SoftLayer_Layout_Container
	return returnValue, nil
}
