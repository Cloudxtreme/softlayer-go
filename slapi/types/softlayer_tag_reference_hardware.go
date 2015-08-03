package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Tag_Reference_Hardware - <nil>
type SoftLayer_Tag_Reference_Hardware struct {
}

// SoftLayer_Tag_Reference_Hardware_Extended is SoftLayer_Tag_Reference_Hardware with all maskable types.
type SoftLayer_Tag_Reference_Hardware_Extended struct {
	SoftLayer_Tag_Reference_Hardware

	// Resource - <nil>
	Resource *SoftLayer_Hardware `json:"resource"`
}

func (softlayer_tag_reference_hardware *SoftLayer_Tag_Reference_Hardware) String() string {
	return "SoftLayer_Tag_Reference_Hardware"
}
