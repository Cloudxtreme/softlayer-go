package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_IPSEvent - The
// McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_IPSEvent data type represents a single IPS
// event. It contains details about the event such as the date the event occurred, the process that
// generated it, the severity of the event, and the action taken.
type McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_IPSEvent struct {

	// SeverityText - no documentation
	SeverityText string `json:"severityText"`

	// IncidentTime - no documentation
	IncidentTime *time.Time `json:"incidentTime"`

	// ProcessName - no documentation
	ProcessName string `json:"processName"`

	// ReactionText - no documentation
	ReactionText string `json:"reactionText"`

	// RemoteIpAddress - no documentation
	RemoteIpAddress string `json:"remoteIpAddress"`
}

func (mcafee_epolicy_orchestrator_version36_hips_version7_ipsevent *McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_IPSEvent) String() string {
	return "McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_IPSEvent"
}

// McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_IPSEvent_Extended is McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_IPSEvent with all maskable types.
type McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_IPSEvent_Extended struct {
	McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_IPSEvent

	// Signature - no documentation
	Signature *McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_Event_Signature `json:"signature"`
}

func (mcafee_epolicy_orchestrator_version36_hips_version7_ipsevent *McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_IPSEvent_Extended) String() string {
	return "McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_IPSEvent"
}
