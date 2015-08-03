package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// McAfee_Epolicy_Orchestrator_Version45_Event - The McAfee_Epolicy_Orchestrator_Version45_Event data
// type represents a single event. It contains details about the event such as the date the event
// occurred, the virus or intrusion that is detected and the action that is taken.
type McAfee_Epolicy_Orchestrator_Version45_Event struct {

	// ThreatName - no documentation
	ThreatName string `json:"threatName"`

	// ThreatSeverityLabel - no documentation
	ThreatSeverityLabel string `json:"threatSeverityLabel"`

	// DetectedUtc - no documentation
	DetectedUtc *time.Time `json:"detectedUtc"`

	// SourceIpv4 - The IP address of the source that generated an event.
	SourceIpv4 string `json:"sourceIpv4"`

	// SourceProcessName - no documentation
	SourceProcessName string `json:"sourceProcessName"`

	// TargetFilename - The name of the file that was the target of the event.
	TargetFilename string `json:"targetFilename"`

	// ThreatActionTaken - no documentation
	ThreatActionTaken string `json:"threatActionTaken"`

	// ThreatType - no documentation
	ThreatType string `json:"threatType"`
}

func (mcafee_epolicy_orchestrator_version45_event *McAfee_Epolicy_Orchestrator_Version45_Event) String() string {
	return "McAfee_Epolicy_Orchestrator_Version45_Event"
}

// McAfee_Epolicy_Orchestrator_Version45_Event_Extended is McAfee_Epolicy_Orchestrator_Version45_Event with all maskable types.
type McAfee_Epolicy_Orchestrator_Version45_Event_Extended struct {
	McAfee_Epolicy_Orchestrator_Version45_Event

	// AgentDetails - no documentation
	AgentDetails *McAfee_Epolicy_Orchestrator_Version45_Agent_Details `json:"agentDetails"`

	// VirusActionTaken - no documentation
	VirusActionTaken *McAfee_Epolicy_Orchestrator_Version45_Event_Filter_Description `json:"virusActionTaken"`
}

func (mcafee_epolicy_orchestrator_version45_event *McAfee_Epolicy_Orchestrator_Version45_Event_Extended) String() string {
	return "McAfee_Epolicy_Orchestrator_Version45_Event"
}
