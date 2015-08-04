package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Tag_Reference - <nil>
type SoftLayer_Tag_Reference struct {

	// TagId - <nil>
	TagId int `json:"tagId,omitempty"`

	// TagTypeId - <nil>
	TagTypeId int `json:"tagTypeId,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// ResourceTableId - <nil>
	ResourceTableId int `json:"resourceTableId,omitempty"`

	// EmpRecordId - <nil>
	EmpRecordId int `json:"empRecordId,omitempty"`

	// UsrRecordId - <nil>
	UsrRecordId int `json:"usrRecordId,omitempty"`
}

func (softlayer_tag_reference *SoftLayer_Tag_Reference) String() string {
	return "SoftLayer_Tag_Reference"
}

// SoftLayer_Tag_Reference_Extended is SoftLayer_Tag_Reference with all maskable types.
type SoftLayer_Tag_Reference_Extended struct {
	SoftLayer_Tag_Reference

	// Tag - <nil>
	Tag *SoftLayer_Tag `json:"tag,omitempty"`

	// TagType - <nil>
	TagType *SoftLayer_Tag_Type `json:"tagType,omitempty"`

	// Customer - <nil>
	Customer *SoftLayer_User_Customer `json:"customer,omitempty"`

	// Employee - <nil>
	Employee *SoftLayer_User_Employee `json:"employee,omitempty"`
}

func (softlayer_tag_reference *SoftLayer_Tag_Reference_Extended) String() string {
	return "SoftLayer_Tag_Reference"
}
