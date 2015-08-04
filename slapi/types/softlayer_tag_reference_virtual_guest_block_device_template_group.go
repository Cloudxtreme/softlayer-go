package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Tag_Reference_Virtual_Guest_Block_Device_Template_Group - <nil>
type SoftLayer_Tag_Reference_Virtual_Guest_Block_Device_Template_Group struct {

	// EmpRecordId - <nil>
	EmpRecordId int `json:"empRecordId,omitempty"`

	// TagTypeId - <nil>
	TagTypeId int `json:"tagTypeId,omitempty"`

	// TagId - <nil>
	TagId int `json:"tagId,omitempty"`

	// ResourceTableId - <nil>
	ResourceTableId int `json:"resourceTableId,omitempty"`

	// UsrRecordId - <nil>
	UsrRecordId int `json:"usrRecordId,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`
}

func (softlayer_tag_reference_virtual_guest_block_device_template_group *SoftLayer_Tag_Reference_Virtual_Guest_Block_Device_Template_Group) String() string {
	return "SoftLayer_Tag_Reference_Virtual_Guest_Block_Device_Template_Group"
}

// SoftLayer_Tag_Reference_Virtual_Guest_Block_Device_Template_Group_Extended is SoftLayer_Tag_Reference_Virtual_Guest_Block_Device_Template_Group with all maskable types.
type SoftLayer_Tag_Reference_Virtual_Guest_Block_Device_Template_Group_Extended struct {
	SoftLayer_Tag_Reference_Virtual_Guest_Block_Device_Template_Group

	// Customer - <nil>
	Customer *SoftLayer_User_Customer `json:"customer,omitempty"`

	// Tag - <nil>
	Tag *SoftLayer_Tag `json:"tag,omitempty"`

	// Employee - <nil>
	Employee *SoftLayer_User_Employee `json:"employee,omitempty"`

	// TagType - <nil>
	TagType *SoftLayer_Tag_Type `json:"tagType,omitempty"`

	// Resource - <nil>
	Resource *SoftLayer_Virtual_Guest_Block_Device_Template_Group `json:"resource,omitempty"`
}

func (softlayer_tag_reference_virtual_guest_block_device_template_group *SoftLayer_Tag_Reference_Virtual_Guest_Block_Device_Template_Group_Extended) String() string {
	return "SoftLayer_Tag_Reference_Virtual_Guest_Block_Device_Template_Group"
}
