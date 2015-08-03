package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Network_IntrusionProtection_Event - The IntrusionProtection_Event object stores
// information about individual intrusion protection events. It is a data container that cannot be
// edited, deleted, or saved. It is returned by many methods in the TippingPointReporting object, but
// never directly, always as a child of another container object.
type SoftLayer_Container_Network_IntrusionProtection_Event struct {

	// AttackName - no documentation
	AttackName string `json:"attackName"`

	// BeginTime - The starting timestamp of the attack recorded, in Y-m-d H:i:s format. May not be set,
	// depending on the type of query performed.
	BeginTime string `json:"beginTime"`

	// Classification - no documentation
	Classification string `json:"classification"`

	// Protocol - no documentation
	Protocol string `json:"protocol"`

	// ActionTaken - The action that was taken when this attack was discovered. Can be either "Block" or
	// "Permit"
	ActionTaken string `json:"actionTaken"`

	// AttackCount - The number of attacks in this block. Attacks are grouped differently based on the
	// query performed on the tippingPointReporting object.
	AttackCount int `json:"attackCount"`

	// SignatureId - Unique ID of the "Signature" in question. The signature determines the type of attack
	// recorded. SignatureId is used in the drillDown() function on the TippingPointReporting service
	SignatureId string `json:"signatureId"`

	// SourcePort - no documentation
	SourcePort int `json:"sourcePort"`

	// CVEId - The CVE ID(s), if any, associated with this attack signature.
	CVEId string `json:"CVEId"`

	// AttackLongDescription - Long description of the attack. May contain links to more information
	AttackLongDescription string `json:"attackLongDescription"`

	// Severity - The human-readable severity of this attack, from "Low" to "Critical"
	Severity string `json:"severity"`

	// SourceIpAddress - The IP Address (as a dotted decimal string) of the machine originating the attack
	SourceIpAddress string `json:"sourceIpAddress"`

	// Platform - no documentation
	Platform string `json:"platform"`

	// BugtraqId - The BugTraq ID(s), if any, associated with this attack signature.
	BugtraqId string `json:"bugtraqId"`

	// DestinationIpAddress - The IP Address (as a dotted decimal string) of the machine that was the
	// target of the attack
	DestinationIpAddress string `json:"destinationIpAddress"`

	// DestinationPort - no documentation
	DestinationPort int `json:"destinationPort"`

	// EndTime - The ending timestamp of the attack recorded, in Y-m-d H:i:s format. May not be set,
	// depending on the type of query performed.
	EndTime string `json:"endTime"`
}

func (softlayer_container_network_intrusionprotection_event *SoftLayer_Container_Network_IntrusionProtection_Event) String() string {
	return "SoftLayer_Container_Network_IntrusionProtection_Event"
}
