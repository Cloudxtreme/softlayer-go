package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// McAfee_Epolicy_Orchestrator_Version45_Agent_Parent_Details - The
// McAfee_Epolicy_Orchestrator_Version45_Agent_Parent_Details data type contains the name of an
// anti-virus policy.
type McAfee_Epolicy_Orchestrator_Version45_Agent_Parent_Details struct {

	// Name - no documentation
	Name string `json:"name,omitempty"`
}

func (mcafee_epolicy_orchestrator_version45_agent_parent_details *McAfee_Epolicy_Orchestrator_Version45_Agent_Parent_Details) String() string {
	return "McAfee_Epolicy_Orchestrator_Version45_Agent_Parent_Details"
}

// McAfee_Epolicy_Orchestrator_Version45_Agent_Parent_Details_Extended is McAfee_Epolicy_Orchestrator_Version45_Agent_Parent_Details with all maskable types.
type McAfee_Epolicy_Orchestrator_Version45_Agent_Parent_Details_Extended struct {
	McAfee_Epolicy_Orchestrator_Version45_Agent_Parent_Details

	// PolicyCount - A count of the current anti-virus policy of an agent.
	PolicyCount uint64 `json:"policyCount,omitempty"`

	// AgentDetails - no documentation
	AgentDetails *McAfee_Epolicy_Orchestrator_Version45_Agent_Details `json:"agentDetails,omitempty"`

	// Policies - no documentation
	Policies []*McAfee_Epolicy_Orchestrator_Version45_Agent_Parent_Details `json:"policies,omitempty"`
}

func (mcafee_epolicy_orchestrator_version45_agent_parent_details *McAfee_Epolicy_Orchestrator_Version45_Agent_Parent_Details_Extended) String() string {
	return "McAfee_Epolicy_Orchestrator_Version45_Agent_Parent_Details"
}
