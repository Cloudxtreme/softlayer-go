package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Tag_Reference_Network_Vlan_Firewall - <nil>
type SoftLayer_Tag_Reference_Network_Vlan_Firewall struct {

	// TagId - <nil>
	TagId int `json:"tagId,omitempty"`

	// TagTypeId - <nil>
	TagTypeId int `json:"tagTypeId,omitempty"`

	// EmpRecordId - <nil>
	EmpRecordId int `json:"empRecordId,omitempty"`

	// ResourceTableId - <nil>
	ResourceTableId int `json:"resourceTableId,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// UsrRecordId - <nil>
	UsrRecordId int `json:"usrRecordId,omitempty"`
}

func (softlayer_tag_reference_network_vlan_firewall *SoftLayer_Tag_Reference_Network_Vlan_Firewall) String() string {
	return "SoftLayer_Tag_Reference_Network_Vlan_Firewall"
}

// SoftLayer_Tag_Reference_Network_Vlan_Firewall_Extended is SoftLayer_Tag_Reference_Network_Vlan_Firewall with all maskable types.
type SoftLayer_Tag_Reference_Network_Vlan_Firewall_Extended struct {
	SoftLayer_Tag_Reference_Network_Vlan_Firewall

	// Employee - <nil>
	Employee *SoftLayer_User_Employee `json:"employee,omitempty"`

	// TagType - <nil>
	TagType *SoftLayer_Tag_Type `json:"tagType,omitempty"`

	// Resource - <nil>
	Resource *SoftLayer_Network_Vlan_Firewall `json:"resource,omitempty"`

	// Customer - <nil>
	Customer *SoftLayer_User_Customer `json:"customer,omitempty"`

	// Tag - <nil>
	Tag *SoftLayer_Tag `json:"tag,omitempty"`
}

func (softlayer_tag_reference_network_vlan_firewall *SoftLayer_Tag_Reference_Network_Vlan_Firewall_Extended) String() string {
	return "SoftLayer_Tag_Reference_Network_Vlan_Firewall"
}
