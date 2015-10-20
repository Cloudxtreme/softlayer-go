package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// McAfee_Epolicy_Orchestrator_Version45_Event_Version8 - The
// McAfee_Epolicy_Orchestrator_Version45_Event_Version8 data type represents a single event. It
// contains details about the event such as the date the event occurred, the virus or intrusion that is
// detected and the action that is taken.
type McAfee_Epolicy_Orchestrator_Version45_Event_Version8 struct {

	// ThreatType - no documentation
	ThreatType string `json:"threatType,omitempty"`

	// DetectedUtc - no documentation
	DetectedUtc *time.Time `json:"detectedUtc,omitempty"`

	// ThreatName - no documentation
	ThreatName string `json:"threatName,omitempty"`

	// ThreatActionTaken - no documentation
	ThreatActionTaken string `json:"threatActionTaken,omitempty"`

	// ThreatSeverityLabel - no documentation
	ThreatSeverityLabel string `json:"threatSeverityLabel,omitempty"`

	// SourceProcessName - no documentation
	SourceProcessName string `json:"sourceProcessName,omitempty"`

	// SourceIpv4 - The IP address of the source that generated an event.
	SourceIpv4 string `json:"sourceIpv4,omitempty"`

	// TargetFilename - The name of the file that was the target of the event.
	TargetFilename string `json:"targetFilename,omitempty"`

	// Signature - no documentation
	Signature *McAfee_Epolicy_Orchestrator_Version45_Hips_Event_Signature_Version8 `json:"signature,omitempty"`

	// VirusActionTaken - no documentation
	VirusActionTaken *McAfee_Epolicy_Orchestrator_Version45_Event_Filter_Description `json:"virusActionTaken,omitempty"`

	// AgentDetails - no documentation
	AgentDetails *McAfee_Epolicy_Orchestrator_Version45_Agent_Details `json:"agentDetails,omitempty"`
}

func (mcafee_epolicy_orchestrator_version45_event_version8 *McAfee_Epolicy_Orchestrator_Version45_Event_Version8) String() string {
	return "McAfee_Epolicy_Orchestrator_Version45_Event_Version8"
}
