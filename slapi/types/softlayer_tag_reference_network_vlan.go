package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Tag_Reference_Network_Vlan - <nil>
type SoftLayer_Tag_Reference_Network_Vlan struct {

	// EmpRecordId - <nil>
	EmpRecordId int `json:"empRecordId,omitempty"`

	// ResourceTableId - <nil>
	ResourceTableId int `json:"resourceTableId,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// TagId - <nil>
	TagId int `json:"tagId,omitempty"`

	// UsrRecordId - <nil>
	UsrRecordId int `json:"usrRecordId,omitempty"`

	// TagTypeId - <nil>
	TagTypeId int `json:"tagTypeId,omitempty"`
}

func (softlayer_tag_reference_network_vlan *SoftLayer_Tag_Reference_Network_Vlan) String() string {
	return "SoftLayer_Tag_Reference_Network_Vlan"
}

// SoftLayer_Tag_Reference_Network_Vlan_Extended is SoftLayer_Tag_Reference_Network_Vlan with all maskable types.
type SoftLayer_Tag_Reference_Network_Vlan_Extended struct {
	SoftLayer_Tag_Reference_Network_Vlan

	// Tag - <nil>
	Tag *SoftLayer_Tag `json:"tag,omitempty"`

	// Customer - <nil>
	Customer *SoftLayer_User_Customer `json:"customer,omitempty"`

	// TagType - <nil>
	TagType *SoftLayer_Tag_Type `json:"tagType,omitempty"`

	// Resource - <nil>
	Resource *SoftLayer_Network_Vlan `json:"resource,omitempty"`

	// Employee - <nil>
	Employee *SoftLayer_User_Employee `json:"employee,omitempty"`
}

func (softlayer_tag_reference_network_vlan *SoftLayer_Tag_Reference_Network_Vlan_Extended) String() string {
	return "SoftLayer_Tag_Reference_Network_Vlan"
}
