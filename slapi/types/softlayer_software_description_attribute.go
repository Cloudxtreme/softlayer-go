package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Software_Description_Attribute - The SoftLayer_Software_Description_Attribute data type
// represents an attributes associated with this software description.
type SoftLayer_Software_Description_Attribute struct {

	// Value - no documentation
	Value string `json:"value"`
}

func (softlayer_software_description_attribute *SoftLayer_Software_Description_Attribute) String() string {
	return "SoftLayer_Software_Description_Attribute"
}

// SoftLayer_Software_Description_Attribute_Extended is SoftLayer_Software_Description_Attribute with all maskable types.
type SoftLayer_Software_Description_Attribute_Extended struct {
	SoftLayer_Software_Description_Attribute

	// SoftwareDescription - <nil>
	SoftwareDescription *SoftLayer_Software_Description `json:"softwareDescription"`

	// Type - <nil>
	Type *SoftLayer_Software_Description_Attribute_Type `json:"type"`
}

func (softlayer_software_description_attribute *SoftLayer_Software_Description_Attribute_Extended) String() string {
	return "SoftLayer_Software_Description_Attribute"
}
