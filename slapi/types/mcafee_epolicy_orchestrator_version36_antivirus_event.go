package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event - The
// McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event data type represents a single anti-virus
// event. It contains details about the event such as the date the event occurred, the virus that is
// detected and the action that is taken.
type McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event struct {

	// Filename - no documentation
	Filename string `json:"filename"`

	// VirusName - no documentation
	VirusName string `json:"virusName"`

	// VirusType - no documentation
	VirusType string `json:"virusType"`

	// EventLocalDateTime - no documentation
	EventLocalDateTime *time.Time `json:"eventLocalDateTime"`
}

// McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event_Extended is McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event with all maskable types.
type McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event_Extended struct {
	McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event

	// VirusActionTaken - no documentation
	VirusActionTaken *McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event_Filter_Description `json:"virusActionTaken"`
}

func (mcafee_epolicy_orchestrator_version36_antivirus_event *McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event) String() string {
	return "McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event"
}
