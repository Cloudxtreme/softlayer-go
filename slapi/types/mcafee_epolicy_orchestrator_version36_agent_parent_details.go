package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// McAfee_Epolicy_Orchestrator_Version36_Agent_Parent_Details - The
// McAfee_Epolicy_Orchestrator_Version36_Agent_Parent_Details data type contains the name of an
// anti-virus policy.
type McAfee_Epolicy_Orchestrator_Version36_Agent_Parent_Details struct {

	// Name - no documentation
	Name string `json:"name,omitempty"`
}

func (mcafee_epolicy_orchestrator_version36_agent_parent_details *McAfee_Epolicy_Orchestrator_Version36_Agent_Parent_Details) String() string {
	return "McAfee_Epolicy_Orchestrator_Version36_Agent_Parent_Details"
}

// McAfee_Epolicy_Orchestrator_Version36_Agent_Parent_Details_Extended is McAfee_Epolicy_Orchestrator_Version36_Agent_Parent_Details with all maskable types.
type McAfee_Epolicy_Orchestrator_Version36_Agent_Parent_Details_Extended struct {
	McAfee_Epolicy_Orchestrator_Version36_Agent_Parent_Details

	// CurrentPolicy - no documentation
	CurrentPolicy *McAfee_Epolicy_Orchestrator_Version36_Agent_Parent_Details `json:"currentPolicy,omitempty"`
}

func (mcafee_epolicy_orchestrator_version36_agent_parent_details *McAfee_Epolicy_Orchestrator_Version36_Agent_Parent_Details_Extended) String() string {
	return "McAfee_Epolicy_Orchestrator_Version36_Agent_Parent_Details"
}
