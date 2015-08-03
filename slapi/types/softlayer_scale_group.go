package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Scale_Group - <nil>
type SoftLayer_Scale_Group struct {

	// TerminationPolicyId - The termination policy for the group. This determines which member to choose
	// to delete when scaling downwards.
	TerminationPolicyId int `json:"terminationPolicyId"`

	// AccountId - The identifier of the account assigned to this group.
	AccountId int `json:"accountId"`

	// Cooldown - The number of seconds this group will wait after lastActionDate before performing another
	// action. Be advised, this can be overridden per policy. While strongly discouraged, a value of 0
	// effectively disables cooldown.
	Cooldown int `json:"cooldown"`

	// MaximumMemberCount - The greatest number of virtual guest members that are allowed on this group.
	// Any attempts to add a guest member will fail if it will result in the total guest member count of
	// this group to be above this number. If this number is edited and is less than the current guest
	// member count, guests will be removed to at least be no greater than this number.
	MaximumMemberCount int `json:"maximumMemberCount"`

	// MinimumMemberCount - The fewest number of virtual guest members that are allowed on this group. Any
	// attempts to remove a guest member will fail if it will result in the total guest member count of
	// this group to be below this number. If this number is edited and is larger than the current guest
	// member count, guests will be added to at least reach this number.
	MinimumMemberCount int `json:"minimumMemberCount"`

	// Id - no documentation
	Id int `json:"id"`

	// SuspendedFlag - no documentation
	SuspendedFlag bool `json:"suspendedFlag"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// VirtualGuestMemberTemplate - This is the template to create guest members with. This is the same
	// template accepted by the createObject call on SoftLayer_Virtual_Guest with some caveats. The
	// hostname provided will have an arbitrary value appended to it for each guest created. Also,
	// hourlyBillingFlag cannot be false, and if the datacenter is provided it must be in the region of
	// this group. Finally, VLANs cannot be provided for the template, it will use VLANs provided to this
	// group instead. Note, if this template is edited on an existing group the previous template values
	// are not kept and are not considered during termination. This means a group's guest members could
	// effectively be a hybrid of multiple templates because this value was changed after some guest
	// members were created but before others were created.
	VirtualGuestMemberTemplate *SoftLayer_Virtual_Guest `json:"virtualGuestMemberTemplate"`

	// DesiredMemberCount - This value is only available on the template for creating and editing a group.
	// It will be null when retrieved. When this value is provided on create or edit, guests will be scaled
	// up or down to meet this number. This number must be in the range provided by minimumMemberCount and
	// maximumMemberCount. This value can only be present during create or edit when this group is active.
	// Note, guests that are created as a result of this value can possibly be removed after cooldown by a
	// policy.
	DesiredMemberCount int `json:"desiredMemberCount"`

	// Name - The name of this scale group. It must be unique on the account.
	Name string `json:"name"`

	// BalancedTerminationFlag - If this is true, this group will scale down members in a way to preserve
	// the balance across VLANs. If there is ambiguity about which member to use to maintain balance, the
	// terminationPolicy is used to resolve it. This is false by default and can only be set to true if
	// there are multiple VLANs that are being balanced across.
	BalancedTerminationFlag bool `json:"balancedTerminationFlag"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// RegionalGroupId - The identifier of the regional group this scaling group is assigned to.
	RegionalGroupId int `json:"regionalGroupId"`

	// LastActionDate - The date of the last action on this group or its create date
	LastActionDate *time.Time `json:"lastActionDate"`
}

func (softlayer_scale_group *SoftLayer_Scale_Group) String() string {
	return "SoftLayer_Scale_Group"
}

// SoftLayer_Scale_Group_Extended is SoftLayer_Scale_Group with all maskable types.
type SoftLayer_Scale_Group_Extended struct {
	SoftLayer_Scale_Group

	// Logs - no documentation
	Logs []*SoftLayer_Scale_Group_Log `json:"logs"`

	// VirtualGuestAssetCount - A count of collection of guests that have been pinned to this group. Guest
	// assets are only used for certain trigger checks such as resource watches. They do not count towards
	// the auto scaling guest counts of this group in anyway and are never automatically added or removed.
	VirtualGuestAssetCount uint64 `json:"virtualGuestAssetCount"`

	// VirtualGuestMemberCount - A count of collection of guests that have been scaled with the group. When
	// this group is active, the count of guests here is guaranteed to be between minimumMemberCount and
	// maximumMemberCount inclusively.
	VirtualGuestMemberCount uint64 `json:"virtualGuestMemberCount"`

	// RegionalGroup - no documentation
	RegionalGroup *SoftLayer_Location_Group_Regional `json:"regionalGroup"`

	// Status - no documentation
	Status *SoftLayer_Scale_Group_Status `json:"status"`

	// TerminationPolicy - no documentation
	TerminationPolicy *SoftLayer_Scale_Termination_Policy `json:"terminationPolicy"`

	// VirtualGuestMembers - Collection of guests that have been scaled with the group. When this group is
	// active, the count of guests here is guaranteed to be between minimumMemberCount and
	// maximumMemberCount inclusively.
	VirtualGuestMembers []*SoftLayer_Scale_Member_Virtual_Guest `json:"virtualGuestMembers"`

	// NetworkVlans - Collection of VLANs for this auto scale group. VLANs are optional. This can contain a
	// public or private or both. When a single for a public/private type is given it can be a
	// non-purchased only if the minimumMemberCount on the group is >= 1. This can also contain any number
	// of public/private purchased VLANs and members are staggered across them when scaled up.
	NetworkVlans []*SoftLayer_Scale_Network_Vlan `json:"networkVlans"`

	// Policies - Collection of policies for this group. This can be empty.
	Policies []*SoftLayer_Scale_Policy `json:"policies"`

	// NetworkVlanCount - A count of collection of VLANs for this auto scale group. VLANs are optional.
	// This can contain a public or private or both. When a single for a public/private type is given it
	// can be a non-purchased only if the minimumMemberCount on the group is >= 1. This can also contain
	// any number of public/private purchased VLANs and members are staggered across them when scaled up.
	NetworkVlanCount uint64 `json:"networkVlanCount"`

	// VirtualGuestAssets - Collection of guests that have been pinned to this group. Guest assets are only
	// used for certain trigger checks such as resource watches. They do not count towards the auto scaling
	// guest counts of this group in anyway and are never automatically added or removed.
	VirtualGuestAssets []*SoftLayer_Scale_Asset_Virtual_Guest `json:"virtualGuestAssets"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// LoadBalancers - Collection of load balancers for this auto scale group.
	LoadBalancers []*SoftLayer_Scale_LoadBalancer `json:"loadBalancers"`

	// PolicyCount - A count of collection of policies for this group. This can be empty.
	PolicyCount uint64 `json:"policyCount"`

	// LoadBalancerCount - A count of collection of load balancers for this auto scale group.
	LoadBalancerCount uint64 `json:"loadBalancerCount"`

	// LogCount - A count of collection of log entries for this group.
	LogCount uint64 `json:"logCount"`
}

func (softlayer_scale_group *SoftLayer_Scale_Group_Extended) String() string {
	return "SoftLayer_Scale_Group"
}
