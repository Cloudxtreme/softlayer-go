package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Search_ObjectType - This data type is a container that stores information about
// a single indexed object type. Object type information can be used for discovery of searchable data
// and for creation or validation of object index search strings. Each of these containers holds a
// collection of [[SoftLayer_Container_Search_ObjectType_Property
// (type)|SoftLayer_Container_Search_ObjectType_Property]] objects, specifying which object properties
// are exposed for the current user. Refer to the the documentation for the
// [[SoftLayer_Search/search|search()]] method for information on using object types in search strings.
type SoftLayer_Container_Search_ObjectType struct {

	// Name - no documentation
	Name string `json:"name"`

	// Properties - A collection of [[SoftLayer_Container_Search_ObjectType_Property|object properties]].
	Properties []*SoftLayer_Container_Search_ObjectType_Property `json:"properties"`
}
