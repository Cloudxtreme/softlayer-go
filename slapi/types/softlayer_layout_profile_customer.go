package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Layout_Profile_Customer - <nil>
type SoftLayer_Layout_Profile_Customer struct {
}

// SoftLayer_Layout_Profile_Customer_Extended is SoftLayer_Layout_Profile_Customer with all maskable types.
type SoftLayer_Layout_Profile_Customer_Extended struct {
	SoftLayer_Layout_Profile_Customer

	// UserRecord - <nil>
	UserRecord *SoftLayer_User_Customer `json:"userRecord"`
}

func (softlayer_layout_profile_customer *SoftLayer_Layout_Profile_Customer) String() string {
	return "SoftLayer_Layout_Profile_Customer"
}
