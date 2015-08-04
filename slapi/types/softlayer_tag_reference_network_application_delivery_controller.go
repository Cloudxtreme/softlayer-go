package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Tag_Reference_Network_Application_Delivery_Controller - <nil>
type SoftLayer_Tag_Reference_Network_Application_Delivery_Controller struct {

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// TagTypeId - <nil>
	TagTypeId int `json:"tagTypeId,omitempty"`

	// ResourceTableId - <nil>
	ResourceTableId int `json:"resourceTableId,omitempty"`

	// EmpRecordId - <nil>
	EmpRecordId int `json:"empRecordId,omitempty"`

	// TagId - <nil>
	TagId int `json:"tagId,omitempty"`

	// UsrRecordId - <nil>
	UsrRecordId int `json:"usrRecordId,omitempty"`
}

func (softlayer_tag_reference_network_application_delivery_controller *SoftLayer_Tag_Reference_Network_Application_Delivery_Controller) String() string {
	return "SoftLayer_Tag_Reference_Network_Application_Delivery_Controller"
}

// SoftLayer_Tag_Reference_Network_Application_Delivery_Controller_Extended is SoftLayer_Tag_Reference_Network_Application_Delivery_Controller with all maskable types.
type SoftLayer_Tag_Reference_Network_Application_Delivery_Controller_Extended struct {
	SoftLayer_Tag_Reference_Network_Application_Delivery_Controller

	// Tag - <nil>
	Tag *SoftLayer_Tag `json:"tag,omitempty"`

	// Employee - <nil>
	Employee *SoftLayer_User_Employee `json:"employee,omitempty"`

	// TagType - <nil>
	TagType *SoftLayer_Tag_Type `json:"tagType,omitempty"`

	// Resource - <nil>
	Resource *SoftLayer_Network_Application_Delivery_Controller `json:"resource,omitempty"`

	// Customer - <nil>
	Customer *SoftLayer_User_Customer `json:"customer,omitempty"`
}

func (softlayer_tag_reference_network_application_delivery_controller *SoftLayer_Tag_Reference_Network_Application_Delivery_Controller_Extended) String() string {
	return "SoftLayer_Tag_Reference_Network_Application_Delivery_Controller"
}
