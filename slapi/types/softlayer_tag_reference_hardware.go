package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Tag_Reference_Hardware - <nil>
type SoftLayer_Tag_Reference_Hardware struct {

	// TagTypeId - <nil>
	TagTypeId int `json:"tagTypeId,omitempty"`

	// UsrRecordId - <nil>
	UsrRecordId int `json:"usrRecordId,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// TagId - <nil>
	TagId int `json:"tagId,omitempty"`

	// EmpRecordId - <nil>
	EmpRecordId int `json:"empRecordId,omitempty"`

	// ResourceTableId - <nil>
	ResourceTableId int `json:"resourceTableId,omitempty"`
}

func (softlayer_tag_reference_hardware *SoftLayer_Tag_Reference_Hardware) String() string {
	return "SoftLayer_Tag_Reference_Hardware"
}

// SoftLayer_Tag_Reference_Hardware_Extended is SoftLayer_Tag_Reference_Hardware with all maskable types.
type SoftLayer_Tag_Reference_Hardware_Extended struct {
	SoftLayer_Tag_Reference_Hardware

	// Tag - <nil>
	Tag *SoftLayer_Tag `json:"tag,omitempty"`

	// Customer - <nil>
	Customer *SoftLayer_User_Customer `json:"customer,omitempty"`

	// Employee - <nil>
	Employee *SoftLayer_User_Employee `json:"employee,omitempty"`

	// TagType - <nil>
	TagType *SoftLayer_Tag_Type `json:"tagType,omitempty"`

	// Resource - <nil>
	Resource *SoftLayer_Hardware `json:"resource,omitempty"`
}

func (softlayer_tag_reference_hardware *SoftLayer_Tag_Reference_Hardware_Extended) String() string {
	return "SoftLayer_Tag_Reference_Hardware"
}
