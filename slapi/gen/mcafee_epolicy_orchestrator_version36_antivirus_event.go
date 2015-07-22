package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event - The
// McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event data type represents a single anti-virus
// event. It contains details about the event such as the date the event occurred, the virus that is
// detected and the action that is taken.
type McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event struct {

	// EventLocalDateTime - no documentation
	EventLocalDateTime *time.Time `json:"eventLocalDateTime"`

	// Filename - no documentation
	Filename string `json:"filename"`

	// VirusActionTaken - no documentation
	VirusActionTaken *McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event_Filter_Description `json:"virusActionTaken"`

	// VirusName - no documentation
	VirusName string `json:"virusName"`

	// VirusType - no documentation
	VirusType string `json:"virusType"`
}
