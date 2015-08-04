package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Tag_Reference_Resource_Group - <nil>
type SoftLayer_Tag_Reference_Resource_Group struct {

	// ResourceTableId - <nil>
	ResourceTableId int `json:"resourceTableId,omitempty"`

	// EmpRecordId - <nil>
	EmpRecordId int `json:"empRecordId,omitempty"`

	// UsrRecordId - <nil>
	UsrRecordId int `json:"usrRecordId,omitempty"`

	// TagTypeId - <nil>
	TagTypeId int `json:"tagTypeId,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// TagId - <nil>
	TagId int `json:"tagId,omitempty"`
}

func (softlayer_tag_reference_resource_group *SoftLayer_Tag_Reference_Resource_Group) String() string {
	return "SoftLayer_Tag_Reference_Resource_Group"
}

// SoftLayer_Tag_Reference_Resource_Group_Extended is SoftLayer_Tag_Reference_Resource_Group with all maskable types.
type SoftLayer_Tag_Reference_Resource_Group_Extended struct {
	SoftLayer_Tag_Reference_Resource_Group

	// Employee - <nil>
	Employee *SoftLayer_User_Employee `json:"employee,omitempty"`

	// Resource - <nil>
	Resource *SoftLayer_Resource_Group `json:"resource,omitempty"`

	// Tag - <nil>
	Tag *SoftLayer_Tag `json:"tag,omitempty"`

	// TagType - <nil>
	TagType *SoftLayer_Tag_Type `json:"tagType,omitempty"`

	// Customer - <nil>
	Customer *SoftLayer_User_Customer `json:"customer,omitempty"`
}

func (softlayer_tag_reference_resource_group *SoftLayer_Tag_Reference_Resource_Group_Extended) String() string {
	return "SoftLayer_Tag_Reference_Resource_Group"
}
