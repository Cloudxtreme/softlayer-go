package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Tag - The SoftLayer_Tag data type is an optional type associated with hardware. The
// account ID that the tag is tied to, and the tag itself are stored in this data type. There is also a
// flag to denote whether the tag is internal or not.
type SoftLayer_Tag struct {

	// AccountId - no documentation
	AccountId int `json:"accountId"`

	// Id - no documentation
	Id int `json:"id"`

	// Internal - no documentation
	Internal int `json:"internal"`

	// Name - Name of the tag. The characters permitted are 0-9, whitespace, _ (underscore),
	Name string `json:"name"`
}

// SoftLayer_Tag_Extended is SoftLayer_Tag with all maskable types.
type SoftLayer_Tag_Extended struct {
	SoftLayer_Tag

	// ReferenceCount - no documentation
	ReferenceCount uint64 `json:"referenceCount"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// References - no documentation
	References []*SoftLayer_Tag_Reference `json:"references"`
}

func (softlayer_tag *SoftLayer_Tag) String() string {
	return "SoftLayer_Tag"
}
