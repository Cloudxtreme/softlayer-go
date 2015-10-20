package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// McAfee_Epolicy_Orchestrator_Version45_Agent_Parent_Details - The
// McAfee_Epolicy_Orchestrator_Version45_Agent_Parent_Details data type contains the name of an
// anti-virus policy.
type McAfee_Epolicy_Orchestrator_Version45_Agent_Parent_Details struct {

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// AgentDetails - no documentation
	AgentDetails *McAfee_Epolicy_Orchestrator_Version45_Agent_Details `json:"agentDetails,omitempty"`

	// Policies - no documentation
	Policies []*McAfee_Epolicy_Orchestrator_Version45_Agent_Parent_Details `json:"policies,omitempty"`

	// PolicyCount - A count of the current anti-virus policy of an agent.
	PolicyCount uint64 `json:"policyCount,omitempty"`
}

func (mcafee_epolicy_orchestrator_version45_agent_parent_details *McAfee_Epolicy_Orchestrator_Version45_Agent_Parent_Details) String() string {
	return "McAfee_Epolicy_Orchestrator_Version45_Agent_Parent_Details"
}
