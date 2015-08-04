package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Tag_Reference_Virtual_Guest - <nil>
type SoftLayer_Tag_Reference_Virtual_Guest struct {

	// UsrRecordId - <nil>
	UsrRecordId int `json:"usrRecordId,omitempty"`

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
}

func (softlayer_tag_reference_virtual_guest *SoftLayer_Tag_Reference_Virtual_Guest) String() string {
	return "SoftLayer_Tag_Reference_Virtual_Guest"
}

// SoftLayer_Tag_Reference_Virtual_Guest_Extended is SoftLayer_Tag_Reference_Virtual_Guest with all maskable types.
type SoftLayer_Tag_Reference_Virtual_Guest_Extended struct {
	SoftLayer_Tag_Reference_Virtual_Guest

	// TagType - <nil>
	TagType *SoftLayer_Tag_Type `json:"tagType,omitempty"`

	// Customer - <nil>
	Customer *SoftLayer_User_Customer `json:"customer,omitempty"`

	// Resource - <nil>
	Resource *SoftLayer_Virtual_Guest `json:"resource,omitempty"`

	// Employee - <nil>
	Employee *SoftLayer_User_Employee `json:"employee,omitempty"`

	// Tag - <nil>
	Tag *SoftLayer_Tag `json:"tag,omitempty"`
}

func (softlayer_tag_reference_virtual_guest *SoftLayer_Tag_Reference_Virtual_Guest_Extended) String() string {
	return "SoftLayer_Tag_Reference_Virtual_Guest"
}
