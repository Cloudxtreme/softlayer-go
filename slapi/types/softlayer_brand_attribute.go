package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Brand_Attribute - <nil>
type SoftLayer_Brand_Attribute struct {
}

// SoftLayer_Brand_Attribute_Extended is SoftLayer_Brand_Attribute with all maskable types.
type SoftLayer_Brand_Attribute_Extended struct {
	SoftLayer_Brand_Attribute

	// Brand - <nil>
	Brand *SoftLayer_Brand `json:"brand"`
}

func (softlayer_brand_attribute *SoftLayer_Brand_Attribute) String() string {
	return "SoftLayer_Brand_Attribute"
}
