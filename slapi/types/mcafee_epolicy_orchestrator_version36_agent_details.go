package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// McAfee_Epolicy_Orchestrator_Version36_Agent_Details - The
// McAfee_Epolicy_Orchestrator_Version36_Agent_Details data type represents a virus scan agent and
// contains details about its version.
type McAfee_Epolicy_Orchestrator_Version36_Agent_Details struct {

	// AgentVersion - no documentation
	AgentVersion string `json:"agentVersion"`

	// LastUpdate - The date of the last time the anti-virus agent checked in.
	LastUpdate string `json:"lastUpdate"`
}

func (mcafee_epolicy_orchestrator_version36_agent_details *McAfee_Epolicy_Orchestrator_Version36_Agent_Details) String() string {
	return "McAfee_Epolicy_Orchestrator_Version36_Agent_Details"
}

// McAfee_Epolicy_Orchestrator_Version36_Agent_Details_Extended is McAfee_Epolicy_Orchestrator_Version36_Agent_Details with all maskable types.
type McAfee_Epolicy_Orchestrator_Version36_Agent_Details_Extended struct {
	McAfee_Epolicy_Orchestrator_Version36_Agent_Details

	// CurrentPolicy - no documentation
	CurrentPolicy *McAfee_Epolicy_Orchestrator_Version36_Agent_Parent_Details `json:"currentPolicy"`
}

func (mcafee_epolicy_orchestrator_version36_agent_details *McAfee_Epolicy_Orchestrator_Version36_Agent_Details_Extended) String() string {
	return "McAfee_Epolicy_Orchestrator_Version36_Agent_Details"
}
