package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Tag_Reference_Virtual_Guest - <nil>
type SoftLayer_Tag_Reference_Virtual_Guest struct {

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// TagTypeId - <nil>
	TagTypeId int `json:"tagTypeId,omitempty"`

	// EmpRecordId - <nil>
	EmpRecordId int `json:"empRecordId,omitempty"`

	// TagId - <nil>
	TagId int `json:"tagId,omitempty"`

	// UsrRecordId - <nil>
	UsrRecordId int `json:"usrRecordId,omitempty"`

	// ResourceTableId - <nil>
	ResourceTableId int `json:"resourceTableId,omitempty"`

	// Resource - <nil>
	Resource *SoftLayer_Virtual_Guest `json:"resource,omitempty"`

	// Customer - <nil>
	Customer *SoftLayer_User_Customer `json:"customer,omitempty"`

	// TagType - <nil>
	TagType *SoftLayer_Tag_Type `json:"tagType,omitempty"`

	// Employee - <nil>
	Employee *SoftLayer_User_Employee `json:"employee,omitempty"`

	// Tag - <nil>
	Tag *SoftLayer_Tag `json:"tag,omitempty"`
}

func (softlayer_tag_reference_virtual_guest *SoftLayer_Tag_Reference_Virtual_Guest) String() string {
	return "SoftLayer_Tag_Reference_Virtual_Guest"
}
