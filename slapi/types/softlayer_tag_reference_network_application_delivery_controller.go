package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Tag_Reference_Network_Application_Delivery_Controller - <nil>
type SoftLayer_Tag_Reference_Network_Application_Delivery_Controller struct {

	// EmpRecordId - <nil>
	EmpRecordId int `json:"empRecordId,omitempty"`

	// ResourceTableId - <nil>
	ResourceTableId int `json:"resourceTableId,omitempty"`

	// TagId - <nil>
	TagId int `json:"tagId,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// TagTypeId - <nil>
	TagTypeId int `json:"tagTypeId,omitempty"`

	// UsrRecordId - <nil>
	UsrRecordId int `json:"usrRecordId,omitempty"`

	// Resource - <nil>
	Resource *SoftLayer_Network_Application_Delivery_Controller `json:"resource,omitempty"`

	// Employee - <nil>
	Employee *SoftLayer_User_Employee `json:"employee,omitempty"`

	// Tag - <nil>
	Tag *SoftLayer_Tag `json:"tag,omitempty"`

	// TagType - <nil>
	TagType *SoftLayer_Tag_Type `json:"tagType,omitempty"`

	// Customer - <nil>
	Customer *SoftLayer_User_Customer `json:"customer,omitempty"`
}

func (softlayer_tag_reference_network_application_delivery_controller *SoftLayer_Tag_Reference_Network_Application_Delivery_Controller) String() string {
	return "SoftLayer_Tag_Reference_Network_Application_Delivery_Controller"
}
