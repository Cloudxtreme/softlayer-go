package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Container_Search_Result - The SoftLayer_Container_Search_Result data type represents a
// result row from an execution of Search service.
type SoftLayer_Container_Search_Result struct {

	// MatchedTerms - An array of terms that were matched in the resource object.
	MatchedTerms []string `json:"matchedTerms,omitempty"`

	// RelevanceScore - The score ratio of the result for relevance to the search criteria.
	RelevanceScore slapi.Float64 `json:"relevanceScore,omitempty"`

	// Resource - A search results resource object that matched search criteria.
	Resource *SoftLayer_Entity `json:"resource,omitempty"`

	// ResourceType - The type of the resource object that matched search criteria.
	ResourceType string `json:"resourceType,omitempty"`
}

func (softlayer_container_search_result *SoftLayer_Container_Search_Result) String() string {
	return "SoftLayer_Container_Search_Result"
}
