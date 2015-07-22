package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Search_Result - The SoftLayer_Container_Search_Result data type represents a
// result row from an execution of Search service.
type SoftLayer_Container_Search_Result struct {

	// MatchedTerms - An array of terms that were matched in the resource object.
	MatchedTerms []string `json:"matchedTerms"`

	// RelevanceScore - The score ratio of the result for relevance to the search criteria.
	RelevanceScore float64 `json:"relevanceScore"`

	// Resource - A search results resource object that matched search criteria.
	Resource *SoftLayer_Entity `json:"resource"`

	// ResourceType - The type of the resource object that matched search criteria.
	ResourceType string `json:"resourceType"`
}
