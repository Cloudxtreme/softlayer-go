package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_BlockedApplicationEvent - The
// McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_BlockedApplicationEvent data type contains a
// single blocked application event. The details of the event are the time the event occurred, the
// process that generated the event and a brief description of the application that was blocked.
type McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_BlockedApplicationEvent struct {

	// ApplicationDescription - A brief description of the application that is blocked.
	ApplicationDescription string `json:"applicationDescription"`

	// IncidentTime - no documentation
	IncidentTime *time.Time `json:"incidentTime"`

	// ProcessName - no documentation
	ProcessName string `json:"processName"`
}
