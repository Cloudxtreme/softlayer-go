package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Resource_Group_Descendant_Reference - The SoftLayer_Resource_Group_Descendant_Reference
// data type simplifies the link between one SoftLayer_Resource_Group_Member object and all of its
// parents.
type SoftLayer_Resource_Group_Descendant_Reference struct {

	// Group - <nil>
	Group *SoftLayer_Resource_Group `json:"group"`

	// GroupMember - <nil>
	GroupMember *SoftLayer_Resource_Group_Member `json:"groupMember"`
}

func (softlayer_resource_group_descendant_reference *SoftLayer_Resource_Group_Descendant_Reference) String() string {
	return "SoftLayer_Resource_Group_Descendant_Reference"
}
