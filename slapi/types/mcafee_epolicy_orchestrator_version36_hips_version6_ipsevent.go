package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_IPSEvent - The
// McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_IPSEvent data type represents a single IPS
// event. It contains details about the event such as the date the event occurred, the process that
// generated it, the severity of the event, and the action taken.
type McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_IPSEvent struct {

	// IncidentTime - no documentation
	IncidentTime *time.Time `json:"incidentTime"`

	// ProcessName - no documentation
	ProcessName string `json:"processName"`

	// ReactionText - no documentation
	ReactionText string `json:"reactionText"`

	// RemoteIpAddress - no documentation
	RemoteIpAddress string `json:"remoteIpAddress"`

	// SeverityText - no documentation
	SeverityText string `json:"severityText"`

	// Signature - no documentation
	Signature *McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_Event_Signature `json:"signature"`
}

func (mcafee_epolicy_orchestrator_version36_hips_version6_ipsevent *McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_IPSEvent) String() string {
	return "McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_IPSEvent"
}
