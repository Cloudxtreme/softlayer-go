package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Tag_Reference_Resource_Group - <nil>
type SoftLayer_Tag_Reference_Resource_Group struct {
}

// SoftLayer_Tag_Reference_Resource_Group_Extended is SoftLayer_Tag_Reference_Resource_Group with all maskable types.
type SoftLayer_Tag_Reference_Resource_Group_Extended struct {
	SoftLayer_Tag_Reference_Resource_Group

	// Resource - <nil>
	Resource *SoftLayer_Resource_Group `json:"resource"`
}

func (softlayer_tag_reference_resource_group *SoftLayer_Tag_Reference_Resource_Group) String() string {
	return "SoftLayer_Tag_Reference_Resource_Group"
}
