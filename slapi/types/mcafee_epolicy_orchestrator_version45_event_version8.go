package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// McAfee_Epolicy_Orchestrator_Version45_Event_Version8 - The
// McAfee_Epolicy_Orchestrator_Version45_Event_Version8 data type represents a single event. It
// contains details about the event such as the date the event occurred, the virus or intrusion that is
// detected and the action that is taken.
type McAfee_Epolicy_Orchestrator_Version45_Event_Version8 struct {
}

func (mcafee_epolicy_orchestrator_version45_event_version8 *McAfee_Epolicy_Orchestrator_Version45_Event_Version8) String() string {
	return "McAfee_Epolicy_Orchestrator_Version45_Event_Version8"
}

// McAfee_Epolicy_Orchestrator_Version45_Event_Version8_Extended is McAfee_Epolicy_Orchestrator_Version45_Event_Version8 with all maskable types.
type McAfee_Epolicy_Orchestrator_Version45_Event_Version8_Extended struct {
	McAfee_Epolicy_Orchestrator_Version45_Event_Version8

	// Signature - no documentation
	Signature *McAfee_Epolicy_Orchestrator_Version45_Hips_Event_Signature_Version8 `json:"signature"`
}

func (mcafee_epolicy_orchestrator_version45_event_version8 *McAfee_Epolicy_Orchestrator_Version45_Event_Version8_Extended) String() string {
	return "McAfee_Epolicy_Orchestrator_Version45_Event_Version8"
}
