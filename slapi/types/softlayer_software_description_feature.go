package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Software_Description_Feature - The SoftLayer_Software_Description_Feature data type
// represents a single software description feature. A feature may show up on more than one software
// description and can not be created, modified, or removed.
type SoftLayer_Software_Description_Feature struct {

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// Vendor - The vendor that a software description feature belongs to.
	Vendor string `json:"vendor,omitempty"`

	// Id - The unique identifier for a software description feature.
	Id int `json:"id,omitempty"`

	// KeyName - A unique name used to reference this software description feature.
	KeyName string `json:"keyName,omitempty"`
}

func (softlayer_software_description_feature *SoftLayer_Software_Description_Feature) String() string {
	return "SoftLayer_Software_Description_Feature"
}
