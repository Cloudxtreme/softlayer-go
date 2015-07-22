package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Tag - The SoftLayer_Tag data type is an optional type associated with hardware. The
// account ID that the tag is tied to, and the tag itself are stored in this data type. There is also a
// flag to denote whether the tag is internal or not.
type SoftLayer_Tag struct {

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// AccountId - no documentation
	AccountId int `json:"accountId"`

	// Id - no documentation
	Id int `json:"id"`

	// Internal - no documentation
	Internal int `json:"internal"`

	// Name - Name of the tag. The characters permitted are 0-9, whitespace, _ (underscore),
	Name string `json:"name"`

	// ReferenceCount - no documentation
	ReferenceCount uint64 `json:"referenceCount"`

	// References - no documentation
	References []*SoftLayer_Tag_Reference `json:"references"`
}

// AutoComplete - This function is responsible for setting the Tags values. The internal flag is set to
// 0 if the user is a customer, and 1 otherwise. AccountId is set to the account bound to the user, and
// the tags name is set to the clean version of the tag inputted by the user.
func (softlayer_tag *SoftLayer_Tag) AutoComplete(commonOptions *slapi.CommonOptions, tag string) ([]*SoftLayer_Tag, error) {
	var returnValue []*SoftLayer_Tag
	return returnValue, nil
}

// GetAllTagTypes - no documentation
func (softlayer_tag *SoftLayer_Tag) GetAllTagTypes(commonOptions *slapi.CommonOptions) ([]*SoftLayer_Tag_Type, error) {
	var returnValue []*SoftLayer_Tag_Type
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_tag *SoftLayer_Tag) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Tag, error) {
	var returnValue *SoftLayer_Tag
	return returnValue, nil
}

// GetTagByTagName - Returns the Tag object with a given name. The user types in the tag name and this
// method returns the tag with that name.
func (softlayer_tag *SoftLayer_Tag) GetTagByTagName(commonOptions *slapi.CommonOptions, tagList string) ([]*SoftLayer_Tag, error) {
	var returnValue []*SoftLayer_Tag
	return returnValue, nil
}

// SetTags - Tag an object by passing in one or more tags separated by a comma. Tag references are
// cleared out every time this method is called. If your object is already tagged you will need to pass
// the current tags along with any new ones. To remove all tag references pass an empty string. To
// remove one or more tags omit them from the tag list. The characters permitted are 0-9, whitespace, _
// (underscore), - (hypen), . (period), and : (colon). All other characters will be stripped away.
func (softlayer_tag *SoftLayer_Tag) SetTags(commonOptions *slapi.CommonOptions, tags string, keyName string, resourceTableId int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}
