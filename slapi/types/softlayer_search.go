package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Search - <nil>
type SoftLayer_Search struct {
}

// AdvancedSearch - This method allows for searching for SoftLayer resources by simple terms and
// operators. Fields that are used for searching will be available at sldn.softlayer.com. It returns a
// collection or array of [[SoftLayer_Container_Search_Result
// (type)|SoftLayer_Container_Search_Result]] objects that have search metadata for each result and the
// resulting resource found. The advancedSearch() method recognizes the special _objectType: quantifier
// in search strings. See the documentation for the [[SoftLayer_Search/search|search()]] method on how
// to restrict searches using object types. The advancedSearch() method recognizes
// [[SoftLayer_Container_Search_ObjectType_Property (type)|object properties]] , which can also be used
// to limit searches. Example: _objectType:Type_1 propertyA: value A search string can specify multiple
// properties, separated with spaces. Example: _objectType:Type_1 propertyA: value propertyB: value A
// collection of available object types and their properties can be retrieved by calling the
// [[SoftLayer_Search/getObjectTypes|getObjectTypes()]] method.
func (softlayer_search *SoftLayer_Search) AdvancedSearch(commonOptions *slapi.CommonOptions, searchString string) ([]*SoftLayer_Container_Search_Result, error) {
	var returnValue []*SoftLayer_Container_Search_Result
	return returnValue, nil
}

// GetObjectTypes - This method returns a collection of [[SoftLayer_Container_Search_ObjectType
// (type)|SoftLayer_Container_Search_ObjectType]] containers that specify which indexed object types
// and properties are exposed for the current user. These object types can be used to discover
// searchable data and to create or validate object index search strings. Refer to the
// [[SoftLayer_Search/search|search()]] and [[SoftLayer_Search/advancedSearch|advancedSearch()]]
// methods for information on using object types and properties in search strings.
func (softlayer_search *SoftLayer_Search) GetObjectTypes(commonOptions *slapi.CommonOptions) ([]*SoftLayer_Container_Search_ObjectType, error) {
	var returnValue []*SoftLayer_Container_Search_ObjectType
	return returnValue, nil
}

// Search - This method allows for searching for SoftLayer resources by simple phrase. It returns a
// collection or array of [[SoftLayer_Container_Search_Result
// (type)|SoftLayer_Container_Search_Result]] objects that have search metadata for each result and the
// resulting resource found. This method recognizes the special _objectType: quantifier in search
// strings. This quantifier can be used to restrict a search to specific object types. Example usage:
// _objectType:Type_1 (other search terms...) A search string can specify multiple object types,
// separated by commas (no spaces are permitted between the type names). Example:
// _objectType:Type_1,Type_2,Type_3 (other search terms...) If the list of object types is prefixed
// with a hyphen or minus sign then the specified types are excluded from the search. Example:
// _objectType:-Type_4,Type_5 (other search terms...) A collection of available object types can be
// retrieved by calling the [[SoftLayer_Search/getObjectTypes|getObjectTypes()]] method.
func (softlayer_search *SoftLayer_Search) Search(commonOptions *slapi.CommonOptions, searchString string) ([]*SoftLayer_Container_Search_Result, error) {
	var returnValue []*SoftLayer_Container_Search_Result
	return returnValue, nil
}
