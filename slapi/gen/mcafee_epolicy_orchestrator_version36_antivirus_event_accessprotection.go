package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event_AccessProtection - The
// McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event_AccessProtection data type represents an
// access protection event. It contains details about the event such as when it occurs, the process
// that caused it, and the rule that triggered the event.
type McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event_AccessProtection struct {

	// EventLocalDateTime - no documentation
	EventLocalDateTime *time.Time `json:"eventLocalDateTime"`

	// Filename - The name of the file that was protected from access.
	Filename string `json:"filename"`

	// ProcessName - The name of the process that was protected from access.
	ProcessName string `json:"processName"`

	// RuleName - The name of the rule that triggered an access protection event.
	RuleName string `json:"ruleName"`

	// Source - The IP address that caused an access protection event.
	Source string `json:"source"`
}