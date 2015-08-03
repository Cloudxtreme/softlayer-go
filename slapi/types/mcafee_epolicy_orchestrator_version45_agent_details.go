package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// McAfee_Epolicy_Orchestrator_Version45_Agent_Details - The
// McAfee_Epolicy_Orchestrator_Version45_Agent_Details data type represents a virus scan agent and
// contains details about its version.
type McAfee_Epolicy_Orchestrator_Version45_Agent_Details struct {

	// AgentVersion - no documentation
	AgentVersion string `json:"agentVersion"`

	// LastUpdate - The date of the last time the anti-virus agent checked in.
	LastUpdate *time.Time `json:"lastUpdate"`
}

func (mcafee_epolicy_orchestrator_version45_agent_details *McAfee_Epolicy_Orchestrator_Version45_Agent_Details) String() string {
	return "McAfee_Epolicy_Orchestrator_Version45_Agent_Details"
}
