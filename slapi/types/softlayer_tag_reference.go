package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Tag_Reference - <nil>
type SoftLayer_Tag_Reference struct {

	// EmpRecordId - <nil>
	EmpRecordId int `json:"empRecordId"`

	// TagId - <nil>
	TagId int `json:"tagId"`

	// TagTypeId - <nil>
	TagTypeId int `json:"tagTypeId"`

	// UsrRecordId - <nil>
	UsrRecordId int `json:"usrRecordId"`

	// Id - <nil>
	Id int `json:"id"`

	// ResourceTableId - <nil>
	ResourceTableId int `json:"resourceTableId"`
}

// SoftLayer_Tag_Reference_Extended is SoftLayer_Tag_Reference with all maskable types.
type SoftLayer_Tag_Reference_Extended struct {
	SoftLayer_Tag_Reference

	// Tag - <nil>
	Tag *SoftLayer_Tag `json:"tag"`

	// TagType - <nil>
	TagType *SoftLayer_Tag_Type `json:"tagType"`

	// Employee - <nil>
	Employee *SoftLayer_User_Employee `json:"employee"`

	// Customer - <nil>
	Customer *SoftLayer_User_Customer `json:"customer"`
}

func (softlayer_tag_reference *SoftLayer_Tag_Reference) String() string {
	return "SoftLayer_Tag_Reference"
}
