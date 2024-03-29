package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// McAfee_Epolicy_Orchestrator_Version45_Event_Version7 - The
// McAfee_Epolicy_Orchestrator_Version45_Event_Version7 data type represents a single event. It
// contains details about the event such as the date the event occurred, the virus or intrusion that is
// detected and the action that is taken.
type McAfee_Epolicy_Orchestrator_Version45_Event_Version7 struct {

	// SourceIpv4 - The IP address of the source that generated an event.
	SourceIpv4 string `json:"sourceIpv4,omitempty"`

	// TargetFilename - The name of the file that was the target of the event.
	TargetFilename string `json:"targetFilename,omitempty"`

	// ThreatName - no documentation
	ThreatName string `json:"threatName,omitempty"`

	// ThreatType - no documentation
	ThreatType string `json:"threatType,omitempty"`

	// DetectedUtc - no documentation
	DetectedUtc *time.Time `json:"detectedUtc,omitempty"`

	// SourceProcessName - no documentation
	SourceProcessName string `json:"sourceProcessName,omitempty"`

	// ThreatActionTaken - no documentation
	ThreatActionTaken string `json:"threatActionTaken,omitempty"`

	// ThreatSeverityLabel - no documentation
	ThreatSeverityLabel string `json:"threatSeverityLabel,omitempty"`

	// AgentDetails - no documentation
	AgentDetails *McAfee_Epolicy_Orchestrator_Version45_Agent_Details `json:"agentDetails,omitempty"`

	// VirusActionTaken - no documentation
	VirusActionTaken *McAfee_Epolicy_Orchestrator_Version45_Event_Filter_Description `json:"virusActionTaken,omitempty"`

	// Signature - no documentation
	Signature *McAfee_Epolicy_Orchestrator_Version45_Hips_Event_Signature_Version7 `json:"signature,omitempty"`
}

func (mcafee_epolicy_orchestrator_version45_event_version7 *McAfee_Epolicy_Orchestrator_Version45_Event_Version7) String() string {
	return "McAfee_Epolicy_Orchestrator_Version45_Event_Version7"
}
