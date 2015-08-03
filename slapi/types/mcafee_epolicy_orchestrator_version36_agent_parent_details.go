package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// McAfee_Epolicy_Orchestrator_Version36_Agent_Parent_Details - The
// McAfee_Epolicy_Orchestrator_Version36_Agent_Parent_Details data type contains the name of an
// anti-virus policy.
type McAfee_Epolicy_Orchestrator_Version36_Agent_Parent_Details struct {

	// CurrentPolicy - no documentation
	CurrentPolicy *McAfee_Epolicy_Orchestrator_Version36_Agent_Parent_Details `json:"currentPolicy"`

	// Name - no documentation
	Name string `json:"name"`
}

func (mcafee_epolicy_orchestrator_version36_agent_parent_details *McAfee_Epolicy_Orchestrator_Version36_Agent_Parent_Details) String() string {
	return "McAfee_Epolicy_Orchestrator_Version36_Agent_Parent_Details"
}
