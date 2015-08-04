package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Resource_Group_Template_Member - <nil>
type SoftLayer_Resource_Group_Template_Member struct {

	// MinQuantity - <nil>
	MinQuantity int `json:"minQuantity,omitempty"`

	// RoleId - <nil>
	RoleId int `json:"roleId,omitempty"`

	// TemplateId - <nil>
	TemplateId int `json:"templateId,omitempty"`

	// MaxQuantity - <nil>
	MaxQuantity int `json:"maxQuantity,omitempty"`
}

func (softlayer_resource_group_template_member *SoftLayer_Resource_Group_Template_Member) String() string {
	return "SoftLayer_Resource_Group_Template_Member"
}

// SoftLayer_Resource_Group_Template_Member_Extended is SoftLayer_Resource_Group_Template_Member with all maskable types.
type SoftLayer_Resource_Group_Template_Member_Extended struct {
	SoftLayer_Resource_Group_Template_Member

	// Role - <nil>
	Role *SoftLayer_Resource_Group_Role `json:"role,omitempty"`

	// Template - <nil>
	Template *SoftLayer_Resource_Group_Template `json:"template,omitempty"`
}

func (softlayer_resource_group_template_member *SoftLayer_Resource_Group_Template_Member_Extended) String() string {
	return "SoftLayer_Resource_Group_Template_Member"
}
