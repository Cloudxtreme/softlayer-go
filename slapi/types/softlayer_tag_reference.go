package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Tag_Reference - <nil>
type SoftLayer_Tag_Reference struct {

	// TagTypeId - <nil>
	TagTypeId int `json:"tagTypeId"`

	// ResourceTableId - <nil>
	ResourceTableId int `json:"resourceTableId"`

	// TagId - <nil>
	TagId int `json:"tagId"`

	// UsrRecordId - <nil>
	UsrRecordId int `json:"usrRecordId"`

	// EmpRecordId - <nil>
	EmpRecordId int `json:"empRecordId"`

	// Id - <nil>
	Id int `json:"id"`
}

func (softlayer_tag_reference *SoftLayer_Tag_Reference) String() string {
	return "SoftLayer_Tag_Reference"
}

// SoftLayer_Tag_Reference_Extended is SoftLayer_Tag_Reference with all maskable types.
type SoftLayer_Tag_Reference_Extended struct {
	SoftLayer_Tag_Reference

	// Customer - <nil>
	Customer *SoftLayer_User_Customer `json:"customer"`

	// Employee - <nil>
	Employee *SoftLayer_User_Employee `json:"employee"`

	// Tag - <nil>
	Tag *SoftLayer_Tag `json:"tag"`

	// TagType - <nil>
	TagType *SoftLayer_Tag_Type `json:"tagType"`
}

func (softlayer_tag_reference *SoftLayer_Tag_Reference_Extended) String() string {
	return "SoftLayer_Tag_Reference"
}
