package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Search_ObjectType_Property - This data type is a container that stores
// information about a single property of a searchable object type. Each
// [[SoftLayer_Container_Search_ObjectType (type)|SoftLayer_Container_Search_ObjectType]] object holds
// a collection of these properties. Property information can be used for discovery of searchable data
// and for the creation or validation of object index search strings. Note that properties are only
// understood by the [[SoftLayer_Search/advancedSearch|advancedSearch()]] method. Refer to the
// advancedSearch() method for information on using properties in search strings.
type SoftLayer_Container_Search_ObjectType_Property struct {

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// SortableFlag - no documentation
	SortableFlag bool `json:"sortableFlag,omitempty"`

	// Type - Property data type. Valid values include 'boolean', 'integer', 'date', 'string' or 'text'.
	Type string `json:"type,omitempty"`
}

func (softlayer_container_search_objecttype_property *SoftLayer_Container_Search_ObjectType_Property) String() string {
	return "SoftLayer_Container_Search_ObjectType_Property"
}
