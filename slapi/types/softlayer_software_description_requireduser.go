package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Software_Description_RequiredUser - This class represents a software description's
// required user
type SoftLayer_Software_Description_RequiredUser struct {

	// DefaultPassword - If the default password is set the user will be created with that password,
	// otherwise a random password is generated.
	DefaultPassword string `json:"defaultPassword"`

	// Username - If this software has a required user (such as "root") this string contains it's name.
	Username string `json:"username"`
}

func (softlayer_software_description_requireduser *SoftLayer_Software_Description_RequiredUser) String() string {
	return "SoftLayer_Software_Description_RequiredUser"
}
