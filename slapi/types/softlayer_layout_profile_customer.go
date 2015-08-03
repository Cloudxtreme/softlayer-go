package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Layout_Profile_Customer - <nil>
type SoftLayer_Layout_Profile_Customer struct {

	// UserRecord - <nil>
	UserRecord *SoftLayer_User_Customer `json:"userRecord"`
}

func (softlayer_layout_profile_customer *SoftLayer_Layout_Profile_Customer) String() string {
	return "SoftLayer_Layout_Profile_Customer"
}

// GetObject - <nil>
func (softlayer_layout_profile_customer *SoftLayer_Layout_Profile_Customer) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Layout_Profile_Customer, error) {
	var returnValue *SoftLayer_Layout_Profile_Customer
	return returnValue, nil
}
