package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Tag_Reference_Virtual_Guest_Block_Device_Template_Group - <nil>
type SoftLayer_Tag_Reference_Virtual_Guest_Block_Device_Template_Group struct {

	// EmpRecordId - <nil>
	EmpRecordId int `json:"empRecordId,omitempty"`

	// UsrRecordId - <nil>
	UsrRecordId int `json:"usrRecordId,omitempty"`

	// TagTypeId - <nil>
	TagTypeId int `json:"tagTypeId,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// ResourceTableId - <nil>
	ResourceTableId int `json:"resourceTableId,omitempty"`

	// TagId - <nil>
	TagId int `json:"tagId,omitempty"`

	// Resource - <nil>
	Resource *SoftLayer_Virtual_Guest_Block_Device_Template_Group `json:"resource,omitempty"`

	// Tag - <nil>
	Tag *SoftLayer_Tag `json:"tag,omitempty"`

	// Customer - <nil>
	Customer *SoftLayer_User_Customer `json:"customer,omitempty"`

	// Employee - <nil>
	Employee *SoftLayer_User_Employee `json:"employee,omitempty"`

	// TagType - <nil>
	TagType *SoftLayer_Tag_Type `json:"tagType,omitempty"`
}

func (softlayer_tag_reference_virtual_guest_block_device_template_group *SoftLayer_Tag_Reference_Virtual_Guest_Block_Device_Template_Group) String() string {
	return "SoftLayer_Tag_Reference_Virtual_Guest_Block_Device_Template_Group"
}
